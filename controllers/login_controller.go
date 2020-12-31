package controllers

import (
	"beego_blogweb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username", username, "password:", password)

	id := models.QueryUserWithUsername(username)
	fmt.Println("登录用户ID：", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Login is sueccessful!!"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Login is failed!!!"}
	}

}
