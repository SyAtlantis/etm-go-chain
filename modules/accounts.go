package modules

import (
	"errors"
	"etm-go-chain/utils"
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"

	"etm-go-chain/models"
)

func init() {
	event.On("load", event.ListenerFunc(onLoadAccounts), event.Normal)
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
	acc.PublicKey = sender
	acc.Address = Address.GenerateAddresss([]byte(sender))
	return acc, acc.SetAccount()
}

func (a account) loadRecipient(recipient string) (models.Account, error) {
	acc := models.Account{}
	if recipient == "" {
		return acc, errors.New("no recipient to load")
	}

	acc.Address = recipient
	return acc, acc.SetAccount()
}

func onLoadAccounts(e event.Event) error {
	logs.Info("onload account", e)
	return nil
}
