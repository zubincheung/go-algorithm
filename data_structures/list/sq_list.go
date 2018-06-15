// Create by Zubin on 2018-06-14 14:31:49
package list

const MAXSIZE = 20

type SqList struct {
	data   [MAXSIZE]interface{}
	length int
}

// 初始化一个新的线性表
func InitList() SqList {
	// var data [MAXSIZE]interface{}
	return SqList{}
}

// 线性表是否为空
func (list *SqList) ListEmpty() bool {
	return list.length == 0
}

// 清空线性表
func (list *SqList) ClearList() {
	list.length = 0
}

// 返回线性表中第i个位置的元素
func (list *SqList) GetElem(index int) interface{} {
	if index < 1 || index > list.length {
		panic("Ivalid index")
	}
	if list.length == 0 {
		panic("Empty list")
	}

	return list.data[index-1]
}

// 在线性表中查找与给定元素e相等的元素，
// 如果查找成果则返回在表中的序号，失败则返回0
func (list *SqList) LocateElem(e interface{}) int {
	if list.length == 0 {
		panic("Empty list")
	}

	for index := 0; index < list.length; index++ {
		if list.data[index] == e {
			return index + 1
		}
	}

	return 0
}

// 返回线性表长度
func (list *SqList) Length() int {
	return list.length
}

// 在指定位置插入新元素e
func (list *SqList) InsertElem(i int, e interface{}) {
	if i < 0 || i > list.length+1 {
		panic("Ivalid i")
	}

	if list.length == MAXSIZE {
		panic("List is full")
	}

	// 线性表不为空且插入的位置在中间，向后移动插入位置后面的元素
	if list.length > 0 && i <= list.length {
		for index := list.length - 1; index >= i-1; index-- {
			list.data[index+1] = list.data[index]
		}
	}

	list.data[i-1] = e
	list.length++
}

// 删除线性表中指定位置的元素，并返回删除的那个元素
func (list *SqList) DeleteElem(i int) interface{} {
	if i < 1 || i > list.length {
		panic("Ivalid i")
	}
	if list.length == 0 {
		panic("Empty list")
	}

	e := list.data[i-1]

	// 如果删除的不是最后的位置
	if i < list.length {
		for index := i - 1; index < list.length-1; index++ {
			// 将删除位置后面的元素前移
			list.data[index] = list.data[index+1]
		}
	}

	list.length--

	return e
}
