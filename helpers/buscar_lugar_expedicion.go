package helpers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarLugarExpedicion(Cedula string) (nombre_lugar_exp string) {

	var nombre_ciudad string
	var temp []models.InformacionPersonaNatural
	var temp2 []models.Ciudad

	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_persona_natural/?limit=-1&query=Id:"+Cedula, &temp); err2 == nil {
		if temp != nil {
			id_ciudad := temp[0].IdCiudadExpedicionDocumento
			if err := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ciudad/?limit=-1&query=Id:"+strconv.Itoa(int(id_ciudad)), &temp2); err2 == nil {
				if temp2 != nil {
					nombre_ciudad = temp2[0].Nombre

				} else {
					nombre_ciudad = "N/A"
				}

			} else {
				fmt.Println("error en json", err)
			}

		} else {
			nombre_ciudad = "N/A"
		}

	} else {
		fmt.Println("error en json", err2)
	}

	return nombre_ciudad

}