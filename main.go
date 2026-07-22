package main

import (
	"fmt"
	leetcode605 "go-leetcode/ArrayString"
)

func main() {

	test := []int{0, 0, 1, 0, 0}

	result := leetcode605.CanPlaceFlowers(test, 1)
	fmt.Println(result)

}
