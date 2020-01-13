package modules

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"

	"etm-go-chain/models"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindPeers), event.Normal)
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
	// TODO load peers
	logs.Warn("TODO load peers!")

	return nil
}

func onBindPeers(e event.Event) error {
	logs.Notice("【onBind】 peers", e)
	err := peers.loadPeers()
	return err
}
