package core

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() {
	dbName := beego.AppConfig.String("db_name")
	dbCover, err := beego.AppConfig.Bool("db_cover")

	err = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err = orm.RegisterDataBase("default", "sqlite3", dbName)
	orm.SetMaxIdleConns("default", 50)
	orm.SetMaxOpenConns("default", 200)

	//自动建表
	err = orm.RunSyncdb("default", dbCover, true)

	//设置数据库时区
	orm.DefaultTimeLoc = time.UTC

	if err != nil {
		logs.Error("【Init】 sqlite3 database failure! ==>", err)
		return
	}

	logs.Info("【Init】 sqlite3 database ok!")
}
