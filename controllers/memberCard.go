// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"net/http"
	"hongId/models"
	"github.com/globalways/utils_go/errors"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/page"
)

// MemberCard API
type MemberCardController struct {
	BaseController
}

// @router / [post]
func (c *MemberCardController) Post() {
	cardCnt, err1 := c.GetInt64("cnt")
	if err1 != nil {
		c.appenWrongParams(errors.NewFieldError("card cnt", err1.Error()))
	}

	merchant, err2 := c.GetInt64("merchant")
	if err2 != nil {
		c.appenWrongParams(errors.NewFieldError("card cnt", err2.Error()))
	}

	card := new(hm.ReqCard)
	if err := json.Unmarshal(c.getHttpBody(), card); err != nil {
		c.appenWrongParams(errors.NewFieldError("memberCard json", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	cards, gErr := hm.GenMemberCards(card, merchant, cardCnt, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_DATA_EXIST {
			c.setHttpStatus(http.StatusOK)
		} else if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.setHttpStatus(http.StatusCreated)
	c.renderJson(cards)
}

// @router / [get]
func (c *MemberCardController) GetAll() {

	var pageNum, pageSize int64
	if page, err := c.GetInt64("page"); err != nil {
		c.appenWrongParams(errors.NewFieldError("page", err.Error()))
	} else {
		pageNum = page
	}
	if size, err := c.GetInt64("size"); err != nil {
		c.appenWrongParams(errors.NewFieldError("size", err.Error()))
	} else {
		pageSize = size
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	pager := page.NewDBPaginator(int(pageNum), int(pageSize))

	cards, gErr := hm.FindMemberCard(pager, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(cards)
}

// @router /id/:id [get]
func (c *MemberCardController) Get() {
	cardId, err := c.GetInt64(":id")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("memberCardId", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	card, gErr := hm.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.renderJson(card)
}

// @router /id/:id/qrcode [get]
func (c *MemberCardController) GetQrCode() {
	cardId, err := c.GetInt64(":id")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("memberCardId", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	card, gErr := hm.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.renderPng(card.GenQrStream())
}

// @router /card/:card/bind/:owner [post]
func (c *MemberCardController) BindCard() {
	cardStr := c.GetString(":card")
	ownerId, err := c.GetInt64(":owner")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError(":owner", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	gErr := hm.BindMemberCardOwner(cardStr, ownerId, models.Writter)

	c.renderJson(errors.NewCommonOutRsp(gErr))
}

// @router /card/:card/unbind [post]
func (c *MemberCardController) UnBindCard() {
	cardStr := c.GetString(":card")

	gErr := hm.UnBindMemberCardOwner(cardStr, models.Writter)

	c.renderJson(errors.NewCommonOutRsp(gErr))
}
