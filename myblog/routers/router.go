package routers

import (
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//登录及注册路由
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/register", &controllers.LoginController{}, "*:Register")
	beego.Router("/addreply", &controllers.TopicController{}, "*:AddReply")
	beego.AutoRouter(&controllers.TopicController{})
}
