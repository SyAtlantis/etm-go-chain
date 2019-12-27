package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"etm-go-chain/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"sort"
)

const (
	TRANSFER   uint8 = 0
	DELEGATE   uint8 = 2
	UNDELEGATE uint8 = 120
	LOCK       uint8 = 101
	UNLOCK     uint8 = 102
	VOTE       uint8 = 3
	DELAY      uint8 = 110
	SECOND     uint8 = 1
	MULTI      uint8 = 4

	UIA_ISSUER   uint8 = 9
	UIA_ASSET    uint8 = 10
	UIA_FALG     uint8 = 11
	UIA_ACL      uint8 = 12
	UIA_ISSUE    uint8 = 13
	UIA_TRANSFER uint8 = 14
)

func init() {
	orm.RegisterModel(new(Transaction))
}

type iTransaction interface {
	IsEmpty() bool
	Create(data TrData) error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature(utils.Keypair) (string, error)
	VerifySignature() (bool error)
	Verify() (bool error)
	Process() error
	Apply() error
	Undo() error
	ApplyUnconfirmed() error
	UndoUnconfirmed() error
	GetTransaction() (Transaction, error)
	SetTransaction() error
}

type Transaction struct {
	Key       int64  `orm:"pk;auto"`
	Id        string `json:"id" `
	Type      uint8  `json:"type"`
	BlockId   *Block `json:"blockId" orm:"rel(fk);column(block_id)"`
	Fee       int64  `json:"fee"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Args      string `json:"args"`
	Signature string `json:"signature"`
}

type TrData struct {
	Id              string
	Type            uint8
	Amount          int64
	Fee             int64
	Timestamp       int64
	RecipientId     string
	SenderPublicKey string
	Asset           Asset
	Args            []string
	Message         string
	Signature       string
	Sender          Account
	Keypair         utils.Keypair
	SecondKeypair   utils.Keypair
	Votes           []string
	Username        string
	Name            string
	Desc            string
	Maximum         string
	Precision       byte
	Strategy        string
	AllowWriteOff   byte
	AllowWhiteList  byte
	AllowBlackList  byte
	Currency        string
	UiaAmount       string
	FlagType        byte
	Flag            byte
	Operator        string
	List            []string
}

type Asset struct {
	Vote     TrVote
	Delegate TrDelegate
	//Signature   TrSecond
	//UiaIssuer   TrUiaIssuer
	//UiaAsset    TrUiaAsset
	//UiaFlags    TrUiaFlags
	//UiaAcl      TrUiaAcl
	//UiaIssue    TrUiaIssue
	//UiaTransfer TrUiaTransfer
}

type TrVote struct {
	Votes []string
}

type TrDelegate struct {
	Username string
}

type SubTr interface {
	create(tr *Transaction, data TrData) error
	getBytes(tr *Transaction) ([]byte, error)
}

var trTypes = make(map[uint8]SubTr)

func RegisterTrs(trType uint8, tr SubTr) {
	trTypes[trType] = tr
}

func (t *Transaction) IsEmpty() bool {
	return t == nil || reflect.DeepEqual(t, Transaction{})
}

func (t *Transaction) Create(data TrData) error {
	var err error
	if data.Sender.IsEmpty() {
		return err
	}
	if data.Keypair.IsEmpty() {
		return err
	}

	t.Type = data.Type
	t.Amount = 0
	t.Fee = data.Fee
	t.Timestamp = data.Timestamp
	t.Sender = data.Sender.PublicKey
	//t.Asset = data.Asset
	args, err := json.Marshal(data.Args)
	t.Args = string(args)
	//t.Message = data.Message

	//err = trTypes[data.Type].create(t, data) //构建对应子交易数据

	t.Signature, err = t.GetSignature(data.Keypair)
	//if data.Type != 1 && !data.SecondKeypair.IsEmpty() {
	//	t.SignSignature = t.GetSignature(data.SecondKeypair)
	//}

	t.Id, err = t.GetId();
	return err
}

func (t *Transaction) GetBytes() ([]byte, error) {
	var err error
	assetBytes, err := trTypes[t.Type].getBytes(t)
	assetSize := len(assetBytes)

	bb := bytes.NewBuffer([]byte{})

	err = binary.Write(bb, binary.LittleEndian, uint8(t.Type))
	err = binary.Write(bb, binary.LittleEndian, uint32(t.Timestamp))

	if t.Sender != "" {
		senderPublicKeyBytes, _ := hex.DecodeString(t.Sender)
		bb.Write(senderPublicKeyBytes)
	}

	if t.Recipient != "" {
		bb.WriteString(t.Recipient)
	} else {
		for i := 0; i < 8; i++ {
			bb.WriteByte(0);
		}
	}

	err = binary.Write(bb, binary.LittleEndian, uint64(t.Amount))

	//if t.Message != "" {
	//	bb.WriteString(string(t.Message))
	//}

	if t.Args != "" {
		//var args []string
		//err = json.Unmarshal([]byte(t.Args), &args)
		//for i := 0; i < len(args); i++ {
		//	bb.WriteString(args[i])
		//}
		bb.WriteString(t.Args)
	}

	if assetSize > 0 {
		bb.Write(assetBytes)
	}

	if t.Signature != "" {
		signatureBytes, _ := hex.DecodeString(t.Signature)
		bb.Write(signatureBytes)
	}

	return bb.Bytes(), err
}

func (t *Transaction) GetHash() ([32]byte, error) {
	bs, err := t.GetBytes()
	hash := sha256.Sum256(bs)
	return hash, err
}

func (t *Transaction) GetId() (string, error) {
	hash, err := t.GetHash()
	return fmt.Sprintf("%x", hash), err
}

func (t *Transaction) GetSignature(keypair utils.Keypair) (string, error) {
	hash, err := t.GetHash()
	sign := ed.Sign(hash[:], keypair)
	return fmt.Sprintf("%x", sign), err
}

func (t *Transaction) VerifySignature() (bool error) {
	panic("implement me")
}

func (t *Transaction) Verify() (bool error) {
	panic("implement me")
}

func (t *Transaction) Process() error {
	panic("implement me")
}

func (t *Transaction) Apply() error {
	logs.Debug("Transaction Apply")
	return nil
}

func (t *Transaction) Undo() error {
	panic("implement me")
}

func (t *Transaction) ApplyUnconfirmed() error {
	panic("implement me")
}

func (t *Transaction) UndoUnconfirmed() error {
	panic("implement me")
}

func (t *Transaction) GetTransaction() (Transaction, error) {
	o := orm.NewOrm()
	err := o.Read(&t)
	return *t, err
}

func (t *Transaction) SetTransaction() error {
	o := orm.NewOrm()
	created, id, err := o.ReadOrCreate(t, "Id")
	if err != nil {
		return err
	}
	if !created {
		err = errors.New("This transaction already exists in the db:" + string(id))
	}
	return err
}

type Trs []*Transaction

type iTransactions interface {
	Sort()
}

func (trs Trs) Len() int {
	return len(trs)
}

func (trs Trs) Swap(i, j int) {
	trs[i], trs[j] = trs[j], trs[i]
}

func (trs Trs) Less(i, j int) bool {
	if trs[i].Type != trs[j].Type {
		if trs[i].Type == 1 {
			return true
		}
		if trs[j].Type == 1 {
			return false
		}
		return trs[i].Type > trs[j].Type
	}
	if trs[i].Amount != trs[j].Amount {
		return trs[i].Amount > trs[j].Amount
	}
	return trs[i].Id > trs[j].Id
}

func (trs Trs) Sort() {
	sort.Sort(trs)
}
