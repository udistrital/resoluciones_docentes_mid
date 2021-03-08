package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func ListarDocentesDesvinculados(query string) (VinculacionDocente []models.VinculacionDocente) {
	v := []models.VinculacionDocente{}

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v); err == nil && response == 200 {
		for x, pos := range v {
			documento_identidad, _ := strconv.Atoi(pos.IdPersona)
			v[x].NombreCompleto = BuscarNombreProveedor(documento_identidad)
			v[x].NumeroDisponibilidad = BuscarNumeroDisponibilidad(pos.Disponibilidad)
			v[x].Dedicacion = BuscarNombreDedicacion(pos.IdDedicacion.Id)
			v[x].LugarExpedicionCedula = BuscarLugarExpedicion(pos.IdPersona)
		}
		if v == nil {
			v = []models.VinculacionDocente{}
		}
		return v
	} else {
		return nil
	}

	return
}
