
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    // Mostrar las imagenes de la carpeta static
    http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("static"))))

    // Obtener el nombre del pod de Kubernetes
    hostname, err := os.Hostname()
    if err != nil {
        hostname = "desconocido"
    }

    // Entorno en el que corre la web (lo pone Jenkins al arrancar el contenedor)
    entorno := os.Getenv("ENTORNO")
    if entorno == "" {
        entorno = "produccion"
    }

    // Crear la pagina principal
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "text/html; charset=utf-8")

        fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <title>RED CLOUD - UOC</title>
        </head>
        <body style="text-align:center; font-family:Arial;">

            <h1>Soy alumno de la UOC</h1>

            <img src="/static/img/uoc.jpg"
                 alt="Imagen de la UOC"
                 width="350">

  <h2>RED CLOUD - Entorno: %s</h2>
            <p>Peticion atendida por el pod: <strong>%s</strong></p>

        </body>
        </html>
        `, entorno, hostname)
    })

    log.Println("Servidor iniciado en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

