package users

import (
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models"
	"time"
)

// User represents a user
type User struct {
	models.Model
	Username  string     `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Firstname string     `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string     `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	Bio       string     `gorm:"column:bio;" json:"bio"`
	Hash      string     `gorm:"column:hash;not null;" json:"hash"`
	Hobbies   []Hobby    `gorm:"many2many:user_hobbies;"`
	Languages []Language `gorm:"many2many:user_languages;"`
	Areas     []Area     `gorm:"many2many:user_areas;"`
	Lunch     Lunch      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Buddies   []*User    `gorm:"many2many:user_buddies;association_joinTable_foreignKey:buddy_id;"`
	Blacklist []*User    `gorm:"many2many:user_blacklists;association_joinTable_foreignKey:blacklist_id;"`
	Likes     []*User    `gorm:"many2many:user_likes;association_joinTable_foreignKey:like_id;"`
}

// BeforeCreate is called before creating a user
// It sets the created and updated at timestamps
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is called before updating a user
// It sets the updated at timestamp
// It returns an error if something went wrong
// It is called by gorm
// It is not intended to be called by the user
func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
