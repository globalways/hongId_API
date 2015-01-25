// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/controller"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	"hongID/models"
	"reflect"
	"strings"
)

var (
	_ controller.BasicController
)

// memberGroup API
type MemberGroupController struct {
	BaseController
}

// @router / [get]
func (c *MemberGroupController) GetGroupALL() {

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	groups := hm.FindMemberGroup(models.Reader)
	if groups == nil || len(groups) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_GROUP)
	}

	c.RenderJson(clientRsp)
}

// @router / [post]
func (c *MemberGroupController) NewGroup() {
	args := make(map[string]interface{})
	err := json.Unmarshal(c.GetHttpBody(), &args)
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("MemberGroup json", err.Error()))
	}

	var name, desc string
	var cont uint
	var status byte

	if v, ok := args["name"].(string); !ok {
		c.Debug("args name interface: %v", reflect.TypeOf(args["name"]))
		c.AppenWrongParams(errors.NewFieldError("name", "会员组name属性值错误."))
	} else {
		name = v
	}

	if v, ok := args["desc"].(string); !ok {
		c.Debug("args desc interface: %v", reflect.TypeOf(args["desc"]))
		c.AppenWrongParams(errors.NewFieldError("desc", "会员组desc属性值错误."))
	} else {
		desc = v
	}

	if v, ok := args["cont"].(float64); !ok {
		c.Debug("args cont interface: %v", reflect.TypeOf(args["cont"]))
		c.AppenWrongParams(errors.NewFieldError("cont", "会员组cont属性值错误."))
	} else {
		cont = uint(v)
	}

	if v, ok := args["status"].(float64); !ok {
		c.Debug("args status interface: %v", reflect.TypeOf(args["status"]))
		c.AppenWrongParams(errors.NewFieldError("status", "会员组status属性值错误."))
	} else {
		status = byte(v)
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	group := hm.NewMemberGroup(name, desc, cont, status, models.Writter)
	if group == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_GEN_GROUP)
	}

	clientRsp.Body = group
	c.SetHttpHeader("Location", c.CombineUrl(beego.UrlFor("MemberGroupController.GetGroup", ":groupId", convert.Int642str(group.Id))))

	c.RenderJson(clientRsp)
}

// @router /id/:groupId [get]
func (c *MemberGroupController) GetGroup() {
	groupId, err := c.GetInt64(":groupId")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":groupId", err.Error()))
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()

	group := hm.GetGroupById(groupId, models.Reader)
	if group == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_GROUP)
	}

	clientRsp.Body = group
	c.RenderJson(clientRsp)
}

// /v1/groups/id/2?fields=group_name,group_desc
// @router /id/:groupId [put]
func (c *MemberGroupController) UpdateALL() {
	groupId, err := c.GetInt64(":groupId")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("groupId", err.Error()))
	}

	argsOriginal := make(map[string]interface{})
	err = json.Unmarshal(c.GetHttpBody(), &argsOriginal)
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("group args", err.Error()))
	}
	args := make(map[string]interface{})

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")
	c.Debug("fields: %+v, len: %v", fields, len(fields))
	if len(fields) != 0 {
		for _, field := range fields {
			if val, ok := argsOriginal[field]; !ok {
				c.AppenWrongParams(errors.NewFieldError(field, fmt.Sprintf("更新参数值%v未传递.", field)))
			} else {
				args[field] = val
			}
		}
	} else {
		args = argsOriginal
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	if !hm.UpdateMemberGroupById(groupId, args, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}
