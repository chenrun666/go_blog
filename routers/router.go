package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 查找数据
    beego.Router("/look", &controllers.MainController{})
    // 添加数据
    beego.Router("/insert", &controllers.InsertController{})
    // 更新
    beego.Router("/update", &controllers.UpdateController{})
    // 删除数据
    beego.Router("/delete", &controllers.DeleteController{})


    // 主页面
    beego.Router("/", &controllers.IndexController{})

    // 登陆注册的url
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/register", &controllers.RegisterController{})

}
