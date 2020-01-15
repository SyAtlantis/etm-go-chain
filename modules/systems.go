package modules

import (
	"etm-go-chain/models"
	"github.com/astaxie/beego/logs"
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

func (s *system) loadBlockChain() error {
	// clear tables accounts
	if err := accounts.RemoveTables(); err != nil {
		return err
	}

	// load blocks offset
	if err := blocks.loadBlocksOffset(0, 1000); err != nil {
		return err
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
