package main

import (
	"fmt"
	"strings"
)

func CountCharactersFrequency(str string, results chan<- [128]int) {
	var cnt [128]int
	for i := 0; i < 128; i++ {
		cnt[i] += strings.Count(str, string(rune(' '+i)))
	}

	results <- cnt
}
func main() {
	numberOfStrings := 2
	strings := []string{"aaabbc", "aaabbd"}
	results := make(chan [128]int, numberOfStrings)

	for stringCounter := 0; stringCounter < numberOfStrings; stringCounter++ {
		go CountCharactersFrequency(strings[stringCounter], results)
	}

	var finalFrequency [128]int

	for _i := 0; _i < numberOfStrings; _i++ {
		result := <-results
		for i := 0; i < 128; i++ {
			finalFrequency[i] += result[i]
		}
	}

	for i := 0; i < 128; i++ {
		fmt.Printf("%c : %d\n", rune(' '+i), finalFrequency[i])
	}
}
