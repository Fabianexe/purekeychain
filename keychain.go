package main

import (
	"github.com/Fabianexe/purekeychain/internal"
)

// SaveAccount saves account data for a given services
func SaveAccount(service string, login string, password string) error {
	return internal.Save(service, login, password)
}

// GetAccount loads account data for a given service
func GetAccount(service string) (login string, password string, err error) {
	// todo

	return "", "", nil
}
