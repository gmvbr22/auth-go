package fingerprint

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type FingerprintSHA256 struct{}

func NewFingerprintSHA256() FingerprintService {
	return &FingerprintSHA256{}
}

func (*FingerprintSHA256) Generate() (fgp string, fgp_hash string, err error) {

	randomFgp := make([]byte, 50)
	if _, err := rand.Read(randomFgp); err != nil {
		return fgp, fgp_hash, err
	}

	fgp = hex.EncodeToString(randomFgp)
	fgp_digest := sha256.Sum256([]byte(fgp))
	fgp_hash = hex.EncodeToString(fgp_digest[:])

	return fgp, fgp_hash, nil
}

func (*FingerprintSHA256) Validate(fgp, fgp_hash string) bool {
	fgp_digest := sha256.Sum256([]byte(fgp))

	return fgp_hash == hex.EncodeToString(fgp_digest[:])
}
