package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/resoluciones_docentes_mid/helpers"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

// ExpedirResolucionController operations for ExpedirResolucion
type ExpedirResolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ExpedirResolucionController) URLMapping() {
	c.Mapping("Expedir", c.Expedir)
	c.Mapping("ValidarDatosExpedicion", c.ValidarDatosExpedicion)
	c.Mapping("ExpedirModificacion", c.ExpedirModificacion)

}

// Expedir ...
// @Title Expedir
// @Description create Expedir
// @Param	body		body 	[]models.ExpedicionResolucion	true		"body for Expedicion Resolucion content"
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedir [post]
func (c *ExpedirResolucionController) Expedir() {
	/*amazon := orm.NewOrm()
	flyway := orm.NewOrm()
	err := amazon.Using("amazonAdmin")
	if err != nil {
		beego.Error(err)
	}
	err = flyway.Using("flywayAdmin")
	if err != nil {
		beego.Error(err)
	}*/
	var m models.ExpedicionResolucion
	var temp int
	var cdve int
	var proveedor []models.InformacionProveedor
	var tipoCon models.TipoContrato
	var disponibilidad models.Disponibilidad
	var dispoap models.DisponibilidadApropiacion
	var response interface{}
	//var CargoId models.CargoSupervisorTemporal
	vigencia, _, _ := time.Now().Date()
	//If 13 - Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		//If 12 - Consecutivo contrato_general
		if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil {
			numeroContratos := cdve
			// for vinculaciones
			for _, vinculacion := range *v {
				numeroContratos = numeroContratos + 1
				v := vinculacion.VinculacionDocente
				idvinculaciondocente := strconv.Itoa(v.Id)
				//if 8 - Vinculacion_docente (GET)
				fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente)
				if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v); err == nil {
					contrato := vinculacion.ContratoGeneral
					fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/tipo_contrato/"+ strconv.Itoa(contrato.TipoContrato.Id))
					if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/tipo_contrato/"+ strconv.Itoa(contrato.TipoContrato.Id), &tipoCon); err == nil {
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
						sup = helpers.SupervisorActual(v.IdResolucion.Id)
						contrato.Supervisor = &sup
						contrato.Condiciones = "Sin condiciones"
						
						
						// If 5 - Informacion_Proveedor
						if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil {
							if proveedor != nil { //Nuevo If
								temp = proveedor[0].Id
								contratoGeneral := make(map[string]interface{})
								contratoGeneral = map[string]interface{}{
									"Id": contrato.Id,
									"VigenciaContrato": contrato.VigenciaContrato,
									"ObjetoContrato": contrato.ObjetoContrato,
									"PlazoEjecucion": contrato.PlazoEjecucion,
									"FormaPago": map[string]interface{}{
										"Id": 240,
										"Descripcion": "TRANSACCIÓN",
										"CodigoContraloria": "'",
            							"EstadoRegistro": true,
            							"FechaRegistro": "2016-10-25T00:00:00Z",
									},
									"OrdenadorGasto": contrato.OrdenadorGasto,
									"SedeSolicitante": contrato.SedeSolicitante,
									"DependenciaSolicitante": contrato.DependenciaSolicitante,
									"Contratista": temp,
									"UnidadEjecucion": map[string]interface{}{
										"Id": 269,
            							"Descripcion": "Semana(s)",
            							"CodigoContraloria": "'",
            							"EstadoRegistro": true,
            							"FechaRegistro": "2018-03-20T00:00:00Z",
									},
									"ValorContrato": int(contrato.ValorContrato),
									"Justificacion": contrato.Justificacion,
									"DescripcionFormaPago": contrato.DescripcionFormaPago,
									"Condiciones": contrato.Condiciones,
									"UnidadEjecutora": contrato.UnidadEjecutora,
									"FechaRegistro": contrato.FechaRegistro.Format(time.RFC3339),
									"TipologiaContrato": contrato.TipologiaContrato,
									"TipoCompromiso": contrato.TipoCompromiso,
									"ModalidadSeleccion": contrato.ModalidadSeleccion,
									"Procedimiento": contrato.Procedimiento,
									"RegimenContratacion": contrato.RegimenContratacion,
									"TipoGasto": contrato.TipoGasto,
									"TemaGastoInversion": contrato.TemaGastoInversion,
									"OrigenPresupueso": contrato.OrigenPresupueso,
									"OrigenRecursos": contrato.OrigenRecursos,
									"TipoMoneda": contrato.TipoMoneda,
									"TipoControl": contrato.TipoControl,
									"Observaciones": contrato.Observaciones,
									"Supervisor": map[string]interface{}{
										"Id": sup.Id,
										"Nombre": sup.Nombre,
										"Documento": sup.Documento,
										"Cargo": sup.Cargo,
            							"SedeSupervisor": sup.SedeSupervisor,
            							"DependenciaSupervisor": sup.DependenciaSupervisor,
            							"Tipo": sup.Tipo,
            							"Estado": sup.Estado,
            							"DigitoVerificacion": sup.DigitoVerificacion,
            							"FechaInicio": sup.FechaInicio,
            							"FechaFin": sup.FechaFin,
										"CargoId": map[string]interface{}{
											"Id": sup.CargoId.Id,
										},
									},
									"ClaseContratista": contrato.ClaseContratista,
									"TipoContrato": map[string]interface{}{
										"Id": 6,
        								"TipoContrato": "Contrato de Prestación de Servicios Profesionales o Apoyo a la Gestión",
        								"Estado": true,
									},
									"LugarEjecucion": map[string]interface{}{
										"Id": 4,
            							"Direccion": "CALLE 40 A No 13-09",
            							"Sede": "00IP",
            							"Dependencia": "DEP39",
            							"Ciudad": 96,
									},
								}
								fmt.Println(contratoGeneral)
								fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general")
								//var response1 models.ContratoGeneral
								if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general", "POST", &response, contratoGeneral); err == nil {
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
									if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response2, &ce); err == nil {
										//var id2 = response
										a := vinculacion.VinculacionDocente
										var ai models.ActaInicio
										ai.NumeroContrato = aux1
										ai.Vigencia = aux2
										ai.Descripcion = acta.Descripcion
										ai.FechaInicio = acta.FechaInicio
										ai.FechaFin = acta.FechaFin
										ai.FechaFin = helpers.CalcularFechaFin(acta.FechaInicio, a.NumeroSemanas)
										// If 3 - Acta_inicio creación
										var response3 models.ActaInicio
										if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response3, &ai); err == nil {
											//var id3 = response
											var cd models.ContratoDisponibilidad
											cd.NumeroContrato = aux1
											cd.Vigencia = aux2
											cd.Estado = true
											cd.FechaRegistro = time.Now()
											// If 2.5.2 - Get disponibildad_apropiacion
											fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/"+strconv.Itoa(v.Disponibilidad))
											if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/"+strconv.Itoa(v.Disponibilidad), &dispoap); err == nil {
												// If 2.5.1 - Get disponibildad
												if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/"+strconv.Itoa(dispoap.Disponibilidad.Id), &disponibilidad); err == nil {
													cd.NumeroCdp = int(disponibilidad.NumeroDisponibilidad)
													cd.VigenciaCdp = int(disponibilidad.Vigencia)
													// If 2 - contrato_disponibilidad
													var response4 models.ContratoDisponibilidad
													if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response4, &cd); err == nil {
														a.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
														a.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
														v := a
														v.NumeroContrato.String = aux1
														v.NumeroContrato.Valid = true
														v.Vigencia.Int64 = int64(aux2)
														v.Vigencia.Valid = true
														v.FechaInicio = acta.FechaInicio
														fmt.Println("CD es: ", response4.Id)
														fmt.Println("AI es: ", response3.Id)
														fmt.Println("CE es: ", response2.Id)
														fmt.Println("CONTRATO GENERAL es: ", contrato.Id)
														// If 1 - vinculacion_docente
														if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.Id), "PUT", &response, &v); err == nil {
															
															fmt.Println()
															fmt.Println("Vinculacion docente actualizada y lista, vamos por la otra")
														} else { // If 1 - vinculacion_docente
															// var response5 interface{}
															// fmt.Println("He fallado un poquito en If 1 - vinculacion_docente, solucioname!!! ", err)
															// fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/"+strconv.Itoa(response4.Id))
															// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/"+strconv.Itoa(response4.Id), "DELETE", &response5, nil)
															// fmt.Println("BORRADO 1: ", response5)
															// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(response3.Id), "DELETE", &response5, nil)
															// fmt.Println("BORRADO 2: ", response5)
															// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/"+strconv.Itoa(response2.Id), "DELETE", &response5, nil)
															// fmt.Println("BORRADO 3: ", response5)
															// erras := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/"+contrato.Id, "DELETE", &response5, nil)
															// fmt.Println("BORRADO 4: ", response5)
															// fmt.Println(erras)
															logs.Error(v)
															c.Data["system"] = v
															c.Abort("400")
														}
													} else { // If 2 - contrato_disponibilidad
														//var response2 interface{}
														fmt.Println("He fallado un poquito en  If 2 - contrato_disponibilidad, solucioname!!!", err)
														// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/"+strconv.Itoa(cd.Id), "DELETE", &response2, nil)
														// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai.Id), "DELETE", &response2, nil)
														// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/"+strconv.Itoa(ce.Id), "DELETE", &response2, nil)
														logs.Error(cd)
														c.Data["system"] = cd
														c.Abort("400")
													}
												} else { // If 2.5.1 - Get disponibildad
													fmt.Println("He fallado un poquito en If 2.5.1 - Get disponibildad, solucioname!!!", err)
													logs.Error(disponibilidad)
													c.Data["system"] = disponibilidad
													c.Abort("404")
												}
											} else { // If 2.5.2 - Get disponibildad_apropiacion
												fmt.Println("He fallado un poquito en If 2.5.2 - Get disponibildad_apropiacion, solucioname!!!", err)
												logs.Error(dispoap)
												c.Data["system"] = dispoap
												c.Abort("404")
											}
										} else { // If 3 - Acta_inicio
											//var response2 interface{}
											fmt.Println("He fallado un poquito en If 3 - Acta_inicio, solucioname!!!", err)
											// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai.Id), "DELETE", &response2, nil)
											// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/"+strconv.Itoa(ce.Id), "DELETE", &response2, nil)
											logs.Error(ai)
											c.Data["system"] = ai
											c.Abort("400")
										}
									} else { // If 4 - contrato_estado
										//var response2 interface{}
										fmt.Println("He fallado un poquito en If 4 - contrato_estado, solucioname!!!", err)
										// helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/"+strconv.Itoa(ce.Id), "DELETE", &response2, nil)
										logs.Error(ce)
										c.Data["system"] = ce
										c.Abort("400")
									}
								} else { //If insert contrato_general
									fmt.Println("He fallado un poquito en insert contrato_general, solucioname!!!", err)
									/*	err = amazon.Rollback()
									if err != nil {
										beego.Error(err)
									}
								err = flyway.Rollback()
								if err != nil {
									beego.Error(err)
									}*/
									return
								}
							} else { // Nuevo If
								fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor nuevo, solucioname!!!", err)
								// err = amazon.Rollback()
								// if err != nil {
								// 	beego.Error(err)
								// }
								// err = flyway.Rollback()
								// if err != nil {
								// 	beego.Error(err)
								// }
								c.Ctx.Output.SetStatus(233)
								err = c.Ctx.Output.Body([]byte("No existe el docente con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
								if err != nil {
									beego.Error(err)
								}
								return
							}
						} else { // If 5 - Informacion_Proveedor
							fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor, solucioname!!!", err)
							// err = amazon.Rollback()
							// if err != nil {
							// 	beego.Error(err)
							// }
							// err = flyway.Rollback()
							// if err != nil {
							// 	beego.Error(err)
							// }
							return
						}
					} else{
						fmt.Println("error")
					}
				} else { //If 8 - Vinculacion_docente (GET)
					fmt.Println("He fallado un poquito en If 8 - Vinculacion_docente (GET), solucioname!!!", err)
					logs.Error(v)
					c.Data["system"] = v
					c.Abort("404")
				}
			} // for vinculaciones
			var r models.Resolucion
			r.Id = m.IdResolucion
			idResolucionDVE := strconv.Itoa(m.IdResolucion)
			//If 11 - Resolucion (GET)
			fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE)
			if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE, &r); err == nil {
				r.FechaExpedicion = m.FechaExpedicion
				//If 10 - Resolucion (PUT)
				if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &response, &r); err == nil {
					var e models.ResolucionEstado
					var er models.EstadoResolucion
					e.Resolucion = &r
					er.Id = 2
					e.Estado = &er
					e.FechaRegistro = time.Now()
					//If 9 - Resolucion_estado
					if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &response, &e); err == nil {
						fmt.Println("Expedición exitosa, ahora va el commit :D")
						c.Data["json"] = v
					} else { //If 9 - Resolucion_estado
						fmt.Println("He fallado un poquito en If 9 - Resolucion_estado, solucioname!!!", err)
						helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado/"+strconv.Itoa(e.Id), "DELETE", &response, &e)
						logs.Error(e)
						c.Data["system"] = e
						c.Abort("400")
					}
				} else { //If 10 - Resolucion (PUT)
					fmt.Println("He fallado un poquito en If 10 - Resolucion (PUT), solucioname!!! ", err)
					logs.Error(r)
					c.Data["system"] = r
					c.Abort("400")
				}
			} else { //If 11 - Resolucion (GET)
				fmt.Println("He fallado un poquito en If 11 - Resolucion (GET), solucioname!!! ", err)
				logs.Error(r)
				c.Data["system"] = r
				c.Abort("404")
			}
		} else { //If 12 - Consecutivo contrato_general
			fmt.Println("He fallado un poquito en If 12 - Consecutivo contrato_general, solucioname!!! ", err)
			logs.Error(cdve)
			c.Data["system"] = cdve
			c.Abort("404")
		}

	} else { //If 13 - Unmarshal
		fmt.Println("He fallado un poquito en If 13 - Unmarshal, solucioname!!! ", err)
		logs.Error(m)
		c.Data["system"] = m
		c.Abort("404")
	}

	/*err = amazon.Commit()
	if err != nil {
		fmt.Println(err)
	}
	err = flyway.Commit()
	if err != nil {
		fmt.Println(err)
	}*/
	c.ServeJSON()
}

// ExpedirResolucionController ...
// @Title ValidarDatosExpedicion
// @Description create ValidarDatosExpedicion
// @Success 201 {int}
// @Failure 403 body is empty
// @router /validar_datos_expedicion [post]
func (c *ExpedirResolucionController) ValidarDatosExpedicion() {
	var m models.ExpedicionResolucion

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	v := m.Vinculaciones
	beego.Info(v)

	for _, vinculacion := range *v {
		v := vinculacion.VinculacionDocente
		idvinculaciondocente := strconv.Itoa(v.Id)

		err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v)
		if err != nil {
			beego.Error("Previnculación no valida", err)
			logs.Error(v)
			c.Data["system"] = v
			c.Abort("233")
		}

		contrato := vinculacion.ContratoGeneral
		var proveedor []models.InformacionProveedor

		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor)
		if err != nil {
			beego.Error("Docente no válido en Ágora, se encuentra identificado con el documento número ", strconv.Itoa(contrato.Contratista), err)
			logs.Error(proveedor)
			c.Data["system"] = proveedor
			c.Abort("233")
		}

		if proveedor == nil {
			beego.Error("No existe el docente con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			logs.Error(proveedor)
			c.Data["system"] = proveedor
			c.Abort("233")
		}

		var dispoap []models.DisponibilidadApropiacion

		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/?query=Id:"+strconv.Itoa(v.Disponibilidad), &dispoap)
		if err != nil {
			beego.Error("Disponibilidad no válida asociada al docente identificado con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			logs.Error(dispoap)
			c.Data["system"] = dispoap
			c.Abort("233")
		}

		if dispoap == nil {
			beego.Error("Disponibilidad no válida asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			logs.Error(dispoap)
			c.Data["system"] = dispoap
			c.Abort("233")
		}

		var proycur []models.Dependencia

		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/?query=Id:"+strconv.Itoa(v.IdProyectoCurricular), &proycur)
		if err != nil {
			beego.Error("Dependencia incorrectamente homologada asociada al docente identificado con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			logs.Error(proycur)
			c.Data["system"] = proycur
			c.Abort("233")
		}

		if proycur == nil {
			beego.Error("Dependencia incorrectamente homologada asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			logs.Error(proycur)
			c.Data["system"] = proycur
			c.Abort("233")
		}
		beego.Info(proycur)

	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()
}

// ExpedirModificacion ...
// @Title ExpedirModificacion
// @Description create ExpedirModificacion
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedirModificacion [post]
func (c *ExpedirResolucionController) ExpedirModificacion() {
	amazon := orm.NewOrm()
	flyway := orm.NewOrm()
	err := amazon.Using("amazonAdmin")
	if err != nil {
		beego.Error(err)
	}
	err = flyway.Using("flywayAdmin")
	if err != nil {
		beego.Error(err)
	}
	var m models.ExpedicionResolucion
	var temp int
	var cdve int
	var proveedor []models.InformacionProveedor
	var disponibilidad models.Disponibilidad
	var dispoap models.DisponibilidadApropiacion
	var modVin []models.ModificacionVinculacion
	var response interface{}
	var resolucion models.Resolucion
	vigencia, _, _ := time.Now().Date()
	// If 13 - Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		// If 12 - Consecutivo contrato_general
		if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil {
			numeroContratos := cdve
			// for vinculaciones
			for _, vinculacion := range *v {
				numeroContratos = numeroContratos + 1
				v := vinculacion.VinculacionDocente
				idvinculaciondocente := strconv.Itoa(v.Id)
				// if 8 - Vinculacion_docente (GET)
				if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v); err == nil {
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
					sup = helpers.SupervisorActual(v.IdResolucion.Id)
					contrato.Supervisor = &sup
					contrato.Condiciones = "Sin condiciones"
					// If 5 - Informacion_Proveedor
					if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil {
						if proveedor != nil { //Nuevo If
							temp = proveedor[0].Id

							// If modificacion_vinculacion
							if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(v.Id), &modVin); err == nil {
								var actaInicioAnterior []models.ActaInicio
								vinculacionModificacion := modVin[0].VinculacionDocenteRegistrada
								vinculacionOriginal := modVin[0].VinculacionDocenteCancelada
								err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(v.IdResolucion.Id), &resolucion)
								if err != nil {
									logs.Error(err)
									c.Data["system"] = err
									c.Abort("400")
								}
								// If get acta_inicio cancelando
								if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+modVin[0].VinculacionDocenteCancelada.NumeroContrato.String+",Vigencia:"+strconv.Itoa(int(modVin[0].VinculacionDocenteCancelada.Vigencia.Int64)), &actaInicioAnterior); err == nil {
									semanasIniciales := vinculacionOriginal.NumeroSemanas
									semanasModificar := vinculacionModificacion.NumeroSemanas
									horasIniciales := vinculacionOriginal.NumeroHorasSemanales
									fechaFinNuevoContrato := helpers.CalcularFechaFin(acta.FechaInicio, semanasModificar)
									horasTotales := horasIniciales + vinculacionModificacion.NumeroHorasSemanales
									// Sólo si es reducción cambia la fecha fin del acta anterior y el valor del nuevo contrato
									if resolucion.IdTipoResolucion.Id == 4 {
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
										if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(aini.Id), "PUT", &response, &aini); err == nil {
											fmt.Println("Acta anterior cancelada en la fecha indicada")
										} else {
											fmt.Println("He fallado un poquito en If put acta_inicio cancelando, solucioname!!!", err)
											logs.Error(aini)
											c.Data["system"] = aini
											c.Abort("400")
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
											IdResolucion:         &models.ResolucionVinculacionDocente{Id: m.IdResolucion},
											IdPersona:            v.IdPersona,
											NumeroHorasSemanales: horasTotales,
											NumeroSemanas:        semanasModificar,
											IdDedicacion:         v.IdDedicacion,
											IdProyectoCurricular: v.IdProyectoCurricular,
											Categoria:            v.Categoria,
											Dedicacion:           v.Dedicacion,
											NivelAcademico:       v.NivelAcademico,
											Vigencia:             v.Vigencia,
											Disponibilidad:       v.Disponibilidad,
										}
										salario, err := CalcularValorContratoReduccion(vinc, semanasRestantes, horasIniciales, v.NivelAcademico)
										if err != nil {
											fmt.Println("He fallado en cálculo del contrato reducción, solucioname!!!", err)
											err = amazon.Rollback()
											if err != nil {
												beego.Error(err)
											}
											err = flyway.Rollback()
											if err != nil {
												beego.Error(err)
											}
											return
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
										_, err = amazon.Raw("INSERT INTO argo.contrato_general(numero_contrato, vigencia, objeto_contrato, plazo_ejecucion, forma_pago, ordenador_gasto, sede_solicitante, dependencia_solicitante, contratista, unidad_ejecucion, valor_contrato, justificacion, descripcion_forma_pago, condiciones, unidad_ejecutora, fecha_registro, tipologia_contrato, tipo_compromiso, modalidad_seleccion, procedimiento, regimen_contratacion, tipo_gasto, tema_gasto_inversion, origen_presupueso, origen_recursos, tipo_moneda, tipo_control, observaciones, supervisor,clase_contratista, tipo_contrato, lugar_ejecucion) VALUES (?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", contrato.Id, contrato.VigenciaContrato, contrato.ObjetoContrato, contrato.PlazoEjecucion, contrato.FormaPago.Id, contrato.OrdenadorGasto, contrato.SedeSolicitante, contrato.DependenciaSolicitante, temp, contrato.UnidadEjecucion.Id, contrato.ValorContrato, contrato.Justificacion, contrato.DescripcionFormaPago, contrato.Condiciones, contrato.UnidadEjecutora, contrato.FechaRegistro.Format(time.RFC1123), contrato.TipologiaContrato, contrato.TipoCompromiso, contrato.ModalidadSeleccion, contrato.Procedimiento, contrato.RegimenContratacion, contrato.TipoGasto, contrato.TemaGastoInversion, contrato.OrigenPresupueso, contrato.OrigenRecursos, contrato.TipoMoneda, contrato.TipoControl, contrato.Observaciones, contrato.Supervisor.Id, contrato.ClaseContratista, contrato.TipoContrato.Id, contrato.LugarEjecucion.Id).Exec()
										// If insert contrato_general
										if err == nil {

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
											if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
												var ai models.ActaInicio
												ai.NumeroContrato = aux1
												ai.Vigencia = aux2
												ai.Descripcion = acta.Descripcion
												ai.FechaInicio = acta.FechaInicio
												ai.FechaFin = fechaFinNuevoContrato
												beego.Info("inicio ", ai.FechaInicio, " fin ", ai.FechaFin)
												// If 3 - Acta_inicio
												if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response, &ai); err == nil {
													var cd models.ContratoDisponibilidad
													cd.NumeroContrato = aux1
													cd.Vigencia = aux2
													cd.Estado = true
													cd.FechaRegistro = time.Now()
													// If 2.5.2 - Get disponibildad_apropiacion
													if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/"+strconv.Itoa(v.Disponibilidad), &dispoap); err == nil {
														// If 2.5.1 - Get disponibildad
														if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/"+strconv.Itoa(dispoap.Disponibilidad.Id), &disponibilidad); err == nil {
															cd.NumeroCdp = int(disponibilidad.NumeroDisponibilidad)
															cd.VigenciaCdp = int(disponibilidad.Vigencia)
															// If 2 - contrato_disponibilidad
															if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response, &cd); err == nil {
																vinculacionModificacion.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
																vinculacionModificacion.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
																vinculacionModificacion.NumeroContrato.String = aux1
																vinculacionModificacion.NumeroContrato.Valid = true
																vinculacionModificacion.Vigencia.Int64 = int64(aux2)
																vinculacionModificacion.Vigencia.Valid = true
																// If 1 - vinculacion_docente
																if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacionModificacion.Id), "PUT", &response, &vinculacionModificacion); err == nil {
																	fmt.Println("Vinculacion docente actualizada y lista, vamos por la otra")
																} else { // If 1 - vinculacion_docente
																	fmt.Println("He fallado un poquito en If 1 - vinculacion_docente, solucioname!!! ", err)
																	err = amazon.Rollback()
																	if err != nil {
																		beego.Error(err)
																	}
																	err = flyway.Rollback()
																	if err != nil {
																		beego.Error(err)
																	}
																	return
																}
															} else { // If 2 - contrato_disponibilidad
																var response2 interface{}
																fmt.Println("He fallado un poquito en  If 2 - contrato_disponibilidad, solucioname!!!", err)
																helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/"+strconv.Itoa(cd.Id), "DELTE", &response2, nil)
																helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai.Id), "DELTE", &response2, nil)
																helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ce.Id), "DELTE", &response2, nil)
																// err = amazon.Rollback()
																// if err != nil {
																// 	beego.Error(err)
																// }
																// err = flyway.Rollback()
																// if err != nil {
																// 	beego.Error(err)
																// }
																return
															}
														} else { // If 2.5.1 - Get disponibildad
															fmt.Println("He fallado un poquito en If 2.5.1 - Get disponibildad, solucioname!!!", err)
															// err = amazon.Rollback()
															// if err != nil {
															// 	beego.Error(err)
															// }
															// err = flyway.Rollback()
															// if err != nil {
															// 	beego.Error(err)
															// }
															return
														}
													} else { // If 2.5.2 - Get disponibildad_apropiacion
														fmt.Println("He fallado un poquito en If 2.5.2 - Get disponibildad_apropiacion, solucioname!!!", err)
														// err = amazon.Rollback()
														// if err != nil {
														// 	beego.Error(err)
														// }
														// err = flyway.Rollback()
														// if err != nil {
														// 	beego.Error(err)
														// }
														return
													}
												} else { // If 3 - Acta_inicio
													var response2 interface{}
													fmt.Println("He fallado un poquito en If 3 - Acta_inicio, solucioname!!!", err)
													helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai.Id), "DELTE", &response2, nil)
													helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ce.Id), "DELTE", &response2, nil)
													err = amazon.Rollback()
													if err != nil {
														beego.Error(err)
													}
													err = flyway.Rollback()
													if err != nil {
														beego.Error(err)
													}
													return
												}
											} else { // If 4 - contrato_estado
												var response2 interface{}
												fmt.Println("He fallado un poquito en If 4 - contrato_estado, solucioname!!!", err)
												helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ce.Id), "DELTE", &response2, nil)
												err = amazon.Rollback()
												if err != nil {
													beego.Error(err)
												}
												err = flyway.Rollback()
												if err != nil {
													beego.Error(err)
												}
												return
											}
										} else { //If insert contrato_general
											fmt.Println("He fallado un poquito en insert contrato_general, solucioname!!!", err)
											err = amazon.Rollback()
											if err != nil {
												beego.Error(err)
											}
											err = flyway.Rollback()
											if err != nil {
												beego.Error(err)
											}
											return
										}
									}
								} else { //If get acta_inicio cancelando
									fmt.Println("He fallado un poquito en If get acta_inicio cancelando, solucioname!!!", err)
									logs.Error(actaInicioAnterior)
									c.Data["system"] = actaInicioAnterior
									c.Abort("400")
								}
							} else { //If modificacion_vinculacion
								fmt.Println("He fallado un poquito en If modificacion_vinculacion, solucioname!!!", err)
								logs.Error(modVin)
								c.Data["system"] = modVin
								c.Abort("404")
							}
						} else { // Nuevo If
							fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor nuevo, solucioname!!!", err)
							logs.Error(proveedor)
							c.Data["system"] = proveedor
							c.Abort("404")
							c.Ctx.Output.SetStatus(233)
							err = c.Ctx.Output.Body([]byte("No existe el docente con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
							if err != nil {
								beego.Error(err)
							}
							return
						}
					} else { // If 5 - Informacion_Proveedor
						fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor, solucioname!!!", err)
						logs.Error(proveedor)
						c.Data["system"] = proveedor
						c.Abort("404")
					}
				} else { //If 8 - Vinculacion_docente (GET)
					fmt.Println("He fallado un poquito en If 8 - Vinculacion_docente (GET), solucioname!!!", err)
					logs.Error(v)
					c.Data["system"] = v
					c.Abort("404")
				}
			} // for vinculaciones
			var r models.Resolucion
			r.Id = m.IdResolucion
			idResolucionDVE := strconv.Itoa(m.IdResolucion)
			// If 11 - Resolucion (GET)
			if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE, &r); err == nil {
				r.FechaExpedicion = m.FechaExpedicion
				// If 10 - Resolucion (PUT)
				if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &response, &r); err == nil {
					var e models.ResolucionEstado
					var er models.EstadoResolucion
					e.Resolucion = &r
					er.Id = 2
					e.Estado = &er
					e.FechaRegistro = time.Now()
					// If 9 - Resolucion_estado
					if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &response, &e); err == nil {
						fmt.Println("Expedición exitosa, ahora va el commit :D")
						c.Data["json"] = v
					} else { //If 9 - Resolucion_estado
						fmt.Println("He fallado un poquito en If 9 - Resolucion_estado, solucioname!!!", err)
						logs.Error(e)
						c.Data["system"] = e
						c.Abort("400")
					}
				} else { //If 10 - Resolucion (PUT)
					fmt.Println("He fallado un poquito en If 10 - Resolucion (PUT), solucioname!!! ", err)
					logs.Error(r)
					c.Data["system"] = r
					c.Abort("400")
				}
			} else { //If 11 - Resolucion (GET)
				fmt.Println("He fallado un poquito en If 11 - Resolucion (GET), solucioname!!! ", err)
				logs.Error(r)
				c.Data["system"] = r
				c.Abort("400")
			}
		} else { //If 12 - Consecutivo contrato_general
			fmt.Println("He fallado un poquito en If 12 - Consecutivo contrato_general, solucioname!!! ", err)
			logs.Error(cdve)
			c.Data["system"] = cdve
			c.Abort("404")
		}

	} else { //If 13 - Unmarshal
		fmt.Println("He fallado un poquito en If 13 - Unmarshal, solucioname!!! ", err)
		err = amazon.Rollback()
		if err != nil {
			beego.Error(err)
		}
		err = flyway.Rollback()
		if err != nil {
			beego.Error(err)
		}
		return
	}
	err = amazon.Commit()
	if err != nil {
		fmt.Println(err)
	}
	err = flyway.Commit()
	if err != nil {
		fmt.Println(err)
	}
	c.ServeJSON()
}

// Cancelar ...
// @Title Cancelar
// @Description create Cancelar
// @Success 201 {int} models.ExpedicionCancelacion
// @Failure 403 body is empty
// @router /cancelar [post]
func (c *ExpedirResolucionController) Cancelar() {
	amazon := orm.NewOrm()
	flyway := orm.NewOrm()
	err := amazon.Using("amazonAdmin")
	if err != nil {
		beego.Error(err)
	}
	err = flyway.Using("flywayAdmin")
	if err != nil {
		beego.Error(err)
	}
	var m models.ExpedicionCancelacion
	var response interface{}
	//var datosAnular models.DatosAnular
	var contratoCancelado models.ContratoCancelado
	//If 13 - Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		// for vinculaciones
		for _, vinculacion := range *v {
			v := vinculacion.VinculacionDocente
			idVinculacionDocente := strconv.Itoa(v.Id)
			//If vinculacion_docente (get)
			if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idVinculacionDocente, &v); err == nil {
				contratoCancelado.NumeroContrato = v.NumeroContrato.String
				contratoCancelado.Vigencia = int(v.Vigencia.Int64)
				contratoCancelado.FechaCancelacion = vinculacion.ContratoCancelado.FechaCancelacion
				contratoCancelado.MotivoCancelacion = vinculacion.ContratoCancelado.MotivoCancelacion
				contratoCancelado.Usuario = vinculacion.ContratoCancelado.Usuario
				contratoCancelado.FechaRegistro = time.Now()
				contratoCancelado.Estado = vinculacion.ContratoCancelado.Estado
				// if contrato_cancelado (post)
				if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_cancelado", "POST", &response, &contratoCancelado); err == nil {
					var ai []models.ActaInicio
					// if acta_inicio (get)
					if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+contratoCancelado.NumeroContrato+",Vigencia:"+strconv.Itoa(contratoCancelado.Vigencia), &ai); err == nil {
						ai[0].FechaFin = helpers.CalcularFechaFin(ai[0].FechaInicio, v.NumeroSemanasNuevas)
						// if acta_inicio (put)
						if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai[0].Id), "PUT", &response, &ai[0]); err == nil {
							var ce models.ContratoEstado
							var ec models.EstadoContrato
							ce.NumeroContrato = contratoCancelado.NumeroContrato
							ce.Vigencia = contratoCancelado.Vigencia
							ce.FechaRegistro = time.Now()
							ec.Id = 7
							ce.Estado = &ec
							// If contrato_estado (post)
							if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
								var r models.Resolucion
								r.Id = m.IdResolucion
								idResolucionDVE := strconv.Itoa(m.IdResolucion)
								//If  Resolucion (GET)
								if err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE, &r); err == nil {
									r.FechaExpedicion = m.FechaExpedicion
									//If Resolucion (PUT)
									if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &response, &r); err == nil {
										var e models.ResolucionEstado
										var er models.EstadoResolucion
										e.Resolucion = &r
										er.Id = 2
										e.Estado = &er
										e.FechaRegistro = time.Now()
										//If  Resolucion_estado (post)
										if err := helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &response, &e); err == nil {
											fmt.Println("Expedición exitosa, ahora va el commit :D")
											c.Data["json"] = v
										} else { //If  Resolucion_estado (post)
											var response2 interface{}
											fmt.Println("He fallado un poquito en If  Resolucion_estado (post), solucioname!!! ", err)
											helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado/"+strconv.Itoa(e.Id), "DELETE", &response2, nil)
											helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/"+strconv.Itoa(ce.Id), "DELETE", &response2, nil)
											helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_cancelado/"+strconv.Itoa(contratoCancelado.Id), "DELETE", &response2, nil)
											logs.Error(e)
											c.Data["system"] = e
											c.Abort("404")
										}
									} else { //If Resolucion (PUT)
										fmt.Println("He fallado un poquito en If Resolucion (PUT), solucioname!!! ", err)
										logs.Error(r)
										c.Data["system"] = r
										c.Abort("404")
									}
								} else { // If Resolucion (GET)
									fmt.Println("He fallado un poquito en If Resolucion (PUT), solucioname!!! ", err)
									logs.Error(r)
									c.Data["system"] = r
									c.Abort("400")
								}
							} else { // If contrato_estado (post)
								var response2 interface{}
								fmt.Println("He fallado un poquito en If Resolucion (GET), solucioname!!! ", err)
								helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/"+strconv.Itoa(ce.Id), "DELETE", &response2, nil)
								helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_cancelado/"+strconv.Itoa(contratoCancelado.Id), "DELETE", &response2, nil)
								logs.Error(ce)
								c.Data["system"] = ce
								c.Abort("404")
							}
						} else { // If acta_inicio (put)
							fmt.Println("He fallado un poquito en If Acta_Inicio (PUT), solucioname!!! ", err)
							logs.Error(ai[0])
							c.Data["system"] = ai[0]
							c.Abort("404")
						}
					} else { // if acta_inicio (get)
						fmt.Println("He fallado un poquito en if acta_inicio (GET), solucioname!!! ", err)
						logs.Error(ai)
						c.Data["system"] = ai
						c.Abort("404")
					}
				} else { // if contrato_cancelado (post)
					var response2 interface{}
					fmt.Println("He fallado un poquito en if contrato_cancelado (post), solucioname!!! ", err)
					helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_cancelado/"+strconv.Itoa(contratoCancelado.Id), "DELETE", &response2, nil)
					logs.Error(contratoCancelado)
					c.Data["system"] = contratoCancelado
					c.Abort("400")
				}
			} else {
				//If vinculacion_docente (get)
				fmt.Println("He fallado un poquito en If vinculacion_docente (get), solucioname!!! ", err)
				logs.Error(v)
				c.Data["system"] = v
				c.Abort("404")
			}
		} // for vinculaciones

	} else { //If 13 - Unmarshal
		fmt.Println("He fallado un poquito en If 13 - Unmarshal, solucioname!!! ", err)
		err = amazon.Rollback()
		if err != nil {
			beego.Error(err)
		}
		err = flyway.Rollback()
		if err != nil {
			beego.Error(err)
		}
		return
	}
	err = amazon.Commit()
	if err != nil {
		fmt.Println(err)
	}
	err = flyway.Commit()
	if err != nil {
		fmt.Println(err)
	}
	c.ServeJSON()
}

// Calcula el valor del contrato a reversar en dos partes:
// (1) las horas a reducir durante las semanas a reducir
// (2) las horas a originales en las semanas restantes (si quedan después de la reducción)
func CalcularValorContratoReduccion(v [1]models.VinculacionDocente, semanasRestantes int, horasOriginales int, nivelAcademico string) (salarioTotal float64, err error) {
	var d []models.VinculacionDocente
	var salarioSemanasReducidas float64
	var salarioSemanasRestantes float64

	jsonEjemplo, err := json.Marshal(v)
	if err != nil {
		return salarioTotal, err
	}
	err = json.Unmarshal(jsonEjemplo, &d)
	if err != nil {
		return salarioTotal, err
	}

	docentes, err := helpers.CalcularSalarioPrecontratacion(d)
	if err != nil {
		return salarioTotal, err
	}
	salarioSemanasReducidas = docentes[0].ValorContrato
	//Para posgrados no se deben tener en cuenta las semanas restantes
	if semanasRestantes > 0 && nivelAcademico == "PREGRADO" {
		d[0].NumeroSemanas = semanasRestantes
		d[0].NumeroHorasSemanales = horasOriginales
		docentes, err := helpers.CalcularSalarioPrecontratacion(d)
		if err != nil {
			return salarioTotal, err
		}
		salarioSemanasRestantes = docentes[0].ValorContrato
	}
	beego.Info("reducidas ", salarioSemanasReducidas, "restantes ", salarioSemanasRestantes)
	salarioTotal = salarioSemanasReducidas + salarioSemanasRestantes
	return salarioTotal, nil
}
