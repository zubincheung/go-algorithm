// Create by Zubin on 2018-08-28 15:34:56

// Create by Zubin on 2018-08-28 15:30:33

package sort

import "testing"

func Test_swap(t *testing.T) {
	t.Run("swap", func(t *testing.T) {
		list := []int{1, 2, 3}
		swap(list, 0, 1)

		if list[0] != 2 || list[1] != 1 || list[2] != 3 {
			t.Error("swap error")
		}
	})
}
