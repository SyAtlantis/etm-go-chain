package models

import (
	"github.com/astaxie/beego/orm"

	_ "workspace/etm-go-chain/models/accounts"
	_ "workspace/etm-go-chain/models/blocks"
	_ "workspace/etm-go-chain/models/peers"
	_ "workspace/etm-go-chain/models/system"
	_ "workspace/etm-go-chain/models/transactions"
)

type Block struct {
	Key                  int    `orm:"pk;auto;column(key)"`
	Id                   string `json:"id"`
	Height               int64  `json:"height"`
	Timestamp            int64  `json:"timestamp"`
	TotalAmount          int64  `json:"totalAmount"`
	TotalFee             int64  `json:"totalFee"`
	Reward               int64  `json:"reward"`
	PayloadHash          string `json:"payloadHash"`
	PayloadLength        int    `json:"payloadLength"`
	NumberOfTransactions int    `json:"numberOfTransactions"`
	PreviousBlock        string `json:"previousBlock"`
	GeneratorPublicKey   string `json:"generatorPublicKey"`
	BlockSignature       string `json:"blockSignature"`
	//Transactions         []Transaction `json:"transactions"`
}

//
//
//type Transaction struct {
//	Type               TrType   `json:"type"`
//	Id                 string   `json:"id"`
//	Fee                int64    `json:"fee"`
//	Amount             int64    `json:"amount"`
//	Timestamp          int64    `json:"timestamp"`
//	RecipientId        string   `json:"recipientId"`
//	Asset              Asset    `json:"asset"`
//	Args               []string `json:"args"`
//	Message            string   `json:"message"`
//	Signature          string   `json:"signature"`
//	SignSignature      string   `json:"signSignature"`
//	SenderPublicKey    string   `json:"senderPublicKey"`
//	RequesterPublicKey string   `json:"requesterPublicKey"`
//}

func init() {
	// 需要在init中注册定义的model,自动生成数据库
	orm.RegisterModel(new(Block))
}
