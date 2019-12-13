package main

import (
	"github.com/astaxie/beego"
	_ "workspace/etm-go-chain/models"
	_ "workspace/etm-go-chain/routers"
)

func main() {
	// 初始化配置
	//Init()

	// 模块运行
	Setup()

	// 启动
	beego.Run()

}
