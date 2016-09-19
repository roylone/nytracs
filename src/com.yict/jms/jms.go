package jms

import (
	_ "com.yict/websocket"
)

//receive jms message and send msg to chan of match

func GetTerminal() (clients []string, msg string) {
	return []string{""}, ""
}

//invoked when jms message arrived
//func noticeSend() {
//	cs, msg := GetTerminal()
//	if ok := websocket.NewMsg(cs, msg); ok {
//		//send success,remove this msg from jms msg queue
//	}
//
//}

//...others func or type
