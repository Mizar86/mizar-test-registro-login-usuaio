package main

import (
    "log"
    "net/http"
    "mizar-test/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	// Routes
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("GET")

    // Iniciar el servidor HTTP en el puerto 8080
    log.Println("Servidor iniciado.")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
