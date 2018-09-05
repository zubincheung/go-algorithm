// Create by Zubin on 2018-09-05 19:00:11
package sort

const MAX_LENGTH = 7

// 快速排序,迭代实现
func QuickSort(list []int) {
	qSort(list, 0, len(list)-1)
}

func qSort(list []int, low, hight int) {
	// hight-low大于MAX_LENGTH使用快速排序，否则使用插入排序
	if hight-low > MAX_LENGTH {
		for low < hight {
			pivot := partition(list, low, hight)

			qSort(list, low, pivot-1) // 对低子表递归排序

			low = pivot + 1 // 尾递归
		}
	} else {
		InsertSort(list)
	}
}

// 将list一分为二，并返回枢轴值
func partition(list []int, low, hight int) int {
	// 三数取中
	medianOfThree(list, low, hight)

	pivotKey := list[low] // 使用list中的第一个记录作枢纽记录

	for low < hight {
		//把小于pivotKey的交换到低端
		for low < hight && list[hight] > pivotKey {
			hight--
		}
		list[low] = list[hight]

		//把大于pivotKey的交换到高端
		for low < hight && list[low] < pivotKey {
			low++
		}
		list[hight] = list[low]
	}

	list[low] = pivotKey

	return low
}

// 三数取中算法
func medianOfThree(list []int, low, hight int) {
	mid := (low + hight) / 2

	if list[low] > list[mid] {
		swap(list, low, mid)
	}

	if list[mid] > list[hight] {
		swap(list, mid, hight)
	}

	if list[low] < list[mid] {
		swap(list, low, mid)
	}
}
