package validators

import (
	"gopkg.in/ldap.v2"
	"regexp"
	"strings"

	"github.com/russellhaering/jsonmap"
)

var LDAPAttrRegex = regexp.MustCompile(`(^$|^[a-zA-Z][[:alnum:]-]+$)`)

type DNValidator struct{}

// Validate confirms that its input is a string, and is a valid IP (IPv4 or IPv6)
func (h *DNValidator) Validate(v interface{}) (munged interface{}, err error) {
	if val, ok := v.(string); ok {
		if val == "" {
			return val, nil
		}
		dn, err := ldap.ParseDN(val)
		if err != nil {
			return munged, jsonmap.NewValidationError("invalid dn")
		}
		if len(dn.RDNs) == 0 {
			return munged, jsonmap.NewValidationError("invalid dn")
		}
		// ParseDN is a bit too forgiving; So I manually verify all the DN type
		// names conform to Spec

		// I'm not using available SYNTAX OIDs for the given type OID to verify
		// values because unfortunately the AD definitions are hugely different
		// from openldap and ds389... so what is a valid value in one is
		// invalid in another. Plus I wouldn't be able to validate custom
		// schema entities
		for _, rdn := range dn.RDNs {
			for _, attr := range rdn.Attributes {
				if !LDAPAttrRegex.MatchString(strings.TrimSpace(attr.Type)) {
					return munged, jsonmap.NewValidationError("invalid dn")
				}
			}
		}
		return val, nil
	}
	return munged, jsonmap.NewValidationError("invalid type")
}

type LDAPFilterValidator struct{}

// Validate ensures that the given value is a string that is a valid LDAP
// filter
func (h *LDAPFilterValidator) Validate(v interface{}) (munged interface{},
	err error) {

	if val, ok := v.(string); ok {
		if val == "" {
			return val, nil
		}

		_, err = ldap.CompileFilter(val)
		if err != nil {
			return munged, jsonmap.NewValidationError("invalid filter")
		}
		return val, nil
	}
	return munged, jsonmap.NewValidationError("invalid type")
}
