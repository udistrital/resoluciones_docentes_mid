package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "github.com/udistrital/golog"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func CalculoSalarios(v []models.VinculacionDocente) (total int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CalculoSalarios", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var totalesDisponibilidad int
	if v, err1 := CalcularSalarioPrecontratacion(v); err1 == nil {
		totalesSalario := CalcularTotalSalario(v)
		vigencia := strconv.Itoa(int(v[0].Vigencia.Int64))
		periodo := strconv.Itoa(v[0].Periodo)
		disponibilidad := strconv.Itoa(v[0].Disponibilidad)

		if request2, err2 := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_valores_totales_x_disponibilidad/"+vigencia+"/"+periodo+"/"+disponibilidad+"", &totalesDisponibilidad); err2 == nil && request2 == 200 {
			total = int(totalesSalario) + totalesDisponibilidad
		} else {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/CalculoSalarios2", "err2": err2.Error(), "status": "502"}
			return total, outputError
		}
	} else {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/CalculoSalarios1", "err1": err1, "status": "502"}
		return total, outputError
	}
	return
}

func CalcularSalarioPrecontratacion(docentes_a_vincular []models.VinculacionDocente) (docentes_a_insertar []models.VinculacionDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	nivelAcademico := docentes_a_vincular[0].NivelAcademico
	vigencia := strconv.Itoa(int(docentes_a_vincular[0].Vigencia.Int64))
	var a string
	var categoria string

	salarioMinimo, err1 := CargarSalarioMinimo(vigencia)
	if err1 != nil {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion1", "err1": err1.Error(), "status": "502"}
		return nil, outputError
	}

	for x, docente := range docentes_a_vincular {
		p, err2 := EsDocentePlanta(docente.IdPersona)
		if err1 != nil {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion2", "err2": err2.Error(), "status": "502"}
			return nil, outputError
		}
		if p && strings.ToLower(nivelAcademico) == "posgrado" {
			categoria = strings.TrimSpace(docente.Categoria) + "ud"
		} else {
			categoria = strings.TrimSpace(docente.Categoria)
		}

		var predicados string
		if strings.ToLower(nivelAcademico) == "posgrado" {
			predicados = "valor_salario_minimo(" + strconv.Itoa(salarioMinimo.Valor) + "," + vigencia + ")." + "\n"
			docente.NumeroSemanas = 1
		} else if strings.ToLower(nivelAcademico) == "pregrado" {
			a, err3 := CargarPuntoSalarial()
			if err3 != nil {
				logs.Error(err2)
				outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion3", "err3": err3.Error(), "status": "502"}
				return nil, outputError
			}
			predicados = "valor_punto(" + strconv.Itoa(a.ValorPunto) + ", " + vigencia + ")." + "\n"
		}

		predicados = predicados + "categoria(" + docente.IdPersona + "," + strings.ToLower(categoria) + ", " + vigencia + ")." + "\n"
		predicados = predicados + "vinculacion(" + docente.IdPersona + "," + strings.ToLower(docente.Dedicacion) + ", " + vigencia + ")." + "\n"
		predicados = predicados + "horas(" + docente.IdPersona + "," + strconv.Itoa(docente.NumeroHorasSemanales*docente.NumeroSemanas) + ", " + vigencia + ")." + "\n"
		reglasbase, err4 := CargarReglasBase("CDVE")
		beego.Info("predicados: ", predicados, "a ", a)
		if err4 != nil {
			logs.Error(err4)
			outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion4", "err4": err4.Error(), "status": "502"}
			return nil, outputError
		}
		reglasbase = reglasbase + predicados
		m := NewMachine().Consult(reglasbase)
		beego.Info("m: ", m)
		contratos := m.ProveAll("valor_contrato(" + strings.ToLower(nivelAcademico) + "," + docente.IdPersona + "," + vigencia + ",X).")
		for _, solution := range contratos {
			a = fmt.Sprintf("%s", solution.ByName_("X"))
			beego.Info("a: ", a)
		}
		f, err5 := strconv.ParseFloat(a, 64)
		if err5 != nil {
			logs.Error(err5)
			outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion5", "err5": err5.Error(), "status": "502"}
			return nil, outputError
		}
		salario := f
		beego.Info("f: ", f, "salario: ", salario)
		docentes_a_vincular[x].ValorContrato = salario

	}

	return

}

func CargarSalarioMinimo(vigencia string) (p models.SalarioMinimo, err error) {
	var v []models.SalarioMinimo

	err = GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/salario_minimo/?limit=1&query=Vigencia:"+vigencia, &v)
	if err != nil {
		err = fmt.Errorf("He fallado en salario_minimo (get) función CargarSalarioMinimo, %s", err)
	}

	return v[0], err
}

func EsDocentePlanta(idPersona string) (docentePlanta bool, err error) {
	var temp map[string]interface{}
	var esDePlanta bool

	err = GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAcademica")+"/"+"consultar_datos_docente/"+idPersona, &temp)
	if err != nil {
		esDePlanta = false
		return false, err
	}
	jsonDocentes, err := json.Marshal(temp)
	if err != nil {
		return false, err
	}

	var tempDocentes models.ObjetoDocentePlanta
	err = json.Unmarshal(jsonDocentes, &tempDocentes)
	if err != nil {
		esDePlanta = false
		return false, err
	}

	if tempDocentes.DocenteCollection.Docente[0].Planta == "true" {
		esDePlanta = true
	} else {
		esDePlanta = false
	}

	return esDePlanta, nil
}

func CargarPuntoSalarial() (p models.PuntoSalarial, err error) {
	var v []models.PuntoSalarial

	err = GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/punto_salarial/?sortby=Vigencia&order=desc&limit=1", &v)
	if err != nil {
		err = fmt.Errorf("He fallado en punto_salarial (get) función CargarPuntoSalarial, %s", err)
	}
	return v[0], err
}

func CalcularTotalSalario(v []models.VinculacionDocente) (total float64) {
	var sumatoria float64
	for _, docente := range v {
		sumatoria = sumatoria + docente.ValorContrato
	}

	return sumatoria
}

func Calcular_totales_vinculacion_pdf_nueva(cedula, id_resolucion string, IdDedicacion int) (suma_total_horas int, suma_total_contrato float64, semanasOriginales int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/Calcular_totales_vinculacion_pdf_nueva", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	query := "?limit=-1&query=IdPersona:" + cedula + ",IdResolucion.Id:" + id_resolucion
	var temp []models.VinculacionDocente
	var total_contrato int
	var total_horas int

	// Busca las vinculaciones del docente en la misma resolución (aplica para diferentes proyectos curriculares en vinculaciones y las de modificación)
	if err2 := GetJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &temp); err2 == nil {

		if IdDedicacion != 3 && IdDedicacion != 4 {
			for _, pos := range temp {
				total_horas = total_horas + pos.NumeroHorasSemanales
				total_contrato = total_contrato + int(pos.ValorContrato)
			}
		} else {
			total_horas = temp[0].NumeroHorasSemanales
			total_contrato = int(temp[0].ValorContrato)
		}

	} else {
		total_horas = 0
		total_contrato = 0
		logs.Error(err2)
		outputError = map[string]interface{}{"funcion": "/Calcular_totales_vinculacion_pdf_nueva", "err2": err2.Error(), "status": "502"}
		return total_horas, float64(total_contrato), temp[0].NumeroSemanas, outputError
	}

	return total_horas, float64(total_contrato), temp[0].NumeroSemanas, outputError
}
