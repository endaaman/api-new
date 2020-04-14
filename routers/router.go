// @APIVersion 1.0.0
// @Title API for endaaman.me
// @Description api.endaaman.me
// @Contact buhibuhidog@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/endaaman/api.endaaman.me/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.AddNamespace(beego.NewNamespace("/v1",
		// beego.NSRouter("/articles", &controllers.ArticleController{}),
		beego.NSNamespace("/articles",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
		beego.NSNamespace("/categories",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/sessions",
			beego.NSInclude(
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace("/files",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
		beego.NSNamespace("/misc",
			beego.NSInclude(
				&controllers.MiscController{},
			),
		),
	))

	// beego.Router("/static", &controllers.StaticController{})
	beego.AddNamespace(beego.NewNamespace("/static",
		beego.NSInclude(
			&controllers.StaticController{},
		),
	))
}
