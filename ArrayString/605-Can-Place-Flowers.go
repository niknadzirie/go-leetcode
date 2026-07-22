package ArrayString

func CanPlaceFlowers(flowerbed []int, n int) bool {

	arrayLength := len(flowerbed)

	for i := 0; i < arrayLength; i++ {

		//no plant to be planted
		if n == 0 {
			break
		}

		//skip if already have plant
		if flowerbed[i] == 1 {
			continue
		}

		// checking for array size = 1
		if arrayLength < 2 && flowerbed[i] == 0 {
			n--
			continue
		}

		//checking at starting of the array
		if i == 0 {
			//only check right side
			if flowerbed[i+1] == 1 {
				continue
			}
			flowerbed[i] = 1
			n--
			continue
		}

		//checking at the end of the array
		if i == arrayLength-1 {
			//only check left side
			if flowerbed[i-1] == 1 {
				continue
			}
			n--
			continue
		}

		//both side empty, can plant here
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}

	if n == 0 {
		return true
	} else {
		return false
	}

}
