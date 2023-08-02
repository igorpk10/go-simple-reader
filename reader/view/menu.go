package view

import (
	"fmt"
	"os"
	"simplereader/reader/logger"
	watcher "simplereader/reader/watchers"
)

func StartMenu() {
	showMenu()
	nextStep()
}

func showMenu() {
	fmt.Println("0- Exit")
	fmt.Println("1- Start monitor")
	fmt.Println("2- Start buildable monitor")
	fmt.Println("3- Show Logs")
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
		redirectToUserWatch()
	case 3:
		redirectToLogs()
	default:
		os.Exit(0)
	}

}

func redirectToWatch() {
	watcher.StartWatch()
}

func redirectToUserWatch() {
	watcher.StartUserWatcher()
}

func redirectToLogs() {
	logger.FetchLogs()
}
