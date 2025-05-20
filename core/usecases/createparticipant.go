package usecases

type CreateParticipant interface {
	Execute(number string, name string, user int) error
}
