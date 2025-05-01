package cfstring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Fabianexe/purekeychain/internal/cfstring"
)

func TestCFString(t *testing.T) {
	type test struct {
		name   string
		target string
	}

	tests := []test{
		{
			name:   "Simple",
			target: "simple",
		},
		{
			name:   "Alphanumeric",
			target: "12AFGfsdf2489",
		},
		{
			name:   "UTF8",
			target: "Ææýµ©©",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := cfstring.Create(test.target)
			assert.Equal(t, test.target, result.String())
		})
	}
}
