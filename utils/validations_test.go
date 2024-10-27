package utils

import (
    "testing"
)

// Pruebas para validar correo

func TestValidarCorreo(t *testing.T) {
    tests := []struct {
        correo   string
        isValid bool
    }{
        {"correo@example.com", true},
        {"incorrecto.com", false},
        {"user@domain", false},
        {"valid@domain.com", true},
    }

    for _, tt := range tests {
        err := ValidarCorreo(tt.correo)
        if (err == nil) != tt.isValid {
            t.Errorf("ValidarCorreo(%v) se espera: %v, --> %v", tt.correo, tt.isValid, err)
        }
    }
}

// pruebas para validar telefono
func TestValidarTelefono(t *testing.T) {
    tests := []struct {
        telefono   string
        isValid bool
    }{
        {"1234567890", true},
        {"12345", false},
        {"0987654321", true},
    }

    for _, tt := range tests {
        err := ValidarTelefono(tt.telefono)
        if (err == nil) != tt.isValid {
            t.Errorf("ValidarTelefono(%v) Se espera : %v, --> %v", tt.telefono, tt.isValid, err)
        }
    }
}

// Pruebas para validar Password
func TestValidarPassword(t *testing.T) {
    tests := []struct {
        password string
        isValid  bool
    }{
        {"Abc$1234", true},
        {"password", false},
        {"Aa1@", false},
        {"Abc@1234567890", false},
    }

    for _, tt := range tests {
        err := ValidarPassword(tt.password)
        if (err == nil) != tt.isValid {
            t.Errorf("ValidarPassword(%v) Se espera: %v, --> %v", tt.password, tt.isValid, err)
        }
    }

}
