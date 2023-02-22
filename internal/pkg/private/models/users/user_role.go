package users

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models"
	"gorm.io/gorm"
	"time"
)

// UserRole represents a user role
// It is used to map users to roles
// It is used to implement a many-to-many relationship between users and roles
type UserRole struct {
	models.Model
	UserID uuid.UUID `gorm:"column:user_id;unique_index:user_role;not null;" json:"user_id"`
	//UserID   uint64 `gorm:"column:user_id;unique_index:user_role;not null;" json:"user_id"`
	RoleName string `gorm:"column:role_name;not null;" json:"role_name"`
}

// BeforeCreate is called before creating a user role
// It sets the created and updated at timestamps
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *UserRole) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is called before updating a user role
// It sets the updated at timestamp
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *UserRole) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
