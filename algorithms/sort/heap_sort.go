// Create by Zubin on 2018-09-03 20:31:50

package sort

// 调整大顶堆
func AdjustHeap(arr []int, index, length int) {
	for i := 2*index + 1; i < length; i = i*2 + 1 {
		if i+1 < length && arr[i+1] > arr[index] && arr[i+1] > arr[i] {
			// 存在右节点，且右节点比父节点和左节点大，交换右节点和父节点
			swap(arr, index, i+1)
			index = i + 1
		} else if i+1 < length && arr[i] > arr[index] && arr[i] > arr[i+1] {
			//存在左节点 左节点比父节点和右节点大，交换左节点和父节点
			swap(arr, index, i)
			index = i
		} else if arr[i] > arr[index] {
			// 左节点比父节点和右节点大，交换左节点和父节点
			swap(arr, index, i)
			index = i
		}
	}
}

// 构建大顶堆
func InitHeap(list []int) {
	for i := len(list)/2 - 1; i >= 0; i-- {
		AdjustHeap(list, i, len(list))
	}
}

// 堆排序
func HeapSort(list []int) {
	InitHeap(list)

	for i := len(list) - 1; i > 0; i-- {
		swap(list, 0, i)       // 将堆顶跟最后一个元素交换
		AdjustHeap(list, 0, i) // 重新调整大顶堆
	}
}
