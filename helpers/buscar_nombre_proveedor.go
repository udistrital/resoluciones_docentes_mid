package helpers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func BuscarNombreProveedor(DocumentoIdentidad int) (nombre_prov string) {

	var nom_proveedor string
	queryInformacionProveedor := "?query=NumDocumento:" + strconv.Itoa(DocumentoIdentidad)
	var informacion_proveedor []models.InformacionProveedor
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/"+queryInformacionProveedor, &informacion_proveedor); err2 == nil {
		if informacion_proveedor != nil {
			nom_proveedor = informacion_proveedor[0].NomProveedor
		} else {
			nom_proveedor = ""
		}

	}

	return nom_proveedor

}