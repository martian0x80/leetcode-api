package responses

import "encoding/json"

type QuestionEditorData struct {
	Data QuestionEditorDataData `json:"data"`
}

type QuestionEditorDataData struct {
	Question QuestionEditorDataQuestion `json:"question"`
}

type QuestionEditorDataQuestion struct {
	QuestionID         string                          `json:"questionId"`
	QuestionFrontendID string                          `json:"questionFrontendId"`
	CodeSnippets       []QuestionEditorDataCodeSnippet `json:"codeSnippets"`
	EnvInfo            string                          `json:"envInfo"`
	EnableRunCode      bool                            `json:"enableRunCode"`
	HasFrontendPreview bool                            `json:"hasFrontendPreview"`
	FrontendPreviews   string                          `json:"frontendPreviews"`
}

type QuestionEditorDataCodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

func (r QuestionEditorData) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r QuestionEditorData) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
