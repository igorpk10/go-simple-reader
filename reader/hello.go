package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	showIntroduction()

	for {
		showMenu()
		comand := readCommand()

		switch comand {
		case 0:
			os.Exit(0)
		case 1:
			startLoopMonitor()
		case 2:
			fmt.Println("Showing logs")
		default:
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	name := "Luffy"
	fmt.Println("Hello sr ", name)
}

func showMenu() {
	fmt.Println("0- Exit")
	fmt.Println("1- Start monitor")
	fmt.Println("2- Show Logas")
}

func readCommand() int {
	var comand int
	fmt.Scanf("%d", &comand)

	return comand
}

func starMonitor(sites []string) {
	fmt.Println("Monitoring...")

	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)

		fmt.Println("")
	}
}

func startLoopMonitor() {
	sites := readFile()

	for i := 0; i < monitoramentos; i++ {
		starMonitor(sites)
		fmt.Println("")
	}

	time.Sleep(delay * time.Second)
	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func readFile() []string {
	var sites []string

	file, err := os.Open("siteslist.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		refinedLine := strings.TrimSpace(line)
		sites = append(sites, refinedLine)

		if err == io.EOF {
			break
		}

	}
	file.Close()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Sprintln(sites)

	return sites
}
