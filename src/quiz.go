package main

import "net/http"

func handlerQuiz(w http.ResponseWriter, r *http.Request) {
	page = "Quiz"
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
	handleFormData(r)
}
