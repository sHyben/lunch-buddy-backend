package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
)

// AreaRepository is a repository for hobbies
// It is used to access the database
// It is a singleton
type AreaRepository struct{}

var areaRepository *AreaRepository

// GetAreaRepository returns the area repository
// It creates a new one if it does not exist
// It returns the singleton instance of the area repository
func GetAreaRepository() *AreaRepository {
	if areaRepository == nil {
		areaRepository = &AreaRepository{}
	}
	return areaRepository
}

// Get returns an area by id
func (r *AreaRepository) Get(id string) (*models.Area, error) {
	var area models.Area
	where := models.Area{}
	stringToUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	where.ID = stringToUuid
	_, err = First(&where, &area, []string{})
	if err != nil {
		return nil, err
	}
	return &area, err
}

// GetByName returns an area by name
func (r *AreaRepository) GetByName(name string) (*models.Area, error) {
	var area models.Area
	where := models.Area{}
	where.Name = name
	_, err := First(&where, &area, []string{})
	if err != nil {
		return nil, err
	}
	return &area, err
}

// All returns all areas
// The hobbies are ordered by id ascending
func (r *AreaRepository) All() (*[]models.Area, error) {
	var areas []models.Area
	err := Find(&models.Area{}, &areas, []string{}, "id asc")
	return &areas, err
}

// Query returns all areas that match the given query
// The query is a user struct with the fields to match
// The fields to match are the fields that are not nil
// The fields to match are the fields that are not empty
// The fields to match are the fields that are not zero
// The fields to match are the fields that are not the zero value for their type
func (r *AreaRepository) Query(q *models.Area) (*[]models.Area, error) {
	var hobbies []models.Area
	err := Find(&q, &hobbies, []string{}, "id asc")
	return &hobbies, err
}

// Add adds an area to the database
func (r *AreaRepository) Add(area *models.Area) error {
	err := Create(&area)
	err = Save(&area)
	return err
}

// Update updates an area in the database
func (r *AreaRepository) Update(area *models.Area) error {
	return db.GetDB().Save(&area).Error
}

// Delete deletes an area from the database
func (r *AreaRepository) Delete(area *models.Area) error {
	return db.GetDB().Unscoped().Delete(&area).Error

}
