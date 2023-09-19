package accountService

type IAccountService interface {
	Connect(userId int, phone string) error
}
