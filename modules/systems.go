package modules

import (
	"etm-go-chain/models"
	"etm-go-chain/utils"
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
	GetLastHeight() int64
	SetLastHeight(int64) error
	GetLastBlock() *models.Block
	SetLastBlock(*models.Block) error
	GetMyDelegates() []string
	SetMyDelegates([]string) error
	GetDelegateList() []string
	SetDelegateList([]string) error

	Loops() error
	LoadBlockChain() error
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

func (s *system) GetLastHeight() int64 {
	return s.LastHeight
}

func (s *system) SetLastHeight(h int64) error {
	s.LastHeight = h
	return nil
}

func (s *system) GetLastBlock() *models.Block {
	return s.LastBlock
}

func (s *system) SetLastBlock(b *models.Block) error {
	s.LastBlock = b
	return nil
}

func (s *system) GetMyDelegates() []string {
	return s.MyDelegates
}

func (s *system) SetMyDelegates(delegates []string) error {
	s.MyDelegates = delegates
	return nil
}

func (s *system) GetDelegateList() []string {
	return s.DelegateList
}

func (s *system) SetDelegateList(delegates []string) error {
	s.DelegateList = delegates
	return nil
}

func (s *system) Loops() error {
	logs.Debug("in loops")

	slot := utils.NewSlots()
	currentSlot := slot.GetSlotNumber()
	lastBlock := systems.GetLastBlock()
	if currentSlot == lastBlock.Timestamp {
		return nil
	}

	timestamp, keypair, err := getBlockSlotData(currentSlot, lastBlock.Height+1)
	if err != nil {
		return err
	}
	if slot.GetSlotNumber(timestamp) == slot.GetSlotNumber() && systems.GetLastBlock().Timestamp < timestamp {
		if err := blocks.generateBlock(keypair, timestamp); err != nil {
			return err
		}
	}

	return nil
}

func getBlockSlotData(slot int64, height int64) (time int64, keypair utils.Keypair, err error) {
	return time, keypair, err
}

func (s *system) LoadBlockChain() error {
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

	if err := systems.LoadBlockChain(); err != nil {
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

	myDelegates := systems.GetMyDelegates()
	if len(myDelegates) > 0 {
		go func() {
			for {
				if err := systems.Loops(); err != nil {
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
