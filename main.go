package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/peteli3/personal-v1/components"
)

func main() {
	http.Handle("/", templ.Handler(components.Homepage()))
	http.Handle("/htmx", templ.Handler(components.HTMXpage()))
	http.Handle("/daisyui", templ.Handler(components.DaisyUIpage()))

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	log.Println("Listening on :6969")
	http.ListenAndServe(":6969", nil)
}
