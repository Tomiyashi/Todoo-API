package memory

import (
	"errors"
	"sync"
	"tasks-api/internal/models"
)

type TaskStorage struct {
	mu     sync.RWMutex
	task   map[int]models.Task
	nextID int
}

func New() *TaskStorage {
	return &TaskStorage{
		task:   map[int]models.Task{},
		nextID: 1,
	}
}

func (t *TaskStorage) Create(tasks models.Task) (models.Task, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	tasks.ID = t.nextID
	t.task[t.nextID] = tasks
	t.nextID++
	return tasks, nil
}

func (t *TaskStorage) Get(id int) (models.Task, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	val, exists := t.task[id]
	return val, exists
}

func (t *TaskStorage) List() []models.Task {
	t.mu.RLock()
	defer t.mu.RUnlock()

	slise := make([]models.Task, 0, len(t.task))

	for _, v := range t.task {
		slise = append(slise, v)
	}
	return slise
}

func (t *TaskStorage) Update(id int, task models.Task) (models.Task, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	val, exists := t.task[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
}
