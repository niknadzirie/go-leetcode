package ArrayString

import (
	"strings"
)

func ReverseWords(s string) string {

	wordsSlice := strings.Fields(s)

	sliceLength := len(wordsSlice)

	var builder strings.Builder

	for i := sliceLength - 1; i >= 0; i-- {
		builder.WriteString(wordsSlice[i])
		if i != 0 {
			builder.WriteString(" ")
		}

	}

	return builder.String()

}
