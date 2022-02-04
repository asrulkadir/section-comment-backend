package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Database DatabaseConfigurations
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func InitViper() (Configurations, error) {
	var config Configurations
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return config, err
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	host := viper.GetString("host")
	port := viper.GetInt("port")
	dbname := viper.GetString("dbname")

	config.Database.DBName = dbname
	config.Database.DBUser = username
	config.Database.DBPassword = password
	config.Database.DBHost = host
	config.Database.DBPort = fmt.Sprintf("%d", port)

	return config, nil
}
