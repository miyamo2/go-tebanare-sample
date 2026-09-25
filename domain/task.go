package domain

import "time"

// Status is the lifecycle state of a Task.
type Status int

const (
	StatusOpen Status = iota
	StatusInProgress
	StatusDone
)

// Priority is how urgently a Task needs attention, higher first.
type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

// Task is a unit of work that belongs to a Project.
type Task struct {
	id          string
	projectID   string
	title       string
	description string
	status      Status
	priority    Priority
	assigneeID  string
	createdAt   time.Time
	dueAt       time.Time
}

// NewTask creates an open Task for projectID, due at dueAt.
func NewTask(id, projectID, title string, dueAt time.Time) (*Task, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	return &Task{
		id:        id,
		projectID: projectID,
		title:     title,
		status:    StatusOpen,
		createdAt: time.Now(),
		dueAt:     dueAt,
	}, nil
}

// ID returns the task's id.
func (t *Task) ID() string { return t.id }

// ProjectID returns the id of the project the task belongs to.
func (t *Task) ProjectID() string { return t.projectID }

// Title returns the task's title.
func (t *Task) Title() string { return t.title }

// Description returns the task's description.
func (t *Task) Description() string { return t.description }

// Status returns the task's current status.
func (t *Task) Status() Status { return t.status }

// AssigneeID returns the id of the member assigned to the task, or "" when
// unassigned.
func (t *Task) AssigneeID() string { return t.assigneeID }

// CreatedAt returns when the task was created.
func (t *Task) CreatedAt() time.Time { return t.createdAt }

// DueAt returns the task's due date.
func (t *Task) DueAt() time.Time { return t.dueAt }

// Priority returns the task's priority.
func (t *Task) Priority() Priority { return t.priority }

// SetPriority changes the task's priority. Not a getter: it takes a
// parameter.
func (t *Task) SetPriority(p Priority) { t.priority = p }

// IsOverdue reports whether the task is still open past its due date. Not a
// getter: it compares two values instead of returning a field.
func (t *Task) IsOverdue() bool {
	return t.status != StatusDone && time.Now().After(t.dueAt)
}

// Assign sets the task's assignee. Not a getter: it takes a parameter.
func (t *Task) Assign(memberID string) { t.assigneeID = memberID }

// Complete marks the task done. It rejects a task that is already done, so
// callers can decide whether that is worth surfacing.
func (t *Task) Complete() error {
	if t.status == StatusDone {
		return ErrAlreadyDone
	}
	t.status = StatusDone
	return nil
}
