// Create by Zubin on 2018-09-03 16:59:03

package sort

// 希尔排序
func ShellSort(list []int) {
	// 每次增量都/2
	for step := len(list) / 2; step > 0; step /= 2 {
		//从增量那组开始进行插入排序，直至完毕
		for i := step; i < len(list); i++ {
			j := i
			tmp := list[j]

			// j - step 就是代表与它同组隔壁的元素
			for j >= step && list[j-step] > tmp {
				list[j] = list[j-step]
				j = j - step
			}

			list[j] = tmp
		}
	}
}
