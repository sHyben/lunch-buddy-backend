package users

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models"
	"time"
)

// Lunch represents a lunch
type Lunch struct {
	models.Model
	UserID   uuid.UUID
	Location string    `gorm:"column:location;unique_index:location;not null;" json:"location"`
	Time     time.Time `gorm:"column:time;unique_index:time;not null;" json:"time"`
}

// BeforeCreate is called before creating a user
// It sets the created and updated at timestamps
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Lunch) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is called before updating a user
// It sets the updated at timestamp
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Lunch) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
