// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/globalways/gws_utils_go/errors"
)

type ChannelType struct {
	Id          int64
	ChannelName string `orm:"column(channel_name);unique"`
	ChannelDesc string `orm:"column(channel_desc);null"`
}

func (c *ChannelType) TableName() string {
	return "channel_type"
}

// 新建渠道
func NewChannelType(channel *ChannelType, ormer orm.Ormer) (int64, errors.GlobalWaysError) {

	if isChannelExist(channel, ormer) {
		return 0, errors.New(errors.CODE_DB_DATA_EXIST)
	}

	if id, err := ormer.Insert(channel); err != nil {
		return 0, errors.Wrap(errors.CODE_DB_ERR_INSERT, err)
	} else {
		return id, errors.ErrorOK()
	}
}

// 检查渠道是否存在
func isChannelExist(channel *ChannelType, ormer orm.Ormer) bool {
	cond := orm.NewCondition()
	cond.Or("channel_name", channel.ChannelName).Or("id", channel.Id)
	return ormer.QueryTable(new(ChannelType)).SetCond(cond).Exist()
}

// 查找渠道
func FindMemberCardChannel(ormer orm.Ormer) ([]*ChannelType, errors.GlobalWaysError) {

	channels := make([]*ChannelType, 0)

	if num, err := ormer.QueryTable(new(ChannelType)).All(&channels); err != nil {
		return nil, errors.Wrap(errors.CODE_DB_ERR_FIND, err)
	} else if num == 0 {
		return nil, errors.New(errors.CODE_DB_ERR_NODATA)
	}

	return channels, errors.ErrorOK()
}

// 根据ID查找渠道
func GetChannelType(id int64, ormer orm.Ormer) (*ChannelType, errors.GlobalWaysError) {
	channelType := &ChannelType{
		Id: id,
	}

	err := ormer.Read(channelType)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New(errors.CODE_DB_ERR_NODATA)
		} else {
			return nil, errors.Wrap(errors.CODE_DB_ERR_GET, err)
		}
	}

	return channelType, errors.ErrorOK()
}

// 更新渠道
func UpdateChannelType(channel *ChannelType, ormer orm.Ormer) (bool, errors.GlobalWaysError) {
	if !isChannelExist(channel, ormer) {
		return false, errors.New(errors.CODE_DB_ERR_NODATA)
	}

	_, err := ormer.Update(channel)
	if err != nil {
		return false, errors.Wrap(errors.CODE_DB_ERR_UPDATE, err)
	}

	return true, errors.ErrorOK()
}
