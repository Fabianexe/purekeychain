package puregokeychain

import (
	"github.com/Fabianexe/purekeychain/internal"
)

type Service struct {
	name string
}

func New(name string) *Service {
	return &Service{name}
}

// SaveData saves account data for the services
func (s *Service) SaveData(login string, password string) error {
	return internal.Save(s.name, login, password)
}

// LoadData loads account data for the service
func (s *Service) LoadData() (login string, password string, err error) {
	return internal.Load(s.name)
}
