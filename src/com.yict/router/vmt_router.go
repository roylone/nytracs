package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type vmtRouter struct{}

func (_ *vmtRouter) createHandle(art *mux.Router) {
	art.HandleFunc("/vmt", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "test vmt func")
	})
}
