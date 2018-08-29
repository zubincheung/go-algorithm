// Create by Zubin on 2018-08-29 17:36:24
package sort

// 简单的选择排序
func SelectSort(list []int) {
	minIndex := 0
	for i := 0; i < len(list); i++ {
		minIndex = i
		for j := i + 1; j < len(list); j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			swap(list, minIndex, i)
		}
	}
}
