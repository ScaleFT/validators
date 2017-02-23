package validators

import (
	"net"

	"github.com/russellhaering/jsonmap"
)

// IPValidator is a type implementing the jsonmap Validator interface for IP Addresses
type IPValidator struct{}

// Validate confirms that its input is a string, and is a valid IP (IPv4 or IPv6)
func (h *IPValidator) Validate(v interface{}) (munged interface{}, err error) {
	if val, ok := v.(string); ok {
		ip := net.ParseIP(val)
		if ip == nil {
			return munged, jsonmap.NewValidationError("invalid ip address")
		}
		return val, nil
	}
	return munged, jsonmap.NewValidationError("invalid type")
}
