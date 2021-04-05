package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

// ExpedirResolucionController operations for ExpedirResolucion
type ExpedirResolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ExpedirResolucionController) URLMapping() {
	c.Mapping("Expedir", c.Expedir)
	c.Mapping("ValidarDatosExpedicion", c.ValidarDatosExpedicion)
	c.Mapping("ExpedirModificacion", c.ExpedirModificacion)

}

// Expedir ...
// @Title Expedir
// @Description create Expedir
// @Param	body		body 	[]models.ExpedicionResolucion	true		"body for Expedicion Resolucion content"
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedir [post]
func (c *ExpedirResolucionController) Expedir() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ExpedirResolucionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()


	var m models.ExpedicionResolucion
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		if err := helpers.Expedir(m); err == nil{
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		}else{
			panic(err)
		}
	} else { //If 13 - Unmarshal
		panic(map[string]interface{}{"funcion": "Expedir", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// ExpedirResolucionController ...
// @Title ValidarDatosExpedicion
// @Description create ValidarDatosExpedicion
// @Param	body		body 	[]models.ExpedicionResolucion	true		"body for Validar Datos Expedición content"
// @Success 201 {int}
// @Failure 403 body is empty
// @router /validar_datos_expedicion [post]
func (c *ExpedirResolucionController) ValidarDatosExpedicion() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ExpedirResolucionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var m models.ExpedicionResolucion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		if err := helpers.ValidarDatosExpedicion(m); err == nil{
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		}else{
			panic(err)
		}
	}else{
		panic(map[string]interface{}{"funcion": "ValidarDatosExpedicion", "err": err.Error(), "status": "400"})
	}	
	c.ServeJSON()
}

// ExpedirModificacion ...
// @Title ExpedirModificacion
// @Description create ExpedirModificacion
// @Param	body		body 	[]models.ExpedicionResolucion	true		"body for Validar Datos Expedición content"
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedirModificacion [post]
func (c *ExpedirResolucionController) ExpedirModificacion() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ExpedirResolucionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var m models.ExpedicionResolucion
	// If 13 - Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		if err := helpers.ExpedirModificacion(m); err == nil{
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		}else{
			panic(err)
		}
	} else { //If 13 - Unmarshal
		panic(map[string]interface{}{"funcion": "ExpedirModificacion", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// Cancelar ...
// @Title Cancelar
// @Description create Cancelar
// @Success 201 {int} models.ExpedicionCancelacion
// @Failure 403 body is empty
// @router /cancelar [post]
func (c *ExpedirResolucionController) Cancelar() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ExpedirResolucionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var m models.ExpedicionCancelacion
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		if err := helpers.Cancelar(m); err == nil{
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		}else{
			panic(err)
		}
	} else { //If 13 - Unmarshal
		panic(map[string]interface{}{"funcion": "Cancelar", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}