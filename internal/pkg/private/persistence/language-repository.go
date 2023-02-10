package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
)

// LanguageRepository is a repository for languages
// It is used to access the database
// It is a singleton
type LanguageRepository struct{}

var languageRepository *LanguageRepository

// GetLanguageRepository returns the language repository
// It creates a new one if it does not exist
// It returns the singleton instance of the language repository
func GetLanguageRepository() *LanguageRepository {
	if languageRepository == nil {
		languageRepository = &LanguageRepository{}
	}
	return languageRepository
}

// Get returns a language by id
func (r *LanguageRepository) Get(id string) (*models.Language, error) {
	var language models.Language
	where := models.Language{}
	stringToUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	where.ID = stringToUuid
	_, err = First(&where, &language, []string{})
	if err != nil {
		return nil, err
	}
	return &language, err
}

// GetByName returns a language by name
func (r *LanguageRepository) GetByName(name string) (*models.Language, error) {
	var language models.Language
	where := models.Language{}
	where.Name = name
	_, err := First(&where, &language, []string{})
	if err != nil {
		return nil, err
	}
	return &language, err
}

// All returns all languages
func (r *LanguageRepository) All() (*[]models.Language, error) {
	var languages []models.Language
	err := Find(&models.Language{}, &languages, []string{}, "id asc")
	return &languages, err
}

// Query returns all languages that match the query
func (r *LanguageRepository) Query(q *models.Language) (*[]models.Language, error) {
	var languages []models.Language
	err := Find(&q, &languages, []string{}, "id asc")
	return &languages, err
}

// Add adds a language to the database
func (r *LanguageRepository) Add(language *models.Language) error {
	err := Create(&language)
	err = Save(&language)
	return err
}

// Update updates a language in the database
func (r *LanguageRepository) Update(language *models.Language) error {
	return db.GetDB().Save(&language).Error
}

// Delete deletes a language from the database
func (r *LanguageRepository) Delete(language *models.Language) error {
	return db.GetDB().Unscoped().Delete(&language).Error
}
