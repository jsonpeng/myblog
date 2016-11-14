package controllers

type LoginController struct {
	Basecontroller
}

func (c *LoginController) Login() {

	c.TplName = c.GetTemplatetype() + "/main/login.html"
}

func (c *LoginController) Register() {
	c.TplName = c.GetTemplatetype() + "/main/register.html"
}
