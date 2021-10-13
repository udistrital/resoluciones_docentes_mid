package controllers

import (
	"net/http"
	"testing"
)

func TestListarDocentesDesvinculados(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_desvinculaciones/docentes_desvinculados?id_resolucion=2511"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestListarDocentesDesvinculados: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesDesvinculados Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesDesvinculados:", err.Error())
		t.Fail()
	}
}

func TestListarDocentesDesvinculadosError(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_desvinculaciones/docentes_desvinculados?id_resolucion=fff"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestListarDocentesDesvinculadosError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesDesvinculadosError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesDesvinculadosError:", err.Error())
		t.Fail()
	}
}

func TestListarDocentesCancelados(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_desvinculaciones/docentes_cancelados?id_resolucion=2511"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestListarDocentesCancelados: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesCancelados Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesCancelados:", err.Error())
		t.Fail()
	}
}

func TestListarDocentesCanceladosError(t *testing.T) {
	if response, err := http.Get("http://localhost:8521/v1/gestion_desvinculaciones/docentes_cancelados?id_resolucion=fff"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestListarDocentesCanceladosError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestListarDocentesCanceladosError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestListarDocentesCanceladosError:", err.Error())
		t.Fail()
	}
}

func TestAnularModificacionesError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_desvinculaciones/anular_modificaciones", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestAnularModificacionesError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestAnularModificacionesError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestAnularModificacionesError:", err.Error())
		t.Fail()
	}
}

func TestAnularAdicionDocenteError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_desvinculaciones/anular_adicion", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestAnularAdicionDocenteError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestAnularAdicionDocenteError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestAnularAdicionDocenteError:", err.Error())
		t.Fail()
	}
}

func TestConsultarCategoriaError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_desvinculaciones/consultar_categoria", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestConsultarCategoriaError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarCategoriaError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarCategoriaError:", err.Error())
		t.Fail()
	}
}

func TestValidarSaldoCDPError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_desvinculaciones/validar_saldo_cdp", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestValidarSaldoCDPError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestValidarSaldoCDPError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestValidarSaldoCDPError:", err.Error())
		t.Fail()
	}
}

func TestAdicionarHorasError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_desvinculaciones/adicionar_horas", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestAdicionarHorasError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestAdicionarHorasError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestAdicionarHorasError:", err.Error())
		t.Fail()
	}
}

func TestActualizarVinculacionesCancelacionError(t *testing.T) {
	if response, err := http.Post("http://localhost:8521/v1/gestion_desvinculaciones/actualizar_vinculaciones_cancelacion", "", nil); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestActualizarVinculacionesCancelacionError: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestActualizarVinculacionesCancelacionError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestActualizarVinculacionesCancelacionError:", err.Error())
		t.Fail()
	}
}
