# Leetcode API for Go

This repository contains a unofficial wrapper for the undocumented Leetcode GraphQL API. It is written in Go and provides a simple interface to interact with the Leetcode API.

This is also my first Go project, so any feedback is appreciated.

## Disclaimer
This project is not affiliated with Leetcode. It is intended for educational purposes only.

## Installation


```bash
go get github.com/martian0x80/leetcode-api
```

## Usage

Example in `main.go`:

```go
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

```
