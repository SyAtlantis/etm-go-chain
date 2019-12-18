package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"workspace/etm-go-chain/core"
	_ "workspace/etm-go-chain/routers"

	_ "workspace/etm-go-chain/models"
	_ "workspace/etm-go-chain/modules/accounts"
	_ "workspace/etm-go-chain/modules/blocks"
	_ "workspace/etm-go-chain/modules/peers"
	_ "workspace/etm-go-chain/modules/system"
	_ "workspace/etm-go-chain/modules/transactions"
)

func init() {
	initLogger()
	initSwagger()
	initI18n()
	initCatch()
	initDb()
	initConfig()
	initGenesis()
	initModels()
}

func initLogger() {
	// log 的配置
	configStr := `{ 
		"filename" : "logs/test.log", 
		"maxlines" : 1000, 
		"maxsize" : 10240
	}`
	err := logs.SetLogger(logs.AdapterFile, configStr)
	if err != nil {
		logs.Error("Set logger file config error!")
	}
	// log打印文件名和行数
	logs.EnableFuncCallDepth(true)

	logs.Info("【Init】 logger ok!")
}

func initSwagger() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		logs.Info("【Init】 swagger ok!")
	}
}

func initI18n() {
	_ = i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	_ = i18n.SetMessage("en-US", "conf/locale_en-US.ini")

	logs.Info("【Init】 i18n ok!")
}

func initCatch() {
	core.InitRedis()
}

func initDb() {
	core.InitSqlite()
}

func initConfig() {
	core.InitConfig()
}

func initGenesis() {
	core.InitGenesisBlock()
}

func initModels() {
	//modelList := models.GetModels()
	//
	//
	//iBlock := modelList["blocks"]
	//blockslist, ok := iBlock.(blocks.Blocks)
	//if ok {
	//	logs.Info(modelList, blockslist.VerifyBlock())
	//} else{
	//	logs.Info("aaaaaaa")
	//}

	//logs.Info(modelList, blocksist)
}
