package main

import (
	"fmt"
	leetcode443 "go-leetcode/ArrayString"
)

func main() {

	var myByte = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	//var myByte = []byte{'a'}

	myNumber := leetcode443.Compress(myByte)

	fmt.Println(myNumber)

}
