package queries

import (
	"fmt"
	"io"
	"strings"
)

type UserProfileCalender struct {
	Username string
}

func (u *UserProfileCalender) GetQuery() io.Reader {
	return strings.NewReader(fmt.Sprintf("{\"query\":\"query userProfileCalendar($username: String!, $year: Int) {\\n  matchedUser(username: $username) {\\n    userCalendar(year: $year) {\\n      activeYears\\n      streak\\n      totalActiveDays\\n      dccBadges {\\n        timestamp\\n        badge {\\n          name\\n          icon\\n        }\\n      }\\n      submissionCalendar\\n    }\\n  }\\n}\",\"variables\":{\"username\":\"%s\"}}", u.Username))
}
