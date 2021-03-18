package helpers

import (
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNumeroDisponibilidad(IdCDP int) (numero_disp int, outputError map[string]interface{}) {

	var temp []models.Disponibilidad
	var numero_disponibilidad int
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad?limit=-1&query=DisponibilidadApropiacion.Id:"+strconv.Itoa(IdCDP), &temp); err == nil && response == 200 {
		if temp != nil {
			numero_disponibilidad = int(temp[0].NumeroDisponibilidad)

		} else {
			numero_disponibilidad = 0
		}

	} else {
		fmt.Println("Error en disponibilidad (get) función BuscarNumeroDisponibilidad:", err)
		outputError = map[string]interface{}{"funcion": "/BuscarNumeroDisponibilidad", "err": err.Error(), "status": "404"}
		return numero_disponibilidad, outputError
	}
	return numero_disponibilidad, nil

}