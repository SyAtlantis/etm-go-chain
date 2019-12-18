package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	_ "workspace/etm-go-chain/routers"
)

func init() {
	initLogger()
	initSwagger()
	initI18n()
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
	if err !=nil{
		logs.Error("【Init】 i18n error! ==>", err)
	}

	logs.Info("【Init】 i18n ok!")
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

func main() {
	// 初始化配置
	//Init()

	// 模块运行
	//Setup()

	//err, _ := event.Fire("bind", event.M{})
	//if err != nil {
	//	panic(err)
	//}
	//_, _ = event.Fire("load", event.M{"name": "bbb"})
	//time.Sleep(10*time.Second)
	//_, _ = event.Fire("ready", event.M{"name": "ccc"})

	// 启动
	beego.Run()

}
