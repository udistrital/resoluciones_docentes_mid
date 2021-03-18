package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNombreProveedor(DocumentoIdentidad int) (nombre_prov string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/BuscarNombreProveedor", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var nom_proveedor string
	queryInformacionProveedor := "?query=NumDocumento:" + strconv.Itoa(DocumentoIdentidad)
	var informacion_proveedor []models.InformacionProveedor
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/"+queryInformacionProveedor, &informacion_proveedor); err2 == nil {
		if informacion_proveedor != nil {
			nom_proveedor = informacion_proveedor[0].NomProveedor
		} else {
			nom_proveedor = ""
		}

	} else {
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/BuscarNombreProveedor2", "err2": err2.Error(), "status": "502"}
		return nom_proveedor, outputError
	}

	return

}
