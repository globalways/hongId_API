package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"],
		beego.ControllerComments{
			"NewAddress",
			`/addresses`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"],
		beego.ControllerComments{
			"DelAddress",
			`/addresses`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"],
		beego.ControllerComments{
			"UpdateAddress",
			`/addresses`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderAddressController"],
		beego.ControllerComments{
			"ListAddress",
			`/addresses`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberCardController"],
		beego.ControllerComments{
			"GenCard",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreController"],
		beego.ControllerComments{
			"NewStore",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreController"],
		beego.ControllerComments{
			"DeleteStore",
			`/:storeid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreController"],
		beego.ControllerComments{
			"UpdateStore",
			`/:storeid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreController"],
		beego.ControllerComments{
			"GetStore",
			`/:storeid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreController"],
		beego.ControllerComments{
			"GetStoresByAdmin",
			`/a/:adminid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreController"],
		beego.ControllerComments{
			"GetStores",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"],
		beego.ControllerComments{
			"NewStoreAdmin",
			`/admins`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"],
		beego.ControllerComments{
			"DeleteStoreAdmin",
			`/admins/:adminid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"],
		beego.ControllerComments{
			"UpdateStoreAdmin",
			`/admins/:adminid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreAdminController"],
		beego.ControllerComments{
			"GetStoreAdmin",
			`/admins/:adminid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberController"],
		beego.ControllerComments{
			"Update",
			`/s`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberController"],
		beego.ControllerComments{
			"GetUnUsed",
			`/u`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberController"],
		beego.ControllerComments{
			"Get",
			`/s`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberController"],
		beego.ControllerComments{
			"SysGenMembers",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreProductController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreProductController"],
		beego.ControllerComments{
			"NewStoreProduct",
			`/products`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreProductController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreProductController"],
		beego.ControllerComments{
			"UpdateStoreProduct",
			`/products/:pid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:PurchaseChannelController"] = append(beego.GlobalControllerRouter["hongID/controllers:PurchaseChannelController"],
		beego.ControllerComments{
			"NewPurchaseChannel",
			`/purchaseChannels`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:PurchaseChannelController"] = append(beego.GlobalControllerRouter["hongID/controllers:PurchaseChannelController"],
		beego.ControllerComments{
			"UpdatePurchaseChannel",
			`/purchaseChannels/:pid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderController"],
		beego.ControllerComments{
			"NewOrder",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderController"],
		beego.ControllerComments{
			"OrderInfo",
			`/:orderid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderController"],
		beego.ControllerComments{
			"BuyerOrders",
			`/buyer/:buyer`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderController"],
		beego.ControllerComments{
			"StoreOrders",
			`/stores/:storeid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderController"],
		beego.ControllerComments{
			"NewProcess",
			`/:orderid/processes`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:OrderController"] = append(beego.GlobalControllerRouter["hongID/controllers:OrderController"],
		beego.ControllerComments{
			"OrderProcesses",
			`/:orderid/processes`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:ProductTagController"] = append(beego.GlobalControllerRouter["hongID/controllers:ProductTagController"],
		beego.ControllerComments{
			"NewProductTag",
			`/products/tags`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:ProductTagController"] = append(beego.GlobalControllerRouter["hongID/controllers:ProductTagController"],
		beego.ControllerComments{
			"DeleteProductTag",
			`/products/tags/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:ProductTagController"] = append(beego.GlobalControllerRouter["hongID/controllers:ProductTagController"],
		beego.ControllerComments{
			"UpdateProductTag",
			`/products/tags/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:ProductTagController"] = append(beego.GlobalControllerRouter["hongID/controllers:ProductTagController"],
		beego.ControllerComments{
			"GetTags",
			`/products/tags`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:ProductTagController"] = append(beego.GlobalControllerRouter["hongID/controllers:ProductTagController"],
		beego.ControllerComments{
			"GetTagsByPage",
			`/products/tags/p`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"],
		beego.ControllerComments{
			"NewIndustry",
			`/industries`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"],
		beego.ControllerComments{
			"DeleteIndustry",
			`/industries/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"],
		beego.ControllerComments{
			"UpdateIndustry",
			`/industries/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"] = append(beego.GlobalControllerRouter["hongID/controllers:StoreIndustryController"],
		beego.ControllerComments{
			"GetIndustries",
			`/industries`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetGroupALL",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"],
		beego.ControllerComments{
			"NewGroup",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"],
		beego.ControllerComments{
			"GetGroup",
			`/id/:groupId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"] = append(beego.GlobalControllerRouter["hongID/controllers:MemberGroupController"],
		beego.ControllerComments{
			"UpdateALL",
			`/id/:groupId`,
			[]string{"put"},
			nil})

}
