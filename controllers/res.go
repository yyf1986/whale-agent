package controllers

import (
	"whale-agent/models"
	"github.com/astaxie/beego"
)

type ResController struct {
	beego.Controller
}

func (d *ResController) Get() {
	d.Data["json"] = models.GetRes()
	d.ServeJSON()
}