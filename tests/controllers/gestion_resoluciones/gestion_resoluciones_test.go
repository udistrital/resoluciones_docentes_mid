package documento

import (
	"net/http"
	"testing"
)

func TestGetResolucionesAprobadas(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/get_resoluciones_aprobadas/?limit=0&offset=0&query="); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}

}

func TestGetResolucionesAprobadasError(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/get_resoluciones_aprobadas/?limit=0&offset=0&query="); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}

}

func TestGetResolucionesInscritas(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/get_resoluciones_inscritas/?limit=0&offset=0&query="); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}
}

func TestGetResolucionesInscritasError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/get_resoluciones_inscritas/?limit=0&offset=0&query="); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}
}

func TestInsertarResolucionCompleta(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/insertar_resolucion_completa"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}
}

func TestInsertarResolucionCompletaError(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/insertar_resolucion_completa"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}
}
