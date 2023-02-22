package persistence

import (
	"errors"
	"github.com/google/uuid"
	database "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	"gorm.io/gorm"
)

// Create a new record
func Create(value interface{}) error {
	return database.GetDB().Create(value).Error
}

// Save a record
func Save(value interface{}) error {
	return database.GetDB().Save(value).Error
}

// Updates a record
func Updates(where interface{}, value interface{}) error {
	return database.GetDB().Model(where).Updates(value).Error
}

// DeleteByModel delete a record
func DeleteByModel(model interface{}) (count int64, err error) {
	db := database.GetDB().Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByWhere delete a record
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := database.GetDB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByID delete a record
func DeleteByID(model interface{}, id uuid.UUID) (count int64, err error) {
	db := database.GetDB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByIDS delete a record
func DeleteByIDS(model interface{}, ids []uuid.UUID) (count int64, err error) {
	db := database.GetDB().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// FirstByID returns the first record found by id
func FirstByID(out interface{}, id uuid.UUID) (notFound bool, err error) {
	err = database.GetDB().First(out, id).Error
	if err != nil {
		//notFound = gorm.IsRecordNotFoundError(err)
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// First returns the first record found by where
func First(where interface{}, out interface{}, associations []string) (notFound bool, err error) {
	db := database.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	err = db.Where(where).First(out).Error
	if err != nil {
		//notFound = gorm.IsRecordNotFoundError(err)
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// Find returns all records found by where
func Find(where interface{}, out interface{}, associations []string, orders ...string) error {
	db := database.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan returns the first record found by where
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = database.GetDB().Model(model).Where(where).Scan(out).Error
	if err != nil {
		//notFound = gorm.IsRecordNotFoundError(err)
		notFound = errors.Is(err, gorm.ErrRecordNotFound)

	}
	return
}

// ScanList returns all records found by where
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	db := database.GetDB().Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}
