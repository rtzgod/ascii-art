package internal

import (
	"strings"
)

func LoadBanner(style string, banner []byte) []string {
	if style == "thinkertoy" {
		var correct string
		correct = strings.ReplaceAll(string(banner), string(13), "")
		banner = []byte(correct)
	}
	var bannerSlice []string
	lines := strings.Split(string(banner), "\n\n")
	// Исключаем пустые строки и добавляем строки баннера в срез
	for _, line := range lines {
		if line != "" {
			bannerSlice = append(bannerSlice, line)
		}
	}
	return bannerSlice
}
