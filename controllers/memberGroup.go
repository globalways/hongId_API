// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hongId/models"
	"net/http"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	hm "github.com/globalways/hongId_models/models"
)

// memberGroup API
type MemberGroupController struct {
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
func (c *MemberGroupController) Post() {
	group := new(hm.MemberGroup)
	err := json.Unmarshal(c.getHttpBody(), group)
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("MemberGroup json", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	id, gErr := hm.NewMemberGroup(group, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_DATA_EXIST {
			c.setHttpStatus(http.StatusOK)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	group.Id = id

	c.setHttpHeader("Location", c.combineUrl(beego.UrlFor("MemberGroupController.Get", ":groupId", convert.Int642str(id))))
	c.setHttpStatus(http.StatusCreated)
	c.renderJson(group)
}

// @Title GetChannels
// @Description get all the channels
// @Success 200 {object} models.ChannelType
// @Failure 404 channel type list is blank
// @Failure 500 internal server error
// @router / [get]
func (c *MemberGroupController) GetAll() {
	groups, gErr := hm.FindMemberGroup(models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(groups)
}

// @Title GetChannel
// @Description get channel by channelid
// @Param channelId path int true "channel id"
// @Success 200 {object} models.ChannelType
// @Failure 400 invalid http request param
// @Failure 404 channel not found
// @Failure 500 interval server error
// @router /:groupId [get]
func (c *MemberGroupController) Get() {
	groupId, err := c.GetInt64(":groupId")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("groupId", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	group, gErr := hm.GetGroupInfo(groupId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(group)
}

// @Title updateChannel
// @Description update channel
// @Param channelId	path string	true "channel id"
// @Param channelType body models.ChannelType true "after update's channelType info"
// @Success 200 {object} models.ChannelType
// @Failure 400 invalid http request param
// @Failure 404 channeltype not found
// @Failure 500 internal server error
// @router /:groupId [put]
func (c *MemberGroupController) Put() {

	groupId, err := c.GetInt64(":groupId")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("groupId", err.Error()))
	}

	group := new(hm.MemberGroup)
	err = json.Unmarshal(c.getHttpBody(), group)
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("group Json", err.Error()))
	}

	if group.Id == 0 {
		group.Id = groupId
	} else if group.Id != groupId {
		c.appenWrongParams(errors.NewFieldError("groupId", "path groupId & json groupId didn't match."))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	_, gErr := hm.UpdateGroupInfo(group, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(group)
}

