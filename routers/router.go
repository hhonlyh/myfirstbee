package routers

import (
	"github.com/astaxie/beego"
	"myfirstbee/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
