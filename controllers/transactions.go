package controllers

import "github.com/astaxie/beego"

// Operations about transactions
type TransactionController struct {
	beego.Controller
}

// @Title getTransactions
// @Description Transactions that meet the conditions
// @Success 200 {Array} transactions
// @Failure 403 transactions not exist
// @router / [get]
func (t *TransactionController) GetTransactions() {
	t.Data["json"] = "00"
	t.ServeJSON()
}