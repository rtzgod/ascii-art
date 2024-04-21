package main

import (
	"ascii_art/src"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect amount of arguments")
		return
	}
	inputString := os.Args[1]
	result := src.AtoAsciiArt(inputString)
	src.PrintAsciiArt(result)
}
