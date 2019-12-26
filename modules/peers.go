package modules

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"

	"etm-go-chain/models"
)

func init() {
	event.On("load", event.ListenerFunc(onLoadPeers), event.Normal)
}

type Peers interface {
	GetPeers()
	loadPeers() error
}

type peer struct {
	models.Peer
}

func NewPeers() Peers {
	return &peer{}
}

func (p peer) GetPeers() {
	panic("implement me")
}

func (p peer) loadPeers() error {
	logs.Debug("load peers")
	return nil
}

func onLoadPeers(e event.Event) error {
	logs.Info("onload peers", e)
	err := peers.loadPeers()
	return err
}
