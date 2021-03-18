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
		outputError = map[string]interface{}{"funcion": "/ListarDocentesHorasLectivas1", "err1": err1, "status": "404"}
		return docentes_a_listar, outputError
	}

	var temp map[string]interface{}
	var docentesXCarga models.ObjetoCargaLectiva

	for _, pos := range tipoVinculacionOld {
		t := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAcademica") + "/" + "carga_lectiva/" + vigencia + "/" + periodo + "/" + pos + "/" + facultadOld + "/" + nivel_academico

		err2 := GetJsonWSO2(t, &temp)
		if err2 != nil {
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
	if err := GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudFinanciera")+"/"+"cdprpdocente/"+numero_cdp+"/"+vigencia_cdp+"/"+identificacion, &temp); err == nil {
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

func ListarDocentesDesvinculados(query string) (VinculacionDocente []models.VinculacionDocente) {
	v := []models.VinculacionDocente{}

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v); err == nil && response == 200 {
		for x, pos := range v {
			documento_identidad, _ := strconv.Atoi(pos.IdPersona)
			v[x].NombreCompleto = BuscarNombreProveedor(documento_identidad)
			v[x].NumeroDisponibilidad = BuscarNumeroDisponibilidad(pos.Disponibilidad)
			v[x].Dedicacion = BuscarNombreDedicacion(pos.IdDedicacion.Id)
			v[x].LugarExpedicionCedula = BuscarLugarExpedicion(pos.IdPersona)
		}
		if v == nil {
			v = []models.VinculacionDocente{}
		}
		return v
	} else {
		return nil
	}

	return
}

func ListarDocentesCancelados(id_resolucion string) (VinculacionDocente []models.VinculacionDocente) {
	var v []models.VinculacionDocente
	var modRes []models.ModificacionResolucion
	var modVin []models.ModificacionVinculacion
	var cv models.VinculacionDocente
	// if 3 - modificacion_resolucion
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=resolucionNueva:"+id_resolucion, &modRes); err == nil && response == 200 {
		// if 2 - modificacion_vinculacion
		t := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAdmin") + "/" + beego.AppConfig.String("NscrudAdmin") + "/modificacion_vinculacion/?limit=-1&query=modificacion_resolucion:" + strconv.Itoa(modRes[0].Id)
		beego.Info(t)
		if response, err := GetJsonTest(t, &modVin); err == nil && response == 200 {
			//for vinculaciones
			for _, vinculacion := range modVin {
				beego.Info(fmt.Sprintf("%+v", vinculacion.VinculacionDocenteCancelada))
				// if 1 - vinculacion_docente
				if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacion.VinculacionDocenteCancelada.Id), &cv); err == nil && response == 200 {
					documento_identidad, _ := strconv.Atoi(vinculacion.VinculacionDocenteCancelada.IdPersona)
					cv.NombreCompleto = BuscarNombreProveedor(documento_identidad)
					cv.NumeroDisponibilidad = BuscarNumeroDisponibilidad(vinculacion.VinculacionDocenteCancelada.Disponibilidad)
					cv.Dedicacion = BuscarNombreDedicacion(vinculacion.VinculacionDocenteCancelada.IdDedicacion.Id)
					cv.LugarExpedicionCedula = BuscarLugarExpedicion(vinculacion.VinculacionDocenteCancelada.IdPersona)
					cv.NumeroSemanasNuevas = vinculacion.VinculacionDocenteCancelada.NumeroSemanas - vinculacion.VinculacionDocenteRegistrada.NumeroSemanas
				} else { // if 1 - vinculacion_docente
					fmt.Println("Error de consulta en vinculacion, solucioname!!!, if 1 - vinculacion_docente: ", err)
				}
				v = append(v, cv)
			} //fin for vinculaciones
			return v
		} else { // if 2 - modificacion_vinculacion
			fmt.Println("Error de consulta en modificacion_vinculacion, solucioname!!!, if 2 - modificacion_vinculacion: ", err)
			v = []models.VinculacionDocente{}
		}
	} else { // if 3 - modificacion_resolucion
		fmt.Println("Error de consulta en modificacion_resolucion, solucioname!!!, if 3 - modificacion_resolucion: ", err)
		return nil
	}
	return
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
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas1", "err1": err1.Error(), "status": "404"}
		return newDocentesXcargaHoraria, outputError
	}

	//BUSCAR CATEGORÍA DE CADA DOCENTE
	for _, pos := range docentesXcargaHoraria.CargasLectivas.CargaLectiva {
		catDocente := models.ObjetoCategoriaDocente{}
		emptyCatDocente := models.ObjetoCategoriaDocente{}
		//TODO: quitar el hardconding para WSO2 cuando ya soporte https:
		q := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudUrano") + "/categoria_docente/" + vigencia + "/" + periodo + "/" + pos.DocDocente
		err2 := GetXml(q, &catDocente.CategoriaDocente)
		if err2 != nil {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas2", "err2": err2, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		var err3 error
		pos.CategoriaNombre, pos.IDCategoria, err3 = Buscar_Categoria_Docente(vigencia, periodo, pos.DocDocente)
		if err3 != nil {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas3", "err3": err3, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		if catDocente.CategoriaDocente != emptyCatDocente.CategoriaDocente {
			newDocentesXcargaHoraria.CargasLectivas.CargaLectiva = append(newDocentesXcargaHoraria.CargasLectivas.CargaLectiva, pos)
		}
	}

	//RETORNAR CON ID DE TIPO DE VINCULACION DE NUEVO MODELO
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		var err4 map[string]interface{}
		pos.IDTipoVinculacion, pos.NombreTipoVinculacion, err4 = HomologarDedicacion_ID("old", pos.IDTipoVinculacion)
		if err4 != nil {
			logs.Error(err4)
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas4", "err4": err4, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
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
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas5", "err5": err5, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}
	//RETORNAR PROYECTOS CURRICUALRES HOMOLOGADOS!!
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		var err6 error
		pos.DependenciaAcademica, err6 = strconv.Atoi(pos.IDProyecto)
		if err6 != nil {
			logs.Error(err6)
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas6", "err6": err6.Error(), "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		var err7 map[string]interface{}
		pos.IDProyecto, err7 = HomologarProyectoCurricular(pos.IDProyecto)
		if err7 != nil {
			logs.Error(err7)
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas7", "err7": err7, "status": "404"}
			return newDocentesXcargaHoraria, outputError
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}

	return
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
	//Devuelve el nivel académico, la dedicación y la facultad de la resolución
	if response, err1 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente/"+idResolucion, &resvinc); err1 != nil && response != 200 {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll1", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	}

	//Devuelve la información básica de la resolución que se está consultando
	if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucion, &res); err2 != nil && response2 != 200 {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll2", "err2": err2.Error(), "status": "502"}
		return nil, outputError
	}

	if res.IdTipoResolucion.Id != tipoVinculacion {
		//Busca el id de la modificación donde se relacionan la resolución original y la de modificación asociada
		if response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres); err3 != nil && response3 != 200 {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll3", "err3": err3.Error(), "status": "502"}
			return nil, outputError
		}
	}

	//Devuelve las vinculaciones presentes en la resolución consultada, agrupadas o no, según el nivel académico
	if resvinc.NivelAcademico == "POSGRADO" {
		if response4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente?limit=-1&query=IdResolucion.Id:"+idResolucion, &vinc); err4 != nil && response4 != 200 {
			logs.Error(err4)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll4", "err4": err4.Error(), "status": "502"}
			return nil, outputError
		}
	}
	if resvinc.NivelAcademico == "PREGRADO" {
		if response5, err5 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_vinculaciones_agrupadas/"+idResolucion, &vinc); err5 != nil && response5 != 200 {
			logs.Error(err5)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll5", "err5": err5.Error(), "status": "502"}
			return nil, outputError
		}
	}
	var err6 map[string]interface{}
	var llenarVinculacion = func(v *models.VinculacionDocente) {
		documentoIdentidad, _ := strconv.Atoi(v.IdPersona)
		v.NombreCompleto, err6 = BuscarNombreProveedor(documentoIdentidad)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.Dedicacion, err6 = BuscarNombreDedicacion(v.IdDedicacion.Id)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.LugarExpedicionCedula, err6 = BuscarLugarExpedicion(v.IdPersona)
		if err6 != nil {
			logs.Error(err6)
			panic(err6)
		}
		v.TipoDocumento, err6 = BuscarTipoDocumento(v.IdPersona)
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
	switch res.IdTipoResolucion.Id {
	case tipoVinculacion:
		for x, pos := range vinc {
			v = append(v, pos)
			llenarVinculacion(&pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err7 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(pos.IdResolucion.Id), pos.IdDedicacion.Id)
				if err7 != nil {
					logs.Error(err7)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll7", "err7": err7, "status": "502"}
					return nil, outputError
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
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas, err7 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				if err8 != nil {
					logs.Error(err8)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll8", "err8": err8, "status": "502"}
					return nil, outputError
				}
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)

				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err9 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
				if err9 != nil {
					logs.Error(err9)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll9", "err9": err9, "status": "502"}
					return nil, outputError
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
				response5, err10 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err10 != nil && response5 != 200 {
					logs.Error(err10)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll10", "err10": err10, "status": "502"}
					return nil, outputError
				}
				var vincOriginal = modvin[0].VinculacionDocenteCancelada
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
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas, err11 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				if err11 != nil {
					logs.Error(err11)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll11", "err11": err11, "status": "502"}
					return nil, outputError
				}
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err12 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
				if err12 != nil {
					logs.Error(err12)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll12", "err12": err12, "status": "502"}
					return nil, outputError
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
				response6, err13 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id)+",VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err13 != nil && response6 != 200 {
					logs.Error(err13)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll13", "err13": err13, "status": "502"}
					return nil, outputError
				}
				var vincOriginal = modvin[0].VinculacionDocenteCancelada
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
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas, err14 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				if err14 != nil {
					logs.Error(err14)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll14", "err14": err14, "status": "502"}
					return nil, outputError
				}
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas, err15 = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
				if err15 != nil {
					logs.Error(err15)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll15", "err15": err15, "status": "502"}
					return nil, outputError
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
				response7, err16 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err16 != nil && response7 != 200 {
					logs.Error(err16)
					outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculadosAll16", "err16": err16, "status": "502"}
					return nil, outputError
				}
				var vincOriginal = modvin[0].VinculacionDocenteCancelada
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

	return v, outputError
}

func ListarDocentesPrevinculados(idResolucion string, tipoVinculacion int) (v []models.VinculacionDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	query := "?limit=-1&query=IdResolucion.Id:" + idResolucion + ",Estado:true"
	var res models.Resolucion
	var modres []models.ModificacionResolucion
	var modvin []models.ModificacionVinculacion

	response1, err1 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucion, &res)
	if err1 != nil && response1 != 200 {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados1", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	}

	if res.IdTipoResolucion.Id == tipoVinculacion {
		response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v)
		if err2 != nil && response2 != 200 {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados2", "err2": err2.Error(), "status": "502"}
			return nil, outputError
		}
	} else {
		response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres)
		if err2 != nil && response2 != 200 {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados2", "err2": err2.Error(), "status": "502"}
			return nil, outputError
		}
		response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err3 != nil && response3 != 200 {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados3", "err3": err3.Error(), "status": "502"}
			return nil, outputError
		}
		if len(modvin) != 0 {
			arreglo := make([]string, len(modvin))
			for x, pos := range modvin {
				arreglo[x] = strconv.Itoa(pos.VinculacionDocenteRegistrada.Id)
			}
			identificadoresvinc := strings.Join(arreglo, "|")
			response4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=Estado:True,Id__in:"+identificadoresvinc+"&limit=-1", &v)
			if err4 != nil && response4 != 200 {
				logs.Error(err4)
				outputError = map[string]interface{}{"funcion": "/ListarDocentesPrevinculados4", "err4": err4.Error(), "status": "502"}
				return nil, outputError
			}
		} else {
			v = nil
		}
	}
	var err5, err6, err7, err8, err9, err10 map[string]interface{}
	for x, pos := range v {
		documentoIdentidad, _ := strconv.Atoi(pos.IdPersona)

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
		pos.Dedicacion, err7 = BuscarNombreDedicacion(pos.IdDedicacion.Id)
		if err7 != nil {
			logs.Error(err7)
			panic(err7)
		}
		pos.LugarExpedicionCedula, err8 = BuscarLugarExpedicion(pos.IdPersona)
		if err8 != nil {
			logs.Error(err8)
			panic(err8)
		}
		pos.TipoDocumento, err9 = BuscarTipoDocumento(pos.IdPersona)
		if err9 != nil {
			logs.Error(err9)
			panic(err9)
		}
		pos.ValorContratoFormato = FormatMoney(int(v[x].ValorContrato), 2)
		pos.ProyectoNombre, err10 = BuscarNombreFacultad(int(v[x].IdProyectoCurricular))
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
	if response1, err1 := GetJsonTest("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+num_vinculacion+",Vigencia:"+vigencia, &contratoDisponibilidad); err1 == nil && response1 == 200 { //If 2  (get)
		//for contrato_disponibilidad
		for _, pos := range contratoDisponibilidad {
			rpdocente = GetInformacionRpDocente(strconv.Itoa(pos.NumeroCdp), strconv.Itoa(pos.VigenciaCdp), identificacion)
			return rpdocente, nil
		}
	} else { //If 1 contrato_disponibilidad (get)
		panic(err1)
	}
	return
}
