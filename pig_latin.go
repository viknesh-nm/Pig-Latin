package igpay

import "strings"

// CompareString compares with the input value string and collection strings
func CompareString(value string, collection []string) bool {
	for _, r := range collection {
		if r == value {
			return true
		}
	}
	return false
}

// PigLatin translates english words into the piglatin format
func PigLatin(text string) string {
	words := strings.Split(text, " ")
	var (
		pLatinWord, consField, temp string
	)
	for _, word := range words {
		consField = ""
		if CompareString(string(word[0]), []string{"a", "e", "i", "o", "u", "y", "x"}) {
			if (string(word[0]) == "y" || string(word[0]) == "x") && string(word[1]) == "e" {
				consField = consField + string(word[0])
				suffix := word[len(consField):]
				pLatinWord = suffix + consField + "ay"
				return pLatinWord
			}
			pLatinWord = word + "ay"
			return pLatinWord
		}
		for i, letter := range word {
			if !CompareString(string(letter), []string{"a", "e", "i", "o"}) {
				if string(letter) != "u" {
					consField = consField + string(letter)
				} else if CompareString(string(word[i+1]), []string{"a", "e", "i", "o"}) {
					consField = consField + string(letter)
				} else {
					break
				}
			} else {
				break
			}
		}
		suffix := word[len(consField):]
		temp = suffix + consField + "ay" + " "
		pLatinWord = pLatinWord + temp
	}
	return strings.TrimSpace(pLatinWord)
}
