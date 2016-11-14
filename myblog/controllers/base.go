package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

const (
	SUCCESS = 0x00 //成功
	FAILED  = 0x01 //失败

	//LOCAL
	NO_USERNAME           = 0x12 //用户名不存在
	INVALID_USERNAME      = 0x13 //用户名非法
	INVALID_PASSWORD      = 0x14 //密码不正确
	REGISTER_USER         = 0x16 //用户已经注册
	LONGINNED_USER        = 0x17 //用户已登录
	ERROR_USERINFORMATION = 0X18 //清除了Session

	NO_LOGIN = 0x25 //用户未登录

)

type Basecontroller struct {
	beego.Controller
	Templatetype string
}

func (c *Basecontroller) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "blog"
	}
	return templatetype
}
func (c *Basecontroller) ReportError(errno int) {
	errno2str := strconv.Itoa(errno)
	c.Data["json"] = &map[string]interface{}{"status": true, "errno": errno2str}
	c.ServeJSON()
}
