// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package websocket

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type WsADT struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

var wsAdt *WsADT = &WsADT{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[string]*Client),
}

func init() {
	go wsAdt.run()
}

func GetWsADT() *WsADT {
	return wsAdt
}

func (wd *WsADT) GetRegister() chan *Client {
	return wd.register
}

func (wd *WsADT) QueryClient(flag string) *Client {
	if client, ok := wd.clients[flag]; ok {
		return client
	}
	return nil
}

func (wd *WsADT) run() {
	for {
		select {
		case client := <-wd.register:
			wd.clients[client.clientIp] = client
		case client := <-wd.unregister:
			if _, ok := wd.clients[client.clientIp]; ok {
				delete(wd.clients, client.clientIp)
				close(client.InMsg)
				close(client.OutMsg)
			}
		case message := <-wd.broadcast:
			for k, v := range wd.clients {
				select {
				case v.OutMsg <- message:
				default:
					close(v.OutMsg)
					delete(wd.clients, k)
				}
			}
		}
	}
}
