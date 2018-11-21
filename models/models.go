package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string
	Age string
	Gender string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/gotest1?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false, true)
}