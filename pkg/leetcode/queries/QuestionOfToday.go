package queries

import (
	"io"
	"strings"
)

type QuestionOfToday struct {
}

func (u *QuestionOfToday) GetQuery() io.Reader {
	return strings.NewReader("{\"query\":\"query questionOfToday {\\n  activeDailyCodingChallengeQuestion {\\n    date\\n    userStatus\\n    link\\n    question {\\n      acRate\\n      difficulty\\n      freqBar\\n      frontendQuestionId: questionFrontendId\\n      isFavor\\n      paidOnly: isPaidOnly\\n      status\\n      title\\n      titleSlug\\n      hasVideoSolution\\n      hasSolution\\n      topicTags {\\n        name\\n        id\\n        slug\\n      }\\n    }\\n  }\\n}\",\"variables\":{}}")
}
