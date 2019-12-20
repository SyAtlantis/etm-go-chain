package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Peer))
}

type Peer struct {
	Id      int    `json:"id" orm:"pk;auto"`
	Ip      string `json:"ip"`
	Port    int64  `json:"port"`
	State   int    `json:"state"`
	Version string `json:"version"`
}
