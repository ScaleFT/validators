package validators

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//#################################################
// rudimentary sanity tests
//#################################################
func TestD(t *testing.T) {
	v := GetJoinType(rune(0x0626))
	require.Equal(t, JoinD, v)
}

func TestL(t *testing.T) {
	v := GetJoinType(rune(0xA872))
	require.Equal(t, JoinL, v)
}

func TestR(t *testing.T) {
	v := GetJoinType(rune(0x0622))
	require.Equal(t, JoinR, v)
}

func TestC(t *testing.T) {
	v := GetJoinType(rune(0x0640))
	require.Equal(t, JoinC, v)
}

func TestU_1(t *testing.T) {
	v := GetJoinType(rune(0x0600))
	require.Equal(t, JoinU, v)
}

func TestU_2(t *testing.T) {
	v := GetJoinType(rune(0x041)) // definitely not a join
	require.Equal(t, JoinU, v)
}

// Picked at random from http://www.unicode.org/Public/8.0.0/ucd/extracted/DerivedJoiningType.txt
func TestT_1(t *testing.T) {
	v := GetJoinType(rune(0x00AD))
	require.Equal(t, JoinT, v)
}

func TestT_2(t *testing.T) {
	v := GetJoinType(rune(0x1B3C))
	require.Equal(t, JoinT, v)
}

func TestT_3(t *testing.T) {
	v := GetJoinType(rune(0x116B7))
	require.Equal(t, JoinT, v)
}
