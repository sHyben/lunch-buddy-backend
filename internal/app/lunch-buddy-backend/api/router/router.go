package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sHyben/lunch-buddy-backend/internal/app/lunch-buddy-backend/api/controllers"
	"github.com/sHyben/lunch-buddy-backend/internal/app/lunch-buddy-backend/api/middlewares"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io"
	"os"
)

// Setup sets up the router
// It returns a gin.Engine
// It is called by main.go
func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())
	//app.Use(middlewares.AuthRequired())

	// Routes
	// ================== Login Routes
	app.POST("/api/login", controllers.Login)
	app.POST("/api/register", controllers.CreateUser)
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ================== User Routes
	app.GET("/api/users", controllers.GetUsers)
	app.GET("/api/users/:id", controllers.GetUserById)
	app.GET("/api/users/username/:username", controllers.GetUserByUsername)
	app.POST("/api/users", controllers.CreateUser)
	app.PUT("/api/users/:id", controllers.UpdateUser)
	app.DELETE("/api/users/:id", controllers.DeleteUser)
	app.POST("/api/users/:id/information", controllers.AddUserInformation)

	app.GET("/api/users/card/:name", controllers.GetUserCard)

	// ================== Hobby Routes
	app.GET("/api/hobbies", controllers.GetHobbies)
	app.GET("/api/hobbies/:id", controllers.GetHobbyById)
	app.POST("/api/hobbies", controllers.CreateHobby)
	app.PUT("/api/hobbies/:id", controllers.UpdateHobby)
	app.DELETE("/api/hobbies/:id", controllers.DeleteHobby)
	// ================== Language Routes
	app.GET("/api/languages", controllers.GetLanguages)
	app.GET("/api/languages/:id", controllers.GetLanguageById)
	app.GET("/api/languages/name/:name", controllers.GetLanguageByName)
	app.POST("/api/languages", controllers.CreateLanguage)
	app.PUT("/api/languages/:id", controllers.UpdateLanguage)
	app.DELETE("/api/languages/:id", controllers.DeleteLanguage)
	// ================== Lunch Routes
	app.GET("/api/lunches", controllers.GetLunches)
	app.GET("/api/lunches/:id", controllers.GetLunchById)
	app.POST("/api/lunches", controllers.CreateLunch)
	app.PUT("/api/lunches/:id", controllers.UpdateLunch)
	app.DELETE("/api/lunches/:id", controllers.DeleteLunch)
	// ================== Area Routes
	app.GET("/api/areas", controllers.GetAreas)
	app.GET("/api/areas/:id", controllers.GetAreaById)
	app.POST("/api/areas", controllers.CreateArea)
	app.PUT("/api/areas/:id", controllers.UpdateArea)
	app.DELETE("/api/areas/:id", controllers.DeleteArea)

	// ================== Tasks Routes
	app.GET("/api/tasks/:id", controllers.GetTaskById)
	app.GET("/api/tasks", controllers.GetTasks)
	app.POST("/api/tasks", controllers.CreateTask)
	app.PUT("/api/tasks/:id", controllers.UpdateTask)
	app.DELETE("/api/tasks/:id", controllers.DeleteTask)

	return app
}
