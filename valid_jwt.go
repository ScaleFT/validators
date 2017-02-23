package validators

import (
	"strconv"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/russellhaering/jsonmap"
	"gopkg.in/dgrijalva/jwt-go.v2"
)

// JWTValidator provides a mechanism to verify that a string is a valid
// representation of a JWT token, and that it has the required "exp" and "iat"
// claims, has not expired, and is not in the future. The validate method does
// *NOT* verfiy that the token's signature is valid
type JWTValidator struct {
	Clock clockwork.Clock
	Key   []byte
}

// Validate confirms that a JWT can be parsed and isn't expired. it does *NOT*
// verify the signature.
func (valid *JWTValidator) Validate(v interface{}) (munged interface{}, err error) {
	if s, ok := v.(string); ok {
		signingKey := []byte("INVALID KEY")
		if valid.Key != nil {
			signingKey = valid.Key
		}

		var dataError error
		_, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
			if t.Header["typ"] != "JWT" {
				dataError = jsonmap.NewValidationError("Invalid JWT")
				return nil, dataError
			}
			if exp, ok := t.Claims["exp"]; ok {
				ts, err := strconv.ParseInt(exp.(string), 10, 64)
				if err != nil {
					dataError = jsonmap.NewValidationError("JWT expiration invalid")
					return nil, dataError
				}
				expTime := time.Unix(ts, 0)
				if valid.Clock.Now().After(expTime) {
					dataError = jsonmap.NewValidationError("JWT expired")
					return nil, dataError
				}
			}
			if iat, ok := t.Claims["iat"]; ok {
				ts, err := strconv.ParseInt(iat.(string), 10, 64)
				if err != nil {
					dataError = jsonmap.NewValidationError("JWT issued time invalid")
					return nil, dataError
				}
				iatTime := time.Unix(ts, 0)
				if valid.Clock.Now().Before(iatTime) {
					dataError = jsonmap.NewValidationError("JWT is not valid yet")
					return nil, dataError
				}
			}
			return signingKey, nil
		})
		if dataError != nil {
			return v, dataError
		}

		if err != nil {
			// Ignore signature errors if the key is unset
			if valid.Key == nil {
				if e, ok := err.(*jwt.ValidationError); ok {
					if e.Inner == jwt.ErrSignatureInvalid {
						return s, nil
					}
				}
			}
			return v, jsonmap.NewValidationError("Invalid JWT")
		}
		return s, nil
	}
	return v, jsonmap.NewValidationError("invalid type")
}
