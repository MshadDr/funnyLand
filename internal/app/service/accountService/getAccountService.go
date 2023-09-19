package accountService

import "errors"

func GetAccountService(accountName string) (IAccountService, error) {
	switch accountName {
	case "whatsapp":
		return &WhatsAppService{}, nil
	default:
		return nil, errors.New("invalid account ID")
	}
}
