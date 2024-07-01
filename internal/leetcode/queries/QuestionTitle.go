package queries

import (
	"fmt"
	"io"
	"strings"
)

type QuestionTitle struct {
	titleSlug string
}

func (q *QuestionTitle) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query questionTitle($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    title\\n    titleSlug\\n    isPaidOnly\\n    difficulty\\n    likes\\n    dislikes\\n    categoryTitle\\n  }\\n}\",\"variables\":{\"titleSlug\":\"%s\"}}", q.titleSlug))
}
