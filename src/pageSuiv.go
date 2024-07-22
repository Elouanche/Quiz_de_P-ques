package main

import (
	"fmt"
	"net/http"
	"strings"
)

//Redirige vers la question suivant après vérif 
func handlerNextPage(w http.ResponseWriter, r *http.Request) {
	page = "NextPage"
	opit++
	if opit < 5 {
		http.Redirect(w, r, "/Quiz", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/endQuiz", http.StatusSeeOther)
}

func handleFormData(r *http.Request) {
	text := strings.ToLower(r.FormValue("String"))
	fmt.Println(text)
}
