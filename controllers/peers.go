package controllers

import "github.com/astaxie/beego"

// Operations about peers
type PeerController struct {
	beego.Controller
}

// @Title getPeers
// @Description The peers of mine
// @Success 200 {array} peers
// @Failure 403 peers not exist
// @router / [get]
func (p *PeerController) GetPeers() {
	p.Data["json"] = "getPeers"
	p.ServeJSON()
}
