package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 查找数据
    beego.Router("/", &controllers.MainController{})
    // 添加数据
    beego.Router("/login", &controllers.LoginController{})
    // 更新
    beego.Router("/update", &controllers.UpdateController{})
    // 删除数据
    beego.Router("/delete", &controllers.DeleteController{})

}
