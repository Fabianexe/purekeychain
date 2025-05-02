package internal

import (
	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/security"
)

func Update(service string, account string, password string) error {
	goDictSearch := make(map[uintptr]uintptr, 10)
	security.AppendSearchData(goDictSearch, service)
	search := cfdictionary.Create(goDictSearch)

	goDictUpdate := make(map[uintptr]uintptr, 10)
	security.AppendAccountData(goDictUpdate, account, password)
	update := cfdictionary.Create(goDictUpdate)

	return security.Update(search, update)
}
