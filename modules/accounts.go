package modules

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
	"workspace/etm-go-chain/models"
)

func init() {
	event.On("load", event.ListenerFunc(onLoadAccounts), event.Normal)
}

type Accounts interface {
	GetAccounts()
}

type account struct {
	models.Account
}

func NewAccounts() Accounts {
	return &account{}
}

func (a account) GetAccounts() {
	panic("implement me")
}

func onLoadAccounts(e event.Event) error {
	logs.Info("onload account", e)
	return nil
}
