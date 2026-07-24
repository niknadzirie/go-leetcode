package ArrayString

func ProductExceptSelf(nums []int) []int {

	arrayLength := len(nums)
	result := make([]int, arrayLength)
	LS := make([]int, arrayLength)
	RS := make([]int, arrayLength)
	j := arrayLength - 1

	leftSide := 1
	rigthSide := 1

	//getting the LS and RS multiplication
	for i := 0; i < arrayLength; i++ {

		if i == 0 {
			leftSide *= 1
			LS[i] = leftSide
		} else {
			leftSide *= nums[i-1]
			LS[i] = leftSide
		}

		if j == arrayLength-1 {
			rigthSide *= 1
			RS[j] = rigthSide
			j--
			continue
		}

		rigthSide *= nums[j+1]
		RS[j] = rigthSide
		j--
	}

	//insert the multiplication except self into result
	for i := 0; i < arrayLength; i++ {
		result[i] = LS[i] * RS[i]
	}

	return result
}
