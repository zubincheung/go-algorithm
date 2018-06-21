// Create by Zubin on 2018-06-14 18:14:47

package list

import (
	"reflect"
	"testing"
)

type check func(list *LinkList) bool

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

func TestLinkList_ListEmpty(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{},
			want:   true,
		},
		{
			fields: fields{head: NewNode("node1", nil), length: 1},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := list.ListEmpty(); got != tt.want {
				t.Errorf("LinkList.ListEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_ClearList(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			fields: fields{},
		},
		{
			fields: fields{head: NewNode("node1", nil), length: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			list.ClearList()
			if !list.ListEmpty() {
				t.Error("list  not empty")
			}
		})
	}
}

func TestLinkList_Size(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{},
			want:   0,
		},
		{
			fields: fields{head: NewNode("node1", nil), length: 1},
			want:   1,
		},
		{
			fields: fields{head: NewNode("node1", NewNode("node2", NewNode("node3", nil))), length: 3},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := list.Size(); got != tt.want {
				t.Errorf("LinkList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_GetElem(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			fields:  fields{},
			args:    args{index: 2},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", nil), length: 1},
			args:    args{index: 2},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", NewNode("node3", nil))), length: 3},
			args:    args{index: 2},
			wantErr: false,
			want:    "node2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			got, err := list.GetElem(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkList.GetElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkList.GetElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_LocateElem(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		e interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields:  fields{},
			args:    args{e: 2},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", nil), length: 1},
			args:    args{e: "node 2"},
			wantErr: false,
			want:    0,
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", NewNode("node3", nil))), length: 3},
			args:    args{e: "node2"},
			wantErr: false,
			want:    2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			got, err := list.LocateElem(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkList.LocateElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LinkList.LocateElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Insert(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}

	type args struct {
		i int
		e interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		checkFunc check
	}{
		// TODO: Add test cases.
		{
			fields:  fields{},
			args:    args{e: "node1", i: 2},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", nil), length: 1},
			args:    args{e: "node 2", i: 3},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", nil), length: 1},
			args:    args{e: "node2", i: 1},
			wantErr: false,
			checkFunc: func(list *LinkList) bool {

				if data, _ := list.GetElem(1); data != "node2" {
					return false
				}

				if data, _ := list.GetElem(2); data != "node1" {
					return false
				}

				if list.Size() != 2 {
					return false
				}

				return true
			},
		},

		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			args:    args{e: "node3", i: 3},
			wantErr: false,
			checkFunc: func(list *LinkList) bool {

				if data, _ := list.GetElem(2); data != "node2" {
					return false
				}

				if data, _ := list.GetElem(3); data != "node3" {
					return false
				}

				if list.Size() != 3 {
					return false
				}

				return true
			},
		},

		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			args:    args{e: "node3", i: 2},
			wantErr: false,
			checkFunc: func(list *LinkList) bool {

				if data, _ := list.GetElem(3); data != "node2" {
					return false
				}

				if data, _ := list.GetElem(2); data != "node3" {
					return false
				}

				if list.Size() != 3 {
					return false
				}

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if err := list.Insert(tt.args.i, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("LinkList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.checkFunc != nil && !tt.checkFunc(list) {
				t.Error("insert error")
			}
		})
	}
}

func TestLinkList_AddToLast(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		e interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		checkFunc check
	}{
		// TODO: Add test cases.
		{
			fields:  fields{},
			args:    args{e: "node2"},
			wantErr: false,
			checkFunc: func(list *LinkList) bool {

				if data, _ := list.GetElem(1); data != "node2" {
					return false
				}

				if list.Size() != 1 {
					return false
				}

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if err := list.AddToLast(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("LinkList.AddToLast() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkList_RemoveFirst(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields:  fields{},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			wantErr: false,
			want:    "node1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			got, err := list.RemoveFirst()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkList.RemoveFirst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkList.RemoveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_RemoveLast(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields:  fields{},
			wantErr: true,
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			wantErr: false,
			want:    "node2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			got, err := list.RemoveLast()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkList.RemoveLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkList.RemoveLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Remove(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields:  fields{},
			wantErr: true,
		},

		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			wantErr: true,
			args:    args{i: 13},
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			wantErr: false,
			args:    args{i: 1},
			want:    "node1",
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", nil)), length: 2},
			wantErr: false,
			args:    args{i: 2},
			want:    "node2",
		},
		{
			fields:  fields{head: NewNode("node1", NewNode("node2", NewNode("node3", NewNode("node4", nil)))), length: 4},
			wantErr: false,
			args:    args{i: 3},
			want:    "node3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			got, err := list.Remove(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkList.Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkList.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
