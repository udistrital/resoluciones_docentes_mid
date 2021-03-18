package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func AnularModificaciones(v []models.VinculacionDocente) (respuesta_total string) {
	var modRes []models.ModificacionResolucion
	var respuesta_vinculacion string
	var vinculacion_cancelada []models.VinculacionDocente
	var respuesta_delete_vin string
	var respuesta_delete string
	var respuesta_modificacion_vinculacion []models.ModificacionVinculacion

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=resolucionNueva:"+strconv.Itoa(v[0].IdResolucion.Id), &modRes); err == nil && response == 200 {
		respuesta_total = "OK"
	} else {
		beego.Error(err)
		respuesta_total = "Error"
	}
	for _, pos := range v {
		//Se trae información de tabla de traza modificacion_vinculacion, para saber cuál vinculación hay que poner en true y cuál eliminar
		query := "?limit=-1&query=ModificacionResolucion.Id:" + strconv.Itoa(modRes[0].Id) + ",VinculacionDocenteRegistrada.Id:" + strconv.Itoa(pos.Id)
		if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion"+query, &respuesta_modificacion_vinculacion); err2 == nil && response2 == 200 {
			fmt.Println("modificacion_vinculacion", respuesta_modificacion_vinculacion)
			respuesta_total = "OK"
		} else {
			beego.Error(err2)
			respuesta_total = "Error"
		}

		//se trae informacion de vinculación que fue cancelada
		query2 := "?limit=-1&query=Id:" + strconv.Itoa(respuesta_modificacion_vinculacion[0].VinculacionDocenteCancelada.Id)
		if response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query2, &vinculacion_cancelada); err3 == nil && response3 == 200 {
			fmt.Println("vinculacion_cancelada", vinculacion_cancelada)
			respuesta_total = "OK"
		} else {
			beego.Error(err3)
			respuesta_total = "Error"
		}
		//se cambia a true vinculación que fue cancelada
		vinculacion_cancelada[0].Estado = true

		//Se le cambia estado en bd a vinculación cancelada
		if err4 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion_cancelada[0].Id), "PUT", &respuesta_vinculacion, vinculacion_cancelada[0]); err4 == nil {
			fmt.Println("respuesta_vinculacion", respuesta_vinculacion)
			respuesta_total = "OK"
		} else {
			beego.Error(err4)
			respuesta_total = "Error"
		}

		//se elimina registro en modificacion_vinculacion
		if err5 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/"+strconv.Itoa(respuesta_modificacion_vinculacion[0].Id), "DELETE", &respuesta_delete, respuesta_modificacion_vinculacion[0]); err5 == nil {
			respuesta_total = "OK"
		} else {
			beego.Error(err5)
			respuesta_total = "Error"
		}

		//Se elimina vinculacion nueva
		if err6 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "DELETE", &respuesta_delete_vin, pos); err6 == nil {
			fmt.Println("respuesta_eliminar_vin_nueva", respuesta_delete_vin)
			respuesta_total = "OK"
		} else {
			beego.Error(err6)
			respuesta_total = "Error"
		}
	}
	return respuesta_total
}

func AnularAdicionDocente(v models.Objeto_Desvinculacion) (respuesta_total string) {
	var respuesta_vinculacion string
	var vinculacion_cancelada []models.VinculacionDocente
	var respuesta_delete_vin string
	var respuesta_delete string
	var respuesta_modificacion_vinculacion []models.ModificacionVinculacion

	//Se trae información de tabla de traza modificacion_vinculacion, para saber cuál vinculación hay que poner en true y cuál eliminar
	query := "?limit=-1&query=ModificacionResolucion.Id:" + strconv.Itoa(v.IdModificacionResolucion) + ",VinculacionDocenteRegistrada.Id:" + strconv.Itoa(v.DocentesDesvincular[0].Id)
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion"+query, &respuesta_modificacion_vinculacion); err == nil && response == 200 {
		fmt.Println("modificacion_vinculacion", respuesta_modificacion_vinculacion)
		respuesta_total = "OK"
	} else {
		respuesta_total = "error"
	}

	//se trae informacion de vinculacion que fue cancelada
	query2 := "?limit=-1&query=Id:" + strconv.Itoa(respuesta_modificacion_vinculacion[0].VinculacionDocenteCancelada.Id)
	if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query2, &vinculacion_cancelada); err2 == nil && response2 == 200 {
		fmt.Println("vinculacion_cancelada", vinculacion_cancelada)
		respuesta_total = "OK"
	} else {
		respuesta_total = "error"
	}
	//se cambia a true vinculacion que fue cancelada
	vinculacion_cancelada[0].Estado = true
	fmt.Println("nuevo estado de vinculacion cancelada", vinculacion_cancelada)

	//Se le cambia estado en bd a vinculacion cancelada

	if err3 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion_cancelada[0].Id), "PUT", &respuesta_vinculacion, vinculacion_cancelada[0]); err3 == nil {
		fmt.Println("respuesta_vinculacion", respuesta_vinculacion)
		respuesta_total = "OK"
	} else {
		respuesta_total = "error"
	}

	//se elimina registro en modificacion_vinculacion

	if err4 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/"+strconv.Itoa(respuesta_modificacion_vinculacion[0].Id), "DELETE", &respuesta_delete, respuesta_modificacion_vinculacion[0]); err4 == nil {
		respuesta_total = "OK"
	} else {
		respuesta_total = "error"
	}

	//Se elimina vinculacion nueva
	if err5 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.DocentesDesvincular[0].Id), "DELETE", &respuesta_delete_vin, v.DocentesDesvincular[0]); err5 == nil {
		fmt.Println("respuesta_eliminar_vin_nueva", respuesta_delete_vin)
		respuesta_total = "OK"
	} else {
		respuesta_total = "error"
	}

	return respuesta_total
}

func ConsultarCategoria(v models.VinculacionDocente) (respuesta string) {

	categoria, _, err := Buscar_Categoria_Docente(strconv.Itoa(v.VigenciaCarga), strconv.Itoa(v.PeriodoCarga), v.IdPersona)
	if err != nil {
		beego.Error(err)
	}

	respuesta = "OK"
	if categoria == "" {
		respuesta = "Sin categoría"
	}

	return respuesta
}

func ValidarSaldoCDP(validacion models.Objeto_Desvinculacion) (respuesta string) {

	var respuestaApropiacion models.DatosApropiacion
	var saldoDisponibilidad float64

	validacion.DocentesDesvincular[0].NumeroHorasSemanales = validacion.DocentesDesvincular[0].NumeroHorasNuevas
	validacion.DocentesDesvincular[0].NumeroSemanas = validacion.DocentesDesvincular[0].NumeroSemanasNuevas

	if docentes, err := CalcularSalarioPrecontratacion(validacion.DocentesDesvincular); err == nil {
		validacion.DocentesDesvincular = docentes
		respuesta = "OK"
	} else {
		respuesta = "Error"
	}

	valorContrato := validacion.DocentesDesvincular[0].ValorContrato

	if err2 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/SaldoCdp", "POST", &respuestaApropiacion, &validacion.DisponibilidadNueva); err2 == nil {
		saldoDisponibilidad = float64(respuestaApropiacion.Saldo)
		respuesta = "OK"
	} else {
		respuesta = "Error"
	}

	if saldoDisponibilidad-valorContrato < 0 {
		respuesta = "Error CDP"
	}
	return respuesta
}

func AdicionarHoras(v models.Objeto_Desvinculacion) (respuesta string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/AprobacionPagosContratistas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var respuesta_mod_vin models.ModificacionVinculacion
	var vinculacion_nueva int
	var temp_vinculacion [1]models.VinculacionDocente

	//CAMBIAR ESTADO DE VINCULACIÓN DOCENTE
	for _, pos := range v.DocentesDesvincular {
		pos.NumeroRp = 0
		pos.VigenciaRp = 0
		err1 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "PUT", &respuesta, pos)
		//TODO: unificar errores
		if err1 != nil {
			logs.Error(err1)
			outputError = map[string]interface{}{"funcion": "/AdicionarHoras", "err1": err1.Error(), "status": "502"}
			return respuesta, outputError
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
		var err2 map[string]interface{}
		vinculacion_nueva, err2 = InsertarDesvinculaciones(temp_vinculacion)
		if err2 != nil {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/AdicionarHoras", "err1": err2, "status": "502"}
			return respuesta, outputError
		}

		//INSERCION  TABLA  DE TRAZA MODIFICACION VINCULACION
		for _, pos := range v.DocentesDesvincular {
			temp := models.ModificacionVinculacion{
				ModificacionResolucion:       &models.ModificacionResolucion{Id: v.IdModificacionResolucion},
				VinculacionDocenteCancelada:  &models.VinculacionDocente{Id: pos.Id},
				VinculacionDocenteRegistrada: &models.VinculacionDocente{Id: vinculacion_nueva},
				Horas:                        pos.NumeroHorasNuevas,
			}
			err3 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/", "POST", &respuesta_mod_vin, temp)

			if err3 != nil {
				logs.Error(err3)
				outputError = map[string]interface{}{"funcion": "/AdicionarHoras", "err3": err3.Error(), "status": "502"}
				return respuesta, outputError
			} else {
				beego.Info("respuesta modificacion vin", respuesta_mod_vin)
				respuesta = "OK"
			}
		}
	}

	return respuesta, outputError
}

func ActualizarVinculacionesCancelacion(v models.Objeto_Desvinculacion) (respuesta string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ActualizarVinculacionesCancelacion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var respuesta_mod_vin models.ModificacionVinculacion
	// var respuesta string
	var vinculacion_nueva int
	var temp_vinculacion [1]models.VinculacionDocente

	beego.Debug("para poner en false", v)

	for _, pos := range v.DocentesDesvincular {

		numerorp := pos.NumeroRp
		vigenciarp := pos.VigenciaRp

		pos.NumeroRp = 0
		pos.VigenciaRp = 0
		err1 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "PUT", &respuesta, pos)
		if err1 != nil {
			logs.Error(err1)
			outputError = map[string]interface{}{"funcion": "/AprobacionPagosContratistas1", "err1": err1.Error(), "status": "502"}
			return respuesta, outputError
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
		var err2 map[string]interface{}
		vinculacion_nueva, err2 = InsertarDesvinculaciones(temp_vinculacion)
		if err2 != nil {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/AprobacionPagosContratistas2", "err2": err2, "status": "502"}
			return respuesta, outputError
		}

		//INSERCION  TABLA  DE TRAZA MODIFICACION VINCULACION
		temp := models.ModificacionVinculacion{
			ModificacionResolucion:       &models.ModificacionResolucion{Id: v.IdModificacionResolucion},
			VinculacionDocenteCancelada:  &models.VinculacionDocente{Id: pos.Id},
			VinculacionDocenteRegistrada: &models.VinculacionDocente{Id: vinculacion_nueva},
			Horas:                        pos.NumeroHorasSemanales,
		}
		errorMod := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/", "POST", &respuesta_mod_vin, temp)

		if errorMod != nil {
			logs.Error(errorMod)
			outputError = map[string]interface{}{"funcion": "/AprobacionPagosContratistas3", "errorMod": errorMod.Error(), "status": "502"}
			return respuesta, outputError
		} else {
			beego.Info("respuesta modificacion vin", respuesta_mod_vin)
			respuesta = "OK"
		}
	}

	return respuesta, outputError

}

func InsertarDesvinculaciones(v [1]models.VinculacionDocente) (id int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var d []models.VinculacionDocente
	json_ejemplo, err1 := json.Marshal(v)
	if err1 != nil {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones1", "err1": err1.Error(), "status": "404"}
		return id, outputError
	}
	err2 := json.Unmarshal(json_ejemplo, &d)

	if err2 != nil {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones2", "err2": err2.Error(), "status": "404"}
		return id, outputError
	}

	//TODO: unificar cont con error
	var err3 map[string]interface{}
	d, err3 = CalcularSalarioPrecontratacion(d)
	if err3 != nil {
		logs.Error(err3)
		outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones3", "err3": err3, "status": "404"}
		return id, outputError
	}

	err4 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &id, &d)
	if err4 != nil {
		logs.Error(err4)
		outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones4", "err4": err4, "status": "404"}
		return id, outputError
	}
	return id, outputError
}
