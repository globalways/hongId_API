// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"hongId/models"
	e "github.com/globalways/gws_utils_go/errors"
)

var (
	_ context.Context
	_ beego.Controller
)

// MemberCard API
type MemberCardController struct {
	BaseController
}

// @Title createMemberCards
// @Description generate member card batch
// @Param reqMemberCards body models.ReqNewMemberCards true "request param json"
// @Success 201 {int[]} models.MemberCard.Id
// @Failure 400 request body is invalid
// @Failure 500 generate member card error
// @router / [post]
func (c *MemberCardController) Post() {
	reqMemberCards := new(models.ReqNewMemberCards)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, reqMemberCards); err != nil {
		c.appenWrongParams(models.NewFieldError("reqNewMemberCards", err.Error()))
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	cardIds, gErr := models.GenMemberCards(reqMemberCards, models.Writter)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_DATA_EXIST {
			c.setHttpStatus(http.StatusOK)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutError(gErr))
		return
	}

	// TODO 以一个产生创建活动的 POST 操作为例, 使用一个 HTTP 201 状态代码 然后包含一个 Location header 来指向新生资源的URL。
	c.setHttpHeader("Location", c.combineUrl(beego.UrlFor("MemberCardController.Get", cardIds[0])))
	c.setHttpStatus(http.StatusCreated)
	c.renderJson(cardIds)
}

// @Title getMemberCards
// @Description get member card list by page & size
// @Param page query int64 false "page number"
// @Param size query int64 false "each page count"
// @Success 200 {object} models.MemberCard
// @Failure 403 request url's parameter is invalid
// @Failure 404 request resource not found
// @Failure 500 get memberCard list wrong
// @router / [get]
func (c *MemberCardController) GetAll() {

	var pageNum, pageSize int64
	if page, err := c.GetInt("page"); err != nil {
		c.appenWrongParams(models.NewFieldError("page", err.Error()))
	} else {
		pageNum = page
	}
	if size, err := c.GetInt("size"); err != nil {
		c.appenWrongParams(models.NewFieldError("size", err.Error()))
	} else {
		pageSize = size
	}

	// handle http request param
	if c.handleParamError() {
		return
	}

	pager := &models.Page{
		Size: pageSize,
		CurPage: pageNum,
	}

	cards, gErr := models.FindMemberCard(pager, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutError(gErr))
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
	cardId, err := c.GetInt(":id")
	if err != nil {
		c.appenWrongParams(models.NewFieldError("memberCardId", err.Error()))
	}

	if c.handleParamError() {
		return
	}

	card, gErr := models.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutError(gErr))
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
	cardId, err := c.GetInt(":id")
	if err != nil {
		c.appenWrongParams(models.NewFieldError("memberCardId", err.Error()))
	}

	if c.handleParamError() {
		return
	}

	card, gErr := models.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		if gErr.GetCode() == e.CODE_DB_ERR_NODATA {
			c.setHttpStatus(http.StatusNotFound)
		} else {
			c.setHttpStatus(http.StatusInternalServerError)
		}

		c.renderJson(models.NewCommonOutError(gErr))
		return
	}

	c.renderPng(card.GenQrStream())
}
