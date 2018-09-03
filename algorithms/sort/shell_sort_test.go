// Create by Zubin on 2018-09-03 16:59:03

package sort

import "testing"

func TestShellSort(t *testing.T) {
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
			ShellSort(tt.args.list)
			if !SliceEquals(tt.args.list, tt.args.want) {

				t.Error("sort error", tt.args.list)
			}
		})
	}
}
