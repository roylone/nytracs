package websocket

import (
	_ "log"
	_ "net/http"
	_ "strings"
)

//NewMsg notice it when other process send msg
func NewMsg(flag string, msg string) (ok bool) {
	if clt := GetWsADT().QueryClient(flag); clt != nil {
		clt.OutMsg <- []byte(msg)
		ok = true
	} else {
		ok = false
	}
	return ok
}

//Broadcast broadcast msg to all clients at list by ip or tid
func Broadcast(list []string, msg string) {
	for _, flag := range list {
		if c := GetWsADT().QueryClient(flag); c != nil {
			c.OutMsg <- []byte(msg)
		}
	}
}
