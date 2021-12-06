package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func ListarDisponibilidades(vigencia, UnidadEjecutora, limit, offset int, query string) (d []models.Disponibilidad, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ListarDisponibilidades", "err": err, "status": "500"}
			panic(outputError)
		}
	}()

	var disp []models.Disponibilidad
	url := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlmidFinanciera") + "/" + beego.AppConfig.String("NsmidFinanciera") + "/disponibilidad/ListaDisponibilidades/"
	params := strconv.Itoa(vigencia) + "?query=" + query + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset) + "&UnidadEjecutora=" + strconv.Itoa(UnidadEjecutora)
	if err := GetJson(url+params, &disp); err != nil {
		panic(err.Error())
	}

	return disp, outputError
}

func GetAllDisponibilidad(query string, limit, offset int) (d []models.Disponibilidad, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetAllDisponibilidad", "err": err, "status": "500"}
			panic(outputError)
		}
	}()

	var disp []models.Disponibilidad
	url := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudKronos") + "/" + beego.AppConfig.String("NscrudKronos") + "/disponibilidad"
	params := "?query=" + query + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
	if err := GetJson(url+params, &disp); err != nil {
		panic(err.Error())
	}

	return disp, outputError
}

func GetTotalDisponibilidades(vigencia, unidadEjecutora int) (total int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetTotalDisponibilidades", "err": err, "status": "500"}
			panic(outputError)
		}
	}()
	var resp map[string]int
	url := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudKronos") + "/" + beego.AppConfig.String("NscrudKronos") + "/disponibilidad/TotalDisponibilidades/"
	params := strconv.Itoa(vigencia) + "?UnidadEjecutora=" + strconv.Itoa(unidadEjecutora)
	if err := GetJson(url+params, &resp); err == nil {
		total = resp["Data"]
	} else {
		panic(err.Error())
	}

	return total, outputError
}
