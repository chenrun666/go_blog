package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 返回登陆页面
	this.TplName = "account/login.html"

}

func (this *LoginController) Post() {

	// 验证登陆路提交的数据
	username := this.GetString("username")
	password := this.GetString("password")

	// orm对象
	o := orm.NewOrm()
	// 获取表对象
	userinfo := models.UserInfo{}
	// 指定查询字段
	userinfo.Name = username
	// 查询
	err := o.Read(&userinfo, "Name")

	if err == nil {
		// 没有错误，有该用户，判断密码是否正确
		if userinfo.Password == password {
			this.Redirect("/", 302)
			return
		} else {
			// 密码错误

		}
	}
	// 没有用户名

	this.TplName = "account/login.html"

}

func (this *RegisterController) Get() {
	// 返回注册页面
	this.TplName = "account/register.html"

}

func (this *RegisterController) Post() {
	// 验证注册的数据，保存到数据库中
	this.TplName = "index.html"

}
