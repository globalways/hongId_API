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
	"encoding/json"
	"github.com/globalways/utils_go/errors"
	sm "github.com/globalways/chain_store_models/models"
	"hongId/models"
	"strings"
	"fmt"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/page"
)

var (
	_ controller.BasicController
	_ beego.Controller
)

type ProductTagController struct {
	BaseController
}

// 新建商品tag
// @router /products/tags [post]
func (c *ProductTagController) NewProductTag() {
	args := make(map[string]interface {})
	if err := json.Unmarshal(c.GetHttpBody(), &args); err != nil {
		c.RenderInternalError();return
	}

	var tagName string
	var tagIcon string
	if v, ok := args["tagname"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("tagname", "参数传递错误."))
	} else {
		tagName = v.(string)
	}

	if v, ok := args["tagicon"]; !ok {
		c.AppenWrongParams(errors.NewFieldError("tagicon", "参数传递错误."))
	} else {
		tagIcon = v.(string)
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := errors.NewClientRspOK()
	productTag := sm.NewProductTag(tagName, tagIcon, models.Writter)
	if productTag == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 删除商品tag
// @router /products/tags/:id [delete]
func (c *ProductTagController) DeleteProductTag() {
	tagId, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":id", "商品tag参数错误."))
	}

	if c.HandleParamError() {return}

	clientRsp := errors.NewClientRspOK()
	if !sm.DeleteProductTag(tagId, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 更新商品tag
// @router /products/tags/:id [put]
func (c *ProductTagController) UpdateProductTag() {
	tagId, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":id", "商品tag参数错误."))
	}

	argsOriginal := make(map[string]interface{})
	args := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &argsOriginal); err != nil {
		c.RenderInternalError();return
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
	if !sm.UpdateProductTag(tagId, args, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 查询商品tag
// @router /products/tags [get]
func (c *ProductTagController) GetTags() {

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")

	clientRsp := errors.NewClientRspOK()
	tags := sm.FindProductTag(fields, models.Reader)
	if tags == nil || len(tags) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	tagsParse := make([]map[string]interface{}, 0)
	for _, tag := range tags {
		tagParse, err := convert.ParseStruct(tag, "orm", "column")
		if err != nil {
			continue
		}

		if len(fields) != 0 {
			body := make(map[string]interface{})
			for _, field := range fields {
				if v, ok := tagParse[field]; ok {
					body[field] = v
				}
			}
			tagsParse = append(tagsParse, body)
		} else {
			tagsParse = append(tagsParse, tagParse)
		}

	}
	clientRsp.Body = tagsParse

	c.RenderJson(clientRsp)
}

// 查询商品tag 分页
// @router /products/tags/p [get]
func (c *ProductTagController) GetTagsByPage() {

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
	tags := sm.FindProductTagByPage(pager, fields, models.Reader)
	if tags == nil || len(tags) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	tagsParse := make([]map[string]interface{}, 0)
	for _, tag := range tags {
		tagParse, err := convert.ParseStruct(tag, "orm", "column")
		if err != nil {
			continue
		}

		if len(fields) != 0 {
			body := make(map[string]interface{})
			for _, field := range fields {
				if v, ok := tagParse[field]; ok {
					body[field] = v
				}
			}
			tagsParse = append(tagsParse, body)
		} else {
			tagsParse = append(tagsParse, tagParse)
		}

	}
	clientRsp.Body = tagsParse

	c.RenderJson(clientRsp)
}
