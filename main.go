package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
	"local/htmx-tt/components"
)

type GlobalState struct {
	Count int
}

var global GlobalState
	

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := components.Counter(global.Count, 0)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state.
	r.ParseForm()

	// Check to see if the global button was pressed.
	if r.Form.Has("global") {
		global.Count++
	}
	//TODO: Update session.

	// Display the form.
	getHandler(w, r)
}


func main() {
	component := components.Hello("Jack")
	
	http.Handle("/", templ.Handler(component))
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		getHandler(w, r)
	})


	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
