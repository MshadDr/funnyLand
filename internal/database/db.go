package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/M.darvish/funtory/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// NewDatabase The function connects to a Postgres SQL database using configuration values from a configuration File
// and returns a pointer to the database connection and an error.
func NewDatabase() (*gorm.DB, error) {
	configs.Setup()
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	dbname := viper.GetString("db.dbname")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tehran", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}

// The Close function closes the database connection and returns any errors encountered.
func Close() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
