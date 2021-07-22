package main

import (
	"fmt"
	"strings"
)

func CountCharactersFrequency(str string, results chan<- [26]int) {
	var cnt [26]int
	for i := 0; i < 26; i++ {
		cnt[i] += strings.Count(str, string(rune('a'+i)))
	}

	results <- cnt
}
func main() {
	numberOfStrings := 2
	Strings := make([]string, 0)
	Results := make(chan [26]int, numberOfStrings)
	Strings = append(Strings, "aaabbc")
	Strings = append(Strings, "aaabbd")

	for stringCounter := 0; stringCounter < numberOfStrings; stringCounter++ {
		go CountCharactersFrequency(Strings[stringCounter], Results)
	}

	var finalFrequency [26]int

	for _i := 0; _i < numberOfStrings; _i++ {
		result := <-Results
		for i := 0; i < 26; i++ {
			finalFrequency[i] += result[i]
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", rune('a'+i), finalFrequency[i])
	}
}
