package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	Basecontroller
}

func (c *LoginController) UserError() {
	c.TplName = c.GetTemplatetype() + "/main/usererror.html"
}

func (c *LoginController) AdminError() {
	c.TplName = c.GetTemplatetype() + "/main/adminerror.html"
}
func (c *LoginController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		user, err := models.FindUser(username)
		if username == beego.AppConfig.String("uname") && password == beego.AppConfig.String("pwd") {
			c.SetSession("admin", beego.AppConfig.String("uname"))
			//c.Redirect("/admin/index", 301)
			c.ReportError(SUCCESS_ADMIN)
		} else if err != nil {
			c.ReportError(NO_USERNAME)

		} else {
			if username != "admin" && password == user.Password {
				c.SetSession("user", username)
				c.ReportError(SUCCESS)
			} else {
				c.ReportError(INVALID_PASSWORD)
			}
		}
	}
	c.Data["IsLogin"] = true
	c.Data["IsModify"] = true
	c.TplName = c.GetTemplatetype() + "/main/login.html"
}

func (c *LoginController) Register() {
	if c.Ctx.Request.Method == "POST" {
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		pass := c.Input().Get("pass")
		email := c.Input().Get("email")
		_, err := models.FindUser(username)
		if err == nil {
			c.ReportError(REGISTER_USER)
		} else {
			models.RegisterUser(username, password, pass, email)
			c.ReportError(SUCCESS)
		}

	}
	c.Data["Reg"] = true
	c.Data["IsModify"] = true
	c.TplName = c.GetTemplatetype() + "/main/register.html"
}
