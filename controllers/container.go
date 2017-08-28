package controllers

import (
	"github.com/astaxie/beego"
	"whale-agent/models"
)

type ContainerController struct {
	beego.Controller
}

func (d *ContainerController) Index() {
	d.Data["json"] = map[string]interface{}{"status": 200}
	d.ServeJSON()
}
func (d *ContainerController) Create() {
	container_name := d.GetString("container_name")
	image_name := d.GetString("image_name")
	env := d.GetString("env")
	if container_name != "" && image_name != "" {
		err := models.CreateContainer(container_name, image_name, env)
		if err != nil {
			d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": err.Error()}
		} else {
			d.Data["json"] = map[string]interface{}{"status": 200}
		}
	} else {
		d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": "Miss agrs"}
	}
	d.ServeJSON()
}

func (d *ContainerController) Start() {
	container_id := d.GetString("container_id")
	if container_id != "" {
		err := models.StartContainer(container_id)
		if err != nil {
			d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": err.Error()}
		} else {
			d.Data["json"] = map[string]interface{}{"status": 200}
		}
	} else {
		d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": "Miss agrs"}
	}
	d.ServeJSON()
}

func (d *ContainerController) Stop() {
	container_id := d.GetString("container_id")
	if container_id != "" {
		err := models.StopContainer(container_id)
		if err != nil {
			d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": err.Error()}
		} else {
			d.Data["json"] = map[string]interface{}{"status": 200}
		}
	} else {
		d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": "Miss agrs"}
	}
	d.ServeJSON()
}

func (d *ContainerController) Del() {
	container_id := d.GetString("container_id")
	if container_id != "" {
		err := models.DelContainer(container_id)
		if err != nil {
			d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": err.Error()}
		} else {
			d.Data["json"] = map[string]interface{}{"status": 200}
		}
	} else {
		d.Data["json"] = map[string]interface{}{"status": 666, "errinfo": "Miss agrs"}
	}
	d.ServeJSON()
}

func (d *ContainerController) GetAll() {
	var containers []models.Container
	containers = models.GetContainer()
	d.Data["json"] = containers
	d.ServeJSON()
}
