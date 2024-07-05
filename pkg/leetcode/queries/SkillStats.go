package queries

import (
	"fmt"
	"io"
	"strings"
)

type SkillStats struct {
	Username string
}

func (u *SkillStats) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query skillStats($username: String!) {\\n  matchedUser(username: $username) {\\n    tagProblemCounts {\\n      advanced {\\n        tagName\\n        tagSlug\\n        problemsSolved\\n      }\\n      intermediate {\\n        tagName\\n        tagSlug\\n        problemsSolved\\n      }\\n      fundamental {\\n        tagName\\n        tagSlug\\n        problemsSolved\\n      }\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
}
