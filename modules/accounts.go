package modules

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"etm-go-chain/core"
	"etm-go-chain/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"

	"etm-go-chain/models"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindMyDelegates), event.Normal)
}

type Accounts interface {
	GetAccounts()
	RemoveTables() error
	loadSender(string) (models.Account, error)
	loadRecipient(string) (models.Account, error)
}

type account struct {
	models.Account
}

func NewAccounts() Accounts {
	return &account{}
}

func (a account) GetAccounts() {
	panic("implement me")
}

func (a account) RemoveTables() error {
	err := a.Account.ClearAccount()
	return err
}

func (a account) loadSender(sender string) (models.Account, error) {
	acc := models.Account{}
	if sender == "" {
		return acc, errors.New("no sender to load")
	}

	Address := utils.Address{}
	pub, err := hex.DecodeString(sender)
	if err != nil {
		return acc, err
	}
	acc.Address = Address.GenerateAddress(pub)

	if acc2, err := acc.GetAccount(); err == nil {
		acc2.PublicKey = sender
		return acc2, nil
	}

	acc.PublicKey = sender
	return acc, acc.SetAccount()
}

func (a account) loadRecipient(recipient string) (models.Account, error) {
	acc := models.Account{}
	if recipient == "" {
		return acc, errors.New("no recipient to load")
	}

	acc.Address = recipient
	if a, err := acc.GetAccount(); err == nil {
		return a, nil
	}
	return acc, acc.SetAccount()
}

func (a account) getMyDelegates() error {

	return nil
}

func onBindMyDelegates(e event.Event) error {
	logs.Notice("【onBind】 myDelegates", e)

	config := core.GetConfig()
	var myDelegates []string
	
	logs.Warn("TODO 需要判断secret是否合法!")
	for _, s := range config.Secrets {
		// TODO 需要判断secret是否合法

		ed := utils.Ed{}
		hash := sha256.Sum256([]byte(s))
		keypair := ed.MakeKeypair(hash[:])
		pub := fmt.Sprintf("%x", keypair.PublicKey)
		myDelegates = append(myDelegates, pub)
	}
	if err := systems.SetMyDelegates(myDelegates); err != nil {
		return err
	}

	return nil
}

