package controllers

import "fmt"

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	fmt.Println("IsLogin", this.IsLogin)
	this.TplName = "home.html"
}
