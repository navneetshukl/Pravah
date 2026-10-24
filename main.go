package main

import (
	"log"
	"sync"
)

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

type MemStore struct {
	store map[int]Task
	mut   *sync.Mutex
	index int
}

func newMemStore() *MemStore {
	return &MemStore{
		store: make(map[int]Task),
		mut:   &sync.Mutex{},
		index: 0,
	}
}

func (m *MemStore) AddTask(value_type string, data interface{}) bool {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.index++
	m.store[m.index] = newTask(m.index, value_type, PENDING.ToString(), data)
	return true
}

func main() {
	store := newMemStore()
	for i := 1; i < 100; i++ {
		store.AddTask("queue", i)
	}
	log.Println(store.store)
}
