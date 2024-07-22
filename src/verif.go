package main

import "net/http"

//Permets de vérifier si la réponse envoyé par le formulaire est égal à la bonne option
func handlerVerif(w http.ResponseWriter, r *http.Request) {
	page = "HandlerVerif"

	userAnswer := r.FormValue("answer")

	var result string
	if userAnswer == currentQuestion.CorrectAns {
		score++
		result = "Bonne réponse !"
	} else {
		result = "Mauvaise réponse ! La réponse était :" + currentQuestion.CorrectAns
	}

	data := struct {
		Score  int
		Page   string
		Phrase string
		Result string
	}{
		Score:  score,
		Page:   page,
		Phrase: phrase,
		Result: result,
	}

	tmpl.Execute(w, data)
}
