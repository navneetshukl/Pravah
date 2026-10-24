package main

type TaskStatus string

var (
	PENDING   TaskStatus = "pending"
	RUNNING   TaskStatus = "running"
	COMPLETED TaskStatus = "completed"
)

func (t TaskStatus) ToString() string {
	return string(t)
}
