package piglatin

import (
	"fmt"
	"testing"
)

// TestPigLatin test the functionality of PigLatin
func TestPigLatin(t *testing.T) {
	var (
		english = []string{"test piglatin", "egg", "phone"}
		pLatin  string
	)
	fmt.Println("English: Pig_Latin:")
	for _, value := range english {
		pLatin = PigLatin(value)
		fmt.Println(value, "\t", pLatin)
	}
}
