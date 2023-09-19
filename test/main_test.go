// accountservicetest.go
package accountService

import (
	"gitlab.com/M.darvish/funtory/internal/app/service/accountService"
	"gitlab.com/M.darvish/funtory/internal/database"
	"gitlab.com/M.darvish/funtory/internal/model"
	"gitlab.com/M.darvish/funtory/internal/model/repository"
	"gitlab.com/M.darvish/funtory/internal/util/security"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

func setupTestDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	return db
}

func createTestUser(t *testing.T, db *gorm.DB, phone string) {
	user := model.User{
		Username: "testuser",
		Password: "testpassword",
		Phone:    phone,
		Session:  "",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
}

func mockEncryptWhatsappToken(userId int, phone string) error {
	_, err := security.EncryptWhatsappToken(userId, phone)
	if err != nil {
		return err
	}
	return nil
}

func queryUserByID(t *testing.T, db *gorm.DB, userId int) model.User {
	var user model.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		t.Fatalf("Failed to query updated user: %v", err)
	}
	return user
}

// TestWhatsAppServiceConnect tests the Connect method of WhatsAppService.
func TestWhatsAppServiceConnect(t *testing.T) {
	db := setupTestDatabase(t)
	defer database.Close()

	userRepo := repository.NewUserImp(db)
	waService := accountService.NewWhatsAppService(userRepo)

	userId := 1
	phone := "1234567890"
	sessionFilePath := "test"

	createTestUser(t, db, phone)

	if err := mockEncryptWhatsappToken(userId, phone); err != nil {
		t.Errorf("Connect returned an error: %v", err)
	}

	if err := waService.Connect(userId, phone); err != nil {
		t.Errorf("Connect returned an error: %v", err)
	}

	updatedUser := queryUserByID(t, db, userId)

	if updatedUser.Session != sessionFilePath {
		t.Errorf("Expected session to be %s, got %s", sessionFilePath, updatedUser.Session)
	}
}

func TestMain(m *testing.M) {
	db := setupTestDatabase(nil)
	defer database.Close()

	db.AutoMigrate(&model.User{})

	// Run the tests.
	exitCode := m.Run()

	os.Remove("test.db")

	os.Exit(exitCode)
}
