package igpay

import (
	"log"
	"regexp"
	"strings"
)

var (
	vowels = regexp.MustCompile("^[aeiou]|^[yx][^aeiou]")
	cons   = regexp.MustCompile(".?(ch|qu|th)[^aeiou]?|.")
)

// PigLatin translates a string into Pig Latin
func PigLatin(text string) string {
	words := strings.Split(text, " ")
	pigLatin := []string{}
	for _, word := range words {
		pigLatin = append(pigLatin, pigLatinWord(word))
		log.Println(pigLatin)
	}
	return strings.Join(pigLatin, " ")
}

func pigLatinWord(word string) string {
	if vowels.MatchString(word) {
		return word + "ay"
	}
	wordCount := cons.FindStringIndex(word)
	return word[wordCount[1]:] + word[wordCount[0]:wordCount[1]] + "ay"
}
