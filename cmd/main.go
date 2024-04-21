package main

import (
	"encoding/json"
	"fmt"
	"local/htmx-tt/internal/aws/dynamodb"
	"local/htmx-tt/internal/domain"
	"local/htmx-tt/internal/services/timetable_service"
	"local/htmx-tt/internal/templates/components"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
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

	// tt service
	tt := timetable_service.NewTTService()

	fmt.Println("Tables:", tables)

	// http.Handle("/", templ.Handler(components.TimeTable(components.CreateTimestable(12)))) // TODO conflict with other roots?
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// http.Handle("/answer", answerQuestion(tt))
	http.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		answerQuestion(w, r, tt)
	})
	http.Handle("/question", returnQuestion(tt))

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// Call handlePostRequest with the service instance
		handlePostRequest(w, r, dbService)
	})

	// http.Handle("/users", templ.Handler(Pages.Users(User.Users, User.SelectedUser)))
	http.HandleFunc("/setUser", updateUser)

	fmt.Println("Listening on :3000")
	http.ListenAndServe("localhost:3000", nil)
}

func returnQuestion(tt *timetable_service.TimeTable) *templ.ComponentHandler {
	q := tt.GetQuestion()
	return templ.Handler(components.QuestionView(q))
}

func answerQuestion(w http.ResponseWriter, r *http.Request, tt *timetable_service.TimeTable) {
	println("answer q")
	err := r.ParseForm()
	if err != nil {
		println("parse err", err)
	}
	answerInput := r.FormValue("numberInput")

	answer, _ := strconv.Atoi(answerInput)

	println("inputIs", answer, answerInput)

	tt.AnswerQuestion(uint16(answer))

	q := tt.GetQuestion()

	println("next Q valueA is:", q.ValueA)

	components.Question(q).Render(r.Context(), w)
}

type RequestBody struct {
	Int1 int `json:"int1"`
	Int2 int `json:"int2"`
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	println("setting user")
	// w.Header().Set("Content-Type", "application/json")
	// Write JSON data to response writer
	// data, _ := json.Marshal(users)
	// w.Write(data)
	w.Write([]byte("hello"))
}

func handlePostRequest(w http.ResponseWriter, r *http.Request, db *dynamodb.DynamoDBService) {
	// Decode JSON body
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	// test := timetable_service.TTResult{TimeMs: uint32(requestBody.Int1), Success: false}
	// x := timetable_service.NormalizeResult(test, uint32(requestBody.Int2))

	tt := domain.TTAnswer{
		Pk:           "jack#multiplication",
		Operation:    "7x3",
		Timestamp:    time.Now().Format(time.RFC850),
		AnswerTimeMs: 200,
		Success:      false,
	}
	db.AddItem("numbers", tt)

	// w.Write([]byte(strconv.Itoa(int(x))))
}
