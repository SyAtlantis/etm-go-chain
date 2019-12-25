package modules

import (
	"crypto/sha256"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"

	"etm-go-chain/core"
	"etm-go-chain/models"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindBlock), event.Normal)
}

type Blocks interface {
	GetBlocks() []models.Block

	GenerateBlock() models.Block
	ProcessBlock(mb models.Block) error

	verifyBlock(mb models.Block) error
	verifyGenesisBlock(mb models.Block) error
	saveBlock(mb models.Block) error
	applyBlock(mb models.Block) error
}

type block struct {
	models.Block
}

func NewBlocks() Blocks {
	return &block{}
}

func (b *block) GetBlocks() []models.Block {
	panic("implement me")
}

func (b *block) GenerateBlock() models.Block {
	panic("implement me")
}

func (b *block) ProcessBlock(mb models.Block) error {
	//trs := b.Transactions
	//for _, tr := range trs {
	//	err := transactions.ProcessTransaction(*tr)
	//	if err != nil {
	//
	//	}
	//	logs.Info(tr)
	//}

	var err error
	err = b.verifyBlock(mb)
	err = b.saveBlock(mb)

	return err
}

func (b *block) verifyBlock(mb models.Block) error {
	var err error

	return err
}

func (b *block) verifyGenesisBlock(mb models.Block) error {
	var size int
	hash := sha256.New()
	trs := mb.Transactions
	for i := 0; i < len(trs); i++ {
		bs, err := trs[i].GetBytes()
		if err != nil {
			return err
		}
		size += len(bs)
		hash.Write(bs)
	}
	payloadHash := fmt.Sprintf("%x", hash.Sum([]byte{}))
	payloadLength := size
	id, err := mb.GetId()

	if payloadLength != mb.PayloadLength || payloadHash != mb.PayloadHash || id != mb.Id || err != nil {
		panic("Verify genesis block error!")
	}

	return nil
}

func (b *block) saveBlock(mb models.Block) error {
	err := mb.SetBlock()
	if err != nil {
		logs.Error("Save block error! ==>", err)
		return err
	}

	return nil
}

func (b *block) applyBlock(mb models.Block) error {
	panic("implement me")
}

func onBindBlock(e event.Event) error {
	logs.Info("onBind block", e.Data())

	genesisBlock := core.GetGenesisBlock()

	if err := blocks.verifyGenesisBlock(genesisBlock); err != nil {
		logs.Error(" 【onBind】verify GenesisBlock error ==>", err)
		return err
	}

	if err := blocks.saveBlock(genesisBlock); err != nil {
		logs.Error(" 【onBind】save GenesisBlock error ==>", err)
		return err
	}

	if err, _ := event.Fire("load", event.M{}); err != nil {
		return err
	}

	return nil
}
