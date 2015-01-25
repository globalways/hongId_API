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
	"github.com/globalways/utils_go/controller"
	"github.com/astaxie/beego"
	sm "github.com/globalways/chain_store_models/models"
	"encoding/json"
	"github.com/globalways/utils_go/errors"
	"hongID/models"
	"strings"
	"fmt"
)

var (
	_ controller.BasicController
	_ beego.Controller
)

type StoreProductController struct {
	BaseController
}

// 新建商铺商品
// @router /products [post]
func (c *StoreProductController) NewStoreProduct() {
	reqMsg := new(sm.StoreProduct)
	if err := json.Unmarshal(c.GetHttpBody(), reqMsg); err != nil {
		c.RenderInternalError();return
	}

	clientRsp := errors.NewClientRspOK()
	product := sm.NewStoreProduct(reqMsg, models.Writter)
	if product == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 更新商铺商品
// @router /products/:pid [put]
func (c *StoreProductController) UpdateStoreProduct() {
	argsOriginal := make(map[string]interface{})
	args := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &argsOriginal); err != nil {
		c.RenderInternalError()
		return
	}

	pid, err := c.GetInt64(":pid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":pid", "商品Id参数值错误."))
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

	clientRsp := errors.NewClientRspOK()
	if !sm.UpdateStoreProduct(pid, args, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

