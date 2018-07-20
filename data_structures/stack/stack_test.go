// Create by Zubin on 2018-07-19 16:47:03

package stack

import (
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		stack := NewStack()

		if !stack.IsEmpty() {
			t.Error("Stack is not empty")
		}

		stack.push(1)
		stack.push(2)
		stack.push(3)
		stack.push(4)

		if stack.Length() != 4 {
			t.Error("Stack size must 4")
		}

		e := stack.Pop()
		if e != 4 || stack.Length() != 3 {
			t.Error("e must equal 4 and stack size must 3")
		}

		e = stack.Pop()
		if e != 3 || stack.Length() != 2 {
			t.Error("e must equal 3 and stack size must 2")
		}

		stack.push(5)

		e = stack.Pop()
		if e != 5 || stack.Length() != 2 {
			t.Error("e must equal 3 and stack size must 2")
		}

		stack.Clear()

		if !stack.IsEmpty() {
			t.Error("Stack is not empty")
		}

		// if got := stack.Pop(); !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
		// }
	})
}
