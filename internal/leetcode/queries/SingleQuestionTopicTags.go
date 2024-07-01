package queries

import (
	"fmt"
	"io"
	"strings"
)

type SingleQuestionTopicTags struct {
	titleSlug string
}

func (q *SingleQuestionTopicTags) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query singleQuestionTopicTags($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    topicTags {\\n      name\\n      slug\\n    }\\n  }\\n}\",\"variables\":{\"titleSlug\":\"%s\"}}", q.titleSlug))
}
