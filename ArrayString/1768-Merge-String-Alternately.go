package leetcode1768

import (
	"strings"
)

func MergeAlternately(word1 string, word2 string) string {
	var stringBuilder strings.Builder
	len1 := len(word1)
	len2 := len(word2)

	maxLength := len1
	if len2 > maxLength {
		maxLength = len2
	}

	for i := 0; i < maxLength; i++ {
		if i < len1 {
			stringBuilder.WriteByte(word1[i])
		}

		if i < len2 {
			stringBuilder.WriteByte(word2[i])
		}
	}

	return stringBuilder.String()
}
