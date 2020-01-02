package modules

import (
	"etm-go-chain/models"
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
)

func init() {
	event.On("load", event.ListenerFunc(onLoadSystem), event.Normal)
	event.On("ready", event.ListenerFunc(onReadySystem), event.Normal)
}

type Systems interface {
	GetVersion() string
	GetLastHeight() int64
	SetLastHeight(int64) error
	LoadBlockChain() error
}

type system struct {
	models.System
}

func NewSystems() Systems {
	return &system{}
}

func (s *system) GetVersion() string {
	panic("implement me")
}

func (s *system) GetLastHeight() int64 {
	return s.LastHeight
}

func (s *system) SetLastHeight(h int64) error {
	s.LastHeight = h
	return nil
}

func (s *system) LoadBlockChain() error {
	logs.Debug("load block chain")
	// clear tables accounts
	if err := accounts.RemoveTables(); err != nil {
		return err
	}

	// load blocks offset
	if err := blocks.loadBlocksOffset(0, 1000); err != nil {
		return err
	}
	return nil
}

func onLoadSystem(e event.Event) error {
	err := systems.LoadBlockChain()
	return err
}

func onReadySystem(e event.Event) error {
	logs.Info("onReady system", e)
	return nil
}
