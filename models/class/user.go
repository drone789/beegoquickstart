package class

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id     string `orm:"pk"`
	Nick   string
	Info   string `orm:"null"`
	Hobby  string `orm:"null"`
	Email  string `orm:"unique"`
	Avator string `orm:"null"`
	Url    string `orm:"null"`

	Posts int

	Followers int
	Following int

	Regtime time.Time `orm:"auto_now_add"`

	Password string
	Private  int
}

// func TestORM() {
// 	o := orm.NewOrm()

// 	u := User{"jike2", "geek", "someinfo", "bst", "132@123.com", "ciscoTest", "t1"}
// 	o.Insert(&u)

// 	u1 := User{Id: "jike2"}
// 	o.Read(&u1)
// 	fmt.Println(u1)

// 	u1.Nick = "lisan"
// 	o.Update(&u1)

// 	// u2 := User{Id: "jike"}
// 	// o.Read(&u2)
// 	// fmt.Println(u2)

// 	// o.Delete(&u2)

// }

const (
	PR_live = iota
	PR_login
	PR_post
)

const (
	DefaultPvt = 1<<3 - 1 //8-1=7
)

//	CRUD
//	create
//	read
//	update
//	delete

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return
}

func (u User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	return
}

func (u User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&u)
	return
}

func (u User) Delete() (err error) {
	//	xxx1 & 1110 = xxx0
	//	~x ==> ^x (-1 ^ x)
	u.Private &= ^1
	err = u.Update()
	return
}

// o.Read() 只能针对主键进行读取
func (u User) ExistId() bool {
	o := orm.NewOrm()
	if err := o.Read(&u); err == orm.ErrNoRows {
		return false
	}
	return true
}

func (u User) ExistEmail() bool {
	o := orm.NewOrm()
	// 表名-条件
	return o.QueryTable("user").Filter("Email", u.Email).Exist()
}
