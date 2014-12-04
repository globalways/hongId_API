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

type EMemberGroupStatus byte

const (
	EMemberGroupStatus_Enable EMemberGroupStatus = iota + 1
	EMemberGroupStatus_Distable
	EMemberGroupStatus_Del
)

var MemberGroupStatusMap map[EMemberGroupStatus]string = map[EMemberGroupStatus]string{
	EMemberGroupStatus_Enable:   "启用",
	EMemberGroupStatus_Distable: "禁用",
	EMemberGroupStatus_Del:      "删除",
}

type MemberGroup struct {
	Id           int64
	GroupName    string `xorm:"varchar(20) unique"`
	GroupDesc    string
	Contribution uint32             //会费
	Status       EMemberGroupStatus `xorm:"tinyint(1)"`
	StatusStr    string             `xorm:"-"`
	CreateTime   string             `xorm:"DateTime created"`
	UpdateTime   string             `xorm:"DateTime updated"`
	Members      []*Member          `orm:"reverse(many)"`
}
