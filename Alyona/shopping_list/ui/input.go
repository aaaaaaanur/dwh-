package ui

import (
	"bufio"
	"os"
	"strconv"
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

func ReadInt() (int, error) {
	str, err := ReadString()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}

func ReadUint() (uint, error) {
	str, err := ReadString()
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
