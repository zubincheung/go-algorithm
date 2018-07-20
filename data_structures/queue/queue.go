// Create by Zubin on 2018-07-19 17:44:48
package queue

const MAXSIZE = 20

type Queue struct {
	data  [MAXSIZE]interface{}
	front int
	rear  int
}

func NewQueue() *Queue {
	return &Queue{}
}

// 返回队列长度
func (queue *Queue) Length() int {
	return (queue.rear - queue.front + MAXSIZE) % MAXSIZE
}

// 队列是否为空
func (queue *Queue) isEmpty() bool {
	return queue.rear == queue.front
}

// 队列是否已满
func (queue *Queue) isFull() bool {
	return (queue.rear+1)%MAXSIZE == queue.front
}

// 入列
func (queue *Queue) EnQueue(elem interface{}) {
	if queue.isFull() {
		panic("queue is full")
	}

	queue.data[queue.rear] = elem
	queue.rear = (queue.rear + 1) % MAXSIZE
}

// 出列
func (queue *Queue) OutQueue() interface{} {
	if queue.isEmpty() {
		panic("queue is empty")
	}

	e := queue.data[queue.front]
	queue.front = (queue.front + 1) % MAXSIZE

	return e
}
