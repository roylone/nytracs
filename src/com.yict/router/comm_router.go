package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

//mux refer to https://github.com/gorilla/mux
type iRouter interface {
	createHandle(art *mux.Router)
}

type commRouter struct {
}

func (_ *commRouter) createHandle(art *mux.Router) {
	art.HandleFunc("/", home)
	art.HandleFunc("/hello/{category}", say)
}

func home(w http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{Directory: "src/com.yict/html", Extensions: []string{".tmpl", ".html"}})
	r.HTML(w, http.StatusOK, "home", nil)
}

func say(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]
	fmt.Fprintln(w, "hi golang,my name is "+category)
}
