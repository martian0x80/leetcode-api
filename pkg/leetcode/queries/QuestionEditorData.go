package queries

import (
	"fmt"
	"io"
	"strings"
)

type QuestionEditorData struct {
	titleSlug string
}

func (q *QuestionEditorData) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query questionEditorData($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    codeSnippets {\\n      lang\\n      langSlug\\n      code\\n    }\\n    envInfo\\n    enableRunCode\\n    hasFrontendPreview\\n    frontendPreviews\\n  }\\n}\",\"variables\":{\"titleSlug\":\"%s\"}}", q.titleSlug))
}
