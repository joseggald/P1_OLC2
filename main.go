package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func saludarHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hola"))
}

type Mensaje struct {
	Contenido string `json:"contenido"`
}

func recibirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var mensaje Mensaje
		err := json.NewDecoder(r.Body).Decode(&mensaje)
		if err != nil {
			http.Error(w, "Error al leer el cuerpo de la solicitud JSON", http.StatusInternalServerError)
			return
		}

		// Cierra el cuerpo de la solicitud para liberar recursos
		defer r.Body.Close()

		contenido := mensaje.Contenido
		fmt.Printf("Contenido recibido del cliente: %s\n", contenido)

		// Aquí puedes realizar alguna acción con el contenido recibido, si lo deseas.

		// Responder al cliente que se recibió el contenido correctamente
		w.Write([]byte("Contenido recibido correctamente"))
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
    http.HandleFunc("/saludar", saludarHandler)
    http.HandleFunc("/recibir", recibirHandler) // Nuevo manejador para recibir el contenido POST
    http.Handle("/", http.FileServer(http.Dir("."))) // Servir archivos estáticos desde la carpeta actual
	fmt.Println("Servidor Activo")
    http.ListenAndServe(":8080", nil)
}