package queries

import (
	"fmt"
	"io"
	"strings"
)

type RecentAcSubmissionList struct {
	Username string
	limit    int
}

func (u *RecentAcSubmissionList) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query recentAcSubmissions($username: String!, $limit: Int!) {\\n  recentAcSubmissionList(username: $username, limit: $limit) {\\n    id\\n    title\\n    titleSlug\\n    timestamp\\n  }\\n}\",\"variables\":{\"username\":\"%s\",\"limit\":%d}}", u.Username, u.limit))
}
