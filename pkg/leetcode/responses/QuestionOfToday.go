package responses

import "encoding/json"

// WHAT THE FUCKKK
// jetbrains - goland supports generating go structs from just the json
// i have been doing this manually, whhaaat the fuck, fuck me
// time it get out of neovim i suppose

type QuestionOfToday struct {
	Data struct {
		ActiveDailyCodingChallengeQuestion struct {
			Date       string `json:"date"`
			UserStatus string `json:"userStatus"`
			Link       string `json:"link"`
			Question   struct {
				AcRate             float64     `json:"acRate"`
				Difficulty         string      `json:"difficulty"`
				FreqBar            interface{} `json:"freqBar"`
				FrontendQuestionId string      `json:"frontendQuestionId"`
				IsFavor            bool        `json:"isFavor"`
				PaidOnly           bool        `json:"paidOnly"`
				Status             interface{} `json:"status"`
				Title              string      `json:"title"`
				TitleSlug          string      `json:"titleSlug"`
				HasVideoSolution   bool        `json:"hasVideoSolution"`
				HasSolution        bool        `json:"hasSolution"`
				TopicTags          []struct {
					Name string `json:"name"`
					Id   string `json:"id"`
					Slug string `json:"slug"`
				} `json:"topicTags"`
			} `json:"question"`
		} `json:"activeDailyCodingChallengeQuestion"`
	} `json:"data"`
}

func (r QuestionOfToday) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r QuestionOfToday) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
