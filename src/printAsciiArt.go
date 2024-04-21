package src

import (
	"fmt"
	"strings"
)

func PrintAsciiArt(result [][8][]byte) {
	var words []string
	for _, word := range result {
		s := ""
		for _, chars := range word {
			s += string(chars)
		}
		words = append(words, s)
	}
	resultString := strings.Join(words, "")
	fmt.Print(resultString)
}
