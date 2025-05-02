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

// SaveData saves data for the services
func (s *Service) SaveData(login string, password string) error {
	return internal.Save(s.name, login, password)
}

// LoadData loads data for the service
func (s *Service) LoadData() (login string, password string, err error) {
	return internal.Load(s.name)
}

// UpdateData updates data for the services
func (s *Service) UpdateData(login string, password string) error {
	return internal.Update(s.name, login, password)
}

// Delete deletes data for the services
func (s *Service) Delete() error {
	return internal.Delete(s.name)
}
