package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type fixRouter struct{}

func (_ *fixRouter) createHandle(art *mux.Router) {
	art.HandleFunc("/fix", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "test fix func")
	})
}
