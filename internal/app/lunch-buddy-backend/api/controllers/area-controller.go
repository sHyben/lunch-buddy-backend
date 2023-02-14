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

// GetAreaById godoc
// @Summary Retrieves area based on given ID
// @Description get Area by ID
// @Produce json
// @Param id path integer true "Area ID"
// @Success 200 {object} users.Area
// @Router /api/areas/{id} [get]
// @Security Authorization Token
func GetAreaById(c *gin.Context) {
	s := persistence.GetAreaRepository()
	id := c.Param("id")
	if area, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("area not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, area)
	}
}

// GetAreas godoc
// @Summary Retrieves all areas
// @Description get all areas
// @Produce json
// @Success 200 {object} users.Area
// @Router /api/areas [get]
// @Security Authorization Token
func GetAreas(c *gin.Context) {
	s := persistence.GetAreaRepository()
	var q models.Area
	_ = c.Bind(&q)
	if areas, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("areas not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, areas)
	}
}

// CreateArea godoc
// @Summary Creates an area
// @Description Creates an area
// @Produce json
// @Param area body users.Area true "Area"
// @Success 201 {object} users.Area
// @Router /api/areas [post]
// @Security Authorization Token
func CreateArea(c *gin.Context) {
	s := persistence.GetAreaRepository()
	var areaInput models.Area
	_ = c.BindJSON(&areaInput)
	if err := s.Add(&areaInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, areaInput)
	}

}

// UpdateArea godoc
// @Summary Updates an area
// @Description Updates an area
// @Produce json
// @Param id path integer true "Area ID"
// @Param area body users.Area true "Area"
// @Success 200 {object} users.Area
// @Router /api/areas/{id} [put]
// @Security Authorization Token
func UpdateArea(c *gin.Context) {
	s := persistence.GetAreaRepository()
	id := c.Params.ByName("id")
	var areaInput models.Area
	_ = c.BindJSON(&areaInput)
	if area, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("area not found"))
		log.Println(err)
	} else {
		if areaInput.Name != "" {
			area.Name = areaInput.Name
		} else {
			http_err.NewError(c, http.StatusBadRequest, errors.New("name is required"))
		}
		if err := s.Update(&areaInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, areaInput)
		}
	}
}

// DeleteArea godoc
// @Summary Deletes an area
// @Description Deletes an area
// @Produce json
// @Param id path integer true "Area ID"
// @Success 204
// @Router /api/areas/{id} [delete]
// @Security Authorization Token
func DeleteArea(c *gin.Context) {
	s := persistence.GetAreaRepository()
	id := c.Params.ByName("id")
	/*	var taskInput models.Task
		_ = c.BindJSON(&taskInput)*/
	if area, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("area not found"))
		log.Println(err)
	} else {
		if err := s.Delete(area); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
