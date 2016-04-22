package main

import "fmt"

type queue struct {
	qu []int
}

var pop bool = false

func main() {
	var t queue
	var p queue

	// Print the starting values
	fmt.Println("--------STARTING--------")
	fmt.Printf("queue: %v \n", t)
	fmt.Printf("Popped numbers: %v \n\n", p)

	// Push some numbers into the queue
	fmt.Println("--------PUSH--------")
	t.push(1)
	t.push(2)
	t.push(3)
	t.push(4)
	fmt.Printf("queue: %v \n\n", t)

	// Pop a number
	fmt.Println("--------POP--------")
	pop = true
	a := t.pop()
	p.push(a)

	// Pop another number
	a = t.pop()
	p.push(a)
	pop = false

	// Print the ending values
	fmt.Println("\n--------END--------")
	fmt.Printf("queue: %v \n", t)
	fmt.Printf("Popped numbers: %v \n", p)
}

func (q *queue) push(v int) {
	q.qu = append(q.qu, v)
	if pop != true {
		fmt.Printf("Pushed value: %v \n", v)
	}
}

func (q *queue) pop() int {
	temp := q.qu[0]
	q.qu = q.qu[1:len(q.qu)]
	fmt.Printf("Popped value: %v \n", temp)
	return temp
}
