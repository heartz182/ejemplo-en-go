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

// mensaje es una estructura que se utiliza para implementar la interfaz http.Handler.
type mensaje struct {
	msg string
}

// ServeHTTP es el método de la interfaz http.Handler que se encarga de procesar las solicitudes HTTP
// que se reciben en la ruta "/hola" y de preparar una respuesta adecuada.
func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.msg)
}

func main() {
	// Creamos una nueva instancia de la estructura mensaje.
	msg := mensaje{
		msg: "hola mundo de nuevo",
	}

	// Creamos un nuevo objeto ServeMux para gestionar las rutas y asociarlas a los handlers correspondientes.
	mux := http.NewServeMux()

	// Asignamos cada handler a la ruta correspondiente utilizando el método HandleFunc del objeto ServeMux.
	mux.HandleFunc("/", holamundo)
	mux.HandleFunc("/prueba", prueba)
	mux.HandleFunc("/usuario", usuario)

	// Asignamos la instancia de la estructura mensaje como handler de la ruta "/hola".
	mux.Handle("/hola", msg)

	// Iniciamos el servidor HTTP en el puerto 8080 y utilizamos el objeto ServeMux como gestor de rutas.
	http.ListenAndServe(":8080", mux)
}
