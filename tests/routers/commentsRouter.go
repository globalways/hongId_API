package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"TelAuthCode",
			`/telAuthCode`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetQrCode",
			`/:id/qrcode`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"Get",
			`/:channelId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["hongId/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"Put",
			`/:channelId`,
			[]string{"put"},
			nil})

}
