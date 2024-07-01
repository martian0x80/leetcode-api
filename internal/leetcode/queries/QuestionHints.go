package queries

import (
	"fmt"
	"io"
	"strings"
)

type QuestionHints struct {
	titleSlug string
}

func (q *QuestionHints) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query questionHints($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    hints\\n  }\\n}\",\"variables\":{\"titleSlug\":\"%s\"}}", q.titleSlug))
}
