package queries

import (
	"fmt"
	"io"
	"strings"
)

type UserProblemsSolved struct {
	Username string
}

func (u *UserProblemsSolved) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query userProblemsSolved($username: String!) {\\n  allQuestionsCount {\\n    difficulty\\n    count\\n  }\\n  matchedUser(username: $username) {\\n    problemsSolvedBeatsStats {\\n      difficulty\\n      percentage\\n    }\\n    submitStatsGlobal {\\n      acSubmissionNum {\\n        difficulty\\n        count\\n      }\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
}
