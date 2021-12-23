package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
)

// ConsultarDisponibilidadesController operations for ConsultarDisponibilidades
type ConsultarDisponibilidadesController struct {
	beego.Controller
}

// URLMapping ...
func (c *ConsultarDisponibilidadesController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("ListarDisponibilidades", c.ListarDisponibilidades)
	c.Mapping("TotalDisponibilidades", c.TotalDisponibilidades)
}

// GetAll ...
// @Title Get All
// @Description get Disponibilidad
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Disponibilidad
// @Failure 400 bad request
// @Failure 404 not found resource
// @router / [get]
func (c *ConsultarDisponibilidadesController) GetAll() {
	defer ErrorController(c)

	var query string
	query = c.GetString("query")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	if d, err3 := helpers.GetAllDisponibilidad(query, limit, offset); err3 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": d}
	} else {
		panic(err3)
	}
	c.ServeJSON()
}

// ListaDisponibilidades ...
// @Title ListaDisponibilidades
// @Description get Disponibilidad by vigencia
// @Param	vigencia	query	string	false	"vigencia de la lista"
// @Param	UnidadEjecutora	query	string	false	"unidad ejecutora de las solicitudes a consultar"
// @Param	query	query	string	false	"query de filtrado para la lista de los cdp"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Disponibilidad
// @Failure 400 ad request
// @Failure 404 aborted by server
// @router /ListaDisponibilidades/:vigencia [get]
func (c *ConsultarDisponibilidadesController) ListarDisponibilidades() {
	defer ErrorController(c)

	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	UnidadEjecutora, err := c.GetInt("UnidadEjecutora")
	query := c.GetString("query")
	limit, err1 := c.GetInt("limit")
	offset, err2 := c.GetInt("offset")

	vigencia, err3 := strconv.Atoi(vigenciaStr)

	if (offset < 0) || (err != nil) || (err1 != nil) || (err2 != nil) || (err3 != nil) {
		panic(map[string]interface{}{"funcion": "ListarDisponibilidades", "err": "Error en los parámetros de ingreso", "status": "400"})
	}

	if respuesta, err := helpers.ListarDisponibilidades(vigencia, UnidadEjecutora, limit, offset, query); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": respuesta}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// TotalDisponibilidades ...
// @Title TotalDisponibilidades
// @Description numero de disponibilidades segun vigencia o rango de fechas
// @Param	vigencia		query 	string	true		"vigencia para la consulta del total de disponibilidades"
// @Param	UnidadEjecutora	query	string	false	"unidad ejecutora de las solicitudes a consultar"
// @Success 200 {int} total
// @Failure 400 bad request
// @Failure 404 aborted by server
// @router /TotalDisponibilidades/:vigencia [get]
func (c *ConsultarDisponibilidadesController) TotalDisponibilidades() {
	defer ErrorController(c)

	UnidadEjecutora, err := c.GetInt("UnidadEjecutora")
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	vigencia, err2 := strconv.Atoi(vigenciaStr)

	if (err != nil) || (err2 != nil) {
		panic(map[string]interface{}{"funcion": "TotalDisponibilidades", "err": "Error en los parámetros de ingreso", "status": "400"})
	}

	if total, err3 := helpers.GetTotalDisponibilidades(vigencia, UnidadEjecutora); err3 == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": total}
	} else {
		panic(err3)
	}
	c.ServeJSON()
}

func ErrorController(c *ConsultarDisponibilidadesController) {
	if err := recover(); err != nil {
		logs.Error(err)
		localError := err.(map[string]interface{})
		c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "GestionPlantillasController" + "/" + (localError["funcion"]).(string))
		c.Data["data"] = (localError["err"])
		if status, ok := localError["status"]; ok {
			c.Abort(status.(string))
		} else {
			c.Abort("500")
		}
	}
}
