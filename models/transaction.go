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
	"github.com/astaxie/beego/orm"
	"reflect"
	"sort"
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
	GetTransaction() (Transaction, error)
	SetTransaction() error
	Trans2Transaction(data interface{}) (Transaction, error)
	Trans2Object() (map[string]interface{}, error)
}

type Transaction struct {
	Key       int64    `orm:"pk;auto"`
	Id        string   `json:"id"`
	Type      uint8    `json:"type"`
	BlockId   *Block   `json:"blockId" orm:"rel(fk);column(block_id)"`
	Fee       int64    `json:"fee"`
	Amount    int64    `json:"amount"`
	Timestamp int64    `json:"timestamp"`
	Sender    *Account `json:"sender" orm:"rel(fk);null"`
	Recipient *Account `json:"recipient" orm:"rel(fk);null"`
	Args      string   `json:"args"`
	Message   string   `json:"message"`
	Signature string   `json:"signature"`
}

type TrData struct {
	Type        uint8
	Amount      int64
	Fee         int64
	Timestamp   int64
	RecipientId string
	//Asset          Asset
	Args           []string
	Message        string
	Sender         Account
	Keypair        utils.Keypair
	SecondKeypair  utils.Keypair
	Votes          []string
	Username       string
	Name           string
	Desc           string
	Maximun        string
	Precision      byte
	Strategy       string
	AllawWriteOff  byte
	AllowWhiteList byte
	AllowBlackList byte
	Currency       string
	UiaAmount      string
	FlagType       byte
	Flag           byte
	Operator       string
	List           []string
}

type SubTr interface {
	create(tr *Transaction, data TrData)
	getBytes(tr *Transaction) []byte
}

var trTypes = make(map[uint8]SubTr)

func (t *Transaction) IsEmpty() bool {
	return reflect.DeepEqual(t, Transaction{})
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
	t.Sender = &Account{PublicKey: data.Sender.PublicKey}
	//t.Asset = data.Asset
	args, err := json.Marshal(data.Args)
	t.Args = string(args)
	t.Message = data.Message

	trTypes[data.Type].create(t, data) //构建对应子交易数据

	t.Signature, err = t.GetSignature(data.Keypair)
	//if data.Type != 1 && !data.SecondKeypair.IsEmpty() {
	//	t.SignSignature = t.GetSignature(data.SecondKeypair)
	//}

	t.Id, err = t.GetId();
	return err
}

func (t *Transaction) GetBytes() ([]byte, error) {
	var err error
	assetBytes := trTypes[t.Type].getBytes(t)
	assetSize := len(assetBytes)

	//bb := bytes.NewBuffer(make([]byte, size+assetSize))
	bb := bytes.NewBuffer([]byte{})

	err = binary.Write(bb, binary.LittleEndian, uint8(t.Type))
	err = binary.Write(bb, binary.LittleEndian, uint32(t.Timestamp))

	if !t.Sender.IsEmpty() {
		senderPublicKeyBytes, _ := hex.DecodeString(t.Sender.PublicKey)
		bb.Write(senderPublicKeyBytes)
	}

	//if !t.Requester.IsEmpty() {
	//	requesterPublicKeyBytes, _ := hex.DecodeString(t.Requester.PublicKey)
	//	bb.Write(requesterPublicKeyBytes)
	//}

	if !t.Recipient.IsEmpty() {
		bb.WriteString(t.Recipient.Address)
	} else {
		for i := 0; i < 8; i++ {
			bb.WriteByte(0);
		}
	}

	err = binary.Write(bb, binary.LittleEndian, uint64(t.Amount))

	if t.Message != "" {
		bb.WriteString(string(t.Message))
	}

	if t.Args != "nil" {
		var args []byte
		err = json.Unmarshal(args, t.Args)
		for i := 0; i < len(args); i++ {
			bb.WriteByte(args[i])
		}
	}

	if assetSize > 0 {
		bb.Write(assetBytes)
	}

	//if !skipSignature && tr.Signature != "" {
	//	signatureBytes, _ := hex.DecodeString(tr.Signature)
	//	bb.Write(signatureBytes)
	//}
	//
	//if !skipSecondSignature && tr.SignSignature != "" {
	//	signSignatureBytes, _ := hex.DecodeString(tr.SignSignature)
	//	bb.Write(signSignatureBytes)
	//}

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

func (t *Transaction) Trans2Transaction(data interface{}) (Transaction, error) {
	var err error
	o, ok := data.(map[string]interface{})

	t.Id, ok = o["id"].(string)
	if fv, ok := o["type"].(float64); ok {
		t.Type = uint8(fv)
	}
	id, ok := o["blockId"].(string)
	t.BlockId = &Block{Id: id,}
	t.Fee, ok = o["fee"].(int64)
	t.Amount, ok = o["amount"].(int64)
	t.Timestamp, ok = o["timestamp"].(int64)
	senderPublicKey, ok := o["senderPublicKey"].(string)
	t.Sender = &Account{PublicKey: senderPublicKey,}
	recipient, ok := o["recipientId"].(string)
	t.Recipient = &Account{Address: recipient,}
	t.Args, ok = o["args"].(string)
	t.Message, ok = o["message"].(string)
	t.Signature, ok = o["signature"].(string)

	if !ok {
		err = errors.New("Transform data to Transaction error")
	}
	return *t, err
}

func (t *Transaction) Trans2Object() (map[string]interface{}, error) {
	panic("implement me")
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
