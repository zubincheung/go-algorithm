// Create by Zubin on 2018-08-29 18:02:17

package sort

// 插入排序
func InsertSort(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && list[j-1] > list[j]; j-- {
			swap(list, j-1, j)
		}
	}
}
