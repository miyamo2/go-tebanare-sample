package domain

// TaskRepository is the port the usecase layer depends on to persist and
// look up tasks. Adapters in the repository package implement it.
type TaskRepository interface {
	Save(t *Task) error
	FindByID(id string) (*Task, error)
	ListByProject(projectID string) ([]*Task, error)
	Delete(id string) error
}
