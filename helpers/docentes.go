package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func ListarDocentesHorasLectivas(vigencia, periodo, tipo_vinculacion, facultad, nivel_academico string) (docentes_a_listar models.ObjetoCargaLectiva, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesHorasLectivas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	tipoVinculacionOld := HomologarDedicacion_nombre(tipo_vinculacion)
	facultadOld, err1 := HomologarFacultad("new", facultad)
	if err1 != nil {
		logs.Error(err1)
		//outputError = map[string]interface{}{"funcion": "/ListarDocentesHorasLectivas1", "err1": err1, "status": "404"}
		return docentes_a_listar, err1
	}

	var temp map[string]interface{}
	var docentesXCarga models.ObjetoCargaLectiva

	for _, pos := range tipoVinculacionOld {
		t := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAcademica") + "/carga_lectiva/" + vigencia + "/" + periodo + "/" + pos + "/" + facultadOld + "/" + nivel_academico
		fmt.Println(t)
		if response, err2 := GetJsonWSO2Test(t, &temp); response == 200 && err2 == nil {

		} else {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesHorasLectivas2", "err2": err2.Error(), "status": "404"}
			return docentesXCarga, outputError

		}
		jsonDocentes, err3 := json.Marshal(temp)
		if err3 != nil {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesHorasLectivas3", "err3": err3.Error(), "status": "404"}
			return docentesXCarga, outputError
		}

		var tempDocentes models.ObjetoCargaLectiva
		err4 := json.Unmarshal(jsonDocentes, &tempDocentes)
		if err4 != nil {
			logs.Error(err4)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesHorasLectivas4", "err4": err4.Error(), "status": "404"}
			return docentesXCarga, outputError
		}
		docentesXCarga.CargasLectivas.CargaLectiva = append(docentesXCarga.CargasLectivas.CargaLectiva, tempDocentes.CargasLectivas.CargaLectiva...)

	}

	return docentesXCarga, nil

}

func GetInformacionRpDocente(numero_cdp string, vigencia_cdp string, identificacion string) (informacion_rp_docente models.RpDocente) {

	var temp map[string]interface{}
	fmt.Println(numero_cdp + " " + vigencia_cdp + " " + identificacion)
	if err := GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudFinanciera")+"/cdprpdocente/"+numero_cdp+"/"+vigencia_cdp+"/"+identificacion, &temp); err == nil {
		json_cdp_rp, error_json := json.Marshal(temp)

		if error_json == nil {
			var rp_docente models.RpDocente
			err = json.Unmarshal(json_cdp_rp, &rp_docente)
			if err != nil {
				fmt.Println(err)
			}
			informacion_rp_docente = rp_docente
			fmt.Println(informacion_rp_docente)
			return informacion_rp_docente
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return informacion_rp_docente
}

func ListarDocentesDesvinculados(query string) (VinculacionDocente []models.VinculacionDocente, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesDesvinculados", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	v := []models.VinculacionDocente{}
	var respuesta_peticion map[string]interface{}

	var err1 map[string]interface{}
	var err2 map[string]interface{}
	var err3 map[string]interface{}
	var err4 map[string]interface{}

	r := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/vinculacion_docente" + query
	if response, err := GetJsonTest(r, &respuesta_peticion); err == nil && response == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &v)
		for x, pos := range v {
			documento_identidad := pos.PersonaId
			v[x].NombreCompleto, err1 = BuscarNombreProveedor(documento_identidad)
			if err1 != nil {
				return v, err1
			}
			v[x].NumeroDisponibilidad, err2 = BuscarNumeroDisponibilidad(pos.Disponibilidad)
			if err2 != nil {
				return v, err2
			}
			v[x].Dedicacion, err3 = BuscarNombreDedicacion(pos.DedicacionId.Id)
			if err3 != nil {
				return v, err3
			}
			v[x].LugarExpedicionCedula, err4 = BuscarLugarExpedicion(strconv.Itoa(pos.PersonaId))
			if err4 != nil {
				return v, err4
			}
		}
		if v == nil {
			v = []models.VinculacionDocente{}
		}
		return v, nil
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesDesvinculados", "err": err.Error(), "status": "404"}
		return nil, outputError
	}
}

func ListarDocentesCancelados(id_resolucion string) (VinculacionDocente []models.VinculacionDocente, outputError map[string]interface{}) {
	var v []models.VinculacionDocente
	var modRes []models.ModificacionResolucion
	var modVin []models.ModificacionVinculacion
	var cv models.VinculacionDocente
	var respuesta_peticion map[string]interface{}
	// if 3 - modificacion_resolucion

	var err1 map[string]interface{}
	var err2 map[string]interface{}
	var err3 map[string]interface{}
	var err4 map[string]interface{}

	r := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/modificacion_resolucion?query=ResolucionNuevaId.Id:" + id_resolucion
	if response, err := GetJsonTest(r, &respuesta_peticion); err == nil && response == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &modRes)
		// if 2 - modificacion_vinculacion
		t := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/modificacion_vinculacion?limit=-1&query=ModificacionResolucionId.Id:" + strconv.Itoa(modRes[0].Id)
		beego.Info(t)
		if response, err := GetJsonTest(t, &respuesta_peticion); err == nil && response == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &modVin)
			//for vinculaciones
			for _, vinculacion := range modVin {
				beego.Info(fmt.Sprintf("%+v", vinculacion.VinculacionDocenteCanceladaId))
				// if 1 - vinculacion_docente
				s := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/vinculacion_docente/" + strconv.Itoa(vinculacion.VinculacionDocenteCanceladaId.Id)
				if response, err := GetJsonTest(s, &respuesta_peticion); err == nil && response == 200 {
					LimpiezaRespuestaRefactor(respuesta_peticion, &cv)
					documento_identidad := vinculacion.VinculacionDocenteCanceladaId.PersonaId
					cv.NombreCompleto, err1 = BuscarNombreProveedor(documento_identidad)
					if err1 != nil {
						return v, err1
					}
					cv.NumeroDisponibilidad, err2 = BuscarNumeroDisponibilidad(vinculacion.VinculacionDocenteCanceladaId.Disponibilidad)
					if err2 != nil {
						return v, err2
					}
					cv.Dedicacion, err3 = BuscarNombreDedicacion(vinculacion.VinculacionDocenteCanceladaId.DedicacionId.Id)
					if err3 != nil {
						return v, err3
					}
					cv.LugarExpedicionCedula, err4 = BuscarLugarExpedicion(strconv.Itoa(vinculacion.VinculacionDocenteCanceladaId.PersonaId))
					if err4 != nil {
						return v, err4
					}
					cv.NumeroSemanasNuevas = vinculacion.VinculacionDocenteCanceladaId.NumeroSemanas - vinculacion.VinculacionDocenteRegistradaId.NumeroSemanas
				} else { // if 1 - vinculacion_docente
					fmt.Println("Error de consulta en vinculacion, solucioname!!!, if 1 - vinculacion_docente: ", err)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesCancelados1", "err": err.Error(), "status": "404"}
					return nil, outputError
				}
				v = append(v, cv)
			} //fin for vinculaciones
			return v, nil
		} else { // if 2 - modificacion_vinculacion
			fmt.Println("Error de consulta en modificacion_vinculacion, solucioname!!!, if 2 - modificacion_vinculacion: ", err)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesCancelados2", "err": err.Error(), "status": "404"}
			return nil, outputError
		}
	} else { // if 3 - modificacion_resolucion
		fmt.Println("Error de consulta en modificacion_resolucion, solucioname!!!, if 3 - modificacion_resolucion: ", err)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesCancelados3", "err": err.Error(), "status": "404"}
		return nil, outputError
	}
}

func ListarDocentesCargaHoraria(vigencia string, periodo string, tipoVinculacion string, facultad string, nivelAcademico string) (newDocentesXcargaHoraria models.ObjetoCargaLectiva, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesCargaHoraria", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	docentesXcargaHoraria, err1 := ListarDocentesHorasLectivas(vigencia, periodo, tipoVinculacion, facultad, nivelAcademico)
	if err1 != nil {
		logs.Error(err1)
		//outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas1", "err1": err1, "status": "404"}
		return newDocentesXcargaHoraria, err1
	}

	//BUSCAR CATEGORÍA DE CADA DOCENTE
	for _, pos := range docentesXcargaHoraria.CargasLectivas.CargaLectiva {
		catDocente := models.ObjetoCategoriaDocente{}
		emptyCatDocente := models.ObjetoCategoriaDocente{}
		//TODO: quitar el hardconding para WSO2 cuando ya soporte https:
		q := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudUrano") + "/categoria_docente/" + vigencia + "/" + periodo + "/" + pos.DocDocente
		response, err2 := GetJsonWSO2Test(q, &catDocente.CategoriaDocente)
		if err2 != nil && response == 200 {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesCargaHoraria", "err2": err2, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		var err1 map[string]interface{}
		pos.CategoriaNombre, pos.IDCategoria, err1 = Buscar_Categoria_Docente(vigencia, periodo, pos.DocDocente)
		fmt.Println(err1)
		if err1 != nil {
			beego.Error(err1)
			return newDocentesXcargaHoraria, outputError
		}
		if catDocente.CategoriaDocente != emptyCatDocente.CategoriaDocente {
			newDocentesXcargaHoraria.CargasLectivas.CargaLectiva = append(newDocentesXcargaHoraria.CargasLectivas.CargaLectiva, pos)
		}
	}

	//RETORNAR CON ID DE TIPO DE VINCULACION DE NUEVO MODELO
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.IDTipoVinculacion, pos.NombreTipoVinculacion = HomologarDedicacion_ID("old", pos.IDTipoVinculacion)
		if pos.IDTipoVinculacion == "3" {
			pos.HorasLectivas = "20"
			pos.NombreTipoVinculacion = "MTO"
		}
		if pos.IDTipoVinculacion == "4" {
			pos.HorasLectivas = "40"
			pos.NombreTipoVinculacion = "TCO"
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}

	//RETORNAR FACULTTADES CON ID DE OIKOS, HOMOLOGACION
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		var err5 map[string]interface{}
		pos.IDFacultad, err5 = HomologarFacultad("old", pos.IDFacultad)
		if err5 != nil {
			logs.Error(err5)
			//outputError = map[string]interface{}{"funcion": "/ListarDocentesCargaHoraria5", "err5": err5, "status": "404"}
			return newDocentesXcargaHoraria, err5
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}
	//RETORNAR PROYECTOS CURRICUALRES HOMOLOGADOS!!
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		var err6 error
		pos.DependenciaAcademica, err6 = strconv.Atoi(pos.IDProyecto)
		if err6 != nil {
			logs.Error(err6)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesCargaHoraria6", "err6": err6.Error(), "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		var err7 map[string]interface{}
		pos.IDProyecto, err7 = HomologarProyectoCurricular(pos.IDProyecto)
		if err7 != nil {
			logs.Error(err7)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesCargaHoraria7", "err7": err7, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}

	return newDocentesXcargaHoraria, nil
}

func ListarDocentesPrevinculadosAll(idResolucion string, tipoVinculacion int, tipoCancelacion int, tipoAdicion int, tipoReduccion int) (v []models.VinculacionDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var res models.Resolucion
	var resvinc models.ResolucionVinculacionDocente
	var modres []models.ModificacionResolucion
	var vinc []models.VinculacionDocente
	var modvin []models.ModificacionVinculacion
	var ValorModificacionContrato float64
	var respuesta_peticion map[string]interface{}

	//Devuelve el nivel académico, la dedicación y la facultad de la resolución
	if response, err1 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion_vinculacion_docente/"+idResolucion, &respuesta_peticion); err1 != nil && response != 200 {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll1", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	} else {
		LimpiezaRespuestaRefactor(respuesta_peticion, &resvinc)
	}

	//Devuelve la información básica de la resolución que se está consultando
	if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+idResolucion, &respuesta_peticion); err2 != nil && response2 != 200 {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll2", "err2": err2.Error(), "status": "502"}
		return nil, outputError
	} else {
		LimpiezaRespuestaRefactor(respuesta_peticion, &res)
	}

	if res.TipoResolucionId.Id != tipoVinculacion {
		//Busca el id de la modificación donde se relacionan la resolución original y la de modificación asociada
		if response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/modificacion_resolucion?query=ResolucionNuevaId.Id:"+idResolucion, &respuesta_peticion); err3 != nil && response3 != 200 {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll3", "err3": err3.Error(), "status": "502"}
			return nil, outputError
		} else {
			LimpiezaRespuestaRefactor(respuesta_peticion, &modres)
		}
	}

	//Devuelve las vinculaciones presentes en la resolución consultada, agrupadas o no, según el nivel académico
	if resvinc.NivelAcademico == "POSGRADO" {
		if response4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente?limit=-1&query=ResolucionVinculacionDocenteId.Id:"+idResolucion, &respuesta_peticion); err4 != nil && response4 != 200 {
			logs.Error(err4)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll4", "err4": err4.Error(), "status": "502"}
			return nil, outputError
		} else {
			LimpiezaRespuestaRefactor(respuesta_peticion, &vinc)
		}
	}
	if resvinc.NivelAcademico == "PREGRADO" {
		if response5, err5 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente/get_vinculaciones_agrupadas/"+idResolucion, &respuesta_peticion); err5 != nil && response5 != 200 {
			logs.Error(err5)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll5", "err5": err5.Error(), "status": "502"}
			return nil, outputError
		} else {
			LimpiezaRespuestaRefactor(respuesta_peticion, &vinc)
		}
	}
	var err6 map[string]interface{}
	var llenarVinculacion = func(v *models.VinculacionDocente) {
		documentoIdentidad := v.PersonaId
		v.NombreCompleto, err6 = BuscarNombreProveedor(documentoIdentidad)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.Dedicacion, err6 = BuscarNombreDedicacion(v.DedicacionId.Id)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.LugarExpedicionCedula, err6 = BuscarLugarExpedicion(strconv.Itoa(v.PersonaId))
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.TipoDocumento, err6 = BuscarTipoDocumento(strconv.Itoa(v.PersonaId))
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.NumeroDisponibilidad, err6 = BuscarNumeroDisponibilidad(v.Disponibilidad)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
	}

	var err7, err8, err9, err11, err12, err14, err15 map[string]interface{}
	switch res.TipoResolucionId.Id {
	case tipoVinculacion:
		for x, pos := range vinc {
			v = append(v, pos)
			llenarVinculacion(&pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err7 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), strconv.Itoa(pos.ResolucionVinculacionDocenteId.Id), pos.DedicacionId.Id)
				if err7 != nil {
					logs.Error(err7)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll7", "err7": err7, "status": "502"}
					return nil, err7
				}
			}
			pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
			pos.ValorContratoFormato = FormatMoney(int(pos.ValorContrato), 2)
			v[x] = pos
		}
		break
	case tipoCancelacion:
		for x, pos := range vinc {
			v = append(v, pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				//Agrupa las modificaciones en la resolución por persona
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas, err8 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), idResolucion, pos.DedicacionId.Id)
				if err8 != nil {
					logs.Error(err8)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll8", "err8": err8, "status": "502"}
					return nil, err8
				}
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)

				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err9 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), strconv.Itoa(modres[0].ResolucionAnteriorId.Id), pos.DedicacionId.Id)
				if err9 != nil {
					logs.Error(err9)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll9", "err9": err9, "status": "502"}
					return nil, err9
				}
				pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanas-pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = FormatMoney(int(pos.ValorContrato), 2)
				pos.ValorContratoFormato = FormatMoney(int(pos.ValorContrato-ValorModificacionContrato), 2)
			}
			if resvinc.NivelAcademico == "POSGRADO" {
				pos.NumeroHorasModificacion = pos.NumeroHorasSemanales
				pos.ValorModificacionFormato = FormatMoney(int(pos.ValorContrato), 2)
				ValorModificacionContrato = pos.ValorContrato

				//Busca la vinculación original a la que está asociada la modificación
				response5, err10 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+
					"/modificacion_vinculacion?query=VinculacionDocenteRegistradaId.Id:"+strconv.Itoa(pos.Id), &respuesta_peticion)
				if err10 != nil && response5 != 200 {
					logs.Error(err10)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll10", "err10": err10.Error(), "status": "502"}
					return nil, outputError
				} else {
					LimpiezaRespuestaRefactor(respuesta_peticion, &modvin)
				}
				var vincOriginal = modvin[0].VinculacionDocenteCanceladaId
				pos.NumeroHorasSemanales = vincOriginal.NumeroHorasSemanales
				pos.NumeroMeses = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas-pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = FormatMoney(int(vincOriginal.ValorContrato), 2)
				pos.ValorContratoFormato = FormatMoney(int(vincOriginal.ValorContrato-ValorModificacionContrato), 2)
			}
			llenarVinculacion(&pos)
			v[x] = pos
		}
		break
	case tipoAdicion:
		for x, pos := range vinc {
			v = append(v, pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				//Agrupa las modificaciones en la resolución por persona
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas, err11 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), idResolucion, pos.DedicacionId.Id)
				if err11 != nil {
					logs.Error(err11)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll11", "err11": err11, "status": "502"}
					return nil, err11
				}
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err12 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), strconv.Itoa(modres[0].ResolucionAnteriorId.Id), pos.DedicacionId.Id)
				if err12 != nil {
					logs.Error(err12)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll12", "err12": err12, "status": "502"}
					return nil, err12
				}
				pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroHorasNuevas = pos.NumeroHorasSemanales + pos.NumeroHorasModificacion
				pos.ValorContratoFormato = FormatMoney(int(pos.ValorContrato+ValorModificacionContrato), 2)
			}
			if resvinc.NivelAcademico == "POSGRADO" {
				pos.NumeroHorasModificacion = pos.NumeroHorasSemanales
				pos.ValorModificacionFormato = FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				ValorModificacionContrato = pos.ValorContrato
				//Busca la vinculación original a la que está asociada la modificación
				response6, err13 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+
					"/modificacion_vinculacion?query=ModificacionResolucionId.Id:"+strconv.Itoa(modres[0].Id)+",VinculacionDocenteRegistradaId.Id:"+strconv.Itoa(pos.Id), &respuesta_peticion)
				if err13 != nil && response6 != 200 {
					logs.Error(err13)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll13", "err13": err13.Error(), "status": "502"}
					return nil, outputError
				} else {
					LimpiezaRespuestaRefactor(respuesta_peticion, &modvin)
				}
				var vincOriginal = modvin[0].VinculacionDocenteCanceladaId
				pos.NumeroHorasSemanales = vincOriginal.NumeroHorasSemanales
				pos.NumeroMeses = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = FormatMoney(int(vincOriginal.ValorContrato), 2)
				pos.NumeroHorasNuevas = vincOriginal.NumeroHorasSemanales + pos.NumeroHorasModificacion
				pos.ValorContratoFormato = FormatMoney(int(vincOriginal.ValorContrato+ValorModificacionContrato), 2)
			}
			llenarVinculacion(&pos)
			v[x] = pos
		}
		break
	case tipoReduccion:
		for x, pos := range vinc {
			v = append(v, pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				//Agrupa las modificaciones en la resolución por persona
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas, err14 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), idResolucion, pos.DedicacionId.Id)
				if err14 != nil {
					logs.Error(err14)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll14", "err14": err14, "status": "502"}
					return nil, err14
				}
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err15 = Calcular_totales_vinculacion_pdf_nueva(strconv.Itoa(pos.PersonaId), strconv.Itoa(modres[0].ResolucionAnteriorId.Id), pos.DedicacionId.Id)
				if err15 != nil {
					logs.Error(err15)
					//outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll15", "err15": err15, "status": "502"}
					return nil, err15
				}
				pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroHorasNuevas = pos.NumeroHorasSemanales - pos.NumeroHorasModificacion
				pos.ValorContratoFormato = FormatMoney(int(pos.ValorContrato-ValorModificacionContrato), 2)

			}
			if resvinc.NivelAcademico == "POSGRADO" {
				pos.NumeroHorasModificacion = pos.NumeroHorasSemanales
				pos.ValorModificacionFormato = FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				ValorModificacionContrato = pos.ValorContrato
				//Busca la vinculación original a la que está asociada la modificación
				response7, err16 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+
					"/modificacion_vinculacion?query=VinculacionDocenteRegistradaId.Id:"+strconv.Itoa(pos.Id), &respuesta_peticion)
				if err16 != nil && response7 != 200 {
					logs.Error(err16)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll16", "err16": err16.Error(), "status": "502"}
					return nil, outputError
				} else {
					LimpiezaRespuestaRefactor(respuesta_peticion, &modvin)
				}
				var vincOriginal = modvin[0].VinculacionDocenteCanceladaId
				pos.NumeroHorasSemanales = vincOriginal.NumeroHorasSemanales
				pos.NumeroMeses = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = FormatMoney(int(vincOriginal.ValorContrato), 2)
				pos.NumeroHorasNuevas = vincOriginal.NumeroHorasSemanales - pos.NumeroHorasModificacion
				pos.ValorContratoFormato = FormatMoney(int(vincOriginal.ValorContrato-ValorModificacionContrato), 2)
			}

			llenarVinculacion(&pos)
			v[x] = pos
		}
		break
	default:
		break
	}

	return v, nil
}

func ListarDocentesPrevinculados(idResolucion string, tipoVinculacion int) (v []models.VinculacionDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	query := "?limit=-1&query=ResolucionVinculacionDocenteId.Id:" + idResolucion + ",Activo:true"
	var res models.Resolucion
	var modres []models.ModificacionResolucion
	var modvin []models.ModificacionVinculacion
	var respuesta_peticion map[string]interface{}

	response1, err1 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+idResolucion, &respuesta_peticion)
	if err1 != nil && response1 != 200 {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados1", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	} else {
		LimpiezaRespuestaRefactor(respuesta_peticion, &res)
	}

	if res.TipoResolucionId.Id == tipoVinculacion {
		response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente"+query, &respuesta_peticion)
		if err2 != nil && response2 != 200 {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados2", "err2": err2.Error(), "status": "502"}
			return nil, outputError
		} else {
			LimpiezaRespuestaRefactor(respuesta_peticion, &v)
		}
	} else {
		response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/modificacion_resolucion?query=ResolucionNuevaId.Id:"+idResolucion, &respuesta_peticion)
		if err2 != nil && response2 != 200 {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados2.1", "err2": err2.Error(), "status": "502"}
			return nil, outputError
		} else {
			LimpiezaRespuestaRefactor(respuesta_peticion, &modres)
		}
		response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/modificacion_vinculacion?query=ModificacionResolucionId.Id:"+strconv.Itoa(modres[0].Id), &respuesta_peticion)
		if err3 != nil && response3 != 200 {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados3", "err3": err3.Error(), "status": "502"}
			return nil, outputError
		} else {
			LimpiezaRespuestaRefactor(respuesta_peticion, &modvin)
		}
		if len(modvin) != 0 {
			arreglo := make([]string, len(modvin))
			for x, pos := range modvin {
				arreglo[x] = strconv.Itoa(pos.VinculacionDocenteRegistradaId.Id)
			}
			identificadoresvinc := strings.Join(arreglo, "|")
			response4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente?query=Activo:True,Id__in:"+identificadoresvinc+"&limit=-1", &respuesta_peticion)
			if err4 != nil && response4 != 200 {
				logs.Error(err4)
				outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados4", "err4": err4.Error(), "status": "502"}
				return nil, outputError
			} else {
				LimpiezaRespuestaRefactor(respuesta_peticion, &v)
			}
		} else {
			v = nil
		}
	}
	var err5, err6, err7, err8, err9, err10 map[string]interface{}
	for x, pos := range v {
		documentoIdentidad := pos.PersonaId

		pos.NombreCompleto, err5 = BuscarNombreProveedor(documentoIdentidad)
		if err5 != nil {
			logs.Error(err5)
			panic(err5)
		}
		pos.NumeroDisponibilidad, err6 = BuscarNumeroDisponibilidad(pos.Disponibilidad)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		pos.Dedicacion, err7 = BuscarNombreDedicacion(pos.DedicacionId.Id)
		if err7 != nil {
			logs.Error(err7)
			panic(err7)
		}
		pos.LugarExpedicionCedula, err8 = BuscarLugarExpedicion(strconv.Itoa(pos.PersonaId))
		if err8 != nil {
			logs.Error(err8)
			panic(err8)
		}
		pos.TipoDocumento, err9 = BuscarTipoDocumento(strconv.Itoa(pos.PersonaId))
		if err9 != nil {
			logs.Error(err9)
			panic(err9)
		}
		pos.ValorContratoFormato = FormatMoney(int(v[x].ValorContrato), 2)
		pos.ProyectoNombre, err10 = BuscarNombreFacultad(int(v[x].ProyectoCurricularId))
		if err10 != nil {
			logs.Error(err10)
			panic(err10)
		}
		pos.Periodo = res.Periodo
		pos.VigenciaCarga = res.VigenciaCarga
		pos.PeriodoCarga = res.PeriodoCarga
		v[x] = pos
	}
	if v == nil {
		v = []models.VinculacionDocente{}
	}

	return v, outputError
}

func GetCdpRpDocente(identificacion string, num_vinculacion string, vigencia string) (rpdocente models.RpDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetCdpRpDocente", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var contratoDisponibilidad []models.ContratoDisponibilidad
	//If 1 contrato_disponibilidad (get)
	if response1, err1 := GetJsonTest("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad?query=NumeroContrato:"+num_vinculacion+",Vigencia:"+vigencia, &contratoDisponibilidad); err1 == nil && response1 == 200 { //If 2  (get)
		//for contrato_disponibilidad
		for _, pos := range contratoDisponibilidad {
			rpdocente = GetInformacionRpDocente(strconv.Itoa(pos.NumeroCdp), strconv.Itoa(pos.VigenciaCdp), identificacion)
			return rpdocente, nil
		}
	} else { //If 1 contrato_disponibilidad (get)
		panic(err1)
	}
	return rpdocente, nil
}
