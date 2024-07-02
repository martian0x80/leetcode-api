package responses

import "encoding/json"

type SkillStats struct {
	Data SkillStatsData `json:"data"`
}

type SkillStatsData struct {
	MatchedUser MatchedUserTags `json:"matchedUser"`
}

type MatchedUserTags struct {
	TagProblemCounts TagProblemCounts `json:"tagProblemCounts"`
}

type TagProblemCounts struct {
	Advanced     []TagProblemCount `json:"advanced"`
	Intermediate []TagProblemCount `json:"intermediate"`
	Fundamental  []TagProblemCount `json:"fundamental"`
}

type TagProblemCount struct {
	TagName        string `json:"tagName"`
	TagSlug        string `json:"tagSlug"`
	ProblemsSolved int    `json:"problemsSolved"`
}

func (r *SkillStats) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *SkillStats) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
