package routers

import (
	"beegoLearn/quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("user/profile", &controllers, `get:Profile`)
	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)

	beego.Router("/api/user/profile", &controllers.UserController{}, `get:API_Profile`)

	beego.Router("/join", &controllers.UserController{}, `get:PageJoin`)
	beego.Router("/login", &controllers.UserController{}, `post:Login`)
	beego.Router("/register", &controllers.UserController{}, `post:Register`)
}
