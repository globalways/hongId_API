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
	"github.com/astaxie/beego/validation"
	sm "github.com/globalways/chain_store_models/models"
	"github.com/globalways/utils_go/consts"
	"github.com/globalways/utils_go/controller"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	"hongId/models"
	"regexp"
)

var (
	_ controller.BasicController
)

type OrderAddressController struct {
	BaseController
}

type ReqAddr struct {
	HongId  int64  `json:"hongid"`
	Contact string `json:"contact"`
	Address string `json:"address"`
	Tel1    string `json:"tel1"`
	Tel2    string `json:"tel2"`
}

func (p *ReqAddr) Valid(v *validation.Validation) {

	if len(p.Contact) == 0 {
		v.SetError("contact", "联系人不能为空.")
	}

	if len(p.Address) == 0 {
		v.SetError("address", "联系地址不能为空.")
	}

	if b, _ := regexp.MatchString(consts.Regexp_Mobile, p.Tel1); !b {
		v.SetError("tel1", "联系电话格式错误.")
	}

	if len(p.Tel2) != 0 {
		if b, _ := regexp.MatchString(consts.Regexp_Mobile, p.Tel2); !b {
			v.SetError("tel2", "备用联系电话格式错误.")
		}
	}
}

// @router /addresses [post]
func (c *OrderAddressController) NewAddress() {
	reqMsg := new(ReqAddr)
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
	addr := sm.NewOrderAddress(reqMsg.HongId, reqMsg.Contact, reqMsg.Address, reqMsg.Tel1, reqMsg.Tel2, models.Writter)
	if addr == nil {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_GEN_ORDER_ADDRESS)
	}

	addrParse, err := convert.ParseStruct(addr, "orm", "column")
	if err != nil {
		c.Debug("parse addr struct err: %v", err)
		clientRsp.Status = errors.NewStatusInternalError()
	}

	clientRsp.Body = addrParse

	c.RenderJson(clientRsp)
}

// @router /addresses [delete]
func (c *OrderAddressController) DelAddress() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("id", "收货地址ID参数格式错误."))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	b := sm.DelOrderAddress(id, models.Writter)
	if !b {
		clientRsp.Status = errors.NewStatus(errors.CODE_DB_ERR_UPDATE)
	} else {
		clientRsp.Status = errors.NewStatusOK()
	}

	c.RenderJson(clientRsp)
}

// @router /addresses [put]
func (c *OrderAddressController) UpdateAddress() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("id", "收货地址ID参数格式错误."))
	}

	reqMsg := new(ReqAddr)
	if err := json.Unmarshal(c.GetHttpBody(), reqMsg); err != nil {
		c.RenderInternalError()
		return
	}

	c.Validation(reqMsg)

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	params := map[string]interface{}{
		"hongid":  reqMsg.HongId,
		"contact": reqMsg.Contact,
		"address": reqMsg.Address,
		"tel1":    reqMsg.Tel1,
		"tel2":    reqMsg.Tel2,
	}
	b := sm.UpdOrderAddress(id, params, models.Writter)
	if !b {
		clientRsp.Status = errors.NewStatus(errors.CODE_DB_ERR_UPDATE)
	} else {
		clientRsp.Status = errors.NewStatusOK()
	}

	c.RenderJson(clientRsp)
}

// @router /addresses [get]
func (c *OrderAddressController) ListAddress() {
	hongid, err := c.GetInt64("hongid")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("hongid", "HongID参数格式错误."))
	}

	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	addrs := sm.FindOrderAddressByHongId(hongid, models.Reader)
	if addrs == nil || len(addrs) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_DB_ERR_NODATA)
	}

	addrsParse := make([]map[string]interface{}, len(addrs))
	for _, addr := range addrs {
		c.Debug("addrs: %+v", addr)
		addrParse, err := convert.ParseStruct(addr, "orm", "column")
		if err != nil {
			continue
		}

		addrsParse = append(addrsParse, addrParse)
	}
	clientRsp.Body = addrsParse

	c.RenderJson(clientRsp)
}
