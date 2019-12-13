package main

import (
	"github.com/gookit/event"
)

func Setup() {
	_, _ = event.Fire("bind", event.M{"name": "aaa"})
	_, _ = event.Fire("load", event.M{"name": "bbb"})
	//time.Sleep(10*time.Second)
	_, _ = event.Fire("ready", event.M{"name": "ccc"})
}
