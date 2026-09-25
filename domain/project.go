package domain

import "time"

// Project groups tasks under a single owner.
type Project struct {
	id        string
	name      string
	ownerID   string
	createdAt time.Time
}

// NewProject creates a Project named name, owned by ownerID.
func NewProject(id, name, ownerID string) *Project {
	return &Project{id: id, name: name, ownerID: ownerID, createdAt: time.Now()}
}

// ID returns the project's id.
func (p *Project) ID() string { return p.id }

// Name returns the project's name.
func (p *Project) Name() string { return p.name }

// OwnerID returns the id of the member who owns the project.
func (p *Project) OwnerID() string { return p.ownerID }

// CreatedAt returns when the project was created.
func (p *Project) CreatedAt() time.Time { return p.createdAt }
