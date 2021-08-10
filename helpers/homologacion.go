package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego"
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
	var id_facultad string
	var temp map[string]interface{}
	var string_consulta_servicio string

	if tipo == "new" {
		string_consulta_servicio = "facultad_gedep_oikos"
	} else {
		string_consulta_servicio = "facultad_oikos_gedep"
	}
	q := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudHomologacion") + "/" + string_consulta_servicio + "/" + facultad
	if response, err := GetJsonWSO2Test(q, &temp); err == nil && response == 200 {
	} else {
		outputError = map[string]interface{}{"funcion": "/HomologarFacultad1", "err": err.Error(), "status": "502"}
		return facultad_old, outputError
	}
	if temp != nil {
		json_facultad, err := json.Marshal(temp)

		if err != nil {
			outputError = map[string]interface{}{"funcion": "/HomologarFacultad2", "err": err.Error(), "status": "502"}
			return facultad_old, outputError
		}

		var temp_proy models.ObjetoFacultad
		err1 := json.Unmarshal(json_facultad, &temp_proy)
		if err != nil {
			outputError = map[string]interface{}{"funcion": "/HomologarFacultad3", "err": err1.Error(), "status": "502"}
			return facultad_old, outputError
		}

		if tipo == "new" {
			id_facultad = temp_proy.Homologacion.IdGeDep
		} else {
			id_facultad = temp_proy.Homologacion.IdOikos
		}

	} else {
		outputError = map[string]interface{}{"funcion": "/HomologarFacultad4", "err": "No hay datos de respuesta de las APIs", "status": "502"}
		return id_facultad, outputError
	}

	return id_facultad, nil

}

func HomologarDedicacion_ID(tipo, dedicacion string) (vinculacion_old, nombre_vinculacion string) {
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

	return id_dedicacion_old, nombre_dedicacion
}

func HomologarProyectoCurricular(proyecto_old string) (proyecto string, outputError map[string]interface{}) {
	var id_proyecto string
	var temp map[string]interface{}

	if response, err := GetJsonWSO2Test("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/proyecto_curricular_cod_proyecto/"+proyecto_old, &temp); err == nil && response == 200 {
	} else {
		outputError = map[string]interface{}{"funcion": "/HomologarProyectoCurricular1", "err": err.Error(), "status": "502"}
		return proyecto, outputError
	}

	json_proyecto_curricular, err := json.Marshal(temp)

	if err != nil {
		outputError = map[string]interface{}{"funcion": "/HomologarProyectoCurricular2", "err": err.Error(), "status": "502"}
		return proyecto, outputError
	}
	var temp_proy models.ObjetoProyectoCurricular
	err = json.Unmarshal(json_proyecto_curricular, &temp_proy)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/HomologarProyectoCurricular3", "err": err.Error(), "status": "502"}
		return proyecto, outputError
	}
	id_proyecto = temp_proy.Homologacion.IDOikos

	return id_proyecto, nil
}
