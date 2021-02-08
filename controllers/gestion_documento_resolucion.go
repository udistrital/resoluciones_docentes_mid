package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
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
	var contenidoResolucion models.ResolucionCompleta
	var ordenador_gasto []models.OrdenadorGasto
	var jefe_dependencia []models.JefeDependencia
	var query string

	if err2 := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contenido_resolucion/"+id_resolucion, &contenidoResolucion); err2 == nil {
		query = "?limit=-1&query=DependenciaId:" + id_facultad

		if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/"+query, &ordenador_gasto); err == nil {
			if ordenador_gasto == nil {
				if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/1", &ordenador_gasto); err == nil {
					contenidoResolucion.OrdenadorGasto = ordenador_gasto[0]
				} else {
					fmt.Println("Error al consultar ordenador 1", err)
				}
			} else {
				contenidoResolucion.OrdenadorGasto = ordenador_gasto[0]
			}

		} else {

			fmt.Println("Error al consultar ordenador del gasto", err)
		}

	} else {
		fmt.Println("Error al consultar contenido", err2)
	}

	fecha_actual := time.Now().Format("2006-01-02")
	query = "?query=DependenciaId:" + id_facultad + ",FechaFin__gte:" + fecha_actual + ",FechaInicio__lte:" + fecha_actual
	if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/"+query, &jefe_dependencia); err == nil {
		contenidoResolucion.OrdenadorGasto.NombreOrdenador = BuscarNombreProveedor(jefe_dependencia[0].TerceroId)

	} else {

	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = contenidoResolucion
	c.ServeJSON()

}
