package controllers

import "github.com/astaxie/beego"

// Operations about blocks
type BlockController struct {
	beego.Controller
}

// @Title getBlocks
// @Description Blocks that meet the conditions
// @Success 200 {Array} blocks
// @Failure 403 blocks not exist
// @router / [get]
func (b *BlockController) GetBlocks() {
	b.Data["json"] = "GetBlocks"
	b.ServeJSON()
}