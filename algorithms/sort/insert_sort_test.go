// Create by Zubin on 2018-08-29 18:02:17

package sort

import "testing"

func TestInsertSort(t *testing.T) {
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
			InsertSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}
