package validators

import (
	"fmt"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

// Verify 1-byte runes

func CompareRune(r rune, wanted pVALUE, t *testing.T) {
	v := classifyRune(r)
	if v != wanted {
		pc, file, line, ok := runtime.Caller(2)
		file = path.Base(file)

		var fnc *runtime.Func

		if ok {
			fnc = runtime.FuncForPC(pc)
		}

		if (v == 1 && wanted != 1 && wanted != 5) || (v != 1 && wanted == 1) {
			fmt.Printf("%s: %s: %d: rune: %c (0x%.4X): %#v != %#v\n", file, fnc.Name(), line, r, r, v, wanted)
			t.Fail()
		}
	}
}

func CompareDirect(r rune, wanted pVALUE, t *testing.T) {
	CompareRune(r, wanted, t)
}

func TestClassifyRuneNonJoiner(t *testing.T) {
	CompareDirect(rune(0x200C), pCONTEXTJ, t)
}

func TestClassifyRuneJoiner(t *testing.T) {
	CompareDirect(rune(0x200D), pCONTEXTJ, t)
}

func TestClassifyRuneOtherContext(t *testing.T) {
	CompareDirect(rune(0x05F3), pCONTEXTO, t)
}

func TestRFC5891DNS(t *testing.T) {
	err := RFC5891DNS("www.google.com")
	require.NoError(t, err)

	err = RFC5891DNS("ip-172-31-43-239")
	require.NoError(t, err)

	err = RFC5891DNS("jenkins.c.scale-ft.internal")
	require.NoError(t, err)

	err = RFC5891DNS("www.google.com.")
	require.NoError(t, err)

	err = RFC5891DNS("öbb.at")
	require.NoError(t, err)

	err = RFC5891DNS("ԛәлп.com")
	require.NoError(t, err)
}

func TestRFC5891DNSNegative(t *testing.T) {
	v := "ÖBB.at"
	err := RFC5891DNS(v)
	require.Error(t, err)

	v = "√.com"
	err = RFC5891DNS(v)
	require.Error(t, err)

	v = "Ⱥbby.com"
	err = RFC5891DNS(v)
	require.Error(t, err)

	v = "-leading-hyphen.com"
	err = RFC5891DNS(v)
	require.Error(t, err)

	v = "trailing-hyphen-.com"
	err = RFC5891DNS(v)
	require.Error(t, err)

	v = "no--hyphens.com"
	err = RFC5891DNS(v)
	require.Error(t, err)

	v = "allowed--hyphens.com"
	err = RFC5891DNS(v)
	require.NoError(t, err)

	v = "contains..nested.null.doma.in"
	err = RFC5891DNS(v)
	require.Error(t, err)

	v = "; rm -rf /;.com"
	err = RFC5891DNS(v)
	require.Error(t, err)
}
