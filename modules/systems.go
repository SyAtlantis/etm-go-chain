package modules

import (
	"encoding/json"
	"etm-go-chain/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gookit/event"
	"time"
)

func init() {
	event.On("load", event.ListenerFunc(onLoadBlockChain), event.Normal)
	event.On("ready", event.ListenerFunc(onReadyLoops), event.Normal)
}

type Systems interface {
	GetVersion() string

	loops() error
	calcRound(height int64) int64
	tick(block models.Block) error
	backwardTick(block models.Block) error
	loadBlockChain() error
}

type system struct {
	models.System
}

func NewSystems() Systems {
	return &system{}
}

func (s *system) GetVersion() string {
	panic("implement me")
}

func (s *system) loops() error {
	currentSlot := slots.GetSlotNumber()
	lastBlock := blocks.GetLastBlock()
	if currentSlot == lastBlock.Timestamp {
		return nil
	}

	blockData, err := blocks.getBlockSlotData(lastBlock.Height+1, currentSlot)
	if err != nil {
		return err
	}
	if slots.GetSlotNumber(blockData.Timestamp) == slots.GetSlotNumber() && blocks.GetLastBlock().Timestamp < blockData.Timestamp {
		if err := blocks.generateBlock(blockData); err != nil {
			return err
		}
	}

	return nil
}

func (s *system) calcRound(height int64) int64 {
	var next int64
	if height%int64(slots.RoundBlocks) > 0 {
		next = 1
	} else {
		next = 0
	}
	return height/int64(slots.RoundBlocks) + next
}

func (s *system) tick(block models.Block) error {
	// TODO tick
	// 处理出块奖励 、分红
	// 处理换轮
	logs.Debug("TODO tick")

	jsonBytes, err := json.MarshalIndent(block, "", "    ")
	if err != nil {
		return err
	}
	logs.Debug("tick block completed:", string(jsonBytes))

	return nil
}

func (s *system) backwardTick(block models.Block) error {
	// TODO backward tick
	logs.Debug("TODO backward tick")
	return nil
}

func (s *system) loadBlockChain() error {
	o := orm.NewOrm()
	qb := o.QueryTable("block")
	count, err := qb.Count()
	if err != nil {
		return err
	}

	var offset int64 = 0
	var limit int64 = 1000
	if count <= 1 {
		// clear tables accounts
		if err := accounts.RemoveTables(); err != nil {
			return err
		}
	} else {
		offset = count - 1
		limit = 1
	}

	// load blocks offset
	for count > offset {
		if err := blocks.loadBlocksOffset(offset, limit); err != nil {
			return err
		}
		offset += limit
	}

	return nil
}

func onLoadBlockChain(e event.Event) error {
	logs.Notice("【onLoad】 blockChain", e.Data())

	if err := systems.loadBlockChain(); err != nil {
		return err
	}

	go func() {
		logs.Notice("Event fire 【onReady】")
		if err, _ := event.Fire("ready", event.M{}); err != nil {
			panic(err)
		}
	}()

	return nil
}

func onReadyLoops(e event.Event) error {
	logs.Notice("【onReady】 loops", e.Data())

	myDelegates := accounts.GetMyKeypairs()
	if len(myDelegates) > 0 {
		go func() {
			for {
				if err := systems.loops(); err != nil {
					panic(err)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()
	} else {
		logs.Debug("no delegates has been found")
	}

	return nil
}
