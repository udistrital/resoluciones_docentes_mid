package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["resoluciones_docentes_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["resoluciones_docentes_mid/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["resoluciones_docentes_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["resoluciones_docentes_mid/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method:           "Expedir",
			Router:           `/expedir`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method:           "ValidarDatosExpedicion",
			Router:           `/validar_datos_expedicion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method:           "ExpedirModificacion",
			Router:           `/expedirModificacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method:           "Cancelar",
			Router:           `/cancelar`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	

}
