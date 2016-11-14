package routers

import (
	"myblog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	//后台管理员路由
	beego.Router("/admin/index", &controllers.AdminController{}, "*:Home")
	beego.Router("/admin/category", &controllers.AdminController{}, "get:Category")
	beego.Router("/admin/topic", &controllers.AdminController{}, "*:Topic")
	beego.Router("/admin/addtopic", &controllers.AdminController{}, "*:AddTopic")
	beego.Router("/admin/liuyan", &controllers.AdminController{}, "*:ManageLiuyan")
	beego.Router("/admin/topic/modify", &controllers.AdminController{}, "*:Modify")

}
