package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//permets de clear le terminal
func ClearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	default:
		fmt.Println("Système d'exploitation non pris en charge pour le nettoyage de la console.")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
