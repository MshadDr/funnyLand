package accountService

import (
	"encoding/gob"
	"fmt"
	"gitlab.com/M.darvish/funtory/internal/model/repository"
	"gitlab.com/M.darvish/funtory/internal/util/security"
	"os"
	"path/filepath"
	"sync"
)

type WhatsAppSession struct {
	UserId      int
	Phone       string
	SessionData map[string]interface{}
}

type IWhatsApp IAccountService

type WhatsAppService struct {
	ur repository.IUser
	mu sync.Mutex
}

func NewWhatsAppService(userRepo repository.IUser) *WhatsAppService {
	return &WhatsAppService{
		ur: userRepo,
	}
}

func saveWhatsAppSession(session *WhatsAppSession, filePath string) error {
	// Create the directory if it doesn't exist.
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(session); err != nil {
		return err
	}

	return nil
}

func (w *WhatsAppService) Connect(userId int, phone string) error {
	// Lock to ensure thread safety when accessing sessions.
	w.mu.Lock()
	defer w.mu.Unlock()

	// Check if a session already exists for this user and phone number.
	userInfo, err := w.ur.GetByUserId(userId)
	if err != nil {
		return err
	}
	session := userInfo.Session
	if session == "nil" {
		// If no session exists, create a new one and save it.
		newSession := &WhatsAppSession{
			UserId:      userId,
			Phone:       userInfo.Phone,
			SessionData: make(map[string]interface{}),
		}

		newSession.SessionData["token"], err = security.EncryptWhatsappToken(userId, userInfo.Phone)
		newSession.SessionData["authenticationData"] = userId

		// Save the new session.
		sessionFilePath := fmt.Sprintf("sessions/%s_session.gob", phone)
		if err := saveWhatsAppSession(newSession, sessionFilePath); err != nil {
			return err
		}

		// Update the session in the database.
		if err := w.ur.UpdateWhatsAppSession(userId, sessionFilePath); err != nil {
			return err
		}
	}
	return nil
}
