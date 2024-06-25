package leetcode

import (
	"encoding/json"
	"fmt"
	"io"
	Query "lc-api/internal/leetcode/queries"
	"lc-api/internal/leetcode/responses"
)

func GetUserPublicProfile(username string) (*responses.UserPublicProfile, error) {
	q := Query.UserPublicProfile{Username: username}
	query := q.GetQuery()
	resp, err := Request(query)
	if err != nil {
		return nil, err
	}
	profile := new(responses.UserPublicProfile)
	err = json.NewDecoder(resp.Body).Decode(profile)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	return profile, nil
}
