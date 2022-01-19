package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcryptAdapter(t *testing.T) {
	adapter := NewBcryptAdapter(11)

	password, err := adapter.Generate([]byte("password"))
	assert.NoError(t, err)
	assert.NotEmpty(t, password)

	err = adapter.Compare(password, []byte("password"))
	assert.NoError(t, err)

	err = adapter.Compare(password, []byte("invalid"))
	assert.Error(t, err)
}
