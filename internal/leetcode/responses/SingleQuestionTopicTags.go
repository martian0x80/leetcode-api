package responses

import "encoding/json"

type SingleQuestionTopicTags struct {
	Data SingleQuestionTopicTagsData `json:"data"`
}

type SingleQuestionTopicTagsData struct {
	Question SingleQuestionTopicTagsQuestion `json:"question"`
}

type SingleQuestionTopicTagsQuestion struct {
	TopicTags []SingleQuestionTopicTagsTopicTag `json:"topicTags"`
}

type SingleQuestionTopicTagsTopicTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r *SingleQuestionTopicTags) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *SingleQuestionTopicTags) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
