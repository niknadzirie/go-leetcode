package ArrayString

import "math"

func IncreasingTriplet(nums []int) bool {

	first := math.Inf(1)
	second := math.Inf(1)

	count := len(nums)

	for i := 0; i < count; i++ {

		currentVal := float64(nums[i])

		if currentVal <= first {
			first = currentVal
		}

		if currentVal > first && currentVal <= second {
			second = currentVal
		}

		if first < second && second < currentVal {
			return true
		}

	}

	return false

}
