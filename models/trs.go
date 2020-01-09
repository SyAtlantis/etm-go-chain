package models

import (
	"errors"
	"etm-go-chain/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"sort"
)

type Trs []*Transaction

type iTransactions interface {
	Sort()
	Apply() error
}

func (trs Trs) Len() int {
	return len(trs)
}

func (trs Trs) Swap(i, j int) {
	trs[i], trs[j] = trs[j], trs[i]
}

func (trs Trs) Less(i, j int) bool {
	if trs[i].Type != trs[j].Type {
		if trs[i].Type == 1 {
			return false
		}
		if trs[j].Type == 1 {
			return true
		}
		return trs[i].Type < trs[j].Type
	}
	if trs[i].Amount != trs[j].Amount {
		return trs[i].Amount < trs[j].Amount
	}
	return trs[i].Id < trs[j].Id
}

func (trs Trs) Sort() {
	sort.Sort(trs)
}

func (trs Trs) Apply() error {
	trs.Sort()

	o := orm.NewOrm()
	if _, err := o.Raw("PRAGMA synchronous = OFF").Exec(); err != nil {
		return err
	}

	qa := o.QueryTable("account")
	ia, _ := qa.PrepareInsert()
	qd := o.QueryTable("delegate")
	id, _ := qd.PrepareInsert()
	qv := o.QueryTable("vote")
	iv, _ := qv.PrepareInsert()
	ql := o.QueryTable("lock")
	il, _ := ql.PrepareInsert()
	defer ia.Close()
	defer id.Close()
	defer iv.Close()
	defer il.Close()

	for i, tr := range trs {
		logs.Debug(i, tr.Id)

		// 获取或者新建交易的account
		if tr.Sender == "" {
			return errors.New("no sender to load")
		}
		Address := utils.Address{}
		addr, err := Address.GenerateAddressByStr(tr.Sender)
		if err != nil {
			return err
		}

		s := &Account{
			Address:   addr,
			PublicKey: tr.Sender,
		}
		if tr.SAccount, err = readOrCreate(qa.Filter("Address", addr), ia, s); err != nil {
			return err
		}

		if tr.Recipient != "" {
			r := &Account{
				Address: tr.Recipient,
			}
			if tr.RAccount, err = readOrCreate(qa.Filter("Address", tr.Recipient), ia, r); err != nil {
				return err
			}
		}

		//更新交易数据到account
		if err := tr.Apply(); err != nil {
			return err
		}

		// 更新到数据库
		if err := updateOrInsert(qa.Filter("Address", addr), ia, tr.SAccount); err != nil {
			return err
		}
		if tr.Recipient != "" {
			if err := updateOrInsert(qa.Filter("Address", tr.Recipient), ia, tr.RAccount); err != nil {
				return err
			}
		}

		// 更新对应的 Delegate，Vote，Locks
		if tr.SAccount.Delegate != nil {
			if err := updateOrInsert(qd.Filter("TransactionId", tr.Id), id, tr.SAccount.Delegate); err != nil {
				return err
			}
		}
		if tr.SAccount.Vote != nil {
			if err := updateOrInsert(qv.Filter("TransactionId", tr.Id), iv, tr.SAccount.Vote); err != nil {
				return err
			}
		}
		if tr.SAccount.Locks != nil {
			if err := updateOrInsert(ql.Filter("TransactionId", tr.Id), il, tr.SAccount.Locks[0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func readOrCreate(qs orm.QuerySeter, i orm.Inserter, a *Account) (*Account, error) {
	if qs.Exist() {
		var account Account
		err := qs.One(&account)
		return &account, err
	} else {
		_, err := i.Insert(a)
		return a, err
	}
}

func updateOrInsert(qs orm.QuerySeter, i orm.Inserter, data interface{}) error {
	if qs.Exist() {
		if _, err := qs.Update(Struct2Map(data)); err != nil {
			return err
		}
	} else {
		_, err := i.Insert(data)
		return err
	}
	return nil
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj) // 获取 obj 的类型信息
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}

	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			//val := v.Field(i)
			//switch val.Kind() {
			//case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			//	if !val.IsNil() {
			//		data[t.Field(i).Name] = val.Interface()
			//	}
			//default:
			//	data[t.Field(i).Name] = val.Interface()
			//}
			name := t.Field(i).Name
			if !(name == "Delegate" || name == "Vote" || name == "Locks" || name == "Voters") {
				data[name] = v.Field(i).Interface()
			}
		}
	}

	return data
}
