// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	sm "github.com/globalways/chain_store_models/models"
	hm "github.com/globalways/hongId_models/models"
	pm "github.com/globalways/points_models/models"
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
)

func init() {

	// 注册模型
	orm.RegisterModelWithPrefix(
		"t_",
		new(hm.MemberProfile),
		new(hm.MemberCard),
		new(hm.Member),
		new(hm.MemberGroup),
		new(hm.MemberGrowth),
		new(pm.MemberPoints),
		new(pm.MemberPointsDetail),
		new(sm.Order),
		new(sm.OrderDetail),
		new(sm.OrderProcess),
		new(sm.OrderAddress),
		new(sm.PurchaseChannel),
		new(sm.PurchaseMaterial),
		new(sm.PurchaseProduct),
		new(sm.Product),
		new(sm.ProductTag),
		new(sm.Store),
		new(sm.StoreMaterial),
		new(sm.StoreProduct),
		new(sm.StoreAdmin),
		new(sm.StoreIndustry),
	)

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 注册数据库
	orm.RegisterDataBase("default", dbDriver,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", userName, userPass, dbHost, dbName, dbEncode), 100, 100)

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
	if !hm.IsGroupExist("环途会员组", Reader) && !hm.IsGroupExist("个人会员组", Reader) && !hm.IsGroupExist("机构会员组", Reader) {
		orm.RunSyncdb("default", true, false)
		go initData()
	}
}

// 初始化系统数据
func initData() {
	// 分组：环途会员, 个人用户，机构用户
	gws := hm.NewMemberGroup("环途会员组", "环途会员分组,只用于系统内部或者内部员工", 0, hm.EMemberGroupStatus_Enable, Writter)
	user := hm.NewMemberGroup("个人会员组", "个人会员分组,用户个人用户注册使用", 0, hm.EMemberGroupStatus_Enable, Writter)
	agency := hm.NewMemberGroup("机构会员组", "机构会员分组,用户企业用户注册使用", 0, hm.EMemberGroupStatus_Enable, Writter)

	// 各自生成1000个会员
	hm.GenMembers(1, 2, 1, gws.Id, Writter)
	hm.GenMembers(10000000, 99999999, 1000, user.Id, Writter)
	hm.GenMembers(100000, 999999, 1000, agency.Id, Writter)

	// 发卡会员，主要使用hongid
	cardAuth := hm.GetUnUsedMember(gws.Id, Reader)
	args := map[string]interface{}{
		"status": hm.EMemberStatus_Sys,
	}
	hm.UpdateMemberById(cardAuth.Id, args, Writter)

	// 环途自营会员卡10000张
	hm.GenMemberCards(6, 32, 86, cardAuth.HongId, 10000, Writter)

	//商铺行业分类添加
	sm.NewStoreIndustry("餐饮", "", Writter)
	sm.NewStoreIndustry("零售", "", Writter)

	//商品tag添加
	sm.NewProductTag("中餐", "", Writter)
	sm.NewProductTag("西餐", "", Writter)
	sm.NewProductTag("小吃", "", Writter)
	sm.NewProductTag("副食", "", Writter)
	sm.NewProductTag("熟食", "", Writter)
	sm.NewProductTag("香烟", "", Writter)
	sm.NewProductTag("洋酒", "", Writter)
}
