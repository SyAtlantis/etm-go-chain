package modules

import (
	"crypto/sha256"
	"etm-go-chain/core"
	"etm-go-chain/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gookit/event"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindBlock), event.Normal)
}

type Blocks interface {
	GetBlocks() []models.Block

	generateBlock() models.Block
	processBlock(mb models.Block) error

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

func (b *block) generateBlock() models.Block {
	panic("implement me")
}

func (b *block) processBlock(mb models.Block) error {
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
	if err != nil {
		return err
	}
	for count > offset {
		qb.Offset(offset).Limit(limit).OrderBy("height")
		var blockList []*models.Block
		if n, err := qb.All(&blockList); err == nil && n > 0 {
			for _, blockItem := range blockList {
				var trList models.Trs
				qt := o.QueryTable("transaction")
				if n, err := qt.Filter("block_id", blockItem.Id).RelatedSel("BlockId").All(&trList); err == nil && n > 0 {
					blockItem.Transactions = trList
				}

				if systems.GetLastHeight() != 0 {
					if err := blocks.verifyBlock(*blockItem); err != nil {
						return err
					}
				}
				if err := blocks.applyBlock(*blockItem); err != nil {
					return err
				}
				if err := systems.SetLastHeight(blockItem.Height); err != nil {
					return err
				}
			}
		}
		offset += limit
	}

	return nil
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

	//var err error
	//appliedTrIds := set.New(set.ThreadSafe)
	trs := mb.Transactions
	if err := trs.Apply(); err != nil {
		return err
	}
	//for i, tr := range trs {
	//	logs.Debug(i,tr.Id)
	//	if tr.SAccount, err = accounts.loadSender(tr.Sender); err != nil {
	//		return err
	//	}
	//
	//	if tr.Recipient != "" {
	//		if tr.RAccount, err = accounts.loadRecipient(tr.Recipient); err != nil {
	//			return err
	//		}
	//	}
	//
	//	//if err = transactions.applyUnconfirmed(*tr); err != nil {
	//	//	return err
	//	//}
	//
	//	if err = tr.Apply(); err != nil {
	//		return err
	//	}
	//
	//	//if err = transactions.removeUnconfirmed(*tr); err != nil {
	//	//	return err
	//	//}
	//
	//	appliedTrIds.Add(tr.Id)
	//}

	logs.Debug("apply block")

	return nil
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
