package system

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
)

func init() {
	//var event = core.NewEvent()
	//ch1 := make(chan core.DataEvent)
	//ch2 := make(chan core.DataEvent)
	//event.Subscribe("load", ch1)
	//event.Subscribe("ready", ch2)
	//
	//go func() {
	//	for {
	//		select {
	//		case d := <-ch1:
	//			go onLoad(d)
	//		case d := <-ch2:
	//			go onReady(d)
	//		}
	//	}
	//}()

	event.On("load", event.ListenerFunc(onLoad), event.Normal)
	event.On("ready", event.ListenerFunc(onReady), event.High)

}

func onLoad(e event.Event) error {
	logs.Info("onload system", e)
	return nil
}

func onReady(e event.Event) error {
	logs.Info("onReady system", e)
	return nil
}
