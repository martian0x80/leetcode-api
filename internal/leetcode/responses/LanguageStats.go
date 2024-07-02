package responses

import "encoding/json"

type LanguageStats struct {
	Data LanguageStatsData `json:"data"`
}

type LanguageStatsData struct {
	MatchedUser MatchedUserLanguages `json:"matchedUser"`
}

type MatchedUserLanguages struct {
	LanguageProblemCount []LanguageProblemCount `json:"languageProblemCount"`
}

type LanguageProblemCount struct {
	LanguageName   string `json:"languageName"`
	ProblemsSolved int    `json:"problemsSolved"`
}

func (r *LanguageStats) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *LanguageStats) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
