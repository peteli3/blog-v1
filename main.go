package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/peteli3/personal-v1/components"
)

var counters components.Counters
var formData components.FormData
var speciesChoices = []string{
	"armadillo",
	"capybara",
	"platypus",
}

func doHTMXGet(w http.ResponseWriter, r *http.Request) {
	props := components.HTMXpageProps{
		Counts:         counters,
		Forms:          formData,
		SpeciesChoices: speciesChoices,
	}
	components.HTMXpage(props).Render(r.Context(), w)
}

func doHTMXIncTemplPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Has("incTempl") {
		counters.Templ++
	}
	// use Redirect instead of GET to implement PRG pattern
	// which prevents unwanted form submission when user refreshes a POST
	http.Redirect(w, r, "/htmx", http.StatusSeeOther)
}

func doHTMXIncHTMXPost(w http.ResponseWriter, r *http.Request) {
	counters.HTMX++
	http.Redirect(w, r, "/htmx/incHTMX/success", http.StatusSeeOther)
}

func doHTMXIncHTMXSuccess(w http.ResponseWriter, r *http.Request) {
	components.HTMXCounter(counters).Render(r.Context(), w)
}

func doHTMXSubmitForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formData.Name = r.Form.Get("formName")
	formData.Species = r.Form.Get("formSpecies")
	formData.FavoriteFood = r.Form.Get("formFavoriteFood")
	http.Redirect(w, r, "/htmx/submitForm/success", http.StatusSeeOther)
}

func doHTMXSubmitFormSuccess(w http.ResponseWriter, r *http.Request) {
	components.FormSuccess(formData).Render(r.Context(), w)
}

func main() {
	http.Handle("/", templ.Handler(components.Homepage()))

	// htmx playground
	http.HandleFunc("/htmx", func(w http.ResponseWriter, r *http.Request) {
		doHTMXGet(w, r)
	})
	http.HandleFunc("/htmx/incTmpl", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			doHTMXIncTemplPost(w, r)
		}
	})
	http.HandleFunc("/htmx/incHTMX", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			doHTMXIncHTMXPost(w, r)
		}
	})
	http.HandleFunc("/htmx/incHTMX/success", func(w http.ResponseWriter, r *http.Request) {
		doHTMXIncHTMXSuccess(w, r)
	})
	http.HandleFunc("/htmx/submitForm", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			doHTMXSubmitForm(w, r)
		}
	})
	http.HandleFunc("/htmx/submitForm/success", func(w http.ResponseWriter, r *http.Request) {
		doHTMXSubmitFormSuccess(w, r)
	})

	// daisyui playground
	http.Handle("/daisyui", templ.Handler(components.DaisyUIpage()))

	// static content
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	log.Println("Listening on :6969")
	http.ListenAndServe(":6969", nil)
}
