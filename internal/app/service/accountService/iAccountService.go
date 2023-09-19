package accountService

type IAccountService interface {
	Connect(userId int, phone string) error
	SendMessage(message string, phone string) error
}
