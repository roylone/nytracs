package router

import (
	myws "com.yict/websocket"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type wsRouter struct{}

func (_ *wsRouter) createHandle(art *mux.Router) {
	art.HandleFunc("/ws", bulidWsClient)
}

// bulidWsClient handles websocket requests from the peer.
func bulidWsClient(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	log.Println(conn.LocalAddr(), conn.RemoteAddr())
	if err != nil {
		log.Println(err)
		return
	}
	ip := strings.Split(conn.RemoteAddr().String(), ":")[0]
	inMsg, outMsg := make(chan []byte, 1024), make(chan []byte, 1024)
	client := myws.NewClient(conn, inMsg, outMsg, ip, "")
	myws.GetWsADT().GetRegister() <- client
	go client.SendMsg()
	client.ReceiveMsg()
}
