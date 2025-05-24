package usecases

type CreateUser interface {
	Execute(name string, number string) error
}
