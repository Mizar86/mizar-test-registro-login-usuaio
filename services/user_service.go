package services

import (
    "errors"
    "mizar-test/models"
    "mizar-test/utils"

    "time"
    "github.com/golang-jwt/jwt/v4"
)

// Array para guardar los registros para pruebas
var users = []models.User{}

// Clave para fimar
var jwtKey = []byte("TGDdT435FGDsdA97D")

// Registro de usuario
func RegistrarUsuario(usuario, correo, telefono, password string) error {

    // Valida campos duplicados
    for _, user := range users {
        if user.Correo == correo {
            return errors.New("El correo ya esta registrado")
        }
        if user.Telefono == telefono {
            return errors.New("El teléfono ya esta registrado")
        }
    }
    
    // Validaciones de campos
    if err := utils.ValidarCorreo(correo); err != nil {
        return err
    }
    if err := utils.ValidarTelefono(telefono); err != nil {
        return err
    }
    if err := utils.ValidarPassword(password); err != nil {
        return err
    }
    
    // Guarda nuevo registro.
    user := models.User{ID: len(users) + 1, Usuario: usuario, Correo: correo, Telefono: telefono, Password: password}
    users = append(users, user)
    return nil
}

// Login de usuario
func LoginUser(usuarioCorreo, password string) (string, error) {
    for _, user := range users {
        if (user.Usuario == usuarioCorreo || user.Correo == usuarioCorreo) && user.Password == password {
            return generarJWT(user.Usuario, user.Correo, user.Telefono )
        }
    }
    return "", errors.New("Usuario / contraseña incorrectos")
}

// Generar token JWT
func generarJWT(usuario, correo, telefono string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "usuario": usuario,
        "correo": correo,
        "telefono": telefono,
        "exp": time.Now().Add(time.Hour * 1).Unix(),
    })

    // Firmar el token
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
