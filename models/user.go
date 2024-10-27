package models

type User struct {
    ID       int    `json:"id"`
    Usuario  string `json:"usuario"`
    Correo   string `json:"correo"`
    Telefono string `json:"telefono"`
    Password string `json:"password"`
}
