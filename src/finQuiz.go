package main

import "net/http"

// Réinitialiser opit à 0 lorsque la partie est terminée
func handlerEndQuiz(w http.ResponseWriter, r *http.Request) {
	page = "endQuiz"
	opit = 0 
	data := struct {
		Score  int
		Page   string
		Phrase string
	}{
		Score:  score,
		Page:   page,
		Phrase: phrase,
	}
	tmpl.Execute(w, data)
	handleFormData(r)
}
