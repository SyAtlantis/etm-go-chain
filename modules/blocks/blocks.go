package blocks

import (
	"github.com/astaxie/beego/logs"
	"github.com/gookit/event"
	"workspace/etm-go-chain/core"
	"workspace/etm-go-chain/models"
)

func init() {
	m := block{}
	models.RegisterModels("blocks", &m)

	event.On("bind", event.ListenerFunc(onBind), event.Normal)
}

type Blocks interface {
	GetBlocks() []block

	GenerateBlock()
	ProcessBlock() error
	ApplyBlock() error
}

type block struct {
	models.Block
}

func (b *block) NewModel() interface{} {
	return &block{}
}

func (b *block) GetBlocks() []block {
	panic("implement me")
}

func (b *block) GenerateBlock() {
	panic("implement me")
}

func (b *block) ProcessBlock() error {
	var err error
	err = varifyBlock(b)
	err = saveBlock(b)

	return err
}

func varifyBlock(b *block) error {
	var err error
	if b.Height == 1 {
		err = verifyGenesisBlock(b)
	} else {

	}
	return err
}
func verifyGenesisBlock(b *block) error {
	var err error

	return err
}
func saveBlock(b *block) error {
	var err error
	_, err = core.Insert(b)
	if err != nil {
		logs.Error("Save block error! ==>", err)
	}
	return err
}

func (b *block) ApplyBlock() error {
	panic("implement me")
}

func onBind(e event.Event) error {
	logs.Info("onBind block", e.Data())

	genesisBlock := core.GetGenesisBlock()
	b := &block{Block: genesisBlock}

	err := b.ProcessBlock()
	if err != nil {
		logs.Error("verifyGenesisBlock err", err)
	}

	return err
}
