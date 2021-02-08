package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNombreFacultad(id_facultad int) (nombre_facultad string) {

	var facultad []models.Facultad
	var nom string
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia?query=Id:"+strconv.Itoa(id_facultad), &facultad); err2 == nil {
		nom = facultad[0].Nombre
	} else {
		nom = "N/A"
	}
	return nom
}
