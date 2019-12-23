package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Peer))
}

type iPeer interface {
	GetPeer() (Peer, error)
	SetPeer() error
	GetPeers() ([]Peer, error)
	SetPeers(ps []Peer) error
}

type Peer struct {
	Id      int    `json:"id" orm:"pk;auto"`
	Ip      string `json:"ip"`
	Port    int64  `json:"port"`
	State   int    `json:"state"`
	Version string `json:"version"`
}

func (p Peer) GetPeer() (Peer, error) {
	panic("implement me")
}

func (p Peer) SetPeer() error {
	o := orm.NewOrm()
	_, _, err := o.ReadOrCreate(p, "ip", "port")
	return err
}

func (p Peer) GetPeers() ([]Peer, error) {
	panic("implement me")
}

func (p Peer) SetPeers(ps []Peer) error {
	panic("implement me")
}
