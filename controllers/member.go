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
	"github.com/globalways/utils_go/errors"
	"hongId/models"
	hm "github.com/globalways/hongId_models/models"
)

// 会员 API
type MemberController struct {
	BaseController
}

// app 手机注册
type ReqRegisterMemberByTel struct {
	Tel string `json:"tel"`
	Group int64 `json:"group"`
}

// @router /register/tel [post]
func (c *MemberController) RegisterByTel() {
	reqMsg := new(ReqRegisterMemberByTel)
	if err := json.Unmarshal(c.GetHttpBody(), reqMsg); err != nil {
		c.RenderInternalError()
		return
	}

	// Validation
	c.Validation(reqMsg)

	// handle error
	if c.HandleParamError() {
		return
	}

	// 根据手机号查找会员信息
	member, gErr := hm.GetMemberByTel(reqMsg.Tel, models.Reader)
	switch gErr.GetCode() {
	case errors.CODE_DB_ERR_GET: // 系统内部错误
		c.RenderInternalError()
		return
	case errors.CODE_DB_ERR_NODATA: // 手机号尚未注册
		if _, gE := hm.RegisterMemberByTel(reqMsg.Group, reqMsg.Tel, models.Writter); gE.IsError() { //注册失败
			c.RenderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_REG))
			return
		}
	case errors.CODE_SUCCESS: // 手机号已经存在
		if member.Status != hm.EMemberStatus_Req { // 手机号已经被其他用户使用
			c.RenderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_TEL_ALREADY_IN))
			return
		}
	}

	// 注册成功
	c.RenderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
}

// @router /id/:memberId [put]
func (c *MemberController) UpdateALL() {
	member := new(hm.Member)
	if err := json.Unmarshal(c.GetHttpBody(), member); err != nil {
		c.RenderInternalError()
		return
	}

	memberId, err := c.GetInt64(":memberId")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":memberId", "会员卡ID参数错误."))
	}

	if memberId != member.Id {
		c.AppenWrongParams(errors.NewFieldError(":memberId", "会员卡ID不匹配."))
	}

	// handle error
	if c.HandleParamError() {
		return
	}

	gErr := hm.UpdateMember(member, models.Writter)
	switch gErr.GetCode() {
	case errors.CODE_SUCCESS: // 更新成功
		c.RenderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
		return
	default:
		c.RenderInternalError()
		return
	}
}

// @router /id/:memberId [patch]
func (c *MemberController) Update() {
	args := make(map[string]interface {})
	if err := json.Unmarshal(c.GetHttpBody(), &args); err != nil {
		c.RenderInternalError()
		return
	}

	memberId, err := c.GetInt64(":memberId")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":memberId", "会员卡ID参数错误."))
	}

	// handle error
	if c.HandleParamError() {
		return
	}

	gErr := hm.UpdateMemberById(memberId, args, models.Writter)
	switch gErr.GetCode() {
	case errors.CODE_SUCCESS: // 更新成功
		c.RenderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
		return
	default:
		c.RenderInternalError()
		return
	}
}

// @router /tel/:tel [get]
func (c *MemberController) GetByTel() {
	tel := c.GetString(":tel")

	member, gErr := hm.GetMemberByTel(tel, models.Reader)
	switch gErr.GetCode() {
	case errors.CODE_DB_ERR_NODATA:
		c.RenderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_USER_NAME))
		return
	case errors.CODE_SUCCESS:
		rspMsg := new(errors.ClientRsp)
		rspMsg.Status = errors.NewStatus(errors.CODE_SUCCESS)
		rspMsg.Body = member

		c.RenderJson(rspMsg)
		return
	default:
		c.RenderInternalError()
		return
	}
}

// @router /id/:id [get]
func (c *MemberController) GetById() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":id", "会员ID参数错误."))
	}

	if c.HandleParamError() {
		return
	}

	member, gErr := hm.GetMemberById(id, models.Reader)
	switch gErr.GetCode() {
	case errors.CODE_DB_ERR_NODATA:
		c.RenderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_USER_ID))
		return
	case errors.CODE_SUCCESS:
		rspMsg := new(errors.ClientRsp)
		rspMsg.Status = errors.NewStatus(errors.CODE_SUCCESS)
		rspMsg.Body = member

		c.RenderJson(rspMsg)
		return
	default:
		c.RenderInternalError()
		return
	}
}

type ReqGenMember struct {
	MinNo  int64 `json:"min"`
	MaxNo  int64 `json:"max"`
	Count  int64 `json:"count"`
	Group  int64 `json:"group"`
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
