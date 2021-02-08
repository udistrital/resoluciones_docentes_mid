package helpers

import (
	"encoding/json"
	"fmt"

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

func HomologarFacultad(tipo, facultad string) (facultad_old string, err error) {
	var id_facultad string
	var temp map[string]interface{}
	var string_consulta_servicio string

	if tipo == "new" {
		string_consulta_servicio = "facultad_gedep_oikos"
	} else {
		string_consulta_servicio = "facultad_oikos_gedep"
	}

	err = GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+string_consulta_servicio+"/"+facultad, &temp)
	if err != nil {
		return facultad_old, err
	}
	if temp != nil {
		json_facultad, err := json.Marshal(temp)

		if err != nil {
			return facultad_old, err
		}

		var temp_proy models.ObjetoFacultad
		err = json.Unmarshal(json_facultad, &temp_proy)
		if err != nil {
			return facultad_old, err
		}

		if tipo == "new" {
			id_facultad = temp_proy.Homologacion.IdGeDep
		} else {
			id_facultad = temp_proy.Homologacion.IdOikos
		}

	} else {
		return id_facultad, fmt.Errorf("No hay datos de respuesta de las APIs")
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

func HomologarProyectoCurricular(proyecto_old string) (proyecto string, err error) {
	var id_proyecto string
	var temp map[string]interface{}

	err = GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+"proyecto_curricular_cod_proyecto/"+proyecto_old, &temp)
	if err != nil {
		return proyecto, err
	}

	json_proyecto_curricular, err := json.Marshal(temp)

	if err != nil {
		return proyecto, err
	}
	var temp_proy models.ObjetoProyectoCurricular
	err = json.Unmarshal(json_proyecto_curricular, &temp_proy)
	if err != nil {
		return proyecto, err
	}
	id_proyecto = temp_proy.Homologacion.IDOikos

	return id_proyecto, nil
}
