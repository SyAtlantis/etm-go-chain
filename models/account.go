package models

import "github.com/astaxie/beego/orm"

type IAccount interface {
	Trans2Account(data interface{}) (Account, error)
	Trans2Object() (map[string]interface{}, error)
}
type Account struct {
	Address   string    `json:"address" orm:"pk"`
	PublicKey string    `json:"publicKey" orm:"column(publicKey)"`
	Balance   int64     `json:"balance"`
	Rewards   int64     `json:"rewards"`
	Bonus     int64     `json:"bonus"`
	Delegate  *Delegate `json:"delegate" orm:"rel(one);null"`
	Vote      *Vote     `json:"vote" orm:"rel(one);null"`
	Locks     []*Lock   `json:"locks" orm:"reverse(many);null"`
}

type Delegate struct {
	Username       string       `json:"username" orm:"pk"`
	Account        *Account     `json:"account" orm:"reverse(one)"`
	TransactionId  *Transaction `json:"transactionId" orm:"rel(one);column(transaction_id)"`
	Voters         []*Vote      `json:"voters" orm:"reverse(many)"`
	Rate           int          `json:"rate"`
	Votes          int64        `json:"votes"`
	ProducedBlocks int64        `json:"producedBlocks" orm:"column(producedBlocks)"`
	MissedBlocks   int64        `json:"missedBlocks" orm:"column(missedBlocks)"`
}

type Vote struct {
	Id            int
	Account       *Account     `json:"account" orm:"reverse(one)"`
	TransactionId *Transaction `json:"transactionId" orm:"rel(one);column(transaction_id)"`
	Delegate      *Delegate    `json:"delegate" orm:"rel(fk)"`
	Locks         []*Lock      `json:"locks" orm:"rel(m2m);null"`
	Votes         int64        `json:"votes"`
}

type Lock struct {
	Id            int
	Account       *Account     `json:"account" orm:"rel(fk)"`
	TransactionId *Transaction `json:"transactionId" orm:"rel(one);column(transaction_id)"`
	LockAmount    int64        `json:"lockAmount" orm:"column(lockAmount)"`
	OriginHeight  int64        `json:"originHeight" orm:"column(originHeight)"`
	CurrentHeight int64        `json:"currentHeight" orm:"column(currentHeight)"`
	Votes         int64        `json:"votes"`
	State         int          `json:"state"`
}

func init() {
	orm.RegisterModel(new(Account), new(Delegate), new(Vote), new(Lock))
}
