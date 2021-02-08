package helpers

import (
	"encoding/json"
	"strconv"
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
	. "github.com/udistrital/golog"
)

func CalcularSalarioPrecontratacion(docentes_a_vincular []models.VinculacionDocente) (docentes_a_insertar []models.VinculacionDocente, err error) {
	nivelAcademico := docentes_a_vincular[0].NivelAcademico
	vigencia := strconv.Itoa(int(docentes_a_vincular[0].Vigencia.Int64))
	var a string
	var categoria string

	salarioMinimo, err := CargarSalarioMinimo(vigencia)
	if err != nil {
		return docentes_a_insertar, err
	}

	for x, docente := range docentes_a_vincular {
		p, err := EsDocentePlanta(docente.IdPersona)
		if err != nil {
			return docentes_a_insertar, err
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
			a, err := CargarPuntoSalarial()
			if err != nil {
				return docentes_a_insertar, err
			}
			predicados = "valor_punto(" + strconv.Itoa(a.ValorPunto) + ", " + vigencia + ")." + "\n"
		}

		predicados = predicados + "categoria(" + docente.IdPersona + "," + strings.ToLower(categoria) + ", " + vigencia + ")." + "\n"
		predicados = predicados + "vinculacion(" + docente.IdPersona + "," + strings.ToLower(docente.Dedicacion) + ", " + vigencia + ")." + "\n"
		predicados = predicados + "horas(" + docente.IdPersona + "," + strconv.Itoa(docente.NumeroHorasSemanales*docente.NumeroSemanas) + ", " + vigencia + ")." + "\n"
		reglasbase, err := CargarReglasBase("CDVE")
		beego.Info("predicados: ", predicados, "a ", a)


		if err != nil {
			return docentes_a_insertar, err
		}
		reglasbase = reglasbase + predicados
		m := NewMachine().Consult(reglasbase)
   	beego.Info("m: ", m)
		contratos := m.ProveAll("valor_contrato(" + strings.ToLower(nivelAcademico) + "," + docente.IdPersona + "," + vigencia + ",X).")
		for _, solution := range contratos {
			a = fmt.Sprintf("%s", solution.ByName_("X"))
				beego.Info("a: ", a)
		}
		f, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return docentes_a_vincular, err
		}
		salario := f
		beego.Info("f: ", f, "salario: ", salario)
		docentes_a_vincular[x].ValorContrato = salario

	}

	return docentes_a_vincular, nil

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