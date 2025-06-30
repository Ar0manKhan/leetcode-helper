package main

import (
	"log"

	"github.com/manifoldco/promptui"
)

func main() {
	log.Println("Getting question")
	var q Question
	var err error

	prompt := promptui.Select{
		Label: "Select Question Type",
		Items: []string{"Daily Question", "Specific Question"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	if result == "Daily Question" {
		q, err = GetDailyQuestion()
		if err != nil {
			log.Fatalf("Error on getting question %v", err)
		}
	} else {
		prompt := promptui.Prompt{
			Label: "Enter Question Slug",
		}
		slug, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		q, err = GetQuestionBySlug(slug)
		if err != nil {
			log.Fatalf("Error on getting question %v", err)
		}
	}

	log.Println("Got the question, generating code")
	testCode, err := GenerateTestCase(&q)
	if err != nil {
		log.Fatalf("Error on generating code %v", err)
	}
	log.Println("Code generated, saving the code")
	err = SaveFile(&q, testCode)
	if err != nil {
		log.Fatalf("Error on saving code %v", err)
	}

}
