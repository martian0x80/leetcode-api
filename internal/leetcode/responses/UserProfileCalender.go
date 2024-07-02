package responses

import "encoding/json"

type UserProfileCalender struct {
	Data struct {
		MatchedUser struct {
			UserCalendar struct {
				ActiveYears        []int         `json:"activeYears"`
				Streak             int           `json:"streak"`
				TotalActiveDays    int           `json:"totalActiveDays"`
				DccBadges          []interface{} `json:"dccBadges"`
				SubmissionCalendar string        `json:"submissionCalendar"`
			} `json:"userCalendar"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func (r *UserProfileCalender) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *UserProfileCalender) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
