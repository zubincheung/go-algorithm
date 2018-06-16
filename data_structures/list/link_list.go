// Create by Zubin on 2018-06-14 18:14:47

package list

import (
	"fmt"
	"strconv"
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

// 初始化一个线性表
func NewLinkList1() *LinkList {
	list := NewLinkList()
	l := list.head

	for index := 0; index < 10; index++ {
		l = NewNode("元素"+strconv.Itoa(index), nil)
		l = l.next
	}

	list.length = 10

	return list
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
	if list.ListEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if index < 0 || index > list.Size() {
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
