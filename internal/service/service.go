package service

type repository interface {

}

type Service struct {
	repository repository
}

func NewService(r repository) *Service {
	return &Service{repository: r}
}