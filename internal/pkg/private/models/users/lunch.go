package users

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models"
	"gorm.io/gorm"
	"time"
)

// Lunch represents a lunch
type Lunch struct {
	models.Model
	UserID   uuid.UUID `gorm:"column:user_id;not null;" json:"user_id"`
	Location string    `gorm:"column:location;not null;" json:"location"`
	Time     time.Time `gorm:"column:time;" json:"time"` //not null;
	Type     string    `gorm:"column:type;not null;" json:"type"`
	Food     string    `gorm:"column:food;not null;" json:"food"`
}

// BeforeCreate is called before creating a user
// It sets the created and updated at timestamps
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Lunch) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is called before updating a user
// It sets the updated at timestamp
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Lunch) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
