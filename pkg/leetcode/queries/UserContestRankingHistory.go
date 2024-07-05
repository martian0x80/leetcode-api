package queries

import (
	"fmt"
	"io"
	"strings"
)

type UserContestRankingHistory struct {
	Username string
}

func (u *UserContestRankingHistory) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query userContestRankingInfo($username: String!) {\\n  userContestRanking(username: $username) {\\n    attendedContestsCount\\n    rating\\n    globalRanking\\n    totalParticipants\\n    topPercentage\\n    badge {\\n      name\\n    }\\n  }\\n  userContestRankingHistory(username: $username) {\\n    attended\\n    trendDirection\\n    problemsSolved\\n    totalProblems\\n    finishTimeInSeconds\\n    rating\\n    ranking\\n    contest {\\n      title\\n      startTime\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
}
