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
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/controller"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	"hongId/models"
	"strings"
)

var (
	_ controller.BasicController
	_ beego.Controller
)

// 会员 API
type MemberController struct {
	BaseController
}

// /v1/members/s?identify=tel&value=18610889275&fields=nickname,password&groupid=2
// @router /s [put]
func (c *MemberController) Update() {
	argsOriginal := make(map[string]interface{})
	args := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &argsOriginal); err != nil {
		c.RenderInternalError()
		return
	}

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

	// handle error
	if c.HandleParamError() {
		return
	}

	identifyFlag := c.GetString("identify")
	identifyVal := c.GetString("value")
	groupId, _ := c.GetInt64("groupid")
	updateFlag := false
	switch identifyFlag {
	case "id":
		updateFlag = hm.UpdateMemberById(convert.Str2Int64(identifyVal), args, models.Writter)
	case "hongid":
		updateFlag = hm.UpdateMemberByHongId(convert.Str2Int64(identifyVal), args, models.Writter)
	case "tel":
		updateFlag = hm.UpdateMemberByTel(identifyVal, groupId, args, models.Writter)
	default:
		c.RenderUnSupportedIdentifyError()
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	if !updateFlag {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// /v1/members/u?groupid=2&fields=id
// @router /u [get]
func (c *MemberController) GetUnUsed() {
	groupid, err := c.GetInt64("groupid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("groupid", "会员组ID参数值错误."))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	member := hm.GetUnUsedMember(groupid, models.Reader)
	if member == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_MEMBER)
	}

	// parse member to map[string]interface{}
	memberParse, e := convert.ParseStruct(member, "orm", "column")
	if e != nil {
		c.Debug("parse struct error: %v", e)
		clientRsp.Status = errors.NewStatusInternalError()
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := memberParse[field]; ok {
				body[field] = v
			}
		}
		clientRsp.Body = body
	} else {
		clientRsp.Body = memberParse
	}

	c.RenderJson(clientRsp)
}

// /v1/members?identify=tel&value=18610889275&fields=id,hongid,tel,email&groupid=2
// @router /s [get]
func (c *MemberController) Get() {

	identifyFlag := c.GetString("identify")
	identifyVal := c.GetString("value")
	groupId, _ := c.GetInt64("groupid")

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	member := new(hm.Member)
	switch identifyFlag {
	case "id":
		member = hm.GetMemberById(convert.Str2Int64(identifyVal), models.Reader)
	case "hongid":
		member = hm.GetMemberByHongId(convert.Str2Int64(identifyVal), models.Reader)
	case "tel":
		member = hm.GetMemberByTel(identifyVal, groupId, models.Reader)
	default:
		c.RenderUnSupportedIdentifyError()
		return
	}
	if member == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_MEMBER)
	}

	// parse member to map[string]interface{}
	memberParse, e := convert.ParseStruct(member, "orm", "column")
	if e != nil {
		c.Debug("parse struct error: %v", e)
		clientRsp.Status = errors.NewStatusInternalError()
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := memberParse[field]; ok {
				body[field] = v
			}
		}
		clientRsp.Body = body
	} else {
		clientRsp.Body = memberParse
	}

	c.RenderJson(clientRsp)
}

type ReqGenMember struct {
	MinNo int64 `json:"min"`
	MaxNo int64 `json:"max"`
	Count int64 `json:"count"`
	Group int64 `json:"group"`
}

// curl -i -H "Content-Type: application/json" -d '{"min":10000,"max":99999,"count":1000,"group":1}' 127.0.0.1:8081/v1/members
// curl -i -H "Content-Type: application/json" -d '{"min":10000,"max":99999,"count":1000,"group":1}' 123.57.132.7:8081/v1/members
// @router / [post]
func (c *MemberController) SysGenMembers() {

	reqMsg := new(ReqGenMember)
	if err := json.Unmarshal(c.GetHttpBody(), reqMsg); err != nil {
		c.RenderInternalError()
		return
	}

	hm.GenMembers(reqMsg.MinNo, reqMsg.MaxNo, reqMsg.Count, reqMsg.Group, models.Writter)

	c.RenderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
}
