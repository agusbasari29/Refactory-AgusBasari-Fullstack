package main

import (
	"fmt"
	"strings"
)

func countWords(str string, keyToFind string) uint8 {
	var counter uint8 = 0
	strSlice := strings.Fields(str)
	for _, word := range strSlice {
		if test := strings.Index(strings.ToLower(word), keyToFind); test > -1 {
			counter++
		}
	}
	return uint8(counter)
}

func main() {
	lyric := `
	Aku ingin begini
	Aku ingin begitu
	Ingin ini itu banyak sekali
	
	Semua semua semua
	Dapat dikabulkan
	Dapat dikabulkan
	Dengan kantong ajaib
	
	Aku ingin terbang bebas
	Di angkasa
	Hei… baling baling bambu
	
	La... la... la...
	Aku sayang sekali
	Doraemon
	
	La... la... la...
	Aku sayang sekali
	`

	var findWords = [3]string{
		`aku`,
		`ingin`,
		`dapat`,
	}

	for _, findWord := range findWords {
		fmt.Println(findWord, " = ", countWords(lyric, findWord))
	}
}
