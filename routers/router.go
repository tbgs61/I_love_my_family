package routers

import (
	"I_love_my_family/love_fa/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
