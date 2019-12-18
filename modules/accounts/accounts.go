package accounts

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
	"workspace/etm-go-chain/models"
)

func init() {
	m := account{}
	models.RegisterModels("account", &m)

	event.On("load", event.ListenerFunc(onLoad), event.Normal)
}

type Accounts interface {
	GetAccounts()
}

type account struct {
	models.Account
}

func (a *account) NewModel() interface{} {
	return &account{}
}

func (a account) GetAccounts() {
	panic("implement me")
}

func onLoad(e event.Event) error {
	logs.Info("onload account", e)
	return nil
}
