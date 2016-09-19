package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type tmtRouter struct{}

func (_ *tmtRouter) createHandle(art *mux.Router) {
	art.HandleFunc("/tmt", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "test tmt func")
	})
}
