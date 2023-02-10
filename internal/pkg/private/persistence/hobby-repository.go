package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
)

// HobbyRepository is a repository for hobbies
// It is used to access the database
// It is a singleton
type HobbyRepository struct{}

var hobbyRepository *HobbyRepository

// GetHobbyRepository returns the hobby repository
// It creates a new one if it does not exist
// It returns the singleton instance of the hobby repository
func GetHobbyRepository() *HobbyRepository {
	if hobbyRepository == nil {
		hobbyRepository = &HobbyRepository{}
	}
	return hobbyRepository
}

// Get returns a hobby by id
func (r *HobbyRepository) Get(id string) (*models.Hobby, error) {
	var hobby models.Hobby
	where := models.Hobby{}
	stringToUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	where.ID = stringToUuid
	_, err = First(&where, &hobby, []string{})
	if err != nil {
		return nil, err
	}
	return &hobby, err
}

// GetByName returns a hobby by name
func (r *HobbyRepository) GetByName(name string) (*models.Hobby, error) {
	var hobby models.Hobby
	where := models.Hobby{}
	where.Name = name
	_, err := First(&where, &hobby, []string{})
	if err != nil {
		return nil, err
	}
	return &hobby, err
}

// All returns all hobbies
// The hobbies are ordered by id ascending
func (r *HobbyRepository) All() (*[]models.Hobby, error) {
	var hobbies []models.Hobby
	err := Find(&models.Hobby{}, &hobbies, []string{}, "id asc")
	return &hobbies, err
}

// Query returns all hobbies that match the given query
// The query is a user struct with the fields to match
// The fields to match are the fields that are not nil
// The fields to match are the fields that are not empty
// The fields to match are the fields that are not zero
// The fields to match are the fields that are not the zero value for their type
func (r *HobbyRepository) Query(q *models.Hobby) (*[]models.Hobby, error) {
	var hobbies []models.Hobby
	err := Find(&q, &hobbies, []string{}, "id asc")
	return &hobbies, err
}

// Add adds a hobby to the database
func (r *HobbyRepository) Add(hobby *models.Hobby) error {
	err := Create(&hobby)
	err = Save(&hobby)
	return err
}

// Update updates a hobby in the database
func (r *HobbyRepository) Update(hobby *models.Hobby) error {
	return db.GetDB().Save(&hobby).Error
}

// Delete deletes a hobby from the database
func (r *HobbyRepository) Delete(hobby *models.Hobby) error {
	return db.GetDB().Unscoped().Delete(&hobby).Error

}
