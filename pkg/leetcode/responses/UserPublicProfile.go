package responses

import (
	"encoding/json"
)

type UserPublicProfile struct {
	Data UserPublicProfileData `json:"data"`
}

type UserPublicProfileData struct {
	MatchedUser UserPublicProfileMatchedUser `json:"matchedUser"`
}

type UserPublicProfileMatchedUser struct {
	ContestBadge interface{}              `json:"contestBadge"` // Assuming null values can be any type, hence using interface{}
	Username     string                   `json:"username"`
	GithubURL    string                   `json:"githubUrl"`
	TwitterURL   interface{}              `json:"twitterUrl"`  // Assuming null values can be any type, hence using interface{}
	LinkedinURL  interface{}              `json:"linkedinUrl"` // Assuming null values can be any type, hence using interface{}
	Profile      UserPublicProfileProfile `json:"profile"`
}

type UserPublicProfileProfile struct {
	Ranking                  int         `json:"ranking"`
	UserAvatar               string      `json:"userAvatar"`
	RealName                 string      `json:"realName"`
	AboutMe                  string      `json:"aboutMe"`
	School                   interface{} `json:"school"` // Assuming null values can be any type, hence using interface{}
	Websites                 []string    `json:"websites"`
	CountryName              interface{} `json:"countryName"` // Assuming null values can be any type, hence using interface{}
	Company                  interface{} `json:"company"`     // Assuming null values can be any type, hence using interface{}
	JobTitle                 interface{} `json:"jobTitle"`    // Assuming null values can be any type, hence using interface{}
	SkillTags                []string    `json:"skillTags"`
	PostViewCount            int         `json:"postViewCount"`
	PostViewCountDiff        int         `json:"postViewCountDiff"`
	Reputation               int         `json:"reputation"`
	ReputationDiff           int         `json:"reputationDiff"`
	SolutionCount            int         `json:"solutionCount"`
	SolutionCountDiff        int         `json:"solutionCountDiff"`
	CategoryDiscussCount     int         `json:"categoryDiscussCount"`
	CategoryDiscussCountDiff int         `json:"categoryDiscussCountDiff"`
}

func (r UserPublicProfile) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r UserPublicProfile) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
