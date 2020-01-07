package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"etm-go-chain/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
)

const maxPayloadLength = 8 * 1024 * 1024

var ed = utils.Ed{}
var blockStatus = utils.BlockStatus{}

func init() {
	orm.RegisterModel(new(Block))
}

type iBlock interface {
	IsEmpty() bool
	Create(BlockData) error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature(utils.Keypair) (string, error)
	VerifySignature() (bool error)
	GetBlock() (Block, error)
	SetBlock() error
}

type Block struct {
	Id                   string `json:"id"`
	Height               int64  `json:"height" orm:"pk"`
	Timestamp            int64  `json:"timestamp"`
	TotalAmount          int64  `json:"totalAmount" orm:"column(totalAmount)"`
	TotalFee             int64  `json:"totalFee" orm:"column(totalFee)"`
	Reward               int64  `json:"reward"`
	PayloadHash          string `json:"payloadHash" orm:"column(payloadHash)"`
	PayloadLength        int    `json:"payloadLength" orm:"column(payloadLength)"`
	PreviousBlock        string `json:"previousBlock" orm:"column(previousBlock)"`
	Generator            string `json:"generator"`
	BlockSignature       string `json:"blockSignature" orm:"column(blockSignature)"`
	NumberOfTransactions int    `json:"numberOfTransactions" orm:"column(numberOfTransactions)"`
	Transactions         Trs    `json:"transactions" orm:"reverse(many)"`
}

type BlockData struct {
	Transactions  Trs
	PreviousBlock Block
	Timestamp     int64
	Keypair       utils.Keypair
}

func (b *Block) IsEmpty() bool {
	return b == nil || reflect.DeepEqual(b, Block{})
}

func (b *Block) Create(data BlockData) error {
	trs := data.Transactions
	trs.Sort()
	var nextHeight int64 = 1
	if !data.PreviousBlock.IsEmpty() {
		nextHeight = data.PreviousBlock.Height + 1
	}
	reward := blockStatus.CalcReward(nextHeight)
	var totalFee int64
	var totalAmount int64
	var size int
	var err error

	blockTrs := trs[:]
	payloadHash := sha256.New()

	for i := 0; i < len(trs); i++ {
		bs, err := trs[i].GetBytes()
		if err != nil {
			return err
		}

		if size+len(bs) > maxPayloadLength {
			blockTrs = trs[:i]
			break
		}

		size += len(bs)
		totalFee += trs[i].Fee
		totalAmount += trs[i].Amount

		payloadHash.Write(bs)
	}

	b.TotalAmount = totalAmount
	b.TotalFee = totalFee
	b.Reward = reward
	b.PayloadHash = fmt.Sprintf("%x", payloadHash.Sum([]byte{}))
	b.Timestamp = data.Timestamp
	b.NumberOfTransactions = len(blockTrs)
	b.PayloadLength = size
	b.PreviousBlock = data.PreviousBlock.Id
	b.Generator = fmt.Sprintf("%x", data.Keypair.PublicKey)
	b.Transactions = blockTrs

	b.BlockSignature, err = b.GetSignature(data.Keypair)

	b.Id, err = b.GetId()
	b.Height = nextHeight

	return err
}

func (b *Block) GetBytes() ([]byte, error) {
	var err error
	bb := bytes.NewBuffer([]byte{})

	err = binary.Write(bb, binary.LittleEndian, uint32(0)) //version
	err = binary.Write(bb, binary.LittleEndian, uint32(b.Timestamp))

	if b.PreviousBlock != "" {
		bb.WriteString(b.PreviousBlock)
	} else {
		bb.WriteString("0")
	}

	err = binary.Write(bb, binary.LittleEndian, uint32(b.NumberOfTransactions))
	err = binary.Write(bb, binary.LittleEndian, uint64(b.TotalAmount))
	err = binary.Write(bb, binary.LittleEndian, uint64(b.TotalFee))
	err = binary.Write(bb, binary.LittleEndian, uint64(b.Reward))
	err = binary.Write(bb, binary.LittleEndian, uint32(b.PayloadLength))

	payloadHashBytes, _ := hex.DecodeString(b.PayloadHash)
	bb.Write(payloadHashBytes)

	generatorPublicKeyBytes, _ := hex.DecodeString(b.Generator)
	bb.Write(generatorPublicKeyBytes)

	if b.BlockSignature != "" {
		blockSignatureBytes, _ := hex.DecodeString(b.BlockSignature)
		bb.Write(blockSignatureBytes)
	}

	return bb.Bytes(), err
}

func (b *Block) GetHash() ([32]byte, error) {
	bs, err := b.GetBytes()
	hash := sha256.Sum256(bs)
	return hash, err
}

func (b *Block) GetId() (string, error) {
	hash, err := b.GetHash()
	return fmt.Sprintf("%x", hash), err
}

func (b *Block) GetSignature(keypair utils.Keypair) (string, error) {
	hash, err := b.GetHash()
	sign := ed.Sign(hash[:], keypair)
	return fmt.Sprintf("%x", sign), err
}

func (b *Block) VerifySignature() (bool error) {
	panic("implement me")
}

func (b *Block) GetBlock() (Block, error) {
	o := orm.NewOrm()
	err := o.Read(&b, "Height")
	return *b, err
}

func (b *Block) SetBlock() error {
	var err error
	o := orm.NewOrm()

	var rollback = func() {
		err = o.Rollback()
		if err != nil {
			panic(err)
		}
	}

	//保存区块错误时，需要事物回滚
	err = o.Begin()
	if err != nil {
		rollback()
		return err
	}

	var created bool
	created, _, err = o.ReadOrCreate(b, "Id")
	if err != nil {
		rollback()
		return err
	}

	//保存区块时，需要保存区块中交易
	if created {
		trs := b.Transactions
		//有回滚，必须在同一个数据库操作中执行
		_, err := o.InsertMulti(20, trs)
		if err != nil {
			rollback()
			return err
		}
	}

	return o.Commit()
}
