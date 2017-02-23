package validators

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHostnameOrIP(t *testing.T) {
	v := &HostnameOrIpValidator{}
	var err error

	_, err = v.Validate("1.2.3.4")
	require.NoError(t, err)

	_, err = v.Validate("192.168.1.1")
	require.NoError(t, err)

	_, err = v.Validate("ip-172-31-43-239")
	require.NoError(t, err)

	_, err = v.Validate("52.32.49.247")
	require.NoError(t, err)

	_, err = v.Validate("bastion")
	require.NoError(t, err)

	_, err = v.Validate("www.google.com")
	require.NoError(t, err)

	_, err = v.Validate("--")
	require.Error(t, err)
}

type Foo struct {
	V []string
}

func TestMulti(t *testing.T) {
	v := NewMultiValidator(&HostnameOrIpValidator{})

	input := []interface{}{"192.168.1.1", "www.google.com"}
	wanted := []string{"192.168.1.1", "www.google.com"}
	output := &Foo{}

	err := v.Unmarshal(nil, nil, input, reflect.ValueOf(output).Elem())
	require.NoError(t, err)

	require.Equal(t, wanted, output.V)
}
