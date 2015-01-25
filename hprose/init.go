// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package hprose

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-go/hprose"
	"fmt"
	"hongID/bussiness/member"
)

var (
	hprose_addr = beego.AppConfig.DefaultString("hprose::hprose_addr", "127.0.0.1")
	hprose_port = beego.AppConfig.DefaultString("hprose::hprose_port", "1234")
)

// start hongid rpc service
func init() {
	server := hprose.NewTcpServer(fmt.Sprintf("tcp://%v:%v", hprose_addr, hprose_port))
	server.AddMethods(member.Member{}, hprose.Raw)
	server.AddFunction("test", member.Test)
	if err := server.Start(); err != nil {
		panic(err)
	}

	beego.BeeLogger.Trace("hprose member service started.")
}
