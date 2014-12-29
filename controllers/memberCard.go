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
		c.AppenWrongParams(errors.NewFieldError("card cnt", err1.Error()))
	}

	merchant, err2 := c.GetInt64("merchant")
	if err2 != nil {
		c.AppenWrongParams(errors.NewFieldError("card cnt", err2.Error()))
	}

	card := new(hm.ReqCard)
	if err := json.Unmarshal(c.GetHttpBody(), card); err != nil {
		c.AppenWrongParams(errors.NewFieldError("memberCard json", err.Error()))
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	cards, gErr := hm.GenMemberCards(card, merchant, cardCnt, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_DATA_EXIST {
			c.SetHttpStatus(http.StatusOK)
		} else if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.SetHttpStatus(http.StatusNotFound)
		} else {
			c.SetHttpStatus(http.StatusInternalServerError)
		}

		c.RenderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.SetHttpStatus(http.StatusCreated)
	c.RenderJson(cards)
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

	pager := page.NewDBPaginator(int(pageNum), int(pageSize))

	cards, gErr := hm.FindMemberCard(pager, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.SetHttpStatus(http.StatusNotFound)
		} else {
			c.SetHttpStatus(http.StatusInternalServerError)
		}

		c.RenderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.RenderJson(cards)
}

// @router /id/:id [get]
func (c *MemberCardController) Get() {
	cardId, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("memberCardId", err.Error()))
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	card, gErr := hm.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.SetHttpStatus(http.StatusNotFound)
		} else {
			c.SetHttpStatus(http.StatusInternalServerError)
		}

		c.RenderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.RenderJson(card)
}

// @router /id/:id/qr [get]
func (c *MemberCardController) GetQrCode() {
	cardId, err := c.GetInt64(":id")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("memberCardId", "会员ID参数错误."))
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	card, gErr := hm.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == errors.CODE_DB_ERR_NODATA {
			c.SetHttpStatus(http.StatusNotFound)
		} else {
			c.SetHttpStatus(http.StatusInternalServerError)
		}

		c.RenderJson(errors.NewCommonOutRsp(gErr))
		return
	}

	c.RenderPng(card.GenQrStream())
}

// @router /card/:card/bind/:owner [post]
func (c *MemberCardController) BindCard() {
	cardStr := c.GetString(":card")
	ownerId, err := c.GetInt64(":owner")
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError(":owner", "会员ID参数错误."))
	}

	// handle http request param
	if c.HandleParamError() {
		return
	}

	gErr := hm.BindMemberCardOwner(cardStr, ownerId, models.Writter)
	switch gErr.GetCode() {
	case errors.CODE_BISS_ERR_HAS_OWNER, errors.CODE_BISS_ERR_VARIFY_CARD, errors.CODE_SUCCESS:
		c.RenderJson(errors.NewClientRsp(gErr.GetCode()))
	default:
		c.RenderInternalError()
	}
}

// @router /card/:card/unbind [post]
func (c *MemberCardController) UnBindCard() {
	cardStr := c.GetString(":card")

	gErr := hm.UnBindMemberCardOwner(cardStr, models.Writter)
	switch gErr.GetCode() {
	case errors.CODE_BISS_ERR_VARIFY_CARD, errors.CODE_SUCCESS:
		c.RenderJson(errors.NewClientRsp(gErr.GetCode()))
	default:
		c.RenderInternalError()
	}
}
