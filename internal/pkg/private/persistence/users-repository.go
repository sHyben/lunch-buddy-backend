package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
)

// UserRepository is a repository for users
// It is used to access the database
// It is a singleton
type UserRepository struct{}

var userRepository *UserRepository

// GetUserRepository returns the user repository
// It creates a new one if it does not exist
// It returns the singleton instance of the user repository
func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

// Get returns a user by id
// The role is eager loaded
func (r *UserRepository) Get(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	//where.ID, _ = strconv.ParseUint(id, 10, 64)
	stringToUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	where.ID = stringToUuid
	_, err = First(&where, &user, []string{"Hobbies", "Languages", "Lunch", "Buddies", "Blacklist", "Likes", "Areas"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

// GetByUsername returns a user by username
// The role is eager loaded
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{"Hobbies", "Languages", "Lunch", "Buddies", "Blacklist", "Likes", "Areas"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

// All returns all users
// The users are ordered by id ascending
// The role is eager loaded
func (r *UserRepository) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Hobbies", "Languages", "Lunch", "Buddies", "Blacklist", "Likes", "Areas"}, "id asc")
	return &users, err
}

// Query returns all users that match the given query
// The query is a user struct with the fields to match
// The fields to match are the fields that are not nil
// The fields to match are the fields that are not empty
// The fields to match are the fields that are not zero
// The fields to match are the fields that are not the zero value for their type
//
// Example: If you want to find all users with the name "test" and the text "test"
// you would create a user struct with the name and text fields set to "test"
// and pass it to the query function
func (r *UserRepository) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{"Hobbies", "Languages", "Lunch", "Buddies", "Blacklist", "Likes", "Areas"}, "id asc")
	return &users, err
}

// Add adds a user to the database
// The role is added to the database
// The user is added to the database
func (r *UserRepository) Add(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}

// Update updates a user in the database
// The role is updated in the database
// The user is updated in the database
func (r *UserRepository) Update(user *models.User) error {
	var userRole models.UserRole
	_, err := First(models.UserRole{UserID: user.ID}, &userRole, []string{})
	//userRole.RoleName = user.Role.RoleName
	err = Save(&userRole)
	err = db.GetDB().Omit("Hobbies", "Languages", "Lunch", "Buddies", "Blacklist", "Likes", "Areas").Save(&user).Error
	//user.Role = userRole
	return err
}

// Delete deletes a user from the database
// The role is deleted from the database
// The user is deleted from the database
func (r *UserRepository) Delete(user *models.User) error {
	err := db.GetDB().Unscoped().Delete(models.UserRole{UserID: user.ID}).Error
	err = db.GetDB().Unscoped().Delete(&user).Error
	return err
}
