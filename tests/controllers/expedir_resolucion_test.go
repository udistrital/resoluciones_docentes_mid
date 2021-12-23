package controllers

import (
	"net/http"
	"testing"
)

func TestExpedirError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/expedir_resolucion/expedir", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestExpedirError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestExpedirError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestExpedirError:", err.Error())
		t.Fail()
	}
}

func TestValidarDatosExpedicionError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/expedir_resolucion/validar_datos_expedicion", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestValidarDatosExpedicionError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestValidarDatosExpedicionError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestValidarDatosExpedicionError:", err.Error())
		t.Fail()
	}
}

func TestExpedirModificacionError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/expedir_resolucion/expedirModificacion", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestExpedirModificacionError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestExpedirModificacionError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestExpedirModificacionError:", err.Error())
		t.Fail()
	}
}

func TestCancelarError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/expedir_resolucion/cancelar", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestCancelarError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestCancelarError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestCancelarError:", err.Error())
		t.Fail()
	}
}
