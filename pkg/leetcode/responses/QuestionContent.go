package responses

import "encoding/json"

type QuestionContentData struct {
	Question QuestionContentQuestion `json:"question"`
}

type QuestionContentQuestion struct {
	Content      string        `json:"content"`
	MySQLSchemas []interface{} `json:"mysqlSchemas"` // Assuming empty array or unknown types, hence using interface{}
	DataSchemas  []interface{} `json:"dataSchemas"`  // Assuming empty array or unknown types, hence using interface{}
}

type QuestionContent struct {
	Data QuestionContentData `json:"data"`
}

func (r QuestionContent) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r QuestionContent) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
