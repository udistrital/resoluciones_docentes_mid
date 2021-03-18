package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

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
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPrevinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	var v []models.VinculacionDocente
	var total int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v, err2 := helpers.CalcularSalarioPrecontratacion(v); err2 == nil {
			total = int(helpers.CalcularTotalSalario(v))
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": total}
		} else {
			panic(err2)
		}
	} else {
		panic(map[string]interface{}{"funcion": "Calcular_total_de_salarios_seleccionados", "err": err.Error(), "status": "400"})
	}
	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos [post]
func (c *GestionPrevinculacionesController) CalcularTotalSalarios() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPrevinculacionesController" + "/" + (localError["funcion"]).(string))
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
		if total, err2 := helpers.CalculoSalarios(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": total}
		} else {
			panic(err2)
		}
	} else {
		panic(map[string]interface{}{"funcion": "CalcularTotalSalarios", "err": err.Error(), "status": "400"})
	}

	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/insertar_previnculaciones [post]
func (c *GestionPrevinculacionesController) InsertarPrevinculaciones() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPrevinculacionesController" + "/" + (localError["funcion"]).(string))
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
		if idRespuesta, err2 := helpers.InsertarPrevinculaciones(v); err2 == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": idRespuesta}
		} else {
			panic(err2)
		}
	} else {
		panic(map[string]interface{}{"funcion": "InsertarPrevinculaciones", "err": err.Error(), "status": "400"})
	}
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
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPrevinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	vig, err1 := strconv.Atoi(vigencia)
	per, err2 := strconv.Atoi(periodo)
	_, err3 := strconv.Atoi(tipoVinculacion)
	_, err4 := strconv.Atoi(facultad)
	_, err5 := strconv.Atoi(nivelAcademico)
	if (err1 != nil) || (err2 != nil) || (err3 != nil) || (err4 != nil) || (err5 != nil) || (vig == 0) || (per == 0) || (len(periodo) != 4) {
		panic(map[string]interface{}{"funcion": "ListarDocentesCargaHoraria1", "err1": "Error en los parametros de ingreso", "status": "400"})
	}

	if respuesta, err6 := helpers.ListarDocentesCargaHoraria(vigencia, periodo, tipoVinculacion, facultad, nivelAcademico); err6 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": respuesta.CargasLectivas.CargaLectiva}
	} else {
		panic(err6)
	}

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
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPrevinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	idRes, err1 := strconv.Atoi(idResolucion)
	if (idRes == 0) || (err1 != nil) {
		panic(map[string]interface{}{"funcion": "ListarDocentesPrevinculadosAll1", "err1": "Error en los parametros de ingreso", "status": "400"})
	}
	if v, err2 := helpers.ListarDocentesPrevinculadosAll(idResolucion, tipoVinculacion, tipoCancelacion, tipoAdicion, tipoReduccion); err2 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": v}
	} else {
		panic(err2)
	}
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
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ListarDocentesPrevinculados" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	idRes, err1 := strconv.Atoi(idResolucion)
	if (idRes == 0) || (err1 != nil) {
		panic(map[string]interface{}{"funcion": "ListarDocentesPrevinculados1", "err1": "Error en los parametros de ingreso", "status": "400"})
	}
	if v, err2 := helpers.ListarDocentesPrevinculados(idResolucion, tipoVinculacion); err2 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": v}
	} else {
		panic(err2)
	}
	c.ServeJSON()
}

// GestionPrevinculacionesController ...
// @Title GetCdpRpDocente
// @Description Get RPs de la vinculacion docente
// @Param num_vinculacion query string true "Número de la vinculación del docente"
// @Param vigencia query string true "Vigencia de la vinculación del docente"
// @Param identificacion query string true "Identificación del docente"
// @Success 201 {object}  models.RpDocente
// @Failure 403 :num_vinculacion is empty
// @Failure 403 :vigencia is empty
// @router /rp_docente/:num_vinculacion/:vigencia/:identificacion [get]
func (c *GestionPrevinculacionesController) GetCdpRpDocente() {
	num_vinculacion := c.Ctx.Input.Param(":num_vinculacion")
	vigencia := c.Ctx.Input.Param(":vigencia")
	identificacion := c.Ctx.Input.Param(":identificacion")
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPrevinculacionesController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	numV, err1 := strconv.Atoi(num_vinculacion)
	vigen, err2 := strconv.Atoi(vigencia)
	ident, err3 := strconv.Atoi(identificacion)
	if (err1 != nil) || (err2 != nil) || (err3 != nil) || (numV == 0) || (vigen == 0) || (ident == 0) {
		panic(map[string]interface{}{"funcion": "GetCertificacionDocumentosAprobados", "err": "Error en los parametros de ingreso", "status": "400"})
	}
	if rpdocente, err4 := helpers.GetCdpRpDocente(identificacion, num_vinculacion, vigencia); err4 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": rpdocente}
	} else {
		panic(err4)
	}
	c.ServeJSON()
}
