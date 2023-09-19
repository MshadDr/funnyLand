package pkg

import (
	"gitlab.com/M.darvish/funtory/internal/model"
	"gorm.io/gorm"
)

// Up The function performs database migration for the tables in GORM.
func Up(db *gorm.DB, modelName string) error {

	modelsName := map[string]interface{}{
		"users": &model.User{},
	}

	if len(modelName) > 0 {
		err := db.AutoMigrate(
			modelsName[modelName],
		)
		if err != nil {
			return err
		}

		return nil
	}

	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Down The function drops the tables from the database using GORM in Go.
func Down(db *gorm.DB, modelName string) error {

	modelsName := map[string]interface{}{
		"users": &model.User{},
	}

	if len(modelName) > 0 {
		err := db.Migrator().DropTable(
			modelsName[modelName],
		)
		if err != nil {
			return err
		}

		return nil
	}

	err := db.Migrator().DropTable(
		&model.User{},
	)
	if err != nil {
		return nil
	}

	return nil
}
