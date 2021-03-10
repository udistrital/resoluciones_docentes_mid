package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
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
// @Failure 403 body is empty
// @router /get_resoluciones_aprobadas [get]
func (c *GestionResolucionesController) GetResolucionesAprobadas() {
	var resolucion_vinculacion_aprobada []models.ResolucionVinculacion

	query := c.GetString("query")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	resolucion_vinculacion_aprobada = helpers.GetResolucionesAprobadas(query, limit, offset)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = resolucion_vinculacion_aprobada
	c.ServeJSON()
}

// GestionResolucionesController ...
// @Title getResolucionesInscritas
// @Description create  getResolucionesInscritas
// @Param vigencia query string false "año a consultar"
// @Success 201 {object} []models.ResolucionVinculacion
// @Failure 403 body is empty
// @router /get_resoluciones_inscritas [get]
func (c *GestionResolucionesController) GetResolucionesInscritas() {
	var resolucion_vinculacion []models.ResolucionVinculacion
	var query []string

	query = c.GetStrings("query")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	resolucion_vinculacion = helpers.GetResolucionesInscritas(query, limit, offset)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = resolucion_vinculacion
	c.ServeJSON()
}

// InsertarResolucionCompleta ...
// @Title InsertarResolucionCompleta
// @Description create InsertarResolucionCompleta
// @Success 201 {int} models.Resolucion
// @Failure 403 body is empty
// @router /insertar_resolucion_completa [post]
func (c *GestionResolucionesController) InsertarResolucionCompleta() {
	var v models.ObjetoResolucion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		id_resolucion_creada, control := helpers.InsertarResolucionCompleta(v)
		if control {
			fmt.Println("okey")
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = id_resolucion_creada
		} else {
			fmt.Println("not okey")
			c.Data["json"] = "Error"
		}
	} else {
		fmt.Println("error al leer objeto resolucion", err)
	}

	c.ServeJSON()
}
