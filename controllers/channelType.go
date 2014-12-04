// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hongId/models"
	e "github.com/globalways/utils_go/errors"
	"net/http"
	"github.com/globalways/utils_go/convert"
)

// channelType API
type ChannelTypeController struct {
	BaseController
}

// @Title createChannelType
// @Description generate a channel type
// @Param	channelType		body 	models.ChannelType true		"channel type json param"
// @Success 201 {object} models.ChannelType
// @Failure 200 channelType exist already
// @Failure 400 invalid http request param
// @Failure 500 internal server error
// @router / [post]
func (c *ChannelTypeController) Post() {
	channelType := new(models.ChannelType)
	err := json.Unmarshal(c.getHttpBody(), channelType)
	if err != nil {
		c.appenWrongParams(models.NewFieldError("channelType json", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	id, gErr := models.NewChannelType(channelType, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_DATA_EXIST {
			c.setHttpStatus(http.StatusOK)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutRsp(gErr))
		return
	}

	channelType.Id = id

	c.setHttpHeader("Location", c.combineUrl(beego.UrlFor("ChannelTypeController.Get", ":channelId", convert.Int642str(id))))
	c.setHttpStatus(http.StatusCreated)
	c.renderJson(channelType)
}

// @Title GetChannels
// @Description get all the channels
// @Success 200 {object} models.ChannelType
// @Failure 404 channel type list is blank
// @Failure 500 internal server error
// @router / [get]
func (c *ChannelTypeController) GetAll() {
	channelList, gErr := models.FindMemberCardChannel(models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(channelList)
}

// @Title GetChannel
// @Description get channel by channelid
// @Param channelId path int true "channel id"
// @Success 200 {object} models.ChannelType
// @Failure 400 invalid http request param
// @Failure 404 channel not found
// @Failure 500 interval server error
// @router /:channelId [get]
func (c *ChannelTypeController) Get() {
	channelId, err := c.GetInt(":channelId")
	if err != nil {
		c.appenWrongParams(models.NewFieldError("channelId", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	channel, gErr := models.GetChannelType(channelId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(channel)
}

// @Title updateChannel
// @Description update channel
// @Param channelId	path string	true "channel id"
// @Param channelType body models.ChannelType true "after update's channelType info"
// @Success 200 {object} models.ChannelType
// @Failure 400 invalid http request param
// @Failure 404 channeltype not found
// @Failure 500 internal server error
// @router /:channelId [put]
func (c *ChannelTypeController) Put() {

	channelId, err := c.GetInt(":channelId")
	if err != nil {
		c.appenWrongParams(models.NewFieldError("channelId", err.Error()))
	}

	channel := new(models.ChannelType)
	err = json.Unmarshal(c.getHttpBody(), channel)
	if err != nil {
		c.appenWrongParams(models.NewFieldError("channelType Json", err.Error()))
	}

	if channel.Id == 0 {
		channel.Id = channelId
	} else if channel.Id != channelId {
		c.appenWrongParams(models.NewFieldError("channelId", "path channelId & json channelId didn't match."))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	_, gErr := models.UpdateChannelType(channel, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(channel)
}

