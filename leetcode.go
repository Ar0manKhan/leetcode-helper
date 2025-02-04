package main

import (
	"io"
	"net/http"
)

// const LeetcodeAPI = "http://localhost:3000/" // for local testing
const LeetcodeAPI = "https://alfa-leetcode-api.onrender.com/"

func GetDailyQuestion() (DailyQuestionResponse, error) {
	// get daily question from leetcode
	// make /dailyQuestion to leetcode api
	dailyQuestionUrl := LeetcodeAPI + "dailyQuestion"
	// make get api call to daily question url
	req, err := http.NewRequest("GET", dailyQuestionUrl, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return ParseDailyQuestion(body)
}
