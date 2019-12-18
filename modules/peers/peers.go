package peers

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
	"workspace/etm-go-chain/models"
)

func init() {
	m := peer{}
	models.RegisterModels("peers", &m)

	event.On("load", event.ListenerFunc(onLoad), event.Normal)
}

type Peers interface {
	GetPeers()
}

type peer struct {
	models.Peer
}

func (p peer) NewModel() interface{} {
	return &peer{}
}

func (p peer) GetPeers() {
	panic("implement me")
}

func onLoad(e event.Event) error {
	logs.Info("onload peers", e)
	return nil
}
