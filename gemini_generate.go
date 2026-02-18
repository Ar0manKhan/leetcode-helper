package main

import (
	"context"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/manifoldco/promptui"
	"google.golang.org/api/option"
)

func GenerateTestCase(q *Question) (*genai.GenerateContentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		return nil, fmt.Errorf("GEMINI_API_KEY not set")
	}

	// context with timeout of 1 minute
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	aiModels := []string{"gemini-3-flash-preview", "gemini-2.5-flash", "gemini-2.5-pro", "gemini-3-pro-preview"}
	prompt := promptui.Select{
		Label: "Select AI Model",
		Items: aiModels,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	model := client.GenerativeModel(result)
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.Text(GenerateQuestionString(q)),
			},
		},
	}

	return session.SendMessage(ctx, genai.Text("INSERT_INPUT_HERE"))
}

func GenerateQuestionString(q *Question) string {
	snippets := q.CodeSnippets
	snippetIdx := slices.IndexFunc(snippets, func(e CodeSnippet) bool {
		return e.LangSlug == "cpp"
	})
	if snippetIdx == -1 {
		return ""
	}

	question := fmt.Sprintf(
		`Write the setup and test case for this leetcode question.
Write the code in %s, use fmt for logging.
The solution class is "%s".

Do not solve the problem, just add the solution class as given here.
Only include the test case given in the question.
Include all the examples mentioned in the question.
Solution will be implemented by the user.
Test case should be in main function and follow this structure:
	Make array of question structures (in array or vector) which include all inputs and expected output.
	Loop through the array and call the function to be tested.
	Display the result, inputs and test case passed or failed.

%s: %s
%s`,
		snippets[snippetIdx].Lang,
		snippets[snippetIdx].Code,
		q.QuestionFrontendID,
		q.Title,
		q.Content,
	)
	return question
}
