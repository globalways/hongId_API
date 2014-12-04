// @APIVersion 1.0.0
// @Title 环途国际会员系统 API
// @Description 会员API，用于GlobalWays会员系统构建
// @Contact mint.zhao.chiu@gmail.com
// @TermsOfServiceUrl http://www.globalways.cn/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hongId/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 会员卡相关router
	memberCard := beego.NewNamespace("/v1",
		beego.NSNamespace("/memberCards",
			beego.NSInclude(
				&controllers.MemberCardController{},
			),
		),
		beego.NSNamespace("/channelTypes",
			beego.NSInclude(
				&controllers.ChannelTypeController{},
			),
		),
	)
	beego.AddNamespace(memberCard)

	// 会员相关router
	member := beego.NewNamespace("/v1",
		beego.NSNamespace("/hongId",
			beego.NSInclude(
				&controllers.MemberController{},
			),
		),
	)
	beego.AddNamespace(member)
}
