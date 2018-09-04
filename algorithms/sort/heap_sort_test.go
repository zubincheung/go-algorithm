// Create by Zubin on 2018-09-03 20:31:50

package sort

import (
	"testing"
)

func TestAdjustHeap(t *testing.T) {
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
				list: []int{6, 3, 1},
				want: []int{6, 3, 1},
			},
		},
		{
			args: args{
				list: []int{3, 6, 1},
				want: []int{6, 3, 1},
			},
		},
		{
			args: args{
				list: []int{1, 3, 6},
				want: []int{6, 3, 1},
			},
		},
		{
			args: args{
				list: []int{1, 2},
				want: []int{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AdjustHeap(tt.args.list, 0, len(tt.args.list))
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}
