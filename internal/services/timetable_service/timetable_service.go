package timetable_service

import (
	"local/htmx-tt/internal/domain"
	"math/rand"
	"sort"
	"time"
)

// to be used by the

type TimeTable struct {
	UserId          string
	QuestionList    domain.QuestionList
	CurrentQuestion domain.TTCurrentQuestion
}

func NewTTService() *TimeTable {
	tt := TimeTable{
		UserId:       "jack",
		QuestionList: setDefaultTT(),
	}
	return &tt
}

func randomizeList(list domain.QuestionList) domain.QuestionList {
	newList := make(domain.QuestionList, len(list))
	copy(newList, list)
	// fisher-Yates shuffle algorithm
	for i := len(newList) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		newList[i], newList[j] = newList[j], newList[i]
	}
	return newList
}

func setDefaultTT() domain.QuestionList {
	const limit = 12
	var questions domain.QuestionList
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			newQuestion := domain.TTQuestion{
				ValueA:          uint16(i),
				ValueB:          uint16(j),
				Ignored:         false,
				NormalizedScore: 6, // highest score == next to attempt
			}
			questions = append(questions, newQuestion)
		}
	}
	return randomizeList(questions)
}

func (tt *TimeTable) GetQuestion() domain.TTCurrentQuestion {
	println("GETq")
	// handle default empty state
	if len(tt.QuestionList) == 0 {
		// TODO try and get from dynamo later...
		tt.QuestionList = setDefaultTT()
		println("new rand list is: ", tt.QuestionList)
	} else {
		// else sort current list
		sort.Sort(tt.QuestionList)
	}

	// get next q... todo
	nextQ := tt.QuestionList[0]

	startTime := time.Now()
	q := domain.TTCurrentQuestion{
		ValueA:    nextQ.ValueA,
		ValueB:    nextQ.ValueA,
		StartTime: startTime,
	}
	println("nextq", q.ValueA, q.ValueB)
	tt.CurrentQuestion = q
	println("tt", tt.CurrentQuestion.ValueA, tt.CurrentQuestion.ValueB)
	return q
}

func (tt *TimeTable) AnswerQuestion(answer uint16) {
	success := tt.CurrentQuestion.ValueA*tt.CurrentQuestion.ValueB == answer
	println(tt.CurrentQuestion.ValueA, "*", tt.CurrentQuestion.ValueA, "==", answer)
	endTime := time.Now()
	timeMs := endTime.Sub(tt.CurrentQuestion.StartTime).Milliseconds()
	println("question correct= ", success, " timeMs=", timeMs)

	tt.QuestionList = tt.QuestionList[1:] //TODO actuall do this just testing
}

// return a value between 0 and 5 where 0 is excellent and 5 is poor
func NormalizeResult(answer domain.TTAnswer, userAvgAnswerTimeMs uint32) uint8 {
	const bound = 0.20 // 20%
	lowerAvgBound := float64(userAvgAnswerTimeMs) * (1 - bound)
	upperAvgBound := float64(userAvgAnswerTimeMs) * (1 + bound)
	inBounds := answer.AnswerTimeMs >= uint32(lowerAvgBound) && answer.AnswerTimeMs <= uint32(upperAvgBound)

	score := uint8(0)
	if inBounds {
		score += 1
	} else if answer.AnswerTimeMs > uint32(upperAvgBound) {
		score += 2
	}

	if !answer.Success {
		score += 3
	}
	return score
}
