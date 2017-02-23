package validators

import (
	"golang.org/x/net/idna"

	"github.com/russellhaering/jsonmap"
)

// HostNameValidator implements the jsonmap Validator interface.
type HostNameValidator struct{}

// Validate ensures that the hostname given is a valid unicode hostname
// v: the the value to be verified, should be a string
// returns:
//   munged: the, possibly modified, value
//   err: any errors encountered
func (h *HostNameValidator) Validate(v interface{}) (munged interface{}, err error) {
	if val, ok := v.(string); ok {
		if err := RFC5891DNSIgnoreCase(val); err != nil {
			return munged, jsonmap.NewValidationError("invalid hostname: %s", err.Error())
		}
		s, err := idna.ToASCII(val)
		if err != nil {
			return munged, jsonmap.NewValidationError("punycode error: %s", err.Error())
		} else if len(s) > 255 {
			return munged, jsonmap.NewValidationError("punycode translation exceeds maximum length")
		}
		return val, nil
	}
	return munged, jsonmap.NewValidationError("invalid type")
}
