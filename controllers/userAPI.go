package controllers

import (
	"beegoLearn/quickstart/models/class"
	// "github.com/astaxie/beego/orm"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	. "fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

func (c *UserController) API_Profile() {
	type userA struct {
		UserID string
		Hobby  []string
	}

	u := userA{
		"jike",
		[]string{"tom", "jerry"},
	}

	c.Data["json"] = u

	c.ServeJSON()
}

type RET struct {
	OK      bool        `json:"success"`
	Content interface{} `json:content`
}

func (c *UserController) Register() {
	ret := RET{
		OK:      true,
		Content: "success",
	}
	defer func() {
		c.Data["json"] = ret
		c.ServeJSON()
	}()

	id := c.GetString("userid")
	nick := c.GetString("nick")
	pwd1 := c.GetString("password1")
	pwd2 := c.GetString("password2")
	email := c.GetString("email")

	if len(nick) < 1 {
		nick = id
	}

	valid := validation.Validation{}

	valid.Email(email, "Email")

	valid.Required(id, "Userid")
	valid.Required(pwd1, "Password")
	valid.Required(pwd2, "Password2")

	valid.MaxSize(id, 20, "UserID")
	valid.MaxSize(nick, 30, "Nick")

	switch {

	case valid.HasErrors():

	case pwd1 != pwd2:
		valid.Error("密码不一致")

	default:
		u := &class.User{
			Id:       id,
			Email:    email,
			Password: PwGen(pwd1),
			Regtime:  time.Now(),
			Private:  class.DefaultPvt,
		}

		switch {
		case u.ExistId():
			valid.Error("用户名被占用")
		case u.ExistEmail():
			valid.Error("邮箱已注册")
		default:
			err := u.Create()
			if err == nil {
				return
			}
			valid.Error(Sprintf("%v", err))
		}

	}

	ret.OK = false
	ret.Content = valid.Errors[0].Key + valid.Errors[0].Message

	return
}
func PwGen(pass string) string {
	salt := strconv.FormatInt(time.Now().UnixNano()%9000+1000, 10)
	return Base64Encode(Sha1(Md5(pass)+salt) + salt)
}

func Sha1(s string) string {
	return Sprintf("%x", sha1.Sum([]byte(s)))
}

func Md5(s string) string {
	return Sprintf("%x", md5.Sum([]byte(s)))
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	res, _ := base64.StdEncoding.DecodeString(s)
	return string(res)
}

func (c *UserController) Login() {
	c.ServeJSON()
}
