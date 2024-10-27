package services

import (
    "testing"
)

func TestRegistrarUsuario(t *testing.T) {

    // Prueba de registro exitosoo
    err := RegistrarUsuario("testuser", "test@example.com", "4491632192", "Abc@1234")
    if err != nil {
        t.Errorf("RegistrarUsuario error: %v", err)
    }

    // Prueba de registro con correo duplicado
    err = RegistrarUsuario("testuser2", "test@example.com", "4491632193", "Abc@5678")
    if err == nil || err.Error() != "El correo ya esta registrado" {
        t.Errorf("Se espera mensaje de correo registrado, --> %v", err)
    }

    // Prueba de registro con teléfono duplicado
    err = RegistrarUsuario("testuser3", "new@example.com", "4491632192", "Xyz@1234")
    if err == nil || err.Error() != "El teléfono ya esta registrado" {
        t.Errorf("Se espera mensaje de telefono registrado, --> %v", err)
    }
}

func TestLoginUser(t *testing.T) {
    // Registrar usuario para hacer pruebas de login
    RegistrarUsuario("franco", "mizar@example.com", "4491632193", "Abc@1234")

    // Login exitoso
    token, err := LoginUser("franco", "Abc@1234")
    if err != nil {
        t.Errorf("LoginUser error: %v", err)
    }
    if token == "" {
        t.Error("Se espera Token")
    }

    // Login con credenciales incorrectas
    _, err = LoginUser("loginuser", "incorrectpassword")
    if err == nil || err.Error() != "Usuario / contraseña incorrectos" {
        t.Errorf("Se espera mensaje de Usuario / contraseña incorrectos, --> %v", err)
    }
}
