package crypto

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	config2 "github.com/sHyben/lunch-buddy-backend/internal/pkg/private/config"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// HashAndSalt hashes a password
// returns a hashed password string
// returns an error if something goes wrong with the hashing
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords compares a hashed password with a plain password
// returns true if they match
// returns false if they don't match
// returns an error if something goes wrong with the hashing
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// CreateToken creates a token
// returns a token string
// returns an error if something goes wrong with the hashing
func CreateToken(username string) (string, error) {
	config := config2.GetConfig()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(config.Server.Secret)) // SECRET
	if err != nil {
		return "token creation error", err
	}
	return token, nil
}

// ValidateToken validates a token
// returns true if the token is valid
// returns false if the token is invalid
// returns an error if something goes wrong with the hashing
func ValidateToken(tokenString string) bool {
	config := config2.GetConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(config.Server.Secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
