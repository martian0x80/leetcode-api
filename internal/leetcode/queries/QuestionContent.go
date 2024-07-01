package queries

import (
	"fmt"
	"io"
	"strings"
)

type QuestionContent struct {
	titleSlug string
}

func (q *QuestionContent) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"\\nquery questionContent($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    content\\n    mysqlSchemas\\n    dataSchemas\\n  }\\n}\\n    \",\"variables\":{\"titleSlug\":\"%s\"}}", q.titleSlug))
}
