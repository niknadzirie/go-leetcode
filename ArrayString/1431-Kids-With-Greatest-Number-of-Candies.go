package ArrayString

import (
	"slices"
)

func KidsWithCandies(candies []int, extraCandies int) []bool {

	maxValue := slices.Max(candies)
	lengthCandies := len(candies)

	result := []bool{}

	for i := 0; i < lengthCandies; i++ {

		addedCandies := candies[i] + extraCandies

		if addedCandies >= maxValue {
			result = append(result, true)
		} else {
			result = append(result, false)
		}

	}

	return result

}
