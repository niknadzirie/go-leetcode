package main

import (
	"fmt"
	leetcode238 "go-leetcode/ArrayString"
)

func main() {

	nums := []int{20, 100, 10, 12, 5, 13}

	myNumber := leetcode238.IncreasingTriplet(nums)

	fmt.Println(myNumber)

}
