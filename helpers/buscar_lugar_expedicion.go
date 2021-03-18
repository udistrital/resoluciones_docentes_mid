package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarLugarExpedicion(Cedula string) (nombre_lugar_exp string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/BuscarLugarExpedicion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var nombre_ciudad string
	var temp []models.InformacionPersonaNatural
	var temp2 []models.Ciudad

	if response, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_persona_natural/?limit=-1&query=Id:"+Cedula, &temp); err2 == nil && response == 200 {
		if temp != nil {
			id_ciudad := temp[0].IdCiudadExpedicionDocumento
			if err3 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ciudad/?limit=-1&query=Id:"+strconv.Itoa(int(id_ciudad)), &temp2); err2 == nil {
				if temp2 != nil {
					nombre_ciudad = temp2[0].Nombre
				} else {
					nombre_ciudad = "N/A"
				}
			} else {
				fmt.Println("error en json", err3)
				outputError = map[string]interface{}{"funcion": "/BuscarLugarExpedicion2", "err": err3.Error(), "status": "404"}
				return nombre_ciudad, outputError
			}
		} else {
			nombre_ciudad = "N/A"
		}
	} else {
		fmt.Println("error en json", err2)
		outputError = map[string]interface{}{"funcion": "/BuscarLugarExpedicion1", "err": err2.Error(), "status": "404"}
		return nombre_ciudad, outputError
	}

	return nombre_ciudad, nil

}
