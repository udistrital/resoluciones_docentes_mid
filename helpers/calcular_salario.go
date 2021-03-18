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
		outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion1", "err1": err1, "status": "502"}
		return nil, outputError
	}

	for x, docente := range docentes_a_vincular {
		p, err2 := EsDocentePlanta(docente.IdPersona)
		if err1 != nil {
			logs.Error(err2)
			outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion2", "err2": err2, "status": "502"}
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
				outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion3", "err3": err3, "status": "502"}
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
			outputError = map[string]interface{}{"funcion": "/CalcularSalarioPrecontratacion4", "err4": err4, "status": "502"}
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

func CargarSalarioMinimo(vigencia string) (p models.SalarioMinimo, outputError map[string]interface{}) {
	var v []models.SalarioMinimo

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/salario_minimo/?limit=1&query=Vigencia:"+vigencia, &v); err == nil && response == 200 {
	} else {
		outputError = map[string]interface{}{"funcion": "/CargarSalarioMinimo", "err": err.Error(), "status": "404"}
		return v[0], outputError
	}

	return v[0], nil
}

func EsDocentePlanta(idPersona string) (docentePlanta bool, outputError map[string]interface{}) {
	var temp map[string]interface{}
	var esDePlanta bool

	if response, err := GetJsonWSO2Test("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAcademica")+"/"+"consultar_datos_docente/"+idPersona, &temp); response == 200 && err == nil {
	} else {
		outputError = map[string]interface{}{"funcion": "/EsDocentePlanta1", "err": err.Error(), "status": "404"}
		return false, outputError
	}
	jsonDocentes, err1 := json.Marshal(temp)
	if err1 != nil {
		outputError = map[string]interface{}{"funcion": "/EsDocentePlanta2", "err": "Error en codificación de datos", "status": "404"}
		return false, outputError
	}

	var tempDocentes models.ObjetoDocentePlanta
	err2 := json.Unmarshal(jsonDocentes, &tempDocentes)
	if err2 != nil {
		outputError = map[string]interface{}{"funcion": "/EsDocentePlanta3", "err": "Error en decodificación de datos", "status": "404"}
		return false, outputError
	}

	if tempDocentes.DocenteCollection.Docente[0].Planta == "true" {
		esDePlanta = true
	} else {
		esDePlanta = false
	}

	return esDePlanta, nil
}

func CargarPuntoSalarial() (p models.PuntoSalarial, outputError map[string]interface{}) {
	var v []models.PuntoSalarial

	if response, err := GetJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/punto_salarial/?sortby=Vigencia&order=desc&limit=1", &v); response == 200 && err == nil {
	} else {
		outputError = map[string]interface{}{"funcion": "/CargarPuntoSalarial", "err": err.Error(), "status": "404"}
		return v[0], outputError
	}
	return v[0], nil
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

// Calcula el valor del contrato a reversar en dos partes:
// (1) las horas a reducir durante las semanas a reducir
// (2) las horas a originales en las semanas restantes (si quedan después de la reducción)
func CalcularValorContratoReduccion(v [1]models.VinculacionDocente, semanasRestantes int, horasOriginales int, nivelAcademico string) (salarioTotal float64, outputError map[string]interface{}) {
	var d []models.VinculacionDocente
	var salarioSemanasReducidas float64
	var salarioSemanasRestantes float64

	jsonEjemplo, err1 := json.Marshal(v)
	if err1 != nil {
		outputError = map[string]interface{}{"funcion": "/CalcularValorContratoReduccion1", "err": err1.Error(), "status": "404"}
		return salarioTotal, outputError
	}
	err2 := json.Unmarshal(jsonEjemplo, &d)
	if err2 != nil {
		outputError = map[string]interface{}{"funcion": "/CalcularValorContratoReduccion2", "err": err2.Error(), "status": "404"}
		return salarioTotal, outputError
	}

	docentes, err := CalcularSalarioPrecontratacion(d)
	if err != nil {
		return salarioTotal, err
	}
	salarioSemanasReducidas = docentes[0].ValorContrato
	//Para posgrados no se deben tener en cuenta las semanas restantes
	if semanasRestantes > 0 && nivelAcademico == "PREGRADO" {
		d[0].NumeroSemanas = semanasRestantes
		d[0].NumeroHorasSemanales = horasOriginales
		docentes, err := CalcularSalarioPrecontratacion(d)
		if err != nil {
			return salarioTotal, err
		}
		salarioSemanasRestantes = docentes[0].ValorContrato
	}
	beego.Info("reducidas ", salarioSemanasReducidas, "restantes ", salarioSemanasRestantes)
	salarioTotal = salarioSemanasReducidas + salarioSemanasRestantes
	return salarioTotal, nil
}
