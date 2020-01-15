package modules

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"etm-go-chain/core"
	"etm-go-chain/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gookit/event"

	"etm-go-chain/models"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindMyDelegates), event.Normal)
}

type Accounts interface {
	GetAccounts()
	GetMyKeypairs() map[string]utils.Keypair
	RemoveTables() error
	
	bindMyKeypairs([]string) error
	getDelegateList(height int64) ([]models.Delegate, error)
	getActiveDelegateKeypairs(height int64) ([]utils.Keypair, error)
	getConsensusDelegate(height int64, slot int64) (models.Delegate, error)
	getDelegateIndex(id string, slot int64, length int) int
	getMySigns(block models.Block) (models.Sign, error)
}

type account struct {
	models.Account
	MyKeypairs map[string]utils.Keypair
}

func NewAccounts() Accounts {
	return &account{}
}

func (a *account) GetAccounts() {
	panic("implement me")
}

func (a *account) GetMyKeypairs() map[string]utils.Keypair {
	return a.MyKeypairs
}

func (a account) RemoveTables() error {
	err := a.Account.ClearAccount()
	return err
}

func (a *account) bindMyKeypairs(secrets []string) error {
	myKeypairs := make(map[string]utils.Keypair)

	logs.Warn("TODO 需要判断secret是否合法!")
	for _, s := range secrets {
		// TODO 需要判断secret是否合法

		ed := utils.Ed{}
		hash := sha256.Sum256([]byte(s))
		keypair := ed.MakeKeypair(hash[:])
		pub := hex.EncodeToString(keypair.PublicKey)
		myKeypairs[pub] = keypair
	}
	a.MyKeypairs = myKeypairs
	return nil
}

func (a *account) getDelegateList(height int64) ([]models.Delegate, error) {
	var delegateList []models.Delegate
	o := orm.NewOrm()
	qd := o.QueryTable("delegate").OrderBy("Rate").Limit(slots.Delegates).RelatedSel("account")
	_, err := qd.All(&delegateList)
	return delegateList, err
}

func (a *account) getActiveDelegateKeypairs(height int64) (keypairs []utils.Keypair, err error) {
	var delegateList []models.Delegate
	if delegateList, err = accounts.getDelegateList(height); err != nil {
		return keypairs, err
	}
	myKeypairs := accounts.GetMyKeypairs()
	for k1, v1 := range myKeypairs {
		for _, v2 := range delegateList {
			if k1 == v2.Account.PublicKey {
				keypairs = append(keypairs, v1)
			}
		}
	}
	return keypairs, err
}

func (a *account) getConsensusDelegate(height int64, slot int64) (d models.Delegate, err error) {
	var delegateList []models.Delegate
	if delegateList, err = accounts.getDelegateList(height); err != nil {
		return d, err
	}
	prevHeight := height - 1
	if height > 2 {
		prevHeight = height - 2
	}
	filter := make(map[string]interface{})
	filter["height"] = prevHeight
	bs, err := blocks.GetBlocks(filter, 0, 0, "")
	if err != nil {
		return d, err
	}
	block := bs[0]
	index := accounts.getDelegateIndex(block.Id, slot, len(delegateList))
	d = delegateList[index]

	return d, err
}

func (a *account) getDelegateIndex(id string, slot int64, length int) int {
	hash := sha256.Sum256([]byte(id))
	h := hex.EncodeToString(hash[:])
	index := utils.Chaos(h, slot, length)
	return index
}

func (a *account) getMySigns(block models.Block) (models.Sign, error) {
	var sign models.Sign
	activeKeypairs, err := accounts.getActiveDelegateKeypairs(block.Height)
	if err != nil {
		return sign, err
	}
	if len(activeKeypairs) <= 0 {
		return sign, errors.New("active keypairs should not be empty")
	}
	if err := sign.Create(activeKeypairs, block); err != nil {
		return sign, err
	}

	return sign, nil
}

func onBindMyDelegates(e event.Event) error {
	logs.Notice("【onBind】 myDelegates", e)

	config := core.GetConfig()
	if err := accounts.bindMyKeypairs(config.Secrets); err != nil {
		return err
	}

	return nil
}
