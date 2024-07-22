package main

import (
	"flag"
	"fmt"
	"os"
)

// cette fonction permets d'afficher les aides essentiels pour le jeu
func help() {
	helpFlag := flag.Bool("h", false, "Affiche les explications")

	flag.Parse()

	// Vérification de l'option -h
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}
}

// fonction qui permets d'afficher les aides
func printHelp() {
	ClearConsole()
	fmt.Println("Bienvenue sur le quiz de pâques. ")
	fmt.Println("\nAfin de jouer au jeu vous devez dans un premier temps écrire dans votre cmd :\n make build puis make run ou .\\QUIZ.exe ")
	fmt.Println("\nEt dans un deuxieme allé sur la bare de recherche de votre navigateur et taper 'localhost:8080' ")
	fmt.Println("\nConditions d'utilisations :")
	fmt.Println("Après avoir faire make build il ne faut pas modifier les fichiers, cela peut poser problème.")
	fmt.Println("N'ajouter pas de phrase ou ne les modifier pas, cela pourrait poser problème")
	fmt.Println("Avant de modifier le programme, faites un 'make clean'")
	fmt.Println("\nObjectif du Jeu")
	fmt.Println("L'objectif du jeu est de répondre aux différentes questions dans les différents niveaux !")
	fmt.Println("\nNiveaux de Difficulté")
	fmt.Println("Facile : Les questions sont très facile.")
	fmt.Println("Moyen : Les questions le sont encore, quelques connaissances sont requises.")
	fmt.Println("Difficile : Attendez-vous à des questions bien plus difficile.")
	fmt.Println("Amusez-vous bien et bonne chance  ! 🎉✨")
}
