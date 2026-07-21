package ArrayString

func canPlaceFlowers(flowerbed []int, n int) bool {
	nextStopCannotPlant := false
	arrayLength := len(flowerbed)

	for i := 0; i < arrayLength; i++ {

		if n == 0 {
			return true
		}

		if nextStopCannotPlant {
			if flowerbed[i] == 1 {
				n++
				nextStopCannotPlant = true
				continue
			}
			nextStopCannotPlant = false
			continue
		}

		if flowerbed[i] == 1 {
			nextStopCannotPlant = true
			continue
		}

		if flowerbed[i] == 0 {
			n--
			nextStopCannotPlant = true
		}
	}

	if n != 0 {
		return false
	} else {
		return true
	}
}
