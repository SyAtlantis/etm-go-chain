package transactions

import "workspace/etm-go-chain/models"

func init() {
	m := transaction{}
	models.RegisterModels("transactions", &m)
}

type Transaction interface {
	Create()
	GetBytes() []byte
	GetHash() [32]byte
	GetId() string
	GetSignature() string
}

type Transactions interface {
	GetTransactions()
	ReceiveTransactions()
	hasUnconfirmed()
	GetUnconfirmed()
	RemoveUnconfirmed()
	ProcessUnconfirmed()
	ApplyUnconfirmed()
	UndoUnconfirmed()
}

type transaction struct {
	models.Transaction
}

func (t transaction) NewModel() interface{} {
	return &transaction{}
}

func (t transaction) GetTransactions() {
	panic("implement me")
}

func (t transaction) ReceiveTransactions() {
	panic("implement me")
}

func (t transaction) hasUnconfirmed() {
	panic("implement me")
}

func (t transaction) GetUnconfirmed() {
	panic("implement me")
}

func (t transaction) RemoveUnconfirmed() {
	panic("implement me")
}

func (t transaction) ProcessUnconfirmed() {
	panic("implement me")
}

func (t transaction) ApplyUnconfirmed() {
	panic("implement me")
}

func (t transaction) UndoUnconfirmed() {
	panic("implement me")
}

func (t transaction) GetBytes() []byte {
	panic("implement me")
}
