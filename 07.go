package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers []int
	numbers = []int{9, 4, 2, 4, 1, 5, 3, 0}
	sort_odd(numbers)
}

func sort_odd(numbers []int) {
	var odd [8]int
	var pos [8]int
	counter := 0
	for p, number := range numbers {
		if isOdd := number % 2; isOdd != 0 {
			odd[counter] = number
			counter++
			pos[counter] = p
		}
	}
	var newNumbers [8]int
	sort.Ints(odd[:])
	for i, j := range numbers {
		if j%2 != 0 {
			for k, l := range odd {
				if l != 0 {
					newNumbers[i] = l
					odd[k] = 0
					break
				}
			}
		} else {
			newNumbers[i] = j
		}
	}
	fmt.Println(newNumbers)
}
