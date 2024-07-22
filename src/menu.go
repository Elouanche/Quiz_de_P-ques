package main

import "net/http"

// Menu du jeu
func handlerMenu(w http.ResponseWriter, r *http.Request) {
	page = "Menu"
	data := struct {
		Score  int
		Page   string
		Phrase string
	}{
		Score:  score,
		Page:   page,
		Phrase: phrase,
	}

	//Quand la partie est terminé 
	if page == "endQuiz" {
		data.Score = 0 
		tmpl.Execute(w, data)
		handleFormData(r)
		return
	}

	tmpl.Execute(w, data)
	handleFormData(r)
}
