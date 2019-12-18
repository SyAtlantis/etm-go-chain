package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:AccountController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetAccounts",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:AccountController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Open",
            Router: `/open`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:BlockController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:BlockController"],
        beego.ControllerComments{
            Method: "GetBlocks",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:PeerController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:PeerController"],
        beego.ControllerComments{
            Method: "GetPeers",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetHeight",
            Router: `/getHeight`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetMilestone",
            Router: `/getMilestone`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetReward",
            Router: `/getReward`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetStatus",
            Router: `/getStatus`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:TransactionController"] = append(beego.GlobalControllerRouter["workspace/etm-go-chain/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "GetTransactions",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
