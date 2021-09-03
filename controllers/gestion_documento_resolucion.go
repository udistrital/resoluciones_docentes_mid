package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
)

//GestionDocumentoResolucionController operations for Preliquidacion
type GestionDocumentoResolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionDocumentoResolucionController) URLMapping() {

}

// GestionPrevinculacionesController ...
// @Title ListarDocentesCargaHoraria
// @Description create GetContenidoResolucion
// @Param id_resolucion query string false "Numero de resolución a consultar"
// @Param id_facultad query string false "dependencia que ordena el gasto"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /get_contenido_resolucion [get]
func (c *GestionDocumentoResolucionController) GetContenidoResolucion() {
	id_resolucion := c.GetString("id_resolucion")
	id_facultad := c.GetString("id_facultad")
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDocumentoResolucionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	idRes, err1 := strconv.Atoi(id_resolucion)
	idFac, err2 := strconv.Atoi(id_facultad)
	if (err1 != nil) || (err2 != nil) || (idRes == 0) || (idFac == 0) {
		panic(map[string]interface{}{"funcion": "GetContenidoResolucion", "err": "Error en los parametros de ingreso", "status": "400"})
	}
	if contenidoResolucion, err := helpers.GetContenidoResolucion(id_resolucion, id_facultad); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": contenidoResolucion}
	} else {
		panic(err)
	}

	c.ServeJSON()

}
