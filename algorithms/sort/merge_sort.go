// Create by Zubin on 2018-09-04 20:04:14

package sort

// 归并排序递归实现
func MergeSort(list []int) {
	var tmp []int = make([]int, len(list)) // 初始化一个和list等长的slice
	mergeGroup(list, tmp, 0, len(list)-1)
}

func mergeGroup(list, tmp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2

		mergeGroup(list, tmp, left, mid)    //左边归并排序，使左子序列有序
		mergeGroup(list, tmp, mid+1, right) //右边归并排序，使右子序列有序

		merge(list, tmp, left, mid, right) // 合并两个有序子序列
	}
}

// 将有二个有序数列list[left...mid]和list[mid...right]合并
func merge(list, tmp []int, left, mid, rigth int) {
	i := left    // 左序列指针
	j := mid + 1 //右序列指针
	k := left    // 临时列表指针

	// 将list中记录由小到大地并入tmp
	for i <= mid && j <= rigth {
		if list[i] < list[j] {
			tmp[k] = list[i]
			i++
		} else {
			tmp[k] = list[j]
			j++
		}
		k++
	}

	// 将左边剩余元素填充到tmp中
	for i <= mid {
		tmp[k] = list[i]
		k++
		i++
	}
	// 将右边剩余元素填充到tmp中
	for j <= rigth {
		tmp[k] = list[j]
		k++
		j++
	}

	// 将tmp中的元素拷贝到list
	for left <= rigth {
		list[left] = tmp[left]
		left++
	}

}

// 归并排序迭代实现
func MergeSortIteration(list []int) {
	var tmp []int = make([]int, len(list)) // 初始化一个和list等长的slice
	i := 1
	for i < len(list) {
		mergePass(list, tmp, i)
		i = 2 * i
		mergePass(tmp, list, i)
		i = 2 * i
	}
}

// list中相邻长度为size的子序列两两合并到tmp
func mergePass(arr, tmp []int, size int) {
	l := len(arr)
	i := 0
	for i < l-2*size+1 {
		// 两两归并
		merge(arr, tmp, i, i+size-1, i+2*size-1)
		i = i + 2*size
	}

	// 归并最后两个序列
	if i < l-size+1 {
		merge(arr, tmp, i, i+size-1, l-1)
	} else {
		// 若最后只剩下单个子序列
		for j := i; j < l; j++ {
			tmp[j] = arr[j]
		}
	}
}
