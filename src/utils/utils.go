package utils

import (
	"os"
	"os/exec"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("clear") // for linux
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func WrapLines(input string, n int) string {
	words := strings.Fields(input)
	var lines []string
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word) <= n {
			currentLine += word + " "
		} else {
			lines = append(lines, strings.TrimSpace(currentLine))
			currentLine = word + " "
		}
	}

	lines = append(lines, strings.TrimSpace(currentLine))
	return strings.Join(lines, "\n")
}
