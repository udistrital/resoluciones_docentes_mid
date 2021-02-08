package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

func SupervisorActual(id_resolucion int) (id_supervisor_actual int) {
	var r models.Resolucion
	var j []models.JefeDependencia
	var s []models.SupervisorContrato
	var fecha = time.Now().Format("2006-01-02")
	//If Resolucion (GET)
	if err := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(id_resolucion), &r); err == nil {
		//If Jefe_dependencia (GET)
		if err := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=DependenciaId:"+strconv.Itoa(r.IdDependencia)+",FechaFin__gte:"+fecha+",FechaInicio__lte:"+fecha, &j); err == nil {
			//If Supervisor (GET)
			if err := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/supervisor_contrato/?query=Documento:"+strconv.Itoa(j[0].TerceroId)+",FechaFin__gte:"+fecha+",FechaInicio__lte:"+fecha+"&CargoId.Cargo__startswith:DECANO|VICE", &s); err == nil {
				return s[0].Id
			} else { //If Jefe_dependencia (GET)
				fmt.Println("He fallado un poquito en If Supervisor (GET) en el método SupervisorActual, solucioname!!! ", err)
				return 0
			}
		} else { //If Jefe_dependencia (GET)
			fmt.Println("He fallado un poquito en If Jefe_dependencia (GET) en el método SupervisorActual, solucioname!!! ", err)
			return 0
		}
	} else { //If Resolucion (GET)
		fmt.Println("He fallado un poquito en If Resolucion (GET) en el método SupervisorActual, solucioname!!! ", err)
		return 0
	}
	return 0
}

func CalcularFechaFin(fecha_inicio time.Time, numero_semanas int) (fecha_fin time.Time) {
	var entero int
	var decimal float32
	meses := float32(numero_semanas) / 4
	entero = int(meses)
	decimal = meses - float32(entero)
	numero_dias := ((decimal * 4) * 7)
	f_i := fecha_inicio
	after := f_i.AddDate(0, entero, int(numero_dias))
	return after
}
