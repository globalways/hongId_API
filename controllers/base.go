// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
	"hongId/models"
	"github.com/astaxie/beego/context"
	"net/http"
	"github.com/globalways/gws_utils_go/errors"
)

var (
	_ context.Context
)

type BaseController struct {
	beego.Controller
	fieldErrors []*models.FieldError
}

// before exec logic, prepare something
func (c *BaseController) Prepare() {
	c.fieldErrors = make([]*models.FieldError, 0)
}

// after exec logic, finish something
func (c *BaseController) Finish() {
	c.fieldErrors = c.fieldErrors[:0]
}

// parse params is wrong, if wrong, fill response with errors
func (c *BaseController) isParamsWrong() bool {
	return len(c.fieldErrors) != 0
}

// append a new parameter wrong info
func (c *BaseController) appenWrongParams(err *models.FieldError) {
	c.fieldErrors = append(c.fieldErrors, err)
}

// http json response
func (c *BaseController) renderJson(data interface {}) {
	c.Data["json"] = data
	c.ServeJson()
}

// http png response
func (c *BaseController) renderPng(data []byte) {
	c.setHttpContentType("image/png")
	c.setHttpBody(data)
}

// set http status
func (c *BaseController) setHttpStatus(status int) {
	c.Ctx.Output.SetStatus(status)
}

// set http response header
func (c *BaseController) setHttpHeader(key, val string) {
	c.Ctx.Output.Header(key, val)
}

// set http response body
func (c *BaseController) setHttpBody(body []byte) {
	c.Ctx.Output.Body(body)
}

// get http request body
func (c *BaseController) getHttpBody() []byte {
	return c.Ctx.Input.RequestBody
}

// set http response contenttype
func (c *BaseController) setHttpContentType(ext string) {
	c.Ctx.Output.ContentType(ext)
}

// combine url
func (c *BaseController) combineUrl(router string) string {
	return c.Ctx.Input.Site() + router
}

// handle http request param error
func (c *BaseController) handleParamError() bool {
	if c.isParamsWrong() {
		c.setHttpStatus(http.StatusBadRequest)
		c.renderJson(models.NewFiledErrors(errors.CODE_HTTP_ERR_INVALID_PARAMS, c.fieldErrors))

		return true
	}

	return false
}
