// @APIVersion 1.0.0
// @Title Resoluciones Docentes API MID
// @Description Middleware para la gestión de la información de resoluciones
// @Contact computo@udistrital.edu.co
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

		beego.NSNamespace("/gestion_previnculacion",
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
