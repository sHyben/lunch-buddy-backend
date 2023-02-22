package db

import (
	"fmt"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/config"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/tasks"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/models/users"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port
	timezone := configuration.Database.TimeZone
	fmt.Println("timezone: ", timezone)

	if driver == "postgres" { // POSTGRES
		dsn := "host=" + host + " port=" + port + " user=" + username + " dbname=" + database + "  sslmode=disable password=" + password + " TimeZone=" + timezone
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("db err: ", err)
		}
	} else if driver == "mysql" { // MYSQL
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}

	// Change this to true if you want to see SQL queries
	//db.LogMode(true)
	dbConfig, _ := db.DB()
	dbConfig.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	dbConfig.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	dbConfig.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)

	/*	db.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
		db.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
		db.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)*/
	DB = db
	migration()
}

// Auto migrate project models
func migration() {
	err := DB.AutoMigrate(&users.User{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&users.UserRole{})
	if err != nil {

		return
	}
	err = DB.AutoMigrate(&tasks.Task{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&users.Hobby{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&users.Language{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&users.Lunch{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&users.Area{})
	if err != nil {
		return
	}
}

func GetDB() *gorm.DB {
	return DB
}
