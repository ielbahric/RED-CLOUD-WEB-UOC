package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Permitir que Go muestre las imagenes de la carpeta static
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// Crear la pagina principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintln(w, `
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <title>Mi primera web con Go</title>
        </head>
        <body style="text-align:center; font-family:Arial;">

            <h1>Soy alumno de la UOC</h1>

            <img src="/static/img/uoc.jpg"
                 alt="Imagen de la UOC"
                 width="350">

        </body>
        </html>
        `)
	})

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
