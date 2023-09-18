package pkg

import (
	"flag"
	"fmt"
	"gitlab.com/M.darvish/funtory/internal/database"
	"log"
)

func main() {

	// Parse command-line flags
	var migrateUp, migrateDown bool
	var modelName string
	flag.BoolVar(&migrateUp, "up", false, "run migrations up")
	flag.BoolVar(&migrateDown, "down", false, "run migrations down")
	flag.Parse()

	// Connect to the database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// Close the database connection when done
	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println("database close process faced a problem")
			log.Fatalln(err)
		}
	}()

	// Run migrations up or down
	if migrateUp {

		if modelName != "" {
			err = Up(db, modelName)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Migrations up completed successfully.")
			return
		}

		err = Up(db, modelName)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Migrations up completed successfully.")

	} else if migrateDown {

		if modelName != "" {
			err = Down(db, modelName)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Migrations down completed successfully.")
			return
		}

		err = Down(db, modelName)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Migrations down completed successfully.")

	} else {

		fmt.Println("Please specify -up or -down flag.")
	}
}
