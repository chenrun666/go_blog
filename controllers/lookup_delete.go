package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"blog/models"
)

type MainController struct {
	beego.Controller
}

type DeleteController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// 获取orm对象
	o := orm.NewOrm()

	// 获取查询对象
	user := models.User{}

	// 指定查询字段
	// 根据ID查询
	// user.Id = 1

	// 根据其他字段
	user.Name = "chenrun"

	// 查询
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	beego.Info("查询成功", user)


	c.TplName = "index.html"
}

func (this *DeleteController) Get() {
	// 获取orm对象
	o := orm.NewOrm()

	// 获取表对象
	user := models.User{}

	// 指定要删除的字段名
	user.Id = 2

	// 进行删除操作
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除失败", err)
		return
	}

	this.TplName = "index.html"
}
