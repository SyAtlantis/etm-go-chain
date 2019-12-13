package controllers

import "github.com/astaxie/beego"

// Operations about accounts
type AccountController struct {
	beego.Controller
}


// @Title getAccounts
// @Description Accounts that meet the conditions
// @Success 200 {Array} accounts
// @Failure 403 accounts not exist
// @router / [get]
func (a *AccountController) GetAccounts() {
	a.Data["json"] = "GetAccounts"
	a.ServeJSON()
}

// @Title open
// @Description account login
// @Success 200 {object} models.Object
// @Failure 403 account not exist
// @router /open [post]
func (a *AccountController) Open() {
	a.Data["json"] = "height 0"
	a.ServeJSON()
}
