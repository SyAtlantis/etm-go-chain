package models

import (
	"github.com/astaxie/beego/orm"
	"reflect"
)

func init() {
	orm.RegisterModel(new(Account), new(Delegate), new(Vote), new(Lock))
	//orm.RegisterModelWithPrefix("account#", new(Delegate), new(Vote), new(Lock))
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
	Delegate  *Delegate `json:"delegate" orm:"reverse(one);on_delete(cascade);null"`
	Vote      *Vote     `json:"vote" orm:"reverse(one);on_delete(cascade);null"`
	Locks     []*Lock   `json:"locks" orm:"reverse(many);on_delete(cascade);null"`
}

type Delegate struct {
	Key            int64    `orm:"pk;auto"`
	Account        *Account `json:"account" orm:"rel(one);on_delete(set_null);null"`
	TransactionId  string   `json:"transactionId" orm:"column(transactionId)"`
	Username       string   `json:"username"`
	Rate           int      `json:"rate"`
	Votes          int64    `json:"votes"`
	ProducedBlocks int64    `json:"producedBlocks" orm:"column(producedBlocks)"`
	MissedBlocks   int64    `json:"missedBlocks" orm:"column(missedBlocks)"`
	Voters         []*Vote  `json:"voters" orm:"-"`
}

type Vote struct {
	Key           int64    `orm:"pk;auto"`
	Account       *Account `json:"account" orm:"rel(one);on_delete(set_null);null"`
	TransactionId string   `json:"transactionId" orm:"column(transactionId)"`
	Delegate      string   `json:"delegate"`
	Votes         int64    `json:"votes"`
	Locks         []*Lock  `json:"locks" orm:"-"`
}

type Lock struct {
	Key           int64    `orm:"pk;auto"`
	Account       *Account `json:"account" orm:"rel(fk);on_delete(set_null);null"`
	TransactionId string   `json:"transactionId" orm:"column(transactionId)"`
	LockAmount    int64    `json:"lockAmount" orm:"column(lockAmount)"`
	OriginHeight  int64    `json:"originHeight" orm:"column(originHeight)"`
	CurrentHeight int64    `json:"currentHeight" orm:"column(currentHeight)"`
	Votes         int64    `json:"votes"`
	State         int      `json:"state"`
}

func (a *Account) IsEmpty() bool {
	return a == nil || reflect.DeepEqual(a, Account{})
}

func (a *Account) Merge() error {
	o := orm.NewOrm()
	if _, err := o.Update(a); err != nil {
		return err
	}

	return nil
}

func (a *Account) GetAccount() (Account, error) {
	o := orm.NewOrm()
	err := o.Read(a, "Address")
	return *a, err
}

func (a *Account) SetAccount() error {
	o := orm.NewOrm()
	//o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	//if a.Delegate != nil {
	//	if _, _, err := o.ReadOrCreate(a.Delegate, "TransactionId"); err != nil {
	//		return err
	//	}
	//}
	//if a.Vote != nil {
	//	if _, _, err := o.ReadOrCreate(a.Vote, "TransactionId"); err != nil {
	//		return err
	//	}
	//}
	//if a.Locks != nil {
	//	if _, _, err := o.ReadOrCreate(a.Locks[0], "TransactionId"); err != nil {
	//		return err
	//	}
	//}

	if _, _, err := o.ReadOrCreate(a, "Address"); err != nil {
		return err
	}

	return nil
}

func (a *Account) ClearAccount() error {
	//var w io.Writer
	//o := orm.NewOrm()
	//o.Raw("UPDATE sqlite_sequence set seq = 0 where name = ?","account")
	//o.Raw("DELETE from sqlite_sequence where name = ?","account")
	//o.Raw("DELETE from sqlite_sequence")
	//orm.DebugLog = orm.NewLog(w)

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
