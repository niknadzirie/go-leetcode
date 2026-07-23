package ArrayString

import (
	"strings"
)

func ReverseVowels(s string) string {

	//first pointer
	i := 0
	//second pointer
	j := len(s) - 1

	//convert string to rune array for swapping operation
	stringInRune := []rune(s)

	//manual control i and j
	for i < j {

		//both are vowels then we only do swap operation
		if isVowels(stringInRune, i) && isVowels(stringInRune, j) {
			stringInRune[i], stringInRune[j] = stringInRune[j], stringInRune[i]
			j--
			i++
		}

		if !isVowels(stringInRune, i) {
			i++
		}

		if !isVowels(stringInRune, j) {
			j--
		}

	}

	//convert the rune array to string
	result := string(stringInRune)

	//return result in string
	return result

}

func isVowels(s []rune, index int) bool {
	//vowel to check
	vowelsString := "AEIOUaeiou"
	return strings.ContainsRune(vowelsString, s[index])
}
