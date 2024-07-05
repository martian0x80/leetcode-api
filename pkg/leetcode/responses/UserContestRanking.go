package responses

import "encoding/json"

type UserContestRanking struct {
	Data UserContestRankingData `json:"data"`
}

type UserContestRankingData struct {
	UserContestRankingDetails UserContestRankingDetails `json:"userContestRanking"`
}

type UserContestRankingDetails struct {
	AttendedContestsCount int     `json:"attendedContestsCount"`
	Rating                float64 `json:"rating"`
	GlobalRanking         int     `json:"globalRanking"`
	TotalParticipants     int     `json:"totalParticipants"`
	TopPercentage         float64 `json:"topPercentage"`
	Badge                 *string `json:"badge"`
}

func (r UserContestRanking) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r UserContestRanking) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
