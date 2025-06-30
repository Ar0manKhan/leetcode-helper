package main

import "encoding/json"

type CodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

type Question struct {
	QuestionID         string        `json:"questionId"`
	QuestionFrontendID string        `json:"questionFrontendId"`
	BoundTopicID       interface{}   `json:"boundTopicId"`
	Title              string        `json:"title"`
	TitleSlug          string        `json:"titleSlug"`
	Content            string        `json:"content"`
	TranslatedTitle    interface{}   `json:"translatedTitle"`
	TranslatedContent  interface{}   `json:"translatedContent"`
	IsPaidOnly         bool          `json:"isPaidOnly"`
	Difficulty         string        `json:"difficulty"`
	Likes              int           `json:"likes"`
	Dislikes           int           `json:"dislikes"`
	IsLiked            interface{}   `json:"isLiked"`
	SimilarQuestions   string        `json:"similarQuestions"`
	ExampleTestcases   string        `json:"exampleTestcases"`
	Contributors       []interface{} `json:"contributors"`
	TopicTags          []struct {
		Name           string      `json:"name"`
		Slug           string      `json:"slug"`
		TranslatedName interface{} `json:"translatedName"`
	} `json:"topicTags"`
	CompanyTagStats interface{}   `json:"companyTagStats"`
	CodeSnippets    []CodeSnippet `json:"codeSnippets"`
	Stats           string        `json:"stats"`
	Hints           []string      `json:"hints"`
	Solution        struct {
		ID               string `json:"id"`
		CanSeeDetail     bool   `json:"canSeeDetail"`
		PaidOnly         bool   `json:"paidOnly"`
		HasVideoSolution bool   `json:"hasVideoSolution"`
		PaidOnlyVideo    bool   `json:"paidOnlyVideo"`
	} `json:"solution"`
	Status            interface{}   `json:"status"`
	SampleTestCase    string        `json:"sampleTestCase"`
	MetaData          string        `json:"metaData"`
	JudgerAvailable   bool          `json:"judgerAvailable"`
	JudgeType         string        `json:"judgeType"`
	MysqlSchemas      []interface{} `json:"mysqlSchemas"`
	EnableRunCode     bool          `json:"enableRunCode"`
	EnableTestMode    bool          `json:"enableTestMode"`
	EnableDebugger    bool          `json:"enableDebugger"`
	EnvInfo           string        `json:"envInfo"`
	LibraryURL        interface{}   `json:"libraryUrl"`
	AdminURL          interface{}   `json:"adminUrl"`
	ChallengeQuestion struct {
		ID                       string `json:"id"`
		Date                     string `json:"date"`
		IncompleteChallengeCount int    `json:"incompleteChallengeCount"`
		StreakCount              int    `json:"streakCount"`
		Type                     string `json:"type"`
	} `json:"challengeQuestion"`
	Note interface{} `json:"note"`
}

type DailyQuestionResponse struct {
	Data struct {
		ActiveDailyCodingChallengeQuestion struct {
			Date     string   `json:"date"`
			Link     string   `json:"link"`
			Question Question `json:"question"`
		} `json:"activeDailyCodingChallengeQuestion"`
	} `json:"data"`
}

type SelectQuestionResponse struct {
	Question Question `json:"question"`
}

func ParseDailyQuestion(body []byte) (DailyQuestionResponse, error) {
	var response DailyQuestionResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func ParseSelectedQuestion(body []byte) (SelectQuestionResponse, error) {
	var response SelectQuestionResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

