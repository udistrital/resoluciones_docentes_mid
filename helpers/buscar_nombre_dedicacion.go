package helpers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNombreDedicacion(id_dedicacion int) (nombre_dedicacion string) {
	var nom_dedicacion string
	query := "?limit=-1&query=Id:" + strconv.Itoa(id_dedicacion)
	var dedicaciones []models.Dedicacion
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/dedicacion"+query, &dedicaciones); err2 == nil {
		if dedicaciones != nil {
			nom_dedicacion = dedicaciones[0].Descripcion
		} else {
			nom_dedicacion = ""
		}

	}

	return nom_dedicacion
}