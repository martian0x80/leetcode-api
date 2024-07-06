package client

import (
	"io"
	"net/http"
)

const graphqlEndpoint string = "https://leetcode.com/graphql"
const method string = "POST"

func newRequest(body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, graphqlEndpoint, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("authority", "leetcode.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("authorization", "")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://leetcode.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://leetcode.com/problems/add-two-numbers/description/")
	req.Header.Add("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Add("x-newrelic-id", "UAQDVFVRGwIAUVhbAAMFXlQ=")

	return req, nil
}

func Request(payload io.Reader) (*http.Response, error) {
	req, err := newRequest(payload)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
