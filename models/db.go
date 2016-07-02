package models

import (
	"beegoLearn/quickstart/models/class"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.Debug = true

	switch beego.AppConfig.String("DB::db") {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8",
			beego.AppConfig.String("DB::user"),
			beego.AppConfig.String("DB::pass"),
			beego.AppConfig.String("DB::name")))
	case "sqlite":
		orm.RegisterDriver("sqlite", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::file"))
	}

	// orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/blog?charset=utf8")

	orm.RegisterModel(new(class.User))

	orm.RunSyncdb("default", false, true)
}
