package main

import (
	"com.yict/router"
	"com.yict/util"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/secure"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	initJms()
	initMsi()

}

func main() {
	//refer to https://github.com/unrolled/secure
	secureMiddleware := negroni.HandlerFunc(secure.New(secure.Options{
		FrameDeny:     true,
		IsDevelopment: true,
	}).HandlerFuncWithNext)

	web := negroni.Classic()
	//	web := negroni.New()
	//
	//	web.Use(negroni.NewRecovery())
	//	web.Use(negronilogrus.NewMiddleware())
	//	web.Use(negroni.NewStatic(http.Dir("public")))

	web.Use(gzip.Gzip(gzip.DefaultCompression))
	web.Use(secureMiddleware)

	web.UseHandler(router.GetRouter())

	go func() {
		log.Fatal(http.ListenAndServe(util.HTTP_HOST, web))
	}()
	log.Fatal(http.ListenAndServeTLS(util.HTTPS_HOST, util.SERVER_CRT, util.SERVER_KEY, web))
}

func initMsi() {
	log.Println("Init msi message handler form nGen!")
}

func initJms() {
	log.Println("Init JMS message handler form nGen!")
}
