package models

import (
	"github.com/astaxie/beego/orm"
	"reflect"
)

func init() {
	orm.RegisterModel(new(Account))
	orm.RegisterModelWithPrefix("account#", new(Delegate), new(Vote), new(Lock))
}

type iAccount interface {
	IsEmpty() bool
	Merge() error
	GetAccount() (Account, error)
	SetAccount() error
	ClearAccount() error
	GetAccounts() ([]Account, error)
	SetAccounts(as []Account) error
}

type Account struct {
	Key       int64     `orm:"pk;auto"`
	Address   string    `json:"address"`
	PublicKey string    `json:"publicKey" orm:"column(publicKey)"`
	Balance   int64     `json:"balance"`
	Rewards   int64     `json:"rewards"`
	Bonus     int64     `json:"bonus"`
	Delegate  *Delegate `json:"delegate" orm:"rel(one);null"`
	Vote      *Vote     `json:"vote" orm:"rel(one);null"`
	Locks     []*Lock   `json:"locks" orm:"reverse(many);null"`
}

type Delegate struct {
	Key            int64        `orm:"pk;auto"`
	Username       string       `json:"username"`
	Account        *Account     `json:"account" orm:"reverse(one);cascade"`
	TransactionId  *Transaction `json:"transactionId" orm:"rel(one);column(transaction_id)"`
	Voters         []*Vote      `json:"voters" orm:"reverse(many)"`
	Rate           int          `json:"rate"`
	Votes          int64        `json:"votes"`
	ProducedBlocks int64        `json:"producedBlocks" orm:"column(producedBlocks)"`
	MissedBlocks   int64        `json:"missedBlocks" orm:"column(missedBlocks)"`
}

type Vote struct {
	Key           int64        `orm:"pk;auto"`
	Account       *Account     `json:"account" orm:"reverse(one);cascade"`
	TransactionId *Transaction `json:"transactionId" orm:"rel(one);column(transaction_id)"`
	Delegate      *Delegate    `json:"delegate" orm:"rel(fk)"`
	Locks         []*Lock      `json:"locks" orm:"rel(m2m);rel_table(account#vote_locks);null"`
	Votes         int64        `json:"votes"`
}

type Lock struct {
	Key           int64        `orm:"pk;auto"`
	Account       *Account     `json:"account" orm:"rel(fk);cascade"`
	TransactionId *Transaction `json:"transactionId" orm:"rel(one);column(transaction_id)"`
	LockAmount    int64        `json:"lockAmount" orm:"column(lockAmount)"`
	OriginHeight  int64        `json:"originHeight" orm:"column(originHeight)"`
	CurrentHeight int64        `json:"currentHeight" orm:"column(currentHeight)"`
	Votes         int64        `json:"votes"`
	State         int          `json:"state"`
}

func (a *Account) IsEmpty() bool {
	return a == nil || reflect.DeepEqual(a, Account{})
}

func (a *Account) Merge() error {
	o := orm.NewOrm()
	_, err := o.Update(a)
	//logs.Debug("Merge account")
	return err
}

func (a *Account) GetAccount() (Account, error) {
	o := orm.NewOrm()
	err := o.Read(&a)
	return *a, err
}

func (a *Account) SetAccount() error {
	o := orm.NewOrm()

	if a.Delegate != nil {
		if _, _, err := o.ReadOrCreate(a.Delegate, "TransactionId"); err != nil {
			return err
		}
	}
	if a.Vote != nil {
		if _, _, err := o.ReadOrCreate(a.Vote, "TransactionId"); err != nil {
			return err
		}
	}
	if a.Locks != nil {
		if _, _, err := o.ReadOrCreate(a.Locks[0], "TransactionId"); err != nil {
			return err
		}
	}

	if _, _, err := o.ReadOrCreate(a, "Address"); err != nil {
		return err
	}
	return nil
}

func (a *Account) ClearAccount() error {
	//o := orm.NewOrm()
	//qs := o.QueryTable("account")
	//if _, err := qs.Delete(); err != nil {
	//	return err
	//}
	return nil
}

func (a *Account) GetAccounts() ([]Account, error) {
	panic("implement me")
}

func (a *Account) SetAccounts(as []Account) error {
	o := orm.NewOrm()
	_, _, err := o.ReadOrCreate(a, "Address")
	return err
}
