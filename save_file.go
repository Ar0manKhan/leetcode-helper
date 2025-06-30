package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func SaveFile(q *Question, generatedCode *genai.GenerateContentResponse) error {
	filename := q.TitleSlug + ".cpp"
	// create this file in current directory
	var sb strings.Builder
	// loop through generatedCode.Candidates[0].Content.Parts and write each part to file
	for _, part := range generatedCode.Candidates[0].Content.Parts {
		sb.WriteString(fmt.Sprintln(part))
	}
	code := TrimCodeMarkdown(sb.String())

	// write the code to the file
	err := os.WriteFile(filename, []byte(code), 0644)
	if err != nil {
		return err
	}
	log.Printf("Code saved to %s", filename)
	return nil
}

func TrimCodeMarkdown(code string) string {
	code = strings.TrimSpace(code)
	code = strings.TrimPrefix(code, "```cpp\n")
	code = strings.TrimSuffix(code, "\n```")
	return code
}
