package responses

import "encoding/json"

type UserProblemsSolved struct {
	Data UserProblemsSolvedData `json:"data"`
}

type UserProblemsSolvedData struct {
	AllQuestionsCount []UserProblemsSolvedQuestionCount `json:"allQuestionsCount"`
	MatchedUser       UserProblemsSolvedMatchedUser     `json:"matchedUser"`
}

type UserProblemsSolvedQuestionCount struct {
	Difficulty string `json:"difficulty"`
	Count      int    `json:"count"`
}

type UserProblemsSolvedMatchedUser struct {
	ProblemsSolvedBeatsStats []UserProblemsSolvedProblemSolvedBeatStat `json:"problemsSolvedBeatsStats"`
	SubmitStatsGlobal        UserProblemsSolvedSubmitStatsGlobal       `json:"submitStatsGlobal"`
}

type UserProblemsSolvedProblemSolvedBeatStat struct {
	Difficulty string  `json:"difficulty"`
	Percentage float64 `json:"percentage"`
}

type UserProblemsSolvedSubmitStatsGlobal struct {
	AcSubmissionNum []UserProblemsSolvedAcSubmission `json:"acSubmissionNum"`
}

type UserProblemsSolvedAcSubmission struct {
	Difficulty string `json:"difficulty"`
	Count      int    `json:"count"`
}

func (r UserProblemsSolved) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r UserProblemsSolved) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
