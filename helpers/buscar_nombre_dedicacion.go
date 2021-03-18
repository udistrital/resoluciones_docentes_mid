package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNombreDedicacion(id_dedicacion int) (nombre_dedicacion string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/BuscarNombreDedicacion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var nom_dedicacion string
	query := "?limit=-1&query=Id:" + strconv.Itoa(id_dedicacion)
	var dedicaciones []models.Dedicacion
	if response, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/dedicacion"+query, &dedicaciones); err2 == nil && response == 200 {
		if dedicaciones != nil {
			nom_dedicacion = dedicaciones[0].Descripcion
		} else {
			nom_dedicacion = ""
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/BuscarNombreDedicacion", "err": err2.Error(), "status": "404"}
		return nom_dedicacion, outputError
	}

	return nom_dedicacion, nil
}
