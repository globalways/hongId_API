package main

import (
	_ "hongId/docs"
	_ "hongId/routers"

	"github.com/astaxie/beego"
	"os"
)

func main() {

	beego.EnableHttpListen = false
	beego.EnableHttpTLS = true

	beego.Run()
}

func init() {

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

	// log
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(4)
	if beego.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}
}
