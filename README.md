# Proyecto Backend en Golang

Backend en Golang que incluye funcionalidades de autenticación básica, como registro de usuarios y login, utilizando JWT para la autenticación.

## Características

*  Registro de usuario que recibe como parámetros usuario, correo, teléfono y contraseña.
*  Valida que el correo y telefono no se encuentren registrados, de lo contrario deberá retornar un mensaje “el correo/telefono ya se encuentra registrado”.
*  Valida que la contraseña sea de 6 caracteres mínimo y 12 máximo y contener al menos una mayúscula, una minúscula, un carácter especial (@ $ o &) y un número.
*  Valida que el teléfono sea a 10 dígitos y el correo tenga un formato válido.
*  Login recibe como parámetros usuario o correo y contraseña.
*  Login retorna un token jwt.
*  Valida que el usuario o correo y contraseña sean válidos, de lo contrario retorna un mensaje “usuario / contraseña incorrectos”.
*  Valida que todos los parámetros solicitados vayan en el cuerpo de la petición, de lo contrario retorna un mensaje con el campo faltante.


### Endpoints

1. **Registro de Usuario**

    - **URL**: `/register`
    - **Método**: `POST`
    - **Parámetros**: JSON en el cuerpo de la solicitud con los siguientes campos:

      ```json
      {
          "usuario": "nombre_usuario",
          "correo": "usuario@example.com",
          "telefono": "1234567890",
          "password": "Ffds123$"
      }
      ```

2. **Login de Usuario**

    - **URL**: `/login`
    - **Método**: `POST`
    - **Parámetros**: JSON en el cuerpo de la solicitud con los siguientes campos:

      ```json
      {
          "usuario": "nombre_usuario",
          "password": "ContraseñaSegura123$"
      }
      ```

    - **Respuesta**: Un token JWT para autenticar solicitudes futuras.

## Ejecutar Pruebas

comando:

```bash
go test -v ./...
