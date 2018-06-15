// Create by Zubin on 2018-06-14 14:54:39

package list

import (
	"reflect"
	"testing"
)

func TestInitList(t *testing.T) {
	tests := []struct {
		name string
		want SqList
	}{
		// TODO: Add test cases.
		{
			name: "list",
			want: SqList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_ListEmpty(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "list1",
			want:   true,
			fields: fields{},
		},
		{
			name: "list2",
			want: false,
			fields: fields{
				data:   [MAXSIZE]ElemType{1},
				length: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := list.ListEmpty(); got != tt.want {
				t.Errorf("SqList.ListEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_ClearList(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "list1",
			fields: fields{
				data:   [MAXSIZE]ElemType{1},
				length: 1,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			list.ClearList()
			if list.length != 0 {
				t.Errorf("List is not empty")
			}
		})
	}
}

func TestSqList_GetElemt(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ElemType
	}{
		{
			name: "list1",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args: args{index: 2},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := list.GetElem(tt.args.index); got != tt.want {
				t.Errorf("SqList.GetElemt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_LocateElem(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	type args struct {
		e ElemType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "list1",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args: args{e: 2},
			want: 2,
		},
		{
			name: "list2",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args: args{e: 4},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := list.LocateElem(tt.args.e); got != tt.want {
				t.Errorf("SqList.LocateElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_InsertElem(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	type args struct {
		i int
		e ElemType
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		listData []ElemType
	}{
		{
			name: "list1",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args:     args{i: 2, e: 8},
			listData: []ElemType{1, 8, 2, 3},
		},
		{
			name:     "list2",
			fields:   fields{},
			args:     args{i: 1, e: 8},
			listData: []ElemType{8},
		},
		{
			name: "list2",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args:     args{i: 4, e: 8},
			listData: []ElemType{1, 2, 3, 8},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			list.InsertElem(tt.args.i, tt.args.e)

			for i := 0; i < len(tt.listData); i++ {
				if list.data[i] != tt.listData[i] {
					t.Errorf("InsertElem Error want:%#v got:%#v", tt.listData, list.data)
				}
			}

		})
	}
}

func TestSqList_Length(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "list1",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			want: 3,
		},
		{
			name:   "list2",
			fields: fields{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := list.Length(); got != tt.want {
				t.Errorf("SqList.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_DeleteElem(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]ElemType
		length int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     ElemType
		listData []ElemType
	}{
		{
			name: "list1",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args:     args{i: 3},
			listData: []ElemType{1, 2},
			want:     3,
		},
		{
			name: "list2",
			fields: fields{
				data:   [MAXSIZE]ElemType{1, 2, 3},
				length: 3,
			},
			args:     args{i: 2},
			listData: []ElemType{1, 3},
			want:     2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := list.DeleteElem(tt.args.i); got != tt.want {
				t.Errorf("SqList.DeleteElem() = %v, want %v", got, tt.want)
			}

			for i := 0; i < len(tt.listData); i++ {
				if list.data[i] != tt.listData[i] {
					t.Errorf("List Delete Error want:%#v got:%#v", tt.listData, list.data)
				}
			}
		})
	}
}
