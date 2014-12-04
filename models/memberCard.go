// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	e "errors"
	"time"
	"github.com/globalways/utils_go/algorith"
	"github.com/globalways/utils_go/qr"
	errors "github.com/globalways/utils_go/errors"
)

const (
	_DEFAULT_INSERT_ERR_CNT = 3
)

const (
	EMemberCardStatus_Inactive byte = iota + 1
	EMemberCardStatus_Active
	EMemberCardStatus_Expired
)

type MemberCard struct {
	Id          int64
	MII         byte      `orm:"default(6);column(mii)"`  // 1     主要产业标识符（Major Industry Identifier (MII)）
	CPI         byte      `orm:"default(32);column(cpi)"` // 2-3   公司标识符，默认: 32
	CDI         uint16    `orm:"default(86);column(cdi)"` // 4-6   国家域标识符（Country Domain Identifier）
	PII         uint64    `orm:"column(pii)"`             // 7-18  个人信息标识（Personal identifying information）
	IVC         byte      `orm:"default(0);column(ivc)"`  // 19    验证码标识（Identity verification code）
	ChannelType int64     `orm:"default(0);column(channel_type)"`
	ChannelId   int64     `orm:"default(0);column(channel_id)"`
	CardStatus  byte      `orm:"column(card_status)"`
	Created     time.Time `orm:"column(created);auto_now_add"`
	Updated     time.Time `orm:"column(updated);auto_now"`
	Member     *Member   `orm:"column(member);rel(fk);null;on_delete(set_null)"`
}

func (c *MemberCard) TableName() string {
	return "member_card"
}

// 输出string
func (c *MemberCard) String() string {
	return fmt.Sprintf("%v%v%.*d%.*d%v", c.MII, c.CPI, 3, c.CDI, 12, c.PII, c.IVC)
}

func (c *MemberCard) genCardIVC() {
	c.IVC = algorith.GenLuhnCheckDigit([]byte(fmt.Sprintf("%v%v%.*d%.*d", c.MII, c.CPI, 3, c.CDI, 12, c.PII)))
}

// 验证会员卡号的正确性
func (c *MemberCard) ValidateCard() bool {
	return algorith.ValidateLuhn(string(c.String()))
}

// 生成二维码png二进制流
func (c *MemberCard) GenQrStream() []byte {
	return qr.GenQRCode(c.String(), qr.H)
}

// 新增会员卡
func genMemberCard(card *MemberCard, ormer orm.Ormer) (bool, errors.GlobalWaysError) {

	if isCardExist(card, ormer) {
		return false, errors.New(errors.CODE_DB_DATA_EXIST)
	}

	if _, err := ormer.Insert(card); err != nil {
		return false, errors.Wrap(errors.CODE_DB_ERR_INSERT, err)
	}

	return true, errors.ErrorOK()
}

// 判断会员卡是否存在
func isCardExist(card *MemberCard, ormer orm.Ormer) bool {
	return ormer.QueryTable(new(MemberCard)).Filter("mii", card.MII).Filter("cpi", card.CPI).Filter("cdi", card.CDI).Filter("pii", card.PII).Filter("ivc", card.IVC).Exist()
}

//批量生成
func GenMemberCards(reqMsg *MemberCard, cardCnt int64, ormer orm.Ormer) ([]string, errors.GlobalWaysError) {

	cardNumbers := make([]string, 0)

	//判断channelType是否存在
	if !isChannelExist(&ChannelType{Id: reqMsg.ChannelType}, ormer) {
		return cardNumbers, errors.Wrap(errors.CODE_DB_ERR_NODATA, e.New("invalid channelType."))
	}

	//获取数据库中最大的PII
	maxPii, _ := getMaxPii(reqMsg.MII, reqMsg.CPI, reqMsg.CDI, ormer)

	affactedTotal := int64(1)
	maxPii += 1
	errCnt := 0
	for affactedTotal <= cardCnt {
		memberCard := &MemberCard{
			MII:         reqMsg.MII,
			CPI:         reqMsg.CPI,
			CDI:         reqMsg.CDI,
			PII:         maxPii,
			ChannelType: reqMsg.ChannelType,
			ChannelId:   reqMsg.ChannelId,
			CardStatus:  EMemberCardStatus_Inactive,
		}
		memberCard.genCardIVC()

		if flag, err := genMemberCard(memberCard, ormer); !flag || err.IsError() {

			errCnt++
			if errCnt >= _DEFAULT_INSERT_ERR_CNT {
				return cardNumbers, err
			}

			continue
		}

		affactedTotal++
		maxPii++

		cardNumbers = append(cardNumbers, memberCard.String())
	}

	return cardNumbers, errors.ErrorOK()
}

// 获取数据库中现存最大的PII
func getMaxPii(mii, cpi byte, cdi uint16, ormer orm.Ormer) (uint64, errors.GlobalWaysError) {

	memberCard := new(MemberCard)

	defer func() {
		println(memberCard.PII)
	}()

	if err := ormer.QueryTable(memberCard).Filter("mii", mii).Filter("cpi", cpi).Filter("cdi", cdi).OrderBy("-pii").One(memberCard, "pii"); err != nil {
		if err != orm.ErrMultiRows {
			return 0, errors.Wrap(errors.CODE_DB_ERR_GET, err)
		}
	}

	return memberCard.PII, errors.ErrorOK()
}

// 获取会员卡列表
func FindMemberCard(pager *Page, ormer orm.Ormer) ([]*MemberCard, errors.GlobalWaysError) {
	memberCardList := make([]*MemberCard, 0)

	if num, err := ormer.QueryTable(new(MemberCard)).Limit(pager.Size, (pager.CurPage-1)*pager.Size).All(&memberCardList); err != nil {
		return memberCardList, errors.Wrap(errors.CODE_DB_ERR_FIND, err)
	} else if num == 0 {
		return memberCardList, errors.New(errors.CODE_DB_ERR_NODATA)
	}

	return memberCardList, errors.ErrorOK()
}

// 获取特定ID会员卡
func GetMemberCardById(id int64, ormer orm.Ormer) (*MemberCard, errors.GlobalWaysError) {

	tmpCard := &MemberCard{
		Id: id,
	}

	err := ormer.Read(tmpCard)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New(errors.CODE_DB_ERR_NODATA)
		} else {
			return nil, errors.Wrap(errors.CODE_DB_ERR_GET, err)
		}
	}

	return tmpCard, errors.ErrorOK()
}
