package usecase

import (
	"fmt"
	"time"

	"github.com/miyamo2/go-tebanare-sample/domain"
)

// TaskUseCase implements the application's task-related use cases on top
// of a TaskRepository.
type TaskUseCase struct {
	repo domain.TaskRepository
	log  Logger
}

// NewTaskUseCase creates a TaskUseCase backed by repo, logging through log.
func NewTaskUseCase(repo domain.TaskRepository, log Logger) *TaskUseCase {
	return &TaskUseCase{repo: repo, log: log}
}

// CreateTask builds a Task for projectID and persists it.
func (uc *TaskUseCase) CreateTask(id, projectID, title string, dueAt time.Time) (*domain.Task, error) {
	t, err := domain.NewTask(id, projectID, title, dueAt)
	if err != nil {
		return nil, err
	}
	err = uc.repo.Save(t)
	if err != nil {
		return nil, err
	}
	uc.log.Info(fmt.Sprintf("created task %s", t.ID()))
	return t, nil
}

// GetTask returns the task with the given id.
func (uc *TaskUseCase) GetTask(id string) (*domain.Task, error) {
	return uc.repo.FindByID(id)
}

// CompleteTask marks a task done and persists the change.
func (uc *TaskUseCase) CompleteTask(id string) error {
	t, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	err = t.Complete()
	if err != nil {
		return fmt.Errorf("complete task %s: %w", id, err)
	}
	return uc.repo.Save(t)
}
