package main

import (
	"fmt"
	"strings"
)

func goCensor(str string, censoredWord []string) string {
	var newSlice []string
	var newWord []string
	strSlice := strings.Fields(str)
	for position, word := range strSlice {
		for _, forbiddenWord := range censoredWord {
			if test := strings.Index(strings.ToLower(word), forbiddenWord); test > -1 {
				explode := strings.SplitAfter(word, "")
				for positionx, char := range explode {
					switch char {
					case ".":
						explode[positionx] = "."
					case ",":
						explode[positionx] = ","
					default:
						explode[positionx] = "*"
					}
					newWord = append(explode[:positionx], explode[positionx:]...)
				}
				strSlice[position] = strings.Join(newWord, "")
				newSlice = append(strSlice[:position], strSlice[position:]...)
			}
		}
	}
	return strings.Join(newSlice, " ")
}

func main() {
	var inputStr = `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	var notAllowed = []string{"dolor", "elit", "quis", "nisi", "fugiat", "proident", "laborum"}
	result := goCensor(inputStr, notAllowed)
	fmt.Println(result)
}
