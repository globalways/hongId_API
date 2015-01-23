package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["hongId/controllers:StoreIndustryController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreIndustryController"],
		beego.ControllerComments{
			"NewIndustry",
			`/industries`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreIndustryController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreIndustryController"],
		beego.ControllerComments{
			"DeleteIndustry",
			`/industries/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"],
		beego.ControllerComments{
			"NewStoreAdmin",
			`/admins`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"],
		beego.ControllerComments{
			"DeleteStoreAdmin",
			`/admins/:adminid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"],
		beego.ControllerComments{
			"UpdateStoreAdmin",
			`/admins/:adminid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreAdminController"],
		beego.ControllerComments{
			"GetStoreAdmin",
			`/admins/:adminid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"Update",
			`/s`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"GetUnUsed",
			`/u`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberController"],
		beego.ControllerComments{
			"Get",
			`/s`,
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
			"GenCard",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetGroupALL",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"NewGroup",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetGroup",
			`/id/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongId/controllers:MemberGroupController"],
		beego.ControllerComments{
			"UpdateALL",
			`/id/:groupId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"],
		beego.ControllerComments{
			"NewAddress",
			`/addresses`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"],
		beego.ControllerComments{
			"DelAddress",
			`/addresses`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"],
		beego.ControllerComments{
			"UpdateAddress",
			`/addresses`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderAddressController"],
		beego.ControllerComments{
			"ListAddress",
			`/addresses`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreController"],
		beego.ControllerComments{
			"NewStore",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreController"],
		beego.ControllerComments{
			"DeleteStore",
			`/:storeid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreController"],
		beego.ControllerComments{
			"UpdateStore",
			`/:storeid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreController"],
		beego.ControllerComments{
			"GetStore",
			`/:storeid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongId/controllers:StoreController"],
		beego.ControllerComments{
			"GetStores",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderController"],
		beego.ControllerComments{
			"NewOrder",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderController"],
		beego.ControllerComments{
			"OrderInfo",
			`/:orderid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderController"],
		beego.ControllerComments{
			"BuyerOrders",
			`/buyer/:buyer`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderController"],
		beego.ControllerComments{
			"StoreOrders",
			`/stores/:storeid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderController"],
		beego.ControllerComments{
			"NewProcess",
			`/:orderid/processes`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongId/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongId/controllers:OrderController"],
		beego.ControllerComments{
			"OrderProcesses",
			`/:orderid/processes`,
			[]string{"get"},
			nil})

}
