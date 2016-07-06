package controllers

import (
// "github.com/astaxie/beego"
)

type UserController struct {
	// beego.Controller
	BaseController
}

func (c *UserController) Profile() {
	c.Data["userid"] = "geek"
	c.Data["tag"] = "xxx"
	c.Data["hobby"] = []string{"bsk", "football"}
	c.TplName = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplName = "user/join.html"
}

func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplName = "user/setting.html"
}
