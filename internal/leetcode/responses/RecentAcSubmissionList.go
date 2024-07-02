package responses

import "encoding/json"

type RecentAcSubmissionList struct {
	Data RecentAcSubmissionListData `json:"data"`
}

type RecentAcSubmissionListData struct {
	RecentAcSubmissionList []RecentAcSubmissionListItem `json:"recentAcSubmissionList"`
}

type RecentAcSubmissionListItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Timestamp string `json:"timestamp"`
}

func (r *RecentAcSubmissionList) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RecentAcSubmissionList) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
