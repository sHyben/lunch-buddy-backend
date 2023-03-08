package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
	"gorm.io/gorm"
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
	/*	var userRole models.UserRole
		_, err := First(models.UserRole{UserID: user.ID}, &userRole, []string{})
		//userRole.RoleName = user.Role.RoleName
		err = Save(&userRole)*/
	err := db.GetDB().Omit("Hobbies", "Languages", "Lunch", "Buddies", "Blacklist", "Likes", "Areas").Save(&user).Error
	//user.Role = userRole
	return err
}

// Delete deletes a user from the database
// The role is deleted from the database
// The user is deleted from the database
func (r *UserRepository) Delete(user *models.User) error {
	//err := db.GetDB().Unscoped().Delete(models.UserRole{UserID: user.ID}).Error
	err := db.GetDB().Unscoped().Delete(&user).Error
	return err
}

func (r *UserRepository) ChangeUserArea(user *models.User, area *models.Area) error {
	err := db.GetDB().Model(&user).Association("Areas").Replace(area)
	return err
}

func (r *UserRepository) ChangeUserHobbies(user *models.User, hobbies []models.Hobby) error {
	err := db.GetDB().Model(&user).Association("Hobbies").Replace(hobbies)
	return err
}

func (r *UserRepository) ChangeUserLanguages(user *models.User, languages []models.Language) error {
	err := db.GetDB().Model(&user).Association("Languages").Replace(languages)
	return err
}

func (r *UserRepository) ChangeUserLunch(user *models.User, lunch *models.Lunch) error {
	err := db.GetDB().Model(&user).Association("Lunch").Replace(lunch)
	return err
}

func (r *UserRepository) AddUserBuddies(user *models.User, buddies []models.User) error {
	err := db.GetDB().Model(&user).Association("Buddies").Append(buddies)
	return err
}

func (r *UserRepository) RemoveUserBuddies(user *models.User, buddies []models.User) error {
	err := db.GetDB().Model(&user).Association("Buddies").Delete(buddies)
	return err
}

func (r *UserRepository) AddUserBlacklist(user *models.User, blacklist []models.User) error {
	err := db.GetDB().Model(&user).Association("Blacklist").Append(blacklist)
	return err
}

func (r *UserRepository) RemoveUserBlacklist(user *models.User, blacklist []models.User) error {
	err := db.GetDB().Model(&user).Association("Blacklist").Delete(blacklist)
	return err
}

func (r *UserRepository) AddUserLikes(user *models.User, likes []models.User) error {
	err := db.GetDB().Model(&user).Association("Likes").Append(likes)
	return err
}

func (r *UserRepository) RemoveUserLikes(user *models.User, likes []models.User) error {
	err := db.GetDB().Model(&user).Association("Likes").Delete(likes)
	return err
}

func (r *UserRepository) GetRandomFiveUsers() ([]models.User, error) {
	var users []models.User
	err := db.GetDB().Order(gorm.Expr("random()")).Limit(5).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetRandomFiveUsersThatShareAtLeastOneAreaAndAtLeastOneHobbyAndAtLeastOneLanguageAndHaveSameLunchTime(user *models.User) ([]models.User, error) {
	var users []models.User
	err := db.GetDB().Where("id != ?", user.ID).Where("id NOT IN (?)", user.Blacklist).Where("id NOT IN (?)", user.Buddies).Where("id NOT IN (?)", user.Likes).Where("id IN (?)", user.Areas).Where("id IN (?)", user.Hobbies).Where("id IN (?)", user.Languages).Order(gorm.Expr("random()")).Limit(5).Find(&users).Error
	return users, err
}
func (r *UserRepository) GetRandomFiveUsersWithAssociation() ([]models.User, error) {
	var users []models.User
	err := db.GetDB().Preload("Hobbies").Preload("Languages").Preload("Lunch").Preload("Buddies").Preload("Blacklist").Preload("Likes").Preload("Areas").Order(gorm.Expr("random()")).Limit(5).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserLunch(user *models.User) (*models.Lunch, error) {
	var lunch models.Lunch
	err := db.GetDB().Model(&user).Association("Lunch").Error
	return &lunch, err
}
