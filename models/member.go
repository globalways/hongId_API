// Copyright 2014 mit.zhao.chiu@gmail.com
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

type EMemberStatus byte

const (
	EMemberStatus_Enable  EMemberStatus = iota + 1 //正常，用户已经在使用
	EMemberStatus_Disable                          //异常，用户在使用，但被锁定
	EMemberStatus_Del                              //删除，已经被用户删除，但系统还未回收
	EMemberStatus_NotUse                           //初始，新生成 & 系统回收
	EMemberStatus_Liang                            //备用，系统生成后经过帐号过滤，被系统留着备用的，靓号..
	EMemberStatus_Reg                              //注册，用户在注册阶段
)

type ECalenDarType byte

const (
	ECalenDarType_Gregorian ECalenDarType = iota + 1 //公历
	ECalenDarType_Lunar                              //农历
)

type EIdentifyCardType byte

const (
	EIdentifyCardType_IDCard   EIdentifyCardType = iota + 1 //身份证
	EIdentifyCardType_Passport                              //护照
)

type ESexType byte

const (
	ESexType_Male ESexType = iota + 1 //男
	ESexType_Female
)

type EEduLevel byte

const (
	EEduLevel_None      EEduLevel = iota + 1 //文盲
	EEduLevel_Primary                        //小学
	EEduLevel_Secondary                      //中学
	EEduLevel_Senior                         //高中
	EEduLevel_Bachelor                       //本科
	EEduLevel_Master                         //研究生
	EEduLevel_Doctor                         //博士
)

type Member struct {
	Id          int64
	UUID        string        `orm:"column(uuid)"`
	HongId      string        `orm:"column(hongid)"`
	Tel         string        `orm:"column(tel);size(11);null"`
	Email       string        `orm:"column(email);null"`
	PassWord    string        `orm:"column(password);null"`
	NickName    string        `orm:"column(nickname);null;size(20)"`
	Group       *MemberGroup  `orm:"column(group);null;rel(fk);on_delete(set_null)"`
	MemberCards []*MemberCard `orm:"reverse(many)"`
	Growth      *MemberGrowth `orm:"column(growth);null;rel(fk);on_delete(set_default);default(1)"`
	GrowthPoint uint32        `orm:"column(growth_point);null"`
}

func (m *Member) TableName() string {
	return "member"
}

// 验证手机号是否已被注册, 注册1阶段
func IsTelRegistered(tel string) bool {

	if tel == "18610889275" {
		return true
	}

	return false
}
