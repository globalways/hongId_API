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
	"github.com/globalways/gws_utils_go/page"
)

var (
	_ context.Context
)

// 会员卡API
type MemberCardController struct {
	beego.Controller
}

// @Title createMemberCards
// @Description 批量生成会员卡
// @Param reqMemberCards body models.ReqNewMemberCards true "请求新建会员卡参数"
// @Success 200 {int} models.MemberCard.Id
// @Failure 400 bad request
// @Failure 403 body is empty
// @Failure 500 internal server error
// @router / [post]
func (c *MemberCardController) Post() {
	reqMemberCards := new(models.ReqNewMemberCards)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, reqMemberCards); err != nil {
		c.Ctx.Output.Status = http.StatusBadRequest
		c.StopRun()
	}

	cardIds, gErr := models.GenMemberCards(reqMemberCards, models.Writter)
	if gErr.IsError() {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.StopRun()
	}

	c.Data["json"] = cardIds
	c.ServeJson()
}

// @Title getMemberCards
// @Description 获取会员卡号列表
// @Param page query int64 false "分页page"
// @Param size query int64 false "分页size"
// @Success 200 {int} models.MemberCard
// @Failure 500 internal server error
// @router / [get]
func (c *MemberCardController) GetAll() {

	pageNum, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("size")
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	pager := &page.Page{
		Perpage:      pageSize,
		Current_page: pageNum,
	}

	cards, gErr := models.FindMemberCard(pager, models.Reader)
	if gErr.IsError() {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.StopRun()
	}

	c.Data["json"] = cards
	c.ServeJson()
}

// @Title getMemberCard
// @Description 通过ID获取会员卡信息
// @Param	cardId		path 	int	true		"会员卡ID"
// @Success 200 {int} models.MemberCard
// @Failure 400 bad request
// @Failure 500 internal server error
// @router /:cardId [get]
func (c *MemberCardController) Get() {
	cardId, err := c.GetInt(":cardId")
	if err != nil {
		c.Ctx.Output.Status = http.StatusBadRequest
		c.StopRun()
	}

	card, gErr := models.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.StopRun()
	}

	c.Data["json"] = card
	c.ServeJson()
}

// @Title getMemberCardQrCode
// @Description 通过ID获取会员卡二维码
// @Param	cardId		path 	int	true		"会员卡ID"
// @Success 200 {image/png} qrCode
// @Failure 400 bad request
// @Failure 500 internal server error
// @router /:cardId/qrcode [get]
func (c *MemberCardController) GetQrCode() {
	cardId, err := c.GetInt(":cardId")
	if err != nil {
		c.Ctx.Output.Status = http.StatusBadRequest
		c.StopRun()
	}

	card, gErr := models.GetMemberCardById(cardId, models.Reader)
	if gErr.IsError() {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.StopRun()
	}

	qrBytes := card.GenQrStream()
	c.Ctx.Output.Body(qrBytes)
	c.Ctx.Output.ContentType("image/png")
}
