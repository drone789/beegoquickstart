package models

import (
	"beegoLearn/quickstart/models/class"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/blog?charset=utf8")

	orm.RegisterModel(new(class.User))

	orm.RunSyncdb("default", false, true)
}
