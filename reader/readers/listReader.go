package listReader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"simplereader/reader/constants"
	"strings"
)

func ReadFile(relativepath string) []string {
	var sites []string

	file, err := os.Open(filepath.Join(constants.BASEPATH, relativepath))

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
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

	return sites
}
