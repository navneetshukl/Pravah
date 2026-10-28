package main

import (
	"sync"
)

type Queue struct {
	arr    []Task
	mut    *sync.Mutex
	isOpen bool
	cond   *sync.Cond
	size   int
}

func NewQueue() *Queue {
	q := &Queue{
		arr:    make([]Task, 0),
		mut:    &sync.Mutex{},
		isOpen: true,
		size:   0,
	}
	q.cond = sync.NewCond(q.mut)
	return q
}

func (q *Queue) Enqueue(task Task) bool {
	q.mut.Lock()
	defer q.mut.Unlock()
	if !q.isOpen {
		return false
	}
	q.arr = append(q.arr, task)
	q.size++
	q.cond.Signal()
	return true
}

func (q *Queue) Dequeue() (Task, bool) {
	q.mut.Lock()
	defer q.mut.Unlock()
	for q.size == 0 && q.isOpen {
		q.cond.Wait()
	}
	if q.size == 0 {
		return Task{}, false
	}
	val := q.arr[0]
	q.arr = q.arr[1:]
	q.size--
	q.cond.Signal()
	return val, true
}

func (q *Queue) Close() {
	q.mut.Lock()
	defer q.mut.Unlock()
	q.isOpen = false
	q.cond.Broadcast()
}

func (q *Queue) Size() int {
	q.mut.Lock()
	defer q.mut.Unlock()
	return q.size
}
