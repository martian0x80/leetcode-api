package queries

import (
	"fmt"
	"io"
	"strings"
)

type QuestionStats struct {
	titleSlug string
}

func (q *QuestionStats) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query questionStats($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    stats\\n  }\\n}\",\"variables\":{\"titleSlug\":\"%s\"}}", q.titleSlug))
}
