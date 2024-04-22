package main

import (
	"ascii_art/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}
	userInput, style := os.Args[1], "standard"
	if len(os.Args) == 3 {
		style = os.Args[2]
	}
	banner, err := os.ReadFile(fmt.Sprintf("../templates/%s.txt", style))
	if err != nil {
		fmt.Println("Error while reading template file")
	}
	bannerSlice := internal.LoadBanner(style, banner)
	internal.PrintAscii(userInput, bannerSlice)
}
