package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
)

// LunchRepository is a repository for lunches
// It is used to access the database
// It is a singleton
type LunchRepository struct{}

var lunchRepository *LunchRepository

// GetLunchRepository returns the lunch repository
// It creates a new one if it does not exist
// It returns the singleton instance of the lunch repository
func GetLunchRepository() *LunchRepository {
	if lunchRepository == nil {
		lunchRepository = &LunchRepository{}
	}
	return lunchRepository
}

// Get returns a lunch by id
func (r *LunchRepository) Get(id string) (*models.Lunch, error) {
	var lunch models.Lunch
	where := models.Lunch{}
	stringToUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	where.ID = stringToUuid
	_, err = First(&where, &lunch, []string{})
	if err != nil {
		return nil, err
	}
	return &lunch, err
}

// All returns all lunches
func (r *LunchRepository) All() (*[]models.Lunch, error) {
	var lunches []models.Lunch
	err := Find(&models.Lunch{}, &lunches, []string{}, "id asc")
	return &lunches, err
}

// Query returns all lunches that match the query
func (r *LunchRepository) Query(q *models.Lunch) (*[]models.Lunch, error) {
	var lunches []models.Lunch
	err := Find(&q, &lunches, []string{}, "id asc")
	return &lunches, err
}

// Add adds a new lunch to the database
func (r *LunchRepository) Add(lunch *models.Lunch) error {
	err := Create(&lunch)
	err = Save(&lunch)
	return err
}

// Update updates a lunch in the database
func (r *LunchRepository) Update(lunch *models.Lunch) error {
	return db.GetDB().Save(&lunch).Error
}

// Delete deletes a lunch from the database
func (r *LunchRepository) Delete(lunch *models.Lunch) error {
	return db.GetDB().Unscoped().Delete(&lunch).Error
}
