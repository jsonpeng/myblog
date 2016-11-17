package routers

import (
	"myblog/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	//前端用户路由
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/category", &controllers.MainController{}, "*:Category")
	beego.Router("/shouji", &controllers.MainController{}, "*:Shouji")
	beego.Router("/liuyan", &controllers.MainController{}, "*:Liuyan")
	beego.Router("/about", &controllers.MainController{}, "*:About")

}
