package controllers

import (
	"encoding/json"
	"fmt"

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

	//c.Mapping("ActualizarVinculaciones", c.ActualizarVinculaciones)
	c.Mapping("AdicionarHoras", c.AdicionarHoras)

}

// GestionDesvinculacionesController ...
// @Title ListarDocentesDesvinculados
// @Description create ListarDocentesDesvinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_desvinculados [get]
func (c *GestionDesvinculacionesController) ListarDocentesDesvinculados() {
	fmt.Println("docentes desvinculados")
	id_resolucion := c.GetString("id_resolucion")
	query := "?limit=-1&query=IdResolucion.Id:" + id_resolucion

	lista_docentes := helpers.ListarDocentesDesvinculados(query)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = lista_docentes
	c.ServeJSON()

}

// GestionCanceladosController ...
// @Title ListarDocentesCancelados
// @Description create ListarDocentesCancelados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_cancelados [get]
func (c *GestionDesvinculacionesController) ListarDocentesCancelados() {
	id_resolucion := c.GetString("id_resolucion")

	lista_docentes := helpers.ListarDocentesCancelados(id_resolucion)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = lista_docentes
	c.ServeJSON()
}

// AnularModificaciones ...
// @Title AnularModificaciones
// @Description create AnularModificaciones
// @Param	body		body 	[]models.VinculacionDocente	true		"body for AnularModificaciones content"
// @Success 201 {string}
// @Failure 403 body is empty
// @router /anular_modificaciones [post]
// Se usa para cuando se anulan resoluciones modificatorias completas
func (c *GestionDesvinculacionesController) AnularModificaciones() {
	var v []models.VinculacionDocente
	var respuesta_total string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta_total = helpers.AnularModificaciones(v)
	} else {
		respuesta_total = "error"
	}

	c.Data["json"] = respuesta_total
	c.ServeJSON()
}

// AnularAdicionDocente ...
// @Title AnularAdicionDocente
// @Description create AnularAdicionDocente
// @Success 201 {string}
// @Failure 403 body is empty
// @router /anular_adicion [post]
// Se usa para adiciones, reducciones y cancelaciones
func (c *GestionDesvinculacionesController) AnularAdicionDocente() {
	fmt.Println("anular adicion")
	var v models.Objeto_Desvinculacion
	var respuesta_total string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta_total = helpers.AnularAdicionDocente(v)
	} else {
		respuesta_total = "error"
	}

	c.Data["json"] = respuesta_total
	c.ServeJSON()
}

// ConsultarCategoria ...
// @Title ConsultarCategoria
// @Description create ConsultarCategoria
// @Success 201 {string}
// @Failure 403 body is empty
// @router /consultar_categoria [post]
// Consulta el servicio de categoría en académica para verificar si el docente tiene el semáforo completo
func (c *GestionDesvinculacionesController) ConsultarCategoria() {
	var v models.VinculacionDocente
	var respuesta string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta = helpers.ConsultarCategoria(v)
	} else {
		beego.Error(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	c.Data["json"] = respuesta
	c.ServeJSON()
}

// ValidarSaldoCDP ...
// @Title ValidarSaldoCDP
// @Description create ValidarSaldoCDP
// @Success 201 {string}
// @Failure 403 body is empty
// @router /validar_saldo_cdp [post]
// Se usa para validar el saldo de la disponibilidad con el valor del contrato de las adiciones
func (c *GestionDesvinculacionesController) ValidarSaldoCDP() {
	var validacion models.Objeto_Desvinculacion
	var respuesta string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &validacion); err == nil {
		respuesta = helpers.ValidarSaldoCDP(validacion)
	} else {
		beego.Error(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}
	c.Data["json"] = respuesta
	c.ServeJSON()
}

// AdicionarHoras ...
// @Title AdicionarHoras
// @Description create AdicionarHoras
// @Success 201 {string}
// @Failure 403 body is empty
// @router /adicionar_horas [post]
// Se usa tanto para adiciones como para reducciones de horas y semanas
func (c *GestionDesvinculacionesController) AdicionarHoras() {

	var v models.Objeto_Desvinculacion
	var respuesta string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta = helpers.AdicionarHoras(v)
	} else {
		beego.Error(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	c.Data["json"] = respuesta

	c.ServeJSON()

}

// ActualizarVinculacionesCancelacion ...
// @Title ActualizarVinculacionesCancelacion
// @Description create ActualizarVinculacionesCancelacion
// @Success 201 {string}
// @Failure 403 body is empty
// @router /actualizar_vinculaciones_cancelacion [post]
func (c *GestionDesvinculacionesController) ActualizarVinculacionesCancelacion() {

	var v models.Objeto_Desvinculacion
	var respuesta interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta = helpers.ActualizarVinculacionesCancelacion(v)
	} else {
		beego.Error(err)
		c.Abort("400")
	}

	c.Data["json"] = respuesta

	c.ServeJSON()

}
