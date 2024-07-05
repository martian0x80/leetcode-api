package responses

import "encoding/json"

type UserContestRankingHistory struct {
	Data UserContestRankingHistoryData `json:"data"`
}

type UserContestRankingHistoryData struct {
	UserContestRankingDetails        UserContestRanking                 `json:"userContestRanking"`
	UserContestRankingHistoryDetails []UserContestRankingHistoryDetails `json:"userContestRankingHistory"`
}

type UserContestRankingHistoryDetails struct {
	Attended            bool    `json:"attended"`
	TrendDirection      string  `json:"trendDirection"`
	ProblemsSolved      int     `json:"problemsSolved"`
	TotalProblems       int     `json:"totalProblems"`
	FinishTimeInSeconds int     `json:"finishTimeInSeconds"`
	Rating              float64 `json:"rating"`
	Ranking             int     `json:"ranking"`
	Contest             Contest `json:"contest"`
}

type Contest struct {
	Title     string `json:"title"`
	StartTime int    `json:"startTime"`
}

func (r UserContestRankingHistory) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r UserContestRankingHistory) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
