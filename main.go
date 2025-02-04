package main

import "log"

func main() {
	log.Println("Getting question")
	q, err := GetDailyQuestion()
	if err != nil {
		log.Fatalf("Error on getting question %v", err)
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
