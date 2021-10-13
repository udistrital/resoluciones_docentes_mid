package controllers

import (
	"net/http"
	"testing"
)

func TestCalcular_total_de_salarios_seleccionadosError(t *testing.T) {

	if response, err := http.Post("http://localhost:8521/v1/gestion_previnculacion/Precontratacion/calcular_valor_contratos_seleccionados", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestCalcular_total_de_salarios_seleccionadosError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestCalcular_total_de_salarios_seleccionadosError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestCalcular_total_de_salarios_seleccionadosError:", err.Error())
		t.Fail()
	}

}

func TestCalcularTotalSalariosError(t *testing.T) {

	if response, err := http.Post("http://localhost:8521/v1/gestion_previnculacion/Precontratacion/calcular_valor_contratos", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestCalcularTotalSalariosError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestCalcularTotalSalariosError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestCalcularTotalSalariosError:", err.Error())
		t.Fail()
	}

}

func TestInsertarPrevinculacionesError(t *testing.T) {

	if response, err := http.Post("http://localhost:8521/v1/gestion_previnculacion/Precontratacion/insertar_previnculaciones", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestInsertarPrevinculacionesError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestInsertarPrevinculacionesError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestInsertarPrevinculacionesError:", err.Error())
		t.Fail()
	}

}

func TestListarDocentesCargaHoraria(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/Precontratacion/docentes_x_carga_horaria?vigencia=2020&periodo=1&tipo_vinculacion=HCH&facultad=14&nivel_academico=PREGRADO"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestListarDocentesCargaHoraria: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesCargaHoraria Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesCargaHoraria:", err.Error())
		t.Fail()
	}

}

func TestListarDocentesCargaHorariaError(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/Precontratacion/docentes_x_carga_horaria?vigencia=0&periodo=0&tipo_vinculacion=HCH&facultad=0&nivel_academico=PREGRADO"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestListarDocentesCargaHorariaError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesCargaHorariaError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesCargaHorariaError:", err.Error())
		t.Fail()
	}

}

func TestListarDocentesPrevinculadosAll(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/docentes_previnculados_all?id_resolucion=219"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestListarDocentesPrevinculadosAll: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesPrevinculadosAll Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesPrevinculadosAll:", err.Error())
		t.Fail()
	}

}

func TestListarDocentesPrevinculadosAllError(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/docentes_previnculados_all"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestListarDocentesPrevinculadosAllError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesPrevinculadosAllError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesPrevinculadosAllError:", err.Error())
		t.Fail()
	}

}

func TestListarDocentesPrevinculados(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/docentes_previnculados?id_resolucion=219"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestListarDocentesPrevinculados: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesPrevinculados Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesPrevinculados:", err.Error())
		t.Fail()
	}

}

func TestListarDocentesPrevinculadosError(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/docentes_previnculados"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestListarDocentesPrevinculadosError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesPrevinculadosError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesPrevinculadosError:", err.Error())
		t.Fail()
	}

}

func TestGetCdpRpDocente(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/rp_docente/1792/2021/79974416"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestGetCdpRpDocente: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetCdpRpDocente Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetCdpRpDocente:", err.Error())
		t.Fail()
	}

}

func TestGetCdpRpDocenteError(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_previnculacion/rp_docente/1792/2021/0"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestGetCdpRpDocenteError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetCdpRpDocenteError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetCdpRpDocenteError:", err.Error())
		t.Fail()
	}

}
