package view

import (
	"fmt"
	"os"
	"simplereader/reader/logger"
)

func StartMenu() {
	showMenu()
	nextStep()
}

func showMenu() {
	fmt.Println("0- Exit")
	fmt.Println("1- Start monitor")
	fmt.Println("2- Show Logs")
}

func nextStep() {
	comand := readCommand()
	redirect(comand)
}

func readCommand() int {
	var comand int
	fmt.Scanf("%d", &comand)

	return comand
}

func redirect(comand int) {

	switch comand {
	case 0:
		os.Exit(0)
	case 1:
		redirectToWatch()
	case 2:
		redirectToLogs()
	}

}

func redirectToWatch() {
	StartWatch()
}

func redirectToLogs() {
	logger.FetchLogs()
}
