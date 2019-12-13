package peers

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
)

func init() {
	event.On("load", event.ListenerFunc(onLoad), event.Normal)
}

func onLoad(e event.Event) error {
	logs.Info("onload peers", e)
	return nil
}
