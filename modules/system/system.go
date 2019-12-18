package system

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
	"workspace/etm-go-chain/models"
)

func init() {
	m := system{}
	models.RegisterModels("system", &m)

	event.On("load", event.ListenerFunc(onLoad), event.Normal)
	event.On("ready", event.ListenerFunc(onReady), event.High)

}

type system struct {
}

func (s system) NewModel() interface{} {
	return &system{}
}

func onLoad(e event.Event) error {
	logs.Info("onload system", e)
	return nil
}

func onReady(e event.Event) error {
	logs.Info("onReady system", e)
	return nil
}
