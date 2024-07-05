package queries

import (
	"fmt"
	"io"
	"strings"
)

type LanguageStats struct {
	Username string
}

func (u *LanguageStats) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query languageStats($username: String!) {\\n  matchedUser(username: $username) {\\n    languageProblemCount {\\n      languageName\\n      problemsSolved\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
}

//type GetStreakCounter struct {
//	Username string
//}
//
//func (u *GetStreakCounter) GetQuery() io.Reader {
//	return strings.NewReader(fmt.Sprintf("{\"query\":\"query getStreakCounter {\\n  streakCounter {\\n    streakCount\\n    daysSkipped\\n    currentDayCompleted\\n  }\\n}\\n    \",\"variables\":{}}"))
//}
