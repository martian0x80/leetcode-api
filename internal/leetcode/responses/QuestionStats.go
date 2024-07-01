package responses

import (
	"encoding/json"
)

type QuestionStats struct {
	Data QuestionStatsData `json:"data"`
}

type QuestionStatsData struct {
	Question QuestionStatsQuestion `json:"question"`
}

type QuestionStatsQuestion struct {
	Stats string `json:"stats"`
}

func (r *QuestionStats) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *QuestionStats) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
