package tasks

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
	"time"
)

// Task represents a task
// It is used to map users to tasks
// It is used to implement a many-to-many relationship between users and tasks
type Task struct {
	models.Model
	Name   string     `gorm:"column:name;not null;" json:"name" form:"name"`
	Text   string     `gorm:"column:text;not null;" json:"text" form:"text"`
	UserID uuid.UUID  `gorm:"column:user_id;unique_index:user_id;not null;" json:"user_id" form:"user_id"`
	User   users.User `json:"user"`
}

// BeforeCreate is called before creating a task
// It sets the created and updated at timestamps
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Task) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is called before updating a task
// It sets the updated at timestamp
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Task) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
