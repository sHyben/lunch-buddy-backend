package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/persistence"
	"github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/crypto"
	httpErr "github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/http-err"
	"log"
	"net/http"
)

// LoginInput godoc
// @type LoginInput
// @property username string
// @property password string
// @required
// @in body
// @name loginInput
// @description Login Input
// @example {"username": "admin", "password": "admin"}
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginOutput struct {
	Token     string    `json:"token"`
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Lastname  string    `json:"lastname"`
	Firstname string    `json:"firstname"`
}

// Login godoc
// @Summary Login user
// @Description Login user
// @Produce json
// @Param loginInput body LoginInput true "Login Input"
func Login(c *gin.Context) {
	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		httpErr.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			httpErr.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		loginOutput := LoginOutput{Token: token, ID: user.ID, Username: user.Username, Lastname: user.Lastname, Firstname: user.Firstname}
		c.JSON(http.StatusOK, loginOutput)
	}
}
