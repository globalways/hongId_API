package main

import (
	_ "hongId/docs"
	_ "hongId/routers"

	"github.com/astaxie/beego"
	"os"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.Run()
}

func init() {

	// log
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(4)

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
