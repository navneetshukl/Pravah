package main

import "log"

type Task struct {
	id         int
	value_type string
	data       interface{}
	status     string
}

func newTask(id int, value_type, status string, data interface{}) Task {
	return Task{
		id:         id,
		value_type: value_type,
		data:       data,
		status:     status,
	}
}

func main() {

	q:=NewQueue()

	log.Println("Before Memory")
	PrintMemoryUsageMB()

	for i:=1;i<10000;i++{
		task:=newTask(i,"int",PENDING.ToString(),i)
		q.Enqueue(task)

	}

	log.Println("Total Added Task ",len(q.arr))

	log.Println("After Memory")
	PrintMemoryUsageMB()

}
