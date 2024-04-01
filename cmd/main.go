package main

import (
	"fmt"
	"local/htmx-tt/internal/templates/components"
	"net/http"
	"github.com/a-h/templ"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func main() {
	http.Handle("/", templ.Handler(components.TimeTable(components.CreateTimestable(12))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
