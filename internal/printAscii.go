package internal

import (
	"fmt"
	"strings"
)

const MaxBannerLines = 8

func PrintAscii(userInput string, bannerSlice []string) {
	userInput = strings.ReplaceAll(userInput, "\\n", "\n")
	linesCount := strings.Count(userInput, "\n")
	if len(userInput) == linesCount {
		for i := 0; i < linesCount; i++ {
			fmt.Println()
		}
		return
	}
	var lines []string
	for _, char := range userInput {
		if char == '\n' {
			if len(lines) != 0 {
				for _, line := range lines {
					fmt.Println(line)
				}
				lines = []string{}
			} else {
				fmt.Println()
			}
			continue
		}

		charIndex := int(char) - 32

		if charIndex >= 0 && charIndex < len(bannerSlice) {

			asciiChar := bannerSlice[charIndex]
			asciiCharLines := strings.Split(asciiChar, "\n")

			for i := 0; i < MaxBannerLines && i < len(asciiCharLines); i++ {
				if i < len(lines) {
					lines[i] += asciiCharLines[i]
				} else {
					lines = append(lines, asciiCharLines[i])
				}
			}
		}
	}
	if len(lines) != 0 {
		for _, line := range lines {
			fmt.Println(line)
		}
	} else {
		fmt.Println()
	}
}
