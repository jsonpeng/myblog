package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	Basecontroller
}

func (c *TopicController) View() {
	c.TplName = c.GetTemplatetype() + "/main/topicview.html"

	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	tid := c.Ctx.Input.Param("0")
	reply, err := models.GetAllReply(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Reply"] = reply
	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
}
func (c *TopicController) AddReply() {
	tid := c.Input().Get("tid")
	err := models.AddReply(tid, c.Input().Get("nickname"), c.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view/"+tid, 302)
}
