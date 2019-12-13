package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"

	"workspace/etm-go-chain/core"
)

func init() {
	// init logger
	initLogger()

	// init swagger
	initSwagger()

	// init i18n
	initI18n()

	// init catch
	initCatch()

	// init db
	initDB()

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

	logs.Info("Init logger ok!")
}

func initSwagger() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func initI18n() {
	_ = i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	_ = i18n.SetMessage("en-US", "conf/locale_en-US.ini")
}

func initCatch() {
	core.InitRedis()
}

func initDB() {
	core.InitSqlite()
}
