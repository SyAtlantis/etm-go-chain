package models

import "github.com/astaxie/beego/orm"

type ITransaction interface {
	Create() error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature() (string, error)
	Trans2Transaction(data interface{}) (Transaction,error)
	Trans2Object() (map[string]interface{}, error)
}

type Transaction struct {
	Id        string   `json:"id" orm:"pk"`
	Type      uint8    `json:"type"`
	BlockId   *Block   `json:"blockId" orm:"rel(fk);column(block_id)"`
	Fee       int64    `json:"fee"`
	Amount    int64    `json:"amount"`
	Timestamp int64    `json:"timestamp"`
	Sender    *Account `json:"sender" orm:"rel(fk)"`
	Recipient *Account `json:"recipient" orm:"rel(fk)"`
	Args      string   `json:"args"`
	Message   string   `json:"message"`
	Signature string   `json:"signature"`
}

func (t *Transaction) Create() error {
	panic("implement me")
}

func (t *Transaction) GetBytes() ([]byte, error) {
	panic("implement me")
}

func (t *Transaction) GetHash() ([32]byte, error) {
	panic("implement me")
}

func (t *Transaction) GetId() (string, error) {
	panic("implement me")
}

func (t *Transaction) GetSignature() (string, error) {
	panic("implement me")
}

func (t *Transaction) Trans2Transaction(data interface{}) (Transaction, error) {
	panic("implement me")
}

func (t *Transaction) Trans2Object() (map[string]interface{}, error) {
	panic("implement me")
}

func init() {
	orm.RegisterModel(new(Transaction))
}
