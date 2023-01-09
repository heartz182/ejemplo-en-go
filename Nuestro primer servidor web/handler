package main

import (
	"fmt"
	"net/http"
)

// holamundo es un handler que se encarga de procesar las solicitudes HTTP
// que se reciben en la ruta "/" y de preparar una respuesta adecuada.
func holamundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hola Mundo :3</>")
}

// prueba es un handler que se encarga de procesar las solicitudes HTTP
// que se reciben en la ruta "/prueba" y de preparar una respuesta adecuada.
func prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hola Mundo desde /prueba :)</>")
}

// usuario es un handler que se encarga de procesar las solicitudes HTTP
// que se reciben en la ruta "/usuario" y de preparar una respuesta adecuada.
func usuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hola desde usuario ;)</>")
}

func main() {
	// Creamos un nuevo objeto ServeMux para gestionar las rutas y asociarlas a los handlers correspondientes.
	mux := http.NewServeMux()

	// Asignamos cada handler a la ruta correspondiente utilizando el método HandleFunc del objeto ServeMux.
	mux.HandleFunc("/", holamundo)
	mux.HandleFunc("/prueba", prueba)
	mux.HandleFunc("/usuario", usuario)

	// Iniciamos el servidor web en el puerto 8080 y utilizamos el objeto ServeMux como manejador de rutas.
	http.ListenAndServe(":8080", mux)
}
