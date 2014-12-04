// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import "github.com/globalways/utils_go/errors"

type CommonResponse struct {
	Code        int    `json:code`
	Message     string `json:message`
	Description string `jsong:description`
}

type FieldError struct {
	Field   string `json:field`
	Message string `json:message`
}

type FieldErrors struct {
	Code    int           `json:code`
	Message string        `json:message`
	Errors  []*FieldError `json:errors`
}

// new common response
func NewCommonOutRsp(gErr errors.GlobalWaysError) *CommonResponse {
	code := gErr.GetCode()
	msg := gErr.GetMessage()
	desc := gErr.GetInner().Error()

	return &CommonResponse{
		Code:        code,
		Message:     msg,
		Description: desc,
	}
}

// new fielderror
func NewFieldError(field string, msg string) *FieldError {
	return &FieldError{
		Field:   field,
		Message: msg,
	}
}

// new fieldErrors
func NewFiledErrors(code int, errs []*FieldError) *FieldErrors {
	return &FieldErrors{
		Code:    code,
		Message: errors.GetCodeMessage(code),
		Errors:  errs,
	}
}
