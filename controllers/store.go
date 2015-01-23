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

type StoreController struct {
	BaseController
}

// 新增商铺
// @router / [post]
func (c *StoreController) NewStore() {
	store := new(sm.Store)
	if err := json.Unmarshal(c.GetHttpBody(), store); err != nil {
		c.RenderInternalError()
		return
	}

	clientRsp := errors.NewClientRspOK()
	store = sm.NewStore(store, models.Writter)
	if store == nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 删除商铺
// @router /:storeid [delete]
func (c *StoreController) DeleteStore() {
	storeid, err := c.GetInt64(":storeid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":storeid", "商铺Id参数值错误."))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := errors.NewClientRspOK()
	if !sm.DeleteStore(storeid, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 更新商铺
// /v1/stores/123483?fields=xx,xx,xx
// @router /:storeid [put]
func (c *StoreController) UpdateStore() {
	storeid, err := c.GetInt64(":storeid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":storeid", "商铺Id参数值错误."))
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
	if !sm.UpdateStore(storeid, args, models.Writter) {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	c.RenderJson(clientRsp)
}

// 查询单个商铺
// /v1/stores/139485?fields=xxx,xxx,xxx
// @router /:storeid [get]
func (c *StoreController) GetStore() {
	storeid, err := c.GetInt64(":storeid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":storeid", "商铺Id参数值错误."))
	}

	if c.HandleParamError() {
		return
	}

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")

	clientRsp := errors.NewClientRspOK()
	store := sm.GetStoreById(storeid, fields, models.Reader)
	if store == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_STORE)
	}

	storeParse, e := convert.ParseStruct(store, "orm", "column")
	if e != nil {
		clientRsp.Status = errors.NewStatusInternalError()
	}

	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := storeParse[field]; ok {
				body[field] = v
			}
		}
		clientRsp.Body = body
	} else {
		clientRsp.Body = storeParse
	}

	c.RenderJson(clientRsp)
}

// 查询商铺by adminid
// /v1/stores/a/234234?fields=xxx,xxxx,xxx
// @router /a/:adminid [get]
func (c *StoreController) GetStoresByAdmin() {
	adminId, err := c.GetInt64(":adminid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":adminid", "参数adminId格式错误."))
	}

	// 拆分fields
	fields := strings.Split(c.GetString("fields"), ",")
	c.Debug("fields: %+v", fields)

	if c.HandleParamError() {
		return
	}

	clientRsp := errors.NewClientRspOK()
	stores := sm.FindStoresByAdmin(adminId, fields, models.Reader)
	if stores == nil || len(stores) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	storesParse := make([]map[string]interface{}, len(stores))
	for _, store := range stores {
		storeParse, err := convert.ParseStruct(store, "orm", "column")
		if err != nil {
			continue
		}

		if len(fields) != 0 {
			body := make(map[string]interface{})
			for _, field := range fields {
				if v, ok := storeParse[field]; ok {
					body[field] = v
				}
			}

			storesParse = append(storesParse, body)
		} else {
			storesParse = append(storesParse, storeParse)
		}

	}
	clientRsp.Body = storesParse

	c.RenderJson(clientRsp)
}

// 查询商铺列表
// /v1/stores?fields=xxx,xxx,xxx&page=1&size=10
// @router / [get]
func (c *StoreController) GetStores() {

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

	stores := sm.FindStores(pager, fields, models.Reader)

	if stores == nil || len(stores) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	storesParse := make([]map[string]interface{}, len(stores))
	for _, store := range stores {
		storeParse, err := convert.ParseStruct(store, "orm", "column")
		if err != nil {
			continue
		}

		if len(fields) != 0 {
			body := make(map[string]interface{})
			for _, field := range fields {
				if v, ok := storeParse[field]; ok {
					body[field] = v
				}
			}

			storesParse = append(storesParse, body)
		} else {
			storesParse = append(storesParse, storeParse)
		}

	}
	clientRsp.Body = storesParse

	c.RenderJson(clientRsp)
}
