package interfaces

type UseCase[T any, G any] interface {
	Execute(*T) (*G, error)
}
