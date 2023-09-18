package configs

import (
	"github.com/spf13/viper"
	"log"
)

// Setup all configs
func Setup() {

	// Set the names of the configuration files and the paths to search for them.
	viper.SetConfigName("app")
	viper.AddConfigPath("./configs/yamls")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Set up the database configuration
	viper.SetConfigName("db")
	viper.AddConfigPath("./configs/yamls")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading db config file: %s", err)
	}

	// Set up the jwt configuration
	viper.SetConfigName("jwt")
	viper.AddConfigPath("./configs/yamls")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading jwt config file: %s", err)
	}
}
