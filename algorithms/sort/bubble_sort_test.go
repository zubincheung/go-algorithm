// Create by Zubin on 2018-08-28 15:25:36

package sort

import (
	"testing"
)

func TestSimpleBubbleSort(t *testing.T) {
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
			SimpleBubbleSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
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
			BubbleSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}

func TestFlagSwapBubbleSort(t *testing.T) {
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
			FlagSwapBubbleSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}

func TestFlagSwapPositionBubbleSort(t *testing.T) {
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
			FlagSwapPositionBubbleSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}
