package models

import (
	"errors"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Block))
}

type iBlock interface {
	Create() error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature() (string, error)
	GetBlock() (Block, error)
	SetBlock() error
	GetBlocks() ([]Block, error)
	SetBlocks(bs []Block) error
	SortTransactions() ([]Transaction, error)
	Trans2Block(data interface{}) (Block, error)
	Trans2Object() (map[string]interface{}, error)
}

type Block struct {
	Id                   string         `json:"id"`
	Height               int64          `json:"height" orm:"pk"`
	Timestamp            int64          `json:"timestamp"`
	TotalAmount          int64          `json:"totalAmount" orm:"column(totalAmount)"`
	TotalFee             int64          `json:"totalFee" orm:"column(totalFee)"`
	Reward               int64          `json:"reward"`
	PayloadHash          string         `json:"payloadHash" orm:"column(payloadHash)"`
	PayloadLength        int            `json:"payloadLength" orm:"column(payloadLength)"`
	PreviousBlock        string         `json:"previousBlock" orm:"column(previousBlock)"`
	Generator            *Delegate      `json:"generator" orm:"rel(fk);column(generator_id)"`
	BlockSignature       string         `json:"blockSignature" orm:"column(blockSignature)"`
	NumberOfTransactions int            `json:"numberOfTransactions" orm:"column(numberOfTransactions)"`
	Transactions         []*Transaction `json:"transactions" orm:"reverse(many)"`
}

func (b *Block) Create() error {
	panic("implement me")
}

func (b *Block) GetBytes() ([]byte, error) {
	panic("implement me")
}

func (b *Block) GetHash() ([32]byte, error) {
	panic("implement me")
}

func (b *Block) GetId() (string, error) {
	panic("implement me")
}

func (b *Block) GetSignature() (string, error) {
	panic("implement me")
}

func (b *Block) GetBlock() (Block, error) {
	o := orm.NewOrm()
	err := o.Read(&b)
	return *b, err
}

func (b *Block) SetBlock() error {
	var err error
	o := orm.NewOrm()

	//保存区块错误时，需要事物回滚
	err = o.Begin()

	var created bool
	created, _, err = o.ReadOrCreate(b, "Id")
	if err != nil {
		rollErr := o.Rollback()
		if rollErr != nil {
			panic(rollErr)
		}
		return err
	}

	if created {
		// 保存区块时，需要保存区块中交易
		trs := b.Transactions
		for _, tr := range trs {
			err := tr.DbSave()
			if err != nil {
				rollErr := o.Rollback()
				if rollErr != nil {
					panic(rollErr)
				}
				return err
			}
		}
	}

	return o.Commit()
}

func (b *Block) GetBlocks() ([]Block, error) {
	panic("implement me")
}

func (b *Block) SetBlocks(bs []Block) error {
	panic("implement me")
}

func (b *Block) SortTransactions() ([]Transaction, error) {
	panic("implement me")
}

func (b *Block) Trans2Block(data interface{}) (Block, error) {
	var err error
	c, ok := data.(config.Configer)
	if !ok {
		err = errors.New("config not type of config.Configer")
	} else {
		b.Id = c.String("id")
		b.Height, err = c.Int64("height")
		b.Timestamp, err = c.Int64("timestamp")
		b.TotalAmount, err = c.Int64("totalAmount")
		b.TotalFee, err = c.Int64("totalFee")
		b.Reward, err = c.Int64("reward")
		b.PayloadHash = c.String("payloadHash")
		b.PayloadLength, err = c.Int("payloadLength")
		b.PreviousBlock = c.String("previousBlock")
		b.BlockSignature = c.String("blockSignature")
		b.NumberOfTransactions, err = c.Int("numberOfTransactions")
		b.Generator = &Delegate{
			Username: c.String("generatorPublicKey"),
		}
		var transactions []*Transaction
		trs, err := c.DIY("transactions")
		if err != nil {
			return *b, err
		}
		for _, tr := range trs.([]interface{}) {
			t := Transaction{}
			tt, err := t.Trans2Transaction(tr)
			if err != nil {

			}
			transactions = append(transactions, &tt)
		}
		b.Transactions = transactions
	}

	return *b, err
}

func (b *Block) Trans2Object() (map[string]interface{}, error) {
	panic("implement me")
}
