package internal

import (
	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/security"
)

func Save(service string, account string, password string) error {
	goDict := make(map[uintptr]uintptr, 10)
	security.AppendSearchData(goDict, service)
	security.AppendAccountData(goDict, account, password)

	cDict := cfdictionary.Create(goDict)

	return security.Save(cDict)
}
