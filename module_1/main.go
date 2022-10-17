package main

import (
	"fmt"
	"sync"
)

func enqueue(queue []Caller, element Caller) []Caller {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}
func dequeue(queue []Caller) []Caller {
	element := queue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
}

type Caller struct {
	operator *Operator
	num      int
}
type Operator struct {
	lock    sync.Mutex
	callers []Caller
	num     int
}

func operatorWorks(operator *Operator) {
	fmt.Println("Hello!")
	for {
		//operator.lock.Lock()
		operator.callers = dequeue(operator.callers)
		//operator.lock.Unlock()
	}
}
func callerCall(caller Caller) {
	caller.operator.callers = enqueue(caller.operator.callers, caller)
}

func main() {
	var arr []Caller
	op1 := Operator{sync.Mutex{}, arr, 1}
	op2 := Operator{sync.Mutex{}, arr, 2}

	cal_1 := Caller{&op1, 1}
	fmt.Print(cal_1)
	cal_2 := Caller{&op2, 2}
	cal_3 := Caller{&op1, 3}
	cal_4 := Caller{&op2, 4}
	cal_5 := Caller{&op1, 5}

	go callerCall(cal_1)
	go callerCall(cal_2)
	go callerCall(cal_3)
	go callerCall(cal_4)
	go callerCall(cal_5)

	go operatorWorks(&op1)
	go operatorWorks(&op2)

}
