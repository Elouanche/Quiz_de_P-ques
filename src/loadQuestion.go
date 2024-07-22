package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//Permets de charger les questions et savoir si elles ont le bon formats 
func loadQuestions(filename string) []Question {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "9")
		if len(parts) < 3 {
			fmt.Println("Format de question incorrect :", line)
			continue
		}

		questionText := strings.TrimSpace(parts[0])
		options := strings.Split(parts[1], ",")
		correctAns := parts[2]

		question := Question{
			Text:       questionText,
			Options:    options,
			CorrectAns: correctAns,
		}

		questions = append(questions, question)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}
	return questions
}
