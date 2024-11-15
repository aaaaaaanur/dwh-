package ui

import (
	"bufio"
	"os"
	"strings"
)

func ReadString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(str, "\n"), nil
}
