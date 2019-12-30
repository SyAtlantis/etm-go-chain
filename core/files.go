package core

import (
	"etm-go-chain/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/goinggo/mapstructure"
)

var (
	appConfig    config.Configer
	genesisBlock models.Block
)

func InitConfig() {
	file := beego.AppConfig.String("file_config")
	conf, err := config.NewConfig("json", file)
	if err != nil {
		logs.Error("【Init】 config error! ==>", err)
		return
	}
	appConfig = conf

	logs.Info("【Init】 config ok!")
}

func InitGenesisBlock() {
	file := beego.AppConfig.String("file_genesisBlock")
	genesis, err := config.NewConfig("json", file)
	genesisBlock, err = Trans2Block(genesis)
	if err != nil {
		logs.Error("【Init】 genesisBlock error! ==>", err)
		return
	}

	logs.Info("【Init】 genesisBlock ok!")
}

func Trans2Block(c config.Configer) (b models.Block, err error) {
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
	b.Generator = c.String("generatorPublicKey")
	var transactions []*models.Transaction
	trs, err2 := c.DIY("transactions")
	if err2 != nil {
		return b, err2
	}
	for _, tr := range trs.([]interface{}) {
		tt, err3 := Trans2Transaction(tr, b)
		if err3 != nil {
			return b, err3
		}
		transactions = append(transactions, &tt)
	}
	b.Transactions = transactions

	return b, err
}

func Trans2Transaction(data interface{}, b models.Block) (t models.Transaction, err error) {
	trData := models.TrData{}
	if err = mapstructure.Decode(data, &trData); err == nil {
		t.Id = trData.Id
		t.Type = trData.Type
		t.BlockId = &b
		t.Fee = trData.Fee
		t.Amount = trData.Amount
		t.Amount = trData.Amount
		t.Timestamp = trData.Timestamp
		t.Sender = trData.SenderPublicKey
		t.Recipient = trData.RecipientId
		
		// Args.Asset全部存在Args中
		if trData.Args != nil && len(trData.Args) > 0 {
			t.Args = trData.Args[0]
		} else {
			if trData.Asset.Delegate.Username != "" {
				t.Args = trData.Asset.Delegate.Username
			}
			if trData.Asset.Vote.Votes != nil {
				// vote中投票和取消投票分离
				t.Args = trData.Asset.Vote.Votes[0]
			}
		}

		t.Signature = trData.Signature
	}

	//obj, ok := data.(map[string]interface{})
	//
	//t.Id, ok = obj["id"].(string)
	//t.BlockId = &b
	//if ty, ok := obj["type"].(float64); ok {
	//	t.Type = uint8(ty)
	//}
	//if fee, ok := obj["fee"].(float64); ok {
	//	t.Fee = int64(fee)
	//}
	//if amount, ok := obj["amount"].(float64); ok {
	//	t.Amount = int64(amount)
	//}
	//if timestamp, ok := obj["timestamp"].(float64); ok {
	//	t.Timestamp = int64(timestamp)
	//}
	//if senderPublicKey, ok := obj["senderPublicKey"].(string); ok && senderPublicKey != "" {
	//	t.Sender = senderPublicKey
	//}
	//if recipient, ok := obj["recipientId"].(string); ok && recipient != "" {
	//	t.Recipient = recipient
	//}
	//if args, ok := obj["args"].([]interface{}); ok {
	//	t.Args = args[0].(string)
	//	//var bs []byte
	//	//if bs, err = json.Marshal(args); err != nil {
	//	//	return t, err
	//	//}
	//	//t.Args = string(bs)
	//}
	//
	////t.Message, ok = obj["message"].(string)
	//t.Signature, ok = obj["signature"].(string)
	//
	//if asset, ok := obj["asset"].(map[string]interface{}); ok {
	//	if asset["delegate"] != nil {
	//		if delegate, ok := asset["delegate"].(map[string]interface{}); ok {
	//			if delegate["username"] != nil {
	//				t.Args = delegate["username"].(string)
	//			}
	//		}
	//
	//	}
	//	//	err = mapstructure.Decode(asset, &t.Asset)
	//}
	//
	//if !ok {
	//	err = errors.New("Transform data to Transaction error")
	//}
	return t, err
}

func GetConfig() config.Configer {
	return appConfig
}

func GetGenesisBlock() models.Block {
	return genesisBlock
}
