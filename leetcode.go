package main

import (
	"io"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
)

func GetDailyQuestion() (DailyQuestionResponse, error) {
	// get daily question from leetcode
	// make /dailyQuestion to leetcode api
	availableUrls := []string{
		"https://alfa-leetcode-api.onrender.com/",
		"http://localhost:3000/",
	}
	log.Println("Asking for prompt")
	prompt := promptui.Select{
		Label: "Select API URL",
		Items: availableUrls,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return DailyQuestionResponse{}, err
	}
	leetcodeAPIUrl := result
	dailyQuestionUrl := leetcodeAPIUrl + "dailyQuestion"
	log.Println("You selected", leetcodeAPIUrl)
	// make get api call to daily question url
	req, err := http.NewRequest("GET", dailyQuestionUrl, nil)
	if err != nil {
		return DailyQuestionResponse{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return DailyQuestionResponse{}, err
	}
	defer resp.Body.Close()
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DailyQuestionResponse{}, err
	}
	return ParseDailyQuestion(body)
}
