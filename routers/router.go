package routers

import (
	"cloud-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/add", &controllers.AddController{})
    beego.Router("/nodes", &controllers.NodeController{})
    beego.Router("/instances", &controllers.InstanceController{})
}
