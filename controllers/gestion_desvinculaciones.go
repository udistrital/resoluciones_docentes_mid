package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
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
	v := []models.VinculacionDocente{}

	err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v)
	beego.Debug(v)
	if err != nil {
		beego.Error("Error de consulta en vinculacion", err)
		c.Abort("403")
	}
	for x, pos := range v {
		documento_identidad, _ := strconv.Atoi(pos.IdPersona)
		v[x].NombreCompleto = helpers.BuscarNombreProveedor(documento_identidad)
		v[x].NumeroDisponibilidad = helpers.BuscarNumeroDisponibilidad(pos.Disponibilidad)
		v[x].Dedicacion = helpers.BuscarNombreDedicacion(pos.IdDedicacion.Id)
		v[x].LugarExpedicionCedula = helpers.BuscarLugarExpedicion(pos.IdPersona)
	}
	if v == nil {
		v = []models.VinculacionDocente{}
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
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
	var v []models.VinculacionDocente
	var modRes []models.ModificacionResolucion
	var modVin []models.ModificacionVinculacion
	var cv models.VinculacionDocente
	// if 3 - modificacion_resolucion
	if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=resolucionNueva:"+id_resolucion, &modRes); err == nil {
		// if 2 - modificacion_vinculacion
		t := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAdmin") + "/" + beego.AppConfig.String("NscrudAdmin") + "/modificacion_vinculacion/?limit=-1&query=modificacion_resolucion:" + strconv.Itoa(modRes[0].Id)
		beego.Info(t)
		if err := helpers.GetJson(t, &modVin); err == nil {
			//for vinculaciones
			for _, vinculacion := range modVin {
				beego.Info(fmt.Sprintf("%+v", vinculacion.VinculacionDocenteCancelada))
				// if 1 - vinculacion_docente
				if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion.VinculacionDocenteCancelada.Id), &cv); err == nil {
					documento_identidad, _ := strconv.Atoi(vinculacion.VinculacionDocenteCancelada.IdPersona)
					cv.NombreCompleto = helpers.BuscarNombreProveedor(documento_identidad)
					cv.NumeroDisponibilidad = helpers.BuscarNumeroDisponibilidad(vinculacion.VinculacionDocenteCancelada.Disponibilidad)
					cv.Dedicacion = helpers.BuscarNombreDedicacion(vinculacion.VinculacionDocenteCancelada.IdDedicacion.Id)
					cv.LugarExpedicionCedula = helpers.BuscarLugarExpedicion(vinculacion.VinculacionDocenteCancelada.IdPersona)
					cv.NumeroSemanasNuevas = vinculacion.VinculacionDocenteCancelada.NumeroSemanas - vinculacion.VinculacionDocenteRegistrada.NumeroSemanas
				} else { // if 1 - vinculacion_docente
					fmt.Println("Error de consulta en vinculacion, solucioname!!!, if 1 - vinculacion_docente: ", err)
				}
				v = append(v, cv)
			} //fin for vinculaciones
		} else { // if 2 - modificacion_vinculacion
			fmt.Println("Error de consulta en modificacion_vinculacion, solucioname!!!, if 2 - modificacion_vinculacion: ", err)
		}
	} else { // if 3 - modificacion_resolucion
		fmt.Println("Error de consulta en modificacion_resolucion, solucioname!!!, if 3 - modificacion_resolucion: ", err)
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()
}

// AnularModificaciones ...
// @Title AnularModificaciones
// @Description create AnularModificaciones
// @Success 201 {string}
// @Failure 403 body is empty
// @router /anular_modificaciones [post]
// Se usa para cuando se anulan resoluciones modificatorias completas
func (c *GestionDesvinculacionesController) AnularModificaciones() {
	var v []models.VinculacionDocente
	var modRes []models.ModificacionResolucion
	var respuesta_vinculacion string
	var vinculacion_cancelada []models.VinculacionDocente
	var respuesta_delete_vin string
	var respuesta_delete string
	var respuesta_total string
	var respuesta_modificacion_vinculacion []models.ModificacionVinculacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta_total = "OK"
		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=resolucionNueva:"+strconv.Itoa(v[0].IdResolucion.Id), &modRes)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
			respuesta_total = "Error"
		}
		for _, pos := range v {
			//Se trae información de tabla de traza modificacion_vinculacion, para saber cuál vinculación hay que poner en true y cuál eliminar
			query := "?limit=-1&query=ModificacionResolucion.Id:" + strconv.Itoa(modRes[0].Id) + ",VinculacionDocenteRegistrada.Id:" + strconv.Itoa(pos.Id)
			err2 := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion"+query, &respuesta_modificacion_vinculacion)
			fmt.Println("modificacion_vinculacion", respuesta_modificacion_vinculacion)
			if err2 != nil {
				beego.Error(err)
				c.Abort("400")
				respuesta_total = "Error"
			}

			//se trae informacion de vinculación que fue cancelada
			query2 := "?limit=-1&query=Id:" + strconv.Itoa(respuesta_modificacion_vinculacion[0].VinculacionDocenteCancelada.Id)
			err2 = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query2, &vinculacion_cancelada)
			fmt.Println("vinculacion_cancelada", vinculacion_cancelada)
			if err2 != nil {
				beego.Error(err)
				c.Abort("400")
				respuesta_total = "Error"
			}
			//se cambia a true vinculación que fue cancelada
			vinculacion_cancelada[0].Estado = true

			//Se le cambia estado en bd a vinculación cancelada
			err2 = helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion_cancelada[0].Id), "PUT", &respuesta_vinculacion, vinculacion_cancelada[0])
			fmt.Println("respuesta_vinculacion", respuesta_vinculacion)
			if err2 != nil {
				beego.Error(err)
				c.Abort("400")
				respuesta_total = "Error"
			}

			//se elimina registro en modificacion_vinculacion
			err2 = helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/"+strconv.Itoa(respuesta_modificacion_vinculacion[0].Id), "DELETE", &respuesta_delete, respuesta_modificacion_vinculacion[0])
			if err2 != nil {
				beego.Error(err)
				c.Abort("400")
				respuesta_total = "Error"
			}

			//Se elimina vinculacion nueva
			err2 = helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "DELETE", &respuesta_delete_vin, pos)
			fmt.Println("respuesta_eliminar_vin_nueva", respuesta_delete_vin)
			if err2 != nil {
				beego.Error(err)
				c.Abort("400")
				respuesta_total = "Error"
			}
		}
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
	var respuesta_vinculacion string
	var vinculacion_cancelada []models.VinculacionDocente
	var respuesta_delete_vin string
	var respuesta_delete string
	var respuesta_total string
	var respuesta_modificacion_vinculacion []models.ModificacionVinculacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta_total = "OK"

		//Se trae información de tabla de traza modificacion_vinculacion, para saber cuál vinculación hay que poner en true y cuál eliminar
		query := "?limit=-1&query=ModificacionResolucion.Id:" + strconv.Itoa(v.IdModificacionResolucion) + ",VinculacionDocenteRegistrada.Id:" + strconv.Itoa(v.DocentesDesvincular[0].Id)
		if err2 := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion"+query, &respuesta_modificacion_vinculacion); err2 == nil {
			fmt.Println("modificacion_vinculacion", respuesta_modificacion_vinculacion)
			respuesta_total = "OK"
		} else {
			respuesta_total = "error"
		}

		//se trae informacion de vinculacion que fue cancelada
		query2 := "?limit=-1&query=Id:" + strconv.Itoa(respuesta_modificacion_vinculacion[0].VinculacionDocenteCancelada.Id)
		if err2 := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query2, &vinculacion_cancelada); err2 == nil {
			fmt.Println("vinculacion_cancelada", vinculacion_cancelada)
			respuesta_total = "OK"
		} else {
			respuesta_total = "error"
		}
		//se cambia a true vinculacion que fue cancelada
		vinculacion_cancelada[0].Estado = true
		fmt.Println("nuevo estado de vinculacion cancelada", vinculacion_cancelada)

		//Se le cambia estado en bd a vinculacion cancelada

		if err2 := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion_cancelada[0].Id), "PUT", &respuesta_vinculacion, vinculacion_cancelada[0]); err2 == nil {
			fmt.Println("respuesta_vinculacion", respuesta_vinculacion)
			respuesta_total = "OK"
		} else {
			respuesta_total = "error"
		}

		//se elimina registro en modificacion_vinculacion

		if err2 := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/"+strconv.Itoa(respuesta_modificacion_vinculacion[0].Id), "DELETE", &respuesta_delete, respuesta_modificacion_vinculacion[0]); err2 == nil {
			respuesta_total = "OK"
		} else {
			respuesta_total = "error"
		}

		//Se elimina vinculacion nueva
		if err2 := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.DocentesDesvincular[0].Id), "DELETE", &respuesta_delete_vin, v.DocentesDesvincular[0]); err2 == nil {
			fmt.Println("respuesta_eliminar_vin_nueva", respuesta_delete_vin)
			respuesta_total = "OK"
		} else {
			respuesta_total = "error"
		}

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

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	categoria, _, err := helpers.Buscar_Categoria_Docente(strconv.Itoa(v.VigenciaCarga), strconv.Itoa(v.PeriodoCarga), v.IdPersona)
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}

	respuesta := "OK"
	if categoria == "" {
		respuesta = "Sin categoría"
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
	var respuesta = "OK"
	var respuestaApropiacion models.DatosApropiacion
	var saldoDisponibilidad float64

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &validacion)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	validacion.DocentesDesvincular[0].NumeroHorasSemanales = validacion.DocentesDesvincular[0].NumeroHorasNuevas
	validacion.DocentesDesvincular[0].NumeroSemanas = validacion.DocentesDesvincular[0].NumeroSemanasNuevas

	validacion.DocentesDesvincular, err = helpers.CalcularSalarioPrecontratacion(validacion.DocentesDesvincular)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	valorContrato := validacion.DocentesDesvincular[0].ValorContrato

	if err2 := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/SaldoCdp", "POST", &respuestaApropiacion, &validacion.DisponibilidadNueva); err2 == nil {
		saldoDisponibilidad = float64(respuestaApropiacion.Saldo)
	} else {
		beego.Error(err)
		c.Abort("400")
	}

	if saldoDisponibilidad-valorContrato < 0 {
		respuesta = "Error CDP"
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
	var respuesta_mod_vin models.ModificacionVinculacion
	var respuesta string
	var vinculacion_nueva int
	var temp_vinculacion [1]models.VinculacionDocente

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	//CAMBIAR ESTADO DE VINCULACIÓN DOCENTE
	for _, pos := range v.DocentesDesvincular {
		pos.NumeroRp = 0
		pos.VigenciaRp = 0
		err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "PUT", &respuesta, pos)
		//TODO: unificar errores
		if err != nil {
			err = fmt.Errorf("error al cambiar estado en vinculación docente al adicionar horas %s", err)
			beego.Error(err)
			c.Abort("403")
		}
		beego.Info("respuesta", respuesta)
		beego.Info("fechaAD", v.DocentesDesvincular[0].FechaInicioNueva)
		temp_vinculacion[0] = models.VinculacionDocente{
			IdPersona:            v.DocentesDesvincular[0].IdPersona,
			NumeroHorasSemanales: v.DocentesDesvincular[0].NumeroHorasNuevas,
			NumeroSemanas:        v.DocentesDesvincular[0].NumeroSemanasNuevas,
			IdResolucion:         &models.ResolucionVinculacionDocente{Id: v.IdNuevaResolucion},
			IdDedicacion:         v.DocentesDesvincular[0].IdDedicacion,
			IdProyectoCurricular: v.DocentesDesvincular[0].IdProyectoCurricular,
			Categoria:            v.DocentesDesvincular[0].Categoria,
			Dedicacion:           v.DocentesDesvincular[0].Dedicacion,
			NivelAcademico:       v.DocentesDesvincular[0].NivelAcademico,
			Disponibilidad:       v.DisponibilidadNueva.Id,
			Vigencia:             v.DocentesDesvincular[0].Vigencia,
			FechaInicio:          v.DocentesDesvincular[0].FechaInicioNueva,
			NumeroRp:             v.DocentesDesvincular[0].NumeroRp,
			VigenciaRp:           v.DocentesDesvincular[0].VigenciaRp,
			DependenciaAcademica: v.DocentesDesvincular[0].DependenciaAcademica,
		}

		//CREAR NUEVA Vinculacion
		vinculacion_nueva, err = InsertarDesvinculaciones(temp_vinculacion)
		if err != nil {
			beego.Error("error al realizar vinculacion nueva", err)
			c.Abort("400")
		}

		//INSERCION  TABLA  DE TRAZA MODIFICACION VINCULACION
		for _, pos := range v.DocentesDesvincular {
			temp := models.ModificacionVinculacion{
				ModificacionResolucion:       &models.ModificacionResolucion{Id: v.IdModificacionResolucion},
				VinculacionDocenteCancelada:  &models.VinculacionDocente{Id: pos.Id},
				VinculacionDocenteRegistrada: &models.VinculacionDocente{Id: vinculacion_nueva},
				Horas: pos.NumeroHorasNuevas,
			}
			err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/", "POST", &respuesta_mod_vin, temp)

			if err != nil {
				beego.Error("error en actualizacion de modificacion vinculacion de modificacion vinculacion", err)
				respuesta = "error"
			} else {
				beego.Info("respuesta modificacion vin", respuesta_mod_vin)
				respuesta = "OK"
			}
		}
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
	var respuesta_mod_vin models.ModificacionVinculacion
	// var respuesta string
	var vinculacion_nueva int
	var temp_vinculacion [1]models.VinculacionDocente

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	beego.Debug("para poner en false", v)

	for _, pos := range v.DocentesDesvincular {

		numerorp := pos.NumeroRp
		vigenciarp := pos.VigenciaRp

		pos.NumeroRp = 0
		pos.VigenciaRp = 0
		err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "PUT", &respuesta, pos)
		if err != nil {
			beego.Error("error en json", err)
			c.Abort("403")
		}

		//Verificar objeto para crear nuevas resoluciones
		temp_vinculacion[0] = models.VinculacionDocente{
			IdPersona:            pos.IdPersona,
			NumeroHorasSemanales: pos.NumeroHorasSemanales,
			NumeroSemanas:        pos.NumeroSemanasNuevas,
			IdResolucion:         &models.ResolucionVinculacionDocente{Id: v.IdNuevaResolucion},
			IdDedicacion:         pos.IdDedicacion,
			IdProyectoCurricular: pos.IdProyectoCurricular,
			Categoria:            pos.Categoria,
			Dedicacion:           pos.Dedicacion,
			NivelAcademico:       pos.NivelAcademico,
			Disponibilidad:       pos.Disponibilidad,
			Vigencia:             pos.Vigencia,
			NumeroRp:             numerorp,
			VigenciaRp:           vigenciarp,
			DependenciaAcademica: pos.DependenciaAcademica,
		}
		fmt.Println("RP: ", temp_vinculacion[0].NumeroRp)
		//CREAR NUEVA Vinculacion
		vinculacion_nueva, err = InsertarDesvinculaciones(temp_vinculacion)
		if err != nil {
			beego.Error("error al realizar vinculacion nueva", err)
			c.Abort("400")
		}

		//INSERCION  TABLA  DE TRAZA MODIFICACION VINCULACION
		temp := models.ModificacionVinculacion{
			ModificacionResolucion:       &models.ModificacionResolucion{Id: v.IdModificacionResolucion},
			VinculacionDocenteCancelada:  &models.VinculacionDocente{Id: pos.Id},
			VinculacionDocenteRegistrada: &models.VinculacionDocente{Id: vinculacion_nueva},
			Horas: pos.NumeroHorasSemanales,
		}
		errorMod := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/", "POST", &respuesta_mod_vin, temp)

		if errorMod != nil {
			beego.Error("error en actualizacion de modificacion vinculacion de modificacion vinculacion", err)
			respuesta = "error"
		} else {
			beego.Info("respuesta modificacion vin", respuesta_mod_vin)
			respuesta = "OK"
		}
	}

	c.Data["json"] = respuesta

	c.ServeJSON()

}



func InsertarDesvinculaciones(v [1]models.VinculacionDocente) (id int, err error) {
	var d []models.VinculacionDocente
	json_ejemplo, err := json.Marshal(v)
	if err != nil {
		beego.Error(err)
		return id, err
	}
	err = json.Unmarshal(json_ejemplo, &d)

	if err != nil {
		beego.Error(err)
		return id, err
	}

	//TODO: unificar cont con error
	d, err = helpers.CalcularSalarioPrecontratacion(d)
	if err != nil {
		return id, err
	}

	err = helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &id, &d)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return id, err
}