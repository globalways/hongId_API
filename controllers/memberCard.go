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

// @Title createMemberCards
// @Description generate member card batch
// @Param cnt query int true "used for generate membercard count, default is 0"
// @Param memberCard body models.MemberCard true "request param json"
// @Success 201 {object} models.MemberCard
// @Failure 400 request body is invalid
// @Failure 500 generate member card error
// @router / [post]
func (c *MemberCardController) Post() {
	cardCnt, err := c.GetInt64("cnt")
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("card cnt", err.Error()))
	}

	memberCard := new(hm.MemberCard)
	if err := json.Unmarshal(c.getHttpBody(), memberCard); err != nil {
		c.appenWrongParams(errors.NewFieldError("memberCard json", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	cards, gErr := hm.GenMemberCards(memberCard, cardCnt, models.Writter)
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

// @Title getMemberCards
// @Description get member card list by page & size
// @Param page query int64 true "page number"
// @Param size query int64 true "each page count"
// @Success 200 {object} models.MemberCard
// @Failure 400 request url's parameter is invalid
// @Failure 404 request resource not found
// @Failure 500 get memberCard list wrong
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

// @Title getMemberCardById
// @Description get member card info by id
// @Param	id		path 	int	true		"member card id"
// @Success 200 {object} models.MemberCard
// @Failure 400 http request param is invalid
// @Failure 404 member card not found
// @Failure 500 internal server error
// @router /:id [get]
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

// @Title getMemberCardQrCode
// @Description get member card qrcode by id
// @Param id path int true "member card id"
// @Success 200 {image/png} qrCode
// @Failure 400 invalid http request param
// @Failure 404 member card not found
// @Failure 500 internal server error
// @router /:id/qrcode [get]
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
