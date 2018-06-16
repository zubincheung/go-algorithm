// Create by Zubin on 2018-06-14 18:14:47

package list

import (
	"reflect"
	"testing"
)

func TestNewLinkList(t *testing.T) {
	tests := []struct {
		name string
		want *LinkList
	}{
		// TODO: Add test cases.
		{
			want: &LinkList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		data interface{}
		next *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			args: args{data: "100", next: nil},
			want: &Node{data: "100", next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.data, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
