// Copyright 2015 mit.zhao.chiu@gmail.com
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
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	sm "github.com/globalways/chain_store_models/models"
	"github.com/globalways/utils_go/controller"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/page"
	"hongID/models"
	"strings"
)

var (
	_ controller.BasicController
	_ beego.Controller
)

type OrderController struct {
	BaseController
}
type ReqNewOrder struct {
	StoreId           int64
	StoreName         string
	HongId            int64
	MemberCard        string
	AddressId         int64
	Payment           byte
	Delivery          byte
	DeliveryPrice     uint64
	CanvassPrice      uint64
	Discounted        uint64
	ProductIds        string
	ProductNames      string
	ProductCounts     string
	ProductUnitPrices string
	Comment           string
}

// 新建订单
// /v1/orders?fields=orderid,store_id,store_name
// @router / [post]
func (c *OrderController) NewOrder() {
	reqMsg := new(ReqNewOrder)
	if err := json.Unmarshal(c.GetHttpBody(), reqMsg); err != nil {
		c.RenderInternalError()
		return
	}

	c.Validation(reqMsg)

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	order := sm.NewOrder(reqMsg.StoreId,
		reqMsg.AddressId,
		reqMsg.StoreName,
		reqMsg.MemberCard,
		reqMsg.Comment,
		reqMsg.HongId,
		reqMsg.Payment,
		reqMsg.Delivery,
		reqMsg.DeliveryPrice,
		reqMsg.Discounted,
		reqMsg.CanvassPrice,
		reqMsg.ProductIds,
		reqMsg.ProductNames,
		reqMsg.ProductCounts,
		reqMsg.ProductUnitPrices,
		models.Writter,
	)
	if order == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_GEN_ORDER)
	}

	orderParse, e := convert.ParseStruct(order, "orm", "column")
	if e != nil {
		c.Debug("parse order err: %v", e)
		clientRsp.Status = errors.NewStatusInternalError()
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := orderParse[field]; ok {
				body[field] = v
			}
		}

		clientRsp.Body = body
	} else {
		clientRsp.Body = orderParse
	}

	c.RenderJson(clientRsp)
}

// 查询特定订单简要信息
// /v1/orders/xxxxxxxxx?fields=store_name,store_id
// @router /:orderid [get]
func (c *OrderController) OrderInfo() {
	orderid := c.GetString("orderid")

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()

	order := sm.FindOrder(orderid, models.Reader)
	if order == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_ORDER)
	}

	orderParse, e := convert.ParseStruct(order, "orm", "column")
	if e != nil {
		c.Debug("parse order err: %v", e)
		clientRsp.Status = errors.NewStatusInternalError()
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := orderParse[field]; ok {
				body[field] = v
			}
		}

		clientRsp.Body = body
	} else {
		clientRsp.Body = orderParse
	}

	c.RenderJson(clientRsp)
}

// 查询用户的订单简要信息
// /v1/orders/buyer/475847?fields=store_name,store_id&page=1&size=10
// @router /buyer/:buyer [get]
func (c *OrderController) BuyerOrders() {
	buyer := c.GetString("buyer")

	var pageNum, pageSize int64
	if page, err := c.GetInt64("page"); err != nil {
		c.AppenWrongParams(errors.NewFieldError("page", err.Error()))
	} else {
		pageNum = page
	}
	if size, err := c.GetInt64("size"); err != nil {
		c.AppenWrongParams(errors.NewFieldError("size", err.Error()))
	} else {
		pageSize = size
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()

	pager := page.NewDBPaginator(int(pageNum), int(pageSize))
	orders := sm.FindOrdersByHongId(buyer, pager, models.Reader)
	if orders == nil || len(orders) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	ordersParse := make([]map[string]interface{}, len(orders))
	for _, order := range orders {
		if orderParse, e := convert.ParseStruct(order, "orm", "column"); e != nil {
			continue
		} else {
			ordersParse = append(ordersParse, orderParse)
		}
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make([]map[string]interface{}, len(ordersParse))
		for idx, order := range ordersParse {
			for _, field := range fields {
				if v, ok := order[field]; ok {
					body[idx][field] = v
				}
			}
		}
	} else {
		clientRsp.Body = ordersParse
	}

	c.RenderJson(clientRsp)
}

// 查询商铺的订单简要信息
// /v1/orders/stores/23?fields=store_name,store_id&page=1&size=10
// @router /stores/:storeid [get]
func (c *OrderController) StoreOrders() {
	storeid, err := c.GetInt64("storeid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("storeid", "商铺ID参数格式错误."))
	}

	var pageNum, pageSize int64
	if page, err := c.GetInt64("page"); err != nil {
		c.AppenWrongParams(errors.NewFieldError("page", err.Error()))
	} else {
		pageNum = page
	}
	if size, err := c.GetInt64("size"); err != nil {
		c.AppenWrongParams(errors.NewFieldError("size", err.Error()))
	} else {
		pageSize = size
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()

	pager := page.NewDBPaginator(int(pageNum), int(pageSize))
	orders := sm.FindOrdersByStoreId(storeid, pager, models.Reader)
	if orders == nil || len(orders) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	ordersParse := make([]map[string]interface{}, len(orders))
	for _, order := range orders {
		if orderParse, e := convert.ParseStruct(order, "orm", "column"); e != nil {
			continue
		} else {
			ordersParse = append(ordersParse, orderParse)
		}
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make([]map[string]interface{}, len(ordersParse))
		for idx, order := range ordersParse {
			for _, field := range fields {
				if v, ok := order[field]; ok {
					body[idx][field] = v
				}
			}
		}
	} else {
		clientRsp.Body = ordersParse
	}

	c.RenderJson(clientRsp)
}

type ReqNewOrderProcess struct {
	OrderId     string `json:"orderid"`
	ProcessDesc string `json:"processdesc"`
	Status      byte   `json:"status"`
}

func (p *ReqNewOrderProcess) Valid(v *validation.Validation) {
	if p.Status < sm.EOrderStatus_error || p.Status > sm.EOrderStatus_finished {
		v.SetError("status", "订单状态码错误.")
	}
}

// 新建订单处理流程
// /v1/orders/xxxxxxxxxxxxxx/processes?fields=orderid,status,desc
// @router /:orderid/processes [post]
func (c *OrderController) NewProcess() {
	reqMsg := new(ReqNewOrderProcess)
	if err := json.Unmarshal(c.GetHttpBody(), reqMsg); err != nil {
		c.RenderInternalError()
		return
	}

	orderid := c.GetString(":orderid")

	c.Validation(reqMsg)

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	process := sm.NewOrderProcess(
		orderid,
		reqMsg.ProcessDesc,
		reqMsg.Status,
		models.Writter,
	)
	if process == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_GEN_ORDER_PROCESS)
	}

	processParse, e := convert.ParseStruct(process, "orm", "column")
	if e != nil {
		c.Debug("parse order process err: %v", e)
		clientRsp.Status = errors.NewStatusInternalError()
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make(map[string]interface{})
		for _, field := range fields {
			if v, ok := processParse[field]; ok {
				body[field] = v
			}
		}

		clientRsp.Body = body
	} else {
		clientRsp.Body = processParse
	}

	c.RenderJson(clientRsp)
}

// 查询订单处理流程
// /v1/orders/xxxxxxxxxxx/processes?fields=orderid,status,desc
// @router /:orderid/processes [get]
func (c *OrderController) OrderProcesses() {
	orderid := c.GetString(":orderid")

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()

	processes := sm.FindOrderProcess(orderid, models.Reader)
	if processes == nil || len(processes) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_NO_ORDER_PROCESS)
	}

	processesParse := make([]map[string]interface{}, len(processes))
	for _, process := range processes {
		if processParse, e := convert.ParseStruct(process, "orm", "column"); e != nil {
			continue
		} else {
			processesParse = append(processesParse, processParse)
		}
	}

	fields := strings.Split(c.GetString("fields"), ",")
	if len(fields) != 0 {
		body := make([]map[string]interface{}, len(processesParse))
		for idx, process := range processesParse {
			for _, field := range fields {
				if v, ok := process[field]; ok {
					body[idx][field] = v
				}
			}
		}
	} else {
		clientRsp.Body = processesParse
	}

	c.RenderJson(clientRsp)
}
