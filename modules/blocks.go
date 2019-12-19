package modules

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
	"workspace/etm-go-chain/core"
	"workspace/etm-go-chain/models"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindBlock), event.Normal)
}

type Blocks interface {
	GetBlocks() []models.Block

	GenerateBlock() models.Block
	ProcessBlock(mb models.Block) error
	ApplyBlock(mb models.Block) error

	verifyBlock(mb models.Block) error
	verifyGenesisBlock(mb models.Block) error
	saveBlock(mb models.Block) error
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

func (b *block) GenerateBlock() models.Block{
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

func (b *block) ApplyBlock(mb models.Block) error {
	panic("implement me")
}

func (b *block) verifyBlock(mb models.Block) error {
	var err error
	if b.Height == 1 {
		err = b.verifyGenesisBlock()
	} else {
		// verify block transactions

		// verify block

	}
	return err
}

func (b *block) verifyGenesisBlock(mb models.Block) error {
	var err error

	return err
}

func (b *block) saveBlock(mb models.Block) error {
	// save block transactions
	trs := b.Transactions
	for _, tr := range trs {
		err := transactions.SaveTransaction(*tr)
		if err != nil {

		}
		logs.Info(tr)
	}

	// save block
	//o := orm.NewOrm()
	//_, _, err := o.ReadOrCreate(&b.Block, "Id")
	err := b.Block.DbSave(b.Block)
	if err != nil {
		logs.Error("Save block error! ==>", err)
	}

	return nil
}

func onBindBlock(e event.Event) error {
	logs.Info("onBind block", e.Data())

	genesisBlock := core.GetGenesisBlock()
	b := &block{genesisBlock}

	err := b.verifyGenesisBlock()
	if err != nil {
		logs.Error(" 【onBind】verify GenesisBlock error ==>", err)
	}

	err = b.saveBlock()
	if err != nil {
		logs.Error(" 【onBind】save GenesisBlock error ==>", err)
	}

	err, _ = event.Fire("load", event.M{})

	return err
}
