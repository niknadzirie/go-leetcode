package main

import (
	"fmt"
	leetcode238 "go-leetcode/ArrayString"
)

func main() {

	nums := []int{1, 2, 3, 4} //[24 , 12 , 8 , 6]

	myNumber := leetcode238.ProductExceptSelf(nums)

	fmt.Println(myNumber)

}
