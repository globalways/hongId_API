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

// curl -i -H "Content-Type: application/json" -d '{"GroupName":"APP会员","GroupDesc":"APP会员哟","Contribution":0,"Status":1}' 127.0.0.1:8081/v1/memberGroups
// curl -i -H "Content-Type: application/json" -d '{"GroupName":"APP会员","GroupDesc":"APP会员哟","Contribution":0,"Status":1}' http://123.57.132.7:8081/v1/memberGroups
// @router / [post]
func (c *MemberGroupController) NewGroup() {
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

	c.setHttpHeader("Location", c.combineUrl(beego.UrlFor("MemberGroupController.GetGroup", ":groupId", convert.Int642str(id))))
	c.setHttpStatus(http.StatusCreated)
	c.renderJson(group)
}

// @router / [get]
func (c *MemberGroupController) GetGroupALL() {
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

// @router /:groupId [get]
func (c *MemberGroupController) GetGroup() {
	groupId, err := c.GetInt64(":groupId")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError(":groupId", err.Error()))
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

// @router /:groupId [put]
func (c *MemberGroupController) UpdateALL() {
	groupId, err := c.GetInt64(":groupId")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("groupId", err.Error()))
	}

	group := new(hm.MemberGroup)
	err = json.Unmarshal(c.getHttpBody(), group)
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("group Json", err.Error()))
	}

	if group.Id != groupId {
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

