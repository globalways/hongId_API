// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/globalways/utils_go/controller"
)

var (
	_         context.Context
	_         logs.BeeLogger
)

type BaseController struct {
	*controller.BasicController
}

