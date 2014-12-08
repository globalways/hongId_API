// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/smsmgr"
	"net/http"
	"github.com/astaxie/beego/logs"
)

var (
	_     context.Context
	_     logs.BeeLogger
	valid = new(validation.Validation)
	sms   = smsmgr.NewDefaultSmsManager()
)

type BaseController struct {
	beego.Controller
	fieldErrors []*errors.FieldError
}

// before exec logic, prepare something
func (c *BaseController) Prepare() {
	c.fieldErrors = make([]*errors.FieldError, 0)
	valid.Clear()

	//prepare for enable gzip
	c.Ctx.Output.EnableGzip = true

	// handle schema error
	//	c.handleConnSchemaError()
}

// api just only allow https connection, if not, throw errors
func (c *BaseController) handleConnSchemaError() {
	if !c.Ctx.Input.IsSecure() {
		c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_HTTP_ERR_NOT_HTTPS)))
	}
}

// after exec logic, finish something
func (c *BaseController) Finish() {
	c.fieldErrors = c.fieldErrors[:0]
}

// http json response
func (c *BaseController) renderJson(data interface{}) {
	c.Data["json"] = data
	c.ServeJson()
}

// http png response
func (c *BaseController) renderPng(data []byte) {
	c.Ctx.Output.EnableGzip = false
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
	beego.BeeLogger.Debug("httpBody: %v", c.Ctx.Input.RequestBody)
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
		c.renderJson(errors.NewFiledErrors(errors.CODE_HTTP_ERR_INVALID_PARAMS, c.fieldErrors))

		for _, err := range c.fieldErrors {
			beego.BeeLogger.Debug("filedError: %v", err)
		}

		return true
	}

	return false
}

// parse params is wrong, if wrong, fill response with errors
func (c *BaseController) isParamsWrong() bool {
	return len(c.fieldErrors) != 0
}

// append a new parameter wrong info
func (c *BaseController) appenWrongParams(err *errors.FieldError) {
	c.fieldErrors = append(c.fieldErrors, err)
}

// valid paramemter
func (c *BaseController) validation(obj interface{}) {
	b, err := valid.Valid(obj)
	if err != nil {
		c.appenWrongParams(errors.NewFieldError("valid", err.Error()))
	}

	if !b {
		for _, err := range valid.Errors {
			c.appenWrongParams(errors.NewFieldError(err.Key, err.Message))
		}
	}
}

// generate sms auth code
func (c *BaseController) genSmsAuthCode(tel string) (string, error) {
	code, err := sms.GenSmsAuthCode(tel)
	beego.BeeLogger.Debug("generate sms auth code: %v", code)
	return code, err
}
