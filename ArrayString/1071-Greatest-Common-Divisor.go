package ArrayString

func GCDOfStrings(str1 string, str2 string) string {

	greatest1 := 0
	greatest2 := 0

	lenStr1 := len(str1)
	lenStr2 := len(str2)

	concatenate1 := str1 + str2
	concatenate2 := str2 + str1

	if concatenate1 != concatenate2 {
		return ""
	}

	maxLength := lenStr1

	if lenStr2 > maxLength {
		maxLength = lenStr2
	}

	for i := maxLength; i > 0; i-- {
		if i <= lenStr1 && lenStr1%i == 0 {
			greatest1 = i
		}

		if i <= lenStr2 && lenStr2%i == 0 {
			greatest2 = i
		}

		if greatest1 == greatest2 {
			break
		}
	}

	result := str1[:greatest1]

	return result

}
