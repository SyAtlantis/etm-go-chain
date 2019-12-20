package core

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"

	"etm-go-chain/models"
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
	if err != nil {
		logs.Error("【Init】 genesisBlock error! ==>", err)
		return
	}
	newBlock := models.Block{}
	b, err := newBlock.Trans2Block(genesis)
	genesisBlock = b

	logs.Info("【Init】 genesisBlock ok!")
}

func GetConfig() config.Configer {
	return appConfig
}

func GetGenesisBlock() models.Block {
	return genesisBlock
}
