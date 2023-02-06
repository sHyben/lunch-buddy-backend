package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/tasks"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/persistence"
	"github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/http-err"
	"log"
	"net/http"
)

// GetTaskById godoc
// @Summary Retrieves task based on given ID
// @Description get Task by ID
// @Produce json
// @Param id path integer true "Task ID"
// @Success 200 {object} tasks.Task
// @Router /api/tasks/{id} [get]
// @Security Authorization Token
func GetTaskById(c *gin.Context) {
	s := persistence.GetTaskRepository()
	id := c.Param("id")
	if task, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

// GetTasks godoc
// @Summary Retrieves tasks based on query
// @Description Get Tasks
// @Produce json
// @Param username query string false "Username"
// @Param taskname query string false "Taskname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []tasks.Task
// @Router /api/tasks [get]
// @Security Authorization Token
// @Tags tasks
// @Accept json
func GetTasks(c *gin.Context) {
	s := persistence.GetTaskRepository()
	var q models.Task
	_ = c.Bind(&q)
	if tasks, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("tasks not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

// CreateTask godoc
// @Summary Creates a new task
// @Description Create Task
// @Produce json
// @Param task body tasks.Task true "Task"
// @Success 201 {object} tasks.Task
// @Router /api/tasks [post]
// @Security Authorization Token
// @Tags tasks
// @Accept json
func CreateTask(c *gin.Context) {
	s := persistence.GetTaskRepository()
	userId := c.Params.ByName("user_id")
	if _, err := s.Get(userId); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		var taskInput models.Task
		_ = c.BindJSON(&taskInput)
		if err := s.Add(&taskInput); err != nil {
			http_err.NewError(c, http.StatusBadRequest, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusCreated, taskInput)
		}
	}
}

// UpdateTask godoc
// @Summary Updates a task
// @Description Update Task
// @Produce json
// @Param id path integer true "Task ID"
// @Param task body tasks.Task true "Task"
// @Success 200 {object} tasks.Task
// @Router /api/tasks/{id} [put]
// @Security Authorization Token
// @Tags tasks
// @Accept json
func UpdateTask(c *gin.Context) {
	s := persistence.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		if err := s.Update(&taskInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, taskInput)
		}
	}
}

// DeleteTask godoc
// @Summary Deletes a task
// @Description Delete Task
// @Produce json
// @Param id path integer true "Task ID"
// @Success 200 {object} tasks.Task
// @Router /api/tasks/{id} [delete]
// @Security Authorization Token
// @Tags tasks
// @Accept json
func DeleteTask(c *gin.Context) {
	s := persistence.GetTaskRepository()
	id := c.Params.ByName("id")
	/*	var taskInput models.Task
		_ = c.BindJSON(&taskInput)*/
	if task, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		if err := s.Delete(task); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
