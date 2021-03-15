package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
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
// @Param id_resolucion query string false "año a consultar"
// @Param id_facultad query string false "periodo a listar"
// @Success 201 {object} models.ResolucionCompleta
// @Failure 403 body is empty
// @router /get_contenido_resolucion [get]
func (c *GestionDocumentoResolucionController) GetContenidoResolucion() {
	id_resolucion := c.GetString("id_resolucion")
	id_facultad := c.GetString("id_facultad")

	fmt.Println("ENTRA")

	contenidoResolucion := helpers.GetContenidoResolucion(id_resolucion, id_facultad)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = contenidoResolucion
	c.ServeJSON()

}
