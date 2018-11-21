package models

import (
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
)

// 表的设计有关
type User struct {
	Id int
	Name string
	Age string
	Gender string
}

func init() {
	// 设置数据库的基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/gotest1?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User))
	// 生成表
	orm.RunSyncdb("default",false, true)
}