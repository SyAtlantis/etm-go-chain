package models

import (
	"github.com/astaxie/beego/config"
)

var modelList = make(map[string]interface{})

type Models interface {
	NewModel() interface{}
}

func RegisterModels(name string, model Models) {
	modelList[name] = model.NewModel()
}

func GetModels() map[string]interface{} {
	return modelList
}

func GetModel(name string) interface{} {
	return modelList[name]
}

func (b *Block) Transform(c config.Configer) {
	var err error
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
		return
	}
	for _, tr := range trs.([]interface{}) {
		t := Transaction{}
		t.Transform(tr.(map[string]interface{}))
		transactions = append(transactions, &t)
	}
	b.Transactions = transactions
}

func (t *Transaction) Transform(o map[string]interface{}) {
	var ok = true
	t.Id, ok = o["id"].(string)
	t.Type, ok = o["type"].(uint8)
	id, ok := o["blockId"].(string)
	t.BlockId = &Block{Id: id,}
	t.Fee, ok = o["fee"].(int64)
	t.Amount, ok = o["amount"].(int64)
	t.Timestamp, ok = o["timestamp"].(int64)
	serder, ok := o["senderPublicKey"].(string)
	t.Sender = &Account{PublicKey: serder,}
	recipient, ok := o["recipientId"].(string)
	t.Recipient = &Account{Address: recipient,}
	t.Args, ok = o["args"].(string)
	t.Message, ok = o["message"].(string)
	t.Signature, ok = o["signature"].(string)

	if !ok {
		return
	}
}
