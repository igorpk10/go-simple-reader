package main

import (
	"fmt"
	"simplereader/reader/view"
)

func main() {
	showIntroduction()
	view.StartMenu()
}

func showIntroduction() {
	fmt.Println("Hello, sr. I'm happy to se you here")
}
