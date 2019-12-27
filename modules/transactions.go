package modules

import (
	"github.com/astaxie/beego/logs"

	"etm-go-chain/models"
)

func init() {
}

type Transactions interface {
	GetTransactions() ([]models.Transaction, error)
	ProcessTransaction(tr models.Transaction) error
	SaveTransaction(tr models.Transaction) error
	SaveTransactions(trs []models.Transaction) error
	ReceiveTransactions()

	hasUnconfirmed() bool
	GetUnconfirmed() ([]models.Transaction, error)
	RemoveUnconfirmed(tr models.Transaction) error
	ProcessUnconfirmed(tr models.Transaction) error
	ApplyUnconfirmed(tr models.Transaction) error
	UndoUnconfirmed(tr models.Transaction) error
}

type transaction struct {
	models.Transaction
}

func (t transaction) ProcessTransaction(tr models.Transaction) error {
	// verify tr

	// save tr

	return nil
}

func NewTransactions() Transactions {
	return &transaction{}
}

func (t transaction) GetTransactions() ([]models.Transaction, error) {
	panic("implement me")
}

func (t transaction) ReceiveTransactions() {
	panic("implement me")
}

func (t transaction) hasUnconfirmed() bool {
	panic("implement me")
}

func (t transaction) GetUnconfirmed() ([]models.Transaction, error) {
	panic("implement me")
}

func (t transaction) RemoveUnconfirmed(tr models.Transaction) error {
	logs.Debug("Remove Unconfirmed transaction")
	return nil
}

func (t transaction) ProcessUnconfirmed(tr models.Transaction) error {
	panic("implement me")
}

func (t transaction) ApplyUnconfirmed(tr models.Transaction) error {
	logs.Debug("Apply Unconfirmed transaction")
	return nil
}

func (t transaction) UndoUnconfirmed(tr models.Transaction) error {
	panic("implement me")
}

func (t transaction) SaveTransaction(tr models.Transaction) error {
	err := tr.SetTransaction()
	if err != nil {
		logs.Error("Save transaction error! ==>", err)
	}

	return err
}

func (t transaction) SaveTransactions(trs []models.Transaction) error {
	//err := t.Transaction.SetTransactions(trs)
	//if err != nil {
	//	logs.Error("Save transaction multi error! ==>", err)
	//}

	return nil
}
