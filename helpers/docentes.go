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

func ListarDocentesHorasLectivas(vigencia, periodo, tipo_vinculacion, facultad, nivel_academico string) (docentes_a_listar models.ObjetoCargaLectiva, err error) {

	tipoVinculacionOld := HomologarDedicacion_nombre(tipo_vinculacion)
	facultadOld, err := HomologarFacultad("new", facultad)
	if err != nil {
		return docentes_a_listar, err
	}

	var temp map[string]interface{}
	var docentesXCarga models.ObjetoCargaLectiva

	for _, pos := range tipoVinculacionOld {
		t := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAcademica") + "/" + "carga_lectiva/" + vigencia + "/" + periodo + "/" + pos + "/" + facultadOld + "/" + nivel_academico

		err = GetJsonWSO2(t, &temp)
		if err != nil {
			return docentesXCarga, err
		}
		jsonDocentes, err := json.Marshal(temp)
		if err != nil {
			return docentesXCarga, err
		}

		var tempDocentes models.ObjetoCargaLectiva
		err = json.Unmarshal(jsonDocentes, &tempDocentes)
		if err != nil {
			return docentesXCarga, err
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

func ListarDocentesDesvinculados(query string) (VinculacionDocente []models.VinculacionDocente, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDocentesDesvinculados", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	v := []models.VinculacionDocente{}

	var err1 map[string]interface{}
	var err2 map[string]interface{}
	var err3 map[string]interface{}
	var err4 map[string]interface{}

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v); err == nil && response == 200 {
		for x, pos := range v {
			documento_identidad, _ := strconv.Atoi(pos.IdPersona)
			v[x].NombreCompleto, err1 = BuscarNombreProveedor(documento_identidad)
			if err1 != nil{
				return v, err1
			}
			v[x].NumeroDisponibilidad, err2 = BuscarNumeroDisponibilidad(pos.Disponibilidad)
			if err2 != nil{
				return v, err2
			}
			v[x].Dedicacion, err3 = BuscarNombreDedicacion(pos.IdDedicacion.Id)
			if err3 != nil{
				return v, err3
			}
			v[x].LugarExpedicionCedula, err4 = BuscarLugarExpedicion(pos.IdPersona)
			if err4 != nil{
				return v, err4
			}
		}
		if v == nil {
			v = []models.VinculacionDocente{}
		}
		return v, nil
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
		return nil, outputError
	}

	return
}

func ListarDocentesCancelados(id_resolucion string) (VinculacionDocente []models.VinculacionDocente, outputError map[string]interface{}) {
	var v []models.VinculacionDocente
	var modRes []models.ModificacionResolucion
	var modVin []models.ModificacionVinculacion
	var cv models.VinculacionDocente
	// if 3 - modificacion_resolucion

	var err1 map[string]interface{}
	var err2 map[string]interface{}
	var err3 map[string]interface{}
	var err4 map[string]interface{}

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
					cv.NombreCompleto, err1 = BuscarNombreProveedor(documento_identidad)
					if err1 != nil{
						return v, err1
					}
					cv.NumeroDisponibilidad, err2 = BuscarNumeroDisponibilidad(vinculacion.VinculacionDocenteCancelada.Disponibilidad)
					if err2 != nil{
						return v, err2
					}
					cv.Dedicacion, err3 = BuscarNombreDedicacion(vinculacion.VinculacionDocenteCancelada.IdDedicacion.Id)
					if err3 != nil{
						return v, err3
					}
					cv.LugarExpedicionCedula, err4 = BuscarLugarExpedicion(vinculacion.VinculacionDocenteCancelada.IdPersona)
					if err4 != nil{
						return v, err4
					}
					cv.NumeroSemanasNuevas = vinculacion.VinculacionDocenteCancelada.NumeroSemanas - vinculacion.VinculacionDocenteRegistrada.NumeroSemanas
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
	return
}

func ListarDocentesCargaHoraria(vigencia string, periodo string, tipoVinculacion string, facultad string, nivelAcademico string) (newDocentesXcargaHoraria models.ObjetoCargaLectiva) {

	docentesXcargaHoraria, err := ListarDocentesHorasLectivas(vigencia, periodo, tipoVinculacion, facultad, nivelAcademico)
	if err != nil {
		beego.Error(err)
	}

	//BUSCAR CATEGORÍA DE CADA DOCENTE
	for _, pos := range docentesXcargaHoraria.CargasLectivas.CargaLectiva {
		catDocente := models.ObjetoCategoriaDocente{}
		emptyCatDocente := models.ObjetoCategoriaDocente{}
		//TODO: quitar el hardconding para WSO2 cuando ya soporte https:
		q := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudUrano") + "/categoria_docente/" + vigencia + "/" + periodo + "/" + pos.DocDocente
		err = GetXml(q, &catDocente.CategoriaDocente)
		if err != nil {
			beego.Error(err)
		}
		var err1 map[string]interface{}
		pos.CategoriaNombre, pos.IDCategoria, err1 = Buscar_Categoria_Docente(vigencia, periodo, pos.DocDocente)
		fmt.Println(err1)
		if err != nil {
			beego.Error(err)
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
		pos.IDFacultad, err = HomologarFacultad("old", pos.IDFacultad)
		if err != nil {
			beego.Error(err)
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}
	//RETORNAR PROYECTOS CURRICUALRES HOMOLOGADOS!!
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.DependenciaAcademica, err = strconv.Atoi(pos.IDProyecto)
		if err != nil {
			beego.Error(err)
		}
		pos.IDProyecto, err = HomologarProyectoCurricular(pos.IDProyecto)
		if err != nil {
			beego.Error(err)
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}

	return newDocentesXcargaHoraria
}

func ListarDocentesPrevinculadosAll(idResolucion string, tipoVinculacion int, tipoCancelacion int, tipoAdicion int, tipoReduccion int) (v []models.VinculacionDocente) {
	var res models.Resolucion
	var resvinc models.ResolucionVinculacionDocente
	var modres []models.ModificacionResolucion
	var vinc []models.VinculacionDocente
	var modvin []models.ModificacionVinculacion
	var ValorModificacionContrato float64

	var err1 map[string]interface{}
	var err2 map[string]interface{}
	var err3 map[string]interface{}
	var err4 map[string]interface{}
	var err5 map[string]interface{}

	//Devuelve el nivel académico, la dedicación y la facultad de la resolución
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente/"+idResolucion, &resvinc); err != nil && response != 200 {
		beego.Error(err)
	}

	//Devuelve la información básica de la resolución que se está consultando
	if response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucion, &res); err2 != nil && response2 != 200 {
		beego.Error(err2)
	}

	if res.IdTipoResolucion.Id != tipoVinculacion {
		//Busca el id de la modificación donde se relacionan la resolución original y la de modificación asociada
		if response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres); err3 != nil && response3 != 200 {
			beego.Error(err3)
		}
	}

	//Devuelve las vinculaciones presentes en la resolución consultada, agrupadas o no, según el nivel académico
	if resvinc.NivelAcademico == "POSGRADO" {
		if response4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente?limit=-1&query=IdResolucion.Id:"+idResolucion, &vinc); err4 != nil && response4 != 200 {
			beego.Error(err4)
		}
	}
	if resvinc.NivelAcademico == "PREGRADO" {
		if response5, err5 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_vinculaciones_agrupadas/"+idResolucion, &vinc); err5 != nil && response5 != 200 {
			beego.Error(err5)
		}
	}

	var llenarVinculacion = func(v *models.VinculacionDocente) {
		documentoIdentidad, _ := strconv.Atoi(v.IdPersona)
		v.NombreCompleto, err1 = BuscarNombreProveedor(documentoIdentidad)
		v.Dedicacion, err2 = BuscarNombreDedicacion(v.IdDedicacion.Id)
		v.LugarExpedicionCedula, err3 = BuscarLugarExpedicion(v.IdPersona)
		v.TipoDocumento, err4 = BuscarTipoDocumento(v.IdPersona)
		v.NumeroDisponibilidad, err5 = BuscarNumeroDisponibilidad(v.Disponibilidad)
	}

	switch res.IdTipoResolucion.Id {
	case tipoVinculacion:
		for x, pos := range vinc {
			v = append(v, pos)
			llenarVinculacion(&pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(pos.IdResolucion.Id), pos.IdDedicacion.Id)
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
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)

				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
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
				response5, err5 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err5 != nil && response5 != 200 {
					beego.Error(err5)
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
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
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
				response6, err6 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id)+",VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err6 != nil && response6 != 200 {
					beego.Error(err6)
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
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				pos.ValorModificacionFormato = FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
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
				response7, err7 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err7 != nil && response7 != 200 {
					beego.Error(err7)
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

	return v
}

func ListarDocentesPrevinculados(idResolucion string, tipoVinculacion int) (v []models.VinculacionDocente) {

	query := "?limit=-1&query=IdResolucion.Id:" + idResolucion + ",Estado:true"
	var res models.Resolucion
	var modres []models.ModificacionResolucion
	var modvin []models.ModificacionVinculacion

	var err1 map[string]interface{}
	var err2 map[string]interface{}
	var err3 map[string]interface{}
	var err4 map[string]interface{}
	var err5 map[string]interface{}

	response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucion, &res)
	if err != nil && response != 200 {
		beego.Error(err)
	}

	if res.IdTipoResolucion.Id == tipoVinculacion {
		response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v)
		if err2 != nil && response2 != 200 {
			beego.Error(err)
		}
	} else {
		response2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres)
		if err2 != nil && response2 != 200 {
			beego.Error(err)
		}
		response3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err3 != nil && response3 != 200 {
			beego.Error(err)
		}
		if len(modvin) != 0 {
			arreglo := make([]string, len(modvin))
			for x, pos := range modvin {
				arreglo[x] = strconv.Itoa(pos.VinculacionDocenteRegistrada.Id)
			}
			identificadoresvinc := strings.Join(arreglo, "|")
			response4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=Estado:True,Id__in:"+identificadoresvinc+"&limit=-1", &v)
			if err4 != nil && response4 != 200 {
				beego.Error(err)
			}
		} else {
			v = nil
		}
	}
	for x, pos := range v {
		documentoIdentidad, _ := strconv.Atoi(pos.IdPersona)

		pos.NombreCompleto, err1 = BuscarNombreProveedor(documentoIdentidad)
		pos.NumeroDisponibilidad, err2 = BuscarNumeroDisponibilidad(pos.Disponibilidad)
		pos.Dedicacion, err3 = BuscarNombreDedicacion(pos.IdDedicacion.Id)
		pos.LugarExpedicionCedula, err4 = BuscarLugarExpedicion(pos.IdPersona)
		pos.TipoDocumento, err5 = BuscarTipoDocumento(pos.IdPersona)
		pos.ValorContratoFormato = FormatMoney(int(v[x].ValorContrato), 2)
		pos.ProyectoNombre = BuscarNombreFacultad(int(v[x].IdProyectoCurricular))
		pos.Periodo = res.Periodo
		pos.VigenciaCarga = res.VigenciaCarga
		pos.PeriodoCarga = res.PeriodoCarga
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
		v[x] = pos
	}
	if v == nil {
		v = []models.VinculacionDocente{}
	}

	return v
}

func GetCdpRpDocente(identificacion string, num_vinculacion string, vigencia string) (rpdocente models.RpDocente) {

	var contratoDisponibilidad []models.ContratoDisponibilidad
	//If 1 contrato_disponibilidad (get)
	if response, err := GetJsonTest("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+num_vinculacion+",Vigencia:"+vigencia, &contratoDisponibilidad); err == nil && response == 200 { //If 2  (get)
		//for contrato_disponibilidad
		for _, pos := range contratoDisponibilidad {
			rpdocente = GetInformacionRpDocente(strconv.Itoa(pos.NumeroCdp), strconv.Itoa(pos.VigenciaCdp), identificacion)
			return rpdocente
		}
	} else { //If 1 contrato_disponibilidad (get)
		fmt.Println("He fallado en If 1 contrato_disponibilidad (get), solucioname!!!", err)
		return rpdocente
	}
	return
}
