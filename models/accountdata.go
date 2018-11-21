package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id int
	Name string
	Password string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/go_blog?charset=utf8")
	orm.RegisterModel(new(UserInfo))
	orm.RunSyncdb("default", false, true)

}