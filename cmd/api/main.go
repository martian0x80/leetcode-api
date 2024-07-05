package main

import (
	"fmt"
	"lc-api/internal/leetcode"
	"lc-api/pkg/leetcode/queries"
	"lc-api/pkg/leetcode/responses"
)

func main() {
	//profile, e := leetcode.GetUserPublicProfile("randomdude")
	profile, e := leetcode.Get[responses.UserPublicProfile](&queries.UserPublicProfile{Username: "randomdude"})
	if e != nil {
		panic(e)
	}

	fmt.Println(profile.ToString())
}
