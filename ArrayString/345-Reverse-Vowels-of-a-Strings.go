package ArrayString

import (
	"strings"
)

func ReverseVowels(s string) string {

	//second pointer
	j := len(s) - 1

	//length of the string
	stringLength := len(s)

	//convert string to byte array, becauseneed to swap the character in string
	stringInRune := []rune(s)

	//iterate each char/index in  string
	for i := 0; i < stringLength; i++ {

		//finish when i and j clash
		if i >= j {
			break
		}

		//j is shifting if i is vowels
		if isVowels(stringInRune, i) && !isVowels(stringInRune, j) {

			for j > 0 && !isVowels(stringInRune, j) {
				j--
			}

		}

		//both are vowels we only do swap operation
		if isVowels(stringInRune, i) && isVowels(stringInRune, j) {
			stringInRune[i], stringInRune[j] = stringInRune[j], stringInRune[i]
			if j > 0 {
				j--
			}
		}

	}

	//convert the byte array to string
	result := string(stringInRune)

	//return result in string
	return result

}

func isVowels(s []rune, index int) bool {
	//vowel to check
	vowelsString := "AEIOUaeiou"
	return strings.ContainsRune(vowelsString, s[index])
}
