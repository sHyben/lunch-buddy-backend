package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

var Config *Configuration

// Configuration is a struct that contains all the configuration data
// for the application
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// DatabaseConfiguration is a struct that contains all the configuration data
// for the database
type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TimeZone     string
}

// ServerConfiguration is a struct that contains all the configuration data
// for the server
type ServerConfiguration struct {
	Port                   string
	Secret                 string
	Mode                   string
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	AccessTokenExpiresIn   time.Duration
	RefreshTokenExpiresIn  time.Duration
	AccessTokenMaxAge      int
	RefreshTokenMaxAge     int
}

// Setup helps you to set up the configuration
// It reads the configuration file and unmarshals it into the configuration struct
// It sets the configuration struct as a global variable
// It panics if something went wrong
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig returns the configuration struct
// It is used to get the configuration from other packages
func GetConfig() *Configuration {
	return Config
}
