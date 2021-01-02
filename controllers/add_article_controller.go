package controllers

import (
	"beego_blogweb/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	article := models.Article{0, title, tags, short, content, "renwoxing", time.Now().Unix()}
	_, err := models.AddArticle(article)

	var response map[string]interface{}
	if err == nil {
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}
	this.Data["json"] = response
	this.ServeJSON()
}
