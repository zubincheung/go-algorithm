// Create by Zubin on 2018-06-21 17:40:59

package list

type node struct {
	data interface{}
	next *node
	prev *node
}

type CircularLinkList struct {
	head   *node
	tail   *node
	length int
}

// 初始化一个循环链表
func NewCircularLinkList() *CircularLinkList {
	return &CircularLinkList{}

}

// 链表是否为空
func (list *CircularLinkList) ListEmpty() bool {
	return list.length == 0
}

// 清空链表
func (list *CircularLinkList) ClearList() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// 返回链表中元素的个数
func (list *CircularLinkList) Size() int {
	return list.length
}

// 返回链表中第i个元素
func (list *CircularLinkList) GetElem(index int) interface{} {
	if list.length <= 0 {
		panic("list is empty")
	}

	if index < 1 || index > list.length {
		panic("invalid index ")
	}

	var e interface{}
	l := list.head

	for i := 0; i < index; i++ {
		e = l.data
		l = l.next
	}

	return e
}

// 链表中查找元素e
// 返回在链表中第1个满足关系的数据元素的位序，如果不存在返回0
func (list *CircularLinkList) LocateElem(e interface{}) int {

	if list.length <= 0 {
		panic("list is empty")
	}

	index := 1
	l := list.head

	for index <= list.length {
		if l.data == e {
			return index
		}

		index++
		l = l.next
	}

	return 0
}

// 在第一个结点插入元素e
func (list *CircularLinkList) AddToFirst(e interface{}) {

	n := &node{
		data: e,
	}

	if list.length == 0 {
		list.head = n
		list.tail = n

		n.next = n
		n.prev = n
	} else {

		n.next = list.head
		n.prev = list.tail

		list.head.prev = n
		list.tail.next = n
		list.head = n
	}

	list.length++
	n = nil
}

// 在最后一个结点插入元素e
func (list *CircularLinkList) AddToLast(e interface{}) {

	n := &node{
		data: e,
	}

	if list.length == 0 {
		list.head = n
		list.tail = n

		n.next = n
		n.prev = n
	} else {
		n.next = list.head
		n.prev = list.tail

		list.tail.next = n
		list.head.prev = n

		list.tail = n
	}

	list.length++
	n = nil
}

// 插入新的结点
func (list *CircularLinkList) Add(index int, elem interface{}) {
	if index < 1 || index > list.length+1 {
		panic("invalid index ")
	}

	if index == 1 {
		list.AddToFirst(elem)
		return
	} else if index == list.length+1 {
		list.AddToLast(elem)
		return
	}

	n := &node{
		data: elem,
	}
	l := list.head
	for i := 1; i < index; i++ {
		l = l.next
	}

	p := l.prev
	n.next = l
	n.prev = p

	l.prev = n
	p.next = n

	list.length++
	p = nil
	l = nil
	n = nil

}

// 移除第一个结点
func (list *CircularLinkList) RemoveFirst() interface{} {
	if list.ListEmpty() {
		panic("list empty")
	}
	e := list.head.data
	if list.length == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.head.next.prev = list.tail
		list.head = list.head.next
		list.tail.next = list.head
	}

	list.length--
	return e
}

// 移除最后一个结点
func (list *CircularLinkList) RemoveLast() interface{} {
	if list.ListEmpty() {
		panic("list empty")
	}
	e := list.tail.data

	if list.length == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.tail.prev.next = list.head
		list.tail = list.tail.prev
		list.head.prev = list.tail
	}

	list.length--
	return e
}

// 删除第i个元素
func (list *CircularLinkList) Remove(index int) interface{} {
	if list.ListEmpty() {
		panic("list empty")
	}
	if index > list.length {
		panic("invalid index ")
	}

	if index == 1 {
		return list.RemoveFirst()
	} else if index == list.length {
		return list.RemoveLast()
	}

	// 查找待删除的结点
	p := list.head

	for i := 1; i < index; i++ {
		p = p.next
	}

	p.prev.next = p.next
	p.next.prev = p.prev

	e := p.data
	p = nil
	list.length--

	return e
}

// 反转链表
func (list *CircularLinkList) Reverse() {
	l := list.head
	list.head = list.tail
	list.tail = l
	for i := 0; i < list.length; i++ {
		t := l.next
		l.next = l.prev
		l.prev = t

		l = t
	}
}
