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
	"github.com/globalways/utils_go/controller"
	"github.com/astaxie/beego"
	sm "github.com/globalways/chain_store_models/models"
	"github.com/globalways/utils_go/errors"
	"hongId/models"
)

var (
	_ controller.BasicController
	_ beego.Controller
)

type StoreIndustryController struct {
	BaseController
}

// 新增行业分类
// @router /industries [post]
func (c *StoreIndustryController) NewIndustry() {
	args := make(map[string]interface {})
	if err := json.Unmarshal(c.GetHttpBody(), &args); err != nil {
		c.RenderInternalError();return
	}

	var industryName string
	if v, ok := args["industryname"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("industryName", "参数传递错误."))
	} else {
		industryName = v.(string)
	}

	if c.HandleParamError() {return}

	clientRsp := errors.NewClientRspOK()
	industry := sm.NewStoreIndustry(industryName, models.Writter)
	if industry == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 删除行业分类
// @router /industries/:id [delete]
func (c *StoreIndustryController) DeleteIndustry() {

}

