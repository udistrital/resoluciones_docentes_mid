package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func AnularModificaciones(v []models.VinculacionDocente) (outputError map[string]interface{}) {
	var modRes []models.ModificacionResolucion
	var respuesta_vinculacion string
	var vinculacion_cancelada []models.VinculacionDocente
	var respuesta_delete_vin string
	var respuesta_delete string
	var respuesta_modificacion_vinculacion []models.ModificacionVinculacion
	var respuesta_peticion map[string]interface{}

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=resolucionNuevaId:"+strconv.Itoa(v[0].ResolucionVinculacionDocenteId.Id), &respuesta_peticion); err == nil && response == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &modRes)
	} else {
		beego.Error(err)
		outputError = map[string]interface{}{"funcion": "/AnularModificaciones6", "err": err.Error(), "status": "404"}
		return outputError
	}
	for _, pos := range v {
		//Se trae información de tabla de traza modificacion_vinculacion, para saber cuál vinculación hay que poner en true y cuál eliminar
		query := "?limit=-1&query=ModificacionResolucionId.Id:" + strconv.Itoa(modRes[0].Id) + ",VinculacionDocenteRegistradaId.Id:" + strconv.Itoa(pos.Id)
		if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion"+query, &respuesta_peticion); err2 == nil && response2 == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_modificacion_vinculacion)
			fmt.Println("modificacion_vinculacion", respuesta_modificacion_vinculacion)
		} else {
			beego.Error(err2)
			outputError = map[string]interface{}{"funcion": "/AnularModificaciones5", "err": err2.Error(), "status": "404"}
			return outputError
		}

		//se trae informacion de vinculación que fue cancelada
		query2 := "?limit=-1&query=Id:" + strconv.Itoa(respuesta_modificacion_vinculacion[0].VinculacionDocenteCanceladaId.Id)
		if response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query2, &respuesta_peticion); err3 == nil && response3 == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &vinculacion_cancelada)
			fmt.Println("vinculacion_cancelada", vinculacion_cancelada)
		} else {
			beego.Error(err3)
			outputError = map[string]interface{}{"funcion": "/AnularModificaciones4", "err": err3.Error(), "status": "404"}
			return outputError
		}
		//se cambia a true vinculación que fue cancelada
		vinculacion_cancelada[0].Activo = true

		//Se le cambia estado en bd a vinculación cancelada
		if err4 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion_cancelada[0].Id), "PUT", &respuesta_peticion, vinculacion_cancelada[0]); err4 == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_vinculacion)
			fmt.Println("respuesta_vinculacion", respuesta_vinculacion)
		} else {
			beego.Error(err4)
			outputError = map[string]interface{}{"funcion": "/AnularModificaciones3", "err": err4.Error(), "status": "404"}
			return outputError
		}

		//se elimina registro en modificacion_vinculacion
		if err5 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/"+strconv.Itoa(respuesta_modificacion_vinculacion[0].Id), "DELETE", &respuesta_peticion, respuesta_modificacion_vinculacion[0]); err5 == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_delete)
		} else {
			beego.Error(err5)
			outputError = map[string]interface{}{"funcion": "/AnularModificaciones2", "err": err5.Error(), "status": "404"}
			return outputError
		}

		//Se elimina vinculacion nueva
		if err6 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "DELETE", &respuesta_peticion, pos); err6 == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_delete_vin)
			fmt.Println("respuesta_eliminar_vin_nueva", respuesta_delete_vin)
		} else {
			beego.Error(err6)
			outputError = map[string]interface{}{"funcion": "/AnularModificaciones1", "err": err6.Error(), "status": "404"}
			return outputError
		}
	}
	return nil
}

func AnularAdicionDocente(v models.Objeto_Desvinculacion) (outputError map[string]interface{}) {
	var respuesta_vinculacion string
	var vinculacion_cancelada []models.VinculacionDocente
	var respuesta_delete_vin string
	var respuesta_delete string
	var respuesta_modificacion_vinculacion []models.ModificacionVinculacion
	var respuesta_peticion map[string]interface{}

	//Se trae información de tabla de traza modificacion_vinculacion, para saber cuál vinculación hay que poner en true y cuál eliminar
	query := "?limit=-1&query=ModificacionResolucionId.Id:" + strconv.Itoa(v.IdModificacionResolucion) + ",VinculacionDocenteRegistradaId.Id:" + strconv.Itoa(v.DocentesDesvincular[0].Id)
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion"+query, &respuesta_peticion); err == nil && response == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_modificacion_vinculacion)
		fmt.Println("modificacion_vinculacion", respuesta_modificacion_vinculacion)
	} else {
		outputError = map[string]interface{}{"funcion": "/AnularAdicionDocente5", "err": err.Error(), "status": "404"}
		return outputError
	}

	//se trae informacion de vinculacion que fue cancelada
	query2 := "?limit=-1&query=Id:" + strconv.Itoa(respuesta_modificacion_vinculacion[0].VinculacionDocenteCanceladaId.Id)
	if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query2, &respuesta_peticion); err2 == nil && response2 == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &vinculacion_cancelada)
		fmt.Println("vinculacion_cancelada", vinculacion_cancelada)
	} else {
		outputError = map[string]interface{}{"funcion": "/AnularAdicionDocente4", "err": err2.Error(), "status": "404"}
		return outputError
	}
	//se cambia a true vinculacion que fue cancelada
	vinculacion_cancelada[0].Activo = true
	fmt.Println("nuevo estado de vinculacion cancelada", vinculacion_cancelada)

	//Se le cambia estado en bd a vinculacion cancelada

	if err3 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion_cancelada[0].Id), "PUT", &respuesta_peticion, vinculacion_cancelada[0]); err3 == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_vinculacion)
		fmt.Println("respuesta_vinculacion", respuesta_vinculacion)
	} else {
		outputError = map[string]interface{}{"funcion": "/AnularAdicionDocente3", "err": err3.Error(), "status": "404"}
		return outputError
	}

	//se elimina registro en modificacion_vinculacion

	if err4 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/"+strconv.Itoa(respuesta_modificacion_vinculacion[0].Id), "DELETE", &respuesta_peticion, respuesta_modificacion_vinculacion[0]); err4 == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_delete)
	} else {
		outputError = map[string]interface{}{"funcion": "/AnularAdicionDocente2", "err": err4.Error(), "status": "404"}
		return outputError
	}

	//Se elimina vinculacion nueva
	if err5 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.DocentesDesvincular[0].Id), "DELETE", &respuesta_peticion, v.DocentesDesvincular[0]); err5 == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_delete_vin)
		fmt.Println("respuesta_eliminar_vin_nueva", respuesta_delete_vin)
	} else {
		outputError = map[string]interface{}{"funcion": "/AnularAdicionDocente1", "err": err5.Error(), "status": "404"}
		return outputError
	}

	return nil
}

func ConsultarCategoria(v models.VinculacionDocente) (respuesta string, outputError map[string]interface{}) {

	categoria, _, err := Buscar_Categoria_Docente(strconv.Itoa(v.VigenciaCarga), strconv.Itoa(v.PeriodoCarga), strconv.Itoa(v.PersonaId))
	if err != nil {
		return respuesta, err
	}

	respuesta = "OK"
	if categoria == "" {
		respuesta = "Sin categoría"
	}

	return respuesta, nil
}

func ValidarSaldoCDP(validacion models.Objeto_Desvinculacion) (respuesta string, outputError map[string]interface{}) {

	var respuestaApropiacion models.DatosApropiacion
	var saldoDisponibilidad float64

	validacion.DocentesDesvincular[0].NumeroHorasSemanales = validacion.DocentesDesvincular[0].NumeroHorasNuevas
	validacion.DocentesDesvincular[0].NumeroSemanas = validacion.DocentesDesvincular[0].NumeroSemanasNuevas

	if docentes, err := CalcularSalarioPrecontratacion(validacion.DocentesDesvincular); err == nil {
		validacion.DocentesDesvincular = docentes
		respuesta = "OK"
	} else {
		respuesta = "Error"
		return respuesta, err
	}

	valorContrato := validacion.DocentesDesvincular[0].ValorContrato

	if err2 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/SaldoCdp", "POST", &respuestaApropiacion, &validacion.DisponibilidadNueva); err2 == nil {
		saldoDisponibilidad = float64(respuestaApropiacion.Saldo)
		respuesta = "OK"
	} else {
		respuesta = "Error"
		outputError = map[string]interface{}{"funcion": "/ValidarSaldoCDP", "err": err2.Error(), "status": "404"}
		return respuesta, outputError
	}

	if saldoDisponibilidad-valorContrato < 0 {
		respuesta = "Error CDP"
	}
	return respuesta, nil
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
	var respuesta_peticion map[string]interface{}

	//CAMBIAR ESTADO DE VINCULACIÓN DOCENTE
	for _, pos := range v.DocentesDesvincular {
		pos.NumeroRp = 0
		pos.VigenciaRp = 0
		err1 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "PUT", &respuesta_peticion, pos)
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta)
		//TODO: unificar errores
		if err1 != nil {
			logs.Error(err1)
			outputError = map[string]interface{}{"funcion": "/AdicionarHoras", "err1": err1.Error(), "status": "502"}
			return respuesta, outputError
		}
		//TODO: unificar errores
		beego.Info("respuesta", respuesta)
		beego.Info("fechaAD", v.DocentesDesvincular[0].FechaInicioNueva)
		temp_vinculacion[0] = models.VinculacionDocente{
			PersonaId:            v.DocentesDesvincular[0].PersonaId,
			NumeroHorasSemanales: v.DocentesDesvincular[0].NumeroHorasNuevas,
			NumeroSemanas:        v.DocentesDesvincular[0].NumeroSemanasNuevas,
			ResolucionVinculacionDocenteId:         &models.ResolucionVinculacionDocente{Id: v.IdNuevaResolucion},
			DedicacionId:         v.DocentesDesvincular[0].DedicacionId,
			ProyectoCurricularId: v.DocentesDesvincular[0].ProyectoCurricularId,
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
			//outputError = map[string]interface{}{"funcion": "/AdicionarHoras", "err1": err2, "status": "502"}
			return respuesta, err2
		}

		//INSERCION  TABLA  DE TRAZA MODIFICACION VINCULACION
		for _, pos := range v.DocentesDesvincular {
			temp := models.ModificacionVinculacion{
				ModificacionResolucionId:       &models.ModificacionResolucion{Id: v.IdModificacionResolucion},
				VinculacionDocenteCanceladaId:  &models.VinculacionDocente{Id: pos.Id},
				VinculacionDocenteRegistradaId: &models.VinculacionDocente{Id: vinculacion_nueva},
				Horas:                        pos.NumeroHorasNuevas,
			}
			if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/", "POST", &respuesta_peticion, temp); err == nil {
				LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_mod_vin)
				beego.Info("respuesta modificacion vin", respuesta_mod_vin)
				respuesta = "OK"
			} else {
				beego.Error("error en actualizacion de modificacion vinculacion de modificacion vinculacion", err)
				outputError = map[string]interface{}{"funcion": "/AdicionarHoras2", "err": err.Error(), "status": "404"}
				return "Error", outputError
			}

		}
	}

	return respuesta, nil
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
	var respuesta_peticion map[string]interface{}

	beego.Debug("para poner en false", v)

	for _, pos := range v.DocentesDesvincular {

		numerorp := pos.NumeroRp
		vigenciarp := pos.VigenciaRp

		pos.NumeroRp = 0
		pos.VigenciaRp = 0
		err1 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id), "PUT", &respuesta_peticion, pos)
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta)
		if err1 != nil {
			logs.Error(err1)
			outputError = map[string]interface{}{"funcion": "/ActualizarVinculacionesCancelacion1", "err1": err1.Error(), "status": "502"}
			return respuesta, outputError
		}
		//Verificar objeto para crear nuevas resoluciones
		temp_vinculacion[0] = models.VinculacionDocente{
			PersonaId:            pos.PersonaId,
			NumeroHorasSemanales: pos.NumeroHorasSemanales,
			NumeroSemanas:        pos.NumeroSemanasNuevas,
			ResolucionVinculacionDocenteId:         &models.ResolucionVinculacionDocente{Id: v.IdNuevaResolucion},
			DedicacionId:         pos.DedicacionId,
			ProyectoCurricularId: pos.ProyectoCurricularId,
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
			//outputError = map[string]interface{}{"funcion": "/ActualizarVinculacionesCancelacion", "err2": err2, "status": "502"}
			return respuesta, err2
		}

		//INSERCION  TABLA  DE TRAZA MODIFICACION VINCULACION
		temp := models.ModificacionVinculacion{
			ModificacionResolucionId:       &models.ModificacionResolucion{Id: v.IdModificacionResolucion},
			VinculacionDocenteCanceladaId:  &models.VinculacionDocente{Id: pos.Id},
			VinculacionDocenteRegistradaId: &models.VinculacionDocente{Id: vinculacion_nueva},
			Horas:                        pos.NumeroHorasSemanales,
		}
		if errorMod := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/", "POST", &respuesta_peticion, temp); errorMod == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_mod_vin)
			beego.Info("respuesta modificacion vin", respuesta_mod_vin)
			respuesta = "OK"
		} else {
			beego.Error("error en actualizacion de modificacion vinculacion de modificacion vinculacion", errorMod)
			outputError = map[string]interface{}{"funcion": "/ActualizarVinculacionesCancelacion2", "err": errorMod.Error(), "status": "404"}
			return "Error", outputError
		}
	}

	return respuesta, nil

}

func InsertarDesvinculaciones(v [1]models.VinculacionDocente) (id int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var d []models.VinculacionDocente
	var respuesta_peticion map[string]interface{}
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
	d, err := CalcularSalarioPrecontratacion(d)
	if err != nil {
		return id, err
	}

	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &respuesta_peticion, &d); err != nil {
		outputError = map[string]interface{}{"funcion": "/InsertarDesvinculaciones", "err": err.Error(), "status": "404"}
		return 0, outputError
	}else{
		LimpiezaRespuestaRefactor(respuesta_peticion, &id)
	}
	return id, nil
}
