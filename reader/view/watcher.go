package view

import (
	"fmt"
	"net/http"
	"simplereader/reader/constants"
	"simplereader/reader/logger"
	listReader "simplereader/reader/readers"
)

func StartWatch() {
	list := listReader.ReadFile(constants.LISTPATH)

	fmt.Println("--- Monitoring ---")
	for _, site := range list {
		startTest(site)
	}
	fmt.Println("--- End Monitor ---")
}

func startTest(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(constants.Houston + "with this request:")
		fmt.Println("--- ", site)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "site has been loaded!")
		logger.LogResults(site, true)
	} else {
		fmt.Println(constants.Houston + " the response code")
		fmt.Sprintln("--- ", site)
		logger.LogResults(site, false)
	}
}
