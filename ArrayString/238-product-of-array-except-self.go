package ArrayString

func ProductExceptSelf(nums []int) []int {

	var result []int
	arrayLength := len(nums)
	maxIndex := arrayLength - 1
	j := maxIndex
	i := 0
	storeHere := 1

	for i < arrayLength {

		if j == 0 {
			//append into result
			result = append(result, storeHere)

			//reset
			i++
			j = maxIndex
			storeHere = 1
			continue
		}

		//skip self multiplication
		if j != i {
			storeHere *= nums[j]
		}
		j--

	}

	return result
}
