package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}



func (this *LoginController) Get() {
	this.Data["username"] = "chenrun"
	this.TplName = "login.html"
}