package main

import (
	_ "myapp/routers"
	"github.com/astaxie/beego"

)

func main() {
	beego.SetStaticPath("/base","/view/base")
	beego.SetStaticPath("/img","/static/image")
	beego.SetStaticPath("/css","/static/css")

	beego.Run()

}