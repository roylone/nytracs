// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is an middleman between the websocket connection and the hub.
type Client struct {
	// The websocket connection.
	wsConn *websocket.Conn

	// Buffered channel of outbound messages.
	InMsg     chan []byte
	OutMsg    chan []byte
	clientIp  string
	clientTid string
}

func NewClient(conn *websocket.Conn, inMsg, outMsg chan []byte, ip, tid string) *Client {
	return &Client{
		wsConn:    conn,
		InMsg:     inMsg,
		OutMsg:    outMsg,
		clientIp:  ip,
		clientTid: tid,
	}
}

// receiveMsg pumps messages from the websocket connection to the hub.
func (c *Client) ReceiveMsg() {
	defer func() {
		GetWsADT().unregister <- c
		c.wsConn.Close()
	}()
	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(pongWait))
	c.wsConn.SetPongHandler(func(string) error { c.wsConn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.wsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("websocket error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		log.Println("receive: " + string(message))
		GetWsADT().broadcast <- message
		//		c.InMsg <- message
	}
}

// write writes a message with the given message type and payload.
func (c *Client) write(mt int, payload []byte) error {
	c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.wsConn.WriteMessage(mt, payload)
}

//sendMsg send messages from the hub to the websocket connection.
func (c *Client) SendMsg() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		GetWsADT().unregister <- c
		ticker.Stop()
		c.wsConn.Close()
	}()
	for {
		select {
		case message, ok := <-c.OutMsg:
			if !ok {
				// The hub closed the channel.
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.wsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.OutMsg)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.OutMsg)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
