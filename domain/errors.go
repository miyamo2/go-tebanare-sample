package domain

import "errors"

var (
	// ErrEmptyTitle is returned when a Task is created with an empty title.
	ErrEmptyTitle = errors.New("title must not be empty")
	// ErrAlreadyDone is returned when Complete is called on a Task that is
	// already done.
	ErrAlreadyDone = errors.New("task is already done")
	// ErrTaskNotFound is returned when a lookup finds no matching task.
	ErrTaskNotFound = errors.New("task not found")
)
