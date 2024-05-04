package model

import "time"

type memoryHandler struct {
	todoMap map[int]*Todo
}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}

func (m *memoryHandler) Close() {

}

func (m *memoryHandler) GetTodos() []*Todo {
	var list []*Todo
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(name string) *Todo {
	id := len(m.todoMap) + 1
	newTodo := &Todo{
		ID:        id,
		Name:      name,
		Completed: false,
		CreatedAt: time.Now(),
	}
	m.todoMap[id] = newTodo
	return newTodo
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}
