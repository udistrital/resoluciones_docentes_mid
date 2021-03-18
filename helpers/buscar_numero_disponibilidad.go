package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNumeroDisponibilidad(IdCDP int) (numero_disp int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/BuscarNumeroDisponibilidad", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var temp []models.Disponibilidad
	var numero_disponibilidad int
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad?limit=-1&query=DisponibilidadApropiacion.Id:"+strconv.Itoa(IdCDP), &temp); err2 == nil {
		if temp != nil {
			numero_disponibilidad = int(temp[0].NumeroDisponibilidad)

		} else {
			numero_disponibilidad = 0
		}

	} else {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/BuscarNumeroDisponibilidad2", "err2": err2.Error(), "status": "502"}
		return numero_disponibilidad, outputError
	}
	return numero_disponibilidad, outputError

}
