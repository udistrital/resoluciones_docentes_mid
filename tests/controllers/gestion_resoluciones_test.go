package controllers

import (
	"net/http"
	"testing"
)

func TestGetResolucionesAprobadas(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_resoluciones/get_resoluciones_aprobadas?limit=10&offset=0&query="); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestGetResolucionesAprobadas: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetResolucionesAprobadas Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetResolucionesAprobadas:", err.Error())
		t.Fail()
	}

}

func TestGetResolucionesAprobadasError(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_resoluciones/get_resoluciones_aprobadas?limit=0&offset=-1&query="); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestGetResolucionesAprobadasError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetResolucionesAprobadasError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetResolucionesAprobadasError:", err.Error())
		t.Fail()
	}

}

func TestGetResolucionesInscritas(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_resoluciones/get_resoluciones_inscritas?limit=10&offset=0&query="); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestGetResolucionesInscritas: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetResolucionesInscritas Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetResolucionesInscritas:", err.Error())
		t.Fail()
	}
}

func TestGetResolucionesInscritasError(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_resoluciones/get_resoluciones_inscritas?limit=0&offset=-1&query="); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestGetResolucionesInscritasError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetResolucionesInscritasError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetResolucionesInscritasError:", err.Error())
		t.Fail()
	}
}

func TestInsertarResolucionCompleta(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_resoluciones/insertar_resolucion_completa", "", nil); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestInsertarResolucionCompleta: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestInsertarResolucionCompleta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestInsertarResolucionCompleta:", err.Error())
		t.Fail()
	}
}

func TestInsertarResolucionCompletaError(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_resoluciones/insertar_resolucion_completa"); err == nil {
		if response.StatusCode != 404 {
			t.Error("Error TestInsertarResolucionCompletaError: Se esperaba 404 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestInsertarResolucionCompletaError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestInsertarResolucionCompletaError:", err.Error())
		t.Fail()
	}
}
