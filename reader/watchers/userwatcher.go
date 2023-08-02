package watcher

import (
	"fmt"
	"simplereader/reader/constants"
	listReader "simplereader/reader/readers"
	"time"
)

func StartUserWatcher() {
	captureEntries()
}

func captureEntries() {
	var times int
	var delay int
	fmt.Println("--- Now, we need whay entries you need ---")

	for {
		fmt.Println("Times ", times)
		fmt.Println("Delay ", delay)

		if times <= 0 {
			fmt.Print("Please, enter how many calls you need in your monitor?: ")
			times = captureEntry()
			fmt.Println(times)
		}

		if delay <= 0 {
			fmt.Print("Please, enter the delay you need: ")
			delay = captureEntry()
			fmt.Println(delay)
		}

		if (times > 0) && (delay > 0) {
			fmt.Println("Times: ", times)
			fmt.Println("Space: ", delay)
			break
		}
	}

	startMonitor(times, delay)
}

func captureEntry() int {
	var command int
	fmt.Scanf("%d", &command)

	return command
}

func startMonitor(times int, delay int) {
	list := listReader.ReadFile(constants.LISTPATH)

	fmt.Println("--- Start monitoring ---")

	for i := 1; i <= times; i++ {

		fmt.Println("--- Monitor ", i, " ---")
		for _, site := range list {
			StartTest(site)
		}
		fmt.Println("--- End Monitor ", i, " ---")
		fmt.Sprintln("")

		time.Sleep(time.Duration(time.Duration(delay).Seconds()))
	}

}
