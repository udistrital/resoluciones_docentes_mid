package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/udistrital/resoluciones_docentes_mid/helpers"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

// PreliquidacionController operations for Preliquidacion
type GestionPrevinculacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionPrevinculacionesController) URLMapping() {
	c.Mapping("InsertarPrevinculaciones", c.InsertarPrevinculaciones)
	c.Mapping("CalcularTotalDeSalarios", c.CalcularTotalSalarios)
	c.Mapping("ListarDocentesCargaHoraria", c.ListarDocentesCargaHoraria)
	c.Mapping("GetCdpRpDocente", c.GetCdpRpDocente)
}

// Calcular_total_de_salarios_seleccionados ...
// @Title Calcular_total_de_salarios_seleccionados
// @Description createCalcular_total_de_salarios_seleccionados
// @Success 201 {int} int
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos_seleccionados [post]
func (c *GestionPrevinculacionesController) Calcular_total_de_salarios_seleccionados() {

	var v []models.VinculacionDocente
	var total int
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	v, err = helpers.CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	total = int(helpers.CalcularTotalSalario(v))
	c.Data["json"] = total

	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos [post]
func (c *GestionPrevinculacionesController) CalcularTotalSalarios() {

	var v []models.VinculacionDocente
	var totalesDisponibilidad int
	var total int
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	v, err = helpers.CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	totalesSalario := helpers.CalcularTotalSalario(v)
	vigencia := strconv.Itoa(int(v[0].Vigencia.Int64))
	periodo := strconv.Itoa(v[0].Periodo)
	disponibilidad := strconv.Itoa(v[0].Disponibilidad)

	err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_valores_totales_x_disponibilidad/"+vigencia+"/"+periodo+"/"+disponibilidad+"", &totalesDisponibilidad)
	if err != nil {
		beego.Error("ERROR al calcular total de contratos", err)
		c.Abort("403")
	}
	total = int(totalesSalario) + totalesDisponibilidad
	c.Data["json"] = total

	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/insertar_previnculaciones [post]
func (c *GestionPrevinculacionesController) InsertarPrevinculaciones() {

	var v []models.VinculacionDocente
	var idRespuesta int

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error("Error al hacer unmarshal", err)
		c.Abort("403")
	}
	v, err = helpers.CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}

	err = helpers.SendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &idRespuesta, &v)
	c.Data["json"] = idRespuesta
	if err != nil {
		beego.Error("Error al insertar docentes", err)
		c.Abort("403")
	}
	c.ServeJSON()
}

// GestionPrevinculacionesController ...
// @Title ListarDocentesCargaHoraria
// @Description create ListarDocentesCargaHoraria
// @Param vigencia query string false "año a consultar"
// @Param periodo query string false "periodo a listar"
// @Param tipo_vinculacion query string false "vinculacion del docente"
// @Param facultad query string false "facultad"
// @Param nivel_academico query string false "nivel_academico"
// @Success 201 {object} models.Docentes_x_Carga
// @Failure 403 body is empty
// @router /Precontratacion/docentes_x_carga_horaria [get]
func (c *GestionPrevinculacionesController) ListarDocentesCargaHoraria() {

	vigencia := c.GetString("vigencia")
	periodo := c.GetString("periodo")
	tipoVinculacion := c.GetString("tipo_vinculacion")
	facultad := c.GetString("facultad")
	nivelAcademico := c.GetString("nivel_academico")

	docentesXcargaHoraria, err := helpers.ListarDocentesHorasLectivas(vigencia, periodo, tipoVinculacion, facultad, nivelAcademico)
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}
	newDocentesXcargaHoraria := models.ObjetoCargaLectiva{}

	//BUSCAR CATEGORÍA DE CADA DOCENTE
	for _, pos := range docentesXcargaHoraria.CargasLectivas.CargaLectiva {
		catDocente := models.ObjetoCategoriaDocente{}
		emptyCatDocente := models.ObjetoCategoriaDocente{}
		//TODO: quitar el hardconding para WSO2 cuando ya soporte https:
		q := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudUrano") + "/categoria_docente/" + vigencia + "/" + periodo + "/" + pos.DocDocente
		err = helpers.GetXml(q, &catDocente.CategoriaDocente)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}

		pos.CategoriaNombre, pos.IDCategoria, err = helpers.Buscar_Categoria_Docente(vigencia, periodo, pos.DocDocente)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		if catDocente.CategoriaDocente != emptyCatDocente.CategoriaDocente {
			newDocentesXcargaHoraria.CargasLectivas.CargaLectiva = append(newDocentesXcargaHoraria.CargasLectivas.CargaLectiva, pos)
		}
	}

	//RETORNAR CON ID DE TIPO DE VINCULACION DE NUEVO MODELO
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.IDTipoVinculacion, pos.NombreTipoVinculacion = helpers.HomologarDedicacion_ID("old", pos.IDTipoVinculacion)
		if pos.IDTipoVinculacion == "3" {
			pos.HorasLectivas = "20"
			pos.NombreTipoVinculacion = "MTO"
		}
		if pos.IDTipoVinculacion == "4" {
			pos.HorasLectivas = "40"
			pos.NombreTipoVinculacion = "TCO"
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}

	//RETORNAR FACULTTADES CON ID DE OIKOS, HOMOLOGACION
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.IDFacultad, err = helpers.HomologarFacultad("old", pos.IDFacultad)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}
	//RETORNAR PROYECTOS CURRICUALRES HOMOLOGADOS!!
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.DependenciaAcademica, err = strconv.Atoi(pos.IDProyecto)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		pos.IDProyecto, err = helpers.HomologarProyectoCurricular(pos.IDProyecto)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos

	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = newDocentesXcargaHoraria.CargasLectivas.CargaLectiva
	c.ServeJSON()

}

const (
	tipoVinculacion = iota + 1
	tipoCancelacion
	tipoAdicion
	tipoReduccion
)

// GestionPrevinculacionesController ...
// @Title ListarDocentesPrevinculadosAll
// @Description create ListarDocentesPrevinculadosAll
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados_all [get]
func (c *GestionPrevinculacionesController) ListarDocentesPrevinculadosAll() {
	idResolucion := c.GetString("id_resolucion")
	var v = []models.VinculacionDocente{}
	var res models.Resolucion
	var resvinc models.ResolucionVinculacionDocente
	var modres []models.ModificacionResolucion
	var vinc []models.VinculacionDocente
	var modvin []models.ModificacionVinculacion
	var ValorModificacionContrato float64

	//Devuelve el nivel académico, la dedicación y la facultad de la resolución
	err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente/"+idResolucion, &resvinc)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	//Devuelve la información básica de la resolución que se está consultando
	err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucion, &res)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	if res.IdTipoResolucion.Id != tipoVinculacion {
		//Busca el id de la modificación donde se relacionan la resolución original y la de modificación asociada
		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
	}

	//Devuelve las vinculaciones presentes en la resolución consultada, agrupadas o no, según el nivel académico
	if resvinc.NivelAcademico == "POSGRADO" {
		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente?limit=-1&query=IdResolucion.Id:"+idResolucion, &vinc)
	}
	if resvinc.NivelAcademico == "PREGRADO" {
		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_vinculaciones_agrupadas/"+idResolucion, &vinc)
	}
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	var llenarVinculacion = func(v *models.VinculacionDocente) {
		documentoIdentidad, _ := strconv.Atoi(v.IdPersona)
		v.NombreCompleto = helpers.BuscarNombreProveedor(documentoIdentidad)
		v.Dedicacion = helpers.BuscarNombreDedicacion(v.IdDedicacion.Id)
		v.LugarExpedicionCedula = helpers.BuscarLugarExpedicion(v.IdPersona)
		v.TipoDocumento = helpers.BuscarTipoDocumento(v.IdPersona)
		v.NumeroDisponibilidad = helpers.BuscarNumeroDisponibilidad(v.Disponibilidad)
	}

	switch res.IdTipoResolucion.Id {
	case tipoVinculacion:
		for x, pos := range vinc {
			v = append(v, pos)
			llenarVinculacion(&pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(pos.IdResolucion.Id), pos.IdDedicacion.Id)
			}
			pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
			pos.ValorContratoFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
			v[x] = pos
		}
		break
	case tipoCancelacion:
		for x, pos := range vinc {
			v = append(v, pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				//Agrupa las modificaciones en la resolución por persona
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				pos.ValorModificacionFormato = helpers.FormatMoney(int(ValorModificacionContrato), 2)

				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
				pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanas-pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
				pos.ValorContratoFormato = helpers.FormatMoney(int(pos.ValorContrato-ValorModificacionContrato), 2)
			}
			if resvinc.NivelAcademico == "POSGRADO" {
				pos.NumeroHorasModificacion = pos.NumeroHorasSemanales
				pos.ValorModificacionFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
				ValorModificacionContrato = pos.ValorContrato

				//Busca la vinculación original a la que está asociada la modificación
				err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err != nil {
					beego.Error(err)
					c.Abort("400")
				}
				var vincOriginal = modvin[0].VinculacionDocenteCancelada
				pos.NumeroHorasSemanales = vincOriginal.NumeroHorasSemanales
				pos.NumeroMeses = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas-pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = helpers.FormatMoney(int(vincOriginal.ValorContrato), 2)
				pos.ValorContratoFormato = helpers.FormatMoney(int(vincOriginal.ValorContrato-ValorModificacionContrato), 2)
			}
			llenarVinculacion(&pos)
			v[x] = pos
		}
		break
	case tipoAdicion:
		for x, pos := range vinc {
			v = append(v, pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				//Agrupa las modificaciones en la resolución por persona
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				pos.ValorModificacionFormato = helpers.FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
				pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroHorasNuevas = pos.NumeroHorasSemanales + pos.NumeroHorasModificacion
				pos.ValorContratoFormato = helpers.FormatMoney(int(pos.ValorContrato+ValorModificacionContrato), 2)
			}
			if resvinc.NivelAcademico == "POSGRADO" {
				pos.NumeroHorasModificacion = pos.NumeroHorasSemanales
				pos.ValorModificacionFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				ValorModificacionContrato = pos.ValorContrato
				//Busca la vinculación original a la que está asociada la modificación
				err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id)+",VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err != nil {
					beego.Error(err)
					c.Abort("400")
				}
				var vincOriginal = modvin[0].VinculacionDocenteCancelada
				pos.NumeroHorasSemanales = vincOriginal.NumeroHorasSemanales
				pos.NumeroMeses = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = helpers.FormatMoney(int(vincOriginal.ValorContrato), 2)
				pos.NumeroHorasNuevas = vincOriginal.NumeroHorasSemanales + pos.NumeroHorasModificacion
				pos.ValorContratoFormato = helpers.FormatMoney(int(vincOriginal.ValorContrato+ValorModificacionContrato), 2)
			}
			llenarVinculacion(&pos)
			v[x] = pos
		}
		break
	case tipoReduccion:
		for x, pos := range vinc {
			v = append(v, pos)
			if resvinc.NivelAcademico == "PREGRADO" {
				//Agrupa las modificaciones en la resolución por persona
				pos.NumeroHorasModificacion, ValorModificacionContrato, pos.NumeroSemanasNuevas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, idResolucion, pos.IdDedicacion.Id)
				pos.ValorModificacionFormato = helpers.FormatMoney(int(ValorModificacionContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanasNuevas)/4, 'f', 2, 64) + " meses"
				//Agrupa las vinculaciones originales por persona
				pos.NumeroHorasSemanales, pos.ValorContrato, pos.NumeroSemanas = helpers.Calcular_totales_vinculacion_pdf_nueva(pos.IdPersona, strconv.Itoa(modres[0].ResolucionAnterior), pos.IdDedicacion.Id)
				pos.NumeroMeses = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroHorasNuevas = pos.NumeroHorasSemanales - pos.NumeroHorasModificacion
				pos.ValorContratoFormato = helpers.FormatMoney(int(pos.ValorContrato-ValorModificacionContrato), 2)

			}
			if resvinc.NivelAcademico == "POSGRADO" {
				pos.NumeroHorasModificacion = pos.NumeroHorasSemanales
				pos.ValorModificacionFormato = helpers.FormatMoney(int(pos.ValorContrato), 2)
				pos.NumeroMesesNuevos = strconv.FormatFloat(float64(pos.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				ValorModificacionContrato = pos.ValorContrato
				//Busca la vinculación original a la que está asociada la modificación
				err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+
					"/modificacion_vinculacion/?query=VinculacionDocenteRegistrada:"+strconv.Itoa(pos.Id), &modvin)
				if err != nil {
					beego.Error(err)
					c.Abort("400")
				}
				var vincOriginal = modvin[0].VinculacionDocenteCancelada
				pos.NumeroHorasSemanales = vincOriginal.NumeroHorasSemanales
				pos.NumeroMeses = strconv.FormatFloat(float64(vincOriginal.NumeroSemanas)/4, 'f', 2, 64) + " meses"
				pos.ValorContratoInicialFormato = helpers.FormatMoney(int(vincOriginal.ValorContrato), 2)
				pos.NumeroHorasNuevas = vincOriginal.NumeroHorasSemanales - pos.NumeroHorasModificacion
				pos.ValorContratoFormato = helpers.FormatMoney(int(vincOriginal.ValorContrato-ValorModificacionContrato), 2)
			}

			llenarVinculacion(&pos)
			v[x] = pos
		}
		break
	default:
		break
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()

}

//ESTA FUNCIÓN LISTA LOS DOCENTES PREVINCULADOS EN TRUE

// GestionPrevinculacionesController ...
// @Title ListarDocentesPrevinculados
// @Description create ListarDocentesPrevinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados [get]
func (c *GestionPrevinculacionesController) ListarDocentesPrevinculados() {
	idResolucion, err := c.GetInt("id_resolucion")
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}

	query := "?limit=-1&query=IdResolucion.Id:" + strconv.Itoa(idResolucion) + ",Estado:true"
	var v = []models.VinculacionDocente{}
	var res models.Resolucion
	var modres []models.ModificacionResolucion
	var modvin []models.ModificacionVinculacion

	err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(idResolucion), &res)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	if res.IdTipoResolucion.Id == tipoVinculacion {
		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
		err := helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+strconv.Itoa(idResolucion), &modres)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		if len(modvin) != 0 {
			arreglo := make([]string, len(modvin))
			for x, pos := range modvin {
				arreglo[x] = strconv.Itoa(pos.VinculacionDocenteRegistrada.Id)
			}
			identificadoresvinc := strings.Join(arreglo, "|")
			err = helpers.GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=Estado:True,Id__in:"+identificadoresvinc+"&limit=-1", &v)
			if err != nil {
				beego.Error(err)
				c.Abort("400")
			}
		} else {
			v = nil
		}
	}
	for x, pos := range v {
		documentoIdentidad, _ := strconv.Atoi(pos.IdPersona)

		pos.NombreCompleto = helpers.BuscarNombreProveedor(documentoIdentidad)
		pos.NumeroDisponibilidad = helpers.BuscarNumeroDisponibilidad(pos.Disponibilidad)
		pos.Dedicacion = helpers.BuscarNombreDedicacion(pos.IdDedicacion.Id)
		pos.LugarExpedicionCedula = helpers.BuscarLugarExpedicion(pos.IdPersona)
		pos.TipoDocumento = helpers.BuscarTipoDocumento(pos.IdPersona)
		pos.ValorContratoFormato = helpers.FormatMoney(int(v[x].ValorContrato), 2)
		pos.ProyectoNombre = helpers.BuscarNombreFacultad(int(v[x].IdProyectoCurricular))
		pos.Periodo = res.Periodo
		pos.VigenciaCarga = res.VigenciaCarga
		pos.PeriodoCarga = res.PeriodoCarga

		v[x] = pos
	}
	if v == nil {
		v = []models.VinculacionDocente{}
	}
	c.Ctx.Output.SetStatus(201)

	c.Data["json"] = v
	c.ServeJSON()

}

// GestionPrevinculacionesController ...
// @Title GetCdpRpDocente
// @Description Get RPs de la vinculacion docente
// @Param num_vinculacion query string true "Número de la vinculación del docente"
// @Param vigencia query string true "Vigencia de la vinculación del docente"
// @Success 201 {object}  models.RpDocente
// @Failure 403 :num_vinculacion is empty
// @Failure 403 :vigencia is empty
// @router /rp_docente/:num_vinculacion/:vigencia/:identificacion [get]
func (c *GestionPrevinculacionesController) GetCdpRpDocente() {
	num_vinculacion := c.Ctx.Input.Param(":num_vinculacion")
	vigencia := c.Ctx.Input.Param(":vigencia")
	identificacion := c.Ctx.Input.Param(":identificacion")

	var contratoDisponibilidad []models.ContratoDisponibilidad
	var rpdocente models.RpDocente

	//If 1 contrato_disponibilidad (get)
	if err := helpers.GetJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+num_vinculacion+",Vigencia:"+vigencia, &contratoDisponibilidad); err == nil { //If 2  (get)
		//for contrato_disponibilidad
		for _, pos := range contratoDisponibilidad {
			rpdocente = helpers.GetInformacionRpDocente(strconv.Itoa(pos.NumeroCdp), strconv.Itoa(pos.VigenciaCdp), identificacion)
			c.Data["json"] = rpdocente
		}

	} else { //If 1 contrato_disponibilidad (get)
		fmt.Println("He fallado en If 1 contrato_disponibilidad (get), solucioname!!!", err)
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = rpdocente
	c.ServeJSON()

}
