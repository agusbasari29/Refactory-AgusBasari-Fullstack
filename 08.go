package main

import (
	"fmt"
	"strings"
)

func main() {
	secret_text := "23dn3ir30fd2eddd"
	fmt.Println(masking(secret_text))
}

func masking(str string) string {
	newSlice := strings.SplitAfter(str, "")
	masking_len := len(newSlice) - 3
	for position, _ := range newSlice {
		if position < masking_len {
			newSlice[position] = "*"
		}
	}
	return strings.Join(newSlice, "")
}
