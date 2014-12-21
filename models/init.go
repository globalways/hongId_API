// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/globalways/hongId_models/models"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//数据库连接
var (
	Reader  orm.Ormer //读数据
	Writter orm.Ormer //写数据
)

// 参数读取
var (
	dbDriver = beego.AppConfig.DefaultString(beego.RunMode+"::dbdriver", "mysql")
	userName = beego.AppConfig.DefaultString(beego.RunMode+"::dbuser", "root")
	userPass = beego.AppConfig.DefaultString(beego.RunMode+"::dbpass", "bigbang990")
	dbHost   = beego.AppConfig.DefaultString(beego.RunMode+"::dbhost", "127.0.0.1:3306")
	dbName   = beego.AppConfig.DefaultString(beego.RunMode+"::dbname", "gws_hongid")
	dbEncode = beego.AppConfig.DefaultString(beego.RunMode+"::dbencode", "utf8")

//	dbDriver = "mysql"
//	userName = "root"
//	userPass = "bigbang990"
//	dbHost   = "127.0.0.1:3306"
//	dbName   = "gws_hongid"
//	dbEncode = "utf8"
)

func init() {

	// 注册模型
	orm.RegisterModelWithPrefix(
		"",
		new(MemberProfile),
		new(MemberCard),
		new(Member),
		new(MemberGroup),
		new(MemberGrowth),
		new(MemberScore),
	)

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 注册数据库
	orm.RegisterDataBase("default", dbDriver,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", userName, userPass, dbHost, dbName, dbEncode), 30, 30)

	if beego.RunMode == "dev" {
		orm.Debug = true
		orm.RunSyncdb("default", false, true)
	} else {
		orm.Debug = false
		orm.RunSyncdb("default", false, true)
	}

	// 注册数据库连接
	Reader = orm.NewOrm()
	Writter = orm.NewOrm()

	SyncData()
}

func SyncData() {
	if !IsGroupExist("环途会员", Reader) && !IsGroupExist("个人会员", Reader) && !IsGroupExist("机构会员", Reader) {
		orm.RunSyncdb("default", true, false)
		go initData()
	}
}

// 初始化系统数据
func initData() {
	// 分组：环途会员, 个人用户，机构用户
	gwsGroup := &MemberGroup{
		GroupName:    "环途会员",
		GroupDesc:    "环途会员分组",
		Contribution: 0,
		Status:       EMemberGroupStatus_Enable,
	}
	gwsId, _ := NewMemberGroup(gwsGroup, Writter)

	userGroup := &MemberGroup{
		GroupName:    "个人会员",
		GroupDesc:    "个人会员分组",
		Contribution: 0,
		Status:       EMemberGroupStatus_Enable,
	}
	userId, _ := NewMemberGroup(userGroup, Writter)

	agencyGroup := &MemberGroup{
		GroupName:    "机构会员",
		GroupDesc:    "机构会员分组",
		Contribution: 100,
		Status:       EMemberGroupStatus_Enable,
	}
	agencyId, _ := NewMemberGroup(agencyGroup, Writter)

	// 各自生成1000个会员
	GenMembers(1, 9999, 1000, gwsId, Writter)
	GenMembers(10000000, 99999999, 1000, userId, Writter)
	GenMembers(100000, 999999, 1000, agencyId, Writter)

	// 生成环途会员卡发布会员
	gwsUser, _ := RegisterMemberByTel(1, "4000285843",Writter)

	// 环途自营会员卡10000张
	card := &ReqCard{
		MII: 6,
		CPI: 32,
		CDI: 86,
	}
	GenMemberCards(card, gwsUser.Id, 10000, Writter)
}
