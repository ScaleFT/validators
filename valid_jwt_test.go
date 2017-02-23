package validators

import (
	"strconv"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"
	"gopkg.in/dgrijalva/jwt-go.v2"
)

func TestJWTToken(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = strconv.FormatInt(time.Unix(1000, 0).Unix(), 10)
	token.Claims["iat"] = strconv.FormatInt(time.Unix(0, 0).Unix(), 10)

	valid, err := token.SignedString([]byte("foobar"))
	require.NoError(t, err)

	v := JWTValidator{Clock: clockwork.NewFakeClockAt(time.Unix(100, 0))}
	_, err = v.Validate(valid)
	require.NoError(t, err)
}

func TestJWTTokenNegative(t *testing.T) {
	valid := "This is not a valid token"

	v := JWTValidator{Clock: clockwork.NewFakeClockAt(time.Unix(100, 0))}
	_, err := v.Validate(valid)
	require.Error(t, err)
}

func TestJWTTokenNegExp(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = strconv.FormatInt(time.Unix(1000, 0).Unix(), 10)
	token.Claims["iat"] = strconv.FormatInt(time.Unix(0, 0).Unix(), 10)

	valid, err := token.SignedString([]byte("foobar"))
	require.NoError(t, err)

	v := JWTValidator{Clock: clockwork.NewFakeClockAt(time.Unix(1001, 0))}
	_, err = v.Validate(valid)
	require.Error(t, err)
}

func TestJWTTokenNegIat(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = strconv.FormatInt(time.Unix(1000, 0).Unix(), 10)
	token.Claims["iat"] = strconv.FormatInt(time.Unix(500, 0).Unix(), 10)

	valid, err := token.SignedString([]byte("foobar"))
	require.NoError(t, err)

	v := JWTValidator{Clock: clockwork.NewFakeClockAt(time.Unix(100, 0))}
	_, err = v.Validate(valid)
	require.Error(t, err)
}

func TestJWTTokenSignatureValid(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = strconv.FormatInt(time.Unix(1000, 0).Unix(), 10)
	token.Claims["iat"] = strconv.FormatInt(time.Unix(10, 0).Unix(), 10)

	valid, err := token.SignedString([]byte("foobar"))
	require.NoError(t, err)

	v := JWTValidator{Clock: clockwork.NewFakeClockAt(time.Unix(100, 0)), Key: []byte("foobar")}
	_, err = v.Validate(valid)
	require.NoError(t, err)

	v = JWTValidator{Clock: clockwork.NewFakeClockAt(time.Unix(100, 0)), Key: []byte("incorrect")}
	_, err = v.Validate(valid)
	require.Error(t, err)
}
