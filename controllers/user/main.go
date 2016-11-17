package controllers

import (
	"myblog/controllers"
	"myblog/models"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	controllers.Basecontroller
}

func (c *MainController) Index() {
	//	user := c.Ctx.Input.Session("user")
	//	//c.GetSession("user")
	//	if user == nil {
	//		//c.Ctx.Output.Body([]byte("<script>alert('请先登录!');parent.window.location.href='/login';</script>"))
	//		c.Redirect("/error/login", 301)
	//	}
	c.Data["IsIndex"] = true
	topics, err := models.GetAllTopics(c.Input().Get("cate"), "", true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["AllTopics"] = topics
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
	c.TplName = c.GetTemplatetype() + "/user/index.html"

}

func (c *MainController) Category() {
	c.Data["IsCategory"] = true
	c.TplName = c.GetTemplatetype() + "/user/category.html"
}

func (c *MainController) Shouji() {
	c.Data["IsShouji"] = true
	c.TplName = c.GetTemplatetype() + "/user/shouji.html"
}

func (c *MainController) Liuyan() {
	if c.Ctx.Request.Method == "POST" {
		allcontent := c.Input().Get("allcontent")
		var err error
		err = models.AddLiuyan(allcontent, time.Now())
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/liuyan", 301)
	}
	c.Data["IsLiuyan"] = true
	liuyan, err := models.GetAllLiuyan(true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Liuyan"] = liuyan
	c.TplName = c.GetTemplatetype() + "/user/liuyan.html"
}

func (c *MainController) About() {
	c.Data["IsAbout"] = true
	c.TplName = c.GetTemplatetype() + "/user/about.html"
}
func (c *MainController) View() {
	c.TplName = "/blog/main/topicview.html"
	//	cc := c.Ctx.Input.Param("0")
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
}
