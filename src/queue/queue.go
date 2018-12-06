package queue

// type Queue [] interface {} 支持任何类型 类似 java T
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q,v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return (*q)[0] == 0
}