package main

import (
	"fmt"
)

func main() {
	// test := make(map[int]int)
	// val, ok := test[1]
	grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	result := numIslands(grid)
	fmt.Println(result)
}
