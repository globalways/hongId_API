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

//数据库连接
var (
	Reader orm.Ormer //读数据

	Writter orm.Ormer //写数据
)

// 参数读取
var (
	dbDriver = beego.AppConfig.String("dbdriver")
	userName = beego.AppConfig.String("dbuser")
	userPass = beego.AppConfig.String("dbpass")
	dbHost   = beego.AppConfig.String("dbhost")
	dbName   = beego.AppConfig.String("dbname")
	dbEncode = beego.AppConfig.String("dbencode")
)

func init() {

	// 注册模型
	orm.RegisterModel(new(ChannelType), new(MemberCard))
	//orm.RegisterModelWithPrefix("prefix_", new(User))

	if len(dbDriver) == 0 {
		dbDriver = "mysql"
	}
	if len(userName) == 0 {
		userName = "root"
	}
	if len(userPass) == 0 {
		userPass = "bigbang990"
	}
	if len(dbHost) == 0 {
		dbHost = "127.0.0.1:3306"
	}
	if len(dbName) == 0 {
		dbName = "gw_memberCard"
	}
	if len(dbEncode) == 0 {
		dbEncode = "utf8"
	}

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
		orm.RunSyncdb("default", false, false)
	}

	// 注册数据库连接
	Reader = orm.NewOrm()
	Writter = orm.NewOrm()
}
