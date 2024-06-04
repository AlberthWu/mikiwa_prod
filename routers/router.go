// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	master "mikiwa_prod/controllers/master"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//company
	beego.Router("/v1/company", &master.CompanyController{}, "get:GetAll;post:Post")

}
