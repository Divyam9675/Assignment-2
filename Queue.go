package main

import "fmt"


type Queue struct {
	items []int
}


func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}


func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	fmt.Println(myQueue)
	myQueue.Dequeue()

}