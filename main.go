package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"github.com/gookit/event"
	"time"

	"etm-go-chain/core"
	"etm-go-chain/modules"
	_ "etm-go-chain/routers"
)

func init() {
	initLogger()
	initSwagger()
	initI18n()

	//core.InitCache()
	core.InitDb()
	core.InitConfig()
	core.InitGenesisBlock()
	modules.InitModules()
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
		logs.Error("【Init】 Set logger file config error! ==>", err)
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
	var err error
	err = i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	err = i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	if err != nil {
		logs.Error("【Init】 i18n error! ==>", err)
	}

	logs.Info("【Init】 i18n ok!")
}

func main() {
	// 启动线程执行出块流程
	go start()

	// 启动服务
	beego.Run()
}

func start() {
	time.Sleep(1 * time.Second)
	if err, _ := event.Fire("bind", event.M{}); err != nil {
		panic(err)
	}

	//_, _ = event.Fire("load", event.M{"name": "bbb"})
	//time.Sleep(10*time.Second)
	//_, _ = event.Fire("ready", event.M{"name": "ccc"})
}
