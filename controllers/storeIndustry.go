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
	"github.com/globalways/utils_go/page"
	"hongId/models"
	"strings"
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
	args := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &args); err != nil {
		c.RenderInternalError()
		return
	}

	var industryName string
	var industryIcon string
	if v, ok := args["industryname"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("industryName", "参数传递错误."))
	} else {
		industryName = v.(string)
	}

	if v, ok := args["industryicon"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("industryIcon", "参数传递错误."))
	} else {
		industryIcon = v.(string)
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := errors.NewClientRspOK()
	industry := sm.NewStoreIndustry(industryName, industryIcon, models.Writter)
	if industry == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 删除行业分类
// @router /industries/:id [delete]
func (c *StoreIndustryController) DeleteIndustry() {

	industryId, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":id", "参数行业Id格式错误."))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := errors.NewClientRspOK()
	if !sm.DeleteStoreIndustry(industryId, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 更新行业分类
// @router /industries/:id [put]
func (c *StoreIndustryController) UpdateIndustry() {
	industryId, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":id", "参数行业Id格式错误."))
	}

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

	clientRsp := errors.NewClientRspOK()
	if !sm.UpdateStoreIndustry(industryId, args, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 查询行业分类(分页)
// @rouer /industries/p [get]
func (c *StoreIndustryController) GetIndustriesByPage() {

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")

	var pageNum, pageSize int64
	if page := c.GetString("page"); len(page) == 0 {
		pageNum = 1
	} else {
		pageNum = convert.Str2Int64(page)
	}
	if size := c.GetString("size"); len(size) == 0 {
		pageSize = 10
	} else {
		pageSize = convert.Str2Int64(size)
	}

	if c.HandleParamError() {
		return
	}

	pager := page.NewDBPaginator(int(pageNum), int(pageSize))

	clientRsp := errors.NewClientRspOK()
	industries := sm.FindStoreIndustry(pager, fields, models.Reader)
	if industries == nil || len(industries) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	industriesParse := make([]map[string]interface{}, len(industries))
	for _, industry := range industries {
		industryParse, err := convert.ParseStruct(industry, "orm", "column")
		if err != nil {
			continue
		}

		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := industryParse[field]; ok {
				body[field] = v
			}
		}

		industriesParse = append(industriesParse, body)
	}
	clientRsp.Body = industriesParse

	c.RenderJson(clientRsp)
}

// 查询行业分类
// @router /industries [get]
func (c *StoreIndustryController) GetIndustries() {

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")

	clientRsp := errors.NewClientRspOK()
	industries := sm.FindALLStoreIndustry(fields, models.Reader)
	if industries == nil || len(industries) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	industriesParse := make([]map[string]interface{}, len(industries))
	for _, industry := range industries {
		industryParse, err := convert.ParseStruct(industry, "orm", "column")
		if err != nil {
			continue
		}

		if len(fields) != 0 {
			body := make(map[string]interface{})
			for _, field := range fields {
				if v, ok := industryParse[field]; ok {
					body[field] = v
				}
			}
			industriesParse = append(industriesParse, body)
		} else {
			industriesParse = append(industriesParse, industryParse)
		}

	}
	clientRsp.Body = industriesParse

	c.RenderJson(clientRsp)
}
