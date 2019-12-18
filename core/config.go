package core

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"workspace/etm-go-chain/models"
)

var (
	appConfig    config.Configer
	genesisBlock models.Block
)
func init(){
	initConfig()
	initGenesisBlock()
}

func initConfig() {
	conf, err := config.NewConfig("json", "conf/config.json")
	if err != nil {
		logs.Error(err)
		return
	}
	appConfig = conf

	logs.Info("【Init】 config ok!")
}

func initGenesisBlock() {
	genesis, err := config.NewConfig("json", "conf/genesisBlock.json")
	if err != nil {
		logs.Error("【Init】 genesisBlock err! ==>", err)
		return
	}

	b := models.Block{}
	b.Transform(genesis)
	genesisBlock = b

	logs.Info("【Init】 genesisBlock ok!")
}

func GetConfig() config.Configer {
	return appConfig
}

func GetGenesisBlock() models.Block {
	return genesisBlock
}
