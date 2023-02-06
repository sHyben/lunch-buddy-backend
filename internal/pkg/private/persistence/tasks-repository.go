package persistence

import (
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/tasks"
)

type TaskRepository struct{}

// Singleton
// The singleton instance of the TaskRepository
var taskRepository *TaskRepository

// GetTaskRepository returns the singleton instance of the TaskRepository
// If the instance is nil, it will be created
// The instance is created with the default values
//
// Example: If you want to get the singleton instance of the TaskRepository
// you would call this function
func GetTaskRepository() *TaskRepository {
	if taskRepository == nil {
		taskRepository = &TaskRepository{}
	}
	return taskRepository
}

// Get returns a task by id
// The user is eager loaded
func (r *TaskRepository) Get(id string) (*models.Task, error) {
	var task models.Task
	where := models.Task{}
	//where.ID, _ = strconv.ParseUint(id, 10, 64)
	stringToUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	where.ID = stringToUuid //uuid.Must(uuid.Parse(id))
	_, err = First(&where, &task, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &task, err
}

// All returns all tasks
// The tasks are ordered by id ascending
// The user is eager loaded
//
// Example: If you want to find all tasks with the name "test" and the text "test"
// you would create a task struct with the name and text fields set to "test"
// and pass it to this function
func (r *TaskRepository) All() (*[]models.Task, error) {
	var tasks []models.Task
	err := Find(&models.Task{}, &tasks, []string{"User"}, "id asc")
	return &tasks, err
}

// Query returns all tasks that match the given query
// The query is a task struct with the fields to match
// The fields to match are the fields that are not nil
// The fields to match are the fields that are not empty
// The fields to match are the fields that are not zero
// The fields to match are the fields that are not the zero value for their type
// The fields to match are the fields that are not the empty value for their type
//
// Example: If you want to find all tasks with the name "test" and the text "test"
// you would create a task struct with the name and text fields set to "test"
// and pass it to this function
func (r *TaskRepository) Query(q *models.Task) (*[]models.Task, error) {
	var tasks []models.Task
	err := Find(&q, &tasks, []string{"User"}, "id asc")
	return &tasks, err
}

// Add adds a new task to the database
// The user is not eager loaded
func (r *TaskRepository) Add(task *models.Task) error {
	err := Create(&task)
	err = Save(&task)
	return err
}

// Update updates a task in the database
// The user is not eager loaded
func (r *TaskRepository) Update(task *models.Task) error {
	return db.GetDB().Omit("User").Save(&task).Error
}

// Delete deletes a task from the database
// The user is not eager loaded
func (r *TaskRepository) Delete(task *models.Task) error {
	return db.GetDB().Unscoped().Delete(&task).Error
}
