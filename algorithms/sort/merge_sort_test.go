// Create by Zubin on 2018-09-04 20:04:14

package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
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
				list: []int{6, 3, 11, 1, 8, 5, 9},
				want: []int{1, 3, 5, 6, 8, 9, 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}

func TestMergeSortIteration(t *testing.T) {
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
				list: []int{6, 3, 11, 1, 8, 5, 9},
				want: []int{1, 3, 5, 6, 8, 9, 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortIteration(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {
				t.Error("sort error", tt.args.list)
			}
		})
	}
}
