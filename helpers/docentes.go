package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
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
