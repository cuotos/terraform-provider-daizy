package daizy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	provider := Provider()
	assert.NoError(t, provider.InternalValidate())
}
