package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

type UpdateController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 获取orm对象
	o := orm.NewOrm()

	// 获取结构体对象
	user := models.User{}

	// 对结构体进行赋值
	user.Name = "chenrun"
	user.Age = "20"
	user.Gender = "男"

	// 插入(返回两个，第一个是影响的行数，第二个有没有错误)
	_, err := o.Insert(&user)

	// 判断错误
	if err != nil {
		beego.Info("插入数据失败", err)
		return
	}

	this.Data["username"] = "chenrun"
	this.TplName = "login.html"
}

func (this *UpdateController) Get() {
	// 获取orm对象
	o := orm.NewOrm()

	// 需要更新的结构
	user := models.User{}

	// 查找到需要更新的数据
	user.Id = 1
	err := o.Read(&user)

	if err == nil {
		user.Name = "penghuixian"
		user.Age = "18"
		user.Gender = "女"
		// 更新
		_, err := o.Update(&user)
		if err != nil {
			beego.Info("更新失败", err)
			return
		}
	}

	this.TplName = "index.html"
}
