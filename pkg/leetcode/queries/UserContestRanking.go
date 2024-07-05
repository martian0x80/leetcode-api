package queries

import (
	"fmt"
	"io"
	"strings"
)

type UserContestRanking struct {
	Username string
}

func (u *UserContestRanking) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query userContestRankingInfo($username: String!) {\\n  userContestRanking(username: $username) {\\n    attendedContestsCount\\n    rating\\n    globalRanking\\n    totalParticipants\\n    topPercentage\\n    badge {\\n      name\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
}
