package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateValid(t *testing.T) {
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
	require.EqualError(t, err, "validation error: invalid hostname: %h: disallowed value '%'")
}
