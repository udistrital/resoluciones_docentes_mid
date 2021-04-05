package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func GetResolucionesAprobadas(query string, limit int, offset int) (resolucion_vinculacion_aprobada []models.ResolucionVinculacion, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetResolucionesAprobadas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var respuesta_peticion map[string]interface{}
	if response1, err1 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion/Aprobada"+"?query="+query+"&offset="+strconv.Itoa(offset)+"&limit="+strconv.Itoa(limit), &respuesta_peticion); err1 == nil && response1 == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &resolucion_vinculacion_aprobada)
		var err2, err3 map[string]interface{}
		for x, pos := range resolucion_vinculacion_aprobada {
			resolucion_vinculacion_aprobada[x].FacultadNombre, err2 = BuscarNombreFacultad(pos.Facultad)
			if err2 != nil {
				panic(err2)
			}
			resolucion_vinculacion_aprobada[x].FacultadFirmaNombre, err3 = BuscarNombreFacultad(pos.IdDependenciaFirma)
			if err3 != nil {
				panic(err3)
			}
		}
		return resolucion_vinculacion_aprobada, nil
	} else {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/GetResolucionesAprobadas", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	}
	return resolucion_vinculacion_aprobada, nil
}

func GetResolucionesInscritas(query []string, limit int, offset int) (resolucion_vinculacion []models.ResolucionVinculacion, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetResolucionesAprobadas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var respuesta_peticion map[string]interface{}
	
	fmt.Println(beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudAdmin") + "/resolucion_vinculacion")
	r := httplib.Get(beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudAdmin") + "/resolucion_vinculacion")
	r.Param("offset", strconv.Itoa(offset))
	r.Param("limit", strconv.Itoa(limit))
	for _, v := range query {
		r.Param("query", v)
	}
	var err2, err3 map[string]interface{}
	if err1 := r.ToJSON(&respuesta_peticion); err1 == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &resolucion_vinculacion)
		for x, pos := range resolucion_vinculacion {
			resolucion_vinculacion[x].FacultadNombre, err2 = BuscarNombreFacultad(pos.Facultad)
			if err2 != nil {
				panic(err2)
			}
			resolucion_vinculacion[x].FacultadFirmaNombre, err3 = BuscarNombreFacultad(pos.IdDependenciaFirma)
			if err3 != nil {
				panic(err3)
			}
		}

		return resolucion_vinculacion, nil

	} else {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/GetResolucionesInscritas", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	}

	return
}

func InsertarResolucionCompleta(v models.ObjetoResolucion) (id_resolucion_creada int, control bool, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var texto_resolucion models.ResolucionCompleta
	var respuesta_peticion map[string]interface{}
	//****MANEJO DE TRANSACCIONES!***!//

	//Se trae cuerpo de resolución según tipo
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contenido_resolucion/ResolucionTemplate/"+
		v.ResolucionVinculacionDocente.Dedicacion+"/"+v.ResolucionVinculacionDocente.NivelAcademico+"/"+strconv.Itoa(v.Resolucion.Periodo)+"/"+strconv.Itoa(v.Resolucion.TipoResolucionId.Id), &respuesta_peticion); err == nil && response == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &texto_resolucion)
		v.Resolucion.ConsideracionResolucion = texto_resolucion.Consideracion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
		return id_resolucion_creada, control, outputError
	}

	//Primero se inserta la resolución, si eso se realiza correctamente
	var err2 map[string]interface{}
	control, id_resolucion_creada, err2 = InsertarResolucion(v)
	if err2 != nil {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas2", "err2": err2, "status": "404"}
		return id_resolucion_creada, control, outputError
	}
	if control {
		//Si se inserta bien en resolución, se puede insertar en resolucion_vinculacion_docente y en resolucion_estado
		control = InsertarResolucionVinDocente(id_resolucion_creada, v.ResolucionVinculacionDocente)
		control = InsertarResolucionEstado(id_resolucion_creada)
		//Si todo sigue bien, se inserta en componente_resolucion
		if control {
			InsertarArticulos(id_resolucion_creada, texto_resolucion.Articulos)
		} else {
			fmt.Println("enviar error al insertar en resolucion_vinculacion_docente")
		}
	} else {
		fmt.Println("envia error al insertar en resolución")
	}

	return id_resolucion_creada, control, outputError
}

func InsertarResolucionEstado(id_res int) (contr bool) {

	var respuesta models.ResolucionEstado
	var respuesta_peticion map[string]interface{}
	var cont bool
	temp := models.ResolucionEstado{
		FechaCreacion: time.Now(),
		EstadoResolucionId:        &models.EstadoResolucion{Id: 1},
		ResolucionId:    &models.Resolucion{Id: id_res},
	}

	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &respuesta_peticion, &temp); err == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta)
		cont = true
	} else {
		cont = false
	}

	return cont
}

func InsertarResolucionVinDocente(id_res int, resvindoc *models.ResolucionVinculacionDocente) (contr bool) {
	var temp = resvindoc
	var respuesta models.ResolucionVinculacionDocente
	var respuesta_peticion map[string]interface{}
	var cont bool
	temp.Id = id_res

	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente", "POST", &respuesta_peticion, &temp); err == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta)
		cont = true
	} else {

		cont = false
	}

	return cont
}

func InsertarArticulos(id_resolucion int, articulos []models.Articulo) {
	fmt.Println("Articulos y parágrafos")
	var respuesta models.ComponenteResolucion
	var respuesta2 models.ComponenteResolucion
	var respuesta_peticion map[string]interface{}

	for x, pos := range articulos {
		temp := models.ComponenteResolucion{
			Numero:         x + 1,
			ResolucionId:   &models.Resolucion{Id: id_resolucion},
			Texto:          pos.Texto,
			TipoComponente: "Articulo"}
		if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta_peticion, &temp); err == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta)
			for y, pos2 := range pos.Paragrafos {
				temp2 := models.ComponenteResolucion{
					Numero:          y + 1,
					ResolucionId:    &models.Resolucion{Id: id_resolucion},
					Texto:           pos2.Texto,
					TipoComponente:  "Paragrafo",
					ComponenteResolucionPadre: &models.ComponenteResolucion{Id: respuesta.Id},
				}
				if err2 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta_peticion, &temp2); err == nil {
					LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta2)
				} else {
					fmt.Println("error al insertar parágrafos", err2)
				}
			}

		} else {
			fmt.Println("error al insertar articulos", err)
		}
	}

}

func InsertarResolucion(resolucion models.ObjetoResolucion) (contr bool, id_cre int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/InsertarResolucion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var err1 map[string]interface{}
	if resolucion.NomDependencia, err1 = BuscarNombreFacultad(resolucion.Resolucion.DependenciaId); err1 != nil {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/InsertarResolucion1", "err1": err1, "status": "404"}
		return contr, id_cre, outputError
	}
	var temp = resolucion.Resolucion
	var respuesta models.Resolucion
	var id_creada int
	var cont bool
	var respuesta_modificacion_res models.ModificacionResolucion
	var resVieja models.Resolucion
	var motivo string
	var dedicacion string
	var articulo string
	var respuesta_peticion map[string]interface{}

	temp.Vigencia, _, _ = time.Now().Date()
	temp.FechaCreacion = time.Now()
	temp.Activo = true
	switch resolucion.ResolucionVinculacionDocente.Dedicacion {
	case "HCH":
		motivo = " RECONOCEN HONORARIOS "
		dedicacion = "hora catedra honorarios"
		articulo = "tercero"
		break
	case "HCP":
		motivo = "vinculan"
		dedicacion = "hora cátedra"
		articulo = "tercero"
		break
	case "TCO-MTO":
		motivo = "vinculan"
		dedicacion = "Tiempo Completo Ocasional o Medio Tiempo Ocasional"
		articulo = "tercero"
	}

	if temp.TipoResolucionId.Id == 1 {
		if resolucion.ResolucionVinculacionDocente.NivelAcademico == "POSGRADO" && resolucion.ResolucionVinculacionDocente.Dedicacion == "HCH" {
			temp.Titulo = "“Por la cual se " + motivo + " a docentes " + " para el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " Periodo Académico de " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de docentes de " + dedicacion + " para la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas (" + resolucion.ResolucionVinculacionDocente.NivelAcademico + ").”"

		} else {
			if resolucion.ResolucionVinculacionDocente.Dedicacion == "HCH" && resolucion.ResolucionVinculacionDocente.NivelAcademico == "PREGRADO" {
				temp.Titulo = "“Por la cual se " + motivo + " a docentes para finalizar el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " PERIODO académico del " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de docentes de " + dedicacion + " para la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas (" + resolucion.ResolucionVinculacionDocente.NivelAcademico + ").”"

			}
			if resolucion.ResolucionVinculacionDocente.Dedicacion == "HCP" && resolucion.ResolucionVinculacionDocente.NivelAcademico == "POSGRADO" {
				temp.Titulo = "“Por la cual se " + motivo + "  docentes para  el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " PERIODO académico de " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de docentes de " + dedicacion + " (vinculación especial) para la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas ( " + resolucion.ResolucionVinculacionDocente.NivelAcademico + ").”"

			}
			if resolucion.ResolucionVinculacionDocente.Dedicacion == "HCP" && resolucion.ResolucionVinculacionDocente.NivelAcademico == "PREGRADO" {
				temp.Titulo = "“Por la cual se " + motivo + "  docentes para finalizar el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " PERIODO académico de " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de docentes de" + dedicacion + " (vinculación especial) para la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas ( " + resolucion.ResolucionVinculacionDocente.NivelAcademico + ").”"

			} else {
				if resolucion.ResolucionVinculacionDocente.Dedicacion == "TCO-MTO" {
					temp.Titulo = "“Por la cual se " + motivo + "  docentes para finalizar el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " PERIODO académico de " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de docentes de " + dedicacion + " (vinculación especial) para la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas ( " + resolucion.ResolucionVinculacionDocente.NivelAcademico + ").”"
				}
			}

		}

	}
	if temp.TipoResolucionId.Id != 1 {
		temp.VigenciaCarga = resVieja.VigenciaCarga
		temp.PeriodoCarga = resVieja.PeriodoCarga
		if response, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(resolucion.ResolucionVieja), &respuesta_peticion); err2 == nil && response == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &resVieja)
			temp.Titulo = "“Por la cual se Modifica la resolución " + resVieja.NumeroResolucion + " de " + cambiarString(resVieja.FechaExpedicion.Month().String()) + " del " + strconv.Itoa(resVieja.FechaExpedicion.Year()) + " en cuanto a carga académica y valor del vínculo para el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " Periodo Académico de " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de Docentes de " + cambiarString(resolucion.ResolucionVinculacionDocente.Dedicacion) + " (Vinculación Especial) para la " + resolucion.NomDependencia + " en " + resolucion.ResolucionVinculacionDocente.NivelAcademico + ".”"
		} else {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/InsertarResolucion2", "err": err2.Error(), "status": "404"}
			return cont, id_cre, outputError
		}
	}
	temp.PreambuloResolucion = "El Decano de la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas, en uso de sus facultades legales y estatutarias, en particular, de las conferidas por el artículo " + articulo + "  de la Resolución de Rectoría Nro. xxx de enero xxx de 2021, y"
	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion", "POST", &respuesta_peticion, &temp); err == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta)
		id_creada = respuesta.Id
		cont = true
	} else {
		cont = false
		id_creada = 0
		outputError = map[string]interface{}{"funcion": "/InsertarResolucion3", "err": err.Error(), "status": "404"}
		return cont, id_creada, outputError
	}

	if temp.TipoResolucionId.Id != 1 {
		objeto_modificacion_res := models.ModificacionResolucion{
			ResolucionAnteriorId: &models.Resolucion{Id: resolucion.ResolucionVieja},
			ResolucionNuevaId:    &models.Resolucion{Id: id_creada},
		}
		if err3 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion", "POST", &respuesta_peticion, &objeto_modificacion_res); err3 == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &respuesta_modificacion_res)
			cont = true
		} else {
			cont = false
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/InsertarResolucion3", "err3": err3.Error(), "status": "404"}
			return cont, id_cre, outputError
		}

	}

	return cont, id_creada, nil
}

func cambiarString(original string) (cambiado string) {
	switch {
	//Periodos
	case original == "1":
		cambiado = "Primer"

	case original == "2":
		cambiado = "Segundo"

	case original == "3":
		cambiado = "Tercer"

		// Meses
	case original == "January":
		cambiado = "Enero"

	case original == "February":
		cambiado = "Febrero"

	case original == "March":
		cambiado = "Marzo"

	case original == "April":
		cambiado = "Abril"

	case original == "May":
		cambiado = "Mayo"

	case original == "June":
		cambiado = "Junio"

	case original == "July":
		cambiado = "Julio"

	case original == "August":
		cambiado = "Agosto"

	case original == "September":
		cambiado = "Septiembre"

	case original == "October":
		cambiado = "Octubre"

	case original == "November":
		cambiado = "Noviembre"

	case original == "December":
		cambiado = "Diciembre"

		//Dedicación
	case original == "HCH":
		cambiado = "Hora Cátedra Honorarios"

	case original == "HCP":
		cambiado = "Hora Cátedra Salarios"

	case original == "TCO-MTO":
		cambiado = "Tiempo Completo Ocasional - Medio Tiempo Ocasional"

	default:
		cambiado = original
	}

	return cambiado
}
