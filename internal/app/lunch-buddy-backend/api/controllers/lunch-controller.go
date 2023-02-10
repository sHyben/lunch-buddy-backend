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

// GetLunchById godoc
// @Summary Get a lunch by id
// @Description Get a lunch by id
// @Param id path string true "Lunch ID"
// @Success 200 {object} users.Lunch
// @Router /api/lunches/{id} [get]
// @Security Authorization Token
func GetLunchById(c *gin.Context) {
	s := persistence.GetLunchRepository()
	id := c.Param("id")
	if lunch, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("lunch not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, lunch)
	}
}

// GetLunches godoc
// @Summary Get all lunches
// @Description Get all lunches
// @Success 200 {object} users.Lunch
// @Router /api/lunches [get]
// @Security Authorization Token
func GetLunches(c *gin.Context) {
	s := persistence.GetLunchRepository()
	var q models.Lunch
	_ = c.Bind(&q)
	if lunches, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("hobbies not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, lunches)
	}
}

// CreateLunch godoc
// @Summary Create a lunch
// @Description Create a lunch
// @Param lunch body users.Lunch true "Lunch"
// @Success 201 {object} users.Lunch
// @Router /api/lunches [post]
// @Security Authorization Token
func CreateLunch(c *gin.Context) {
	s := persistence.GetLunchRepository()
	var lunchInput models.Lunch
	_ = c.BindJSON(&lunchInput)
	if err := s.Add(&lunchInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, lunchInput)
	}
}

// UpdateLunch godoc
// @Summary Update a lunch
// @Description Update a lunch
// @Param id path string true "Lunch ID"
// @Param lunch body users.Lunch true "Lunch"
// @Success 200 {object} users.Lunch
// @Router /api/lunches/{id} [put]
// @Security Authorization Token
func UpdateLunch(c *gin.Context) {
	s := persistence.GetLunchRepository()
	id := c.Params.ByName("id")
	var lunchInput models.Lunch
	_ = c.BindJSON(&lunchInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("lunch not found"))
		log.Println(err)
	} else {
		if err := s.Update(&lunchInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, lunchInput)
		}
	}
}

// DeleteLunch godoc
// @Summary Delete a lunch
// @Description Delete a lunch
// @Param id path string true "Lunch ID"
// @Success 204
// @Router /api/lunches/{id} [delete]
// @Security Authorization Token
func DeleteLunch(c *gin.Context) {
	s := persistence.GetLunchRepository()
	id := c.Params.ByName("id")
	if lunch, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("lunch not found"))
		log.Println(err)
	} else {
		if err := s.Delete(lunch); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
