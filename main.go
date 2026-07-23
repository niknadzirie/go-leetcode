package main

import (
	"fmt"
	leetcode238 "go-leetcode/ArrayString"
)

func main() {

	nums := []int{-1, 1, 0, -3, 3}

	myNumber := leetcode238.ProductExceptSelf(nums)

	fmt.Println(myNumber)

}
