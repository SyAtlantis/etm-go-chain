package modules

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"etm-go-chain/core"
	"etm-go-chain/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gookit/event"
)

func init() {
	event.On("bind", event.ListenerFunc(onBindGenesisBlock), event.Normal)
}

type Blocks interface {
	GetBlocks(map[string]interface{}, int64, int64, string) ([]models.Block, error)
	GetLastBlock() models.Block

	getBlockSlotData(int64, int64) (models.BlockData, error)
	generateBlock(models.BlockData) error
	processBlock(mb models.Block) error

	loadBlocksOffset(offset int64, limit int64) error
	verifyGenesisBlock(mb models.Block) error
	verifyBlock(mb models.Block) error
	saveBlock(mb models.Block) error
	applyBlock(mb models.Block) error
}

type block struct {
	models.Block
	LastBlock models.Block
}

func NewBlocks() Blocks {
	return &block{}
}

func (b *block) GetBlocks(filter map[string]interface{}, limit int64, offset int64, order string) ([]models.Block, error) {
	o := orm.NewOrm()
	qb := o.QueryTable("block")
	for k, v := range filter {
		qb.Filter(k, v)
	}
	if limit > 0 {
		qb.Limit(limit)
	}
	if offset > 0 {
		qb.Offset(offset)
	}
	if order != "" {
		qb.OrderBy(order)
	}
	var bs []models.Block
	if _, err := qb.All(&bs); err != nil {
		return bs, err
	}
	return bs, nil
}

func (b *block) GetLastBlock() models.Block {
	return b.LastBlock
}

func (b *block) getBlockSlotData(height int64, slot int64) (models.BlockData, error) {
	var bd models.BlockData
	delegate, err := accounts.getConsensusDelegate(height, slot)
	if err != nil {
		return bd, err
	}

	myKeypairs := accounts.GetMyKeypairs()
	selectDelegate := myKeypairs[delegate.Account.PublicKey]
	if !selectDelegate.IsEmpty() {
		bd.Keypair = selectDelegate
	}
	bd.Timestamp = slots.GetSlotTime(slot)

	return bd, nil
}

func (b *block) generateBlock(bd models.BlockData) error {
	var trList models.Trs
	bd.Transactions = trList
	bd.PreviousBlock = blocks.GetLastBlock()

	var err error
	var newBlock models.Block
	if err = newBlock.Create(bd); err != nil {
		return err
	}
	if newBlock.Id, err = newBlock.GetId(); err != nil {
		return err
	}
	newBlock.Height = bd.PreviousBlock.Height + 1

	var sign models.Sign
	if sign, err = accounts.getMySigns(newBlock); err != nil {
		return err
	}
	if sign.HasEnoughSigns() {
		if err := blocks.processBlock(newBlock); err != nil {
			return err
		}
	} else {
		// TODO createPropose
	}

	return nil
}

func (b *block) processBlock(mb models.Block) error {
	if err := b.verifyBlock(mb); err != nil {
		return err
	}
	if err := b.saveBlock(mb); err != nil {
		return err
	}
	if err := b.applyBlock(mb); err != nil {
		return err
	}

	return nil
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

				if !b.LastBlock.IsEmpty() && b.LastBlock.Height != 0 {
					if err := blocks.verifyBlock(*blockItem); err != nil {
						return err
					}
				}
				if err := blocks.applyBlock(*blockItem); err != nil {
					return err
				}
			}
		}
		offset += limit
	}

	return nil
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
	payloadHash := hex.EncodeToString(hash.Sum([]byte{}))
	id, err := mb.GetId()

	if payloadLength != mb.PayloadLength || payloadHash != mb.PayloadHash || id != mb.Id || err != nil {
		panic("Verify genesis block error!")
	}

	return nil
}

func (b *block) verifyBlock(mb models.Block) error {
	var err error
	logs.Debug("verify block")
	return err
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
	trs := mb.Transactions
	if err := trs.Apply(); err != nil {
		return err
	}
	b.LastBlock = mb

	jsonBytes, err := json.MarshalIndent(mb, "", "    ")
	if err != nil {
		return err
	}
	logs.Debug("applied block:", string(jsonBytes))

	return nil
}

func onBindGenesisBlock(e event.Event) error {
	logs.Notice("【onBind】 genesisBlock", e.Data())

	genesisBlock := core.GetGenesisBlock()

	if err := blocks.verifyGenesisBlock(genesisBlock); err != nil {
		logs.Error(" 【onBind】verify GenesisBlock error ==>", err)
		return err
	}

	if err := blocks.saveBlock(genesisBlock); err != nil {
		logs.Error(" 【onBind】save GenesisBlock error ==>", err)
		return err
	}

	go func() {
		logs.Notice("Event fire 【onLoad】")
		if err, _ := event.Fire("load", event.M{}); err != nil {
			panic(err)
		}
	}()

	return nil
}
