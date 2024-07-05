package queries

import (
	"io"
	"strings"
)

type ContestRatingHistogram struct {
}

func (u *ContestRatingHistogram) GetQuery() io.Reader {
	return strings.NewReader("{\"query\":\"query contestRatingHistogram {\\n  contestRatingHistogram {\\n    userCount\\n    ratingStart\\n    ratingEnd\\n    topPercentage\\n  }\\n}\",\"variables\":{}}")
}
