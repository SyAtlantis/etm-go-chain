package blocks

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
)

func init() {
	event.On("bind", event.ListenerFunc(onBind), event.Normal)
}

func onBind(e event.Event) error {
	logs.Info("onBind block", e.Data())
	return nil
}