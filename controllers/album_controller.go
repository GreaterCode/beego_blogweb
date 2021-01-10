package controllers

import (
	"beego_blogweb/models"
	"github.com/astaxie/beego/logs"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		logs.Error(err)
	}

	this.Data["Album"] = albums
	this.TplName = "album.html"
}
