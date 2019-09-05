package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateHostname(t *testing.T) {
	v := &HostNameValidator{}

	_, err := v.Validate("www.google.com")
	require.NoError(t, err)

	_, err = v.Validate("الجزائر.بھارت.فلسطين.台湾.台湾.ভারত")
	require.NoError(t, err)

	// This should be way too long
	_, err = v.Validate("الجزائر.بھارت.فلسطين.台湾.台湾.ভারত.الجزائر.بھارت.فلسطين.台湾.台湾.ভারত.الجزائر.بھارت.فلسطين.台湾.台湾.ভারত.الجزائر.بھارت.فلسطين.台湾.台湾.ভারত")
	require.Error(t, err)

	// No attempt to expand user supplied values
	_, err = v.Validate("%h")
	require.EqualError(t, err, "invalid hostname: %h: disallowed value '%'")
}

func TestValidatePort(t *testing.T) {
	v := &PortValidator{}

	p, err := v.Validate("1")
	require.NoError(t, err)
	require.Equal(t, "1", p)

	p, err = v.Validate("80")
	require.NoError(t, err)
	require.Equal(t, "80", p)

	p, err = v.Validate("443")
	require.NoError(t, err)
	require.Equal(t, "443", p)

	p, err = v.Validate("65535")
	require.NoError(t, err)
	require.Equal(t, "65535", p)

	p, err = v.Validate("")
	require.NoError(t, err)
	require.Equal(t, "", p)

	_, err = v.Validate("0")
	require.Error(t, err)

	_, err = v.Validate("-1")
	require.Error(t, err)

	_, err = v.Validate("65536")
	require.Error(t, err)

	_, err = v.Validate("abc")
	require.Error(t, err)

}

func TestValidateHostAndPort(t *testing.T) {

	v := &HostNameAndPortValidator{}

	hp, err := v.Validate("www.example.com:1")
	require.NoError(t, err)
	require.Equal(t, "www.example.com:1", hp)

	hp, err = v.Validate("example.com:443")
	require.NoError(t, err)
	require.Equal(t, "example.com:443", hp)

	hp, err = v.Validate("localhost:65535")
	require.NoError(t, err)
	require.Equal(t, "localhost:65535", hp)

	hp, err = v.Validate("www.example.com:")
	require.NoError(t, err)
	require.Equal(t, "www.example.com", hp)

	_, err = v.Validate("www.example.com:0")
	require.Error(t, err)

	_, err = v.Validate("www.example.com:-443")
	require.Error(t, err)

	_, err = v.Validate("www.example.com:abc")
	require.Error(t, err)

	_, err = v.Validate("www.example.com:80:443")
	require.Error(t, err)

}
