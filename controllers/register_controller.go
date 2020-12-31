package controllers

import (
	"beego_blogweb/models"
	"beego_blogweb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Printf("id", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Username has already exists"}
		this.ServeJSON()
		return
	}

	password = utils.MD5(password)
	fmt.Println("md5后的password:", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Register failed"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Register successfully"}
	}
	this.ServeJSON()
}
