package router

import (
	"github.com/gorilla/mux"
)

var apps = []string{"common", "websocket", "hht", "fix", "vmt", "tmt"}

func GetRouter() *mux.Router {
	art := mux.NewRouter()
	addAppHandler(art)
	return art
}

//Add all handler when get router
func addAppHandler(art *mux.Router) {
	for _, a := range apps {
		factoryRoute(a).createHandle(art)
	}
}

//To create each router from factory method
func factoryRoute(app string) iRouter {
	var appRouter iRouter
	switch app {
	case "common":
		appRouter = new(commRouter)
	case "hht":
		appRouter = new(hhtRouter)
	case "fix":
		appRouter = new(fixRouter)
	case "vmt":
		appRouter = new(vmtRouter)
	case "tmt":
		appRouter = new(tmtRouter)
	case "websocket":
		appRouter = new(wsRouter)
	}
	return appRouter
}
