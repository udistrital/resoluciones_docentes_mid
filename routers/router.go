// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/resoluciones_docentes_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/gestion_documento_resolucion",
			beego.NSInclude(
				&controllers.GestionDocumentoResolucionController{},
			),
		),

		beego.NSNamespace("/expedir_resolucion",
			beego.NSInclude(
				&controllers.ExpedirResolucionController{},
			),
		),

		beego.NSNamespace("/gestion_desvinculaciones",
			beego.NSInclude(
				&controllers.GestionDesvinculacionesController{},
			),
		),

		beego.NSNamespace("/gestion_previnculaciones",
			beego.NSInclude(
				&controllers.GestionPrevinculacionesController{},
			),
		),

		beego.NSNamespace("/gestion_resoluciones",
			beego.NSInclude(
				&controllers.GestionResolucionesController{},
			),
		),
	/*  beego.NSNamespace("/user",
	    beego.NSInclude(
	        &controllers.UserController{},
	    ),
	), */
	)
	beego.AddNamespace(ns)
}
