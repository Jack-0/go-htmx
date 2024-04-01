package main

import (
	"fmt"
	"github.com/a-h/templ"
	"local/htmx-tt/internal/templates/components"
	"local/htmx-tt/internal/aws/dynamodb"
	"net/http"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func main() {
	// simple dynamo db
	dbService, err := dynamodb.NewDynamoDBService("us-west-2", "http://localhost:8000")
	if err != nil {
		fmt.Println("Error initializing DynamoDB service:", err)
		return
	}
	tables, err := dbService.ListTables()
	if err != nil {
		fmt.Println("Error listing tables:", err)
		return
	}
	fmt.Println("Tables:", tables)


	http.Handle("/", templ.Handler(components.TimeTable(components.CreateTimestable(12))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
