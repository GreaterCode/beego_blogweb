package controllers

import "beego_blogweb/models"

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	this.Data["Tags"] = models.HandleTagsListData(tags)
	this.TplName = "tags.html"
}
