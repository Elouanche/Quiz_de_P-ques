package main

import "net/http"

//Permets la redirection vers la page d'accueil du site
func handlerAccueil(w http.ResponseWriter, r *http.Request) {
	page = "Accueil"
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
