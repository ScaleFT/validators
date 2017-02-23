package validators

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/russellhaering/jsonmap"
)

// IPValidator is a type implementing the jsonmap Validator interface for IP Addresses
type HostnameOrIpValidator struct{}

// Validate confirms that its input is a string, and is a valid IP (IPv4 or IPv6)
func (h *HostnameOrIpValidator) Validate(v interface{}) (munged interface{}, err error) {
	ipv := &IPValidator{}
	if munged, err = ipv.Validate(v); err == nil {
		return munged, err
	}
	hnv := &HostNameValidator{}
	return hnv.Validate(v)
}

// shamelessly stolen from
// https://github.com/russellhaering/jsonmap/blob/master/validators.go but made to
// accept any validator instead of explicitely a StringValidator
type MultiValidator struct {
	elementValidator jsonmap.Validator
}

// Used for StringsArray, which has a "V" field containing []string.
// Optionally can take a string validator to apply to each entry.
func NewMultiValidator(v jsonmap.Validator) jsonmap.TypeMap {
	return &MultiValidator{elementValidator: v}
}

func (ss *MultiValidator) Unmarshal(ctx jsonmap.Context, parent *reflect.Value, partial interface{}, dstValue reflect.Value) error {
	v := dstValue.FieldByName("V")

	underlying := v.Interface()
	if _, ok := underlying.([]string); !ok {
		panic("target field V for MultiValidator is not a []string")
	}

	if partial == nil {
		v.Set(reflect.ValueOf([]string{}))
		return nil
	}

	data, ok := partial.([]interface{})
	if !ok {
		return jsonmap.NewValidationError("expected a list")
	}

	rv := make([]string, len(data))

	for i, dv := range data {
		s, ok := dv.(string)
		if !ok {
			return fmt.Errorf("Error converting %#v to string", dv)
		}

		if ss.elementValidator != nil {
			munged, err := ss.elementValidator.Validate(s)
			if err != nil {
				return err
			}
			s = munged.(string)
		}

		rv[i] = s
	}

	v.Set(reflect.ValueOf(rv))

	return nil
}

func (s *MultiValidator) Marshal(ctx jsonmap.Context, parent *reflect.Value, src reflect.Value) (json.Marshaler, error) {
	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}

	v := src.FieldByName("V")

	data, err := json.Marshal(v.Interface())
	if err != nil {
		return nil, err
	}

	return jsonmap.RawMessage{data}, nil
}
