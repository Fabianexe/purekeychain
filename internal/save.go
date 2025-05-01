package internal

import (
	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/security"
)

func Save(service string, account string, password string) error {
	goDict := security.CreateSaveDict(service, account, password)
	cDict := cfdictionary.CreateDictionary(goDict)

	return security.Save(cDict)
}
