package routers

import (
	"beego_blogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/article/add", &controllers.AddArticleController{})
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	beego.Router("/tags", &controllers.TagsController{})
	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/aboutme", &controllers.AboutMeController{})

}
