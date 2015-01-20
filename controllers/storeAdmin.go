// Copyright 2015 mint.zhao.chiu@gmail.com
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
	sm "github.com/globalways/chain_store_models/models"
	"github.com/globalways/utils_go/controller"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	"hongId/models"
	"strings"
)

var (
	_ beego.Controller
	_ controller.BasicController
)

type StoreAdminController struct {
	BaseController
}

// 新增商铺管理员
// /v1/stores/admins
// @router /admins [post]
func (c *StoreAdminController) NewStoreAdmin() {
	params := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &params); err != nil {
		c.RenderInternalError()
		return
	}

	var adminId int64
	var role byte

	if v, ok := params["adminid"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("adminid", "必须参数adminID未传递."))
	} else {
		adminId = int64(v.(float64))
	}

	if v, ok := params["role"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("role", "必须参数role未传递."))
	} else {
		role = byte(v.(float64))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	admin := sm.NewStoreAdmin(adminId, role, models.Writter)
	if admin == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 删除商铺管理员
// /v1/stores/admins/234958
// @router /admins/:adminid [delete]
func (c *StoreAdminController) DeleteStoreAdmin() {
	adminId, err := c.GetInt64(":adminid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":adminid", "参数adminId格式错误."))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	if !sm.DeleteStoreAdmin(adminId, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 更新商铺管理员
// /v1/stores/admins/184765?fields=role
// @router /admins/:adminid [put]
func (c *StoreAdminController) UpdateStoreAdmin() {
	args := make(map[string]interface{})
	argsOriginal := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &argsOriginal); err != nil {
		c.RenderInternalError()
		return
	}

	adminId, err := c.GetInt64("adminid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("adminid", "参数adminId格式错误."))
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

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	if !sm.UpdateStoreAdmin(adminId, args, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 查询商铺管理员
// /v1/stores/admins/1234558?fields=role
// @router /admins/:adminid [get]
func (c *StoreAdminController) GetStoreAdmin() {
	adminId, err := c.GetInt64(":adminid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("adminid", "参数adminId格式错误."))
	}

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	admin := sm.GetStoreAdmin(adminId, fields, models.Reader)
	if admin == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	adminParse, e := convert.ParseStruct(admin, "orm", "column")
	if e != nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := adminParse[field]; ok {
				body[field] = v
			}
		}
		clientRsp.Body = body
	} else {
		clientRsp.Body = adminParse
	}

	c.RenderJson(clientRsp)
}
