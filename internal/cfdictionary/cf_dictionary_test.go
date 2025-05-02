package cfdictionary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Fabianexe/purekeychain/internal/cfdictionary"
	"github.com/Fabianexe/purekeychain/internal/cfstring"
)

func TestCFString(t *testing.T) {
	type test struct {
		name string
		m    map[string]string
	}

	tests := []test{
		{
			name: "Simple",
			m:    map[string]string{"k": "v"},
		},
		{
			name: "Multiple 1",
			m:    map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		},
		{
			name: "Multiple 2",
			m:    map[string]string{"k3": "v1", "k2": "v2", "k1": "v3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cValues := make(map[cfstring.CFString]cfstring.CFString)
			for k, v := range test.m {
				cValues[cfstring.Create(k)] = cfstring.Create(v)
			}
			cDict := cfdictionary.Create(cValues)
			result := cfdictionary.ToMap(
				cDict,
				cfstring.CFString.String,
				cfstring.CFString.String,
			)
			assert.Equal(t, test.m, result)
		})
	}
}
