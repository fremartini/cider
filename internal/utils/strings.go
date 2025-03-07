package utils

import "strings"

func PadLeft(s string, paddingChar rune, totalWidth int) string {
	if len(s) >= totalWidth {
		return s
	}

	padding := totalWidth - len(s)

	return strings.Repeat(string(paddingChar), padding) + s
}

func PadRight(s string, paddingChar rune, totalWidth int) string {
	if len(s) >= totalWidth {
		return s
	}

	padding := totalWidth - len(s)

	return s + strings.Repeat(string(paddingChar), padding)
}
