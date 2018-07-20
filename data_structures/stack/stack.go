// Create by Zubin on 2018-07-19 16:24:02
package stack

const MAXSIZE = 20

type Stack struct {
	data [MAXSIZE]interface{}
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

// 获取栈的长度
func (stack *Stack) Length() int {
	return stack.size
}

// 清空栈
func (stack *Stack) Clear() {

	stack.size = 0
}

// 是否空栈
func (stack *Stack) IsEmpty() bool {
	return stack.Length() == 0
}

// 入栈
func (stack *Stack) push(elem interface{}) {
	if stack.Length() >= MAXSIZE {
		panic("Stack is full")
	}

	index := stack.Length()

	stack.data[index] = elem
	stack.size++
}

// 出栈
func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic("Stack is empty")
	}

	e := stack.data[stack.size-1]
	stack.data[stack.size-1] = nil

	stack.size--

	return e
}
