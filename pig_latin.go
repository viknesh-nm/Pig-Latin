package piglatin

import (
	"fmt"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

// CompareString compares with the input value string and collection strings
func CompareString(value string, collection []string) bool {
	for _, r := range collection {
		if r == value {
			return true
		}
	}
	return false
}

// PigLatin translates english words into the pig-latin format
func PigLatin(text string) string {
	words := strings.Split(text, " ")
	var (
		pLatinWord, consField, temp string
	)
	for _, word := range words {
		textField := word
		if CompareString(string(textField[0]), vowels) {
			pLatinWord = text + "way"
		} else {
			consField = ""
			for _, letter := range textField {
				if !CompareString(string(letter), vowels) {
					consField = consField + string(letter)
				} else {
					break
				}
			}
			suffix := textField[len(consField):]
			temp = fmt.Sprintf("%s%say", suffix, consField)
		}
		pLatinWord = pLatinWord + temp + " "
	}
	return pLatinWord
}
