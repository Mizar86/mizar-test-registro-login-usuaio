package utils

import (
    "errors"
    "regexp"
)

// Valida correo
func ValidarCorreo(correo string) error {
    regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
    if matched, _ := regexp.MatchString(regex, correo); !matched {
        return errors.New("Formato de correo inválido")
    }
    return nil
}

// Valida telefono de 10 dígitos
func ValidarTelefono(telefono string) error {
    if len(telefono) != 10 {
        return errors.New("El teléfono debe tener 10 dígitos")
    }
    return nil
}

// Valida contraseña
func ValidarPassword(password string) error {
    if len(password) < 6 || len(password) > 12 {
        return errors.New("La contraseña debe tener entre 6 y 12 caracteres")
    }

    var (
        hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
        hasLower   = regexp.MustCompile(`[a-z]`).MatchString
        hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
        hasSpecial = regexp.MustCompile(`[@$&]`).MatchString
    )

    if !hasUpper(password) {
        return errors.New("La contraseña debe incluir al menos una mayúscula")
    }
    if !hasLower(password) {
        return errors.New("La contraseña debe incluir al menos una minúscula")
    }
    if !hasNumber(password) {
        return errors.New("La contraseña debe incluir al menos un número")
    }
    if !hasSpecial(password) {
        return errors.New("La contraseña debe incluir al menos un carácter especial (@ $ &)")
    }

    return nil
}