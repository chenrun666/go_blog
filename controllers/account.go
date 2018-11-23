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

	// 初步验证用户名和密码
	if username == "" || password == "" {
		beego.Info("用户名或密码不能为空")
		this.Redirect("/redister", 302)
		return
	}

	// orm对象
	o := orm.NewOrm()
	// 获取表对象
	userinfo := models.UserInfo{}
	// 指定查询字段
	userinfo.Name = username
	// 查询
	err := o.Read(&userinfo, "Name")
	if err != nil {
		beego.Info("查询失败")
		this.Ctx.WriteString("没有该用户！！！")
		return
	}
	if userinfo.Password == password {
		this.Redirect("/", 302)
		return
	}
	beego.Info("密码错误")
	this.Ctx.WriteString("密码错误")
	return

}

func (this *RegisterController) Get() {
	// 返回注册页面
	this.TplName = "account/register.html"

}

func (this *RegisterController) Post() {

	// 获取提交的数据
	username := this.GetString("username")
	password := this.GetString("password")
	againpwd := this.GetString("againpwd")

	// 校验数据
	if username == "" || password == "" || againpwd == "" {
		beego.Info("数据不能为空")
		this.Ctx.WriteString("数据不能为空")
		return
	}
	if password != againpwd {
		beego.Info("两次密码不一致")
		this.Ctx.WriteString("两次密码不一致")
		return
	}

	// 获取orm对象
	o := orm.NewOrm()
	// 获取操作的表
	userdata := models.UserInfo{}
	// 赋值
	userdata.Name = username
	userdata.Password = password
	// 插入数据库中
	_, err := o.Insert(&userdata)
	if err != nil {
		beego.Info("添加失败")
		this.Ctx.WriteString("注册失败")
		return
	}

	// 注册成功跳转页面
	this.Redirect("/login", 301)

}
