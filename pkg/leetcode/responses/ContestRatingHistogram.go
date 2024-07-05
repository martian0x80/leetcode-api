package responses

import "encoding/json"

type ContestRatingHistogram struct {
	Data ContestRatingHistogramData `json:"data"`
}

type ContestRatingHistogramData struct {
	ContestRatingHistogram []RatingHistogramEntry `json:"contestRatingHistogram"`
}

type RatingHistogramEntry struct {
	UserCount     int     `json:"userCount"`
	RatingStart   int     `json:"ratingStart"`
	RatingEnd     int     `json:"ratingEnd"`
	TopPercentage float64 `json:"topPercentage"`
}

func (r ContestRatingHistogram) ToByte() ([]byte, error) {
	return json.Marshal(r)
}

func (r ContestRatingHistogram) ToString() (string, error) {
	b, e := r.ToByte()
	return string(b), e
}
