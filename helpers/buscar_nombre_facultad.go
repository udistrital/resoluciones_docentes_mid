package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNombreFacultad(id_facultad int) (nombre_facultad string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/BuscarNombreFacultad", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var facultad []models.Facultad
	var nom string
	if response, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia?query=Id:"+strconv.Itoa(id_facultad), &facultad); err2 == nil && response == 200{
		nom = facultad[0].Nombre
	} else {
		nom = "N/A"
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/BuscarNombreFacultad2", "err2": err2.Error(), "status": "502"}
		return nom, outputError
	}
	return nom, outputError
}
