package controllers

import (
	"beego_blogweb/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}

	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	fmt.Println("IsLogin", this.IsLogin)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	this.Data["HasFooter"] = true
	this.TplName = "home.html"
}
