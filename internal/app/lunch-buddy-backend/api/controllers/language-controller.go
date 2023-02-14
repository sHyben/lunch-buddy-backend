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

// GetLanguageById godoc
// @Summary Get a language by id
// @Description Get a language by id
// @Param id path string true "Language ID"
// @Success 200 {object} users.Language
// @Router /api/languages/{id} [get]
// @Security Authorization Token
func GetLanguageById(c *gin.Context) {
	s := persistence.GetLanguageRepository()
	id := c.Param("id")
	if language, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("language not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, language)
	}
}

// GetLanguages godoc
// @Summary Get all languages
// @Description Get all languages
// @Success 200 {object} users.Language
// @Router /api/languages [get]
// @Security Authorization Token
func GetLanguages(c *gin.Context) {
	s := persistence.GetLanguageRepository()
	var q models.Language
	_ = c.Bind(&q)
	if languages, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("hobbies not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, languages)
	}
}

// CreateLanguage godoc
// @Summary Create a language
// @Description Create a language
// @Param language body users.Language true "Language"
// @Success 201 {object} users.Language
// @Router /api/languages [post]
// @Security Authorization Token
func CreateLanguage(c *gin.Context) {
	s := persistence.GetLanguageRepository()
	var languageInput models.Language
	_ = c.BindJSON(&languageInput)
	if err := s.Add(&languageInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, languageInput)
	}

}

// UpdateLanguage godoc
// @Summary Update a language
// @Description Update a language
// @Param id path string true "Language ID"
// @Param language body users.Language true "Language"
// @Success 200 {object} users.Language
// @Router /api/languages/{id} [put]
// @Security Authorization Token
func UpdateLanguage(c *gin.Context) {
	s := persistence.GetLanguageRepository()
	id := c.Params.ByName("id")
	var languageInput models.Language
	_ = c.BindJSON(&languageInput)
	if language, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("language not found"))
		log.Println(err)
	} else {
		if languageInput.Name != "" {
			language.Name = languageInput.Name
		} else {
			http_err.NewError(c, http.StatusBadRequest, errors.New("name is required"))
		}

		if err := s.Update(language); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, language)
		}
	}
}

// DeleteLanguage godoc
// @Summary Delete a language
// @Description Delete a language
// @Param id path string true "Language ID"
// @Success 204
// @Router /api/languages/{id} [delete]
// @Security Authorization Token
func DeleteLanguage(c *gin.Context) {
	s := persistence.GetLanguageRepository()
	id := c.Params.ByName("id")
	if language, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("language not found"))
		log.Println(err)
	} else {
		if err := s.Delete(language); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}

func GetLanguageByName(c *gin.Context) {
	s := persistence.GetLanguageRepository()
	name := c.Param("name")
	if language, err := s.GetByName(name); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("language not found"))
		log.Println(err)
	} else {
		//c.JSON(http.StatusOK, user)
		languageResponse := UserResponse{Username: language.Name}
		c.JSON(http.StatusOK, languageResponse)
	}
}
