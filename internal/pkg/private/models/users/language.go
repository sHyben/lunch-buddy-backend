package users

import (
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models"
	"gorm.io/gorm"
	"time"
)

// Language represents a language
type Language struct {
	models.Model
	//Location column is an enum representation of the location of the area
	Name string `gorm:"column:name;unique_index:name;not null;" json:"name"`
}

// BeforeCreate is called before creating a user
// It sets the created and updated at timestamps
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Language) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is called before updating a user
// It sets the updated at timestamp
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *Language) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
