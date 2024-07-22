package main

import "net/http"

//Redirige vers le niveau facile du quiz 
func handlerEasy(w http.ResponseWriter, r *http.Request) {
	resetGame()
	page = "Quiz"
	questions = loadQuestions("ressources/quiz.txt")

	if len(questions) > 0 && opit < len(questions) {
		currentQuestion = questions[opit]
		data := struct {
			Score    int
			Page     string
			Phrase   string
			Question Question
		}{
			Score:    score,
			Page:     page,
			Phrase:   phrase,
			Question: currentQuestion,
		}

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Aucune question n'est disponible", http.StatusInternalServerError)
	}

	handleFormData(r)
}
