package responses

import "encoding/json"

type QuestionHints struct {
	Data QuestionHintsData `json:"data"`
}

type QuestionHintsData struct {
	Question QuestionHintsQuestion `json:"question"`
}

type QuestionHintsQuestion struct {
	Hints []string `json:"hints"`
}

func (r *QuestionHints) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *QuestionHints) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
