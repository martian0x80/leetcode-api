package main

import (
	"fmt"
	"lc-api/internal/leetcode"
)

func main() {
	profile, e := leetcode.GetUserPublicProfile("randomdude")
	if e != nil {
		panic(e)
	}

	fmt.Println(profile.ToString())
}
