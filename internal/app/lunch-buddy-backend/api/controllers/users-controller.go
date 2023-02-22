package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	models "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/persistence"
	"github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/crypto"
	"github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/http-err"
	"log"
	"net/http"
	"time"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
}

type UserResponse struct {
	Username      string   `json:"username"`
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	Bio           string   `json:"bio"`
	IsSetup       bool     `json:"isSetup"`
	Hobbies       []string `json:"hobbies"`
	Languages     []string `json:"languages"`
	Areas         []string `json:"areas"`
	LunchStart    string   `json:"lunchStart"`
	LunchEnd      string   `json:"lunchEnd"`
	LunchType     string   `json:"lunchType"`
	LunchFood     string   `json:"lunchFood"`
	LunchLocation string   `json:"lunchLocation"`
	Buddies       []string `json:"buddies"`
	Blacklist     []string `json:"blacklist"`
	Likes         []string `json:"likes"`
}

type UserInformation struct {
	AreaNames     []string `json:"areaName"`
	HobbyNames    []string `json:"hobbyNames"`
	LanguageNames []string `json:"languageNames"`
	LunchLocation string   `json:"lunchLocation"`
	LunchTime     string   `json:"lunchTime"`
	LunchType     string   `json:"lunchType"`
	LunchFood     string   `json:"lunchFood"`
	Bio           string   `json:"bio"`
}

// GetUserById godoc
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} users.User
// @Router /api/users/{id} [get]
// @Security Authorization Token
func GetUserById(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Param("id")
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers godoc
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param username query string false "Username"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []users.User
// @Router /api/users [get]
// @Security Authorization Token
func GetUsers(c *gin.Context) {
	s := persistence.GetUserRepository()
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("users not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// CreateUser godoc
// @Summary Creates a new user
// @Description Create User
// @Accept json
// @Produce json
// @Param user body UserInput true "User"
// @Success 201 {object} users.User
// @Router /api/users [post]
// @Security Authorization Token
func CreateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		//Role:      models.UserRole{RoleName: userInput.Role},
	}
	if err := s.Add(&user); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

// UpdateUser godoc
// @Summary Updates an existing user
// @Description Update User
// @Accept json
// @Produce json
// @Param id path integer true "User ID"
// @Param user body UserInput true "User"
// @Success 200 {object} users.User
// @Router /api/users/{id} [put]
// @Security Authorization Token
func UpdateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		user.Username = userInput.Username
		user.Lastname = userInput.Lastname
		user.Firstname = userInput.Firstname
		user.Hash = crypto.HashAndSalt([]byte(userInput.Password))
		//user.Role = models.UserRole{RoleName: userInput.Role}
		if err := s.Update(user); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// DeleteUser godoc
// @Summary Deletes a user
// @Description Delete User
// @Produce json
// @Param id path integer true "User ID"
// @Success 204
// @Router /api/users/{id} [delete]
// @Security Authorization Token
func DeleteUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	/*	var userInput UserInput
		_ = c.BindJSON(&userInput)*/
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if err := s.Delete(user); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}

// GetUserByUsername godoc
// @Summary Retrieves user based on given username
// @Description get User by username
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} users.User
// @Router /api/users/username/{username} [get]
// @Security Authorization Token
func GetUserByUsername(c *gin.Context) {
	s := persistence.GetUserRepository()
	username := c.Param("username")
	if user, err := s.GetByUsername(username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		//c.JSON(http.StatusOK, user)
		userResponse := UserResponse{Username: user.Username, FirstName: user.Firstname, LastName: user.Lastname}
		c.JSON(http.StatusOK, userResponse)
	}
}

func AddUserInformation(c *gin.Context) {
	u := persistence.GetUserRepository()

	id := c.Param("id")
	if user, err := u.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		var userInformation UserInformation
		_ = c.BindJSON(&userInformation)

		if userInformation.Bio != "" {
			user.Bio = userInformation.Bio
			if err := u.Update(user); err != nil {
				http_err.NewError(c, http.StatusNotFound, err)
				log.Println(err)
			}
		}
		AddUserAreas(c, userInformation, user)
		AddUserLunch(c, userInformation, user)
		AddUserHobbies(c, userInformation, user)
		AddUserLanguages(c, userInformation, user)
		c.JSON(http.StatusOK, user)
	}
}

func AddUserAreas(c *gin.Context, userInformation UserInformation, user *models.User) {
	u := persistence.GetUserRepository()
	a := persistence.GetAreaRepository()

	if userInformation.AreaNames != nil && len(userInformation.AreaNames) > 0 {
		for _, areaName := range userInformation.AreaNames {
			if area, err := a.GetByName(areaName); err != nil {
				newArea := models.Area{Name: areaName}
				if err := a.Add(&newArea); err != nil {
					http_err.NewError(c, http.StatusNotFound, err)
					log.Println(err)
				} else {
					if err := u.ChangeUserArea(user, newArea); err != nil {
						http_err.NewError(c, http.StatusNotFound, err)
						log.Println(err)
					}
				}
			} else {
				if err := u.ChangeUserArea(user, *area); err != nil {
					http_err.NewError(c, http.StatusNotFound, err)
					log.Println(err)
				}
			}
		}
	}
}

func AddUserLunch(c *gin.Context, userInformation UserInformation, user *models.User) {
	//u := persistence.GetUserRepository()
	l := persistence.GetLunchRepository()

	if userInformation.LunchLocation != "" && userInformation.LunchTime != "" && userInformation.LunchType != "" && userInformation.LunchFood != "" {
		loc, _ := time.LoadLocation("Europe/Bratislava")
		currentTime := time.Now()
		fmt.Println(currentTime.Format("2006-01-02"))
		//if lunchTime, err := time.Parse("2006-01-01 15:04:05 +01:00", "1970-01-01 "+userInformation.LunchTime+" +01:00"); err != nil {
		if lunchTime, err := time.ParseInLocation("2006-01-02 15:04:05 (MST)", currentTime.Format("2006-01-02")+" "+userInformation.LunchTime+" (CET)", loc); err != nil {
			http_err.NewError(c, http.StatusBadRequest, err)
			log.Println(err)
		} else {
			fmt.Println("Lunch time parsed: ", lunchTime)
			if existingLunch, err := l.Get(user.Lunch.ID.String()); err != nil {
				lunch := models.Lunch{Location: userInformation.LunchLocation, Time: lunchTime, Type: userInformation.LunchType, Food: userInformation.LunchFood, UserID: user.ID}
				if err := l.Add(&lunch); err != nil {
					http_err.NewError(c, http.StatusNotFound, err)
					log.Println(err)
				}
			} else {
				existingLunch.Location = userInformation.LunchLocation
				existingLunch.Time = lunchTime
				existingLunch.Type = userInformation.LunchType
				existingLunch.Food = userInformation.LunchFood
				if err := l.Update(existingLunch); err != nil {
					http_err.NewError(c, http.StatusNotFound, err)
					log.Println(err)
				}
			}

		}
	}
}

func AddUserHobbies(c *gin.Context, userInformation UserInformation, user *models.User) {
	u := persistence.GetUserRepository()
	h := persistence.GetHobbyRepository()

	if userInformation.HobbyNames != nil && len(userInformation.HobbyNames) > 0 {
		log.Println(userInformation.HobbyNames)
		var hobbies []models.Hobby
		for _, hobbyName := range userInformation.HobbyNames {
			if hobby, err := h.GetByName(hobbyName); err != nil {
				hobby = &models.Hobby{Name: hobbyName}
				if err := h.Add(hobby); err != nil {
					http_err.NewError(c, http.StatusNotFound, err)
					log.Println(err)
				} else {
					hobbies = append(hobbies, *hobby)
				}
			} else {
				hobbies = append(hobbies, *hobby)
			}
		}

		if err := u.ChangeUserHobbies(user, hobbies); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		}
	}
}

func AddUserLanguages(c *gin.Context, userInformation UserInformation, user *models.User) {
	u := persistence.GetUserRepository()
	l := persistence.GetLanguageRepository()

	if userInformation.LanguageNames != nil && len(userInformation.LanguageNames) > 0 {
		var languages []models.Language
		for _, languageName := range userInformation.LanguageNames {
			if language, err := l.GetByName(languageName); err != nil {
				language = &models.Language{Name: languageName}
				if err := l.Add(language); err != nil {
					http_err.NewError(c, http.StatusNotFound, err)
					log.Println(err)
				} else {
					languages = append(languages, *language)
				}
			} else {
				languages = append(languages, *language)
			}
		}

		if err := u.ChangeUserLanguages(user, languages); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		}
	}
}

func GetUserCard(c *gin.Context) {
	u := persistence.GetUserRepository()

	name := c.Param("name")
	if user, err := u.GetByUsername(name); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		hobbiesNames := make([]string, len(user.Hobbies))
		for i, hobby := range user.Hobbies {
			hobbiesNames[i] = hobby.Name
		}
		languageNames := make([]string, len(user.Languages))
		for i, language := range user.Languages {
			languageNames[i] = language.Name
		}
		areasNames := make([]string, len(user.Areas))
		for i, area := range user.Areas {
			areasNames[i] = area.Name
		}
		buddiesNames := make([]string, len(user.Buddies))
		for i, buddy := range user.Buddies {
			buddiesNames[i] = buddy.Username
		}
		blackListNames := make([]string, len(user.Blacklist))
		for i, blackList := range user.Blacklist {
			blackListNames[i] = blackList.Username
		}
		likesNames := make([]string, len(user.Likes))
		for i, like := range user.Likes {
			likesNames[i] = like.Username
		}

		c.JSON(http.StatusOK, UserResponse{
			Username:      user.Username,
			FirstName:     user.Firstname,
			LastName:      user.Lastname,
			Bio:           user.Bio,
			IsSetup:       user.IsSetup,
			Hobbies:       hobbiesNames,
			Languages:     languageNames,
			Areas:         areasNames,
			Buddies:       buddiesNames,
			Blacklist:     blackListNames,
			Likes:         likesNames,
			LunchLocation: user.Lunch.Location,
			LunchStart:    user.Lunch.Time.Format("15:04"),
			LunchEnd:      user.Lunch.Time.Add(time.Hour / 2).Format("15:04"),
			LunchType:     user.Lunch.Type,
			LunchFood:     user.Lunch.Food,
		})
	}
}
