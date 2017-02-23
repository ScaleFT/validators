package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDNValidaterPositive(t *testing.T) {
	valid := "ou=Hello,ou=World,dc=foo,dc=example,dc=com"
	v := &DNValidator{}

	munged, err := v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)

	valid = "ou=Hello, ou=World, dc=foo, dc=example, dc=com"
	munged, err = v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)

	valid = "ou=Hello World, ou=From, ou=Space, dc=foo, dc=example, dc=com"
	munged, err = v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)
}

func TestDnValidaterNegative(t *testing.T) {
	invalid := "$ou=Hello,ou=World,dc=foo,dc=example,dc=com"
	v := &DNValidator{}

	_, err := v.Validate(invalid)
	require.Error(t, err)

	invalid = "Common_Name=Hello,ou=World,dc=foo,dc=example,dc=com"
	_, err = v.Validate(invalid)
	require.Error(t, err)
}

func TestValidFilterPositive(t *testing.T) {
	v := &LDAPFilterValidator{}

	valid := "(&(foo=bar)(baz=quux))"
	munged, err := v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)

	valid = "(&(foo=bar)(baz=quux))"
	munged, err = v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)

	valid = "(|(foo=bar)(baz=quux))"
	munged, err = v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)

	valid = "(foo=bar)"
	munged, err = v.Validate(valid)
	require.NoError(t, err)
	require.Equal(t, valid, munged)
}

func TestValidFilterNegative(t *testing.T) {
	v := &LDAPFilterValidator{}

	invalid := "(&(foo=bar)(baz=quux)"
	_, err := v.Validate(invalid)
	require.Error(t, err)

	invalid = "((foo=bar)&(baz=quux))"
	_, err = v.Validate(invalid)
	require.Error(t, err)

	invalid = "&(foo=bar)(baz=quux)"
	_, err = v.Validate(invalid)
	require.Error(t, err)

	invalid = "foo=bar"
	_, err = v.Validate(invalid)
	require.Error(t, err)
}
