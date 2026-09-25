package repository

import (
	"sync"

	"github.com/miyamo2/go-tebanare-sample/domain"
)

// InMemoryTaskRepository stores tasks in memory. It exists so the sample
// runs without a database; a real deployment would swap in a SQL or NoSQL
// adapter behind the same domain.TaskRepository port.
type InMemoryTaskRepository struct {
	mu    sync.Mutex
	tasks map[string]*domain.Task
}

// NewInMemoryTaskRepository creates an empty InMemoryTaskRepository.
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{tasks: make(map[string]*domain.Task)}
}

// Save stores t, overwriting any task with the same id.
func (r *InMemoryTaskRepository) Save(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[t.ID()] = t
	return nil
}

// FindByID returns the task with the given id.
func (r *InMemoryTaskRepository) FindByID(id string) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.tasks[id]
	if !ok {
		return nil, domain.ErrTaskNotFound
	}
	return t, nil
}

// ListByProject returns every task that belongs to projectID.
func (r *InMemoryTaskRepository) ListByProject(projectID string) ([]*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var tasks []*domain.Task
	for _, t := range r.tasks {
		if t.ProjectID() == projectID {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

// Close satisfies io.Closer for adapters that need it. Not a noop match: it
// returns a value.
func (r *InMemoryTaskRepository) Close() error { return nil }

// Migrate is a no-op: the in-memory store needs no schema migration.
func (r *InMemoryTaskRepository) Migrate() {}
