package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func SupervisorActual(id_resolucion int) (supervisor_actual models.SupervisorContrato, outputError map[string]interface{}) {
	var r models.Resolucion
	var j []models.JefeDependencia
	var s []models.SupervisorContrato
	var respuesta_peticion map[string]interface{}
	//var fecha = time.Now().Format("2006-01-02")   -- Se debe dejar este una vez se suba
	var fecha = "2018-01-01"
	//If Resolucion (GET)
	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+strconv.Itoa(id_resolucion), &respuesta_peticion); err == nil && response == 200 {
		LimpiezaRespuestaRefactor(respuesta_peticion, &r)
		//If Jefe_dependencia (GET)
		q := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudCore") + "/" + beego.AppConfig.String("NscrudCore") + "/jefe_dependencia?query=DependenciaId:" + strconv.Itoa(r.DependenciaId) + ",FechaFin__gte:" + fecha + ",FechaInicio__lte:" + fecha
		if response, err := GetJsonTest(q, &j); err == nil && response == 200 {
			//If Supervisor (GET)
			t := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAgora") + "/" + beego.AppConfig.String("NscrudAgora") + "/supervisor_contrato?query=Documento:" + strconv.Itoa(j[0].TerceroId) + ",FechaFin__gte:" + fecha + ",FechaInicio__lte:" + fecha + "&CargoId.Cargo__startswith:DECANO|VICE"
			if response, err := GetJsonTest(t, &s); err == nil && response == 200 {
				fmt.Println(s[0])
				return s[0], nil
			} else { //If Jefe_dependencia (GET)
				fmt.Println("He fallado un poquito en If Supervisor 1 (GET) en el método SupervisorActual, solucioname!!! ", err)
				outputError = map[string]interface{}{"funcion": "/SupervisorActual3", "err": err.Error(), "status": "404"}
				return s[0], outputError
			}
		} else { //If Jefe_dependencia (GET)
			fmt.Println("He fallado un poquito en If Jefe_dependencia 2 (GET) en el método SupervisorActual, solucioname!!! ", err)
			outputError = map[string]interface{}{"funcion": "/SupervisorActua2", "err": err.Error(), "status": "404"}
			return s[0], outputError
		}
	} else { //If Resolucion (GET)
		fmt.Println("He fallado un poquito en If Resolucion 3 (GET) en el método SupervisorActual, solucioname!!! ", err)
		outputError = map[string]interface{}{"funcion": "/SupervisorActual", "err": err.Error(), "status": "404"}
		return s[0], outputError
	}
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

func GetContenidoResolucion(id_resolucion string, id_facultad string) (contenidoResolucion models.ResolucionCompleta, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContenidoResolucion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var ordenador_gasto []models.OrdenadorGasto
	var jefe_dependencia []models.JefeDependencia
	var query string
	var respuesta_peticion map[string]interface{}

	if request, err1 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/contenido_resolucion/"+id_resolucion, &respuesta_peticion); err1 == nil && request == 200 {
		query = "?limit=-1&query=DependenciaId:" + id_facultad
		LimpiezaRespuestaRefactor(respuesta_peticion, &contenidoResolucion)
		if request2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto"+query, &ordenador_gasto); err2 == nil && request2 == 200 {
			fmt.Println(ordenador_gasto)
			if ordenador_gasto == nil || len(ordenador_gasto) == 0 {
				if request3, err3 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto?query=Id:1", &ordenador_gasto); err3 == nil && request3 == 200 {
					contenidoResolucion.OrdenadorGasto = ordenador_gasto[0]
				} else {
					logs.Error(err3)
					outputError = map[string]interface{}{"funcion": "/GetContenidoResolucion3", "err3": err3, "status": "502"}
					return contenidoResolucion, outputError
				}
			} else {
				contenidoResolucion.OrdenadorGasto = ordenador_gasto[0]
			}

		} else {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/GetContenidoResolucion2", "err2": err2, "status": "502"}
			return contenidoResolucion, outputError
		}
	} else {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/GetContenidoResolucion1", "err": err1, "status": "502"}
		return contenidoResolucion, outputError
	}

	fecha_actual := time.Now().Format("2006-01-02")
	query = "?query=DependenciaId:" + id_facultad + ",FechaFin__gte:" + fecha_actual + ",FechaInicio__lte:" + fecha_actual
	var err5 map[string]interface{}
	if request4, err4 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia"+query, &jefe_dependencia); err4 == nil && request4 == 200 {
		contenidoResolucion.OrdenadorGasto.NombreOrdenador, err5 = BuscarNombreProveedor(jefe_dependencia[0].TerceroId)
		if err5 != nil {
			logs.Error(err4)
			//outputError = map[string]interface{}{"funcion": "/GetContenidoResolucion5", "err5": err5, "status": "502"}
			return contenidoResolucion, err5
		}
	} else {
		logs.Error(err4)
		outputError = map[string]interface{}{"funcion": "/GetContenidoResolucion4", "err4": err4, "status": "502"}
		return contenidoResolucion, outputError
	}

	return contenidoResolucion, nil
}
