// Create by Zubin on 2018-07-20 11:56:27

package queue

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		queue := NewQueue()

		if !queue.isEmpty() {
			t.Error("queue must empty")
		}

		if queue.Length() != 0 {
			t.Error("queue length must 0")
		}

		for index := 0; index < MAXSIZE-1; index++ {
			queue.EnQueue(index)
		}

		if !queue.isFull() {
			t.Error("queue must full")
		}

		if queue.Length() != MAXSIZE-1 {
			t.Errorf("queue length must %d", MAXSIZE-1)
		}

		if queue.OutQueue().(int) != 0 {
			t.Error("queue out must 0")
		}

		if queue.OutQueue().(int) != 1 {
			t.Error("queue out must 1")
		}

		queue.EnQueue(21)

	})
}
