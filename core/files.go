package core

import (
	"etm-go-chain/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/goinggo/mapstructure"
)

var (
	appConfig    Config
	genesisBlock models.Block
)

type Config struct {
	Port     int
	PublicIp string
	LogLevel string
	Magic    string
	PeerList []string
	Secrets  []string
}

func InitConfig() {
	file := beego.AppConfig.String("file_config")
	conf, err := config.NewConfig("json", file)
	if err != nil {
		logs.Error("【Init】 read config error! ==>", err)
		return
	}
	if appConfig, err = trans2Config(conf); err != nil {
		logs.Error("【Init】 transfer config error! ==>", err)
		return
	}

	logs.Info("【Init】 config ok!")
}

func trans2Config(conf config.Configer) (c Config, err error) {
	if c.Port, err = conf.Int("port"); err != nil {
		return c, err
	}
	c.PublicIp = conf.String("publicIp")
	c.LogLevel = conf.String("logLevel")
	c.Magic = conf.String("magic")
	c.PeerList = conf.Strings("peerList")
	c.Secrets = conf.Strings("secrets")
	//secrets, err := conf.DIY("secrets")
	//if err != nil {
	//	return c, err
	//}
	//c.Secrets = []string(secrets.([]interface{}))

	return c, nil
}

func InitGenesisBlock() {
	file := beego.AppConfig.String("file_genesisBlock")
	genesis, err := config.NewConfig("json", file)
	if err != nil {
		logs.Error("【Init】 read genesisBlock error! ==>", err)
		return
	}

	if genesisBlock, err = trans2Block(genesis); err != nil {
		logs.Error("【Init】 transfer genesisBlock error! ==>", err)
		return
	}

	logs.Info("【Init】 genesisBlock ok!")
}

func trans2Block(c config.Configer) (b models.Block, err error) {
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
		tt, err3 := trans2Transaction(tr, b)
		if err3 != nil {
			return b, err3
		}
		transactions = append(transactions, &tt)
	}
	b.Transactions = transactions

	return b, err
}

func trans2Transaction(data interface{}, b models.Block) (t models.Transaction, err error) {
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

	return t, err
}

func GetConfig() Config {
	return appConfig
}

func GetGenesisBlock() models.Block {
	return genesisBlock
}
