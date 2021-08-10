package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func InsertarPrevinculaciones(v []models.VinculacionDocente) (respuesta int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/InsertarPrevinculaciones", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var idRespuesta int
	var respuesta_peticion map[string]interface{}
	v, err := CalcularSalarioPrecontratacion(v)
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InsertarPrevinculaciones", "err": err, "status": "404"}
		return respuesta, outputError
	}

	if err2 := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &respuesta_peticion, &v); err2 == nil {
		LimpiezaRespuestaRefactor(respuesta_peticion, &idRespuesta)
		respuesta = idRespuesta
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InsertarPrevinculaciones", "err2": err2, "status": "404"}
		return respuesta, outputError
	}

	return respuesta, nil
}
