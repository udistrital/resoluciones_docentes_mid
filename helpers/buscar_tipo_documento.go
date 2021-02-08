package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarTipoDocumento(Cedula string) (nombre_tipo_doc string) {
	var tipo_documento string
	var temp []models.InformacionPersonaNatural
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_persona_natural/?limit=-1&query=Id:"+Cedula, &temp); err2 == nil {
		if temp != nil {
			tipo_documento = temp[0].TipoDocumento.ValorParametro
		} else {
			tipo_documento = "N/A"
		}
	} else {
		fmt.Println("error en json", err2)
		tipo_documento = "N/A"
	}

	return tipo_documento

}
