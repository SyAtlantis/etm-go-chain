package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type iTransaction interface {
	Create() error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature() (string, error)
	DbRead() (Transaction, error)
	DbSave(data Transaction) error
	Trans2Transaction(data interface{}) (Transaction, error)
	Trans2Object() (map[string]interface{}, error)
}

type Transaction struct {
	Id        string   `json:"id" orm:"pk"`
	Type      uint8    `json:"type"`
	BlockId   *Block   `json:"blockId" orm:"rel(fk);column(block_id)"`
	Fee       int64    `json:"fee"`
	Amount    int64    `json:"amount"`
	Timestamp int64    `json:"timestamp"`
	Sender    *Account `json:"sender" orm:"rel(fk);null"`
	Recipient *Account `json:"recipient" orm:"rel(fk);null"`
	Args      string   `json:"args"`
	Message   string   `json:"message"`
	Signature string   `json:"signature"`
}

func (t *Transaction) DbRead() (Transaction, error) {
	o := orm.NewOrm()
	err := o.Read(&t)
	return *t, err
}

func (t *Transaction) DbSave(data Transaction) error {
	o := orm.NewOrm()
	_, _, err := o.ReadOrCreate(&data, "Id")
	return err
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
	var err error
	o, ok := data.(map[string]interface{})

	t.Id, ok = o["id"].(string)
	t.Type, ok = o["type"].(uint8)
	id, ok := o["blockId"].(string)
	t.BlockId = &Block{Id: id,}
	t.Fee, ok = o["fee"].(int64)
	t.Amount, ok = o["amount"].(int64)
	t.Timestamp, ok = o["timestamp"].(int64)
	senderPublicKey, ok := o["senderPublicKey"].(string)
	t.Sender = &Account{PublicKey: senderPublicKey,}
	recipient, ok := o["recipientId"].(string)
	t.Recipient = &Account{Address: recipient,}
	t.Args, ok = o["args"].(string)
	t.Message, ok = o["message"].(string)
	t.Signature, ok = o["signature"].(string)

	if !ok {
		err = errors.New("Transform data to Transaction error")
	}
	return *t, err
}

func (t *Transaction) Trans2Object() (map[string]interface{}, error) {
	panic("implement me")
}

func init() {
	orm.RegisterModel(new(Transaction))
}
