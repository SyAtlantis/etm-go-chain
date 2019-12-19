package modules

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
)

func init() {
	event.On("load", event.ListenerFunc(onLoadSystem), event.Normal)
	event.On("ready", event.ListenerFunc(onReadySystem), event.Normal)
}

type System interface {
	GetVersion() string
}

type sys struct {
}

func (s sys) GetVersion() string {
	panic("implement me")
}

func NewSystem() System {
	return &sys{}
}

func onLoadSystem(e event.Event) error {
	logs.Info("onload system", e)
	return nil
}

func onReadySystem(e event.Event) error {
	logs.Info("onReady system", e)
	return nil
}
