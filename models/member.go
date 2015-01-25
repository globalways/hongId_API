// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/globalways/utils_go/random"
	"time"
)

const (
	ECalenDarType_Gregorian byte = iota + 1 //公历
	ECalenDarType_Lunar                     //农历
)

const (
	ESexType_Male byte = iota + 1 //男
	ESexType_Female
)

const (
	EEduLevel_None      byte = iota + 1 //文盲
	EEduLevel_Primary                   //小学
	EEduLevel_Secondary                 //中学
	EEduLevel_Senior                    //高中
	EEduLevel_Bachelor                  //本科
	EEduLevel_Master                    //研究生
	EEduLevel_Doctor                    //博士
)

const (
	EMemberStatus_NotUse byte = iota + 1 //初始，新生成 & 系统回收
	EMemberStatus_Liang                  //备用，系统生成后经过帐号过滤，被系统留着备用的，靓号..
	EMemberStatus_Use                    //正常，用户已经在使用
	EMemberStatus_Lock                   //异常，用户在使用，但被锁定
	EMemberStatus_Del                    //删除，已经被用户删除，但系统还未回收
	EMemberStatus_Sys                    //系统用户，不准删除
)

type Member struct {
	Id          int64
	HongId      int64     `orm:"column(hongid);unique"`
	Tel         string    `orm:"column(tel);null"`
	Email       string    `orm:"column(email);null"`
	PassWord    string    `orm:"column(password);null"`
	NickName    string    `orm:"column(nick_name);null;size(20)"`
	Avatar      string    `orm:"column(avatar);null"`
	GroupId     int64     `orm:"column(groupid)"`
	GrowthId    int64     `orm:"column(growthid)"`
	GrowthPoint uint      `orm:"column(growth_point);null"`
	Status      byte      `orm:"column(status)"`
	ProfileId   int64     `orm:"column(profileid)"`
	Created     time.Time `orm:"column(created);auto_now_add"`
	Updated     time.Time `orm:"column(updated);auto_now"`
}

func (m *Member) TableName() string {
	return "member1"
}

//批量生成会员
func GenMembers(minNo int64, maxNo int64, count int64, groupId int64) (affactedTotal int64) {
	affactedTotal = 0

	qs := ormer.QueryTable(new(Member))
	inserter, _ := qs.PrepareInsert()
	for i := int64(1); i <= count; i++ {
		randId := random.RandomInt64(minNo, maxNo)
		//如果存在当前hongid，放弃randId
		if HasHongId(randId) {
			continue
		}

		member := &Member{
			HongId:  randId,
			GroupId: groupId,
			Status:  EMemberStatus_NotUse,
		}

		//TODO 靓号逻辑判断

		_, err := inserter.Insert(member)
		if err != nil {
			continue
		}

		affactedTotal++
	}
	inserter.Close()

	return
}

// 返回一个未被使用的会员记录
func GetUnUsedMember(groupId int64, fields []string) *Member {
	member := new(Member)
	if err := ormer.QueryTable(member).Filter("status", EMemberStatus_NotUse).Filter("groupid", groupId).Limit(1).One(member, fields...); err != nil {
		return nil
	}

	return member
}

// 更新会员信息
func UpdateMemberById(memberId int64, args map[string]interface{}) bool {
	if _, err := ormer.QueryTable(new(Member)).Filter("id", memberId).Update(orm.Params(args)); err != nil {
		return false
	}

	return true
}

func UpdateMemberByHongId(hongid int64, args map[string]interface{}) bool {
	if _, err := ormer.QueryTable(new(Member)).Filter("hongid", hongid).Update(orm.Params(args)); err != nil {
		return false
	}

	return true
}

func UpdateMemberByTel(tel string, groupId int64, args map[string]interface{}) bool {
	if _, err := ormer.QueryTable(new(Member)).Filter("tel", tel).Filter("groupid", groupId).Update(orm.Params(args)); err != nil {
		return false
	}

	return true
}

// 查找会员信息(手机号)
func GetMemberByTel(tel string, groupId int64, fields []string) *Member {
	member := new(Member)
	if err := ormer.QueryTable(member).Filter("tel", tel).Filter("groupid", groupId).Limit(1).One(member, fields...); err != nil {
		return nil
	}

	return member
}

// 查找会员信息(ID)
func GetMemberById(id int64, fields []string) *Member {
	member := new(Member)
	if err := ormer.QueryTable(member).Filter("id", id).Limit(1).One(member, fields...); err != nil {
		return nil
	}

	return member
}

func GetMemberByHongId(hongid int64, fields []string) *Member {
	member := new(Member)
	if err := ormer.QueryTable(member).Filter("hongid", hongid).Limit(1).One(member, fields...); err != nil {
		return nil
	}

	return member
}

// hongid是否已经存在
func HasHongId(hongid int64) bool {
	return ormer.QueryTable(new(Member)).Filter("hongid", hongid).Exist()
}

// 验证手机号是否已被注册
func HasTel(tel string, groupid int64) bool {
	return ormer.QueryTable(new(Member)).Filter("tel", tel).Filter("groupid", groupid).Exist()
}

// 通过hongid，返回会员ID
func GetMemberIdByHongid(hongid int64) int64 {
	member := new(Member)
	err := ormer.QueryTable(member).Filter("hongid", hongid).Limit(1).One(member, "id")
	if err != nil {
		return 0
	}

	return member.Id
}
