package main

import (
	_ "beegoLearn/quickstart/routers"
	// "fmt"
	_ "beegoLearn/quickstart/models"
	"beegoLearn/quickstart/models/class"
	"github.com/astaxie/beego"
)

func main() {

	beego.AppConfig.Set("key", "value")
	// x := beego.AppConfig.String("mysqldb")
	// fmt.Println(x)
	beego.SetStaticPath("download", "download")

	class.TestORM()
	beego.Run()
}
