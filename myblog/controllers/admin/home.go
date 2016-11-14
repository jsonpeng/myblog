package controllers

import (
	"myblog/controllers"
	"myblog/models"
	"time"

	"github.com/astaxie/beego"
)

type AdminController struct {
	controllers.Basecontroller
}

func (c *AdminController) Home() {
	c.Data["IsHome"] = true
	c.Data["IsAdmin"] = true
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["AllTopics"] = topics

	c.TplName = c.GetTemplatetype() + "/admin/home.html"
}
func (c *AdminController) Topic() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["IsTopic"] = true
		tid := c.Input().Get("tid")
		topic, _ := models.GetTopic(tid)

		c.Data["Topic"] = topic
		c.Data["Tid"] = tid

		var err error
		c.Data["AllTopics"], err = models.GetAllTopics(false)
		if err != nil {
			beego.Error(err)
		}
		action := c.Input().Get("action")
		switch action {
		case "del":
			id := c.Input().Get("id")
			err := models.DelTopics(id)
			if err != nil {
				beego.Error(err)
			}
			//c.ReportError(0x00)
			c.Redirect("/admin/topic", 301)
		}
		c.TplName = c.GetTemplatetype() + "/admin/topic.html"
	}
}

func (c *AdminController) AddTopic() {
	if c.Ctx.Request.Method == "POST" {
		title := c.Input().Get("title")
		category := c.Input().Get("category")
		desc := c.Input().Get("desc")
		content := c.Input().Get("content")
		var err error
		err = models.AddTopic(title, category, desc, content)
		if err != nil {
			beego.Error(err)
		}
		c.Ctx.Output.Body([]byte("<script>alert('添加成功!');parent.window.location.href='/admin/topic';</script>"))
		//c.Redirect("/admin/topic", 301)
	}
	c.Data["IsModify"] = true
	c.TplName = c.GetTemplatetype() + "/admin/addtopic.html"
}
func (c *AdminController) Category() {

	action := c.Input().Get("action")

	switch action {
	case "add":
		name := c.Input().Get("name")
		err := models.AddCategory(name, time.Now())
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/category", 301)
		return
	}

	c.TplName = c.GetTemplatetype() + "/admin/category.html"
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	c.Data["IsCate"] = true
	if err != nil {
		beego.Error(err)
	}
}
func (c *AdminController) Modify() {
	c.TplName = c.GetTemplatetype() + "/admin/modifytopic.html"
	if c.Ctx.Request.Method == "POST" {
		tid := c.Input().Get("tid")
		title := c.Input().Get("title")
		category := c.Input().Get("category")
		desc := c.Input().Get("desc")
		content := c.Input().Get("content")
		models.ModifyTopic(tid, title, category, desc, content)
		c.Ctx.Output.Body([]byte("<script>alert('修改成功！');window.location.href='/admin/topic';</script>"))

	}
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}
func (c *AdminController) ManageLiuyan() {
	action := c.Input().Get("action")
	switch action {
	case "del":
		id := c.Input().Get("id")
		err := models.DelLiuyan(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/liuyan", 301)
	}
	c.Data["IsLiu"] = true
	c.TplName = c.GetTemplatetype() + "/admin/liuyan.html"
	var err error
	c.Data["AllLiuyan"], err = models.GetAllLiuyan(false)
	if err != nil {
		beego.Error(err)
	}
}
