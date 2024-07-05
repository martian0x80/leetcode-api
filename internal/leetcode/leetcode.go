package leetcode

import (
	"encoding/json"
	"fmt"
	"io"
	Query "lc-api/pkg/leetcode/queries"
	"lc-api/pkg/leetcode/responses"
)

//func GetUserPublicProfile(username string) (*responses.UserPublicProfile, error) {
//	q := Query.UserPublicProfile{Username: username}
//	query := q.GetQuery()
//	resp, err := Request(query)
//	if err != nil {
//		return nil, err
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(resp.Body)
//	profile := new(responses.UserPublicProfile)
//	err = json.NewDecoder(resp.Body).Decode(profile)
//	if err != nil {
//		return nil, err
//	}
//	return profile, nil
//}

func Get[T responses.IResponse](query Query.IQuery) (*T, error) {
	q := query.GetQuery()
	resp, err := Request(q)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	data := new(T)
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
