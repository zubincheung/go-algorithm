// Create by Zubin on 2018-06-14 18:14:47

package list

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkList struct {
	head   *Node
	length int
}

// 新建一个节点
func NewNode(data interface{}, next *Node) *Node {
	return &Node{data, next}
}

// 初始化一个线性表
func NewLinkList() *LinkList {
	return &LinkList{}
}

// 线性表是否为空
func (list *LinkList) ListEmpty() bool {
	return list.head == nil
}

// 清空线性表
func (list *LinkList) ClearList() {
	var l1, l2 *Node
	l1 = list.head

	for l1 != nil {
		l1.data = nil

		l2 = l1.next
		l1.next = nil
		l1 = l2
	}

	// 头结点指针域为空
	list.head = nil
	list.length = 0
}

// 返回线性表中元素的个数
func (list *LinkList) Size() int {
	return list.length
}

// 返回线性表中第i个元素
func (list *LinkList) GetElem(index int) (interface{}, error) {
	if list.length <= 0 {
		return nil, fmt.Errorf("list is empty")
	}

	if index < 1 || index > list.length {
		return nil, fmt.Errorf("invalid index value %d", index)
	}

	var e interface{}
	l := list.head

	for i := 0; i < index; i++ {
		e = l.data
		l = l.next
	}

	return e, nil
}

// 线性表中查找元素e
// 返回在线性表中第1个满足关系的数据元素的位序，如果不存在返回0
func (list *LinkList) LocateElem(e interface{}) (int, error) {

	if list.length <= 0 {
		return 0, fmt.Errorf("list is empty")
	}

	index := 1
	l := list.head

	for l != nil {
		if l.data == e {
			return index, nil
		}

		index++
		l = l.next
	}

	return 0, nil
}

// 在第一个结点插入元素e
func (list *LinkList) AddToFirst(e interface{}) error {
	list.head = NewNode(e, list.head)
	list.length++
	return nil
}

// 在最后一个结点插入元素e
func (list *LinkList) AddToLast(e interface{}) error {
	if list.ListEmpty() {
		// 线性表为空插入第一个结点
		return list.AddToFirst(e)
	}

	// 查找最后一个节点
	l := list.head
	for l.next != nil {
		l = l.next
	}

	l.next = NewNode(e, nil)
	list.length++
	return nil
}

// 在线性表第i个位置之前插入元素e
func (list *LinkList) Insert(i int, e interface{}) error {
	if i < 1 || i > list.length+1 {
		return fmt.Errorf("invalid index value %d", i)
	}

	if i == 1 {
		// 插入第一个结点
		return list.AddToFirst(e)
	} else if i == list.length+1 {
		// 插入最后一个节点
		return list.AddToLast(e)
	}

	l := list.head
	j := 1

	// 查找第i个结点
	for j < i {
		l = l.next
		j++
	}

	node := NewNode(l.data, l.next) // 把新节点的后继指向l的后继节点，

	l.data = e
	l.next = node //l的后继节点指向新节点

	list.length++
	return nil

}

// 移除第一个结点
func (list *LinkList) RemoveFirst() (interface{}, error) {
	if list.ListEmpty() {
		return nil, fmt.Errorf("list empty")
	}

	e := list.head.data

	list.head = list.head.next
	list.length--
	return e, nil
}

// 移除最后一个结点
func (list *LinkList) RemoveLast() (interface{}, error) {
	return list.Remove(list.length)
}

// 删除第i个元素
func (list *LinkList) Remove(i int) (interface{}, error) {
	if list.ListEmpty() {
		return nil, fmt.Errorf("list empty")
	}
	if i > list.length {
		return nil, fmt.Errorf("invalid index value %d", i)
	}

	if i == 1 {
		return list.RemoveFirst()
	}

	// 查找待删除元素的上一个结点
	l := list.head
	for j := 1; j < i-1; j++ {
		l = l.next
	}

	e := l.next.data
	l.next = l.next.next
	list.length--

	return e, nil
}
