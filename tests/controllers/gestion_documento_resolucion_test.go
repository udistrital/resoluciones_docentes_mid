package controllers

import (
	"net/http"
	"testing"
)

func TestGetContenidoResolucion(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_documento_resolucion/get_contenido_resolucion?id_resolucion=333&id_facultad=66"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestGetContenidoResolucion: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetContenidoResolucion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetContenidoResolucion:", err.Error())
		t.Fail()
	}

}

func TestGetContenidoResolucionError(t *testing.T) {

	if response, err := http.Get("http://localhost:8521/v1/gestion_documento_resolucion/get_contenido_resolucion?id_resolucion=36&id_faculdat=14"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestGetContenidoResolucionError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetContenidoResolucionError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestGetContenidoResolucionError:", err.Error())
		t.Fail()
	}

}
