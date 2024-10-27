package handlers

import (
    "encoding/json"
    "net/http"
    "mizar-test/services"
)

type RegisterRequest struct {
    Usuario  string `json:"usuario"`
    Correo   string `json:"correo"`
    Telefono string `json:"telefono"`
    Password string `json:"password"`
}

// Estructura para el login
type LoginRequest struct {
    Usuario string `json:"usuario"`
    Password   string `json:"password"`
}

// Estructura para respuesta de login
type LoginResponse struct {
    Token string `json:"token"`
}

// Función para registro.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Error en la solicitud", http.StatusBadRequest)
        return
    }

    // Validar campos requeridos
    if req.Usuario == "" {
        http.Error(w, "Falta el usuario", http.StatusBadRequest)
        return
    }
    if req.Correo == "" {
        http.Error(w, "Falta el correo", http.StatusBadRequest)
        return
    }
    if req.Telefono == "" {
        http.Error(w, "Falta el telefono", http.StatusBadRequest)
        return
    }
    if req.Password == "" {
        http.Error(w, "Falta el password", http.StatusBadRequest)
        return
    }

    // Se registra el usuario
    if err := services.RegistrarUsuario(req.Usuario, req.Correo, req.Telefono, req.Password); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Registrado con éxito"))
}

// Función para login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest

    // Decodifica el cuerpo de la solicitud
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Solicitud inválida", http.StatusBadRequest)
        return
    }

    // Valida campos obligatorios.
    if req.Usuario == "" {
        http.Error(w, "Falta el Usuario/Correo", http.StatusBadRequest)
        return
    }
    if req.Password == "" {
        http.Error(w, "Falta el password", http.StatusBadRequest)
        return
    }

    // Autenticar el usuario
    token, err := services.LoginUser(req.Usuario, req.Password)
    if err != nil {
        // Si las credenciales son incorrectas, retornamos un mensaje de error.
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Regresamos el TokenJWT
    resp := LoginResponse{Token: token}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}