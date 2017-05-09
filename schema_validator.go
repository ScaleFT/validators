package validators

import (
	"net"
	"strconv"

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

// PortValidator implements the jsonmap Validator interface.
type PortValidator struct{}

// Validate ensures that the port given is valid (empty OR non-zero uint16).
// 0 ports are legal but will fail this validator.
// v: the the value to be verified, should be a string
// returns:
//   munged: the, possibly modified, value
//   err: any errors encountered
func (h *PortValidator) Validate(v interface{}) (munged interface{}, err error) {
	if val, ok := v.(string); ok {
		if val == "" {
			return "", nil
		}

		p, err := strconv.ParseUint(val, 10, 16)
		if err != nil {
			return munged, jsonmap.NewValidationError("invalid port: %s", err.Error())
		}
		if p == 0 {
			return munged, jsonmap.NewValidationError("invalid port: Port cannot be 0")
		}
		return strconv.FormatUint(p, 10), nil
	}
	return munged, jsonmap.NewValidationError("invalid type")
}

// HostNameAndPortValidator implements the jsonmap Validator interface.
type HostNameAndPortValidator struct{}

// Validate ensures that the hostname given is a valid unicode hostname and optional port.
// v: the the value to be verified, should be a string
// returns:
//   munged: the, possibly modified, value
//   err: any errors encountered
func (h *HostNameAndPortValidator) Validate(v interface{}) (munged interface{}, err error) {
	if val, ok := v.(string); ok {
		hostname, port, err := net.SplitHostPort(val)
		if err != nil {
			return munged, jsonmap.NewValidationError("invalid hostname/port: %s", err.Error())
		}

		// Validate hostname
		hv := &HostNameValidator{}
		munged, err := hv.Validate(hostname)
		if err != nil {
			return munged, err
		}
		if hostname, ok = munged.(string); !ok {
			return munged, jsonmap.NewValidationError("invalid hostname type")
		}

		// Validate port
		pv := &PortValidator{}
		munged, err = pv.Validate(port)
		if err != nil {
			return munged, err
		}
		if port, ok = munged.(string); !ok {
			return munged, jsonmap.NewValidationError("invalid port type")
		}

		// Omit empty port
		if port != "" {
			return net.JoinHostPort(hostname, port), nil
		}
		return hostname, nil
	}
	return munged, jsonmap.NewValidationError("invalid type")
}
