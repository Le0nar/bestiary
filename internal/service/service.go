package service

import "github.com/Le0nar/bestiary/internal/npc"

type repository interface {
	CreateNpc(dto npc.CreateNpcDto) error
	GetNpcList() ([]npc.Npc, error)
}

type Service struct {
	repository repository
}

func NewService(r repository) *Service {
	return &Service{repository: r}
}

func (s *Service) CreateNpc(dto npc.CreateNpcDto) error {
	return s.repository.CreateNpc(dto)
}

func (s *Service) GetNpcList() ([]npc.Npc, error)  {
	return s.repository.GetNpcList()
}
