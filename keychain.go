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

// Save saves data for the services
func (s *Service) Save(login string, password string) error {
	return internal.Save(s.name, login, password)
}

// Load loads data for the service
func (s *Service) Load() (login string, password string, err error) {
	return internal.Load(s.name)
}

// Update updates data for the services
func (s *Service) Update(login string, password string) error {
	return internal.Update(s.name, login, password)
}

// Delete deletes data for the services
func (s *Service) Delete() error {
	return internal.Delete(s.name)
}
