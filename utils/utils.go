package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ParseCoord(coord string) (int, int, error) {
	if len(coord) != 2 {
		return 0, 0, fmt.Errorf("invalid coord format. expected something like a1 but was %v", coord)
	}
	x := (int)(coord[0]-'a') + 1
	y, err := strconv.Atoi(coord[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse coordinate: %w", err)
	}
	return x - 1, y - 1, nil
}

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
