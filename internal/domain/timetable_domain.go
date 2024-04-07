package domain

import "time"

type TTCurrentQuestion struct {
	StartTime time.Time
	EndTime   time.Time
	Success   bool
	ValueA    uint16
	ValueB    uint16
}

type TTAnswer struct {
	Pk           string `json:"pk"`
	Operation    string `json:"operation"`
	Timestamp    string `json:"timestamp"`
	AnswerTimeMs uint32 `json:"answerTimeMs"`
	Success      bool   `json:"successs"`
}

type TTQuestion struct {
	ValueA          uint16 `json:"valueA"`
	ValueB          uint16 `json:"valueB"`
	NormalizedScore uint8  `json:"normalizedScore"`
	AnsweredTime    string `json:"answeredTime"`
	Ignored         bool   `json:"ignored"`
}

type QuestionList []TTQuestion

type TTList struct {
	Pk        string       `json:"pk"`
	Operation string       `json:"operation"`
	Questions QuestionList `json:"questions"`
}

func (ql QuestionList) Less(i, j int) bool {
	// set highest normalised score to be first
	return ql[i].NormalizedScore < ql[j].NormalizedScore
}

func (ql QuestionList) Len() int {
	return len(ql)
}
func (ql QuestionList) Swap(i, j int) {
    ql[i], ql[j] = ql[j], ql[i]
}
