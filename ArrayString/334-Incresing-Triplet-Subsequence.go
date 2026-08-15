package ArrayString

func IncreasingTriplet(nums []int) bool {

	ln := len(nums)
	minNum := 0
	maxNum := 0

	minIndex := 0
	midIndex := 0
	maxIndex := 0

	for i := 0; i < ln; i++ {
		j := ln - 1 - i

		if i == 0 {
			minNum, minIndex = nums[i], i
		}

		if j == ln-1 {
			maxNum, maxIndex = nums[j], j
		}

		if nums[i] < minNum {
			minNum, minIndex = nums[i], i
		}

		if nums[j] > maxNum {
			maxNum, maxIndex = nums[j], j
		}

	}

	for i := 0; i < ln; i++ {
		if minNum < nums[i] && nums[i] < maxNum {
			midIndex = i
		}

		if minIndex < midIndex && midIndex < maxIndex {
			return true
		}
	}

	return false

}
