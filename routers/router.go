// @APIVersion 1.0.0
// @Title 环途国际会员卡系统 API
// @Description 会员卡API，用于GlobalWays会员卡系统构建
// @Contact mint.zhao.chiu@gmail.com
// @TermsOfServiceUrl http://www.globalways.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hongId/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
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
	beego.AddNamespace(ns)
}
