package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/administrativa_mid_api/models"
)

func HomologarDedicacion_nombre(dedicacion string) (vinculacion_old []string) {
	var id_dedicacion_old []string
	homologacion_dedicacion := `[
						{
							"nombre": "HCH",
							"old": "5",
							"new": "1"
						},
						{
							"nombre": "HCP",
							"old": "4",
							"new": "2"
						},
						{
							"nombre": "TCO|MTO",
							"old": "2",
							"new": "4"
						},{
							"nombre": "TCO|MTO",
							"old": "3",
							"new": "3"
						}
						]`

	byt := []byte(homologacion_dedicacion)
	var arreglo_homologacion []models.HomologacionDedicacion
	if err := json.Unmarshal(byt, &arreglo_homologacion); err != nil {
		panic(err)
	}

	for _, pos := range arreglo_homologacion {
		if pos.Nombre == dedicacion {
			id_dedicacion_old = append(id_dedicacion_old, pos.Old)
		}
	}

	return id_dedicacion_old
}

func HomologarFacultad(tipo, facultad string) (facultad_old string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/HomologarFacultad", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var id_facultad string
	var temp map[string]interface{}
	var string_consulta_servicio string

	if tipo == "new" {
		string_consulta_servicio = "facultad_gedep_oikos"
	} else {
		string_consulta_servicio = "facultad_oikos_gedep"
	}

	err1 := GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+string_consulta_servicio+"/"+facultad, &temp)
	if err1 != nil {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/HomologarFacultad1", "err1": err1, "status": "502"}
		return facultad_old, outputError
	}
	if temp != nil {
		json_facultad, err2 := json.Marshal(temp)

		if err2 != nil {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/HomologarFacultad2", "err2": err2, "status": "502"}
			return facultad_old, outputError
		}

		var temp_proy models.ObjetoFacultad
		err3 := json.Unmarshal(json_facultad, &temp_proy)
		if err3 != nil {
			logs.Error(err3)
			outputError = map[string]interface{}{"funcion": "/HomologarFacultad3", "err3": err3, "status": "502"}
			return facultad_old, outputError
		}

		if tipo == "new" {
			id_facultad = temp_proy.Homologacion.IdGeDep
		} else {
			id_facultad = temp_proy.Homologacion.IdOikos
		}

	} else {
		outputError = map[string]interface{}{"funcion": "/HomologarFacultad3", "errTemp": "No hay datos de respuesta de las APIs", "status": "502"}
		return facultad_old, outputError
	}

	return id_facultad, nil

}

func HomologarDedicacion_ID(tipo, dedicacion string) (vinculacion_old, nombre_vinculacion string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var id_dedicacion_old string
	var nombre_dedicacion string
	var comparacion string
	var resultado string
	homologacion_dedicacion := `[
						{
							"nombre": "HCH",
							"old": "5",
							"new": "1"
						},
						{
							"nombre": "HCP",
							"old": "4",
							"new": "2"
						},
						{
							"nombre": "TCO|MTO",
							"old": "2",
							"new": "4"
						},{
							"nombre": "TCO|MTO",
							"old": "3",
							"new": "3"
						}
						]`

	byt := []byte(homologacion_dedicacion)
	var arreglo_homologacion []models.HomologacionDedicacion
	if err := json.Unmarshal(byt, &arreglo_homologacion); err != nil {
		panic(err) //nunca esperado
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
		return vinculacion_old, nombre_vinculacion, outputError
	}

	for _, pos := range arreglo_homologacion {
		if tipo == "new" {
			comparacion = pos.New
			resultado = pos.Old
		} else {
			comparacion = pos.Old
			resultado = pos.New
		}

		if comparacion == dedicacion {
			id_dedicacion_old = resultado
			nombre_dedicacion = pos.Nombre
		}
	}

	return id_dedicacion_old, nombre_dedicacion, outputError
}

func HomologarProyectoCurricular(proyecto_old string) (proyecto string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CertificacionDocumentosAprobados3", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var id_proyecto string
	var temp map[string]interface{}

	err1 := GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+"proyecto_curricular_cod_proyecto/"+proyecto_old, &temp)
	if err1 != nil {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/HomologarProyectoCurricular1", "err1": err1.Error(), "status": "502"}
		return proyecto, outputError
	}

	json_proyecto_curricular, err2 := json.Marshal(temp)

	if err2 != nil {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/HomologarProyectoCurricular2", "err2": err2.Error(), "status": "502"}
		return proyecto, outputError
	}
	var temp_proy models.ObjetoProyectoCurricular
	err3 := json.Unmarshal(json_proyecto_curricular, &temp_proy)
	if err3 != nil {
		logs.Error(err3)
		outputError = map[string]interface{}{"funcion": "/HomologarProyectoCurricular3", "err3": err3.Error(), "status": "502"}
		return proyecto, outputError
	}
	id_proyecto = temp_proy.Homologacion.IDOikos

	return id_proyecto, nil
}
