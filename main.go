package main

import (
	"beegoLearn/quickstart/models/class"
	_ "beegoLearn/quickstart/routers"
	"encoding/gob"
	// "fmt"
	_ "beegoLearn/quickstart/models"
	// "beegoLearn/quickstart/models/class"
	"github.com/astaxie/beego"
)

func init() {
	gob.Register(class.User{})
}
func main() {

	// beego.AppConfig.Set("key", "value")
	// // x := beego.AppConfig.String("mysqldb")
	// // fmt.Println(x)
	// beego.SetStaticPath("download", "download")

	// class.TestORM()
	beego.Run()
}
