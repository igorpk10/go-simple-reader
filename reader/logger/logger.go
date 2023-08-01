package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"simplereader/reader/constants"
	listReader "simplereader/reader/readers"
	"time"
)

func LogResults(site string, status bool) {

	filepath := filepath.Join(constants.BASEPATH, constants.LOGPATH)
	file, error := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println("Register log error")
	}

	if status {
		file.WriteString(time.Now().Format(constants.DataFormat) + " site: " + site + " - active" + "\n")
	} else {
		file.WriteString(time.Now().Format(constants.DataFormat) + " site: " + site + " - inactive" + "\n")
	}

	file.Close()
}

func FetchLogs() {

	logs := listReader.ReadFile(constants.LOGPATH)

	fmt.Println("--- Start Logs --- \n")

	for _, text := range logs {
		fmt.Println(text)
	}
	fmt.Println("--- End Logs ---")

}
