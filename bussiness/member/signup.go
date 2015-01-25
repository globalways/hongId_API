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
package member

import (
	"github.com/globalways/errors"
	"github.com/globalways/utils_go/security"
	"hongID/models"
)

const (
	err_code_member_base           int = -100
	err_code_member_tel_already_in     = err_code_member_base - 1
	err_code_member_no_more            = err_code_member_base - 2
	err_code_member_sign_up            = err_code_member_base - 3
)

var (
	errMsgs = map[int]string{
		err_code_member_tel_already_in: "该手机号已被注册.",
		err_code_member_no_more:        "没有更多会员资源可供使用,请联系系统管理员.",
		err_code_member_sign_up:        "会员注册失败.",
	}
)

type Member struct{}

// 注册会员 tel
// @param tel 手机号
// @param nick 昵称
// @param password 密码
// @param groupid 会员组id
// @return error
func (b Member) SignUpMemberByTel(tel, nick, password string, groupid int64) errors.GlobalWaysError {
	// 会员是否存在
	if models.HasTel(tel, groupid) {
		return errors.New(err_code_member_tel_already_in, errMsgs[err_code_member_tel_already_in])
	}

	// 返回未使用的会员
	member := models.GetUnUsedMember(groupid, []string{"id"})
	if member == nil {
		return errors.New(err_code_member_no_more, errMsgs[err_code_member_no_more])
	}

	// 注册会员
	params := map[string]interface{}{
		"tel":       tel,
		"nick_name": nick,
		"password":  security.GenerateFromPassword(password),
		"status":    models.EMemberStatus_Use,
	}
	if !models.UpdateMemberById(member.Id, params) {
		return errors.New(err_code_member_sign_up, errMsgs[err_code_member_sign_up])
	}

	return errors.ErrorOK()
}

func Test() errors.GlobalWaysError {
	return errors.ErrorOK()
}
