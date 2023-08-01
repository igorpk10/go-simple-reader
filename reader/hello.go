package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 10
const dataFormat = "2/1/2006 15:04:05"

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
	fmt.Println("Hello, sr. I'm happy to se you here")
}

func showMenu() {
	fmt.Println("0- Exit")
	fmt.Println("1- Start monitor")
	fmt.Println("2- Show Logs")
}

func readCommand() int {
	var comand int
	fmt.Scanf("%d", &comand)

	return comand
}

func starMonitor(sites []string) {
	fmt.Println("Monitoring...")

	for i, site := range sites {
		fmt.Println("Testing", i, ":", site)
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
		fmt.Println("Site:", site, "site has been loaded!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "Houston, we have a problem. Status Code:", resp.StatusCode)
		logRegister(site, false)
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

func logRegister(site string, status bool) {
	file, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println("Register log error")
	}

	file.WriteString(time.Now().Format(dataFormat) + " site: " + site + " " + strconv.FormatBool(status) + "\n")

	file.Close()
}
