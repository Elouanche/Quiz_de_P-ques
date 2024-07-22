package main

import (
	"fmt"
	"log"
	"net/http"
)

//Permets à la redirection vers les différentes pages 
func route(h1, h2, h3, h4, h5, h6, h7, h8, h9 http.HandlerFunc) {
	http.HandleFunc("/", h1)
	http.HandleFunc("/Quiz", h2)
	http.HandleFunc("/Menu", h3)
	http.HandleFunc("/Easy", h4)
	http.HandleFunc("/Medium", h5)
	http.HandleFunc("/Hard", h6)
	http.HandleFunc("/endQuiz", h7)
	http.HandleFunc("/HandlerVerif", h8)
	http.HandleFunc("/NextPage", h9)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	ClearConsole()
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
