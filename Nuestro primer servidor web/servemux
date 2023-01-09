package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Creamos un nuevo objeto ServeMux para gestionar las rutas y asociarlas a los handlers correspondientes.
	mux := http.NewServeMux()

	// Asignamos cada handler anónimo a la ruta correspondiente utilizando el método HandleFunc del objeto ServeMux.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hola Mundo :3</>")
	})
	mux.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hola Mundo desde /prueba :)</>")
	})
	mux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hola desde usuario ;)</>")
	})

	// Iniciamos el servidor web en el puerto 8080 y utilizamos el objeto ServeMux como manejador de rutas.
	http.ListenAndServe(":8080", mux)
}
