package main

import "fmt"

type vowelsFinder interface {
	findVowels() []rune
}

type myString string

func (ms myString) findVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	var name myString = "Anna Xaviour"
	fmt.Printf("vowels are %c", name.findVowels())
}
