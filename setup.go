package main

import (
	"github.com/gookit/event"
)

//var models = make(map[string]Models)
//
//type Models interface {
//	NewModel() interface{}
//}
//
//func RegisterModels(name string, model Models) {
//	models[name] = model
//}

func Setup() {

	_, _ = event.Fire("load", event.M{"name": "bbb"})
	//time.Sleep(10*time.Second)
	_, _ = event.Fire("ready", event.M{"name": "ccc"})

	//block := blocks.NewBlock()
	//block.VerifyBlock()
	//logs.Info(block)
}
