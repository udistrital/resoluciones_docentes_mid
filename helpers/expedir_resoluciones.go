package helpers

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func Expedir(m models.ExpedicionResolucion) (outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/Expedir", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	vigencia, _, _ := time.Now().Date()
	var cdve int
	var temp int
	var tipoCon models.TipoContrato
	var proveedor []models.InformacionProveedor
	var response interface{}
	var disponibilidad models.Disponibilidad
	var dispoap models.DisponibilidadApropiacion
	var respuesta_peticion map[string]interface{}
	//var vincDocente models.VinculacionDocente

	v := m.Vinculaciones
	//If 12 - Consecutivo contrato_general
	if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil && responseG == 201 {
		numeroContratos := cdve
		// for vinculaciones
		for _, vinculacion := range *v {
			numeroContratos = numeroContratos + 1
			v := vinculacion.VinculacionDocente
			idvinculaciondocente := strconv.Itoa(v.Id)
			//if 8 - Vinculacion_docente (GET)
			q := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/vinculacion_docente/" + idvinculaciondocente
			if responseG, err := GetJsonTest(q, &respuesta_peticion); err == nil && responseG == 200 {
				fmt.Println("RESPUESTA ", respuesta_peticion)
				//vincDocente = models.VinculacionDocente{}
				LimpiezaRespuestaRefactor(respuesta_peticion, &v)
				fmt.Println("V ", v.DedicacionId)
				contrato := vinculacion.ContratoGeneral
				t := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAgora") + "/" + beego.AppConfig.String("NscrudAgora") + "/tipo_contrato/" + strconv.Itoa(contrato.TipoContrato.Id)
				if responseG, err := GetJsonTest(t, &tipoCon); err == nil && responseG == 200 {
					var sup models.SupervisorContrato
					acta := vinculacion.ActaInicio
					aux1 := 181
					contrato.VigenciaContrato = vigencia
					contrato.Id = "DVE" + strconv.Itoa(numeroContratos)
					contrato.FormaPago.Id = 240
					contrato.DescripcionFormaPago = "Abono a Cuenta Mensual de acuerdo a puntos y horas laboradas"
					contrato.Justificacion = "Docente de Vinculacion Especial"
					contrato.UnidadEjecucion.Id = 269
					contrato.LugarEjecucion.Id = 4
					contrato.TipoControl = aux1
					contrato.ClaseContratista = 33
					contrato.TipoMoneda = 137
					contrato.OrigenRecursos = 149
					contrato.OrigenPresupueso = 156
					contrato.TemaGastoInversion = 166
					contrato.TipoGasto = 146
					contrato.RegimenContratacion = 136
					contrato.Procedimiento = 132
					contrato.ModalidadSeleccion = 123
					contrato.TipoCompromiso = 35
					contrato.TipologiaContrato = 46
					contrato.FechaRegistro = time.Now()
					fmt.Println("FECHA ", contrato.FechaRegistro.Format(time.RFC3339))
					contrato.UnidadEjecutora = 1
					fmt.Println("LLEGO ", v.ResolucionVinculacionDocenteId.Id)
					sup, err := SupervisorActual(v.ResolucionVinculacionDocenteId.Id)
					if err != nil {
						fmt.Println(err)
						return err
					}
					contrato.Supervisor = &sup
					contrato.Condiciones = "Sin condiciones"

					// If 5 - Informacion_Proveedor
					m := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAgora") + "/" + beego.AppConfig.String("NscrudAgora") + "/informacion_proveedor?query=NumDocumento:" + strconv.Itoa(contrato.Contratista)
					if responseG, err := GetJsonTest(m, &proveedor); err == nil && responseG == 200 {
						if proveedor != nil { //Nuevo If
							temp = proveedor[0].Id
							contratoGeneral := make(map[string]interface{})
							contratoGeneral = map[string]interface{}{
								"Id":               contrato.Id,
								"VigenciaContrato": contrato.VigenciaContrato,
								"ObjetoContrato":   contrato.ObjetoContrato,
								"PlazoEjecucion":   contrato.PlazoEjecucion,
								"FormaPago": map[string]interface{}{
									"Id":                240,
									"Descripcion":       "TRANSACCIÓN",
									"CodigoContraloria": "'",
									"EstadoRegistro":    true,
									"FechaRegistro":     "2016-10-25T00:00:00Z",
								},
								"OrdenadorGasto":         contrato.OrdenadorGasto,
								"SedeSolicitante":        contrato.SedeSolicitante,
								"DependenciaSolicitante": contrato.DependenciaSolicitante,
								"Contratista":            temp,
								"UnidadEjecucion": map[string]interface{}{
									"Id":                269,
									"Descripcion":       "Semana(s)",
									"CodigoContraloria": "'",
									"EstadoRegistro":    true,
									"FechaRegistro":     "2018-03-20T00:00:00Z",
								},
								"ValorContrato":        int(contrato.ValorContrato),
								"Justificacion":        contrato.Justificacion,
								"DescripcionFormaPago": contrato.DescripcionFormaPago,
								"Condiciones":          contrato.Condiciones,
								"UnidadEjecutora":      contrato.UnidadEjecutora,
								"FechaRegistro":        contrato.FechaRegistro.Format(time.RFC3339),
								"TipologiaContrato":    contrato.TipologiaContrato,
								"TipoCompromiso":       contrato.TipoCompromiso,
								"ModalidadSeleccion":   contrato.ModalidadSeleccion,
								"Procedimiento":        contrato.Procedimiento,
								"RegimenContratacion":  contrato.RegimenContratacion,
								"TipoGasto":            contrato.TipoGasto,
								"TemaGastoInversion":   contrato.TemaGastoInversion,
								"OrigenPresupueso":     contrato.OrigenPresupueso,
								"OrigenRecursos":       contrato.OrigenRecursos,
								"TipoMoneda":           contrato.TipoMoneda,
								"TipoControl":          contrato.TipoControl,
								"Observaciones":        contrato.Observaciones,
								"Supervisor": map[string]interface{}{
									"Id":                    sup.Id,
									"Nombre":                sup.Nombre,
									"Documento":             sup.Documento,
									"Cargo":                 sup.Cargo,
									"SedeSupervisor":        sup.SedeSupervisor,
									"DependenciaSupervisor": sup.DependenciaSupervisor,
									"Tipo":                  sup.Tipo,
									"Estado":                sup.Estado,
									"DigitoVerificacion":    sup.DigitoVerificacion,
									"FechaInicio":           sup.FechaInicio,
									"FechaFin":              sup.FechaFin,
									"CargoId": map[string]interface{}{
										"Id": sup.CargoId.Id,
									},
								},
								"ClaseContratista": contrato.ClaseContratista,
								"TipoContrato": map[string]interface{}{
									"Id":           6,
									"TipoContrato": "Contrato de Prestación de Servicios Profesionales o Apoyo a la Gestión",
									"Estado":       true,
								},
								"LugarEjecucion": map[string]interface{}{
									"Id":          4,
									"Direccion":   "CALLE 40 A No 13-09",
									"Sede":        "00IP",
									"Dependencia": "DEP39",
									"Ciudad":      96,
								},
							}
							fmt.Println(contratoGeneral)
							p := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAgora") + "/" + beego.AppConfig.String("NscrudAgora") + "/contrato_general"
							//var response1 models.ContratoGeneral
							if err := SendJson(p, "POST", &response, contratoGeneral); err == nil {
								//var id1 = response
								fmt.Println("response")
								fmt.Println(response)
								aux1 := contrato.Id
								aux2 := contrato.VigenciaContrato
								var ce models.ContratoEstado
								var ec models.EstadoContrato
								ce.NumeroContrato = aux1
								ce.Vigencia = aux2
								ce.FechaRegistro = time.Now()
								ec.Id = 4
								ce.Estado = &ec
								// If 4 - contrato_estado
								var response2 models.ContratoEstado
								if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response2, &ce); err == nil {
									//var id2 = response
									a := vinculacion.VinculacionDocente
									var ai models.ActaInicio
									ai.NumeroContrato = aux1
									ai.Vigencia = aux2
									ai.Descripcion = acta.Descripcion
									ai.FechaInicio = acta.FechaInicio
									ai.FechaFin = acta.FechaFin
									ai.FechaFin = CalcularFechaFin(acta.FechaInicio, a.NumeroSemanas)
									ai.FechaRegistro = time.Now()
									// If 3 - Acta_inicio creación
									var response3 models.ActaInicio
									if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response3, &ai); err == nil {
										//var id3 = response
										var cd models.ContratoDisponibilidad
										cd.NumeroContrato = aux1
										fmt.Println("aux1 ", aux1)
										cd.Vigencia = aux2
										cd.Estado = true
										cd.FechaRegistro = time.Now()
										// If 2.5.2 - Get disponibildad_apropiacion
										f := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudKronos") + "/" + beego.AppConfig.String("NscrudKronos") + "/disponibilidad_apropiacion/" + strconv.Itoa(v.Disponibilidad)
										if responseG, err := GetJsonTest(f, &dispoap); err == nil && responseG == 200 {
											// If 2.5.1 - Get disponibildad
											if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/"+strconv.Itoa(dispoap.Disponibilidad.Id), &disponibilidad); err == nil && responseG == 200 {
												cd.NumeroCdp = int(disponibilidad.NumeroDisponibilidad)
												cd.VigenciaCdp = int(disponibilidad.Vigencia)
												// If 2 - contrato_disponibilidad
												var response4 models.ContratoDisponibilidad
												if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response4, &cd); err == nil {
													a.PuntoSalarialId = vinculacion.VinculacionDocente.PuntoSalarialId
													a.SalarioMinimoId = vinculacion.VinculacionDocente.SalarioMinimoId
													v := a
													v.NumeroContrato = aux1
													//v.NumeroContrato.Valid = true
													v.Vigencia = aux2
													//v.Vigencia.Valid = true
													v.FechaInicio = acta.FechaInicio
													//v.FechaModificacion = time.Now()
													fmt.Println("CD es: ", response4.Id)
													fmt.Println("AI es: ", response3.Id)
													fmt.Println("CE es: ", response2.Id)
													fmt.Println("CONTRATO GENERAL es: ", contrato.Id)
													// If 1 - vinculacion_docente
													if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente/"+strconv.Itoa(v.Id), "PUT", &respuesta_peticion, &v); err == nil {
														LimpiezaRespuestaRefactor(respuesta_peticion, &response)
														fmt.Println("Response ", response)
														fmt.Println("Vinculacion docente actualizada y lista, vamos por la otra")
													} else { // If 1 - vinculacion_docente
														logs.Error(v)
														outputError = map[string]interface{}{"funcion": "/Expedir1.0", "err": err.Error(), "status": "502"}
														return outputError
													}
												} else { // If 2 - contrato_disponibilidad
													fmt.Println("He fallado un poquito en  If 2 - contrato_disponibilidad, solucioname!!!", err)
													logs.Error(cd)
													outputError = map[string]interface{}{"funcion": "/Expedir2", "err": err.Error(), "status": "502"}
													return outputError
												}
											} else { // If 2.5.1 - Get disponibildad
												fmt.Println("He fallado un poquito en If 2.5.1 - Get disponibildad, solucioname!!!", err)
												logs.Error(disponibilidad)
												outputError = map[string]interface{}{"funcion": "/Expedir2.5.1", "err": err.Error(), "status": "502"}
												return outputError
											}
										} else { // If 2.5.2 - Get disponibildad_apropiacion
											fmt.Println("He fallado un poquito en If 2.5.2 - Get disponibildad_apropiacion, solucioname!!!", err)
											logs.Error(dispoap)
											outputError = map[string]interface{}{"funcion": "/Expedir2.5.2", "err": err.Error(), "status": "502"}
											return outputError
										}
									} else { // If 3 - Acta_inicio
										fmt.Println("He fallado un poquito en If 3 - Acta_inicio, solucioname!!!", err)
										logs.Error(ai)
										outputError = map[string]interface{}{"funcion": "/Expedir3", "err": err.Error(), "status": "502"}
										return outputError
									}
								} else { // If 4 - contrato_estado
									fmt.Println("He fallado un poquito en If 4 - contrato_estado, solucioname!!!", err)
									logs.Error(ce)
									outputError = map[string]interface{}{"funcion": "/Expedir4", "err": err.Error(), "status": "502"}
									return outputError
								}
							} else { //If insert contrato_general
								fmt.Println("He fallado un poquito en insert contrato_general, solucioname!!!", err)
								outputError = map[string]interface{}{"funcion": "/Expedir4.1", "err": err.Error(), "status": "502"}
								return outputError
							}
						} else { // Nuevo If
							fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor nuevo, solucioname!!!", err)
							outputError = map[string]interface{}{"funcion": "/Expedir4.2", "err": err.Error(), "status": "502"}
							return outputError
						}
					} else { // If 5 - Informacion_Proveedor
						fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor, solucioname!!!", err)
						outputError = map[string]interface{}{"funcion": "/Expedir5 ", "err": err.Error(), "status": "502"}
						return outputError
					}
				} else {
					fmt.Println("error")
					outputError = map[string]interface{}{"funcion": "/Expedir7 ", "err": err.Error(), "status": "502"}
					return outputError
				}
			} else { //If 8 - Vinculacion_docente (GET)
				fmt.Println("He fallado un poquito en If 8 - Vinculacion_docente (GET), solucioname!!!", err)
				logs.Error(v)
				outputError = map[string]interface{}{"funcion": "/Expedir8 ", "err": err.Error(), "status": "502"}
				return outputError
			}
		} // for vinculaciones
		var r models.Resolucion
		r.Id = m.IdResolucion
		idResolucionDVE := strconv.Itoa(m.IdResolucion)
		//If 11 - Resolucion (GET)
		s := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/resolucion/" + idResolucionDVE
		if responseG, err := GetJsonTest(s, &respuesta_peticion); err == nil && responseG == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &r)
			r.FechaExpedicion = m.FechaExpedicion
			//If 10 - Resolucion (PUT)
			if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &respuesta_peticion, &r); err == nil {
				LimpiezaRespuestaRefactor(respuesta_peticion, &response)
				var e models.ResolucionEstado
				var er models.EstadoResolucion
				e.ResolucionId = &r
				er.Id = 2
				e.EstadoResolucionId = &er
				//e.FechaCreacion = time.Now()
				//e.FechaModificacion = time.Now()
				//If 9 - Resolucion_estado
				if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion_estado", "POST", &respuesta_peticion, &e); err == nil {
					LimpiezaRespuestaRefactor(respuesta_peticion, &response)
					fmt.Println("ResponsePOST ", response)
					fmt.Println("Expedición exitosa, ahora va el commit :D")
					//c.Data["json"] = v
				} else { //If 9 - Resolucion_estado
					fmt.Println("He fallado un poquito en If 9 - Resolucion_estado, solucioname!!!", err)
					logs.Error(e)
					outputError = map[string]interface{}{"funcion": "/Expedir9 ", "err": err.Error(), "status": "502"}
					return outputError
				}
			} else { //If 10 - Resolucion (PUT)
				fmt.Println("He fallado un poquito en If 10 - Resolucion (PUT), solucioname!!! ", err)
				logs.Error(r)
				outputError = map[string]interface{}{"funcion": "/Expedir10 ", "err": err.Error(), "status": "502"}
				return outputError
			}
		} else { //If 11 - Resolucion (GET)
			fmt.Println("He fallado un poquito en If 11 - Resolucion (GET), solucioname!!! ", err)
			logs.Error(r)
			outputError = map[string]interface{}{"funcion": "/Expedir11 ", "err": err.Error(), "status": "502"}
			return outputError
		}
	} else { //If 12 - Consecutivo contrato_general
		fmt.Println("He fallado un poquito en If 12 - Consecutivo contrato_general, solucioname!!! ", err)
		logs.Error(cdve)
		outputError = map[string]interface{}{"funcion": "/Expedir12 ", "err": err.Error(), "status": "502"}
		return outputError
	}
	return nil
}

func ValidarDatosExpedicion(m models.ExpedicionResolucion) (outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	v := m.Vinculaciones
	beego.Info(v)

	var respuesta_peticion map[string]interface{}

	for _, vinculacion := range *v {
		v := vinculacion.VinculacionDocente
		idvinculaciondocente := strconv.Itoa(v.Id)

		if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente/"+idvinculaciondocente, &respuesta_peticion); responseG == 200 && err == nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &v)
		} else {
			beego.Error("Previnculación no valida", err)
			logs.Error(v)
			//c.Data["system"] = v
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion1", "err": err.Error(), "status": "502"}
			return outputError
		}

		contrato := vinculacion.ContratoGeneral
		var proveedor []models.InformacionProveedor

		if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); responseG == 200 && err == nil {
		} else {
			beego.Error("Docente no válido en Ágora, se encuentra identificado con el documento número ", strconv.Itoa(contrato.Contratista), err)
			logs.Error(proveedor)
			//c.Data["system"] = proveedor
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion2", "err": err.Error(), "status": "502"}
			return outputError
		}

		if proveedor == nil {
			beego.Error("No existe el docente con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			logs.Error(proveedor)
			//c.Data["system"] = proveedor
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion3", "err": "No existe el docente con este numero de documento", "status": "502"}
			return outputError
		}

		var dispoap []models.DisponibilidadApropiacion

		if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion?query=Id:"+strconv.Itoa(v.Disponibilidad), &dispoap); responseG == 200 && err == nil {
		} else {
			beego.Error("Disponibilidad no válida asociada al docente identificado con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			logs.Error(dispoap)
			//c.Data["system"] = dispoap
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion4", "err": err.Error(), "status": "502"}
			return outputError
		}

		if dispoap == nil {
			beego.Error("Disponibilidad no válida asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			logs.Error(dispoap)
			//c.Data["system"] = dispoap
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion5", "err": "Disponibilidad no válida asociada al docente identificado con número de documento", "status": "502"}
			return outputError
		}

		var proycur []models.Dependencia

		if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia?query=Id:"+strconv.Itoa(v.ProyectoCurricularId), &proycur); responseG == 200 && err == nil {
		} else {
			beego.Error("Dependencia incorrectamente homologada asociada al docente identificado con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			logs.Error(proycur)
			//c.Data["system"] = proycur
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion6", "err": err.Error(), "status": "502"}
			return outputError
		}

		if proycur == nil {
			beego.Error("Dependencia incorrectamente homologada asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			logs.Error(proycur)
			//c.Data["system"] = proycur
			//c.Abort("233")
			outputError = map[string]interface{}{"funcion": "/ValidarDatosExpedicion7", "err": "ependencia incorrectamente homologada asociada al docente identificado con ese numero de documento", "status": "502"}
			return outputError
		}
		beego.Info(proycur)

	}
	return
}

func ExpedirModificacion(m models.ExpedicionResolucion) (outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ExpedirModificacion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp int
	var cdve int
	var proveedor []models.InformacionProveedor
	var disponibilidad models.Disponibilidad
	var dispoap models.DisponibilidadApropiacion
	var modVin []models.ModificacionVinculacion
	var response interface{}
	var resolucion models.Resolucion
	var respuesta_peticion map[string]interface{}
	vigencia, _, _ := time.Now().Date()
	v := m.Vinculaciones

	// If 12 - Consecutivo contrato_general
	if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil && responseG == 201 {
		numeroContratos := cdve
		// for vinculaciones
		for _, vinculacion := range *v {
			numeroContratos = numeroContratos + 1
			v := vinculacion.VinculacionDocente
			fmt.Println("V es ", v)
			idvinculaciondocente := strconv.Itoa(v.Id)
			// if 8 - Vinculacion_docente (GET)
			q := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/vinculacion_docente/" + idvinculaciondocente
			if responseG, err := GetJsonTest(q, &respuesta_peticion); err == nil && responseG == 200 {
				LimpiezaRespuestaRefactor(respuesta_peticion, &v)
				contrato := vinculacion.ContratoGeneral
				var sup models.SupervisorContrato
				acta := vinculacion.ActaInicio
				aux1 := 181
				contrato.VigenciaContrato = vigencia
				contrato.Id = "DVE" + strconv.Itoa(numeroContratos)
				contrato.FormaPago.Id = 240
				contrato.DescripcionFormaPago = "Abono a Cuenta Mensual de acuerdo a puntos y horas laboradas"
				contrato.Justificacion = "Docente de Vinculacion Especial"
				contrato.UnidadEjecucion.Id = 269
				contrato.LugarEjecucion.Id = 4
				contrato.TipoControl = aux1
				contrato.ClaseContratista = 33
				contrato.TipoMoneda = 137
				contrato.OrigenRecursos = 149
				contrato.OrigenPresupueso = 156
				contrato.TemaGastoInversion = 166
				contrato.TipoGasto = 146
				contrato.RegimenContratacion = 136
				contrato.Procedimiento = 132
				contrato.ModalidadSeleccion = 123
				contrato.TipoCompromiso = 35
				contrato.TipologiaContrato = 46
				contrato.FechaRegistro = time.Now()
				contrato.UnidadEjecutora = 1
				sup, err := SupervisorActual(v.ResolucionVinculacionDocenteId.Id)
				if err != nil {
					return err
				}
				contrato.Supervisor = &sup
				contrato.Condiciones = "Sin condiciones"
				// If 5 - Informacion_Proveedor
				if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil && responseG == 200 {
					if proveedor != nil { //Nuevo If
						temp = proveedor[0].Id
						contratoGeneral := make(map[string]interface{})
						contratoGeneral = map[string]interface{}{
							"Id":               contrato.Id,
							"VigenciaContrato": contrato.VigenciaContrato,
							"ObjetoContrato":   contrato.ObjetoContrato,
							"PlazoEjecucion":   contrato.PlazoEjecucion,
							"FormaPago": map[string]interface{}{
								"Id":                240,
								"Descripcion":       "TRANSACCIÓN",
								"CodigoContraloria": "'",
								"EstadoRegistro":    true,
								"FechaRegistro":     "2016-10-25T00:00:00Z",
							},
							"OrdenadorGasto":         contrato.OrdenadorGasto,
							"SedeSolicitante":        contrato.SedeSolicitante,
							"DependenciaSolicitante": contrato.DependenciaSolicitante,
							"Contratista":            temp,
							"UnidadEjecucion": map[string]interface{}{
								"Id":                269,
								"Descripcion":       "Semana(s)",
								"CodigoContraloria": "'",
								"EstadoRegistro":    true,
								"FechaRegistro":     "2018-03-20T00:00:00Z",
							},
							"ValorContrato":        int(contrato.ValorContrato),
							"Justificacion":        contrato.Justificacion,
							"DescripcionFormaPago": contrato.DescripcionFormaPago,
							"Condiciones":          contrato.Condiciones,
							"UnidadEjecutora":      contrato.UnidadEjecutora,
							"FechaRegistro":        contrato.FechaRegistro.Format(time.RFC3339),
							"TipologiaContrato":    contrato.TipologiaContrato,
							"TipoCompromiso":       contrato.TipoCompromiso,
							"ModalidadSeleccion":   contrato.ModalidadSeleccion,
							"Procedimiento":        contrato.Procedimiento,
							"RegimenContratacion":  contrato.RegimenContratacion,
							"TipoGasto":            contrato.TipoGasto,
							"TemaGastoInversion":   contrato.TemaGastoInversion,
							"OrigenPresupueso":     contrato.OrigenPresupueso,
							"OrigenRecursos":       contrato.OrigenRecursos,
							"TipoMoneda":           contrato.TipoMoneda,
							"TipoControl":          contrato.TipoControl,
							"Observaciones":        contrato.Observaciones,
							"Supervisor": map[string]interface{}{
								"Id":                    sup.Id,
								"Nombre":                sup.Nombre,
								"Documento":             sup.Documento,
								"Cargo":                 sup.Cargo,
								"SedeSupervisor":        sup.SedeSupervisor,
								"DependenciaSupervisor": sup.DependenciaSupervisor,
								"Tipo":                  sup.Tipo,
								"Estado":                sup.Estado,
								"DigitoVerificacion":    sup.DigitoVerificacion,
								"FechaInicio":           sup.FechaInicio,
								"FechaFin":              sup.FechaFin,
								"CargoId": map[string]interface{}{
									"Id": sup.CargoId.Id,
								},
							},
							"ClaseContratista": contrato.ClaseContratista,
							"TipoContrato": map[string]interface{}{
								"Id":           6,
								"TipoContrato": "Contrato de Prestación de Servicios Profesionales o Apoyo a la Gestión",
								"Estado":       true,
							},
							"LugarEjecucion": map[string]interface{}{
								"Id":          4,
								"Direccion":   "CALLE 40 A No 13-09",
								"Sede":        "00IP",
								"Dependencia": "DEP39",
								"Ciudad":      96,
							},
						}
						fmt.Println(contratoGeneral)
						// If modificacion_vinculacion
						s := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlCrudResoluciones") + "/" + beego.AppConfig.String("NscrudResoluciones") + "/modificacion_vinculacion?query=VinculacionDocenteRegistradaId.Id:" + strconv.Itoa(v.Id)
						if responseG, err := GetJsonTest(s, &respuesta_peticion); err == nil && responseG == 200 {
							LimpiezaRespuestaRefactor(respuesta_peticion, &modVin)
							fmt.Println("1")
							fmt.Println(respuesta_peticion)
							fmt.Println(modVin)
							var actaInicioAnterior []models.ActaInicio
							vinculacionModificacion := modVin[0].VinculacionDocenteRegistradaId
							vinculacionOriginal := modVin[0].VinculacionDocenteCanceladaId
							if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+strconv.Itoa(v.ResolucionVinculacionDocenteId.Id), &respuesta_peticion); err == nil && responseG == 200 {
								LimpiezaRespuestaRefactor(respuesta_peticion, &resolucion)
							} else {
								logs.Error(err)
								outputError = map[string]interface{}{"funcion": "/ExpedirModificacion19", "err": err.Error(), "status": "502"}
								return outputError
							}
							// If get acta_inicio cancelando
							if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio?query=NumeroContrato:"+modVin[0].VinculacionDocenteCanceladaId.NumeroContrato+",Vigencia:"+strconv.Itoa(int(modVin[0].VinculacionDocenteCanceladaId.Vigencia)), &actaInicioAnterior); err == nil && responseG == 200 {
								fmt.Println("3")
								semanasIniciales := vinculacionOriginal.NumeroSemanas
								semanasModificar := vinculacionModificacion.NumeroSemanas
								horasIniciales := vinculacionOriginal.NumeroHorasSemanales
								fechaFinNuevoContrato := CalcularFechaFin(acta.FechaInicio, semanasModificar)
								horasTotales := horasIniciales + vinculacionModificacion.NumeroHorasSemanales
								// Sólo si es reducción cambia la fecha fin del acta anterior y el valor del nuevo contrato
								if resolucion.TipoResolucionId.Id == 4 {
									var aini models.ActaInicio
									aini.Id = actaInicioAnterior[0].Id
									aini.NumeroContrato = actaInicioAnterior[0].NumeroContrato
									aini.Vigencia = actaInicioAnterior[0].Vigencia
									aini.Descripcion = actaInicioAnterior[0].Descripcion
									aini.FechaInicio = actaInicioAnterior[0].FechaInicio
									aini.FechaFin = acta.FechaInicio
									fechaFinNuevoContrato = actaInicioAnterior[0].FechaFin
									beego.Info("fin nuevo ", fechaFinNuevoContrato)
									beego.Info("fin viejo", aini.FechaFin)
									// If put acta_inicio cancelando - cambia fecha fin del acta anterior por la fecha inicio escogida por el usuario
									if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(aini.Id), "PUT", &response, &aini); err == nil {
										fmt.Println("Acta anterior cancelada en la fecha indicada")
									} else {
										fmt.Println("He fallado un poquito en If put acta_inicio cancelando, solucioname!!!", err)
										logs.Error(aini)
										outputError = map[string]interface{}{"funcion": "/ExpedirModificacion18", "err": err.Error(), "status": "502"}
										return outputError
									}
									// Calcula el valor del nuevo contrato con base en las semanas desde la fecha inicio escogida hasta la nueva fecha fin y las nuevas horas
									semanasTranscurridasDecimal := (acta.FechaInicio.Sub(actaInicioAnterior[0].FechaInicio).Hours()) / 24 / 30 * 4 // cálculo con base en meses de 30 días y 4 semanas
									semanasTranscurridas, decimal := math.Modf(semanasTranscurridasDecimal)
									if decimal > 0 {
										semanasTranscurridas = semanasTranscurridas + 1
									}
									var semanasTranscurridasInt = int(semanasTranscurridas)
									semanasRestantes := semanasIniciales - semanasTranscurridasInt - semanasModificar
									horasTotales = horasIniciales - vinculacionModificacion.NumeroHorasSemanales
									var vinc [1]models.VinculacionDocente
									vinc[0] = models.VinculacionDocente{
										ResolucionVinculacionDocenteId: &models.ResolucionVinculacionDocente{Id: m.IdResolucion},
										PersonaId:                      v.PersonaId,
										NumeroHorasSemanales:           horasTotales,
										NumeroSemanas:                  semanasModificar,
										DedicacionId:                   v.DedicacionId,
										ProyectoCurricularId:           v.ProyectoCurricularId,
										Categoria:                      v.Categoria,
										Dedicacion:                     v.Dedicacion,
										NivelAcademico:                 v.NivelAcademico,
										Vigencia:                       v.Vigencia,
										Disponibilidad:                 v.Disponibilidad,
									}
									salario, err := CalcularValorContratoReduccion(vinc, semanasRestantes, horasIniciales, v.NivelAcademico)
									if err != nil {
										fmt.Println("He fallado en cálculo del contrato reducción, solucioname!!!", err)
										//outputError = map[string]interface{}{"funcion": "/ExpedirModificacion17", "err": err.Error(), "status": "502"}
										return err
									}
									// Si es de posgrado calcula el valor que se le ha pagado hasta la fecha de inicio y se resta del total que debe quedar con la reducción
									if v.NivelAcademico == "POSGRADO" {
										diasOriginales, _ := math.Modf((actaInicioAnterior[0].FechaFin.Sub(actaInicioAnterior[0].FechaInicio).Hours()) / 24)
										diasTranscurridos, _ := math.Modf((acta.FechaInicio.Sub(actaInicioAnterior[0].FechaInicio).Hours()) / 24)
										valorDiario := vinculacionOriginal.ValorContrato / diasOriginales
										valorPagado := valorDiario * diasTranscurridos
										salario = salario - valorPagado
									}
									contrato.ValorContrato = salario
									beego.Info(contrato.ValorContrato)
								}
								if contrato.ValorContrato > 0 {
									//_, err = amazon.Raw("INSERT INTO argo.contrato_general(numero_contrato, vigencia, objeto_contrato, plazo_ejecucion, forma_pago, ordenador_gasto, sede_solicitante, dependencia_solicitante, contratista, unidad_ejecucion, valor_contrato, justificacion, descripcion_forma_pago, condiciones, unidad_ejecutora, fecha_registro, tipologia_contrato, tipo_compromiso, modalidad_seleccion, procedimiento, regimen_contratacion, tipo_gasto, tema_gasto_inversion, origen_presupueso, origen_recursos, tipo_moneda, tipo_control, observaciones, supervisor,clase_contratista, tipo_contrato, lugar_ejecucion) VALUES (?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", contrato.Id, contrato.VigenciaContrato, contrato.ObjetoContrato, contrato.PlazoEjecucion, contrato.FormaPago.Id, contrato.OrdenadorGasto, contrato.SedeSolicitante, contrato.DependenciaSolicitante, temp, contrato.UnidadEjecucion.Id, contrato.ValorContrato, contrato.Justificacion, contrato.DescripcionFormaPago, contrato.Condiciones, contrato.UnidadEjecutora, contrato.FechaRegistro.Format(time.RFC1123), contrato.TipologiaContrato, contrato.TipoCompromiso, contrato.ModalidadSeleccion, contrato.Procedimiento, contrato.RegimenContratacion, contrato.TipoGasto, contrato.TemaGastoInversion, contrato.OrigenPresupueso, contrato.OrigenRecursos, contrato.TipoMoneda, contrato.TipoControl, contrato.Observaciones, contrato.Supervisor.Id, contrato.ClaseContratista, contrato.TipoContrato.Id, contrato.LugarEjecucion.Id).Exec()
									// if contrato_general (antes insert)
									if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general", "POST", &response, contratoGeneral); err == nil {
										aux1 := contrato.Id
										aux2 := contrato.VigenciaContrato
										var ce models.ContratoEstado
										var ec models.EstadoContrato
										ce.NumeroContrato = aux1
										ce.Vigencia = aux2
										ce.FechaRegistro = time.Now()
										ec.Id = 4
										ce.Estado = &ec
										// If 4 - contrato_estado
										if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
											var ai models.ActaInicio
											ai.NumeroContrato = aux1
											ai.Vigencia = aux2
											ai.Descripcion = acta.Descripcion
											ai.FechaInicio = acta.FechaInicio
											ai.FechaFin = fechaFinNuevoContrato
											beego.Info("inicio ", ai.FechaInicio, " fin ", ai.FechaFin)
											// If 3 - Acta_inicio
											if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response, &ai); err == nil {
												var cd models.ContratoDisponibilidad
												cd.NumeroContrato = aux1
												cd.Vigencia = aux2
												cd.Estado = true
												cd.FechaRegistro = time.Now()
												// If 2.5.2 - Get disponibildad_apropiacion
												q := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudKronos") + "/" + beego.AppConfig.String("NscrudKronos") + "/disponibilidad_apropiacion/" + strconv.Itoa(v.Disponibilidad)
												if responseG, err := GetJsonTest(q, &dispoap); err == nil && responseG == 200 {
													// If 2.5.1 - Get disponibildad
													if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/"+strconv.Itoa(dispoap.Disponibilidad.Id), &disponibilidad); err == nil && responseG == 200 {
														cd.NumeroCdp = int(disponibilidad.NumeroDisponibilidad)
														cd.VigenciaCdp = int(disponibilidad.Vigencia)
														// If 2 - contrato_disponibilidad
														if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response, &cd); err == nil {
															vinculacionModificacion.PuntoSalarialId = vinculacion.VinculacionDocente.PuntoSalarialId
															vinculacionModificacion.SalarioMinimoId = vinculacion.VinculacionDocente.SalarioMinimoId
															vinculacionModificacion.NumeroContrato = aux1
															//vinculacionModificacion.NumeroContrato.Valid = true
															vinculacionModificacion.Vigencia = aux2
															//vinculacionModificacion.Vigencia.Valid = true
															// If 1 - vinculacion_docente
															if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente/"+strconv.Itoa(vinculacionModificacion.Id), "PUT", &respuesta_peticion, &vinculacionModificacion); err == nil {
																LimpiezaRespuestaRefactor(respuesta_peticion, &response)
																fmt.Println("Vinculacion docente actualizada y lista, vamos por la otra")
															} else { // If 1 - vinculacion_docente
																fmt.Println("He fallado un poquito en If 1 - vinculacion_docente, solucioname!!! ", err)
																outputError = map[string]interface{}{"funcion": "/ExpedirModificacion16", "err": err.Error(), "status": "502"}
																return outputError
															}
														} else { // If 2 - contrato_disponibilidad
															fmt.Println("He fallado un poquito en  If 2 - contrato_disponibilidad, solucioname!!!", err)
															outputError = map[string]interface{}{"funcion": "/ExpedirModificacion15", "err": err.Error(), "status": "502"}
															return outputError
														}
													} else { // If 2.5.1 - Get disponibildad
														fmt.Println("He fallado un poquito en If 2.5.1 - Get disponibildad, solucioname!!!", err)
														outputError = map[string]interface{}{"funcion": "/ExpedirModificacion14", "err": err.Error(), "status": "502"}
														return outputError
													}
												} else { // If 2.5.2 - Get disponibildad_apropiacion
													fmt.Println("He fallado un poquito en If 2.5.2 - Get disponibildad_apropiacion, solucioname!!!", err)
													outputError = map[string]interface{}{"funcion": "/ExpedirModificacion13", "err": err.Error(), "status": "502"}
													return outputError
												}
											} else { // If 3 - Acta_inicio
												//var response2 interface{}
												fmt.Println("He fallado un poquito en If 3 - Acta_inicio, solucioname!!!", err)
												outputError = map[string]interface{}{"funcion": "/ExpedirModificacion12", "err": err.Error(), "status": "502"}
												return outputError
											}
										} else { // If 4 - contrato_estado
											//var response2 interface{}
											fmt.Println("He fallado un poquito en If 4 - contrato_estado, solucioname!!!", err)
											outputError = map[string]interface{}{"funcion": "/ExpedirModificacion11", "err": err.Error(), "status": "502"}
											return outputError
										}
									} else { //If insert contrato_general
										fmt.Println("He fallado un poquito en insert contrato_general, solucioname!!!", err)
										outputError = map[string]interface{}{"funcion": "/ExpedirModificacion10", "err": err.Error(), "status": "502"}
										return outputError
									}
								}
							} else { //If get acta_inicio cancelando
								fmt.Println("He fallado un poquito en If get acta_inicio cancelando, solucioname!!!", err)
								logs.Error(actaInicioAnterior)
								outputError = map[string]interface{}{"funcion": "/ExpedirModificacion9", "err": err.Error(), "status": "502"}
								return outputError
							}
						} else { //If modificacion_vinculacion
							fmt.Println("He fallado un poquito en If modificacion_vinculacion, solucioname!!!", err)
							logs.Error(modVin)
							outputError = map[string]interface{}{"funcion": "/ExpedirModificacion8", "err": err.Error(), "status": "502"}
							return outputError
						}
					} else { // Nuevo If
						fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor nuevo, solucioname!!!", err)
						logs.Error(proveedor)
						//c.Ctx.Output.SetStatus(233)
						//err = c.Ctx.Output.Body([]byte("No existe el docente con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
						// if err != nil {
						// beego.Error(err)
						// }
						outputError = map[string]interface{}{"funcion": "/ExpedirModificacion7", "err": err.Error(), "status": "502"}
						return outputError
					}
				} else { // If 5 - Informacion_Proveedor
					fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor, solucioname!!!", err)
					logs.Error(proveedor)
					outputError = map[string]interface{}{"funcion": "/ExpedirModificacion6", "err": err.Error(), "status": "502"}
					return outputError
				}
			} else { //If 8 - Vinculacion_docente (GET)
				fmt.Println("He fallado un poquito en If 8 - Vinculacion_docente (GET), solucioname!!!", err)
				logs.Error(v)
				outputError = map[string]interface{}{"funcion": "/ExpedirModificacion5", "err": err.Error(), "status": "502"}
				return outputError
			}
		} // for vinculaciones
		var r models.Resolucion
		r.Id = m.IdResolucion
		idResolucionDVE := strconv.Itoa(m.IdResolucion)
		// If 11 - Resolucion (GET)
		if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+idResolucionDVE, &respuesta_peticion); err == nil && responseG == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &r)
			r.FechaExpedicion = m.FechaExpedicion
			// If 10 - Resolucion (PUT)
			if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &respuesta_peticion, &r); err == nil {
				LimpiezaRespuestaRefactor(respuesta_peticion, &response)
				var e models.ResolucionEstado
				var er models.EstadoResolucion
				e.ResolucionId = &r
				er.Id = 2
				e.EstadoResolucionId = &er
				//e.FechaCreacion = time.Now()
				// If 9 - Resolucion_estado
				if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion_estado", "POST", &respuesta_peticion, &e); err == nil {
					LimpiezaRespuestaRefactor(respuesta_peticion, &response)
					fmt.Println("Expedición exitosa, ahora va el commit :D")
					//c.Data["json"] = v
				} else { //If 9 - Resolucion_estado
					fmt.Println("He fallado un poquito en If 9 - Resolucion_estado, solucioname!!!", err)
					logs.Error(e)
					outputError = map[string]interface{}{"funcion": "/ExpedirModificacion4", "err": err.Error(), "status": "502"}
					return outputError
				}
			} else { //If 10 - Resolucion (PUT)
				fmt.Println("He fallado un poquito en If 10 - Resolucion (PUT), solucioname!!! ", err)
				logs.Error(r)
				outputError = map[string]interface{}{"funcion": "/ExpedirModificacion3", "err": err.Error(), "status": "502"}
				return outputError
			}
		} else { //If 11 - Resolucion (GET)
			fmt.Println("He fallado un poquito en If 11 - Resolucion (GET), solucioname!!! ", err)
			logs.Error(r)
			outputError = map[string]interface{}{"funcion": "/ExpedirModificacion2", "err": err.Error(), "status": "502"}
			return outputError
		}
	} else { //If 12 - Consecutivo contrato_general
		fmt.Println("He fallado un poquito en If 12 - Consecutivo contrato_general, solucioname!!! ", err)
		logs.Error(cdve)
		outputError = map[string]interface{}{"funcion": "/ExpedirModificacion1", "err": err.Error(), "status": "502"}
		return outputError
	}
	return
}

func Cancelar(m models.ExpedicionCancelacion) (outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/Cancelar", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	v := m.Vinculaciones
	var contratoCancelado models.ContratoCancelado
	var response interface{}
	var respuesta_peticion map[string]interface{}
	// for vinculaciones
	for _, vinculacion := range *v {
		v := vinculacion.VinculacionDocente
		idVinculacionDocente := strconv.Itoa(v.Id)
		//If vinculacion_docente (get)
		if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/vinculacion_docente/"+idVinculacionDocente, &respuesta_peticion); err == nil && responseG == 200 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &v)
			contratoCancelado.NumeroContrato = v.NumeroContrato
			contratoCancelado.Vigencia = int(v.Vigencia)
			contratoCancelado.FechaCancelacion = vinculacion.ContratoCancelado.FechaCancelacion
			contratoCancelado.MotivoCancelacion = vinculacion.ContratoCancelado.MotivoCancelacion
			contratoCancelado.Usuario = vinculacion.ContratoCancelado.Usuario
			contratoCancelado.FechaRegistro = time.Now()
			contratoCancelado.Estado = vinculacion.ContratoCancelado.Estado
			// if contrato_cancelado (post)
			if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_cancelado", "POST", &response, &contratoCancelado); err == nil {
				var ai []models.ActaInicio
				// if acta_inicio (get)
				if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio?query=NumeroContrato:"+contratoCancelado.NumeroContrato+",Vigencia:"+strconv.Itoa(contratoCancelado.Vigencia), &ai); err == nil && responseG == 200 {
					ai[0].FechaFin = CalcularFechaFin(ai[0].FechaInicio, v.NumeroSemanasNuevas)
					// if acta_inicio (put)
					if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai[0].Id), "PUT", &response, &ai[0]); err == nil {
						var ce models.ContratoEstado
						var ec models.EstadoContrato
						ce.NumeroContrato = contratoCancelado.NumeroContrato
						ce.Vigencia = contratoCancelado.Vigencia
						ce.FechaRegistro = time.Now()
						ec.Id = 7
						ce.Estado = &ec
						// If contrato_estado (post)
						if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
							var r models.Resolucion
							r.Id = m.IdResolucion
							idResolucionDVE := strconv.Itoa(m.IdResolucion)
							//If  Resolucion (GET)
							if responseG, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+idResolucionDVE, &r); err == nil && responseG == 200 {
								LimpiezaRespuestaRefactor(respuesta_peticion, &r)
								r.FechaExpedicion = m.FechaExpedicion
								//If Resolucion (PUT)
								if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &respuesta_peticion, &r); err == nil {
									LimpiezaRespuestaRefactor(respuesta_peticion, &response)
									var e models.ResolucionEstado
									var er models.EstadoResolucion
									e.ResolucionId = &models.Resolucion{Id: m.IdResolucion}
									er.Id = 2
									e.EstadoResolucionId = &er
									//e.FechaCreacion = time.Now()
									//If  Resolucion_estado (post)
									if err := SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlCrudResoluciones")+"/"+beego.AppConfig.String("NscrudResoluciones")+"/resolucion_estado", "POST", &respuesta_peticion, &e); err == nil {
										LimpiezaRespuestaRefactor(respuesta_peticion, &response)
										fmt.Println("Expedición exitosa, ahora va el commit :D")
									} else { //If  Resolucion_estado (post)
										//var response2 interface{}
										fmt.Println("He fallado un poquito en If  Resolucion_estado (post), solucioname!!! ", err)
										logs.Error(e)
										outputError = map[string]interface{}{"funcion": "/Cancelar8", "err": err.Error(), "status": "502"}
										return outputError
									}
								} else { //If Resolucion (PUT)
									fmt.Println("He fallado un poquito en If Resolucion (PUT), solucioname!!! ", err)
									logs.Error(r)
									outputError = map[string]interface{}{"funcion": "/Cancelar7", "err": err.Error(), "status": "502"}
									return outputError
								}
							} else { // If Resolucion (GET)
								fmt.Println("He fallado un poquito en If Resolucion (PUT), solucioname!!! ", err)
								logs.Error(r)
								outputError = map[string]interface{}{"funcion": "/Cancelar6", "err": err.Error(), "status": "502"}
								return outputError
							}
						} else { // If contrato_estado (post)
							//var response2 interface{}
							fmt.Println("He fallado un poquito en If Resolucion (GET), solucioname!!! ", err)
							logs.Error(ce)
							outputError = map[string]interface{}{"funcion": "/Cancelar5", "err": err.Error(), "status": "502"}
							return outputError
						}
					} else { // If acta_inicio (put)
						fmt.Println("He fallado un poquito en If Acta_Inicio (PUT), solucioname!!! ", err)
						logs.Error(ai[0])
						outputError = map[string]interface{}{"funcion": "/Cancelar4", "err": err.Error(), "status": "502"}
						return outputError
					}
				} else { // if acta_inicio (get)
					fmt.Println("He fallado un poquito en if acta_inicio (GET), solucioname!!! ", err)
					logs.Error(ai)
					outputError = map[string]interface{}{"funcion": "/Cancelar3", "err": err.Error(), "status": "502"}
					return outputError
				}
			} else { // if contrato_cancelado (post)
				//var response2 interface{}
				fmt.Println("He fallado un poquito en if contrato_cancelado (post), solucioname!!! ", err)
				logs.Error(contratoCancelado)
				outputError = map[string]interface{}{"funcion": "/Cancelar2", "err": err.Error(), "status": "502"}
				return outputError
			}
		} else {
			//If vinculacion_docente (get)
			fmt.Println("He fallado un poquito en If vinculacion_docente (get), solucioname!!! ", err)
			logs.Error(v)
			outputError = map[string]interface{}{"funcion": "/Cancelar1", "err": err.Error(), "status": "502"}
			return outputError
		}
	}
	return
}
