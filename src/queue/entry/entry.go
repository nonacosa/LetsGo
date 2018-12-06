package main

import "queue"

func main() {
	// 这里可以存放原始值
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	q.Pop()
	q.IsEmpty()
	q.Pop()
	q.IsEmpty()
}
