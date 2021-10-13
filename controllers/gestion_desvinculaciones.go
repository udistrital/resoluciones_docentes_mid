package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

// PreliquidacionController operations for Preliquidacion
type GestionDesvinculacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionDesvinculacionesController) URLMapping() {
	c.Mapping("ListarDocentesDesvinculados", c.ListarDocentesDesvinculados)
	c.Mapping("ListarDocentesCancelados", c.ListarDocentesCancelados)
	c.Mapping("AnularModificaciones", c.AnularModificaciones)
	c.Mapping("AnularAdicionDocente", c.AnularAdicionDocente)
	c.Mapping("ConsultarCategoria", c.ConsultarCategoria)
	c.Mapping("ValidarSaldoCDP", c.ValidarSaldoCDP)
	c.Mapping("AdicionarHoras", c.AdicionarHoras)
	c.Mapping("ActualizarVinculacionesCancelacion", c.ActualizarVinculacionesCancelacion)
}

// GestionDesvinculacionesController ...
// @Title ListarDocentesDesvinculados
// @Description create ListarDocentesDesvinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 200 {int} []models.VinculacionDocente
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /docentes_desvinculados [get]
func (c *GestionDesvinculacionesController) ListarDocentesDesvinculados() {
	id_resolucion := c.GetString("id_resolucion")
	query := "?limit=-1&query=ResolucionVinculacionDocenteId.Id:" + id_resolucion

	_, err1 := strconv.Atoi(id_resolucion)

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	if err1 != nil {
		panic(map[string]interface{}{"funcion": "ListarDocentesDesvinculados", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if lista_docentes, err := helpers.ListarDocentesDesvinculados(query); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": lista_docentes}
	} else {
		panic(err)
	}
	c.ServeJSON()

}

// GestionCanceladosController ...
// @Title ListarDocentesCancelados
// @Description create ListarDocentesCancelados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 200 {int} []models.VinculacionDocente
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /docentes_cancelados [get]
func (c *GestionDesvinculacionesController) ListarDocentesCancelados() {
	id_resolucion := c.GetString("id_resolucion")

	_, err1 := strconv.Atoi(id_resolucion)

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	if err1 != nil {
		panic(map[string]interface{}{"funcion": "ListarDocentesCancelados", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if lista_docentes, err := helpers.ListarDocentesCancelados(id_resolucion); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": lista_docentes}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// AnularModificaciones ...
// @Title AnularModificaciones
// @Description Se usa para cuando se anulan resoluciones modificatorias completas
// @Param	body		body 	[]models.VinculacionDocente	true		"body for AnularModificaciones content"
// @Success 201 {string}
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /anular_modificaciones [post]
func (c *GestionDesvinculacionesController) AnularModificaciones() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var v []models.VinculacionDocente

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := helpers.AnularModificaciones(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "AnularModificaciones", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// AnularAdicionDocente ...
// @Title AnularAdicionDocente
// @Param	body		body 	models.Objeto_Desvinculacion	true		"body for Anular Adiciones content"
// @Description Se usa para adiciones, reducciones y cancelaciones
// @Success 201 {string}
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /anular_adicion [post]
func (c *GestionDesvinculacionesController) AnularAdicionDocente() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	fmt.Println("anular adicion")
	var v models.Objeto_Desvinculacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := helpers.AnularAdicionDocente(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "AnularAdicionDocente", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// ConsultarCategoria ...
// @Title ConsultarCategoria
// @Param	body		body 	models.VinculacionDocente	true		"body for Consultar Categoria content"
// @Description Consulta el servicio de categoría en académica para verificar si el docente tiene el semáforo completo
// @Success 201 {string}
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /consultar_categoria [post]
func (c *GestionDesvinculacionesController) ConsultarCategoria() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var v models.VinculacionDocente
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if respuesta, err := helpers.ConsultarCategoria(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": respuesta}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarCategoria", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// ValidarSaldoCDP ...
// @Title ValidarSaldoCDP
// @Description Se usa para validar el saldo de la disponibilidad con el valor del contrato de las adiciones
// @Param	body		body 	models.Objeto_Desvinculacion	true		"body for Objeto Desvinculacion content"
// @Success 201 {string}
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /validar_saldo_cdp [post]
func (c *GestionDesvinculacionesController) ValidarSaldoCDP() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var validacion models.Objeto_Desvinculacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &validacion); err == nil {
		if respuesta, err := helpers.ValidarSaldoCDP(validacion); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": respuesta}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "ValidarSaldoCDP", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// AdicionarHoras ...
// @Title AdicionarHoras
// @Description Se usa tanto para adiciones como para reducciones de horas y semanas
// @Param	body		body 	models.Objeto_Desvinculacion	true		"body for Objeto Desvinculacion content"
// @Success 201 {string}
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /adicionar_horas [post]
func (c *GestionDesvinculacionesController) AdicionarHoras() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var v models.Objeto_Desvinculacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if respuesta, err := helpers.AdicionarHoras(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": respuesta}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "AdicionarHoras", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()

}

// ActualizarVinculacionesCancelacion ...
// @Title ActualizarVinculacionesCancelacion
// @Description create ActualizarVinculacionesCancelacion
// @Param	body		body 	models.Objeto_Desvinculacion	true		"body for Objeto Desvinculacion content"
// @Success 201 {string}
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /actualizar_vinculaciones_cancelacion [post]
func (c *GestionDesvinculacionesController) ActualizarVinculacionesCancelacion() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "GestionDesvinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var v models.Objeto_Desvinculacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if respuesta, err := helpers.ActualizarVinculacionesCancelacion(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": respuesta}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "ActualizarVinculacionesCancelacion", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()

}
