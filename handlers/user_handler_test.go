package handlers

import (
	"strings"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRegisterHandler(t *testing.T) {

    payload := map[string]string{
        "usuario":  "newuser",
        "correo":   "newuser@example.com",
        "telefono": "1234567890",
        "password": "Abc@1234",
    }

    body, _ := json.Marshal(payload)
    req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RegisterHandler)

    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusCreated {
        t.Errorf("se espera cofigo 201, --> %v", rr.Code)
    }

    if rr.Body.String() != "Registrado con éxito" {
        t.Errorf("Se espera (Registrado con éxito): %v", rr.Body.String())
    }

}

func TestRegisterHandlerFaltaCampo(t *testing.T) {
    payload := map[string]string{
        "correo":   "newuser@example.com",
        "telefono": "1234567890",
        "password": "Abc@1234",
    }
    body, _ := json.Marshal(payload)
    req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RegisterHandler)
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
        t.Errorf("Se espera codigo 400, --> %v", rr.Code)
    }

    expectedMessage := "Falta el usuario"
    if strings.TrimSpace(rr.Body.String()) != expectedMessage {
        t.Errorf("Se espera: (%v), --> : (%v)", expectedMessage, strings.TrimSpace(rr.Body.String()))
    }
}
