// @APIVersion 1.0.0
// @Title 环途国际会员系统 API
// @Description 会员API，用于GlobalWays会员系统构建
// @Contact mint.zhao.chiu@gmail.com
// @TermsOfServiceUrl http://www.globalways.cn/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"hongID/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/memberCards",
			beego.NSInclude(
				&controllers.MemberCardController{},
			),
		),
		beego.NSNamespace("/members",
			beego.NSInclude(
				&controllers.MemberController{},
			),
		),
		beego.NSNamespace("/memberGroups",
			beego.NSInclude(
				&controllers.MemberGroupController{},
			),
		),
		beego.NSNamespace("/orders",
			beego.NSInclude(
				&controllers.OrderController{},
				&controllers.OrderAddressController{},
			),
		),
		beego.NSNamespace("/stores",
			beego.NSInclude(
				&controllers.StoreAdminController{},
				&controllers.StoreController{},
				&controllers.StoreIndustryController{},
				&controllers.ProductTagController{},
				&controllers.StoreProductController{},
				&controllers.PurchaseChannelController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
