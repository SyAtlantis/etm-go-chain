package main

import (
	"github.com/astaxie/beego"
)

func main() {
	// 初始化配置
	//Init()

	// 模块运行
	//Setup()

	//err, _ := event.Fire("bind", event.M{})
	//if err != nil {
	//	panic(err)
	//}

	// 启动
	beego.Run()

}
