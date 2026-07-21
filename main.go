package main

import (
	"fmt"
	leetcode1431 "go-leetcode/ArrayString"
)

func main() {

	candies := []int{2, 3, 5, 1, 3}

	result := leetcode1431.KidsWithCandies(candies, 3)
	fmt.Println(result)

}
