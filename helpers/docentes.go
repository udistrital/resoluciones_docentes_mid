package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
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
