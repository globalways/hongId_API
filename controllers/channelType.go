// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hongId/models"
)

// 分发渠道API
type ChannelTypeController struct {
	beego.Controller
}

// @Title createChannelType
// @Description 创建一个分发渠道
// @Param	channelType		body 	models.ChannelType true		"分发渠道json"
// @Success 200 {object} models.ChannelType
// @Failure 403 body is empty
// @Failure 500 json unmarshal error
// @router / [post]
func (c *ChannelTypeController) Post() {
	channelType := new(models.ChannelType)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, channelType)
	if err != nil {
		c.Data["json"] = err
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	id, gErr := models.NewChannelType(channelType, models.Writter)
	if gErr.IsError() {
		c.Data["json"] = gErr.ErrorMessage()
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	channelType.Id = id
	c.Data["json"] = channelType
	c.ServeJson()
}

// @Title GetChannels
// @Description 获取所有的分发渠道
// @Success 200 {object} models.ChannelType
// @router / [get]
func (c *ChannelTypeController) GetAll() {
	channelList, gErr := models.FindMemberCardChannel(models.Reader)
	if gErr.IsError() {
		c.Data["json"] = gErr.ErrorMessage()
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	c.Data["json"] = channelList
	c.ServeJson()
}

// @Title GetChannel
// @Description 通过ID获取分发渠道
// @Param	channelId		path 	int64	true		"分发渠道ID"
// @Success 200 {object} models.ChannelType
// @Failure 403 :channelId is empty
// @router /:channelId [get]
func (c *ChannelTypeController) Get() {
	channelId, err := c.GetInt(":channelId")
	if err != nil {
		c.Data["json"] = err
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	channel, gErr := models.GetChannelType(channelId, models.Reader)
	if gErr.IsError() {
		c.Data["json"] = gErr.ErrorMessage()
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	c.Data["json"] = channel
	c.ServeJson()
}

// @Title updateChannel
// @Description 更新分发渠道信息
// @Param	channelId		path 	string	true		"想更新的渠道ID"
// @Param	channelType		body 	models.ChannelType	true		"更新的渠道信息"
// @Success 200 {object} models.ChannelType
// @Failure 403 :channelId is not int
// @router /:channelId [put]
func (c *ChannelTypeController) Put() {

	channelId, err := c.GetInt(":channelId")
	if err != nil {
		c.Data["json"] = err
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	channel := new(models.ChannelType)
	err = json.Unmarshal(c.Ctx.Input.RequestBody, channel)
	if err != nil {
		c.Data["json"] = err
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}
	if channel.Id == 0 {
		channel.Id = channelId
	}

	_, gErr := models.UpdateChannelType(channel, models.Writter)
	if gErr.IsError() {
		c.Data["json"] = gErr.ErrorMessage()
		c.Ctx.Output.Status = 203
		c.ServeJson()
		c.StopRun()
	}

	c.Data["json"] = channel
	c.ServeJson()
}

