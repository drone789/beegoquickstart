package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `orm:pk`
	Nick  string
	Email string
}

func TestORM() {
	o := orm.NewOrm()

	u := User{"jike", "geek", "132@123.com"}
	o.Insert(&u)

	u1 := User{Id: "jike"}
	o.Read(&u1)
	fmt.Println(u1)

	u1.Nick = "lisan"
	o.Update(&u1)

	u2 := User{Id: "jike"}
	o.Read(&u2)
	fmt.Println(u2)

	o.Delete(&u2)

}
