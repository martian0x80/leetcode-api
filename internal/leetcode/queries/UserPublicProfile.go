package queries

import (
	"fmt"
	"io"
	"strings"
)

type UserPublicProfile struct {
	Username string
}

func (u *UserPublicProfile) GetQuery() io.Reader {
	payload := strings.NewReader(fmt.Sprintf("{\"query\":\"query userPublicProfile($username: String!) {\\n  matchedUser(username: $username) {\\n    contestBadge {\\n      name\\n      expired\\n      hoverText\\n      icon\\n    }\\n    username\\n    githubUrl\\n    twitterUrl\\n    linkedinUrl\\n    profile {\\n      ranking\\n      userAvatar\\n      realName\\n      aboutMe\\n      school\\n      websites\\n      countryName\\n      company\\n      jobTitle\\n      skillTags\\n      postViewCount\\n      postViewCountDiff\\n      reputation\\n      reputationDiff\\n      solutionCount\\n      solutionCountDiff\\n      categoryDiscussCount\\n      categoryDiscussCountDiff\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
	return payload
}
