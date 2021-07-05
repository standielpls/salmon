package io

import (
	"bufio"
	"io"
	"strings"
)

func ReadFrom(r io.Reader) string {
	reader := bufio.NewReader(r)
	input, _ := reader.ReadString('\n')
	return strings.Trim(input, " \n")
}
