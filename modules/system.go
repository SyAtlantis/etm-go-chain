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
	GetLastHeight() int64
	LoadBlockChain() error
}

type sys struct {
	Version    string
	LastHeight int64
}

func NewSystem() System {
	return &sys{}
}

func (s *sys) GetVersion() string {
	panic("implement me")
}

func (s *sys) GetLastHeight() int64 {
	panic("implement me")
}

func (s *sys) LoadBlockChain() error {
	logs.Debug("load block chain")
	return nil
}

func onLoadSystem(e event.Event) error {
	err := system.LoadBlockChain()
	return err
}

func onReadySystem(e event.Event) error {
	logs.Info("onReady system", e)
	return nil
}
