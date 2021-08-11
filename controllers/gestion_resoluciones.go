package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

//GestionResolucionesController operations for Preliquidacion
type GestionResolucionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionResolucionesController) URLMapping() {
	c.Mapping("InsertarResolucionCompleta", c.InsertarResolucionCompleta)

}

// GestionResolucionesController ...
// @Title getResolucionesAprobadas
// @Description create  getResolucionesAprobadas
// @Param limit query int false "Limit the size of result set. Must be an integer"
// @Param offset query int false "Start position of result set. Must be an integer"
// @Param query query string false "Filter. e.g. col1:v1,col2:v2 ..."
// @Success 201 {object} []models.ResolucionVinculacion
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /get_resoluciones_aprobadas [get]
func (c *GestionResolucionesController) GetResolucionesAprobadas() {
	query := c.GetString("query")
	limit, err1 := c.GetInt("limit")
	offset, err2 := c.GetInt("offset")
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionResolucionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	var resolucion_vinculacion_aprobada []models.ResolucionVinculacion
	var err3 map[string]interface{}
	if (limit == 0) || (len(query) == 0) || (offset == 0) || (err1 != nil) || (err2 != nil) {
		panic(map[string]interface{}{"funcion": "GetResolucionesAprobadas", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if resolucion_vinculacion_aprobada, err3 = helpers.GetResolucionesAprobadas(query, limit, offset); err3 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": resolucion_vinculacion_aprobada}
	} else {
		panic(err3)
	}

	c.ServeJSON()
}

// GestionResolucionesController ...
// @Title getResolucionesInscritas
// @Description create  getResolucionesInscritas
// @Param limit query int false "Limit the size of result set. Must be an integer"
// @Param offset query int false "Start position of result set. Must be an integer"
// @Param query query string false "Filter. e.g. col1:v1,col2:v2 ..."
// @Success 201 {object} []models.ResolucionVinculacion
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /get_resoluciones_inscritas [get]
func (c *GestionResolucionesController) GetResolucionesInscritas() {
	var query []string
	query = c.GetStrings("query")
	limit, err1 := c.GetInt("limit")
	offset, err2 := c.GetInt("offset")
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "CertificacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	var resolucion_vinculacion []models.ResolucionVinculacion
	if (limit == 0) || (len(query) == 0) || (offset == 0) || (err1 != nil) || (err2 != nil) {
		panic(map[string]interface{}{"funcion": "GetResolucionesInscritas", "err": "Error en los parametros de ingreso", "status": "400"})
	}
	var err3 map[string]interface{}
	if resolucion_vinculacion, err3 = helpers.GetResolucionesInscritas(query, limit, offset); err3 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": resolucion_vinculacion}
	} else {
		panic(err3)
	}
	c.ServeJSON()
}

// InsertarResolucionCompleta ...
// @Title InsertarResolucionCompleta
// @Description create InsertarResolucionCompleta
// @Success 201 {int} models.Resolucion
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /insertar_resolucion_completa [post]
func (c *GestionResolucionesController) InsertarResolucionCompleta() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "CertificacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	var v models.ObjetoResolucion
	if err1 := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err1 == nil {
		if id_resolucion_creada, control, err2 := helpers.InsertarResolucionCompleta(v); err2 == nil || control == false {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": id_resolucion_creada}
		} else {
			panic(err2)
		}
	} else {
		panic(err1)
	}

	c.ServeJSON()
}
