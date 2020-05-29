package main

import "fmt"

func main() {
	test := make(map[int]int)
	val, ok := test[1]
	fmt.Println(val, ok)
}
