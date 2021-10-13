package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
)

var parameters struct {
	ProtocolAdmin       string
	UrlcrudResoluciones string
	UrlcrudCore         string
	UrlcrudAgora        string
	UrlcrudKronos       string
	Urlruler            string
	UrlcrudOikos        string
	UrlcrudWSO2         string
}

func TestMain(m *testing.M) {
	parameters.ProtocolAdmin = os.Getenv("RESOLUCIONES_MID_PROTOCOL_ADMIN")
	parameters.UrlcrudResoluciones = os.Getenv("RESOLUCIONES_MID_RESOLUCIONES_CRUD_URL")
	parameters.UrlcrudCore = os.Getenv("RESOLUCIONES_MID_CORE_URL")
	parameters.UrlcrudAgora = os.Getenv("RESOLUCIONES_MID_AGORA_URL")
	parameters.UrlcrudKronos = os.Getenv("RESOLUCIONES_MID_KRONOS_URL")
	parameters.Urlruler = os.Getenv("RESOLUCIONES_MID_RULER_URL")
	parameters.UrlcrudOikos = os.Getenv("RESOLUCIONES_MID_OIKOS_URL")
	parameters.UrlcrudWSO2 = os.Getenv("RESOLUCIONES_MID_WSO2_URL")
	flag.Parse()
	os.Exit(m.Run())
}

func TestEndPointResolucionesCrud(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.UrlcrudResoluciones
	BaseTestEndpoint(t, endpoint)
}
func TestEndPointCore(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.UrlcrudCore
	BaseTestEndpoint(t, endpoint)
}
func TestEndPointAgora(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.UrlcrudAgora
	BaseTestEndpoint(t, endpoint)
}
func TestEndPointKronos(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.UrlcrudKronos
	BaseTestEndpoint(t, endpoint)
}
func TestEndPointRuler(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.Urlruler
	BaseTestEndpoint(t, endpoint)
}
func TestEndPointOikos(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.UrlcrudOikos
	BaseTestEndpoint(t, endpoint)
}
func TestEndPointWSO2(t *testing.T) {
	endpoint := parameters.ProtocolAdmin + "://" + parameters.UrlcrudWSO2
	BaseTestEndpoint(t, endpoint)
}

func BaseTestEndpoint(t *testing.T, endpoint string) {
	t.Log(endpoint)
	if response, err := http.Get(endpoint); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndpoint:", endpoint, "Estado: ", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint", endpoint, "Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}
}
