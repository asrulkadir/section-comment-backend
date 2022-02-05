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
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	// viper.SetConfigType("yaml")

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return config, err
	}

	// user := viper.GetString("database.user")
	// password := viper.GetString("database.password")
	// host := viper.GetString("database.host")
	// port := viper.GetInt("database.port")
	// dbname := viper.GetString("database.dbname")

	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	dbname := viper.GetString("DB_NAME")

	config.Database.DBName = dbname
	config.Database.DBUser = user
	config.Database.DBPassword = password
	config.Database.DBHost = host
	config.Database.DBPort = fmt.Sprintf("%d", port)

	return config, nil
}
