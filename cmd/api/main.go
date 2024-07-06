package main

import (
	"fmt"
	"github.com/martian0x80/leetcode-api/pkg/leetcode"
	"github.com/martian0x80/leetcode-api/pkg/leetcode/queries"
	"github.com/martian0x80/leetcode-api/pkg/leetcode/responses"
)

func main() {
	
	// Example of using the leetcode API

	// Get user profile

	profile, e := leetcode.Get[responses.UserPublicProfile](&queries.UserPublicProfile{Username: "randomdude"})
	if e != nil {
		panic(e)
	}

	fmt.Println(profile.ToString())
}
