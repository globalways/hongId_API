// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/controller"
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/page"
	"hongID/models"
	"reflect"
)

var (
	_ controller.BasicController
)

// MemberCard API
type MemberCardController struct {
	BaseController
}

type ReqGenCard struct {
	MII      byte   `json:"mii"`
	CPI      byte   `json:"cpi"`
	CDI      uint16 `json:"cdi"`
	Count    int64  `json:"count"`
	Merchant string `json:"merchant"`
}

// @router / [post]
func (c *MemberCardController) GenCard() {
	args := make(map[string]interface{})
	if err := json.Unmarshal(c.GetHttpBody(), &args); err != nil {
		c.AppenWrongParams(errors.NewFieldError("memberCard json", err.Error()))
	}

	var mii, cpi byte
	var cdi uint16
	var count int64
	var merchant int64

	if v, ok := args["mii"].(float64); !ok {
		c.Debug("args mii interfact: %v", reflect.TypeOf(args["mii"]))
		c.AppenWrongParams(errors.NewFieldError("mii", "mii参数格式错误."))
	} else {
		mii = byte(v)
	}

	if v, ok := args["cpi"].(float64); !ok {
		c.Debug("args cpi interfact: %v", reflect.TypeOf(args["cpi"]))
		c.AppenWrongParams(errors.NewFieldError("cpi", "cpi参数格式错误."))
	} else {
		cpi = byte(v)
	}

	if v, ok := args["cdi"].(float64); !ok {
		c.Debug("args cdi interfact: %v", reflect.TypeOf(args["cdi"]))
		c.AppenWrongParams(errors.NewFieldError("cdi", "cdi参数格式错误."))
	} else {
		cdi = uint16(v)
	}

	if v, ok := args["count"].(float64); !ok {
		c.Debug("args count interfact: %v", reflect.TypeOf(args["count"]))
		c.AppenWrongParams(errors.NewFieldError("count", "count参数格式错误."))
	} else {
		count = int64(v)
	}

	if v, ok := args["merchant"].(float64); !ok {
		c.Debug("args merchant interfact: %v", reflect.TypeOf(args["merchant"]))
		c.AppenWrongParams(errors.NewFieldError("merchant", "merchant参数格式错误."))
	} else {
		merchant = int64(v)
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()
	cards := hm.GenMemberCards(mii, cpi, cdi, merchant, count, models.Writter)
	if cards == nil || len(cards) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_BISS_ERR_GEN_CARD)
	}

	clientRsp.Body = cards

	c.RenderJson(clientRsp)
}

// @router / [get]
func (c *MemberCardController) GetAll() {

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

	// handle http request param
	if c.HandleParamError() {
		return
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatusOK()

	pager := page.NewDBPaginator(int(pageNum), int(pageSize))
	cards := hm.FindMemberCard(pager, models.Reader)
	if cards == nil || len(cards) == 0 {
		clientRsp.Status = errors.NewStatus(errors.CODE_OPT_NO_MORE_DATA)
	}

	clientRsp.Body = cards

	c.RenderJson(clientRsp)
}
