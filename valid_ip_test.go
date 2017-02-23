package validators

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIP(t *testing.T) {
	v := &IPValidator{}

	_, err := v.Validate("192.168.1.1")
	require.NoError(t, err)

	_, err = v.Validate("1.2.3.4")
	require.NoError(t, err)

	/* Bug in ParseIP
	_, err = v.Validate("127")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	*/

	_, err = v.Validate("::1")
	require.NoError(t, err)

	/* Bug in ParseIP
	_, err = v.Validate("2130706433")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	*/

	_, err = v.Validate("1234:1234:1::a")
	require.NoError(t, err)

	_, err = v.Validate("1.2.3.4.5")
	require.Error(t, err)

	_, err = v.Validate("256.3.4.5")
	require.Error(t, err)

	_, err = v.Validate("256.1")
	require.Error(t, err)

	_, err = v.Validate("--")
	require.Error(t, err)

	_, err = v.Validate("12001:0000:1234:0000:0000:C1C0:ABCD:0876")
	require.Error(t, err)

	_, err = v.Validate("::1111:2222:3333:4444:5555:6666::")
	require.Error(t, err)

	_, err = v.Validate("1234:1234::1::a")
	require.Error(t, err)

}
