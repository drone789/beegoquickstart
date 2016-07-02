package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.Data["userid"] = "geek"
	c.Data["tag"] = "xxx"
	c.Data["hobby"] = []string{"bsk", "football"}
	c.TplName = "user/profile.html"
}
