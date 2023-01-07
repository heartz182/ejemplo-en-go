package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Asignamos un handler anónimo a la ruta raíz ("/").
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Este handler escribe un mensaje simple en la respuesta HTTP.
		fmt.Fprintf(w, "<h1> Hola Mundo </h1>")
	})

	// Iniciamos el servidor web en el puerto 8080.
	http.ListenAndServe(":8080", nil)
}
