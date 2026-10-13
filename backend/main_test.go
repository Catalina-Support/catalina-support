package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type pingerDePrueba struct{ err error }

func (p pingerDePrueba) PingContext(context.Context) error { return p.err }

func respuestaDeHealth(t *testing.T, pinger databasePinger, probe schemaProbe) (int, map[string]string) {
	t.Helper()
	response := httptest.NewRecorder()
	health(pinger, probe).ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/api/health", nil))

	body := map[string]string{}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("la respuesta no es JSON válido: %v", err)
	}
	return response.Code, body
}

func TestHealthDistingueEsquemaPendiente(t *testing.T) {
	status, body := respuestaDeHealth(t, pingerDePrueba{}, func(context.Context) (bool, error) {
		return false, nil
	})
	if status != http.StatusServiceUnavailable || body["status"] != "error" || body["errorKey"] != "database.schema_pending" {
		t.Fatalf("respuesta inesperada: status=%d body=%v", status, body)
	}
	if len(body) != 2 {
		t.Fatalf("el error de migración debe conservar la forma exacta acordada: %v", body)
	}
}

func TestHealthConBaseDesconectadaConservaErrorGenerico(t *testing.T) {
	status, body := respuestaDeHealth(t, pingerDePrueba{err: errors.New("sin conexión")}, func(context.Context) (bool, error) {
		t.Fatal("no debe consultar el esquema si PostgreSQL no responde")
		return false, nil
	})
	if status != http.StatusServiceUnavailable || body["status"] != "error" || body["database"] != "error" {
		t.Fatalf("respuesta inesperada: status=%d body=%v", status, body)
	}
}

func TestHealthConEsquemaListo(t *testing.T) {
	status, body := respuestaDeHealth(t, pingerDePrueba{}, func(context.Context) (bool, error) {
		return true, nil
	})
	if status != http.StatusOK || body["status"] != "ok" || body["database"] != "ok" {
		t.Fatalf("respuesta inesperada: status=%d body=%v", status, body)
	}
}

func TestEsquemaPendienteHaceVisibleLaPantallaDeServidorNoDisponible(t *testing.T) {
	handler := pendingSchemaHandler(pingerDePrueba{}, func(context.Context) (bool, error) { return false, nil })
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/api/settings/brand", nil))

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status=%d; la interfaz sólo reconoce el esquema pendiente como indisponibilidad mediante 503", response.Code)
	}
}
