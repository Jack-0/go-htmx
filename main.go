package main

import (
	"fmt"
	"github.com/a-h/templ"
	"local/htmx-tt/templates/components"
	"net/http"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func main() {

	http.Handle("/", templ.Handler(components.TimeTable(components.CreateTimetable(12))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
