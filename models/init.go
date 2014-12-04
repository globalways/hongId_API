// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// page
type Page struct {
	Size    int64 // per page count
	CurPage int64 // current page
}

//数据库连接
var (
	Reader  orm.Ormer //读数据
	Writter orm.Ormer //写数据
)

// 参数读取
var (
//	dbDriver = beego.AppConfig.String(beego.RunMode + "::dbdriver")
//	userName = beego.AppConfig.String(beego.RunMode + "::dbuser")
//	userPass = beego.AppConfig.String(beego.RunMode + "::dbpass")
//	dbHost   = beego.AppConfig.String(beego.RunMode + "::dbhost")
//	dbName   = beego.AppConfig.String(beego.RunMode + "::dbname")
//	dbEncode = beego.AppConfig.String(beego.RunMode + "::dbencode")
	dbDriver = "mysql"
	userName = "root"
	userPass = "bigbang990"
	dbHost   = "127.0.0.1:3306"
	dbName   = "gws_hongid"
	dbEncode = "utf8"
)

func init() {

	// 注册模型
	orm.RegisterModelWithPrefix("", new(ChannelType), new(MemberCard), new(Member), new(MemberGroup), new(MemberGrowth))

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 注册数据库
	orm.RegisterDataBase("default", dbDriver,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", userName, userPass, dbHost, dbName, dbEncode), 30, 30)

	if beego.RunMode == "dev" {
		// 调试模式
		orm.Debug = true
		// 自动建表
		orm.RunSyncdb("default", true, true)
	} else {
		// 调试模式
		orm.Debug = false
		// 自动建表
		orm.RunSyncdb("default", false, true)
	}

	// 注册数据库连接
	Reader = orm.NewOrm()
	Writter = orm.NewOrm()
}
