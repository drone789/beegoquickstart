package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `orm:"pk"`
	Nick  string
	Email string
	Test  string
	T1    string
}

func TestORM() {
	o := orm.NewOrm()

	u := User{"jike2", "geek", "132@123.com", "ciscoTest", "t1"}
	o.Insert(&u)

	u1 := User{Id: "jike2"}
	o.Read(&u1)
	fmt.Println(u1)

	u1.Nick = "lisan"
	o.Update(&u1)

	// u2 := User{Id: "jike"}
	// o.Read(&u2)
	// fmt.Println(u2)

	// o.Delete(&u2)

}
