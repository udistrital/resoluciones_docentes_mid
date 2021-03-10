package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/udistrital/resoluciones_docentes_mid/helpers"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

// PreliquidacionController operations for Preliquidacion
type GestionPrevinculacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionPrevinculacionesController) URLMapping() {
	c.Mapping("InsertarPrevinculaciones", c.InsertarPrevinculaciones)
	c.Mapping("CalcularTotalDeSalarios", c.CalcularTotalSalarios)
	c.Mapping("ListarDocentesCargaHoraria", c.ListarDocentesCargaHoraria)
	c.Mapping("GetCdpRpDocente", c.GetCdpRpDocente)
}

// Calcular_total_de_salarios_seleccionados ...
// @Title Calcular_total_de_salarios_seleccionados
// @Description createCalcular_total_de_salarios_seleccionados
// @Success 201 {int} int
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos_seleccionados [post]
func (c *GestionPrevinculacionesController) Calcular_total_de_salarios_seleccionados() {

	var v []models.VinculacionDocente
	var total int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v, err2 := helpers.CalcularSalarioPrecontratacion(v); err2 == nil {
			total = int(helpers.CalcularTotalSalario(v))
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
		beego.Error(err)
		c.Abort("400")
	}

	c.Data["json"] = total
	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos [post]
func (c *GestionPrevinculacionesController) CalcularTotalSalarios() {
	var v []models.VinculacionDocente
	var total int

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		total = helpers.CalculoSalarios(v)
	} else {
		beego.Error(err)
		c.Abort("400")
	}

	c.Data["json"] = total
	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/insertar_previnculaciones [post]
func (c *GestionPrevinculacionesController) InsertarPrevinculaciones() {

	var v []models.VinculacionDocente
	var idRespuesta int

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		idRespuesta = helpers.InsertarPrevinculaciones(v)
	} else {
		beego.Error("Error al hacer unmarshal", err)
		c.Abort("403")
	}
	c.Data["json"] = idRespuesta
	c.ServeJSON()
}

// GestionPrevinculacionesController ...
// @Title ListarDocentesCargaHoraria
// @Description create ListarDocentesCargaHoraria
// @Param vigencia query string false "año a consultar"
// @Param periodo query string false "periodo a listar"
// @Param tipo_vinculacion query string false "vinculacion del docente"
// @Param facultad query string false "facultad"
// @Param nivel_academico query string false "nivel_academico"
// @Success 201 {object} models.Docentes_x_Carga
// @Failure 403 body is empty
// @router /Precontratacion/docentes_x_carga_horaria [get]
func (c *GestionPrevinculacionesController) ListarDocentesCargaHoraria() {

	vigencia := c.GetString("vigencia")
	periodo := c.GetString("periodo")
	tipoVinculacion := c.GetString("tipo_vinculacion")
	facultad := c.GetString("facultad")
	nivelAcademico := c.GetString("nivel_academico")

	respuesta := helpers.ListarDocentesCargaHoraria(vigencia, periodo, tipoVinculacion, facultad, nivelAcademico)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = respuesta.CargasLectivas.CargaLectiva
	c.ServeJSON()

}

const (
	tipoVinculacion = iota + 1
	tipoCancelacion
	tipoAdicion
	tipoReduccion
)

// GestionPrevinculacionesController ...
// @Title ListarDocentesPrevinculadosAll
// @Description create ListarDocentesPrevinculadosAll
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados_all [get]
func (c *GestionPrevinculacionesController) ListarDocentesPrevinculadosAll() {
	idResolucion := c.GetString("id_resolucion")

	v := helpers.ListarDocentesPrevinculadosAll(idResolucion, tipoVinculacion, tipoCancelacion, tipoAdicion, tipoReduccion)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()

}

//ESTA FUNCIÓN LISTA LOS DOCENTES PREVINCULADOS EN TRUE

// GestionPrevinculacionesController ...
// @Title ListarDocentesPrevinculados
// @Description create ListarDocentesPrevinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados [get]
func (c *GestionPrevinculacionesController) ListarDocentesPrevinculados() {
	idResolucion := c.GetString("id_resolucion")

	v := helpers.ListarDocentesPrevinculados(idResolucion, tipoVinculacion)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()

}

// GestionPrevinculacionesController ...
// @Title GetCdpRpDocente
// @Description Get RPs de la vinculacion docente
// @Param num_vinculacion query string true "Número de la vinculación del docente"
// @Param vigencia query string true "Vigencia de la vinculación del docente"
// @Success 201 {object}  models.RpDocente
// @Failure 403 :num_vinculacion is empty
// @Failure 403 :vigencia is empty
// @router /rp_docente/:num_vinculacion/:vigencia/:identificacion [get]
func (c *GestionPrevinculacionesController) GetCdpRpDocente() {
	num_vinculacion := c.Ctx.Input.Param(":num_vinculacion")
	vigencia := c.Ctx.Input.Param(":vigencia")
	identificacion := c.Ctx.Input.Param(":identificacion")

	rpdocente := helpers.GetCdpRpDocente(identificacion, num_vinculacion, vigencia)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = rpdocente
	c.ServeJSON()

}
