package responses

import "encoding/json"

type QuestionTitle struct {
	Data QuestionTitleData `json:"data"`
}

type QuestionTitleData struct {
	Question QuestionTitleDetails `json:"question"`
}

type QuestionTitleDetails struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	IsPaidOnly         bool   `json:"isPaidOnly"`
	Difficulty         string `json:"difficulty"`
	Likes              int    `json:"likes"`
	Dislikes           int    `json:"dislikes"`
	CategoryTitle      string `json:"categoryTitle"`
}

func (r *QuestionTitle) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *QuestionTitle) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
