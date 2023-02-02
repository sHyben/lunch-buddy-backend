package main

import "github.com/sHyben/lunch-buddy-backend/internal/app/lunch-buddy-backend/api"

// @Golang Lunch Buddy API REST
// @version 1.0
// @description API REST in Golang with Gin Framework for application Lunch Buddy

// @contact.name Šimon Hyben
// @contact.url https://shyben.github.io/
// @contact.email hyben.simon@gmail.com

// @license.name MIT
// @license.url

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
