package goqueue

// Queue is the int queue type
type Queue []int

// Get new int queue
func New() *Queue {
	return &Queue{}
}

// Push values (int's) into the queue
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop values (int's) from the queue
func (q Queue) Pop() int {
	temp := q[0]
	q = q[1:len(q)]
	return temp
}
