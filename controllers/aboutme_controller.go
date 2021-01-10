package controllers

type AboutMeController struct {
	BaseController
}

func (this *AboutMeController) Get() {
	this.Data["wechat"] = "微信：xxx"
	this.Data["qq"] = "qq:xxxxx"
	this.Data["tel"] = "tel: xxx"
	this.TplName = "aboutme.html"
}
