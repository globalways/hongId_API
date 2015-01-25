package main

import (
	_ "hongID/docs"
	_ "hongID/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"os"
	"runtime"
	_ "hongID/hprose"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	beego.Run()
}

func init() {

	// log
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(3)

	switch beego.RunMode {
	case "dev":
		// api document
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"

		corsFunc := func(ctx *context.Context) {
			ctx.Output.Header("Access-Control-Allow-Origin", "*")
		}

		beego.InsertFilter("*", beego.BeforeRouter, corsFunc)
	case "prod":
		// log
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}
}
