package internal

import (
	"github.com/Fabianexe/purekeychain/internal/cfdata"
	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/cfstring"
	"github.com/Fabianexe/purekeychain/internal/security"
)

func Load(service string) (string, string, error) {
	goDict := security.CreateLoadDict(service)
	cDict := cfdictionary.Create(goDict)
	result, err := security.Load(cDict)
	if err != nil {
		return "", "", err
	}

	goResult := cfdictionary.ToMap(result, cfstring.CFString.String, func(cv uintptr) uintptr { return cv })
	login := cfstring.CFString(goResult["acct"]).String()
	password := cfdata.CFData(goResult["v_Data"]).String()

	return login, password, nil
}
