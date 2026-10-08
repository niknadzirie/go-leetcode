package Stack

func removeStars(s string) string {
	var strStack []byte
	deleteAt := -1

	sLength := len(s)

	for i := 0; i < sLength; i++ {
		if s[i] != '*' {
			strStack = append(strStack, s[i])
			deleteAt++
			continue
		}

		strStack = strStack[:len(strStack)-1]
	}

	return string(strStack)
}
