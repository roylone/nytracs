package util

import (
	"github.com/widuu/goini"
)

const (
	ENV string = "dev"
)

var CONF *goini.Config
var (
	HTTPS_HOST,
	HTTP_HOST,
	SERVER_CRT,
	SERVER_KEY string
)
var CONF_PATH = "src/com.yict/conf.ini"

//refer to https://github.com/widuu/goini
func init() {
	CONF = goini.SetConfig(CONF_PATH)
	HTTPS_HOST = CONF.GetValue(ENV, "https_host")
	HTTP_HOST = CONF.GetValue(ENV, "http_host")
	SERVER_CRT = CONF.GetValue(ENV, "SERVER_CRT")
	SERVER_KEY = CONF.GetValue(ENV, "SERVER_KEY")

}
