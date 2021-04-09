package documento

import (
	"net/http"
	"testing"
)

func TestCalcular_total_de_salarios_seleccionados(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/calcular_valor_contratos_seleccionados"); err == nil {
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

func TestCalcular_total_de_salarios_seleccionadosError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/calcular_valor_contratos_seleccionados"); err == nil {
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

func TestCalcularTotalSalarios(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/calcular_valor_contratos"); err == nil {
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

func TestCalcularTotalSalariosError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/calcular_valor_contratos"); err == nil {
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

func TestInsertarPrevinculaciones(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/insertar_previnculaciones"); err == nil {
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

func TestInsertarPrevinculacionesError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/insertar_previnculaciones"); err == nil {
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

func TestListarDocentesCargaHoraria(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/docentes_x_carga_horaria/2020/3/HCH/14/PREGRADO"); err == nil {
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

func TestListarDocentesCargaHorariaError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/Precontratacion/docentes_x_carga_horaria/2020a/3e/HCH/14/PREGRADO"); err == nil {
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

func TestListarDocentesPrevinculadosAll(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/docentes_previnculados_all/219"); err == nil {
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

func TestListarDocentesPrevinculadosAllError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/docentes_previnculados_all/219a"); err == nil {
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

func TestListarDocentesPrevinculados(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/docentes_previnculados/219"); err == nil {
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

func TestListarDocentesPrevinculadosError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/docentes_previnculados/219"); err == nil {
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

func TestGetCdpRpDocente(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1//rp_docente/:num_vinculacion/:vigencia/:identificacion"); err == nil {
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

func TestGetCdpRpDocenteError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1//rp_docente/:num_vinculacion/:vigencia/:identificacion"); err == nil {
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
