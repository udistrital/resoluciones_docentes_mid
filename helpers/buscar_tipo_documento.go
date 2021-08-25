package helpers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarTipoDocumento(Cedula string) (nombre_tipo_doc string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/BuscarTipoDocumento", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var tipo_documento string
	var temp []models.InformacionPersonaNatural
	if response, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_persona_natural?limit=-1&query=Id:"+Cedula, &temp); err2 == nil && response == 200 {
		if temp != nil {
			tipo_documento = temp[0].TipoDocumento.ValorParametro
		} else {
			tipo_documento = "N/A"
		}
	} else {
		tipo_documento = "N/A"
		outputError = map[string]interface{}{"funcion": "/BuscarTipoDocumento1", "err": err2.Error(), "status": "404"}
		return tipo_documento, outputError
	}

	return tipo_documento, nil

}
