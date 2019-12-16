package models

import (
	"github.com/astaxie/beego/orm"

	_ "workspace/etm-go-chain/models/accounts"
	_ "workspace/etm-go-chain/models/blocks"
	_ "workspace/etm-go-chain/models/peers"
	_ "workspace/etm-go-chain/models/system"
	_ "workspace/etm-go-chain/models/transactions"
)

type Peer struct {
	Id      int    `json:"id" orm:"pk;auto"`
	Ip      string `json:"ip"`
	Port    int64  `json:"port"`
	State   int    `json:"state"`
	Version string `json:"version"`
}

type Account struct {
	Address   string    `json:"address" orm:"pk"`
	PublicKey string    `json:"publicKey"`
	Balance   int64     `json:"balance"`
	Rewards   int64     `json:"rewards"`
	Bonus     int64     `json:"bonus"`
	Delegate  *Delegate `json:"delegate" orm:"rel(fk);null;column(delegate)"`
	Vote      *Vote     `json:"vote" orm:"rel(fk);null;column(vote)"`
}

type Delegate struct {
	Username       string  `json:"username" orm:"pk"`
	Rate           int     `json:"rate"`
	Votes          int64   `json:"votes"`
	Voters         []*Vote `json:"voters" orm:"reverse(many)"`
	ProducedBlocks int64   `json:"producedBlocks"`
	MissedBlocks   int64   `json:"missedBlocks"`
}

type Vote struct {
	TransactionId string    `json:"transactionId" orm:"pk;column(transactionId)"`
	Votes         int64     `json:"votes"`
	Delegate      *Delegate `json:"delegate" orm:"rel(fk);column(delegate)"`
	Locks         []*Lock   `json:"locks" orm:"rel(m2m)"`
}
type Lock struct {
	TransactionId string `json:"transactionId" orm:"pk;column(transactionId)"`
	LockAmount    int64  `json:"lockAmount"`
	OriginHeight  int64  `json:"originHeight"`
	CurrentHeight int64  `json:"currentHeight"`
	Votes         int64  `json:"votes"`
	State         int    `json:"state"`
}

type Block struct {
	Id                   string         `json:"id" orm:"pk"`
	Height               int64          `json:"height"`
	Timestamp            int64          `json:"timestamp"`
	TotalAmount          int64          `json:"totalAmount" orm:"column(totalAmount)"`
	TotalFee             int64          `json:"totalFee" orm:"column(totalFee)"`
	Reward               int64          `json:"reward"`
	PayloadHash          string         `json:"payloadHash" orm:"column(payloadHash)"`
	PayloadLength        int            `json:"payloadLength" orm:"column(payloadLength)"`
	PreviousBlock        string         `json:"previousBlock" orm:"column(previousBlock)"`
	Generator            *Delegate      `json:"generator" orm:"rel(fk);column(generator)"`
	BlockSignature       string         `json:"blockSignature" orm:"column(blockSignature)"`
	NumberOfTransactions int            `json:"numberOfTransactions" orm:"column(numberOfTransactions)"`
	Transactions         []*Transaction `json:"transactions" orm:"reverse(many)"`
}

type Transaction struct {
	Id        string `json:"id" orm:"pk"`
	Type      uint8  `json:"type"`
	BlockId   *Block `json:"blockId" orm:"rel(fk);column(blockId)"`
	Fee       int64  `json:"fee"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Args      string `json:"args"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

func init() {
	// 需要在init中注册定义的model,自动生成数据库
	orm.RegisterModel(new(Peer), new(Account), new(Delegate), new(Vote), new(Lock), new(Block), new(Transaction))
}
