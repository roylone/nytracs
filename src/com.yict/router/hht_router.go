package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type hhtRouter struct{}

func (_ *hhtRouter) createHandle(art *mux.Router) {
	art.HandleFunc("/hht", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "test hht func")
	})
}
