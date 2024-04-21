package src

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func AtoAsciiArt(userInput string) [][8][]byte {
	banner, err := os.ReadFile("templates/standard.txt")
	if err != nil {
		fmt.Println("Error while reading template file")
	}
	standardSymbols := makeBannerSymbols(banner)
	var inputInAscii [][][]byte
	splittedInput := strings.Split(userInput, "\\n")
	for _, input := range splittedInput {
		var temp [][]byte
		for _, c := range input {
			if !(' ' <= c && c <= '~') {
				fmt.Println("Incorrect char in input string")
				return nil
			}
			temp = append(temp, standardSymbols[c])
		}
		inputInAscii = append(inputInAscii, temp)
	}
	var result [][8][]byte

	for _, word := range inputInAscii {
		var temp [8][]byte
		for _, chars := range word {

			start := 1
			iter := 0
			for i := 1; i < len(chars); i++ {
				if chars[i] == 10 {
					temp[iter] = slices.Concat(temp[iter], chars[start:i])
					start = i + 1
					iter++
				}
			}
		}
		for i := 0; i < len(temp); i++ {
			if string(temp[i]) == "" {
				continue
			}
			temp[i] = append(temp[i], 10)
		}
		result = append(result, temp)
	}
	return result
}

func makeBannerSymbols(banner []byte) map[rune][]byte {
	Symbols := make(map[rune][]byte)
	curr := ' '
	count := 0
	start := 0
	for i := 1; i < len(banner); i++ {
		if i+1 == len(banner) {
			curr++
			Symbols[curr] = banner[start : i+1]
		}
		if banner[i] == 10 {
			count++
		}
		if count == 9 {
			Symbols[curr] = banner[start:i]
			curr++
			count = 0
			start = i
		}
	}
	return Symbols
}
