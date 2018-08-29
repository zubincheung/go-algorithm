package main

import (
	"fmt"
	"zubin/go-algorithm/algorithms/sort"
)

func main() {
	list := []int{6, 3, 1, 8, 5, 9}

	sort.InsertSort(list)

	fmt.Println("------------\n", list)
}
