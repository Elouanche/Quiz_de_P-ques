package main

import (
	"fmt"
	"text/template"
)

var (
	tmpl            *template.Template
	score           int
	page            = " "
	phrase          = ""
	opit            = 0 
	currentQuestion Question 
	questions       []Question
)

//Structure pour les questions 
type Question struct {
	Text       string   // La question
	Options    []string // Les options de réponse
	CorrectAns string   // Var reponse correct
}

//Permets l'initialisation du jeu et des redirections
func main() {
	help()
	setup()
	route(
		handlerAccueil,
		handlerQuiz,
		handlerMenu,
		handlerEasy,
		handlerMedium,
		handlerHard,
		handlerEndQuiz,
		handlerVerif,
		handlerNextPage,
	)
}

func setup() {
	var err error
	tmpl, err = template.ParseFiles("./static/html/index.html")
	if err != nil {
		fmt.Println(err)
	}
}
