package controllers

import (
	"goweb/beeblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.TplName = "topic.html"
	this.Data["IsTopic"] = true
	this.Data["isLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
	this.Data["IsTopic"] = true
	this.Data["isLogin"] = checkAccount(this.Ctx)
	// 获取category
	categorys, err := models.GetAllCategory()
	if err != nil {
		this.Redirect("/topic", 302)
		return
	}
	this.Data["Categorys"] = categorys

}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	this.Data["IsTopic"] = true
	isLogin := checkAccount(this.Ctx)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/home", 302)
	}
	this.Data["Topic"] = topic
}

func (this *TopicController) Delete() {
	isLogin := checkAccount(this.Ctx)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Ctx.Input.Params()["0"]
	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.GetString("title")
	content := this.GetString("content")
	tid := this.GetString("tid")
	categoryId := this.GetString("categoryId") // 获取分类Id

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content, categoryId)
	} else {
		err = models.UpdateTopic(title, content, tid)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)

}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	this.Data["IsTopic"] = true
	this.Data["isLogin"] = checkAccount(this.Ctx)

	// tid := this.GetString("tid")
	tid := this.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/home", 302)
	}
	this.Data["Topic"] = topic
}
