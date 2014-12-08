package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"Get",
			`/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"Put",
			`/:groupId`,
			[]string{"put"},
			nil})

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

}
