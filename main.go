package main

import (
	_ "blog/routers"
	_ "blog/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

