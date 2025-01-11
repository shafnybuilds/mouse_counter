package main

import (
	"example.com/mouse-counter/views"
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	component := views.Hello("World")

	http.Handle("/", templ.Handler(component))

	http.ListenAndServe(":8080", nil)
}
