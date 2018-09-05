// Create by Zubin on 2018-09-04 19:00:11
package sort

import "testing"

func TestQuickSort(t *testing.T) {
	type args struct {
		list []int
		want []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				list: []int{6, 3, 1, 8, 5, 9},
				want: []int{1, 3, 5, 6, 8, 9},
			},
		},
		{
			args: args{
				list: []int{6, 3, 11, 10, 1, 18, 9, 8, 5, 4},
				want: []int{1, 3, 4, 5, 6, 8, 9, 10, 11, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {
				t.Error("sort error", tt.args.list)
			}
		})
	}
}
