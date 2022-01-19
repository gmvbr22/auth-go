package fingerprint

type FingerprintService interface {
	Generate() (string, string, error)
	Validate(fgp, fgp_hash string) bool
}
