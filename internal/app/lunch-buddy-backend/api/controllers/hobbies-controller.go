package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/persistence"
	"github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/http-err"
	"log"
	"net/http"
)

// GetHobbyById godoc
// @Summary Retrieves hobby based on given ID
// @Description get Hobby by ID
// @Produce json
// @Param id path integer true "Hobby ID"
// @Success 200 {object} users.Hobby
// @Router /api/hobbies/{id} [get]
// @Security Authorization Token
func GetHobbyById(c *gin.Context) {
	s := persistence.GetHobbyRepository()
	id := c.Param("id")
	if hobby, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("hobby not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, hobby)
	}
}

// GetHobbies godoc
// @Summary Retrieves all hobbies
// @Description get all hobbies
// @Produce json
// @Success 200 {object} users.Hobby
// @Router /api/hobbies [get]
// @Security Authorization Token
func GetHobbies(c *gin.Context) {
	s := persistence.GetHobbyRepository()
	var q models.Hobby
	_ = c.Bind(&q)
	if hobbies, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("hobbies not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, hobbies)
	}
}

// CreateHobby godoc
// @Summary Creates a hobby
// @Description Creates a hobby
// @Produce json
// @Param hobby body users.Hobby true "Hobby"
// @Success 201 {object} users.Hobby
// @Router /api/hobbies [post]
// @Security Authorization Token
func CreateHobby(c *gin.Context) {
	s := persistence.GetHobbyRepository()
	var hobbyInput models.Hobby
	_ = c.BindJSON(&hobbyInput)
	if err := s.Add(&hobbyInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, hobbyInput)
	}

}

// UpdateHobby godoc
// @Summary Updates a hobby
// @Description Updates a hobby
// @Produce json
// @Param id path integer true "Hobby ID"
// @Param hobby body users.Hobby true "Hobby"
// @Success 200 {object} users.Hobby
// @Router /api/hobbies/{id} [put]
// @Security Authorization Token
func UpdateHobby(c *gin.Context) {
	s := persistence.GetHobbyRepository()
	id := c.Params.ByName("id")
	var hobbyInput models.Hobby
	_ = c.BindJSON(&hobbyInput)
	if hobby, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("hobby not found"))
		log.Println(err)
	} else {
		if hobbyInput.Name != "" {
			hobby.Name = hobbyInput.Name
		} else {
			http_err.NewError(c, http.StatusBadRequest, errors.New("name is required"))
		}
		hobby.Name = hobbyInput.Name
		if err := s.Update(hobby); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, hobby)
		}
	}
}

// DeleteHobby godoc
// @Summary Deletes a hobby
// @Description Deletes a hobby
// @Produce json
// @Param id path integer true "Hobby ID"
// @Success 204
// @Router /api/hobbies/{id} [delete]
// @Security Authorization Token
func DeleteHobby(c *gin.Context) {
	s := persistence.GetHobbyRepository()
	id := c.Params.ByName("id")
	/*	var taskInput models.Task
		_ = c.BindJSON(&taskInput)*/
	if hobby, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("hobby not found"))
		log.Println(err)
	} else {
		if err := s.Delete(hobby); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
