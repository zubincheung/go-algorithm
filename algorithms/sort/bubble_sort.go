package sort

func SimpleBubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				swap(list, i, j)
			}
		}
	}

}

func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list); j++ {
			if list[j-1] > list[j] {
				swap(list, j-1, j)
			}
		}
	}
}

func FlagSwapBubbleSort(list []int) {
	hasSwaped := true

	for i := 0; i < len(list) && hasSwaped; i++ {
		hasSwaped = false
		for j := 1; j < len(list); j++ {
			if list[j-1] > list[j] {
				swap(list, j-1, j)
				hasSwaped = true
			}
		}
	}
}

func FlagSwapPositionBubbleSort(list []int) {
	var flag int
	laseSwapPosition := len(list)
	hasSwaped := true

	for i := 0; i < len(list) && hasSwaped; i++ {
		hasSwaped = false
		flag = laseSwapPosition
		for j := 1; j < flag; j++ {
			if list[j-1] > list[j] {
				swap(list, j-1, j)
				hasSwaped = true
				laseSwapPosition = j
			}
		}
	}
}
