package validators

import (
	"crypto/x509"
	"encoding/pem"

	"github.com/jonboulle/clockwork"
	"github.com/russellhaering/jsonmap"
)

type X509CertificateValidator struct {
	Clock clockwork.Clock
}

func (valid *X509CertificateValidator) Validate(v interface{}) (munged interface{}, err error) {
	if val, ok := v.(string); ok {
		if val == "" {
			return val, nil
		}
		rest := []byte(val)
		var p *pem.Block

		for len(rest) > 0 {
			p, rest = pem.Decode(rest)

			if p == nil {
				return v, jsonmap.NewValidationError("invalid certificate")
			}

			cert, err := x509.ParseCertificate(p.Bytes)
			if err != nil {
				return v, jsonmap.NewValidationError("invalid certificate")
			}
			if valid.Clock.Now().Before(cert.NotBefore) {
				return v, jsonmap.NewValidationError("Certificate is not yet valid (current time is before certificate's NotBefore)")
			}
			if valid.Clock.Now().After(cert.NotAfter) {
				return v, jsonmap.NewValidationError("Certificate has expired (current time is after certificate's NotAfter)")
			}
		}
		return val, nil
	}
	return v, jsonmap.NewValidationError("invalid type")
}
