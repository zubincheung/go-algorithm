// Create by Zubin on 2018-06-14 14:51:30

package main

import (
	"zubin/go-algorithm/algorithms/sort"

	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recover", r)
		debug.PrintStack()
	}
}

func main() {
	defer r()

	list := []int{2, 1, 3, 4, 5, 6, 7, 8, 9}
	sort.FlagSwapPositionBubbleSort(list)
	fmt.Println("-------------\n", list)
}
