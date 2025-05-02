package internal

import (
	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/security"
)

func Delete(service string) error {
	goDictSearch := make(map[uintptr]uintptr, 10)
	security.AppendSearchData(goDictSearch, service)

	search := cfdictionary.Create(goDictSearch)

	return security.Delete(search)
}
