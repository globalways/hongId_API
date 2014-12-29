package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"RegisterByTel",
			`/register/tel`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"UpdateALL",
			`/id/:memberId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"Update",
			`/id/:memberId`,
			[]string{"patch"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"GetByTel",
			`/tel/:tel`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"GetById",
			`/id/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"SysGenMembers",
			`/`,
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
			`/id/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetQrCode",
			`/id/:id/qr`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"BindCard",
			`/card/:card/bind/:owner`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"UnBindCard",
			`/card/:card/unbind`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"NewGroup",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetGroupALL",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetGroup",
			`/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"UpdateALL",
			`/:groupId`,
			[]string{"put"},
			nil})

}
