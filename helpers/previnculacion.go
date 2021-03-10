package helpers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func InsertarPrevinculaciones(v []models.VinculacionDocente) (respuesta int) {
	var idRespuesta int

	v, err := CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
	}

	if err2 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &idRespuesta, &v); err2 == nil {
		respuesta = idRespuesta
	} else {
		beego.Error("Error al insertar docentes", err2)
	}

	return respuesta
}
