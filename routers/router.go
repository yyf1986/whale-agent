// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"whale-agent/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.ContainerController{}, "Get:Index")
	//beego.Router("/v1/container/create", &controllers.ContainerController{}, "Get:Create")
	//beego.Router("/v1/container/start", &controllers.ContainerController{}, "Get:Start")
	//beego.Router("/v1/container/stop", &controllers.ContainerController{}, "Get:Stop")
	//beego.Router("/v1/container/del", &controllers.ContainerController{}, "Get:Del")
	//beego.Router("/v1/container/getall", &controllers.ContainerController{}, "Get:GetAll")
	beego.Router("/v1/res/host", &controllers.ResController{})
	beego.Router("/v1/res/getport", &controllers.ResController{}, "Get:GetPort")
}
