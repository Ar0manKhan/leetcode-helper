package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
)

func selectAPIURL() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Enter localhost port",
		Default: "3000",
	}
	port, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://localhost:%s/", port), nil
}

func GetDailyQuestion() (Question, error) {
	// get daily question from leetcode
	// make /dailyQuestion to leetcode api
	leetcodeAPIUrl, err := selectAPIURL()
	if err != nil {
		return Question{}, err
	}
	dailyQuestionUrl := leetcodeAPIUrl + "daily/raw"
	log.Println("You selected", leetcodeAPIUrl)
	// make get api call to daily question url
	req, err := http.NewRequest("GET", dailyQuestionUrl, nil)
	if err != nil {
		return Question{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Question{}, err
	}
	defer resp.Body.Close()
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Question{}, err
	}
	dailyQuestion, err := ParseDailyQuestion(body)
	if err != nil {
		return Question{}, err
	}
	return dailyQuestion.Data.ActiveDailyCodingChallengeQuestion.Question, nil
}

func GetQuestionBySlug(slug string) (Question, error) {
	leetcodeAPIUrl, err := selectAPIURL()
	if err != nil {
		return Question{}, err
	}
	questionUrl := fmt.Sprintf("%sselectQuestion?titleSlug=%s", leetcodeAPIUrl, slug)
	log.Println("You selected", leetcodeAPIUrl)
	// make get api call to daily question url
	req, err := http.NewRequest("GET", questionUrl, nil)
	if err != nil {
		return Question{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Question{}, err
	}
	defer resp.Body.Close()
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Question{}, err
	}
	selectedQuestion, err := ParseSelectedQuestion(body)
	if err != nil {
		return Question{}, err
	}
	return selectedQuestion.Question, nil
}
