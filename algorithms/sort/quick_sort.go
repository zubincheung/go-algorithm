// Create by Zubin on 2018-09-05 19:00:11
package sort

// 快速排序,迭代实现
func QuickSort(list []int) {
	qSort(list, 0, len(list)-1)
}

func qSort(list []int, low, hight int) {
	if low < hight {
		pivot := partition(list, low, hight)

		qSort(list, low, pivot-1)   // 对低子表递归排序
		qSort(list, pivot+1, hight) // 对高子表递归排序
	}
}

// 将list一分为二，并返回枢轴值
func partition(list []int, low, hight int) int {
	pivotKey := list[low]

	for low < hight {
		//把小于pivotKey的交换到低端
		for low < hight && list[hight] > pivotKey {
			hight--
		}
		swap(list, low, hight)
		//把大于pivotKey的交换到高端
		for low < hight && list[low] < pivotKey {
			low++
		}
		swap(list, low, hight)
	}

	return low
}
