// @APIVersion 1.0.0
// @Title EnTanMo API
// @Description The Entanmo Blockchain API 
// @Contact developer@entanmo.com
// @TermsOfServiceUrl http://www.entanmo.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"workspace/etm-go-chain/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/blocks",
			beego.NSInclude(
				&controllers.BlockController{},
			),
		),
		beego.NSNamespace("/transactions",
			beego.NSInclude(
				&controllers.TransactionController{},
			),
		),
		beego.NSNamespace("/peers",
			beego.NSInclude(
				&controllers.PeerController{},
			),
		),
		beego.NSNamespace("/system",
			beego.NSInclude(
				&controllers.SystemController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
