package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func GetResolucionesAprobadas(query string, limit int, offset int) (resolucion_vinculacion_aprobada []models.ResolucionVinculacion) {
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion/Aprobada"+"?query="+query+"&offset="+strconv.Itoa(offset)+"&limit="+strconv.Itoa(limit), &resolucion_vinculacion_aprobada); err == nil && response == 200 {
		for x, pos := range resolucion_vinculacion_aprobada {
			resolucion_vinculacion_aprobada[x].FacultadNombre = BuscarNombreFacultad(pos.Facultad)
			resolucion_vinculacion_aprobada[x].FacultadFirmaNombre = BuscarNombreFacultad(pos.IdDependenciaFirma)
		}
		return resolucion_vinculacion_aprobada
	} else {
		beego.Error("Error de consulta en resolucion_vinculacion_aprobada", err)
		return resolucion_vinculacion_aprobada
	}
	return
}

func GetResolucionesInscritas(query []string, limit int, offset int) (resolucion_vinculacion []models.ResolucionVinculacion) {
	r := httplib.Get(beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAdmin") + "/" + beego.AppConfig.String("NscrudAdmin") + "/resolucion_vinculacion")
	r.Param("offset", strconv.Itoa(offset))
	r.Param("limit", strconv.Itoa(limit))
	for _, v := range query {
		r.Param("query", v)

	}

	if err := r.ToJSON(&resolucion_vinculacion); err == nil {
		for x, pos := range resolucion_vinculacion {
			resolucion_vinculacion[x].FacultadNombre = BuscarNombreFacultad(pos.Facultad)
			resolucion_vinculacion[x].FacultadFirmaNombre = BuscarNombreFacultad(pos.IdDependenciaFirma)
		}

		return resolucion_vinculacion

	} else {
		beego.Error("Error de consulta en resolucion_vinculacion", err)
		resolucion_vinculacion = []models.ResolucionVinculacion{}
	}

	return resolucion_vinculacion
}

func InsertarResolucionCompleta(v models.ObjetoResolucion) (id_resolucion_creada int, control bool) {
	var texto_resolucion models.ResolucionCompleta

	//****MANEJO DE TRANSACCIONES!***!//

	//Se trae cuerpo de resolución según tipo
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contenido_resolucion/ResolucionTemplate/"+
		v.ResolucionVinculacionDocente.Dedicacion+"/"+v.ResolucionVinculacionDocente.NivelAcademico+"/"+strconv.Itoa(v.Resolucion.Periodo)+"/"+strconv.Itoa(v.Resolucion.IdTipoResolucion.Id), &texto_resolucion); err == nil && response == 200 {
		v.Resolucion.ConsideracionResolucion = texto_resolucion.Consideracion
	} else {
		fmt.Println("Error de consulta en texto de resolucion", err)
	}

	//Primero se inserta la resolución, si eso se realiza correctamente
	control, id_resolucion_creada = InsertarResolucion(v)
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

	return id_resolucion_creada, control
}

func InsertarResolucionEstado(id_res int) (contr bool) {

	var respuesta models.ResolucionEstado
	var cont bool
	temp := models.ResolucionEstado{
		FechaRegistro: time.Now(),
		Estado:        &models.EstadoResolucion{Id: 1},
		Resolucion:    &models.Resolucion{Id: id_res},
	}

	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &respuesta, &temp); err == nil {
		cont = true
	} else {
		cont = false
	}

	return cont
}

func InsertarResolucionVinDocente(id_res int, resvindoc *models.ResolucionVinculacionDocente) (contr bool) {
	var temp = resvindoc
	var respuesta models.ResolucionVinculacionDocente

	var cont bool
	temp.Id = id_res

	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente", "POST", &respuesta, &temp); err == nil {

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

	for x, pos := range articulos {
		temp := models.ComponenteResolucion{
			Numero:         x + 1,
			ResolucionId:   &models.Resolucion{Id: id_resolucion},
			Texto:          pos.Texto,
			TipoComponente: "Articulo"}
		if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta, &temp); err == nil {
			for y, pos2 := range pos.Paragrafos {
				temp2 := models.ComponenteResolucion{
					Numero:          y + 1,
					ResolucionId:    &models.Resolucion{Id: id_resolucion},
					Texto:           pos2.Texto,
					TipoComponente:  "Paragrafo",
					ComponentePadre: &models.ComponenteResolucion{Id: respuesta.Id},
				}
				if err2 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta2, &temp2); err == nil {

				} else {
					fmt.Println("error al insertar parágrafos", err2)
				}
			}

		} else {
			fmt.Println("error al insertar articulos", err)
		}
	}

}

func InsertarResolucion(resolucion models.ObjetoResolucion) (contr bool, id_cre int) {
	resolucion.NomDependencia = BuscarNombreFacultad(resolucion.Resolucion.IdDependencia)
	var temp = resolucion.Resolucion
	var respuesta models.Resolucion
	var id_creada int
	var cont bool
	var respuesta_modificacion_res models.ModificacionResolucion
	var resVieja models.Resolucion
	var motivo string
	var dedicacion string
	var articulo string

	temp.Vigencia, _, _ = time.Now().Date()
	temp.FechaRegistro = time.Now()
	temp.Estado = true
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

	if temp.IdTipoResolucion.Id == 1 {
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
	if temp.IdTipoResolucion.Id != 1 {
		temp.VigenciaCarga = resVieja.VigenciaCarga
		temp.PeriodoCarga = resVieja.PeriodoCarga
		if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(resolucion.ResolucionVieja), &resVieja); err == nil && response == 200 {
			temp.Titulo = "“Por la cual se Modifica la resolución " + resVieja.NumeroResolucion + " de " + cambiarString(resVieja.FechaExpedicion.Month().String()) + " del " + strconv.Itoa(resVieja.FechaExpedicion.Year()) + " en cuanto a carga académica y valor del vínculo para el " + cambiarString(strconv.Itoa(temp.PeriodoCarga)) + " Periodo Académico de " + strconv.Itoa(temp.VigenciaCarga) + " en la modalidad de Docentes de " + cambiarString(resolucion.ResolucionVinculacionDocente.Dedicacion) + " (Vinculación Especial) para la " + resolucion.NomDependencia + " en " + resolucion.ResolucionVinculacionDocente.NivelAcademico + ".”"
		} else {
			fmt.Println("Error al consultar resolución vieja", err)
		}
	}
	temp.PreambuloResolucion = "El Decano de la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas, en uso de sus facultades legales y estatutarias, en particular, de las conferidas por el artículo " + articulo + "  de la Resolución de Rectoría Nro. xxx de enero xxx de 2021, y"
	if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion", "POST", &respuesta, &temp); err == nil {
		id_creada = respuesta.Id
		cont = true
	} else {
		cont = false
		id_creada = 0
	}

	if temp.IdTipoResolucion.Id != 1 {
		objeto_modificacion_res := models.ModificacionResolucion{
			ResolucionAnterior: resolucion.ResolucionVieja,
			ResolucionNueva:    id_creada,
		}
		if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion", "POST", &respuesta_modificacion_res, &objeto_modificacion_res); err == nil {
			cont = true
		} else {
			fmt.Println("error al insertar en modificacion resolucion", err)
			cont = false

		}

	}

	return cont, id_creada
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
