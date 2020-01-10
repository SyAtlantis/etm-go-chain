package modules

import (
	"encoding/hex"
	"errors"
	"etm-go-chain/utils"
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"

	"etm-go-chain/models"
)

func init() {
	event.On("ready", event.ListenerFunc(onReadyMyDelegates), event.Normal)
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

func onReadyMyDelegates(e event.Event) error {
	logs.Info("onReady myDelegates", e)
	return nil
}
