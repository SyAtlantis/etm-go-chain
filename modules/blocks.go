package modules

import (
	"crypto/sha256"
	"etm-go-chain/core"
	"etm-go-chain/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/fatih/set"
	"github.com/gookit/event"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindBlock), event.Normal)
}

type Blocks interface {
	GetBlocks() []models.Block

	GenerateBlock() models.Block
	ProcessBlock(mb models.Block) error

	loadBlocksOffset(offset int64, limit int64) error
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

func (b *block) loadBlocksOffset(offset int64, limit int64) error {
	// get blocks offset limit
	// get block trs
	// sort trs
	// if no last block applyBlock
	// else verifyBlock applyBlock
	// set last block

	o := orm.NewOrm()
	qb := o.QueryTable("block")

	count, err := qb.Count()
	for count > offset {
		qb.Offset(offset).Limit(limit).OrderBy("height")
		var n int64
		var blockList []*models.Block
		n, err = qb.All(&blockList)
		if err == nil && n > 0 {
			for _, blockItem := range blockList {
				var trList models.Trs
				qt := o.QueryTable("transaction")
				n, err = qt.Filter("block_id", blockItem.Height).All(&trList)
				if err == nil && n > 0 {
					//trList.Sort()
					blockItem.Transactions = trList
				}

				if systems.GetLastHeight() != 0 {
					err = blocks.verifyBlock(*blockItem)
				}
				err = blocks.applyBlock(*blockItem)

				err = systems.SetLastHeight(blockItem.Height)
			}
		}
		offset += limit
	}

	return err
}

func (b *block) verifyBlock(mb models.Block) error {
	var err error
	logs.Debug("verify block")
	return err
}

func (b *block) verifyGenesisBlock(mb models.Block) error {
	var payloadLength int
	hash := sha256.New()
	trs := mb.Transactions
	for i := 0; i < len(trs); i++ {
		bs, err := trs[i].GetBytes()
		if err != nil {
			return err
		}
		payloadLength += len(bs)
		hash.Write(bs)
	}
	payloadHash := fmt.Sprintf("%x", hash.Sum([]byte{}))
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
	// get unconfirmedList
	// undo unconfirmedList
	// do apply

	var err error
	appliedTrIds := set.New(set.ThreadSafe)
	trs := mb.Transactions
	trs.Sort()
	for _, tr := range trs {
		if mb.Height == 1 {
			// create account
		} else {
			// update account
		}

		if err = transactions.ApplyUnconfirmed(*tr); err != nil {
			return err
		}

		if err = tr.Apply(); err != nil {
			return err
		}

		if err = transactions.RemoveUnconfirmed(*tr); err != nil {
			return err
		}

		appliedTrIds.Add(tr.Id)
	}

	logs.Debug("apply block")
	return err
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
