package fingerprint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFgpSha256(t *testing.T) {

	fingerprint := NewFingerprintSHA256()

	fgp, fgp_hash, err := fingerprint.Generate()

	assert.NoError(t, err)
	assert.NotEmpty(t, fgp)
	assert.NotEmpty(t, fgp_hash)
}

func TestValidateFgpSha256(t *testing.T) {

	fingerprint := NewFingerprintSHA256()
	fgp, fgp_hash, err := fingerprint.Generate()

	assert.NoError(t, err)

	ok := fingerprint.Validate(fgp, fgp_hash)
	assert.True(t, ok)

	ok = fingerprint.Validate("FFFFFFFFFFFFFFFFFF", fgp_hash)
	assert.False(t, ok)
}
