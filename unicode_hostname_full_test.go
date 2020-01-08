package validators

/*
generated with:

curl -o - https://tools.ietf.org/html/rfc5892 2>/dev/null | awk '/^[0-9A-F]{4}\ .*(DISALLOWED|PVALID|CONTEXT.|UNASSIGNED)/ {print "func Test" $1 "(t *testing.T) {\n\tValidateRune(t, 0x"$1", p"$3")\n}\n"} /^[0-9A-F]{5,8}\ .*(DISALLOWED|PVALID|CONTEXT.|UNASSIGNED)/ {print "func Test" $1 "(t *testing.T) {\n\tValidateRune(t, 0x"$1", p"$3")\n}\n"} /^[0-9A-F]{4}..[0-9A-F]{4}.*(DISALLOWED|PVALID|CONTEXT.|UNASSIGNED)/ {n1=substr($1,0,index($1,"..")-1); n2=substr($1,index($1,"..")+2); print "func Test" n1 "_" n2"(t *testing.T) {\n\tValidateRuneRange(t, 0x" n1 ", 0x" n2 ", p" $3 ")\n}\n"} /^[0-9A-F]{5,8}..[0-9A-F]{5,8}.*(DISALLOWED|PVALID|CONTEXT.|UNASSIGNED)/ {n1=substr($1,0,index($1,"..")-1); n2=substr($1,index($1,"..")+2, index($1,";") - (index($1,"..")+2)); print "func Test" n1 "_" n2"(t *testing.T) {\n\tValidateRuneRange(t, 0x" n1 ", 0x" n2 ", p" $2 ")\n}\n"}' >> THIS_FILE


Because the RFC example used to generate this was built from unicode 5.2, quite
a lot has changed, and so many tests now fail.

This is precisely why the RFC outlines the algorithms to use to identify valid
labels rather than a fixed list. However, that means that quite a few tests
here in fact fail that would otherwise pass.

As such, this test should be used as a guideline, and not requried to pass. As
such, I'm setting the build-directive on this to disable it under normal
circumstanes (as it would break the build otherwise.
*/

import (
	"testing"
)

const (
	DisableFullTests = false
)

var (
	empty                     struct{}
	expectedFailuresUnicodeV8 = map[rune]struct{}{
		0x0CF1: empty,
		0x0CF2: empty,
		0x19DA: empty,
		0x1F80: empty,
		0x1F81: empty,
		0x1F82: empty,
		0x1F83: empty,
		0x1F84: empty,
		0x1F85: empty,
		0x1F86: empty,
		0x1F87: empty,
		0x1F90: empty,
		0x1F91: empty,
		0x1F92: empty,
		0x1F93: empty,
		0x1F94: empty,
		0x1F95: empty,
		0x1F96: empty,
		0x1F97: empty,
		0x1FA0: empty,
		0x1FA1: empty,
		0x1FA2: empty,
		0x1FA3: empty,
		0x1FA4: empty,
		0x1FA5: empty,
		0x1FA6: empty,
		0x1FA7: empty,
		0x1FB2: empty,
		0x1FB3: empty,
		0x1FB4: empty,
		0x1FB7: empty,
		0x1FC2: empty,
		0x1FC3: empty,
		0x1FC4: empty,
		0x1FC7: empty,
		0x1FF2: empty,
		0x1FF3: empty,
		0x1FF4: empty,
		0x1FF7: empty,
	}
)

func ValidateRune(t *testing.T, r rune, wanted pVALUE) {
	_, skip := expectedFailuresUnicodeV8[r]
	if !DisableFullTests || !skip {
		CompareRune(r, wanted, t)
	}
}

func ValidateRuneRange(t *testing.T, first, second rune, wanted pVALUE) {
	if !DisableFullTests {
		for i := first; i <= second; i++ {
			if _, skip := expectedFailuresUnicodeV8[i]; skip {
				continue
			}
			CompareRune(rune(i), wanted, t)
		}
	}
}

//-------------------CUT HERE---------------------------
func Test0000_002C(t *testing.T) {
	ValidateRuneRange(t, 0x0000, 0x002C, pDISALLOWED)
}

func Test002D(t *testing.T) {
	ValidateRune(t, 0x002D, pPVALID)
}

func Test002E_002F(t *testing.T) {
	ValidateRuneRange(t, 0x002E, 0x002F, pDISALLOWED)
}

func Test0030_0039(t *testing.T) {
	ValidateRuneRange(t, 0x0030, 0x0039, pPVALID)
}

func Test003A_005E(t *testing.T) {
	ValidateRuneRange(t, 0x003A, 0x005E, pDISALLOWED)
}

func Test005F_005F(t *testing.T) {
	// Exception to add support for underscores in hostnames -- not from a generated test case
	ValidateRuneRange(t, 0x005F, 0x005F, pPVALID)
}

func Test0060_0060(t *testing.T) {
	ValidateRuneRange(t, 0x0060, 0x0060, pDISALLOWED)
}

func Test0061_007A(t *testing.T) {
	ValidateRuneRange(t, 0x0061, 0x007A, pPVALID)
}

func Test007B_00B6(t *testing.T) {
	ValidateRuneRange(t, 0x007B, 0x00B6, pDISALLOWED)
}

func Test00B7(t *testing.T) {
	ValidateRune(t, 0x00B7, pCONTEXTO)
}

func Test00B8_00DE(t *testing.T) {
	ValidateRuneRange(t, 0x00B8, 0x00DE, pDISALLOWED)
}

func Test00DF_00F6(t *testing.T) {
	ValidateRuneRange(t, 0x00DF, 0x00F6, pPVALID)
}

func Test00F7(t *testing.T) {
	ValidateRune(t, 0x00F7, pDISALLOWED)
}

func Test00F8_00FF(t *testing.T) {
	ValidateRuneRange(t, 0x00F8, 0x00FF, pPVALID)
}

func Test010A(t *testing.T) {
	ValidateRune(t, 0x010A, pDISALLOWED)
}

func Test010B(t *testing.T) {
	ValidateRune(t, 0x010B, pPVALID)
}

func Test010C(t *testing.T) {
	ValidateRune(t, 0x010C, pDISALLOWED)
}

func Test010D(t *testing.T) {
	ValidateRune(t, 0x010D, pPVALID)
}

func Test010E(t *testing.T) {
	ValidateRune(t, 0x010E, pDISALLOWED)
}

func Test010F(t *testing.T) {
	ValidateRune(t, 0x010F, pPVALID)
}

func Test011A(t *testing.T) {
	ValidateRune(t, 0x011A, pDISALLOWED)
}

func Test011B(t *testing.T) {
	ValidateRune(t, 0x011B, pPVALID)
}

func Test011C(t *testing.T) {
	ValidateRune(t, 0x011C, pDISALLOWED)
}

func Test011D(t *testing.T) {
	ValidateRune(t, 0x011D, pPVALID)
}

func Test011E(t *testing.T) {
	ValidateRune(t, 0x011E, pDISALLOWED)
}

func Test011F(t *testing.T) {
	ValidateRune(t, 0x011F, pPVALID)
}

func Test012A(t *testing.T) {
	ValidateRune(t, 0x012A, pDISALLOWED)
}

func Test012B(t *testing.T) {
	ValidateRune(t, 0x012B, pPVALID)
}

func Test012C(t *testing.T) {
	ValidateRune(t, 0x012C, pDISALLOWED)
}

func Test012D(t *testing.T) {
	ValidateRune(t, 0x012D, pPVALID)
}

func Test012E(t *testing.T) {
	ValidateRune(t, 0x012E, pDISALLOWED)
}

func Test012F(t *testing.T) {
	ValidateRune(t, 0x012F, pPVALID)
}

func Test0132_0134(t *testing.T) {
	ValidateRuneRange(t, 0x0132, 0x0134, pDISALLOWED)
}

func Test0137_0138(t *testing.T) {
	ValidateRuneRange(t, 0x0137, 0x0138, pPVALID)
}

func Test013A(t *testing.T) {
	ValidateRune(t, 0x013A, pPVALID)
}

func Test013B(t *testing.T) {
	ValidateRune(t, 0x013B, pDISALLOWED)
}

func Test013C(t *testing.T) {
	ValidateRune(t, 0x013C, pPVALID)
}

func Test013D(t *testing.T) {
	ValidateRune(t, 0x013D, pDISALLOWED)
}

func Test013E(t *testing.T) {
	ValidateRune(t, 0x013E, pPVALID)
}

func Test013F_0141(t *testing.T) {
	ValidateRuneRange(t, 0x013F, 0x0141, pDISALLOWED)
}

func Test0149_014A(t *testing.T) {
	ValidateRuneRange(t, 0x0149, 0x014A, pDISALLOWED)
}

func Test014B(t *testing.T) {
	ValidateRune(t, 0x014B, pPVALID)
}

func Test014C(t *testing.T) {
	ValidateRune(t, 0x014C, pDISALLOWED)
}

func Test014D(t *testing.T) {
	ValidateRune(t, 0x014D, pPVALID)
}

func Test014E(t *testing.T) {
	ValidateRune(t, 0x014E, pDISALLOWED)
}

func Test014F(t *testing.T) {
	ValidateRune(t, 0x014F, pPVALID)
}

func Test015A(t *testing.T) {
	ValidateRune(t, 0x015A, pDISALLOWED)
}

func Test015B(t *testing.T) {
	ValidateRune(t, 0x015B, pPVALID)
}

func Test015C(t *testing.T) {
	ValidateRune(t, 0x015C, pDISALLOWED)
}

func Test015D(t *testing.T) {
	ValidateRune(t, 0x015D, pPVALID)
}

func Test015E(t *testing.T) {
	ValidateRune(t, 0x015E, pDISALLOWED)
}

func Test015F(t *testing.T) {
	ValidateRune(t, 0x015F, pPVALID)
}

func Test016A(t *testing.T) {
	ValidateRune(t, 0x016A, pDISALLOWED)
}

func Test016B(t *testing.T) {
	ValidateRune(t, 0x016B, pPVALID)
}

func Test016C(t *testing.T) {
	ValidateRune(t, 0x016C, pDISALLOWED)
}

func Test016D(t *testing.T) {
	ValidateRune(t, 0x016D, pPVALID)
}

func Test016E(t *testing.T) {
	ValidateRune(t, 0x016E, pDISALLOWED)
}

func Test016F(t *testing.T) {
	ValidateRune(t, 0x016F, pPVALID)
}

func Test0178_0179(t *testing.T) {
	ValidateRuneRange(t, 0x0178, 0x0179, pDISALLOWED)
}

func Test017A(t *testing.T) {
	ValidateRune(t, 0x017A, pPVALID)
}

func Test017B(t *testing.T) {
	ValidateRune(t, 0x017B, pDISALLOWED)
}

func Test017C(t *testing.T) {
	ValidateRune(t, 0x017C, pPVALID)
}

func Test017D(t *testing.T) {
	ValidateRune(t, 0x017D, pDISALLOWED)
}

func Test017E(t *testing.T) {
	ValidateRune(t, 0x017E, pPVALID)
}

func Test017F(t *testing.T) {
	ValidateRune(t, 0x017F, pDISALLOWED)
}

func Test0181_0182(t *testing.T) {
	ValidateRuneRange(t, 0x0181, 0x0182, pDISALLOWED)
}

func Test0186_0187(t *testing.T) {
	ValidateRuneRange(t, 0x0186, 0x0187, pDISALLOWED)
}

func Test0189_018B(t *testing.T) {
	ValidateRuneRange(t, 0x0189, 0x018B, pDISALLOWED)
}

func Test018C_018D(t *testing.T) {
	ValidateRuneRange(t, 0x018C, 0x018D, pPVALID)
}

func Test018E_0191(t *testing.T) {
	ValidateRuneRange(t, 0x018E, 0x0191, pDISALLOWED)
}

func Test0193_0194(t *testing.T) {
	ValidateRuneRange(t, 0x0193, 0x0194, pDISALLOWED)
}

func Test0196_0198(t *testing.T) {
	ValidateRuneRange(t, 0x0196, 0x0198, pDISALLOWED)
}

func Test0199_019B(t *testing.T) {
	ValidateRuneRange(t, 0x0199, 0x019B, pPVALID)
}

func Test019C_019D(t *testing.T) {
	ValidateRuneRange(t, 0x019C, 0x019D, pDISALLOWED)
}

func Test019E(t *testing.T) {
	ValidateRune(t, 0x019E, pPVALID)
}

func Test019F_01A0(t *testing.T) {
	ValidateRuneRange(t, 0x019F, 0x01A0, pDISALLOWED)
}

func Test01A1(t *testing.T) {
	ValidateRune(t, 0x01A1, pPVALID)
}

func Test01A2(t *testing.T) {
	ValidateRune(t, 0x01A2, pDISALLOWED)
}

func Test01A3(t *testing.T) {
	ValidateRune(t, 0x01A3, pPVALID)
}

func Test01A4(t *testing.T) {
	ValidateRune(t, 0x01A4, pDISALLOWED)
}

func Test01A5(t *testing.T) {
	ValidateRune(t, 0x01A5, pPVALID)
}

func Test01A6_01A7(t *testing.T) {
	ValidateRuneRange(t, 0x01A6, 0x01A7, pDISALLOWED)
}

func Test01A8(t *testing.T) {
	ValidateRune(t, 0x01A8, pPVALID)
}

func Test01A9(t *testing.T) {
	ValidateRune(t, 0x01A9, pDISALLOWED)
}

func Test01AA_01AB(t *testing.T) {
	ValidateRuneRange(t, 0x01AA, 0x01AB, pPVALID)
}

func Test01AC(t *testing.T) {
	ValidateRune(t, 0x01AC, pDISALLOWED)
}

func Test01AD(t *testing.T) {
	ValidateRune(t, 0x01AD, pPVALID)
}

func Test01AE_01AF(t *testing.T) {
	ValidateRuneRange(t, 0x01AE, 0x01AF, pDISALLOWED)
}

func Test01B0(t *testing.T) {
	ValidateRune(t, 0x01B0, pPVALID)
}

func Test01B1_01B3(t *testing.T) {
	ValidateRuneRange(t, 0x01B1, 0x01B3, pDISALLOWED)
}

func Test01B4(t *testing.T) {
	ValidateRune(t, 0x01B4, pPVALID)
}

func Test01B5(t *testing.T) {
	ValidateRune(t, 0x01B5, pDISALLOWED)
}

func Test01B6(t *testing.T) {
	ValidateRune(t, 0x01B6, pPVALID)
}

func Test01B7_01B8(t *testing.T) {
	ValidateRuneRange(t, 0x01B7, 0x01B8, pDISALLOWED)
}

func Test01B9_01BB(t *testing.T) {
	ValidateRuneRange(t, 0x01B9, 0x01BB, pPVALID)
}

func Test01BC(t *testing.T) {
	ValidateRune(t, 0x01BC, pDISALLOWED)
}

func Test01BD_01C3(t *testing.T) {
	ValidateRuneRange(t, 0x01BD, 0x01C3, pPVALID)
}

func Test01C4_01CD(t *testing.T) {
	ValidateRuneRange(t, 0x01C4, 0x01CD, pDISALLOWED)
}

func Test01CE(t *testing.T) {
	ValidateRune(t, 0x01CE, pPVALID)
}

func Test01CF(t *testing.T) {
	ValidateRune(t, 0x01CF, pDISALLOWED)
}

func Test01D0(t *testing.T) {
	ValidateRune(t, 0x01D0, pPVALID)
}

func Test01D1(t *testing.T) {
	ValidateRune(t, 0x01D1, pDISALLOWED)
}

func Test01D2(t *testing.T) {
	ValidateRune(t, 0x01D2, pPVALID)
}

func Test01D3(t *testing.T) {
	ValidateRune(t, 0x01D3, pDISALLOWED)
}

func Test01D4(t *testing.T) {
	ValidateRune(t, 0x01D4, pPVALID)
}

func Test01D5(t *testing.T) {
	ValidateRune(t, 0x01D5, pDISALLOWED)
}

func Test01D6(t *testing.T) {
	ValidateRune(t, 0x01D6, pPVALID)
}

func Test01D7(t *testing.T) {
	ValidateRune(t, 0x01D7, pDISALLOWED)
}

func Test01D8(t *testing.T) {
	ValidateRune(t, 0x01D8, pPVALID)
}

func Test01D9(t *testing.T) {
	ValidateRune(t, 0x01D9, pDISALLOWED)
}

func Test01DA(t *testing.T) {
	ValidateRune(t, 0x01DA, pPVALID)
}

func Test01DB(t *testing.T) {
	ValidateRune(t, 0x01DB, pDISALLOWED)
}

func Test01DC_01DD(t *testing.T) {
	ValidateRuneRange(t, 0x01DC, 0x01DD, pPVALID)
}

func Test01DE(t *testing.T) {
	ValidateRune(t, 0x01DE, pDISALLOWED)
}

func Test01DF(t *testing.T) {
	ValidateRune(t, 0x01DF, pPVALID)
}

func Test01E0(t *testing.T) {
	ValidateRune(t, 0x01E0, pDISALLOWED)
}

func Test01E1(t *testing.T) {
	ValidateRune(t, 0x01E1, pPVALID)
}

func Test01E2(t *testing.T) {
	ValidateRune(t, 0x01E2, pDISALLOWED)
}

func Test01E3(t *testing.T) {
	ValidateRune(t, 0x01E3, pPVALID)
}

func Test01E4(t *testing.T) {
	ValidateRune(t, 0x01E4, pDISALLOWED)
}

func Test01E5(t *testing.T) {
	ValidateRune(t, 0x01E5, pPVALID)
}

func Test01E6(t *testing.T) {
	ValidateRune(t, 0x01E6, pDISALLOWED)
}

func Test01E7(t *testing.T) {
	ValidateRune(t, 0x01E7, pPVALID)
}

func Test01E8(t *testing.T) {
	ValidateRune(t, 0x01E8, pDISALLOWED)
}

func Test01E9(t *testing.T) {
	ValidateRune(t, 0x01E9, pPVALID)
}

func Test01EA(t *testing.T) {
	ValidateRune(t, 0x01EA, pDISALLOWED)
}

func Test01EB(t *testing.T) {
	ValidateRune(t, 0x01EB, pPVALID)
}

func Test01EC(t *testing.T) {
	ValidateRune(t, 0x01EC, pDISALLOWED)
}

func Test01ED(t *testing.T) {
	ValidateRune(t, 0x01ED, pPVALID)
}

func Test01EE(t *testing.T) {
	ValidateRune(t, 0x01EE, pDISALLOWED)
}

func Test01EF_01F0(t *testing.T) {
	ValidateRuneRange(t, 0x01EF, 0x01F0, pPVALID)
}

func Test01F1_01F4(t *testing.T) {
	ValidateRuneRange(t, 0x01F1, 0x01F4, pDISALLOWED)
}

func Test01F5(t *testing.T) {
	ValidateRune(t, 0x01F5, pPVALID)
}

func Test01F6_01F8(t *testing.T) {
	ValidateRuneRange(t, 0x01F6, 0x01F8, pDISALLOWED)
}

func Test01F9(t *testing.T) {
	ValidateRune(t, 0x01F9, pPVALID)
}

func Test01FA(t *testing.T) {
	ValidateRune(t, 0x01FA, pDISALLOWED)
}

func Test01FB(t *testing.T) {
	ValidateRune(t, 0x01FB, pPVALID)
}

func Test01FC(t *testing.T) {
	ValidateRune(t, 0x01FC, pDISALLOWED)
}

func Test01FD(t *testing.T) {
	ValidateRune(t, 0x01FD, pPVALID)
}

func Test01FE(t *testing.T) {
	ValidateRune(t, 0x01FE, pDISALLOWED)
}

func Test01FF(t *testing.T) {
	ValidateRune(t, 0x01FF, pPVALID)
}

func Test020A(t *testing.T) {
	ValidateRune(t, 0x020A, pDISALLOWED)
}

func Test020B(t *testing.T) {
	ValidateRune(t, 0x020B, pPVALID)
}

func Test020C(t *testing.T) {
	ValidateRune(t, 0x020C, pDISALLOWED)
}

func Test020D(t *testing.T) {
	ValidateRune(t, 0x020D, pPVALID)
}

func Test020E(t *testing.T) {
	ValidateRune(t, 0x020E, pDISALLOWED)
}

func Test020F(t *testing.T) {
	ValidateRune(t, 0x020F, pPVALID)
}

func Test021A(t *testing.T) {
	ValidateRune(t, 0x021A, pDISALLOWED)
}

func Test021B(t *testing.T) {
	ValidateRune(t, 0x021B, pPVALID)
}

func Test021C(t *testing.T) {
	ValidateRune(t, 0x021C, pDISALLOWED)
}

func Test021D(t *testing.T) {
	ValidateRune(t, 0x021D, pPVALID)
}

func Test021E(t *testing.T) {
	ValidateRune(t, 0x021E, pDISALLOWED)
}

func Test021F(t *testing.T) {
	ValidateRune(t, 0x021F, pPVALID)
}

func Test022A(t *testing.T) {
	ValidateRune(t, 0x022A, pDISALLOWED)
}

func Test022B(t *testing.T) {
	ValidateRune(t, 0x022B, pPVALID)
}

func Test022C(t *testing.T) {
	ValidateRune(t, 0x022C, pDISALLOWED)
}

func Test022D(t *testing.T) {
	ValidateRune(t, 0x022D, pPVALID)
}

func Test022E(t *testing.T) {
	ValidateRune(t, 0x022E, pDISALLOWED)
}

func Test022F(t *testing.T) {
	ValidateRune(t, 0x022F, pPVALID)
}

func Test0233_0239(t *testing.T) {
	ValidateRuneRange(t, 0x0233, 0x0239, pPVALID)
}

func Test023A_023B(t *testing.T) {
	ValidateRuneRange(t, 0x023A, 0x023B, pDISALLOWED)
}

func Test023C(t *testing.T) {
	ValidateRune(t, 0x023C, pPVALID)
}

func Test023D_023E(t *testing.T) {
	ValidateRuneRange(t, 0x023D, 0x023E, pDISALLOWED)
}

func Test023F_0240(t *testing.T) {
	ValidateRuneRange(t, 0x023F, 0x0240, pPVALID)
}

func Test0243_0246(t *testing.T) {
	ValidateRuneRange(t, 0x0243, 0x0246, pDISALLOWED)
}

func Test024A(t *testing.T) {
	ValidateRune(t, 0x024A, pDISALLOWED)
}

func Test024B(t *testing.T) {
	ValidateRune(t, 0x024B, pPVALID)
}

func Test024C(t *testing.T) {
	ValidateRune(t, 0x024C, pDISALLOWED)
}

func Test024D(t *testing.T) {
	ValidateRune(t, 0x024D, pPVALID)
}

func Test024E(t *testing.T) {
	ValidateRune(t, 0x024E, pDISALLOWED)
}

func Test024F_02AF(t *testing.T) {
	ValidateRuneRange(t, 0x024F, 0x02AF, pPVALID)
}

func Test02B0_02B8(t *testing.T) {
	ValidateRuneRange(t, 0x02B0, 0x02B8, pDISALLOWED)
}

func Test02B9_02C1(t *testing.T) {
	ValidateRuneRange(t, 0x02B9, 0x02C1, pPVALID)
}

func Test02C2_02C5(t *testing.T) {
	ValidateRuneRange(t, 0x02C2, 0x02C5, pDISALLOWED)
}

func Test02C6_02D1(t *testing.T) {
	ValidateRuneRange(t, 0x02C6, 0x02D1, pPVALID)
}

func Test02D2_02EB(t *testing.T) {
	ValidateRuneRange(t, 0x02D2, 0x02EB, pDISALLOWED)
}

func Test02EC(t *testing.T) {
	ValidateRune(t, 0x02EC, pPVALID)
}

func Test02ED(t *testing.T) {
	ValidateRune(t, 0x02ED, pDISALLOWED)
}

func Test02EE(t *testing.T) {
	ValidateRune(t, 0x02EE, pPVALID)
}

func Test02EF_02FF(t *testing.T) {
	ValidateRuneRange(t, 0x02EF, 0x02FF, pDISALLOWED)
}

func Test0300_033F(t *testing.T) {
	ValidateRuneRange(t, 0x0300, 0x033F, pPVALID)
}

func Test0340_0341(t *testing.T) {
	ValidateRuneRange(t, 0x0340, 0x0341, pDISALLOWED)
}

func Test0343_0345(t *testing.T) {
	ValidateRuneRange(t, 0x0343, 0x0345, pDISALLOWED)
}

func Test0346_034E(t *testing.T) {
	ValidateRuneRange(t, 0x0346, 0x034E, pPVALID)
}

func Test034F(t *testing.T) {
	ValidateRune(t, 0x034F, pDISALLOWED)
}

func Test0350_036F(t *testing.T) {
	ValidateRuneRange(t, 0x0350, 0x036F, pPVALID)
}

func Test0378_0379(t *testing.T) {
	ValidateRuneRange(t, 0x0378, 0x0379, pUNASSIGNED)
}

func Test037A(t *testing.T) {
	ValidateRune(t, 0x037A, pDISALLOWED)
}

func Test037B_037D(t *testing.T) {
	ValidateRuneRange(t, 0x037B, 0x037D, pPVALID)
}

func Test037E(t *testing.T) {
	ValidateRune(t, 0x037E, pDISALLOWED)
}

func Test037F_0383(t *testing.T) {
	ValidateRuneRange(t, 0x037F, 0x0383, pUNASSIGNED)
}

func Test0384_038A(t *testing.T) {
	ValidateRuneRange(t, 0x0384, 0x038A, pDISALLOWED)
}

func Test038B(t *testing.T) {
	ValidateRune(t, 0x038B, pUNASSIGNED)
}

func Test038C(t *testing.T) {
	ValidateRune(t, 0x038C, pDISALLOWED)
}

func Test038D(t *testing.T) {
	ValidateRune(t, 0x038D, pUNASSIGNED)
}

func Test038E_038F(t *testing.T) {
	ValidateRuneRange(t, 0x038E, 0x038F, pDISALLOWED)
}

func Test0391_03A1(t *testing.T) {
	ValidateRuneRange(t, 0x0391, 0x03A1, pDISALLOWED)
}

func Test03A2(t *testing.T) {
	ValidateRune(t, 0x03A2, pUNASSIGNED)
}

func Test03A3_03AB(t *testing.T) {
	ValidateRuneRange(t, 0x03A3, 0x03AB, pDISALLOWED)
}

func Test03AC_03CE(t *testing.T) {
	ValidateRuneRange(t, 0x03AC, 0x03CE, pPVALID)
}

func Test03CF_03D6(t *testing.T) {
	ValidateRuneRange(t, 0x03CF, 0x03D6, pDISALLOWED)
}

func Test03D7(t *testing.T) {
	ValidateRune(t, 0x03D7, pPVALID)
}

func Test03D8(t *testing.T) {
	ValidateRune(t, 0x03D8, pDISALLOWED)
}

func Test03D9(t *testing.T) {
	ValidateRune(t, 0x03D9, pPVALID)
}

func Test03DA(t *testing.T) {
	ValidateRune(t, 0x03DA, pDISALLOWED)
}

func Test03DB(t *testing.T) {
	ValidateRune(t, 0x03DB, pPVALID)
}

func Test03DC(t *testing.T) {
	ValidateRune(t, 0x03DC, pDISALLOWED)
}

func Test03DD(t *testing.T) {
	ValidateRune(t, 0x03DD, pPVALID)
}

func Test03DE(t *testing.T) {
	ValidateRune(t, 0x03DE, pDISALLOWED)
}

func Test03DF(t *testing.T) {
	ValidateRune(t, 0x03DF, pPVALID)
}

func Test03E0(t *testing.T) {
	ValidateRune(t, 0x03E0, pDISALLOWED)
}

func Test03E1(t *testing.T) {
	ValidateRune(t, 0x03E1, pPVALID)
}

func Test03E2(t *testing.T) {
	ValidateRune(t, 0x03E2, pDISALLOWED)
}

func Test03E3(t *testing.T) {
	ValidateRune(t, 0x03E3, pPVALID)
}

func Test03E4(t *testing.T) {
	ValidateRune(t, 0x03E4, pDISALLOWED)
}

func Test03E5(t *testing.T) {
	ValidateRune(t, 0x03E5, pPVALID)
}

func Test03E6(t *testing.T) {
	ValidateRune(t, 0x03E6, pDISALLOWED)
}

func Test03E7(t *testing.T) {
	ValidateRune(t, 0x03E7, pPVALID)
}

func Test03E8(t *testing.T) {
	ValidateRune(t, 0x03E8, pDISALLOWED)
}

func Test03E9(t *testing.T) {
	ValidateRune(t, 0x03E9, pPVALID)
}

func Test03EA(t *testing.T) {
	ValidateRune(t, 0x03EA, pDISALLOWED)
}

func Test03EB(t *testing.T) {
	ValidateRune(t, 0x03EB, pPVALID)
}

func Test03EC(t *testing.T) {
	ValidateRune(t, 0x03EC, pDISALLOWED)
}

func Test03ED(t *testing.T) {
	ValidateRune(t, 0x03ED, pPVALID)
}

func Test03EE(t *testing.T) {
	ValidateRune(t, 0x03EE, pDISALLOWED)
}

func Test03EF(t *testing.T) {
	ValidateRune(t, 0x03EF, pPVALID)
}

func Test03F0_03F2(t *testing.T) {
	ValidateRuneRange(t, 0x03F0, 0x03F2, pDISALLOWED)
}

func Test03F3(t *testing.T) {
	ValidateRune(t, 0x03F3, pPVALID)
}

func Test03F4_03F7(t *testing.T) {
	ValidateRuneRange(t, 0x03F4, 0x03F7, pDISALLOWED)
}

func Test03F8(t *testing.T) {
	ValidateRune(t, 0x03F8, pPVALID)
}

func Test03F9_03FA(t *testing.T) {
	ValidateRuneRange(t, 0x03F9, 0x03FA, pDISALLOWED)
}

func Test03FB_03FC(t *testing.T) {
	ValidateRuneRange(t, 0x03FB, 0x03FC, pPVALID)
}

func Test03FD_042F(t *testing.T) {
	ValidateRuneRange(t, 0x03FD, 0x042F, pDISALLOWED)
}

func Test0430_045F(t *testing.T) {
	ValidateRuneRange(t, 0x0430, 0x045F, pPVALID)
}

func Test046A(t *testing.T) {
	ValidateRune(t, 0x046A, pDISALLOWED)
}

func Test046B(t *testing.T) {
	ValidateRune(t, 0x046B, pPVALID)
}

func Test046C(t *testing.T) {
	ValidateRune(t, 0x046C, pDISALLOWED)
}

func Test046D(t *testing.T) {
	ValidateRune(t, 0x046D, pPVALID)
}

func Test046E(t *testing.T) {
	ValidateRune(t, 0x046E, pDISALLOWED)
}

func Test046F(t *testing.T) {
	ValidateRune(t, 0x046F, pPVALID)
}

func Test047A(t *testing.T) {
	ValidateRune(t, 0x047A, pDISALLOWED)
}

func Test047B(t *testing.T) {
	ValidateRune(t, 0x047B, pPVALID)
}

func Test047C(t *testing.T) {
	ValidateRune(t, 0x047C, pDISALLOWED)
}

func Test047D(t *testing.T) {
	ValidateRune(t, 0x047D, pPVALID)
}

func Test047E(t *testing.T) {
	ValidateRune(t, 0x047E, pDISALLOWED)
}

func Test047F(t *testing.T) {
	ValidateRune(t, 0x047F, pPVALID)
}

func Test0483_0487(t *testing.T) {
	ValidateRuneRange(t, 0x0483, 0x0487, pPVALID)
}

func Test0488_048A(t *testing.T) {
	ValidateRuneRange(t, 0x0488, 0x048A, pDISALLOWED)
}

func Test048B(t *testing.T) {
	ValidateRune(t, 0x048B, pPVALID)
}

func Test048C(t *testing.T) {
	ValidateRune(t, 0x048C, pDISALLOWED)
}

func Test048D(t *testing.T) {
	ValidateRune(t, 0x048D, pPVALID)
}

func Test048E(t *testing.T) {
	ValidateRune(t, 0x048E, pDISALLOWED)
}

func Test048F(t *testing.T) {
	ValidateRune(t, 0x048F, pPVALID)
}

func Test049A(t *testing.T) {
	ValidateRune(t, 0x049A, pDISALLOWED)
}

func Test049B(t *testing.T) {
	ValidateRune(t, 0x049B, pPVALID)
}

func Test049C(t *testing.T) {
	ValidateRune(t, 0x049C, pDISALLOWED)
}

func Test049D(t *testing.T) {
	ValidateRune(t, 0x049D, pPVALID)
}

func Test049E(t *testing.T) {
	ValidateRune(t, 0x049E, pDISALLOWED)
}

func Test049F(t *testing.T) {
	ValidateRune(t, 0x049F, pPVALID)
}

func Test04A0(t *testing.T) {
	ValidateRune(t, 0x04A0, pDISALLOWED)
}

func Test04A1(t *testing.T) {
	ValidateRune(t, 0x04A1, pPVALID)
}

func Test04A2(t *testing.T) {
	ValidateRune(t, 0x04A2, pDISALLOWED)
}

func Test04A3(t *testing.T) {
	ValidateRune(t, 0x04A3, pPVALID)
}

func Test04A4(t *testing.T) {
	ValidateRune(t, 0x04A4, pDISALLOWED)
}

func Test04A5(t *testing.T) {
	ValidateRune(t, 0x04A5, pPVALID)
}

func Test04A6(t *testing.T) {
	ValidateRune(t, 0x04A6, pDISALLOWED)
}

func Test04A7(t *testing.T) {
	ValidateRune(t, 0x04A7, pPVALID)
}

func Test04A8(t *testing.T) {
	ValidateRune(t, 0x04A8, pDISALLOWED)
}

func Test04A9(t *testing.T) {
	ValidateRune(t, 0x04A9, pPVALID)
}

func Test04AA(t *testing.T) {
	ValidateRune(t, 0x04AA, pDISALLOWED)
}

func Test04AB(t *testing.T) {
	ValidateRune(t, 0x04AB, pPVALID)
}

func Test04AC(t *testing.T) {
	ValidateRune(t, 0x04AC, pDISALLOWED)
}

func Test04AD(t *testing.T) {
	ValidateRune(t, 0x04AD, pPVALID)
}

func Test04AE(t *testing.T) {
	ValidateRune(t, 0x04AE, pDISALLOWED)
}

func Test04AF(t *testing.T) {
	ValidateRune(t, 0x04AF, pPVALID)
}

func Test04B0(t *testing.T) {
	ValidateRune(t, 0x04B0, pDISALLOWED)
}

func Test04B1(t *testing.T) {
	ValidateRune(t, 0x04B1, pPVALID)
}

func Test04B2(t *testing.T) {
	ValidateRune(t, 0x04B2, pDISALLOWED)
}

func Test04B3(t *testing.T) {
	ValidateRune(t, 0x04B3, pPVALID)
}

func Test04B4(t *testing.T) {
	ValidateRune(t, 0x04B4, pDISALLOWED)
}

func Test04B5(t *testing.T) {
	ValidateRune(t, 0x04B5, pPVALID)
}

func Test04B6(t *testing.T) {
	ValidateRune(t, 0x04B6, pDISALLOWED)
}

func Test04B7(t *testing.T) {
	ValidateRune(t, 0x04B7, pPVALID)
}

func Test04B8(t *testing.T) {
	ValidateRune(t, 0x04B8, pDISALLOWED)
}

func Test04B9(t *testing.T) {
	ValidateRune(t, 0x04B9, pPVALID)
}

func Test04BA(t *testing.T) {
	ValidateRune(t, 0x04BA, pDISALLOWED)
}

func Test04BB(t *testing.T) {
	ValidateRune(t, 0x04BB, pPVALID)
}

func Test04BC(t *testing.T) {
	ValidateRune(t, 0x04BC, pDISALLOWED)
}

func Test04BD(t *testing.T) {
	ValidateRune(t, 0x04BD, pPVALID)
}

func Test04BE(t *testing.T) {
	ValidateRune(t, 0x04BE, pDISALLOWED)
}

func Test04BF(t *testing.T) {
	ValidateRune(t, 0x04BF, pPVALID)
}

func Test04C0_04C1(t *testing.T) {
	ValidateRuneRange(t, 0x04C0, 0x04C1, pDISALLOWED)
}

func Test04C2(t *testing.T) {
	ValidateRune(t, 0x04C2, pPVALID)
}

func Test04C3(t *testing.T) {
	ValidateRune(t, 0x04C3, pDISALLOWED)
}

func Test04C4(t *testing.T) {
	ValidateRune(t, 0x04C4, pPVALID)
}

func Test04C5(t *testing.T) {
	ValidateRune(t, 0x04C5, pDISALLOWED)
}

func Test04C6(t *testing.T) {
	ValidateRune(t, 0x04C6, pPVALID)
}

func Test04C7(t *testing.T) {
	ValidateRune(t, 0x04C7, pDISALLOWED)
}

func Test04C8(t *testing.T) {
	ValidateRune(t, 0x04C8, pPVALID)
}

func Test04C9(t *testing.T) {
	ValidateRune(t, 0x04C9, pDISALLOWED)
}

func Test04CA(t *testing.T) {
	ValidateRune(t, 0x04CA, pPVALID)
}

func Test04CB(t *testing.T) {
	ValidateRune(t, 0x04CB, pDISALLOWED)
}

func Test04CC(t *testing.T) {
	ValidateRune(t, 0x04CC, pPVALID)
}

func Test04CD(t *testing.T) {
	ValidateRune(t, 0x04CD, pDISALLOWED)
}

func Test04CE_04CF(t *testing.T) {
	ValidateRuneRange(t, 0x04CE, 0x04CF, pPVALID)
}

func Test04D0(t *testing.T) {
	ValidateRune(t, 0x04D0, pDISALLOWED)
}

func Test04D1(t *testing.T) {
	ValidateRune(t, 0x04D1, pPVALID)
}

func Test04D2(t *testing.T) {
	ValidateRune(t, 0x04D2, pDISALLOWED)
}

func Test04D3(t *testing.T) {
	ValidateRune(t, 0x04D3, pPVALID)
}

func Test04D4(t *testing.T) {
	ValidateRune(t, 0x04D4, pDISALLOWED)
}

func Test04D5(t *testing.T) {
	ValidateRune(t, 0x04D5, pPVALID)
}

func Test04D6(t *testing.T) {
	ValidateRune(t, 0x04D6, pDISALLOWED)
}

func Test04D7(t *testing.T) {
	ValidateRune(t, 0x04D7, pPVALID)
}

func Test04D8(t *testing.T) {
	ValidateRune(t, 0x04D8, pDISALLOWED)
}

func Test04D9(t *testing.T) {
	ValidateRune(t, 0x04D9, pPVALID)
}

func Test04DA(t *testing.T) {
	ValidateRune(t, 0x04DA, pDISALLOWED)
}

func Test04DB(t *testing.T) {
	ValidateRune(t, 0x04DB, pPVALID)
}

func Test04DC(t *testing.T) {
	ValidateRune(t, 0x04DC, pDISALLOWED)
}

func Test04DD(t *testing.T) {
	ValidateRune(t, 0x04DD, pPVALID)
}

func Test04DE(t *testing.T) {
	ValidateRune(t, 0x04DE, pDISALLOWED)
}

func Test04DF(t *testing.T) {
	ValidateRune(t, 0x04DF, pPVALID)
}

func Test04E0(t *testing.T) {
	ValidateRune(t, 0x04E0, pDISALLOWED)
}

func Test04E1(t *testing.T) {
	ValidateRune(t, 0x04E1, pPVALID)
}

func Test04E2(t *testing.T) {
	ValidateRune(t, 0x04E2, pDISALLOWED)
}

func Test04E3(t *testing.T) {
	ValidateRune(t, 0x04E3, pPVALID)
}

func Test04E4(t *testing.T) {
	ValidateRune(t, 0x04E4, pDISALLOWED)
}

func Test04E5(t *testing.T) {
	ValidateRune(t, 0x04E5, pPVALID)
}

func Test04E6(t *testing.T) {
	ValidateRune(t, 0x04E6, pDISALLOWED)
}

func Test04E7(t *testing.T) {
	ValidateRune(t, 0x04E7, pPVALID)
}

func Test04E8(t *testing.T) {
	ValidateRune(t, 0x04E8, pDISALLOWED)
}

func Test04E9(t *testing.T) {
	ValidateRune(t, 0x04E9, pPVALID)
}

func Test04EA(t *testing.T) {
	ValidateRune(t, 0x04EA, pDISALLOWED)
}

func Test04EB(t *testing.T) {
	ValidateRune(t, 0x04EB, pPVALID)
}

func Test04EC(t *testing.T) {
	ValidateRune(t, 0x04EC, pDISALLOWED)
}

func Test04ED(t *testing.T) {
	ValidateRune(t, 0x04ED, pPVALID)
}

func Test04EE(t *testing.T) {
	ValidateRune(t, 0x04EE, pDISALLOWED)
}

func Test04EF(t *testing.T) {
	ValidateRune(t, 0x04EF, pPVALID)
}

func Test04F0(t *testing.T) {
	ValidateRune(t, 0x04F0, pDISALLOWED)
}

func Test04F1(t *testing.T) {
	ValidateRune(t, 0x04F1, pPVALID)
}

func Test04F2(t *testing.T) {
	ValidateRune(t, 0x04F2, pDISALLOWED)
}

func Test04F3(t *testing.T) {
	ValidateRune(t, 0x04F3, pPVALID)
}

func Test04F4(t *testing.T) {
	ValidateRune(t, 0x04F4, pDISALLOWED)
}

func Test04F5(t *testing.T) {
	ValidateRune(t, 0x04F5, pPVALID)
}

func Test04F6(t *testing.T) {
	ValidateRune(t, 0x04F6, pDISALLOWED)
}

func Test04F7(t *testing.T) {
	ValidateRune(t, 0x04F7, pPVALID)
}

func Test04F8(t *testing.T) {
	ValidateRune(t, 0x04F8, pDISALLOWED)
}

func Test04F9(t *testing.T) {
	ValidateRune(t, 0x04F9, pPVALID)
}

func Test04FA(t *testing.T) {
	ValidateRune(t, 0x04FA, pDISALLOWED)
}

func Test04FB(t *testing.T) {
	ValidateRune(t, 0x04FB, pPVALID)
}

func Test04FC(t *testing.T) {
	ValidateRune(t, 0x04FC, pDISALLOWED)
}

func Test04FD(t *testing.T) {
	ValidateRune(t, 0x04FD, pPVALID)
}

func Test04FE(t *testing.T) {
	ValidateRune(t, 0x04FE, pDISALLOWED)
}

func Test04FF(t *testing.T) {
	ValidateRune(t, 0x04FF, pPVALID)
}

func Test050A(t *testing.T) {
	ValidateRune(t, 0x050A, pDISALLOWED)
}

func Test050B(t *testing.T) {
	ValidateRune(t, 0x050B, pPVALID)
}

func Test050C(t *testing.T) {
	ValidateRune(t, 0x050C, pDISALLOWED)
}

func Test050D(t *testing.T) {
	ValidateRune(t, 0x050D, pPVALID)
}

func Test050E(t *testing.T) {
	ValidateRune(t, 0x050E, pDISALLOWED)
}

func Test050F(t *testing.T) {
	ValidateRune(t, 0x050F, pPVALID)
}

func Test051A(t *testing.T) {
	ValidateRune(t, 0x051A, pDISALLOWED)
}

func Test051B(t *testing.T) {
	ValidateRune(t, 0x051B, pPVALID)
}

func Test051C(t *testing.T) {
	ValidateRune(t, 0x051C, pDISALLOWED)
}

func Test051D(t *testing.T) {
	ValidateRune(t, 0x051D, pPVALID)
}

func Test051E(t *testing.T) {
	ValidateRune(t, 0x051E, pDISALLOWED)
}

func Test051F(t *testing.T) {
	ValidateRune(t, 0x051F, pPVALID)
}

func Test0526_0530(t *testing.T) {
	ValidateRuneRange(t, 0x0526, 0x0530, pUNASSIGNED)
}

func Test0531_0556(t *testing.T) {
	ValidateRuneRange(t, 0x0531, 0x0556, pDISALLOWED)
}

func Test0557_0558(t *testing.T) {
	ValidateRuneRange(t, 0x0557, 0x0558, pUNASSIGNED)
}

func Test055A_055F(t *testing.T) {
	ValidateRuneRange(t, 0x055A, 0x055F, pDISALLOWED)
}

func Test0561_0586(t *testing.T) {
	ValidateRuneRange(t, 0x0561, 0x0586, pPVALID)
}

func Test0589_058A(t *testing.T) {
	ValidateRuneRange(t, 0x0589, 0x058A, pDISALLOWED)
}

func Test058B_0590(t *testing.T) {
	ValidateRuneRange(t, 0x058B, 0x0590, pUNASSIGNED)
}

func Test0591_05BD(t *testing.T) {
	ValidateRuneRange(t, 0x0591, 0x05BD, pPVALID)
}

func Test05BE(t *testing.T) {
	ValidateRune(t, 0x05BE, pDISALLOWED)
}

func Test05BF(t *testing.T) {
	ValidateRune(t, 0x05BF, pPVALID)
}

func Test05C0(t *testing.T) {
	ValidateRune(t, 0x05C0, pDISALLOWED)
}

func Test05C1_05C2(t *testing.T) {
	ValidateRuneRange(t, 0x05C1, 0x05C2, pPVALID)
}

func Test05C3(t *testing.T) {
	ValidateRune(t, 0x05C3, pDISALLOWED)
}

func Test05C4_05C5(t *testing.T) {
	ValidateRuneRange(t, 0x05C4, 0x05C5, pPVALID)
}

func Test05C6(t *testing.T) {
	ValidateRune(t, 0x05C6, pDISALLOWED)
}

func Test05C7(t *testing.T) {
	ValidateRune(t, 0x05C7, pPVALID)
}

func Test05C8_05CF(t *testing.T) {
	ValidateRuneRange(t, 0x05C8, 0x05CF, pUNASSIGNED)
}

func Test05D0_05EA(t *testing.T) {
	ValidateRuneRange(t, 0x05D0, 0x05EA, pPVALID)
}

func Test05EB_05EF(t *testing.T) {
	ValidateRuneRange(t, 0x05EB, 0x05EF, pUNASSIGNED)
}

func Test05F0_05F2(t *testing.T) {
	ValidateRuneRange(t, 0x05F0, 0x05F2, pPVALID)
}

func Test05F3_05F4(t *testing.T) {
	ValidateRuneRange(t, 0x05F3, 0x05F4, pCONTEXTO)
}

func Test05F5_05FF(t *testing.T) {
	ValidateRuneRange(t, 0x05F5, 0x05FF, pUNASSIGNED)
}

func Test0600_0603(t *testing.T) {
	ValidateRuneRange(t, 0x0600, 0x0603, pDISALLOWED)
}

func Test0604_0605(t *testing.T) {
	ValidateRuneRange(t, 0x0604, 0x0605, pUNASSIGNED)
}

func Test0606_060F(t *testing.T) {
	ValidateRuneRange(t, 0x0606, 0x060F, pDISALLOWED)
}

func Test0610_061A(t *testing.T) {
	ValidateRuneRange(t, 0x0610, 0x061A, pPVALID)
}

func Test061B(t *testing.T) {
	ValidateRune(t, 0x061B, pDISALLOWED)
}

func Test061C_061D(t *testing.T) {
	ValidateRuneRange(t, 0x061C, 0x061D, pUNASSIGNED)
}

func Test061E_061F(t *testing.T) {
	ValidateRuneRange(t, 0x061E, 0x061F, pDISALLOWED)
}

func Test0621_063F(t *testing.T) {
	ValidateRuneRange(t, 0x0621, 0x063F, pPVALID)
}

func Test0641_065E(t *testing.T) {
	ValidateRuneRange(t, 0x0641, 0x065E, pPVALID)
}

func Test065F(t *testing.T) {
	ValidateRune(t, 0x065F, pUNASSIGNED)
}

func Test0660_0669(t *testing.T) {
	ValidateRuneRange(t, 0x0660, 0x0669, pCONTEXTO)
}

func Test066A_066D(t *testing.T) {
	ValidateRuneRange(t, 0x066A, 0x066D, pDISALLOWED)
}

func Test066E_0674(t *testing.T) {
	ValidateRuneRange(t, 0x066E, 0x0674, pPVALID)
}

func Test0675_0678(t *testing.T) {
	ValidateRuneRange(t, 0x0675, 0x0678, pDISALLOWED)
}

func Test0679_06D3(t *testing.T) {
	ValidateRuneRange(t, 0x0679, 0x06D3, pPVALID)
}

func Test06D4(t *testing.T) {
	ValidateRune(t, 0x06D4, pDISALLOWED)
}

func Test06D5_06DC(t *testing.T) {
	ValidateRuneRange(t, 0x06D5, 0x06DC, pPVALID)
}

func Test06DD_06DE(t *testing.T) {
	ValidateRuneRange(t, 0x06DD, 0x06DE, pDISALLOWED)
}

func Test06DF_06E8(t *testing.T) {
	ValidateRuneRange(t, 0x06DF, 0x06E8, pPVALID)
}

func Test06E9(t *testing.T) {
	ValidateRune(t, 0x06E9, pDISALLOWED)
}

func Test06EA_06EF(t *testing.T) {
	ValidateRuneRange(t, 0x06EA, 0x06EF, pPVALID)
}

func Test06F0_06F9(t *testing.T) {
	ValidateRuneRange(t, 0x06F0, 0x06F9, pCONTEXTO)
}

func Test06FA_06FF(t *testing.T) {
	ValidateRuneRange(t, 0x06FA, 0x06FF, pPVALID)
}

func Test0700_070D(t *testing.T) {
	ValidateRuneRange(t, 0x0700, 0x070D, pDISALLOWED)
}

func Test070E(t *testing.T) {
	ValidateRune(t, 0x070E, pUNASSIGNED)
}

func Test070F(t *testing.T) {
	ValidateRune(t, 0x070F, pDISALLOWED)
}

func Test0710_074A(t *testing.T) {
	ValidateRuneRange(t, 0x0710, 0x074A, pPVALID)
}

func Test074B_074C(t *testing.T) {
	ValidateRuneRange(t, 0x074B, 0x074C, pUNASSIGNED)
}

func Test074D_07B1(t *testing.T) {
	ValidateRuneRange(t, 0x074D, 0x07B1, pPVALID)
}

func Test07B2_07BF(t *testing.T) {
	ValidateRuneRange(t, 0x07B2, 0x07BF, pUNASSIGNED)
}

func Test07C0_07F5(t *testing.T) {
	ValidateRuneRange(t, 0x07C0, 0x07F5, pPVALID)
}

func Test07F6_07FA(t *testing.T) {
	ValidateRuneRange(t, 0x07F6, 0x07FA, pDISALLOWED)
}

func Test07FB_07FF(t *testing.T) {
	ValidateRuneRange(t, 0x07FB, 0x07FF, pUNASSIGNED)
}

func Test0800_082D(t *testing.T) {
	ValidateRuneRange(t, 0x0800, 0x082D, pPVALID)
}

func Test082E_082F(t *testing.T) {
	ValidateRuneRange(t, 0x082E, 0x082F, pUNASSIGNED)
}

func Test0830_083E(t *testing.T) {
	ValidateRuneRange(t, 0x0830, 0x083E, pDISALLOWED)
}

func Test083F_08FF(t *testing.T) {
	ValidateRuneRange(t, 0x083F, 0x08FF, pUNASSIGNED)
}

func Test0900_0939(t *testing.T) {
	ValidateRuneRange(t, 0x0900, 0x0939, pPVALID)
}

func Test093A_093B(t *testing.T) {
	ValidateRuneRange(t, 0x093A, 0x093B, pUNASSIGNED)
}

func Test093C_094E(t *testing.T) {
	ValidateRuneRange(t, 0x093C, 0x094E, pPVALID)
}

func Test094F(t *testing.T) {
	ValidateRune(t, 0x094F, pUNASSIGNED)
}

func Test0950_0955(t *testing.T) {
	ValidateRuneRange(t, 0x0950, 0x0955, pPVALID)
}

func Test0956_0957(t *testing.T) {
	ValidateRuneRange(t, 0x0956, 0x0957, pUNASSIGNED)
}

func Test0958_095F(t *testing.T) {
	ValidateRuneRange(t, 0x0958, 0x095F, pDISALLOWED)
}

func Test0960_0963(t *testing.T) {
	ValidateRuneRange(t, 0x0960, 0x0963, pPVALID)
}

func Test0964_0965(t *testing.T) {
	ValidateRuneRange(t, 0x0964, 0x0965, pDISALLOWED)
}

func Test0966_096F(t *testing.T) {
	ValidateRuneRange(t, 0x0966, 0x096F, pPVALID)
}

func Test0971_0972(t *testing.T) {
	ValidateRuneRange(t, 0x0971, 0x0972, pPVALID)
}

func Test0973_0978(t *testing.T) {
	ValidateRuneRange(t, 0x0973, 0x0978, pUNASSIGNED)
}

func Test0979_097F(t *testing.T) {
	ValidateRuneRange(t, 0x0979, 0x097F, pPVALID)
}

func Test0981_0983(t *testing.T) {
	ValidateRuneRange(t, 0x0981, 0x0983, pPVALID)
}

func Test0985_098C(t *testing.T) {
	ValidateRuneRange(t, 0x0985, 0x098C, pPVALID)
}

func Test098D_098E(t *testing.T) {
	ValidateRuneRange(t, 0x098D, 0x098E, pUNASSIGNED)
}

func Test098F_0990(t *testing.T) {
	ValidateRuneRange(t, 0x098F, 0x0990, pPVALID)
}

func Test0991_0992(t *testing.T) {
	ValidateRuneRange(t, 0x0991, 0x0992, pUNASSIGNED)
}

func Test0993_09A8(t *testing.T) {
	ValidateRuneRange(t, 0x0993, 0x09A8, pPVALID)
}

func Test09A9(t *testing.T) {
	ValidateRune(t, 0x09A9, pUNASSIGNED)
}

func Test09AA_09B0(t *testing.T) {
	ValidateRuneRange(t, 0x09AA, 0x09B0, pPVALID)
}

func Test09B1(t *testing.T) {
	ValidateRune(t, 0x09B1, pUNASSIGNED)
}

func Test09B2(t *testing.T) {
	ValidateRune(t, 0x09B2, pPVALID)
}

func Test09B3_09B5(t *testing.T) {
	ValidateRuneRange(t, 0x09B3, 0x09B5, pUNASSIGNED)
}

func Test09B6_09B9(t *testing.T) {
	ValidateRuneRange(t, 0x09B6, 0x09B9, pPVALID)
}

func Test09BA_09BB(t *testing.T) {
	ValidateRuneRange(t, 0x09BA, 0x09BB, pUNASSIGNED)
}

func Test09BC_09C4(t *testing.T) {
	ValidateRuneRange(t, 0x09BC, 0x09C4, pPVALID)
}

func Test09C5_09C6(t *testing.T) {
	ValidateRuneRange(t, 0x09C5, 0x09C6, pUNASSIGNED)
}

func Test09C7_09C8(t *testing.T) {
	ValidateRuneRange(t, 0x09C7, 0x09C8, pPVALID)
}

func Test09C9_09CA(t *testing.T) {
	ValidateRuneRange(t, 0x09C9, 0x09CA, pUNASSIGNED)
}

func Test09CB_09CE(t *testing.T) {
	ValidateRuneRange(t, 0x09CB, 0x09CE, pPVALID)
}

func Test09CF_09D6(t *testing.T) {
	ValidateRuneRange(t, 0x09CF, 0x09D6, pUNASSIGNED)
}

func Test09D7(t *testing.T) {
	ValidateRune(t, 0x09D7, pPVALID)
}

func Test09D8_09DB(t *testing.T) {
	ValidateRuneRange(t, 0x09D8, 0x09DB, pUNASSIGNED)
}

func Test09DC_09DD(t *testing.T) {
	ValidateRuneRange(t, 0x09DC, 0x09DD, pDISALLOWED)
}

func Test09DE(t *testing.T) {
	ValidateRune(t, 0x09DE, pUNASSIGNED)
}

func Test09DF(t *testing.T) {
	ValidateRune(t, 0x09DF, pDISALLOWED)
}

func Test09E0_09E3(t *testing.T) {
	ValidateRuneRange(t, 0x09E0, 0x09E3, pPVALID)
}

func Test09E4_09E5(t *testing.T) {
	ValidateRuneRange(t, 0x09E4, 0x09E5, pUNASSIGNED)
}

func Test09E6_09F1(t *testing.T) {
	ValidateRuneRange(t, 0x09E6, 0x09F1, pPVALID)
}

func Test09F2_09FB(t *testing.T) {
	ValidateRuneRange(t, 0x09F2, 0x09FB, pDISALLOWED)
}

func Test09FC_0A00(t *testing.T) {
	ValidateRuneRange(t, 0x09FC, 0x0A00, pUNASSIGNED)
}

func Test0A01_0A03(t *testing.T) {
	ValidateRuneRange(t, 0x0A01, 0x0A03, pPVALID)
}

func Test0A04(t *testing.T) {
	ValidateRune(t, 0x0A04, pUNASSIGNED)
}

func Test0A05_0A0A(t *testing.T) {
	ValidateRuneRange(t, 0x0A05, 0x0A0A, pPVALID)
}

func Test0A0B_0A0E(t *testing.T) {
	ValidateRuneRange(t, 0x0A0B, 0x0A0E, pUNASSIGNED)
}

func Test0A0F_0A10(t *testing.T) {
	ValidateRuneRange(t, 0x0A0F, 0x0A10, pPVALID)
}

func Test0A11_0A12(t *testing.T) {
	ValidateRuneRange(t, 0x0A11, 0x0A12, pUNASSIGNED)
}

func Test0A13_0A28(t *testing.T) {
	ValidateRuneRange(t, 0x0A13, 0x0A28, pPVALID)
}

func Test0A29(t *testing.T) {
	ValidateRune(t, 0x0A29, pUNASSIGNED)
}

func Test0A2A_0A30(t *testing.T) {
	ValidateRuneRange(t, 0x0A2A, 0x0A30, pPVALID)
}

func Test0A31(t *testing.T) {
	ValidateRune(t, 0x0A31, pUNASSIGNED)
}

func Test0A32(t *testing.T) {
	ValidateRune(t, 0x0A32, pPVALID)
}

func Test0A33(t *testing.T) {
	ValidateRune(t, 0x0A33, pDISALLOWED)
}

func Test0A34(t *testing.T) {
	ValidateRune(t, 0x0A34, pUNASSIGNED)
}

func Test0A35(t *testing.T) {
	ValidateRune(t, 0x0A35, pPVALID)
}

func Test0A36(t *testing.T) {
	ValidateRune(t, 0x0A36, pDISALLOWED)
}

func Test0A37(t *testing.T) {
	ValidateRune(t, 0x0A37, pUNASSIGNED)
}

func Test0A38_0A39(t *testing.T) {
	ValidateRuneRange(t, 0x0A38, 0x0A39, pPVALID)
}

func Test0A3A_0A3B(t *testing.T) {
	ValidateRuneRange(t, 0x0A3A, 0x0A3B, pUNASSIGNED)
}

func Test0A3C(t *testing.T) {
	ValidateRune(t, 0x0A3C, pPVALID)
}

func Test0A3D(t *testing.T) {
	ValidateRune(t, 0x0A3D, pUNASSIGNED)
}

func Test0A3E_0A42(t *testing.T) {
	ValidateRuneRange(t, 0x0A3E, 0x0A42, pPVALID)
}

func Test0A43_0A46(t *testing.T) {
	ValidateRuneRange(t, 0x0A43, 0x0A46, pUNASSIGNED)
}

func Test0A47_0A48(t *testing.T) {
	ValidateRuneRange(t, 0x0A47, 0x0A48, pPVALID)
}

func Test0A49_0A4A(t *testing.T) {
	ValidateRuneRange(t, 0x0A49, 0x0A4A, pUNASSIGNED)
}

func Test0A4B_0A4D(t *testing.T) {
	ValidateRuneRange(t, 0x0A4B, 0x0A4D, pPVALID)
}

func Test0A4E_0A50(t *testing.T) {
	ValidateRuneRange(t, 0x0A4E, 0x0A50, pUNASSIGNED)
}

func Test0A51(t *testing.T) {
	ValidateRune(t, 0x0A51, pPVALID)
}

func Test0A52_0A58(t *testing.T) {
	ValidateRuneRange(t, 0x0A52, 0x0A58, pUNASSIGNED)
}

func Test0A59_0A5B(t *testing.T) {
	ValidateRuneRange(t, 0x0A59, 0x0A5B, pDISALLOWED)
}

func Test0A5C(t *testing.T) {
	ValidateRune(t, 0x0A5C, pPVALID)
}

func Test0A5D(t *testing.T) {
	ValidateRune(t, 0x0A5D, pUNASSIGNED)
}

func Test0A5E(t *testing.T) {
	ValidateRune(t, 0x0A5E, pDISALLOWED)
}

func Test0A5F_0A65(t *testing.T) {
	ValidateRuneRange(t, 0x0A5F, 0x0A65, pUNASSIGNED)
}

func Test0A66_0A75(t *testing.T) {
	ValidateRuneRange(t, 0x0A66, 0x0A75, pPVALID)
}

func Test0A76_0A80(t *testing.T) {
	ValidateRuneRange(t, 0x0A76, 0x0A80, pUNASSIGNED)
}

func Test0A81_0A83(t *testing.T) {
	ValidateRuneRange(t, 0x0A81, 0x0A83, pPVALID)
}

func Test0A84(t *testing.T) {
	ValidateRune(t, 0x0A84, pUNASSIGNED)
}

func Test0A85_0A8D(t *testing.T) {
	ValidateRuneRange(t, 0x0A85, 0x0A8D, pPVALID)
}

func Test0A8E(t *testing.T) {
	ValidateRune(t, 0x0A8E, pUNASSIGNED)
}

func Test0A8F_0A91(t *testing.T) {
	ValidateRuneRange(t, 0x0A8F, 0x0A91, pPVALID)
}

func Test0A92(t *testing.T) {
	ValidateRune(t, 0x0A92, pUNASSIGNED)
}

func Test0A93_0AA8(t *testing.T) {
	ValidateRuneRange(t, 0x0A93, 0x0AA8, pPVALID)
}

func Test0AA9(t *testing.T) {
	ValidateRune(t, 0x0AA9, pUNASSIGNED)
}

func Test0AAA_0AB0(t *testing.T) {
	ValidateRuneRange(t, 0x0AAA, 0x0AB0, pPVALID)
}

func Test0AB1(t *testing.T) {
	ValidateRune(t, 0x0AB1, pUNASSIGNED)
}

func Test0AB2_0AB3(t *testing.T) {
	ValidateRuneRange(t, 0x0AB2, 0x0AB3, pPVALID)
}

func Test0AB4(t *testing.T) {
	ValidateRune(t, 0x0AB4, pUNASSIGNED)
}

func Test0AB5_0AB9(t *testing.T) {
	ValidateRuneRange(t, 0x0AB5, 0x0AB9, pPVALID)
}

func Test0ABA_0ABB(t *testing.T) {
	ValidateRuneRange(t, 0x0ABA, 0x0ABB, pUNASSIGNED)
}

func Test0ABC_0AC5(t *testing.T) {
	ValidateRuneRange(t, 0x0ABC, 0x0AC5, pPVALID)
}

func Test0AC6(t *testing.T) {
	ValidateRune(t, 0x0AC6, pUNASSIGNED)
}

func Test0AC7_0AC9(t *testing.T) {
	ValidateRuneRange(t, 0x0AC7, 0x0AC9, pPVALID)
}

func Test0ACA(t *testing.T) {
	ValidateRune(t, 0x0ACA, pUNASSIGNED)
}

func Test0ACB_0ACD(t *testing.T) {
	ValidateRuneRange(t, 0x0ACB, 0x0ACD, pPVALID)
}

func Test0ACE_0ACF(t *testing.T) {
	ValidateRuneRange(t, 0x0ACE, 0x0ACF, pUNASSIGNED)
}

func Test0AD0(t *testing.T) {
	ValidateRune(t, 0x0AD0, pPVALID)
}

func Test0AD1_0ADF(t *testing.T) {
	ValidateRuneRange(t, 0x0AD1, 0x0ADF, pUNASSIGNED)
}

func Test0AE0_0AE3(t *testing.T) {
	ValidateRuneRange(t, 0x0AE0, 0x0AE3, pPVALID)
}

func Test0AE4_0AE5(t *testing.T) {
	ValidateRuneRange(t, 0x0AE4, 0x0AE5, pUNASSIGNED)
}

func Test0AE6_0AEF(t *testing.T) {
	ValidateRuneRange(t, 0x0AE6, 0x0AEF, pPVALID)
}

func Test0AF0(t *testing.T) {
	ValidateRune(t, 0x0AF0, pUNASSIGNED)
}

func Test0AF1(t *testing.T) {
	ValidateRune(t, 0x0AF1, pDISALLOWED)
}

func Test0AF2_0B00(t *testing.T) {
	ValidateRuneRange(t, 0x0AF2, 0x0B00, pUNASSIGNED)
}

func Test0B01_0B03(t *testing.T) {
	ValidateRuneRange(t, 0x0B01, 0x0B03, pPVALID)
}

func Test0B04(t *testing.T) {
	ValidateRune(t, 0x0B04, pUNASSIGNED)
}

func Test0B05_0B0C(t *testing.T) {
	ValidateRuneRange(t, 0x0B05, 0x0B0C, pPVALID)
}

func Test0B0D_0B0E(t *testing.T) {
	ValidateRuneRange(t, 0x0B0D, 0x0B0E, pUNASSIGNED)
}

func Test0B0F_0B10(t *testing.T) {
	ValidateRuneRange(t, 0x0B0F, 0x0B10, pPVALID)
}

func Test0B11_0B12(t *testing.T) {
	ValidateRuneRange(t, 0x0B11, 0x0B12, pUNASSIGNED)
}

func Test0B13_0B28(t *testing.T) {
	ValidateRuneRange(t, 0x0B13, 0x0B28, pPVALID)
}

func Test0B29(t *testing.T) {
	ValidateRune(t, 0x0B29, pUNASSIGNED)
}

func Test0B2A_0B30(t *testing.T) {
	ValidateRuneRange(t, 0x0B2A, 0x0B30, pPVALID)
}

func Test0B31(t *testing.T) {
	ValidateRune(t, 0x0B31, pUNASSIGNED)
}

func Test0B32_0B33(t *testing.T) {
	ValidateRuneRange(t, 0x0B32, 0x0B33, pPVALID)
}

func Test0B34(t *testing.T) {
	ValidateRune(t, 0x0B34, pUNASSIGNED)
}

func Test0B35_0B39(t *testing.T) {
	ValidateRuneRange(t, 0x0B35, 0x0B39, pPVALID)
}

func Test0B3A_0B3B(t *testing.T) {
	ValidateRuneRange(t, 0x0B3A, 0x0B3B, pUNASSIGNED)
}

func Test0B3C_0B44(t *testing.T) {
	ValidateRuneRange(t, 0x0B3C, 0x0B44, pPVALID)
}

func Test0B45_0B46(t *testing.T) {
	ValidateRuneRange(t, 0x0B45, 0x0B46, pUNASSIGNED)
}

func Test0B47_0B48(t *testing.T) {
	ValidateRuneRange(t, 0x0B47, 0x0B48, pPVALID)
}

func Test0B49_0B4A(t *testing.T) {
	ValidateRuneRange(t, 0x0B49, 0x0B4A, pUNASSIGNED)
}

func Test0B4B_0B4D(t *testing.T) {
	ValidateRuneRange(t, 0x0B4B, 0x0B4D, pPVALID)
}

func Test0B4E_0B55(t *testing.T) {
	ValidateRuneRange(t, 0x0B4E, 0x0B55, pUNASSIGNED)
}

func Test0B56_0B57(t *testing.T) {
	ValidateRuneRange(t, 0x0B56, 0x0B57, pPVALID)
}

func Test0B58_0B5B(t *testing.T) {
	ValidateRuneRange(t, 0x0B58, 0x0B5B, pUNASSIGNED)
}

func Test0B5C_0B5D(t *testing.T) {
	ValidateRuneRange(t, 0x0B5C, 0x0B5D, pDISALLOWED)
}

func Test0B5E(t *testing.T) {
	ValidateRune(t, 0x0B5E, pUNASSIGNED)
}

func Test0B5F_0B63(t *testing.T) {
	ValidateRuneRange(t, 0x0B5F, 0x0B63, pPVALID)
}

func Test0B64_0B65(t *testing.T) {
	ValidateRuneRange(t, 0x0B64, 0x0B65, pUNASSIGNED)
}

func Test0B66_0B6F(t *testing.T) {
	ValidateRuneRange(t, 0x0B66, 0x0B6F, pPVALID)
}

func Test0B70(t *testing.T) {
	ValidateRune(t, 0x0B70, pDISALLOWED)
}

func Test0B71(t *testing.T) {
	ValidateRune(t, 0x0B71, pPVALID)
}

func Test0B72_0B81(t *testing.T) {
	ValidateRuneRange(t, 0x0B72, 0x0B81, pUNASSIGNED)
}

func Test0B82_0B83(t *testing.T) {
	ValidateRuneRange(t, 0x0B82, 0x0B83, pPVALID)
}

func Test0B84(t *testing.T) {
	ValidateRune(t, 0x0B84, pUNASSIGNED)
}

func Test0B85_0B8A(t *testing.T) {
	ValidateRuneRange(t, 0x0B85, 0x0B8A, pPVALID)
}

func Test0B8B_0B8D(t *testing.T) {
	ValidateRuneRange(t, 0x0B8B, 0x0B8D, pUNASSIGNED)
}

func Test0B8E_0B90(t *testing.T) {
	ValidateRuneRange(t, 0x0B8E, 0x0B90, pPVALID)
}

func Test0B91(t *testing.T) {
	ValidateRune(t, 0x0B91, pUNASSIGNED)
}

func Test0B92_0B95(t *testing.T) {
	ValidateRuneRange(t, 0x0B92, 0x0B95, pPVALID)
}

func Test0B96_0B98(t *testing.T) {
	ValidateRuneRange(t, 0x0B96, 0x0B98, pUNASSIGNED)
}

func Test0B99_0B9A(t *testing.T) {
	ValidateRuneRange(t, 0x0B99, 0x0B9A, pPVALID)
}

func Test0B9B(t *testing.T) {
	ValidateRune(t, 0x0B9B, pUNASSIGNED)
}

func Test0B9C(t *testing.T) {
	ValidateRune(t, 0x0B9C, pPVALID)
}

func Test0B9D(t *testing.T) {
	ValidateRune(t, 0x0B9D, pUNASSIGNED)
}

func Test0B9E_0B9F(t *testing.T) {
	ValidateRuneRange(t, 0x0B9E, 0x0B9F, pPVALID)
}

func Test0BA0_0BA2(t *testing.T) {
	ValidateRuneRange(t, 0x0BA0, 0x0BA2, pUNASSIGNED)
}

func Test0BA3_0BA4(t *testing.T) {
	ValidateRuneRange(t, 0x0BA3, 0x0BA4, pPVALID)
}

func Test0BA5_0BA7(t *testing.T) {
	ValidateRuneRange(t, 0x0BA5, 0x0BA7, pUNASSIGNED)
}

func Test0BA8_0BAA(t *testing.T) {
	ValidateRuneRange(t, 0x0BA8, 0x0BAA, pPVALID)
}

func Test0BAB_0BAD(t *testing.T) {
	ValidateRuneRange(t, 0x0BAB, 0x0BAD, pUNASSIGNED)
}

func Test0BAE_0BB9(t *testing.T) {
	ValidateRuneRange(t, 0x0BAE, 0x0BB9, pPVALID)
}

func Test0BBA_0BBD(t *testing.T) {
	ValidateRuneRange(t, 0x0BBA, 0x0BBD, pUNASSIGNED)
}

func Test0BBE_0BC2(t *testing.T) {
	ValidateRuneRange(t, 0x0BBE, 0x0BC2, pPVALID)
}

func Test0BC3_0BC5(t *testing.T) {
	ValidateRuneRange(t, 0x0BC3, 0x0BC5, pUNASSIGNED)
}

func Test0BC6_0BC8(t *testing.T) {
	ValidateRuneRange(t, 0x0BC6, 0x0BC8, pPVALID)
}

func Test0BC9(t *testing.T) {
	ValidateRune(t, 0x0BC9, pUNASSIGNED)
}

func Test0BCA_0BCD(t *testing.T) {
	ValidateRuneRange(t, 0x0BCA, 0x0BCD, pPVALID)
}

func Test0BCE_0BCF(t *testing.T) {
	ValidateRuneRange(t, 0x0BCE, 0x0BCF, pUNASSIGNED)
}

func Test0BD0(t *testing.T) {
	ValidateRune(t, 0x0BD0, pPVALID)
}

func Test0BD1_0BD6(t *testing.T) {
	ValidateRuneRange(t, 0x0BD1, 0x0BD6, pUNASSIGNED)
}

func Test0BD7(t *testing.T) {
	ValidateRune(t, 0x0BD7, pPVALID)
}

func Test0BD8_0BE5(t *testing.T) {
	ValidateRuneRange(t, 0x0BD8, 0x0BE5, pUNASSIGNED)
}

func Test0BE6_0BEF(t *testing.T) {
	ValidateRuneRange(t, 0x0BE6, 0x0BEF, pPVALID)
}

func Test0BF0_0BFA(t *testing.T) {
	ValidateRuneRange(t, 0x0BF0, 0x0BFA, pDISALLOWED)
}

func Test0BFB_0C00(t *testing.T) {
	ValidateRuneRange(t, 0x0BFB, 0x0C00, pUNASSIGNED)
}

func Test0C01_0C03(t *testing.T) {
	ValidateRuneRange(t, 0x0C01, 0x0C03, pPVALID)
}

func Test0C04(t *testing.T) {
	ValidateRune(t, 0x0C04, pUNASSIGNED)
}

func Test0C05_0C0C(t *testing.T) {
	ValidateRuneRange(t, 0x0C05, 0x0C0C, pPVALID)
}

func Test0C0D(t *testing.T) {
	ValidateRune(t, 0x0C0D, pUNASSIGNED)
}

func Test0C0E_0C10(t *testing.T) {
	ValidateRuneRange(t, 0x0C0E, 0x0C10, pPVALID)
}

func Test0C11(t *testing.T) {
	ValidateRune(t, 0x0C11, pUNASSIGNED)
}

func Test0C12_0C28(t *testing.T) {
	ValidateRuneRange(t, 0x0C12, 0x0C28, pPVALID)
}

func Test0C29(t *testing.T) {
	ValidateRune(t, 0x0C29, pUNASSIGNED)
}

func Test0C2A_0C33(t *testing.T) {
	ValidateRuneRange(t, 0x0C2A, 0x0C33, pPVALID)
}

func Test0C34(t *testing.T) {
	ValidateRune(t, 0x0C34, pUNASSIGNED)
}

func Test0C35_0C39(t *testing.T) {
	ValidateRuneRange(t, 0x0C35, 0x0C39, pPVALID)
}

func Test0C3A_0C3C(t *testing.T) {
	ValidateRuneRange(t, 0x0C3A, 0x0C3C, pUNASSIGNED)
}

func Test0C3D_0C44(t *testing.T) {
	ValidateRuneRange(t, 0x0C3D, 0x0C44, pPVALID)
}

func Test0C45(t *testing.T) {
	ValidateRune(t, 0x0C45, pUNASSIGNED)
}

func Test0C46_0C48(t *testing.T) {
	ValidateRuneRange(t, 0x0C46, 0x0C48, pPVALID)
}

func Test0C49(t *testing.T) {
	ValidateRune(t, 0x0C49, pUNASSIGNED)
}

func Test0C4A_0C4D(t *testing.T) {
	ValidateRuneRange(t, 0x0C4A, 0x0C4D, pPVALID)
}

func Test0C4E_0C54(t *testing.T) {
	ValidateRuneRange(t, 0x0C4E, 0x0C54, pUNASSIGNED)
}

func Test0C55_0C56(t *testing.T) {
	ValidateRuneRange(t, 0x0C55, 0x0C56, pPVALID)
}

func Test0C57(t *testing.T) {
	ValidateRune(t, 0x0C57, pUNASSIGNED)
}

func Test0C58_0C59(t *testing.T) {
	ValidateRuneRange(t, 0x0C58, 0x0C59, pPVALID)
}

func Test0C5A_0C5F(t *testing.T) {
	ValidateRuneRange(t, 0x0C5A, 0x0C5F, pUNASSIGNED)
}

func Test0C60_0C63(t *testing.T) {
	ValidateRuneRange(t, 0x0C60, 0x0C63, pPVALID)
}

func Test0C64_0C65(t *testing.T) {
	ValidateRuneRange(t, 0x0C64, 0x0C65, pUNASSIGNED)
}

func Test0C66_0C6F(t *testing.T) {
	ValidateRuneRange(t, 0x0C66, 0x0C6F, pPVALID)
}

func Test0C70_0C77(t *testing.T) {
	ValidateRuneRange(t, 0x0C70, 0x0C77, pUNASSIGNED)
}

func Test0C78_0C7F(t *testing.T) {
	ValidateRuneRange(t, 0x0C78, 0x0C7F, pDISALLOWED)
}

func Test0C80_0C81(t *testing.T) {
	ValidateRuneRange(t, 0x0C80, 0x0C81, pUNASSIGNED)
}

func Test0C82_0C83(t *testing.T) {
	ValidateRuneRange(t, 0x0C82, 0x0C83, pPVALID)
}

func Test0C84(t *testing.T) {
	ValidateRune(t, 0x0C84, pUNASSIGNED)
}

func Test0C85_0C8C(t *testing.T) {
	ValidateRuneRange(t, 0x0C85, 0x0C8C, pPVALID)
}

func Test0C8D(t *testing.T) {
	ValidateRune(t, 0x0C8D, pUNASSIGNED)
}

func Test0C8E_0C90(t *testing.T) {
	ValidateRuneRange(t, 0x0C8E, 0x0C90, pPVALID)
}

func Test0C91(t *testing.T) {
	ValidateRune(t, 0x0C91, pUNASSIGNED)
}

func Test0C92_0CA8(t *testing.T) {
	ValidateRuneRange(t, 0x0C92, 0x0CA8, pPVALID)
}

func Test0CA9(t *testing.T) {
	ValidateRune(t, 0x0CA9, pUNASSIGNED)
}

func Test0CAA_0CB3(t *testing.T) {
	ValidateRuneRange(t, 0x0CAA, 0x0CB3, pPVALID)
}

func Test0CB4(t *testing.T) {
	ValidateRune(t, 0x0CB4, pUNASSIGNED)
}

func Test0CB5_0CB9(t *testing.T) {
	ValidateRuneRange(t, 0x0CB5, 0x0CB9, pPVALID)
}

func Test0CBA_0CBB(t *testing.T) {
	ValidateRuneRange(t, 0x0CBA, 0x0CBB, pUNASSIGNED)
}

func Test0CBC_0CC4(t *testing.T) {
	ValidateRuneRange(t, 0x0CBC, 0x0CC4, pPVALID)
}

func Test0CC5(t *testing.T) {
	ValidateRune(t, 0x0CC5, pUNASSIGNED)
}

func Test0CC6_0CC8(t *testing.T) {
	ValidateRuneRange(t, 0x0CC6, 0x0CC8, pPVALID)
}

func Test0CC9(t *testing.T) {
	ValidateRune(t, 0x0CC9, pUNASSIGNED)
}

func Test0CCA_0CCD(t *testing.T) {
	ValidateRuneRange(t, 0x0CCA, 0x0CCD, pPVALID)
}

func Test0CCE_0CD4(t *testing.T) {
	ValidateRuneRange(t, 0x0CCE, 0x0CD4, pUNASSIGNED)
}

func Test0CD5_0CD6(t *testing.T) {
	ValidateRuneRange(t, 0x0CD5, 0x0CD6, pPVALID)
}

func Test0CD7_0CDD(t *testing.T) {
	ValidateRuneRange(t, 0x0CD7, 0x0CDD, pUNASSIGNED)
}

func Test0CDE(t *testing.T) {
	ValidateRune(t, 0x0CDE, pPVALID)
}

func Test0CDF(t *testing.T) {
	ValidateRune(t, 0x0CDF, pUNASSIGNED)
}

func Test0CE0_0CE3(t *testing.T) {
	ValidateRuneRange(t, 0x0CE0, 0x0CE3, pPVALID)
}

func Test0CE4_0CE5(t *testing.T) {
	ValidateRuneRange(t, 0x0CE4, 0x0CE5, pUNASSIGNED)
}

func Test0CE6_0CEF(t *testing.T) {
	ValidateRuneRange(t, 0x0CE6, 0x0CEF, pPVALID)
}

func Test0CF0(t *testing.T) {
	ValidateRune(t, 0x0CF0, pUNASSIGNED)
}

func Test0CF1_0CF2(t *testing.T) {
	ValidateRuneRange(t, 0x0CF1, 0x0CF2, pDISALLOWED)
}

func Test0CF3_0D01(t *testing.T) {
	ValidateRuneRange(t, 0x0CF3, 0x0D01, pUNASSIGNED)
}

func Test0D02_0D03(t *testing.T) {
	ValidateRuneRange(t, 0x0D02, 0x0D03, pPVALID)
}

func Test0D04(t *testing.T) {
	ValidateRune(t, 0x0D04, pUNASSIGNED)
}

func Test0D05_0D0C(t *testing.T) {
	ValidateRuneRange(t, 0x0D05, 0x0D0C, pPVALID)
}

func Test0D0D(t *testing.T) {
	ValidateRune(t, 0x0D0D, pUNASSIGNED)
}

func Test0D0E_0D10(t *testing.T) {
	ValidateRuneRange(t, 0x0D0E, 0x0D10, pPVALID)
}

func Test0D11(t *testing.T) {
	ValidateRune(t, 0x0D11, pUNASSIGNED)
}

func Test0D12_0D28(t *testing.T) {
	ValidateRuneRange(t, 0x0D12, 0x0D28, pPVALID)
}

func Test0D29(t *testing.T) {
	ValidateRune(t, 0x0D29, pUNASSIGNED)
}

func Test0D2A_0D39(t *testing.T) {
	ValidateRuneRange(t, 0x0D2A, 0x0D39, pPVALID)
}

func Test0D3A_0D3C(t *testing.T) {
	ValidateRuneRange(t, 0x0D3A, 0x0D3C, pUNASSIGNED)
}

func Test0D3D_0D44(t *testing.T) {
	ValidateRuneRange(t, 0x0D3D, 0x0D44, pPVALID)
}

func Test0D45(t *testing.T) {
	ValidateRune(t, 0x0D45, pUNASSIGNED)
}

func Test0D46_0D48(t *testing.T) {
	ValidateRuneRange(t, 0x0D46, 0x0D48, pPVALID)
}

func Test0D49(t *testing.T) {
	ValidateRune(t, 0x0D49, pUNASSIGNED)
}

func Test0D4A_0D4D(t *testing.T) {
	ValidateRuneRange(t, 0x0D4A, 0x0D4D, pPVALID)
}

func Test0D4E_0D56(t *testing.T) {
	ValidateRuneRange(t, 0x0D4E, 0x0D56, pUNASSIGNED)
}

func Test0D57(t *testing.T) {
	ValidateRune(t, 0x0D57, pPVALID)
}

func Test0D58_0D5F(t *testing.T) {
	ValidateRuneRange(t, 0x0D58, 0x0D5F, pUNASSIGNED)
}

func Test0D60_0D63(t *testing.T) {
	ValidateRuneRange(t, 0x0D60, 0x0D63, pPVALID)
}

func Test0D64_0D65(t *testing.T) {
	ValidateRuneRange(t, 0x0D64, 0x0D65, pUNASSIGNED)
}

func Test0D66_0D6F(t *testing.T) {
	ValidateRuneRange(t, 0x0D66, 0x0D6F, pPVALID)
}

func Test0D70_0D75(t *testing.T) {
	ValidateRuneRange(t, 0x0D70, 0x0D75, pDISALLOWED)
}

func Test0D76_0D78(t *testing.T) {
	ValidateRuneRange(t, 0x0D76, 0x0D78, pUNASSIGNED)
}

func Test0D79(t *testing.T) {
	ValidateRune(t, 0x0D79, pDISALLOWED)
}

func Test0D7A_0D7F(t *testing.T) {
	ValidateRuneRange(t, 0x0D7A, 0x0D7F, pPVALID)
}

func Test0D80_0D81(t *testing.T) {
	ValidateRuneRange(t, 0x0D80, 0x0D81, pUNASSIGNED)
}

func Test0D82_0D83(t *testing.T) {
	ValidateRuneRange(t, 0x0D82, 0x0D83, pPVALID)
}

func Test0D84(t *testing.T) {
	ValidateRune(t, 0x0D84, pUNASSIGNED)
}

func Test0D85_0D96(t *testing.T) {
	ValidateRuneRange(t, 0x0D85, 0x0D96, pPVALID)
}

func Test0D97_0D99(t *testing.T) {
	ValidateRuneRange(t, 0x0D97, 0x0D99, pUNASSIGNED)
}

func Test0D9A_0DB1(t *testing.T) {
	ValidateRuneRange(t, 0x0D9A, 0x0DB1, pPVALID)
}

func Test0DB2(t *testing.T) {
	ValidateRune(t, 0x0DB2, pUNASSIGNED)
}

func Test0DB3_0DBB(t *testing.T) {
	ValidateRuneRange(t, 0x0DB3, 0x0DBB, pPVALID)
}

func Test0DBC(t *testing.T) {
	ValidateRune(t, 0x0DBC, pUNASSIGNED)
}

func Test0DBD(t *testing.T) {
	ValidateRune(t, 0x0DBD, pPVALID)
}

func Test0DBE_0DBF(t *testing.T) {
	ValidateRuneRange(t, 0x0DBE, 0x0DBF, pUNASSIGNED)
}

func Test0DC0_0DC6(t *testing.T) {
	ValidateRuneRange(t, 0x0DC0, 0x0DC6, pPVALID)
}

func Test0DC7_0DC9(t *testing.T) {
	ValidateRuneRange(t, 0x0DC7, 0x0DC9, pUNASSIGNED)
}

func Test0DCA(t *testing.T) {
	ValidateRune(t, 0x0DCA, pPVALID)
}

func Test0DCB_0DCE(t *testing.T) {
	ValidateRuneRange(t, 0x0DCB, 0x0DCE, pUNASSIGNED)
}

func Test0DCF_0DD4(t *testing.T) {
	ValidateRuneRange(t, 0x0DCF, 0x0DD4, pPVALID)
}

func Test0DD5(t *testing.T) {
	ValidateRune(t, 0x0DD5, pUNASSIGNED)
}

func Test0DD6(t *testing.T) {
	ValidateRune(t, 0x0DD6, pPVALID)
}

func Test0DD7(t *testing.T) {
	ValidateRune(t, 0x0DD7, pUNASSIGNED)
}

func Test0DD8_0DDF(t *testing.T) {
	ValidateRuneRange(t, 0x0DD8, 0x0DDF, pPVALID)
}

func Test0DE0_0DF1(t *testing.T) {
	ValidateRuneRange(t, 0x0DE0, 0x0DF1, pUNASSIGNED)
}

func Test0DF2_0DF3(t *testing.T) {
	ValidateRuneRange(t, 0x0DF2, 0x0DF3, pPVALID)
}

func Test0DF4(t *testing.T) {
	ValidateRune(t, 0x0DF4, pDISALLOWED)
}

func Test0DF5_0E00(t *testing.T) {
	ValidateRuneRange(t, 0x0DF5, 0x0E00, pUNASSIGNED)
}

func Test0E01_0E32(t *testing.T) {
	ValidateRuneRange(t, 0x0E01, 0x0E32, pPVALID)
}

func Test0E33(t *testing.T) {
	ValidateRune(t, 0x0E33, pDISALLOWED)
}

func Test0E34_0E3A(t *testing.T) {
	ValidateRuneRange(t, 0x0E34, 0x0E3A, pPVALID)
}

func Test0E3B_0E3E(t *testing.T) {
	ValidateRuneRange(t, 0x0E3B, 0x0E3E, pUNASSIGNED)
}

func Test0E3F(t *testing.T) {
	ValidateRune(t, 0x0E3F, pDISALLOWED)
}

func Test0E40_0E4E(t *testing.T) {
	ValidateRuneRange(t, 0x0E40, 0x0E4E, pPVALID)
}

func Test0E4F(t *testing.T) {
	ValidateRune(t, 0x0E4F, pDISALLOWED)
}

func Test0E50_0E59(t *testing.T) {
	ValidateRuneRange(t, 0x0E50, 0x0E59, pPVALID)
}

func Test0E5A_0E5B(t *testing.T) {
	ValidateRuneRange(t, 0x0E5A, 0x0E5B, pDISALLOWED)
}

func Test0E5C_0E80(t *testing.T) {
	ValidateRuneRange(t, 0x0E5C, 0x0E80, pUNASSIGNED)
}

func Test0E81_0E82(t *testing.T) {
	ValidateRuneRange(t, 0x0E81, 0x0E82, pPVALID)
}

func Test0E83(t *testing.T) {
	ValidateRune(t, 0x0E83, pUNASSIGNED)
}

func Test0E84(t *testing.T) {
	ValidateRune(t, 0x0E84, pPVALID)
}

func Test0E85_0E86(t *testing.T) {
	ValidateRuneRange(t, 0x0E85, 0x0E86, pUNASSIGNED)
}

func Test0E87_0E88(t *testing.T) {
	ValidateRuneRange(t, 0x0E87, 0x0E88, pPVALID)
}

func Test0E89(t *testing.T) {
	ValidateRune(t, 0x0E89, pUNASSIGNED)
}

func Test0E8A(t *testing.T) {
	ValidateRune(t, 0x0E8A, pPVALID)
}

func Test0E8B_0E8C(t *testing.T) {
	ValidateRuneRange(t, 0x0E8B, 0x0E8C, pUNASSIGNED)
}

func Test0E8D(t *testing.T) {
	ValidateRune(t, 0x0E8D, pPVALID)
}

func Test0E8E_0E93(t *testing.T) {
	ValidateRuneRange(t, 0x0E8E, 0x0E93, pUNASSIGNED)
}

func Test0E94_0E97(t *testing.T) {
	ValidateRuneRange(t, 0x0E94, 0x0E97, pPVALID)
}

func Test0E98(t *testing.T) {
	ValidateRune(t, 0x0E98, pUNASSIGNED)
}

func Test0E99_0E9F(t *testing.T) {
	ValidateRuneRange(t, 0x0E99, 0x0E9F, pPVALID)
}

func Test0EA0(t *testing.T) {
	ValidateRune(t, 0x0EA0, pUNASSIGNED)
}

func Test0EA1_0EA3(t *testing.T) {
	ValidateRuneRange(t, 0x0EA1, 0x0EA3, pPVALID)
}

func Test0EA4(t *testing.T) {
	ValidateRune(t, 0x0EA4, pUNASSIGNED)
}

func Test0EA5(t *testing.T) {
	ValidateRune(t, 0x0EA5, pPVALID)
}

func Test0EA6(t *testing.T) {
	ValidateRune(t, 0x0EA6, pUNASSIGNED)
}

func Test0EA7(t *testing.T) {
	ValidateRune(t, 0x0EA7, pPVALID)
}

func Test0EA8_0EA9(t *testing.T) {
	ValidateRuneRange(t, 0x0EA8, 0x0EA9, pUNASSIGNED)
}

func Test0EAA_0EAB(t *testing.T) {
	ValidateRuneRange(t, 0x0EAA, 0x0EAB, pPVALID)
}

func Test0EAC(t *testing.T) {
	ValidateRune(t, 0x0EAC, pUNASSIGNED)
}

func Test0EAD_0EB2(t *testing.T) {
	ValidateRuneRange(t, 0x0EAD, 0x0EB2, pPVALID)
}

func Test0EB3(t *testing.T) {
	ValidateRune(t, 0x0EB3, pDISALLOWED)
}

func Test0EB4_0EB9(t *testing.T) {
	ValidateRuneRange(t, 0x0EB4, 0x0EB9, pPVALID)
}

func Test0EBA(t *testing.T) {
	ValidateRune(t, 0x0EBA, pUNASSIGNED)
}

func Test0EBB_0EBD(t *testing.T) {
	ValidateRuneRange(t, 0x0EBB, 0x0EBD, pPVALID)
}

func Test0EBE_0EBF(t *testing.T) {
	ValidateRuneRange(t, 0x0EBE, 0x0EBF, pUNASSIGNED)
}

func Test0EC0_0EC4(t *testing.T) {
	ValidateRuneRange(t, 0x0EC0, 0x0EC4, pPVALID)
}

func Test0EC5(t *testing.T) {
	ValidateRune(t, 0x0EC5, pUNASSIGNED)
}

func Test0EC6(t *testing.T) {
	ValidateRune(t, 0x0EC6, pPVALID)
}

func Test0EC7(t *testing.T) {
	ValidateRune(t, 0x0EC7, pUNASSIGNED)
}

func Test0EC8_0ECD(t *testing.T) {
	ValidateRuneRange(t, 0x0EC8, 0x0ECD, pPVALID)
}

func Test0ECE_0ECF(t *testing.T) {
	ValidateRuneRange(t, 0x0ECE, 0x0ECF, pUNASSIGNED)
}

func Test0ED0_0ED9(t *testing.T) {
	ValidateRuneRange(t, 0x0ED0, 0x0ED9, pPVALID)
}

func Test0EDA_0EDB(t *testing.T) {
	ValidateRuneRange(t, 0x0EDA, 0x0EDB, pUNASSIGNED)
}

func Test0EDC_0EDD(t *testing.T) {
	ValidateRuneRange(t, 0x0EDC, 0x0EDD, pDISALLOWED)
}

func Test0EDE_0EFF(t *testing.T) {
	ValidateRuneRange(t, 0x0EDE, 0x0EFF, pUNASSIGNED)
}

func Test0F00(t *testing.T) {
	ValidateRune(t, 0x0F00, pPVALID)
}

func Test0F01_0F0A(t *testing.T) {
	ValidateRuneRange(t, 0x0F01, 0x0F0A, pDISALLOWED)
}

func Test0F0B(t *testing.T) {
	ValidateRune(t, 0x0F0B, pPVALID)
}

func Test0F0C_0F17(t *testing.T) {
	ValidateRuneRange(t, 0x0F0C, 0x0F17, pDISALLOWED)
}

func Test0F18_0F19(t *testing.T) {
	ValidateRuneRange(t, 0x0F18, 0x0F19, pPVALID)
}

func Test0F1A_0F1F(t *testing.T) {
	ValidateRuneRange(t, 0x0F1A, 0x0F1F, pDISALLOWED)
}

func Test0F20_0F29(t *testing.T) {
	ValidateRuneRange(t, 0x0F20, 0x0F29, pPVALID)
}

func Test0F2A_0F34(t *testing.T) {
	ValidateRuneRange(t, 0x0F2A, 0x0F34, pDISALLOWED)
}

func Test0F35(t *testing.T) {
	ValidateRune(t, 0x0F35, pPVALID)
}

func Test0F36(t *testing.T) {
	ValidateRune(t, 0x0F36, pDISALLOWED)
}

func Test0F37(t *testing.T) {
	ValidateRune(t, 0x0F37, pPVALID)
}

func Test0F38(t *testing.T) {
	ValidateRune(t, 0x0F38, pDISALLOWED)
}

func Test0F39(t *testing.T) {
	ValidateRune(t, 0x0F39, pPVALID)
}

func Test0F3A_0F3D(t *testing.T) {
	ValidateRuneRange(t, 0x0F3A, 0x0F3D, pDISALLOWED)
}

func Test0F3E_0F42(t *testing.T) {
	ValidateRuneRange(t, 0x0F3E, 0x0F42, pPVALID)
}

func Test0F43(t *testing.T) {
	ValidateRune(t, 0x0F43, pDISALLOWED)
}

func Test0F44_0F47(t *testing.T) {
	ValidateRuneRange(t, 0x0F44, 0x0F47, pPVALID)
}

func Test0F48(t *testing.T) {
	ValidateRune(t, 0x0F48, pUNASSIGNED)
}

func Test0F49_0F4C(t *testing.T) {
	ValidateRuneRange(t, 0x0F49, 0x0F4C, pPVALID)
}

func Test0F4D(t *testing.T) {
	ValidateRune(t, 0x0F4D, pDISALLOWED)
}

func Test0F4E_0F51(t *testing.T) {
	ValidateRuneRange(t, 0x0F4E, 0x0F51, pPVALID)
}

func Test0F52(t *testing.T) {
	ValidateRune(t, 0x0F52, pDISALLOWED)
}

func Test0F53_0F56(t *testing.T) {
	ValidateRuneRange(t, 0x0F53, 0x0F56, pPVALID)
}

func Test0F57(t *testing.T) {
	ValidateRune(t, 0x0F57, pDISALLOWED)
}

func Test0F58_0F5B(t *testing.T) {
	ValidateRuneRange(t, 0x0F58, 0x0F5B, pPVALID)
}

func Test0F5C(t *testing.T) {
	ValidateRune(t, 0x0F5C, pDISALLOWED)
}

func Test0F5D_0F68(t *testing.T) {
	ValidateRuneRange(t, 0x0F5D, 0x0F68, pPVALID)
}

func Test0F69(t *testing.T) {
	ValidateRune(t, 0x0F69, pDISALLOWED)
}

func Test0F6A_0F6C(t *testing.T) {
	ValidateRuneRange(t, 0x0F6A, 0x0F6C, pPVALID)
}

func Test0F6D_0F70(t *testing.T) {
	ValidateRuneRange(t, 0x0F6D, 0x0F70, pUNASSIGNED)
}

func Test0F71_0F72(t *testing.T) {
	ValidateRuneRange(t, 0x0F71, 0x0F72, pPVALID)
}

func Test0F73(t *testing.T) {
	ValidateRune(t, 0x0F73, pDISALLOWED)
}

func Test0F74(t *testing.T) {
	ValidateRune(t, 0x0F74, pPVALID)
}

func Test0F75_0F79(t *testing.T) {
	ValidateRuneRange(t, 0x0F75, 0x0F79, pDISALLOWED)
}

func Test0F7A_0F80(t *testing.T) {
	ValidateRuneRange(t, 0x0F7A, 0x0F80, pPVALID)
}

func Test0F81(t *testing.T) {
	ValidateRune(t, 0x0F81, pDISALLOWED)
}

func Test0F82_0F84(t *testing.T) {
	ValidateRuneRange(t, 0x0F82, 0x0F84, pPVALID)
}

func Test0F85(t *testing.T) {
	ValidateRune(t, 0x0F85, pDISALLOWED)
}

func Test0F86_0F8B(t *testing.T) {
	ValidateRuneRange(t, 0x0F86, 0x0F8B, pPVALID)
}

func Test0F8C_0F8F(t *testing.T) {
	ValidateRuneRange(t, 0x0F8C, 0x0F8F, pUNASSIGNED)
}

func Test0F90_0F92(t *testing.T) {
	ValidateRuneRange(t, 0x0F90, 0x0F92, pPVALID)
}

func Test0F93(t *testing.T) {
	ValidateRune(t, 0x0F93, pDISALLOWED)
}

func Test0F94_0F97(t *testing.T) {
	ValidateRuneRange(t, 0x0F94, 0x0F97, pPVALID)
}

func Test0F98(t *testing.T) {
	ValidateRune(t, 0x0F98, pUNASSIGNED)
}

func Test0F99_0F9C(t *testing.T) {
	ValidateRuneRange(t, 0x0F99, 0x0F9C, pPVALID)
}

func Test0F9D(t *testing.T) {
	ValidateRune(t, 0x0F9D, pDISALLOWED)
}

func Test0F9E_0FA1(t *testing.T) {
	ValidateRuneRange(t, 0x0F9E, 0x0FA1, pPVALID)
}

func Test0FA2(t *testing.T) {
	ValidateRune(t, 0x0FA2, pDISALLOWED)
}

func Test0FA3_0FA6(t *testing.T) {
	ValidateRuneRange(t, 0x0FA3, 0x0FA6, pPVALID)
}

func Test0FA7(t *testing.T) {
	ValidateRune(t, 0x0FA7, pDISALLOWED)
}

func Test0FA8_0FAB(t *testing.T) {
	ValidateRuneRange(t, 0x0FA8, 0x0FAB, pPVALID)
}

func Test0FAC(t *testing.T) {
	ValidateRune(t, 0x0FAC, pDISALLOWED)
}

func Test0FAD_0FB8(t *testing.T) {
	ValidateRuneRange(t, 0x0FAD, 0x0FB8, pPVALID)
}

func Test0FB9(t *testing.T) {
	ValidateRune(t, 0x0FB9, pDISALLOWED)
}

func Test0FBA_0FBC(t *testing.T) {
	ValidateRuneRange(t, 0x0FBA, 0x0FBC, pPVALID)
}

func Test0FBD(t *testing.T) {
	ValidateRune(t, 0x0FBD, pUNASSIGNED)
}

func Test0FBE_0FC5(t *testing.T) {
	ValidateRuneRange(t, 0x0FBE, 0x0FC5, pDISALLOWED)
}

func Test0FC6(t *testing.T) {
	ValidateRune(t, 0x0FC6, pPVALID)
}

func Test0FC7_0FCC(t *testing.T) {
	ValidateRuneRange(t, 0x0FC7, 0x0FCC, pDISALLOWED)
}

func Test0FCD(t *testing.T) {
	ValidateRune(t, 0x0FCD, pUNASSIGNED)
}

func Test0FCE_0FD8(t *testing.T) {
	ValidateRuneRange(t, 0x0FCE, 0x0FD8, pDISALLOWED)
}

func Test0FD9_0FFF(t *testing.T) {
	ValidateRuneRange(t, 0x0FD9, 0x0FFF, pUNASSIGNED)
}

func Test1000_1049(t *testing.T) {
	ValidateRuneRange(t, 0x1000, 0x1049, pPVALID)
}

func Test104A_104F(t *testing.T) {
	ValidateRuneRange(t, 0x104A, 0x104F, pDISALLOWED)
}

func Test1050_109D(t *testing.T) {
	ValidateRuneRange(t, 0x1050, 0x109D, pPVALID)
}

func Test109E_10C5(t *testing.T) {
	ValidateRuneRange(t, 0x109E, 0x10C5, pDISALLOWED)
}

func Test10C6_10CF(t *testing.T) {
	ValidateRuneRange(t, 0x10C6, 0x10CF, pUNASSIGNED)
}

func Test10D0_10FA(t *testing.T) {
	ValidateRuneRange(t, 0x10D0, 0x10FA, pPVALID)
}

func Test10FB_10FC(t *testing.T) {
	ValidateRuneRange(t, 0x10FB, 0x10FC, pDISALLOWED)
}

func Test10FD_10FF(t *testing.T) {
	ValidateRuneRange(t, 0x10FD, 0x10FF, pUNASSIGNED)
}

func Test1100_11FF(t *testing.T) {
	ValidateRuneRange(t, 0x1100, 0x11FF, pDISALLOWED)
}

func Test1200_1248(t *testing.T) {
	ValidateRuneRange(t, 0x1200, 0x1248, pPVALID)
}

func Test124A_124D(t *testing.T) {
	ValidateRuneRange(t, 0x124A, 0x124D, pPVALID)
}

func Test124E_124F(t *testing.T) {
	ValidateRuneRange(t, 0x124E, 0x124F, pUNASSIGNED)
}

func Test1250_1256(t *testing.T) {
	ValidateRuneRange(t, 0x1250, 0x1256, pPVALID)
}

func Test125A_125D(t *testing.T) {
	ValidateRuneRange(t, 0x125A, 0x125D, pPVALID)
}

func Test125E_125F(t *testing.T) {
	ValidateRuneRange(t, 0x125E, 0x125F, pUNASSIGNED)
}

func Test1260_1288(t *testing.T) {
	ValidateRuneRange(t, 0x1260, 0x1288, pPVALID)
}

func Test128A_128D(t *testing.T) {
	ValidateRuneRange(t, 0x128A, 0x128D, pPVALID)
}

func Test128E_128F(t *testing.T) {
	ValidateRuneRange(t, 0x128E, 0x128F, pUNASSIGNED)
}

func Test1290_12B0(t *testing.T) {
	ValidateRuneRange(t, 0x1290, 0x12B0, pPVALID)
}

func Test12B1(t *testing.T) {
	ValidateRune(t, 0x12B1, pUNASSIGNED)
}

func Test12B2_12B5(t *testing.T) {
	ValidateRuneRange(t, 0x12B2, 0x12B5, pPVALID)
}

func Test12B6_12B7(t *testing.T) {
	ValidateRuneRange(t, 0x12B6, 0x12B7, pUNASSIGNED)
}

func Test12B8_12BE(t *testing.T) {
	ValidateRuneRange(t, 0x12B8, 0x12BE, pPVALID)
}

func Test12BF(t *testing.T) {
	ValidateRune(t, 0x12BF, pUNASSIGNED)
}

func Test12C0(t *testing.T) {
	ValidateRune(t, 0x12C0, pPVALID)
}

func Test12C1(t *testing.T) {
	ValidateRune(t, 0x12C1, pUNASSIGNED)
}

func Test12C2_12C5(t *testing.T) {
	ValidateRuneRange(t, 0x12C2, 0x12C5, pPVALID)
}

func Test12C6_12C7(t *testing.T) {
	ValidateRuneRange(t, 0x12C6, 0x12C7, pUNASSIGNED)
}

func Test12C8_12D6(t *testing.T) {
	ValidateRuneRange(t, 0x12C8, 0x12D6, pPVALID)
}

func Test12D7(t *testing.T) {
	ValidateRune(t, 0x12D7, pUNASSIGNED)
}

func Test12D8_1310(t *testing.T) {
	ValidateRuneRange(t, 0x12D8, 0x1310, pPVALID)
}

func Test1312_1315(t *testing.T) {
	ValidateRuneRange(t, 0x1312, 0x1315, pPVALID)
}

func Test1316_1317(t *testing.T) {
	ValidateRuneRange(t, 0x1316, 0x1317, pUNASSIGNED)
}

func Test1318_135A(t *testing.T) {
	ValidateRuneRange(t, 0x1318, 0x135A, pPVALID)
}

func Test135B_135E(t *testing.T) {
	ValidateRuneRange(t, 0x135B, 0x135E, pUNASSIGNED)
}

func Test135F(t *testing.T) {
	ValidateRune(t, 0x135F, pPVALID)
}

func Test1360_137C(t *testing.T) {
	ValidateRuneRange(t, 0x1360, 0x137C, pDISALLOWED)
}

func Test137D_137F(t *testing.T) {
	ValidateRuneRange(t, 0x137D, 0x137F, pUNASSIGNED)
}

func Test1380_138F(t *testing.T) {
	ValidateRuneRange(t, 0x1380, 0x138F, pPVALID)
}

func Test1390_1399(t *testing.T) {
	ValidateRuneRange(t, 0x1390, 0x1399, pDISALLOWED)
}

func Test139A_139F(t *testing.T) {
	ValidateRuneRange(t, 0x139A, 0x139F, pUNASSIGNED)
}

func Test13A0_13F4(t *testing.T) {
	ValidateRuneRange(t, 0x13A0, 0x13F4, pPVALID)
}

func Test13F5_13FF(t *testing.T) {
	ValidateRuneRange(t, 0x13F5, 0x13FF, pUNASSIGNED)
}

func Test1401_166C(t *testing.T) {
	ValidateRuneRange(t, 0x1401, 0x166C, pPVALID)
}

func Test166D_166E(t *testing.T) {
	ValidateRuneRange(t, 0x166D, 0x166E, pDISALLOWED)
}

func Test166F_167F(t *testing.T) {
	ValidateRuneRange(t, 0x166F, 0x167F, pPVALID)
}

func Test1681_169A(t *testing.T) {
	ValidateRuneRange(t, 0x1681, 0x169A, pPVALID)
}

func Test169B_169C(t *testing.T) {
	ValidateRuneRange(t, 0x169B, 0x169C, pDISALLOWED)
}

func Test169D_169F(t *testing.T) {
	ValidateRuneRange(t, 0x169D, 0x169F, pUNASSIGNED)
}

func Test16A0_16EA(t *testing.T) {
	ValidateRuneRange(t, 0x16A0, 0x16EA, pPVALID)
}

func Test16EB_16F0(t *testing.T) {
	ValidateRuneRange(t, 0x16EB, 0x16F0, pDISALLOWED)
}

func Test16F1_16FF(t *testing.T) {
	ValidateRuneRange(t, 0x16F1, 0x16FF, pUNASSIGNED)
}

func Test1700_170C(t *testing.T) {
	ValidateRuneRange(t, 0x1700, 0x170C, pPVALID)
}

func Test170D(t *testing.T) {
	ValidateRune(t, 0x170D, pUNASSIGNED)
}

func Test170E_1714(t *testing.T) {
	ValidateRuneRange(t, 0x170E, 0x1714, pPVALID)
}

func Test1715_171F(t *testing.T) {
	ValidateRuneRange(t, 0x1715, 0x171F, pUNASSIGNED)
}

func Test1720_1734(t *testing.T) {
	ValidateRuneRange(t, 0x1720, 0x1734, pPVALID)
}

func Test1735_1736(t *testing.T) {
	ValidateRuneRange(t, 0x1735, 0x1736, pDISALLOWED)
}

func Test1737_173F(t *testing.T) {
	ValidateRuneRange(t, 0x1737, 0x173F, pUNASSIGNED)
}

func Test1740_1753(t *testing.T) {
	ValidateRuneRange(t, 0x1740, 0x1753, pPVALID)
}

func Test1754_175F(t *testing.T) {
	ValidateRuneRange(t, 0x1754, 0x175F, pUNASSIGNED)
}

func Test1760_176C(t *testing.T) {
	ValidateRuneRange(t, 0x1760, 0x176C, pPVALID)
}

func Test176D(t *testing.T) {
	ValidateRune(t, 0x176D, pUNASSIGNED)
}

func Test176E_1770(t *testing.T) {
	ValidateRuneRange(t, 0x176E, 0x1770, pPVALID)
}

func Test1772_1773(t *testing.T) {
	ValidateRuneRange(t, 0x1772, 0x1773, pPVALID)
}

func Test1774_177F(t *testing.T) {
	ValidateRuneRange(t, 0x1774, 0x177F, pUNASSIGNED)
}

func Test1780_17B3(t *testing.T) {
	ValidateRuneRange(t, 0x1780, 0x17B3, pPVALID)
}

func Test17B4_17B5(t *testing.T) {
	ValidateRuneRange(t, 0x17B4, 0x17B5, pDISALLOWED)
}

func Test17B6_17D3(t *testing.T) {
	ValidateRuneRange(t, 0x17B6, 0x17D3, pPVALID)
}

func Test17D4_17D6(t *testing.T) {
	ValidateRuneRange(t, 0x17D4, 0x17D6, pDISALLOWED)
}

func Test17D7(t *testing.T) {
	ValidateRune(t, 0x17D7, pPVALID)
}

func Test17D8_17DB(t *testing.T) {
	ValidateRuneRange(t, 0x17D8, 0x17DB, pDISALLOWED)
}

func Test17DC_17DD(t *testing.T) {
	ValidateRuneRange(t, 0x17DC, 0x17DD, pPVALID)
}

func Test17DE_17DF(t *testing.T) {
	ValidateRuneRange(t, 0x17DE, 0x17DF, pUNASSIGNED)
}

func Test17E0_17E9(t *testing.T) {
	ValidateRuneRange(t, 0x17E0, 0x17E9, pPVALID)
}

func Test17EA_17EF(t *testing.T) {
	ValidateRuneRange(t, 0x17EA, 0x17EF, pUNASSIGNED)
}

func Test17F0_17F9(t *testing.T) {
	ValidateRuneRange(t, 0x17F0, 0x17F9, pDISALLOWED)
}

func Test17FA_17FF(t *testing.T) {
	ValidateRuneRange(t, 0x17FA, 0x17FF, pUNASSIGNED)
}

func Test1800_180E(t *testing.T) {
	ValidateRuneRange(t, 0x1800, 0x180E, pDISALLOWED)
}

func Test180F(t *testing.T) {
	ValidateRune(t, 0x180F, pUNASSIGNED)
}

func Test1810_1819(t *testing.T) {
	ValidateRuneRange(t, 0x1810, 0x1819, pPVALID)
}

func Test181A_181F(t *testing.T) {
	ValidateRuneRange(t, 0x181A, 0x181F, pUNASSIGNED)
}

func Test1820_1877(t *testing.T) {
	ValidateRuneRange(t, 0x1820, 0x1877, pPVALID)
}

func Test1878_187F(t *testing.T) {
	ValidateRuneRange(t, 0x1878, 0x187F, pUNASSIGNED)
}

func Test1880_18AA(t *testing.T) {
	ValidateRuneRange(t, 0x1880, 0x18AA, pPVALID)
}

func Test18AB_18AF(t *testing.T) {
	ValidateRuneRange(t, 0x18AB, 0x18AF, pUNASSIGNED)
}

func Test18B0_18F5(t *testing.T) {
	ValidateRuneRange(t, 0x18B0, 0x18F5, pPVALID)
}

func Test18F6_18FF(t *testing.T) {
	ValidateRuneRange(t, 0x18F6, 0x18FF, pUNASSIGNED)
}

func Test1900_191C(t *testing.T) {
	ValidateRuneRange(t, 0x1900, 0x191C, pPVALID)
}

func Test191D_191F(t *testing.T) {
	ValidateRuneRange(t, 0x191D, 0x191F, pUNASSIGNED)
}

func Test1920_192B(t *testing.T) {
	ValidateRuneRange(t, 0x1920, 0x192B, pPVALID)
}

func Test192C_192F(t *testing.T) {
	ValidateRuneRange(t, 0x192C, 0x192F, pUNASSIGNED)
}

func Test1930_193B(t *testing.T) {
	ValidateRuneRange(t, 0x1930, 0x193B, pPVALID)
}

func Test193C_193F(t *testing.T) {
	ValidateRuneRange(t, 0x193C, 0x193F, pUNASSIGNED)
}

func Test1941_1943(t *testing.T) {
	ValidateRuneRange(t, 0x1941, 0x1943, pUNASSIGNED)
}

func Test1944_1945(t *testing.T) {
	ValidateRuneRange(t, 0x1944, 0x1945, pDISALLOWED)
}

func Test1946_196D(t *testing.T) {
	ValidateRuneRange(t, 0x1946, 0x196D, pPVALID)
}

func Test196E_196F(t *testing.T) {
	ValidateRuneRange(t, 0x196E, 0x196F, pUNASSIGNED)
}

func Test1970_1974(t *testing.T) {
	ValidateRuneRange(t, 0x1970, 0x1974, pPVALID)
}

func Test1975_197F(t *testing.T) {
	ValidateRuneRange(t, 0x1975, 0x197F, pUNASSIGNED)
}

func Test1980_19AB(t *testing.T) {
	ValidateRuneRange(t, 0x1980, 0x19AB, pPVALID)
}

func Test19AC_19AF(t *testing.T) {
	ValidateRuneRange(t, 0x19AC, 0x19AF, pUNASSIGNED)
}

func Test19B0_19C9(t *testing.T) {
	ValidateRuneRange(t, 0x19B0, 0x19C9, pPVALID)
}

func Test19CA_19CF(t *testing.T) {
	ValidateRuneRange(t, 0x19CA, 0x19CF, pUNASSIGNED)
}

func Test19D0_19DA(t *testing.T) {
	ValidateRuneRange(t, 0x19D0, 0x19DA, pPVALID)
}

func Test19DB_19DD(t *testing.T) {
	ValidateRuneRange(t, 0x19DB, 0x19DD, pUNASSIGNED)
}

func Test19DE_19FF(t *testing.T) {
	ValidateRuneRange(t, 0x19DE, 0x19FF, pDISALLOWED)
}

func Test1A00_1A1B(t *testing.T) {
	ValidateRuneRange(t, 0x1A00, 0x1A1B, pPVALID)
}

func Test1A1C_1A1D(t *testing.T) {
	ValidateRuneRange(t, 0x1A1C, 0x1A1D, pUNASSIGNED)
}

func Test1A1E_1A1F(t *testing.T) {
	ValidateRuneRange(t, 0x1A1E, 0x1A1F, pDISALLOWED)
}

func Test1A5F(t *testing.T) {
	ValidateRune(t, 0x1A5F, pUNASSIGNED)
}

func Test1A60_1A7C(t *testing.T) {
	ValidateRuneRange(t, 0x1A60, 0x1A7C, pPVALID)
}

func Test1A7D_1A7E(t *testing.T) {
	ValidateRuneRange(t, 0x1A7D, 0x1A7E, pUNASSIGNED)
}

func Test1A7F_1A89(t *testing.T) {
	ValidateRuneRange(t, 0x1A7F, 0x1A89, pPVALID)
}

func Test1A8A_1A8F(t *testing.T) {
	ValidateRuneRange(t, 0x1A8A, 0x1A8F, pUNASSIGNED)
}

func Test1A90_1A99(t *testing.T) {
	ValidateRuneRange(t, 0x1A90, 0x1A99, pPVALID)
}

func Test1A9A_1A9F(t *testing.T) {
	ValidateRuneRange(t, 0x1A9A, 0x1A9F, pUNASSIGNED)
}

func Test1AA0_1AA6(t *testing.T) {
	ValidateRuneRange(t, 0x1AA0, 0x1AA6, pDISALLOWED)
}

func Test1AA7(t *testing.T) {
	ValidateRune(t, 0x1AA7, pPVALID)
}

func Test1AA8_1AAD(t *testing.T) {
	ValidateRuneRange(t, 0x1AA8, 0x1AAD, pDISALLOWED)
}

func Test1AAE_1AFF(t *testing.T) {
	ValidateRuneRange(t, 0x1AAE, 0x1AFF, pUNASSIGNED)
}

func Test1B00_1B4B(t *testing.T) {
	ValidateRuneRange(t, 0x1B00, 0x1B4B, pPVALID)
}

func Test1B4C_1B4F(t *testing.T) {
	ValidateRuneRange(t, 0x1B4C, 0x1B4F, pUNASSIGNED)
}

func Test1B50_1B59(t *testing.T) {
	ValidateRuneRange(t, 0x1B50, 0x1B59, pPVALID)
}

func Test1B5A_1B6A(t *testing.T) {
	ValidateRuneRange(t, 0x1B5A, 0x1B6A, pDISALLOWED)
}

func Test1B6B_1B73(t *testing.T) {
	ValidateRuneRange(t, 0x1B6B, 0x1B73, pPVALID)
}

func Test1B74_1B7C(t *testing.T) {
	ValidateRuneRange(t, 0x1B74, 0x1B7C, pDISALLOWED)
}

func Test1B7D_1B7F(t *testing.T) {
	ValidateRuneRange(t, 0x1B7D, 0x1B7F, pUNASSIGNED)
}

func Test1B80_1BAA(t *testing.T) {
	ValidateRuneRange(t, 0x1B80, 0x1BAA, pPVALID)
}

func Test1BAB_1BAD(t *testing.T) {
	ValidateRuneRange(t, 0x1BAB, 0x1BAD, pUNASSIGNED)
}

func Test1BAE_1BB9(t *testing.T) {
	ValidateRuneRange(t, 0x1BAE, 0x1BB9, pPVALID)
}

func Test1BBA_1BFF(t *testing.T) {
	ValidateRuneRange(t, 0x1BBA, 0x1BFF, pUNASSIGNED)
}

func Test1C00_1C37(t *testing.T) {
	ValidateRuneRange(t, 0x1C00, 0x1C37, pPVALID)
}

func Test1C38_1C3A(t *testing.T) {
	ValidateRuneRange(t, 0x1C38, 0x1C3A, pUNASSIGNED)
}

func Test1C3B_1C3F(t *testing.T) {
	ValidateRuneRange(t, 0x1C3B, 0x1C3F, pDISALLOWED)
}

func Test1C40_1C49(t *testing.T) {
	ValidateRuneRange(t, 0x1C40, 0x1C49, pPVALID)
}

func Test1C4A_1C4C(t *testing.T) {
	ValidateRuneRange(t, 0x1C4A, 0x1C4C, pUNASSIGNED)
}

func Test1C4D_1C7D(t *testing.T) {
	ValidateRuneRange(t, 0x1C4D, 0x1C7D, pPVALID)
}

func Test1C7E_1C7F(t *testing.T) {
	ValidateRuneRange(t, 0x1C7E, 0x1C7F, pDISALLOWED)
}

func Test1C80_1CCF(t *testing.T) {
	ValidateRuneRange(t, 0x1C80, 0x1CCF, pUNASSIGNED)
}

func Test1CD0_1CD2(t *testing.T) {
	ValidateRuneRange(t, 0x1CD0, 0x1CD2, pPVALID)
}

func Test1CD3(t *testing.T) {
	ValidateRune(t, 0x1CD3, pDISALLOWED)
}

func Test1CD4_1CF2(t *testing.T) {
	ValidateRuneRange(t, 0x1CD4, 0x1CF2, pPVALID)
}

func Test1CF3_1CFF(t *testing.T) {
	ValidateRuneRange(t, 0x1CF3, 0x1CFF, pUNASSIGNED)
}

func Test1D00_1D2B(t *testing.T) {
	ValidateRuneRange(t, 0x1D00, 0x1D2B, pPVALID)
}

func Test1D2C_1D2E(t *testing.T) {
	ValidateRuneRange(t, 0x1D2C, 0x1D2E, pDISALLOWED)
}

func Test1D2F(t *testing.T) {
	ValidateRune(t, 0x1D2F, pPVALID)
}

func Test1D30_1D3A(t *testing.T) {
	ValidateRuneRange(t, 0x1D30, 0x1D3A, pDISALLOWED)
}

func Test1D3B(t *testing.T) {
	ValidateRune(t, 0x1D3B, pPVALID)
}

func Test1D3C_1D4D(t *testing.T) {
	ValidateRuneRange(t, 0x1D3C, 0x1D4D, pDISALLOWED)
}

func Test1D4E(t *testing.T) {
	ValidateRune(t, 0x1D4E, pPVALID)
}

func Test1D4F_1D6A(t *testing.T) {
	ValidateRuneRange(t, 0x1D4F, 0x1D6A, pDISALLOWED)
}

func Test1D6B_1D77(t *testing.T) {
	ValidateRuneRange(t, 0x1D6B, 0x1D77, pPVALID)
}

func Test1D78(t *testing.T) {
	ValidateRune(t, 0x1D78, pDISALLOWED)
}

func Test1D79_1D9A(t *testing.T) {
	ValidateRuneRange(t, 0x1D79, 0x1D9A, pPVALID)
}

func Test1D9B_1DBF(t *testing.T) {
	ValidateRuneRange(t, 0x1D9B, 0x1DBF, pDISALLOWED)
}

func Test1DC0_1DE6(t *testing.T) {
	ValidateRuneRange(t, 0x1DC0, 0x1DE6, pPVALID)
}

func Test1DE7_1DFC(t *testing.T) {
	ValidateRuneRange(t, 0x1DE7, 0x1DFC, pUNASSIGNED)
}

func Test1DFD_1DFF(t *testing.T) {
	ValidateRuneRange(t, 0x1DFD, 0x1DFF, pPVALID)
}

func Test1E00(t *testing.T) {
	ValidateRune(t, 0x1E00, pDISALLOWED)
}

func Test1E01(t *testing.T) {
	ValidateRune(t, 0x1E01, pPVALID)
}

func Test1E02(t *testing.T) {
	ValidateRune(t, 0x1E02, pDISALLOWED)
}

func Test1E03(t *testing.T) {
	ValidateRune(t, 0x1E03, pPVALID)
}

func Test1E04(t *testing.T) {
	ValidateRune(t, 0x1E04, pDISALLOWED)
}

func Test1E05(t *testing.T) {
	ValidateRune(t, 0x1E05, pPVALID)
}

func Test1E06(t *testing.T) {
	ValidateRune(t, 0x1E06, pDISALLOWED)
}

func Test1E07(t *testing.T) {
	ValidateRune(t, 0x1E07, pPVALID)
}

func Test1E08(t *testing.T) {
	ValidateRune(t, 0x1E08, pDISALLOWED)
}

func Test1E09(t *testing.T) {
	ValidateRune(t, 0x1E09, pPVALID)
}

func Test1E0A(t *testing.T) {
	ValidateRune(t, 0x1E0A, pDISALLOWED)
}

func Test1E0B(t *testing.T) {
	ValidateRune(t, 0x1E0B, pPVALID)
}

func Test1E0C(t *testing.T) {
	ValidateRune(t, 0x1E0C, pDISALLOWED)
}

func Test1E0D(t *testing.T) {
	ValidateRune(t, 0x1E0D, pPVALID)
}

func Test1E0E(t *testing.T) {
	ValidateRune(t, 0x1E0E, pDISALLOWED)
}

func Test1E0F(t *testing.T) {
	ValidateRune(t, 0x1E0F, pPVALID)
}

func Test1E10(t *testing.T) {
	ValidateRune(t, 0x1E10, pDISALLOWED)
}

func Test1E11(t *testing.T) {
	ValidateRune(t, 0x1E11, pPVALID)
}

func Test1E12(t *testing.T) {
	ValidateRune(t, 0x1E12, pDISALLOWED)
}

func Test1E13(t *testing.T) {
	ValidateRune(t, 0x1E13, pPVALID)
}

func Test1E14(t *testing.T) {
	ValidateRune(t, 0x1E14, pDISALLOWED)
}

func Test1E15(t *testing.T) {
	ValidateRune(t, 0x1E15, pPVALID)
}

func Test1E16(t *testing.T) {
	ValidateRune(t, 0x1E16, pDISALLOWED)
}

func Test1E17(t *testing.T) {
	ValidateRune(t, 0x1E17, pPVALID)
}

func Test1E18(t *testing.T) {
	ValidateRune(t, 0x1E18, pDISALLOWED)
}

func Test1E19(t *testing.T) {
	ValidateRune(t, 0x1E19, pPVALID)
}

func Test1E1A(t *testing.T) {
	ValidateRune(t, 0x1E1A, pDISALLOWED)
}

func Test1E1B(t *testing.T) {
	ValidateRune(t, 0x1E1B, pPVALID)
}

func Test1E1C(t *testing.T) {
	ValidateRune(t, 0x1E1C, pDISALLOWED)
}

func Test1E1D(t *testing.T) {
	ValidateRune(t, 0x1E1D, pPVALID)
}

func Test1E1E(t *testing.T) {
	ValidateRune(t, 0x1E1E, pDISALLOWED)
}

func Test1E1F(t *testing.T) {
	ValidateRune(t, 0x1E1F, pPVALID)
}

func Test1E20(t *testing.T) {
	ValidateRune(t, 0x1E20, pDISALLOWED)
}

func Test1E21(t *testing.T) {
	ValidateRune(t, 0x1E21, pPVALID)
}

func Test1E22(t *testing.T) {
	ValidateRune(t, 0x1E22, pDISALLOWED)
}

func Test1E23(t *testing.T) {
	ValidateRune(t, 0x1E23, pPVALID)
}

func Test1E24(t *testing.T) {
	ValidateRune(t, 0x1E24, pDISALLOWED)
}

func Test1E25(t *testing.T) {
	ValidateRune(t, 0x1E25, pPVALID)
}

func Test1E26(t *testing.T) {
	ValidateRune(t, 0x1E26, pDISALLOWED)
}

func Test1E27(t *testing.T) {
	ValidateRune(t, 0x1E27, pPVALID)
}

func Test1E28(t *testing.T) {
	ValidateRune(t, 0x1E28, pDISALLOWED)
}

func Test1E29(t *testing.T) {
	ValidateRune(t, 0x1E29, pPVALID)
}

func Test1E2A(t *testing.T) {
	ValidateRune(t, 0x1E2A, pDISALLOWED)
}

func Test1E2B(t *testing.T) {
	ValidateRune(t, 0x1E2B, pPVALID)
}

func Test1E2C(t *testing.T) {
	ValidateRune(t, 0x1E2C, pDISALLOWED)
}

func Test1E2D(t *testing.T) {
	ValidateRune(t, 0x1E2D, pPVALID)
}

func Test1E2E(t *testing.T) {
	ValidateRune(t, 0x1E2E, pDISALLOWED)
}

func Test1E2F(t *testing.T) {
	ValidateRune(t, 0x1E2F, pPVALID)
}

func Test1E30(t *testing.T) {
	ValidateRune(t, 0x1E30, pDISALLOWED)
}

func Test1E31(t *testing.T) {
	ValidateRune(t, 0x1E31, pPVALID)
}

func Test1E32(t *testing.T) {
	ValidateRune(t, 0x1E32, pDISALLOWED)
}

func Test1E33(t *testing.T) {
	ValidateRune(t, 0x1E33, pPVALID)
}

func Test1E34(t *testing.T) {
	ValidateRune(t, 0x1E34, pDISALLOWED)
}

func Test1E35(t *testing.T) {
	ValidateRune(t, 0x1E35, pPVALID)
}

func Test1E36(t *testing.T) {
	ValidateRune(t, 0x1E36, pDISALLOWED)
}

func Test1E37(t *testing.T) {
	ValidateRune(t, 0x1E37, pPVALID)
}

func Test1E38(t *testing.T) {
	ValidateRune(t, 0x1E38, pDISALLOWED)
}

func Test1E39(t *testing.T) {
	ValidateRune(t, 0x1E39, pPVALID)
}

func Test1E3A(t *testing.T) {
	ValidateRune(t, 0x1E3A, pDISALLOWED)
}

func Test1E3B(t *testing.T) {
	ValidateRune(t, 0x1E3B, pPVALID)
}

func Test1E3C(t *testing.T) {
	ValidateRune(t, 0x1E3C, pDISALLOWED)
}

func Test1E3D(t *testing.T) {
	ValidateRune(t, 0x1E3D, pPVALID)
}

func Test1E3E(t *testing.T) {
	ValidateRune(t, 0x1E3E, pDISALLOWED)
}

func Test1E3F(t *testing.T) {
	ValidateRune(t, 0x1E3F, pPVALID)
}

func Test1E40(t *testing.T) {
	ValidateRune(t, 0x1E40, pDISALLOWED)
}

func Test1E41(t *testing.T) {
	ValidateRune(t, 0x1E41, pPVALID)
}

func Test1E42(t *testing.T) {
	ValidateRune(t, 0x1E42, pDISALLOWED)
}

func Test1E43(t *testing.T) {
	ValidateRune(t, 0x1E43, pPVALID)
}

func Test1E44(t *testing.T) {
	ValidateRune(t, 0x1E44, pDISALLOWED)
}

func Test1E45(t *testing.T) {
	ValidateRune(t, 0x1E45, pPVALID)
}

func Test1E46(t *testing.T) {
	ValidateRune(t, 0x1E46, pDISALLOWED)
}

func Test1E47(t *testing.T) {
	ValidateRune(t, 0x1E47, pPVALID)
}

func Test1E48(t *testing.T) {
	ValidateRune(t, 0x1E48, pDISALLOWED)
}

func Test1E49(t *testing.T) {
	ValidateRune(t, 0x1E49, pPVALID)
}

func Test1E4A(t *testing.T) {
	ValidateRune(t, 0x1E4A, pDISALLOWED)
}

func Test1E4B(t *testing.T) {
	ValidateRune(t, 0x1E4B, pPVALID)
}

func Test1E4C(t *testing.T) {
	ValidateRune(t, 0x1E4C, pDISALLOWED)
}

func Test1E4D(t *testing.T) {
	ValidateRune(t, 0x1E4D, pPVALID)
}

func Test1E4E(t *testing.T) {
	ValidateRune(t, 0x1E4E, pDISALLOWED)
}

func Test1E4F(t *testing.T) {
	ValidateRune(t, 0x1E4F, pPVALID)
}

func Test1E50(t *testing.T) {
	ValidateRune(t, 0x1E50, pDISALLOWED)
}

func Test1E51(t *testing.T) {
	ValidateRune(t, 0x1E51, pPVALID)
}

func Test1E52(t *testing.T) {
	ValidateRune(t, 0x1E52, pDISALLOWED)
}

func Test1E53(t *testing.T) {
	ValidateRune(t, 0x1E53, pPVALID)
}

func Test1E54(t *testing.T) {
	ValidateRune(t, 0x1E54, pDISALLOWED)
}

func Test1E55(t *testing.T) {
	ValidateRune(t, 0x1E55, pPVALID)
}

func Test1E56(t *testing.T) {
	ValidateRune(t, 0x1E56, pDISALLOWED)
}

func Test1E57(t *testing.T) {
	ValidateRune(t, 0x1E57, pPVALID)
}

func Test1E58(t *testing.T) {
	ValidateRune(t, 0x1E58, pDISALLOWED)
}

func Test1E59(t *testing.T) {
	ValidateRune(t, 0x1E59, pPVALID)
}

func Test1E5A(t *testing.T) {
	ValidateRune(t, 0x1E5A, pDISALLOWED)
}

func Test1E5B(t *testing.T) {
	ValidateRune(t, 0x1E5B, pPVALID)
}

func Test1E5C(t *testing.T) {
	ValidateRune(t, 0x1E5C, pDISALLOWED)
}

func Test1E5D(t *testing.T) {
	ValidateRune(t, 0x1E5D, pPVALID)
}

func Test1E5E(t *testing.T) {
	ValidateRune(t, 0x1E5E, pDISALLOWED)
}

func Test1E5F(t *testing.T) {
	ValidateRune(t, 0x1E5F, pPVALID)
}

func Test1E60(t *testing.T) {
	ValidateRune(t, 0x1E60, pDISALLOWED)
}

func Test1E61(t *testing.T) {
	ValidateRune(t, 0x1E61, pPVALID)
}

func Test1E62(t *testing.T) {
	ValidateRune(t, 0x1E62, pDISALLOWED)
}

func Test1E63(t *testing.T) {
	ValidateRune(t, 0x1E63, pPVALID)
}

func Test1E64(t *testing.T) {
	ValidateRune(t, 0x1E64, pDISALLOWED)
}

func Test1E65(t *testing.T) {
	ValidateRune(t, 0x1E65, pPVALID)
}

func Test1E66(t *testing.T) {
	ValidateRune(t, 0x1E66, pDISALLOWED)
}

func Test1E67(t *testing.T) {
	ValidateRune(t, 0x1E67, pPVALID)
}

func Test1E68(t *testing.T) {
	ValidateRune(t, 0x1E68, pDISALLOWED)
}

func Test1E69(t *testing.T) {
	ValidateRune(t, 0x1E69, pPVALID)
}

func Test1E6A(t *testing.T) {
	ValidateRune(t, 0x1E6A, pDISALLOWED)
}

func Test1E6B(t *testing.T) {
	ValidateRune(t, 0x1E6B, pPVALID)
}

func Test1E6C(t *testing.T) {
	ValidateRune(t, 0x1E6C, pDISALLOWED)
}

func Test1E6D(t *testing.T) {
	ValidateRune(t, 0x1E6D, pPVALID)
}

func Test1E6E(t *testing.T) {
	ValidateRune(t, 0x1E6E, pDISALLOWED)
}

func Test1E6F(t *testing.T) {
	ValidateRune(t, 0x1E6F, pPVALID)
}

func Test1E70(t *testing.T) {
	ValidateRune(t, 0x1E70, pDISALLOWED)
}

func Test1E71(t *testing.T) {
	ValidateRune(t, 0x1E71, pPVALID)
}

func Test1E72(t *testing.T) {
	ValidateRune(t, 0x1E72, pDISALLOWED)
}

func Test1E73(t *testing.T) {
	ValidateRune(t, 0x1E73, pPVALID)
}

func Test1E74(t *testing.T) {
	ValidateRune(t, 0x1E74, pDISALLOWED)
}

func Test1E75(t *testing.T) {
	ValidateRune(t, 0x1E75, pPVALID)
}

func Test1E76(t *testing.T) {
	ValidateRune(t, 0x1E76, pDISALLOWED)
}

func Test1E77(t *testing.T) {
	ValidateRune(t, 0x1E77, pPVALID)
}

func Test1E78(t *testing.T) {
	ValidateRune(t, 0x1E78, pDISALLOWED)
}

func Test1E79(t *testing.T) {
	ValidateRune(t, 0x1E79, pPVALID)
}

func Test1E7A(t *testing.T) {
	ValidateRune(t, 0x1E7A, pDISALLOWED)
}

func Test1E7B(t *testing.T) {
	ValidateRune(t, 0x1E7B, pPVALID)
}

func Test1E7C(t *testing.T) {
	ValidateRune(t, 0x1E7C, pDISALLOWED)
}

func Test1E7D(t *testing.T) {
	ValidateRune(t, 0x1E7D, pPVALID)
}

func Test1E7E(t *testing.T) {
	ValidateRune(t, 0x1E7E, pDISALLOWED)
}

func Test1E7F(t *testing.T) {
	ValidateRune(t, 0x1E7F, pPVALID)
}

func Test1E80(t *testing.T) {
	ValidateRune(t, 0x1E80, pDISALLOWED)
}

func Test1E81(t *testing.T) {
	ValidateRune(t, 0x1E81, pPVALID)
}

func Test1E82(t *testing.T) {
	ValidateRune(t, 0x1E82, pDISALLOWED)
}

func Test1E83(t *testing.T) {
	ValidateRune(t, 0x1E83, pPVALID)
}

func Test1E84(t *testing.T) {
	ValidateRune(t, 0x1E84, pDISALLOWED)
}

func Test1E85(t *testing.T) {
	ValidateRune(t, 0x1E85, pPVALID)
}

func Test1E86(t *testing.T) {
	ValidateRune(t, 0x1E86, pDISALLOWED)
}

func Test1E87(t *testing.T) {
	ValidateRune(t, 0x1E87, pPVALID)
}

func Test1E88(t *testing.T) {
	ValidateRune(t, 0x1E88, pDISALLOWED)
}

func Test1E89(t *testing.T) {
	ValidateRune(t, 0x1E89, pPVALID)
}

func Test1E8A(t *testing.T) {
	ValidateRune(t, 0x1E8A, pDISALLOWED)
}

func Test1E8B(t *testing.T) {
	ValidateRune(t, 0x1E8B, pPVALID)
}

func Test1E8C(t *testing.T) {
	ValidateRune(t, 0x1E8C, pDISALLOWED)
}

func Test1E8D(t *testing.T) {
	ValidateRune(t, 0x1E8D, pPVALID)
}

func Test1E8E(t *testing.T) {
	ValidateRune(t, 0x1E8E, pDISALLOWED)
}

func Test1E8F(t *testing.T) {
	ValidateRune(t, 0x1E8F, pPVALID)
}

func Test1E90(t *testing.T) {
	ValidateRune(t, 0x1E90, pDISALLOWED)
}

func Test1E91(t *testing.T) {
	ValidateRune(t, 0x1E91, pPVALID)
}

func Test1E92(t *testing.T) {
	ValidateRune(t, 0x1E92, pDISALLOWED)
}

func Test1E93(t *testing.T) {
	ValidateRune(t, 0x1E93, pPVALID)
}

func Test1E94(t *testing.T) {
	ValidateRune(t, 0x1E94, pDISALLOWED)
}

func Test1E95_1E99(t *testing.T) {
	ValidateRuneRange(t, 0x1E95, 0x1E99, pPVALID)
}

func Test1E9A_1E9B(t *testing.T) {
	ValidateRuneRange(t, 0x1E9A, 0x1E9B, pDISALLOWED)
}

func Test1E9C_1E9D(t *testing.T) {
	ValidateRuneRange(t, 0x1E9C, 0x1E9D, pPVALID)
}

func Test1E9E(t *testing.T) {
	ValidateRune(t, 0x1E9E, pDISALLOWED)
}

func Test1E9F(t *testing.T) {
	ValidateRune(t, 0x1E9F, pPVALID)
}

func Test1EA0(t *testing.T) {
	ValidateRune(t, 0x1EA0, pDISALLOWED)
}

func Test1EA1(t *testing.T) {
	ValidateRune(t, 0x1EA1, pPVALID)
}

func Test1EA2(t *testing.T) {
	ValidateRune(t, 0x1EA2, pDISALLOWED)
}

func Test1EA3(t *testing.T) {
	ValidateRune(t, 0x1EA3, pPVALID)
}

func Test1EA4(t *testing.T) {
	ValidateRune(t, 0x1EA4, pDISALLOWED)
}

func Test1EA5(t *testing.T) {
	ValidateRune(t, 0x1EA5, pPVALID)
}

func Test1EA6(t *testing.T) {
	ValidateRune(t, 0x1EA6, pDISALLOWED)
}

func Test1EA7(t *testing.T) {
	ValidateRune(t, 0x1EA7, pPVALID)
}

func Test1EA8(t *testing.T) {
	ValidateRune(t, 0x1EA8, pDISALLOWED)
}

func Test1EA9(t *testing.T) {
	ValidateRune(t, 0x1EA9, pPVALID)
}

func Test1EAA(t *testing.T) {
	ValidateRune(t, 0x1EAA, pDISALLOWED)
}

func Test1EAB(t *testing.T) {
	ValidateRune(t, 0x1EAB, pPVALID)
}

func Test1EAC(t *testing.T) {
	ValidateRune(t, 0x1EAC, pDISALLOWED)
}

func Test1EAD(t *testing.T) {
	ValidateRune(t, 0x1EAD, pPVALID)
}

func Test1EAE(t *testing.T) {
	ValidateRune(t, 0x1EAE, pDISALLOWED)
}

func Test1EAF(t *testing.T) {
	ValidateRune(t, 0x1EAF, pPVALID)
}

func Test1EB0(t *testing.T) {
	ValidateRune(t, 0x1EB0, pDISALLOWED)
}

func Test1EB1(t *testing.T) {
	ValidateRune(t, 0x1EB1, pPVALID)
}

func Test1EB2(t *testing.T) {
	ValidateRune(t, 0x1EB2, pDISALLOWED)
}

func Test1EB3(t *testing.T) {
	ValidateRune(t, 0x1EB3, pPVALID)
}

func Test1EB4(t *testing.T) {
	ValidateRune(t, 0x1EB4, pDISALLOWED)
}

func Test1EB5(t *testing.T) {
	ValidateRune(t, 0x1EB5, pPVALID)
}

func Test1EB6(t *testing.T) {
	ValidateRune(t, 0x1EB6, pDISALLOWED)
}

func Test1EB7(t *testing.T) {
	ValidateRune(t, 0x1EB7, pPVALID)
}

func Test1EB8(t *testing.T) {
	ValidateRune(t, 0x1EB8, pDISALLOWED)
}

func Test1EB9(t *testing.T) {
	ValidateRune(t, 0x1EB9, pPVALID)
}

func Test1EBA(t *testing.T) {
	ValidateRune(t, 0x1EBA, pDISALLOWED)
}

func Test1EBB(t *testing.T) {
	ValidateRune(t, 0x1EBB, pPVALID)
}

func Test1EBC(t *testing.T) {
	ValidateRune(t, 0x1EBC, pDISALLOWED)
}

func Test1EBD(t *testing.T) {
	ValidateRune(t, 0x1EBD, pPVALID)
}

func Test1EBE(t *testing.T) {
	ValidateRune(t, 0x1EBE, pDISALLOWED)
}

func Test1EBF(t *testing.T) {
	ValidateRune(t, 0x1EBF, pPVALID)
}

func Test1EC0(t *testing.T) {
	ValidateRune(t, 0x1EC0, pDISALLOWED)
}

func Test1EC1(t *testing.T) {
	ValidateRune(t, 0x1EC1, pPVALID)
}

func Test1EC2(t *testing.T) {
	ValidateRune(t, 0x1EC2, pDISALLOWED)
}

func Test1EC3(t *testing.T) {
	ValidateRune(t, 0x1EC3, pPVALID)
}

func Test1EC4(t *testing.T) {
	ValidateRune(t, 0x1EC4, pDISALLOWED)
}

func Test1EC5(t *testing.T) {
	ValidateRune(t, 0x1EC5, pPVALID)
}

func Test1EC6(t *testing.T) {
	ValidateRune(t, 0x1EC6, pDISALLOWED)
}

func Test1EC7(t *testing.T) {
	ValidateRune(t, 0x1EC7, pPVALID)
}

func Test1EC8(t *testing.T) {
	ValidateRune(t, 0x1EC8, pDISALLOWED)
}

func Test1EC9(t *testing.T) {
	ValidateRune(t, 0x1EC9, pPVALID)
}

func Test1ECA(t *testing.T) {
	ValidateRune(t, 0x1ECA, pDISALLOWED)
}

func Test1ECB(t *testing.T) {
	ValidateRune(t, 0x1ECB, pPVALID)
}

func Test1ECC(t *testing.T) {
	ValidateRune(t, 0x1ECC, pDISALLOWED)
}

func Test1ECD(t *testing.T) {
	ValidateRune(t, 0x1ECD, pPVALID)
}

func Test1ECE(t *testing.T) {
	ValidateRune(t, 0x1ECE, pDISALLOWED)
}

func Test1ECF(t *testing.T) {
	ValidateRune(t, 0x1ECF, pPVALID)
}

func Test1ED0(t *testing.T) {
	ValidateRune(t, 0x1ED0, pDISALLOWED)
}

func Test1ED1(t *testing.T) {
	ValidateRune(t, 0x1ED1, pPVALID)
}

func Test1ED2(t *testing.T) {
	ValidateRune(t, 0x1ED2, pDISALLOWED)
}

func Test1ED3(t *testing.T) {
	ValidateRune(t, 0x1ED3, pPVALID)
}

func Test1ED4(t *testing.T) {
	ValidateRune(t, 0x1ED4, pDISALLOWED)
}

func Test1ED5(t *testing.T) {
	ValidateRune(t, 0x1ED5, pPVALID)
}

func Test1ED6(t *testing.T) {
	ValidateRune(t, 0x1ED6, pDISALLOWED)
}

func Test1ED7(t *testing.T) {
	ValidateRune(t, 0x1ED7, pPVALID)
}

func Test1ED8(t *testing.T) {
	ValidateRune(t, 0x1ED8, pDISALLOWED)
}

func Test1ED9(t *testing.T) {
	ValidateRune(t, 0x1ED9, pPVALID)
}

func Test1EDA(t *testing.T) {
	ValidateRune(t, 0x1EDA, pDISALLOWED)
}

func Test1EDB(t *testing.T) {
	ValidateRune(t, 0x1EDB, pPVALID)
}

func Test1EDC(t *testing.T) {
	ValidateRune(t, 0x1EDC, pDISALLOWED)
}

func Test1EDD(t *testing.T) {
	ValidateRune(t, 0x1EDD, pPVALID)
}

func Test1EDE(t *testing.T) {
	ValidateRune(t, 0x1EDE, pDISALLOWED)
}

func Test1EDF(t *testing.T) {
	ValidateRune(t, 0x1EDF, pPVALID)
}

func Test1EE0(t *testing.T) {
	ValidateRune(t, 0x1EE0, pDISALLOWED)
}

func Test1EE1(t *testing.T) {
	ValidateRune(t, 0x1EE1, pPVALID)
}

func Test1EE2(t *testing.T) {
	ValidateRune(t, 0x1EE2, pDISALLOWED)
}

func Test1EE3(t *testing.T) {
	ValidateRune(t, 0x1EE3, pPVALID)
}

func Test1EE4(t *testing.T) {
	ValidateRune(t, 0x1EE4, pDISALLOWED)
}

func Test1EE5(t *testing.T) {
	ValidateRune(t, 0x1EE5, pPVALID)
}

func Test1EE6(t *testing.T) {
	ValidateRune(t, 0x1EE6, pDISALLOWED)
}

func Test1EE7(t *testing.T) {
	ValidateRune(t, 0x1EE7, pPVALID)
}

func Test1EE8(t *testing.T) {
	ValidateRune(t, 0x1EE8, pDISALLOWED)
}

func Test1EE9(t *testing.T) {
	ValidateRune(t, 0x1EE9, pPVALID)
}

func Test1EEA(t *testing.T) {
	ValidateRune(t, 0x1EEA, pDISALLOWED)
}

func Test1EEB(t *testing.T) {
	ValidateRune(t, 0x1EEB, pPVALID)
}

func Test1EEC(t *testing.T) {
	ValidateRune(t, 0x1EEC, pDISALLOWED)
}

func Test1EED(t *testing.T) {
	ValidateRune(t, 0x1EED, pPVALID)
}

func Test1EEE(t *testing.T) {
	ValidateRune(t, 0x1EEE, pDISALLOWED)
}

func Test1EEF(t *testing.T) {
	ValidateRune(t, 0x1EEF, pPVALID)
}

func Test1EF0(t *testing.T) {
	ValidateRune(t, 0x1EF0, pDISALLOWED)
}

func Test1EF1(t *testing.T) {
	ValidateRune(t, 0x1EF1, pPVALID)
}

func Test1EF2(t *testing.T) {
	ValidateRune(t, 0x1EF2, pDISALLOWED)
}

func Test1EF3(t *testing.T) {
	ValidateRune(t, 0x1EF3, pPVALID)
}

func Test1EF4(t *testing.T) {
	ValidateRune(t, 0x1EF4, pDISALLOWED)
}

func Test1EF5(t *testing.T) {
	ValidateRune(t, 0x1EF5, pPVALID)
}

func Test1EF6(t *testing.T) {
	ValidateRune(t, 0x1EF6, pDISALLOWED)
}

func Test1EF7(t *testing.T) {
	ValidateRune(t, 0x1EF7, pPVALID)
}

func Test1EF8(t *testing.T) {
	ValidateRune(t, 0x1EF8, pDISALLOWED)
}

func Test1EF9(t *testing.T) {
	ValidateRune(t, 0x1EF9, pPVALID)
}

func Test1EFA(t *testing.T) {
	ValidateRune(t, 0x1EFA, pDISALLOWED)
}

func Test1EFB(t *testing.T) {
	ValidateRune(t, 0x1EFB, pPVALID)
}

func Test1EFC(t *testing.T) {
	ValidateRune(t, 0x1EFC, pDISALLOWED)
}

func Test1EFD(t *testing.T) {
	ValidateRune(t, 0x1EFD, pPVALID)
}

func Test1EFE(t *testing.T) {
	ValidateRune(t, 0x1EFE, pDISALLOWED)
}

func Test1EFF_1F07(t *testing.T) {
	ValidateRuneRange(t, 0x1EFF, 0x1F07, pPVALID)
}

func Test1F08_1F0F(t *testing.T) {
	ValidateRuneRange(t, 0x1F08, 0x1F0F, pDISALLOWED)
}

func Test1F10_1F15(t *testing.T) {
	ValidateRuneRange(t, 0x1F10, 0x1F15, pPVALID)
}

func Test1F16_1F17(t *testing.T) {
	ValidateRuneRange(t, 0x1F16, 0x1F17, pUNASSIGNED)
}

func Test1F18_1F1D(t *testing.T) {
	ValidateRuneRange(t, 0x1F18, 0x1F1D, pDISALLOWED)
}

func Test1F1E_1F1F(t *testing.T) {
	ValidateRuneRange(t, 0x1F1E, 0x1F1F, pUNASSIGNED)
}

func Test1F20_1F27(t *testing.T) {
	ValidateRuneRange(t, 0x1F20, 0x1F27, pPVALID)
}

func Test1F28_1F2F(t *testing.T) {
	ValidateRuneRange(t, 0x1F28, 0x1F2F, pDISALLOWED)
}

func Test1F30_1F37(t *testing.T) {
	ValidateRuneRange(t, 0x1F30, 0x1F37, pPVALID)
}

func Test1F38_1F3F(t *testing.T) {
	ValidateRuneRange(t, 0x1F38, 0x1F3F, pDISALLOWED)
}

func Test1F40_1F45(t *testing.T) {
	ValidateRuneRange(t, 0x1F40, 0x1F45, pPVALID)
}

func Test1F46_1F47(t *testing.T) {
	ValidateRuneRange(t, 0x1F46, 0x1F47, pUNASSIGNED)
}

func Test1F48_1F4D(t *testing.T) {
	ValidateRuneRange(t, 0x1F48, 0x1F4D, pDISALLOWED)
}

func Test1F4E_1F4F(t *testing.T) {
	ValidateRuneRange(t, 0x1F4E, 0x1F4F, pUNASSIGNED)
}

func Test1F50_1F57(t *testing.T) {
	ValidateRuneRange(t, 0x1F50, 0x1F57, pPVALID)
}

func Test1F58(t *testing.T) {
	ValidateRune(t, 0x1F58, pUNASSIGNED)
}

func Test1F59(t *testing.T) {
	ValidateRune(t, 0x1F59, pDISALLOWED)
}

func Test1F5A(t *testing.T) {
	ValidateRune(t, 0x1F5A, pUNASSIGNED)
}

func Test1F5B(t *testing.T) {
	ValidateRune(t, 0x1F5B, pDISALLOWED)
}

func Test1F5C(t *testing.T) {
	ValidateRune(t, 0x1F5C, pUNASSIGNED)
}

func Test1F5D(t *testing.T) {
	ValidateRune(t, 0x1F5D, pDISALLOWED)
}

func Test1F5E(t *testing.T) {
	ValidateRune(t, 0x1F5E, pUNASSIGNED)
}

func Test1F5F(t *testing.T) {
	ValidateRune(t, 0x1F5F, pDISALLOWED)
}

func Test1F60_1F67(t *testing.T) {
	ValidateRuneRange(t, 0x1F60, 0x1F67, pPVALID)
}

func Test1F68_1F6F(t *testing.T) {
	ValidateRuneRange(t, 0x1F68, 0x1F6F, pDISALLOWED)
}

func Test1F70(t *testing.T) {
	ValidateRune(t, 0x1F70, pPVALID)
}

func Test1F71(t *testing.T) {
	ValidateRune(t, 0x1F71, pDISALLOWED)
}

func Test1F72(t *testing.T) {
	ValidateRune(t, 0x1F72, pPVALID)
}

func Test1F73(t *testing.T) {
	ValidateRune(t, 0x1F73, pDISALLOWED)
}

func Test1F74(t *testing.T) {
	ValidateRune(t, 0x1F74, pPVALID)
}

func Test1F75(t *testing.T) {
	ValidateRune(t, 0x1F75, pDISALLOWED)
}

func Test1F76(t *testing.T) {
	ValidateRune(t, 0x1F76, pPVALID)
}

func Test1F77(t *testing.T) {
	ValidateRune(t, 0x1F77, pDISALLOWED)
}

func Test1F78(t *testing.T) {
	ValidateRune(t, 0x1F78, pPVALID)
}

func Test1F79(t *testing.T) {
	ValidateRune(t, 0x1F79, pDISALLOWED)
}

func Test1F7A(t *testing.T) {
	ValidateRune(t, 0x1F7A, pPVALID)
}

func Test1F7B(t *testing.T) {
	ValidateRune(t, 0x1F7B, pDISALLOWED)
}

func Test1F7C(t *testing.T) {
	ValidateRune(t, 0x1F7C, pPVALID)
}

func Test1F7D(t *testing.T) {
	ValidateRune(t, 0x1F7D, pDISALLOWED)
}

func Test1F7E_1F7F(t *testing.T) {
	ValidateRuneRange(t, 0x1F7E, 0x1F7F, pUNASSIGNED)
}

func Test1F80_1FAF(t *testing.T) {
	ValidateRuneRange(t, 0x1F80, 0x1FAF, pDISALLOWED)
}

func Test1FB0_1FB1(t *testing.T) {
	ValidateRuneRange(t, 0x1FB0, 0x1FB1, pPVALID)
}

func Test1FB2_1FB4(t *testing.T) {
	ValidateRuneRange(t, 0x1FB2, 0x1FB4, pDISALLOWED)
}

func Test1FB5(t *testing.T) {
	ValidateRune(t, 0x1FB5, pUNASSIGNED)
}

func Test1FB6(t *testing.T) {
	ValidateRune(t, 0x1FB6, pPVALID)
}

func Test1FB7_1FC4(t *testing.T) {
	ValidateRuneRange(t, 0x1FB7, 0x1FC4, pDISALLOWED)
}

func Test1FC5(t *testing.T) {
	ValidateRune(t, 0x1FC5, pUNASSIGNED)
}

func Test1FC6(t *testing.T) {
	ValidateRune(t, 0x1FC6, pPVALID)
}

func Test1FC7_1FCF(t *testing.T) {
	ValidateRuneRange(t, 0x1FC7, 0x1FCF, pDISALLOWED)
}

func Test1FD0_1FD2(t *testing.T) {
	ValidateRuneRange(t, 0x1FD0, 0x1FD2, pPVALID)
}

func Test1FD3(t *testing.T) {
	ValidateRune(t, 0x1FD3, pDISALLOWED)
}

func Test1FD4_1FD5(t *testing.T) {
	ValidateRuneRange(t, 0x1FD4, 0x1FD5, pUNASSIGNED)
}

func Test1FD6_1FD7(t *testing.T) {
	ValidateRuneRange(t, 0x1FD6, 0x1FD7, pPVALID)
}

func Test1FD8_1FDB(t *testing.T) {
	ValidateRuneRange(t, 0x1FD8, 0x1FDB, pDISALLOWED)
}

func Test1FDC(t *testing.T) {
	ValidateRune(t, 0x1FDC, pUNASSIGNED)
}

func Test1FDD_1FDF(t *testing.T) {
	ValidateRuneRange(t, 0x1FDD, 0x1FDF, pDISALLOWED)
}

func Test1FE0_1FE2(t *testing.T) {
	ValidateRuneRange(t, 0x1FE0, 0x1FE2, pPVALID)
}

func Test1FE3(t *testing.T) {
	ValidateRune(t, 0x1FE3, pDISALLOWED)
}

func Test1FE4_1FE7(t *testing.T) {
	ValidateRuneRange(t, 0x1FE4, 0x1FE7, pPVALID)
}

func Test1FE8_1FEF(t *testing.T) {
	ValidateRuneRange(t, 0x1FE8, 0x1FEF, pDISALLOWED)
}

func Test1FF0_1FF1(t *testing.T) {
	ValidateRuneRange(t, 0x1FF0, 0x1FF1, pUNASSIGNED)
}

func Test1FF2_1FF4(t *testing.T) {
	ValidateRuneRange(t, 0x1FF2, 0x1FF4, pDISALLOWED)
}

func Test1FF5(t *testing.T) {
	ValidateRune(t, 0x1FF5, pUNASSIGNED)
}

func Test1FF6(t *testing.T) {
	ValidateRune(t, 0x1FF6, pPVALID)
}

func Test1FF7_1FFE(t *testing.T) {
	ValidateRuneRange(t, 0x1FF7, 0x1FFE, pDISALLOWED)
}

func Test1FFF(t *testing.T) {
	ValidateRune(t, 0x1FFF, pUNASSIGNED)
}

func Test2000_200B(t *testing.T) {
	ValidateRuneRange(t, 0x2000, 0x200B, pDISALLOWED)
}

func Test200C_200D(t *testing.T) {
	ValidateRuneRange(t, 0x200C, 0x200D, pCONTEXTJ)
}

func Test200E_2064(t *testing.T) {
	ValidateRuneRange(t, 0x200E, 0x2064, pDISALLOWED)
}

func Test2065_2069(t *testing.T) {
	ValidateRuneRange(t, 0x2065, 0x2069, pUNASSIGNED)
}

func Test206A_2071(t *testing.T) {
	ValidateRuneRange(t, 0x206A, 0x2071, pDISALLOWED)
}

func Test2072_2073(t *testing.T) {
	ValidateRuneRange(t, 0x2072, 0x2073, pUNASSIGNED)
}

func Test2074_208E(t *testing.T) {
	ValidateRuneRange(t, 0x2074, 0x208E, pDISALLOWED)
}

func Test208F(t *testing.T) {
	ValidateRune(t, 0x208F, pUNASSIGNED)
}

func Test2090_2094(t *testing.T) {
	ValidateRuneRange(t, 0x2090, 0x2094, pDISALLOWED)
}

func Test2095_209F(t *testing.T) {
	ValidateRuneRange(t, 0x2095, 0x209F, pUNASSIGNED)
}

func Test20A0_20B8(t *testing.T) {
	ValidateRuneRange(t, 0x20A0, 0x20B8, pDISALLOWED)
}

func Test20B9_20CF(t *testing.T) {
	ValidateRuneRange(t, 0x20B9, 0x20CF, pUNASSIGNED)
}

func Test20D0_20F0(t *testing.T) {
	ValidateRuneRange(t, 0x20D0, 0x20F0, pDISALLOWED)
}

func Test20F1_20FF(t *testing.T) {
	ValidateRuneRange(t, 0x20F1, 0x20FF, pUNASSIGNED)
}

func Test2100_214D(t *testing.T) {
	ValidateRuneRange(t, 0x2100, 0x214D, pDISALLOWED)
}

func Test214E(t *testing.T) {
	ValidateRune(t, 0x214E, pPVALID)
}

func Test214F_2183(t *testing.T) {
	ValidateRuneRange(t, 0x214F, 0x2183, pDISALLOWED)
}

func Test2185_2189(t *testing.T) {
	ValidateRuneRange(t, 0x2185, 0x2189, pDISALLOWED)
}

func Test218A_218F(t *testing.T) {
	ValidateRuneRange(t, 0x218A, 0x218F, pUNASSIGNED)
}

func Test2190_23E8(t *testing.T) {
	ValidateRuneRange(t, 0x2190, 0x23E8, pDISALLOWED)
}

func Test23E9_23FF(t *testing.T) {
	ValidateRuneRange(t, 0x23E9, 0x23FF, pUNASSIGNED)
}

func Test2400_2426(t *testing.T) {
	ValidateRuneRange(t, 0x2400, 0x2426, pDISALLOWED)
}

func Test2427_243F(t *testing.T) {
	ValidateRuneRange(t, 0x2427, 0x243F, pUNASSIGNED)
}

func Test2440_244A(t *testing.T) {
	ValidateRuneRange(t, 0x2440, 0x244A, pDISALLOWED)
}

func Test244B_245F(t *testing.T) {
	ValidateRuneRange(t, 0x244B, 0x245F, pUNASSIGNED)
}

func Test2460_26CD(t *testing.T) {
	ValidateRuneRange(t, 0x2460, 0x26CD, pDISALLOWED)
}

func Test26CE(t *testing.T) {
	ValidateRune(t, 0x26CE, pUNASSIGNED)
}

func Test26CF_26E1(t *testing.T) {
	ValidateRuneRange(t, 0x26CF, 0x26E1, pDISALLOWED)
}

func Test26E2(t *testing.T) {
	ValidateRune(t, 0x26E2, pUNASSIGNED)
}

func Test26E3(t *testing.T) {
	ValidateRune(t, 0x26E3, pDISALLOWED)
}

func Test26E4_26E7(t *testing.T) {
	ValidateRuneRange(t, 0x26E4, 0x26E7, pUNASSIGNED)
}

func Test26E8_26FF(t *testing.T) {
	ValidateRuneRange(t, 0x26E8, 0x26FF, pDISALLOWED)
}

func Test2701_2704(t *testing.T) {
	ValidateRuneRange(t, 0x2701, 0x2704, pDISALLOWED)
}

func Test2706_2709(t *testing.T) {
	ValidateRuneRange(t, 0x2706, 0x2709, pDISALLOWED)
}

func Test270A_270B(t *testing.T) {
	ValidateRuneRange(t, 0x270A, 0x270B, pUNASSIGNED)
}

func Test270C_2727(t *testing.T) {
	ValidateRuneRange(t, 0x270C, 0x2727, pDISALLOWED)
}

func Test2729_274B(t *testing.T) {
	ValidateRuneRange(t, 0x2729, 0x274B, pDISALLOWED)
}

func Test274C(t *testing.T) {
	ValidateRune(t, 0x274C, pUNASSIGNED)
}

func Test274D(t *testing.T) {
	ValidateRune(t, 0x274D, pDISALLOWED)
}

func Test274E(t *testing.T) {
	ValidateRune(t, 0x274E, pUNASSIGNED)
}

func Test274F_2752(t *testing.T) {
	ValidateRuneRange(t, 0x274F, 0x2752, pDISALLOWED)
}

func Test2753_2755(t *testing.T) {
	ValidateRuneRange(t, 0x2753, 0x2755, pUNASSIGNED)
}

func Test2756_275E(t *testing.T) {
	ValidateRuneRange(t, 0x2756, 0x275E, pDISALLOWED)
}

func Test275F_2760(t *testing.T) {
	ValidateRuneRange(t, 0x275F, 0x2760, pUNASSIGNED)
}

func Test2761_2794(t *testing.T) {
	ValidateRuneRange(t, 0x2761, 0x2794, pDISALLOWED)
}

func Test2795_2797(t *testing.T) {
	ValidateRuneRange(t, 0x2795, 0x2797, pUNASSIGNED)
}

func Test2798_27AF(t *testing.T) {
	ValidateRuneRange(t, 0x2798, 0x27AF, pDISALLOWED)
}

func Test27B0(t *testing.T) {
	ValidateRune(t, 0x27B0, pUNASSIGNED)
}

func Test27B1_27BE(t *testing.T) {
	ValidateRuneRange(t, 0x27B1, 0x27BE, pDISALLOWED)
}

func Test27BF(t *testing.T) {
	ValidateRune(t, 0x27BF, pUNASSIGNED)
}

func Test27C0_27CA(t *testing.T) {
	ValidateRuneRange(t, 0x27C0, 0x27CA, pDISALLOWED)
}

func Test27CB(t *testing.T) {
	ValidateRune(t, 0x27CB, pUNASSIGNED)
}

func Test27CC(t *testing.T) {
	ValidateRune(t, 0x27CC, pDISALLOWED)
}

func Test27CD_27CF(t *testing.T) {
	ValidateRuneRange(t, 0x27CD, 0x27CF, pUNASSIGNED)
}

func Test27D0_2B4C(t *testing.T) {
	ValidateRuneRange(t, 0x27D0, 0x2B4C, pDISALLOWED)
}

func Test2B4D_2B4F(t *testing.T) {
	ValidateRuneRange(t, 0x2B4D, 0x2B4F, pUNASSIGNED)
}

func Test2B50_2B59(t *testing.T) {
	ValidateRuneRange(t, 0x2B50, 0x2B59, pDISALLOWED)
}

func Test2B5A_2BFF(t *testing.T) {
	ValidateRuneRange(t, 0x2B5A, 0x2BFF, pUNASSIGNED)
}

func Test2C00_2C2E(t *testing.T) {
	ValidateRuneRange(t, 0x2C00, 0x2C2E, pDISALLOWED)
}

func Test2C2F(t *testing.T) {
	ValidateRune(t, 0x2C2F, pUNASSIGNED)
}

func Test2C30_2C5E(t *testing.T) {
	ValidateRuneRange(t, 0x2C30, 0x2C5E, pPVALID)
}

func Test2C5F(t *testing.T) {
	ValidateRune(t, 0x2C5F, pUNASSIGNED)
}

func Test2C60(t *testing.T) {
	ValidateRune(t, 0x2C60, pDISALLOWED)
}

func Test2C61(t *testing.T) {
	ValidateRune(t, 0x2C61, pPVALID)
}

func Test2C62_2C64(t *testing.T) {
	ValidateRuneRange(t, 0x2C62, 0x2C64, pDISALLOWED)
}

func Test2C65_2C66(t *testing.T) {
	ValidateRuneRange(t, 0x2C65, 0x2C66, pPVALID)
}

func Test2C67(t *testing.T) {
	ValidateRune(t, 0x2C67, pDISALLOWED)
}

func Test2C68(t *testing.T) {
	ValidateRune(t, 0x2C68, pPVALID)
}

func Test2C69(t *testing.T) {
	ValidateRune(t, 0x2C69, pDISALLOWED)
}

func Test2C6A(t *testing.T) {
	ValidateRune(t, 0x2C6A, pPVALID)
}

func Test2C6B(t *testing.T) {
	ValidateRune(t, 0x2C6B, pDISALLOWED)
}

func Test2C6C(t *testing.T) {
	ValidateRune(t, 0x2C6C, pPVALID)
}

func Test2C6D_2C70(t *testing.T) {
	ValidateRuneRange(t, 0x2C6D, 0x2C70, pDISALLOWED)
}

func Test2C71(t *testing.T) {
	ValidateRune(t, 0x2C71, pPVALID)
}

func Test2C72(t *testing.T) {
	ValidateRune(t, 0x2C72, pDISALLOWED)
}

func Test2C73_2C74(t *testing.T) {
	ValidateRuneRange(t, 0x2C73, 0x2C74, pPVALID)
}

func Test2C75(t *testing.T) {
	ValidateRune(t, 0x2C75, pDISALLOWED)
}

func Test2C76_2C7B(t *testing.T) {
	ValidateRuneRange(t, 0x2C76, 0x2C7B, pPVALID)
}

func Test2C7C_2C80(t *testing.T) {
	ValidateRuneRange(t, 0x2C7C, 0x2C80, pDISALLOWED)
}

func Test2C81(t *testing.T) {
	ValidateRune(t, 0x2C81, pPVALID)
}

func Test2C82(t *testing.T) {
	ValidateRune(t, 0x2C82, pDISALLOWED)
}

func Test2C83(t *testing.T) {
	ValidateRune(t, 0x2C83, pPVALID)
}

func Test2C84(t *testing.T) {
	ValidateRune(t, 0x2C84, pDISALLOWED)
}

func Test2C85(t *testing.T) {
	ValidateRune(t, 0x2C85, pPVALID)
}

func Test2C86(t *testing.T) {
	ValidateRune(t, 0x2C86, pDISALLOWED)
}

func Test2C87(t *testing.T) {
	ValidateRune(t, 0x2C87, pPVALID)
}

func Test2C88(t *testing.T) {
	ValidateRune(t, 0x2C88, pDISALLOWED)
}

func Test2C89(t *testing.T) {
	ValidateRune(t, 0x2C89, pPVALID)
}

func Test2C8A(t *testing.T) {
	ValidateRune(t, 0x2C8A, pDISALLOWED)
}

func Test2C8B(t *testing.T) {
	ValidateRune(t, 0x2C8B, pPVALID)
}

func Test2C8C(t *testing.T) {
	ValidateRune(t, 0x2C8C, pDISALLOWED)
}

func Test2C8D(t *testing.T) {
	ValidateRune(t, 0x2C8D, pPVALID)
}

func Test2C8E(t *testing.T) {
	ValidateRune(t, 0x2C8E, pDISALLOWED)
}

func Test2C8F(t *testing.T) {
	ValidateRune(t, 0x2C8F, pPVALID)
}

func Test2C90(t *testing.T) {
	ValidateRune(t, 0x2C90, pDISALLOWED)
}

func Test2C91(t *testing.T) {
	ValidateRune(t, 0x2C91, pPVALID)
}

func Test2C92(t *testing.T) {
	ValidateRune(t, 0x2C92, pDISALLOWED)
}

func Test2C93(t *testing.T) {
	ValidateRune(t, 0x2C93, pPVALID)
}

func Test2C94(t *testing.T) {
	ValidateRune(t, 0x2C94, pDISALLOWED)
}

func Test2C95(t *testing.T) {
	ValidateRune(t, 0x2C95, pPVALID)
}

func Test2C96(t *testing.T) {
	ValidateRune(t, 0x2C96, pDISALLOWED)
}

func Test2C97(t *testing.T) {
	ValidateRune(t, 0x2C97, pPVALID)
}

func Test2C98(t *testing.T) {
	ValidateRune(t, 0x2C98, pDISALLOWED)
}

func Test2C99(t *testing.T) {
	ValidateRune(t, 0x2C99, pPVALID)
}

func Test2C9A(t *testing.T) {
	ValidateRune(t, 0x2C9A, pDISALLOWED)
}

func Test2C9B(t *testing.T) {
	ValidateRune(t, 0x2C9B, pPVALID)
}

func Test2C9C(t *testing.T) {
	ValidateRune(t, 0x2C9C, pDISALLOWED)
}

func Test2C9D(t *testing.T) {
	ValidateRune(t, 0x2C9D, pPVALID)
}

func Test2C9E(t *testing.T) {
	ValidateRune(t, 0x2C9E, pDISALLOWED)
}

func Test2C9F(t *testing.T) {
	ValidateRune(t, 0x2C9F, pPVALID)
}

func Test2CA0(t *testing.T) {
	ValidateRune(t, 0x2CA0, pDISALLOWED)
}

func Test2CA1(t *testing.T) {
	ValidateRune(t, 0x2CA1, pPVALID)
}

func Test2CA2(t *testing.T) {
	ValidateRune(t, 0x2CA2, pDISALLOWED)
}

func Test2CA3(t *testing.T) {
	ValidateRune(t, 0x2CA3, pPVALID)
}

func Test2CA4(t *testing.T) {
	ValidateRune(t, 0x2CA4, pDISALLOWED)
}

func Test2CA5(t *testing.T) {
	ValidateRune(t, 0x2CA5, pPVALID)
}

func Test2CA6(t *testing.T) {
	ValidateRune(t, 0x2CA6, pDISALLOWED)
}

func Test2CA7(t *testing.T) {
	ValidateRune(t, 0x2CA7, pPVALID)
}

func Test2CA8(t *testing.T) {
	ValidateRune(t, 0x2CA8, pDISALLOWED)
}

func Test2CA9(t *testing.T) {
	ValidateRune(t, 0x2CA9, pPVALID)
}

func Test2CAA(t *testing.T) {
	ValidateRune(t, 0x2CAA, pDISALLOWED)
}

func Test2CAB(t *testing.T) {
	ValidateRune(t, 0x2CAB, pPVALID)
}

func Test2CAC(t *testing.T) {
	ValidateRune(t, 0x2CAC, pDISALLOWED)
}

func Test2CAD(t *testing.T) {
	ValidateRune(t, 0x2CAD, pPVALID)
}

func Test2CAE(t *testing.T) {
	ValidateRune(t, 0x2CAE, pDISALLOWED)
}

func Test2CAF(t *testing.T) {
	ValidateRune(t, 0x2CAF, pPVALID)
}

func Test2CB0(t *testing.T) {
	ValidateRune(t, 0x2CB0, pDISALLOWED)
}

func Test2CB1(t *testing.T) {
	ValidateRune(t, 0x2CB1, pPVALID)
}

func Test2CB2(t *testing.T) {
	ValidateRune(t, 0x2CB2, pDISALLOWED)
}

func Test2CB3(t *testing.T) {
	ValidateRune(t, 0x2CB3, pPVALID)
}

func Test2CB4(t *testing.T) {
	ValidateRune(t, 0x2CB4, pDISALLOWED)
}

func Test2CB5(t *testing.T) {
	ValidateRune(t, 0x2CB5, pPVALID)
}

func Test2CB6(t *testing.T) {
	ValidateRune(t, 0x2CB6, pDISALLOWED)
}

func Test2CB7(t *testing.T) {
	ValidateRune(t, 0x2CB7, pPVALID)
}

func Test2CB8(t *testing.T) {
	ValidateRune(t, 0x2CB8, pDISALLOWED)
}

func Test2CB9(t *testing.T) {
	ValidateRune(t, 0x2CB9, pPVALID)
}

func Test2CBA(t *testing.T) {
	ValidateRune(t, 0x2CBA, pDISALLOWED)
}

func Test2CBB(t *testing.T) {
	ValidateRune(t, 0x2CBB, pPVALID)
}

func Test2CBC(t *testing.T) {
	ValidateRune(t, 0x2CBC, pDISALLOWED)
}

func Test2CBD(t *testing.T) {
	ValidateRune(t, 0x2CBD, pPVALID)
}

func Test2CBE(t *testing.T) {
	ValidateRune(t, 0x2CBE, pDISALLOWED)
}

func Test2CBF(t *testing.T) {
	ValidateRune(t, 0x2CBF, pPVALID)
}

func Test2CC0(t *testing.T) {
	ValidateRune(t, 0x2CC0, pDISALLOWED)
}

func Test2CC1(t *testing.T) {
	ValidateRune(t, 0x2CC1, pPVALID)
}

func Test2CC2(t *testing.T) {
	ValidateRune(t, 0x2CC2, pDISALLOWED)
}

func Test2CC3(t *testing.T) {
	ValidateRune(t, 0x2CC3, pPVALID)
}

func Test2CC4(t *testing.T) {
	ValidateRune(t, 0x2CC4, pDISALLOWED)
}

func Test2CC5(t *testing.T) {
	ValidateRune(t, 0x2CC5, pPVALID)
}

func Test2CC6(t *testing.T) {
	ValidateRune(t, 0x2CC6, pDISALLOWED)
}

func Test2CC7(t *testing.T) {
	ValidateRune(t, 0x2CC7, pPVALID)
}

func Test2CC8(t *testing.T) {
	ValidateRune(t, 0x2CC8, pDISALLOWED)
}

func Test2CC9(t *testing.T) {
	ValidateRune(t, 0x2CC9, pPVALID)
}

func Test2CCA(t *testing.T) {
	ValidateRune(t, 0x2CCA, pDISALLOWED)
}

func Test2CCB(t *testing.T) {
	ValidateRune(t, 0x2CCB, pPVALID)
}

func Test2CCC(t *testing.T) {
	ValidateRune(t, 0x2CCC, pDISALLOWED)
}

func Test2CCD(t *testing.T) {
	ValidateRune(t, 0x2CCD, pPVALID)
}

func Test2CCE(t *testing.T) {
	ValidateRune(t, 0x2CCE, pDISALLOWED)
}

func Test2CCF(t *testing.T) {
	ValidateRune(t, 0x2CCF, pPVALID)
}

func Test2CD0(t *testing.T) {
	ValidateRune(t, 0x2CD0, pDISALLOWED)
}

func Test2CD1(t *testing.T) {
	ValidateRune(t, 0x2CD1, pPVALID)
}

func Test2CD2(t *testing.T) {
	ValidateRune(t, 0x2CD2, pDISALLOWED)
}

func Test2CD3(t *testing.T) {
	ValidateRune(t, 0x2CD3, pPVALID)
}

func Test2CD4(t *testing.T) {
	ValidateRune(t, 0x2CD4, pDISALLOWED)
}

func Test2CD5(t *testing.T) {
	ValidateRune(t, 0x2CD5, pPVALID)
}

func Test2CD6(t *testing.T) {
	ValidateRune(t, 0x2CD6, pDISALLOWED)
}

func Test2CD7(t *testing.T) {
	ValidateRune(t, 0x2CD7, pPVALID)
}

func Test2CD8(t *testing.T) {
	ValidateRune(t, 0x2CD8, pDISALLOWED)
}

func Test2CD9(t *testing.T) {
	ValidateRune(t, 0x2CD9, pPVALID)
}

func Test2CDA(t *testing.T) {
	ValidateRune(t, 0x2CDA, pDISALLOWED)
}

func Test2CDB(t *testing.T) {
	ValidateRune(t, 0x2CDB, pPVALID)
}

func Test2CDC(t *testing.T) {
	ValidateRune(t, 0x2CDC, pDISALLOWED)
}

func Test2CDD(t *testing.T) {
	ValidateRune(t, 0x2CDD, pPVALID)
}

func Test2CDE(t *testing.T) {
	ValidateRune(t, 0x2CDE, pDISALLOWED)
}

func Test2CDF(t *testing.T) {
	ValidateRune(t, 0x2CDF, pPVALID)
}

func Test2CE0(t *testing.T) {
	ValidateRune(t, 0x2CE0, pDISALLOWED)
}

func Test2CE1(t *testing.T) {
	ValidateRune(t, 0x2CE1, pPVALID)
}

func Test2CE2(t *testing.T) {
	ValidateRune(t, 0x2CE2, pDISALLOWED)
}

func Test2CE3_2CE4(t *testing.T) {
	ValidateRuneRange(t, 0x2CE3, 0x2CE4, pPVALID)
}

func Test2CE5_2CEB(t *testing.T) {
	ValidateRuneRange(t, 0x2CE5, 0x2CEB, pDISALLOWED)
}

func Test2CEC(t *testing.T) {
	ValidateRune(t, 0x2CEC, pPVALID)
}

func Test2CED(t *testing.T) {
	ValidateRune(t, 0x2CED, pDISALLOWED)
}

func Test2CEE_2CF1(t *testing.T) {
	ValidateRuneRange(t, 0x2CEE, 0x2CF1, pPVALID)
}

func Test2CF2_2CF8(t *testing.T) {
	ValidateRuneRange(t, 0x2CF2, 0x2CF8, pUNASSIGNED)
}

func Test2CF9_2CFF(t *testing.T) {
	ValidateRuneRange(t, 0x2CF9, 0x2CFF, pDISALLOWED)
}

func Test2D00_2D25(t *testing.T) {
	ValidateRuneRange(t, 0x2D00, 0x2D25, pPVALID)
}

func Test2D26_2D2F(t *testing.T) {
	ValidateRuneRange(t, 0x2D26, 0x2D2F, pUNASSIGNED)
}

func Test2D30_2D65(t *testing.T) {
	ValidateRuneRange(t, 0x2D30, 0x2D65, pPVALID)
}

func Test2D66_2D6E(t *testing.T) {
	ValidateRuneRange(t, 0x2D66, 0x2D6E, pUNASSIGNED)
}

func Test2D6F(t *testing.T) {
	ValidateRune(t, 0x2D6F, pDISALLOWED)
}

func Test2D70_2D7F(t *testing.T) {
	ValidateRuneRange(t, 0x2D70, 0x2D7F, pUNASSIGNED)
}

func Test2D80_2D96(t *testing.T) {
	ValidateRuneRange(t, 0x2D80, 0x2D96, pPVALID)
}

func Test2D97_2D9F(t *testing.T) {
	ValidateRuneRange(t, 0x2D97, 0x2D9F, pUNASSIGNED)
}

func Test2DA0_2DA6(t *testing.T) {
	ValidateRuneRange(t, 0x2DA0, 0x2DA6, pPVALID)
}

func Test2DA7(t *testing.T) {
	ValidateRune(t, 0x2DA7, pUNASSIGNED)
}

func Test2DA8_2DAE(t *testing.T) {
	ValidateRuneRange(t, 0x2DA8, 0x2DAE, pPVALID)
}

func Test2DAF(t *testing.T) {
	ValidateRune(t, 0x2DAF, pUNASSIGNED)
}

func Test2DB0_2DB6(t *testing.T) {
	ValidateRuneRange(t, 0x2DB0, 0x2DB6, pPVALID)
}

func Test2DB7(t *testing.T) {
	ValidateRune(t, 0x2DB7, pUNASSIGNED)
}

func Test2DB8_2DBE(t *testing.T) {
	ValidateRuneRange(t, 0x2DB8, 0x2DBE, pPVALID)
}

func Test2DBF(t *testing.T) {
	ValidateRune(t, 0x2DBF, pUNASSIGNED)
}

func Test2DC0_2DC6(t *testing.T) {
	ValidateRuneRange(t, 0x2DC0, 0x2DC6, pPVALID)
}

func Test2DC7(t *testing.T) {
	ValidateRune(t, 0x2DC7, pUNASSIGNED)
}

func Test2DC8_2DCE(t *testing.T) {
	ValidateRuneRange(t, 0x2DC8, 0x2DCE, pPVALID)
}

func Test2DCF(t *testing.T) {
	ValidateRune(t, 0x2DCF, pUNASSIGNED)
}

func Test2DD0_2DD6(t *testing.T) {
	ValidateRuneRange(t, 0x2DD0, 0x2DD6, pPVALID)
}

func Test2DD7(t *testing.T) {
	ValidateRune(t, 0x2DD7, pUNASSIGNED)
}

func Test2DD8_2DDE(t *testing.T) {
	ValidateRuneRange(t, 0x2DD8, 0x2DDE, pPVALID)
}

func Test2DDF(t *testing.T) {
	ValidateRune(t, 0x2DDF, pUNASSIGNED)
}

func Test2DE0_2DFF(t *testing.T) {
	ValidateRuneRange(t, 0x2DE0, 0x2DFF, pPVALID)
}

func Test2E00_2E2E(t *testing.T) {
	ValidateRuneRange(t, 0x2E00, 0x2E2E, pDISALLOWED)
}

func Test2E2F(t *testing.T) {
	ValidateRune(t, 0x2E2F, pPVALID)
}

func Test2E30_2E31(t *testing.T) {
	ValidateRuneRange(t, 0x2E30, 0x2E31, pDISALLOWED)
}

func Test2E32_2E7F(t *testing.T) {
	ValidateRuneRange(t, 0x2E32, 0x2E7F, pUNASSIGNED)
}

func Test2E80_2E99(t *testing.T) {
	ValidateRuneRange(t, 0x2E80, 0x2E99, pDISALLOWED)
}

func Test2E9A(t *testing.T) {
	ValidateRune(t, 0x2E9A, pUNASSIGNED)
}

func Test2E9B_2EF3(t *testing.T) {
	ValidateRuneRange(t, 0x2E9B, 0x2EF3, pDISALLOWED)
}

func Test2EF4_2EFF(t *testing.T) {
	ValidateRuneRange(t, 0x2EF4, 0x2EFF, pUNASSIGNED)
}

func Test2F00_2FD5(t *testing.T) {
	ValidateRuneRange(t, 0x2F00, 0x2FD5, pDISALLOWED)
}

func Test2FD6_2FEF(t *testing.T) {
	ValidateRuneRange(t, 0x2FD6, 0x2FEF, pUNASSIGNED)
}

func Test2FF0_2FFB(t *testing.T) {
	ValidateRuneRange(t, 0x2FF0, 0x2FFB, pDISALLOWED)
}

func Test2FFC_2FFF(t *testing.T) {
	ValidateRuneRange(t, 0x2FFC, 0x2FFF, pUNASSIGNED)
}

func Test3000_3004(t *testing.T) {
	ValidateRuneRange(t, 0x3000, 0x3004, pDISALLOWED)
}

func Test3005_3007(t *testing.T) {
	ValidateRuneRange(t, 0x3005, 0x3007, pPVALID)
}

func Test3008_3029(t *testing.T) {
	ValidateRuneRange(t, 0x3008, 0x3029, pDISALLOWED)
}

func Test302A_302D(t *testing.T) {
	ValidateRuneRange(t, 0x302A, 0x302D, pPVALID)
}

func Test302E_303B(t *testing.T) {
	ValidateRuneRange(t, 0x302E, 0x303B, pDISALLOWED)
}

func Test303C(t *testing.T) {
	ValidateRune(t, 0x303C, pPVALID)
}

func Test303D_303F(t *testing.T) {
	ValidateRuneRange(t, 0x303D, 0x303F, pDISALLOWED)
}

func Test3041_3096(t *testing.T) {
	ValidateRuneRange(t, 0x3041, 0x3096, pPVALID)
}

func Test3097_3098(t *testing.T) {
	ValidateRuneRange(t, 0x3097, 0x3098, pUNASSIGNED)
}

func Test3099_309A(t *testing.T) {
	ValidateRuneRange(t, 0x3099, 0x309A, pPVALID)
}

func Test309B_309C(t *testing.T) {
	ValidateRuneRange(t, 0x309B, 0x309C, pDISALLOWED)
}

func Test309D_309E(t *testing.T) {
	ValidateRuneRange(t, 0x309D, 0x309E, pPVALID)
}

func Test309F_30A0(t *testing.T) {
	ValidateRuneRange(t, 0x309F, 0x30A0, pDISALLOWED)
}

func Test30A1_30FA(t *testing.T) {
	ValidateRuneRange(t, 0x30A1, 0x30FA, pPVALID)
}

func Test30FB(t *testing.T) {
	ValidateRune(t, 0x30FB, pCONTEXTO)
}

func Test30FC_30FE(t *testing.T) {
	ValidateRuneRange(t, 0x30FC, 0x30FE, pPVALID)
}

func Test30FF(t *testing.T) {
	ValidateRune(t, 0x30FF, pDISALLOWED)
}

func Test3100_3104(t *testing.T) {
	ValidateRuneRange(t, 0x3100, 0x3104, pUNASSIGNED)
}

func Test3105_312D(t *testing.T) {
	ValidateRuneRange(t, 0x3105, 0x312D, pPVALID)
}

func Test312E_3130(t *testing.T) {
	ValidateRuneRange(t, 0x312E, 0x3130, pUNASSIGNED)
}

func Test3131_318E(t *testing.T) {
	ValidateRuneRange(t, 0x3131, 0x318E, pDISALLOWED)
}

func Test318F(t *testing.T) {
	ValidateRune(t, 0x318F, pUNASSIGNED)
}

func Test3190_319F(t *testing.T) {
	ValidateRuneRange(t, 0x3190, 0x319F, pDISALLOWED)
}

func Test31A0_31B7(t *testing.T) {
	ValidateRuneRange(t, 0x31A0, 0x31B7, pPVALID)
}

func Test31B8_31BF(t *testing.T) {
	ValidateRuneRange(t, 0x31B8, 0x31BF, pUNASSIGNED)
}

func Test31C0_31E3(t *testing.T) {
	ValidateRuneRange(t, 0x31C0, 0x31E3, pDISALLOWED)
}

func Test31E4_31EF(t *testing.T) {
	ValidateRuneRange(t, 0x31E4, 0x31EF, pUNASSIGNED)
}

func Test31F0_31FF(t *testing.T) {
	ValidateRuneRange(t, 0x31F0, 0x31FF, pPVALID)
}

func Test3200_321E(t *testing.T) {
	ValidateRuneRange(t, 0x3200, 0x321E, pDISALLOWED)
}

func Test321F(t *testing.T) {
	ValidateRune(t, 0x321F, pUNASSIGNED)
}

func Test3220_32FE(t *testing.T) {
	ValidateRuneRange(t, 0x3220, 0x32FE, pDISALLOWED)
}

func Test32FF(t *testing.T) {
	ValidateRune(t, 0x32FF, pUNASSIGNED)
}

func Test3300_33FF(t *testing.T) {
	ValidateRuneRange(t, 0x3300, 0x33FF, pDISALLOWED)
}

func Test3400_4DB5(t *testing.T) {
	ValidateRuneRange(t, 0x3400, 0x4DB5, pPVALID)
}

func Test4DB6_4DBF(t *testing.T) {
	ValidateRuneRange(t, 0x4DB6, 0x4DBF, pUNASSIGNED)
}

func Test4DC0_4DFF(t *testing.T) {
	ValidateRuneRange(t, 0x4DC0, 0x4DFF, pDISALLOWED)
}

func Test4E00_9FCB(t *testing.T) {
	ValidateRuneRange(t, 0x4E00, 0x9FCB, pPVALID)
}

func Test9FCC_9FFF(t *testing.T) {
	ValidateRuneRange(t, 0x9FCC, 0x9FFF, pUNASSIGNED)
}

func TestA000_A48C(t *testing.T) {
	ValidateRuneRange(t, 0xA000, 0xA48C, pPVALID)
}

func TestA48D_A48F(t *testing.T) {
	ValidateRuneRange(t, 0xA48D, 0xA48F, pUNASSIGNED)
}

func TestA490_A4C6(t *testing.T) {
	ValidateRuneRange(t, 0xA490, 0xA4C6, pDISALLOWED)
}

func TestA4C7_A4CF(t *testing.T) {
	ValidateRuneRange(t, 0xA4C7, 0xA4CF, pUNASSIGNED)
}

func TestA4D0_A4FD(t *testing.T) {
	ValidateRuneRange(t, 0xA4D0, 0xA4FD, pPVALID)
}

func TestA4FE_A4FF(t *testing.T) {
	ValidateRuneRange(t, 0xA4FE, 0xA4FF, pDISALLOWED)
}

func TestA500_A60C(t *testing.T) {
	ValidateRuneRange(t, 0xA500, 0xA60C, pPVALID)
}

func TestA60D_A60F(t *testing.T) {
	ValidateRuneRange(t, 0xA60D, 0xA60F, pDISALLOWED)
}

func TestA610_A62B(t *testing.T) {
	ValidateRuneRange(t, 0xA610, 0xA62B, pPVALID)
}

func TestA62C_A63F(t *testing.T) {
	ValidateRuneRange(t, 0xA62C, 0xA63F, pUNASSIGNED)
}

func TestA640(t *testing.T) {
	ValidateRune(t, 0xA640, pDISALLOWED)
}

func TestA641(t *testing.T) {
	ValidateRune(t, 0xA641, pPVALID)
}

func TestA642(t *testing.T) {
	ValidateRune(t, 0xA642, pDISALLOWED)
}

func TestA643(t *testing.T) {
	ValidateRune(t, 0xA643, pPVALID)
}

func TestA644(t *testing.T) {
	ValidateRune(t, 0xA644, pDISALLOWED)
}

func TestA645(t *testing.T) {
	ValidateRune(t, 0xA645, pPVALID)
}

func TestA646(t *testing.T) {
	ValidateRune(t, 0xA646, pDISALLOWED)
}

func TestA647(t *testing.T) {
	ValidateRune(t, 0xA647, pPVALID)
}

func TestA648(t *testing.T) {
	ValidateRune(t, 0xA648, pDISALLOWED)
}

func TestA649(t *testing.T) {
	ValidateRune(t, 0xA649, pPVALID)
}

func TestA64A(t *testing.T) {
	ValidateRune(t, 0xA64A, pDISALLOWED)
}

func TestA64B(t *testing.T) {
	ValidateRune(t, 0xA64B, pPVALID)
}

func TestA64C(t *testing.T) {
	ValidateRune(t, 0xA64C, pDISALLOWED)
}

func TestA64D(t *testing.T) {
	ValidateRune(t, 0xA64D, pPVALID)
}

func TestA64E(t *testing.T) {
	ValidateRune(t, 0xA64E, pDISALLOWED)
}

func TestA64F(t *testing.T) {
	ValidateRune(t, 0xA64F, pPVALID)
}

func TestA650(t *testing.T) {
	ValidateRune(t, 0xA650, pDISALLOWED)
}

func TestA651(t *testing.T) {
	ValidateRune(t, 0xA651, pPVALID)
}

func TestA652(t *testing.T) {
	ValidateRune(t, 0xA652, pDISALLOWED)
}

func TestA653(t *testing.T) {
	ValidateRune(t, 0xA653, pPVALID)
}

func TestA654(t *testing.T) {
	ValidateRune(t, 0xA654, pDISALLOWED)
}

func TestA655(t *testing.T) {
	ValidateRune(t, 0xA655, pPVALID)
}

func TestA656(t *testing.T) {
	ValidateRune(t, 0xA656, pDISALLOWED)
}

func TestA657(t *testing.T) {
	ValidateRune(t, 0xA657, pPVALID)
}

func TestA658(t *testing.T) {
	ValidateRune(t, 0xA658, pDISALLOWED)
}

func TestA659(t *testing.T) {
	ValidateRune(t, 0xA659, pPVALID)
}

func TestA65A(t *testing.T) {
	ValidateRune(t, 0xA65A, pDISALLOWED)
}

func TestA65B(t *testing.T) {
	ValidateRune(t, 0xA65B, pPVALID)
}

func TestA65C(t *testing.T) {
	ValidateRune(t, 0xA65C, pDISALLOWED)
}

func TestA65D(t *testing.T) {
	ValidateRune(t, 0xA65D, pPVALID)
}

func TestA65E(t *testing.T) {
	ValidateRune(t, 0xA65E, pDISALLOWED)
}

func TestA65F(t *testing.T) {
	ValidateRune(t, 0xA65F, pPVALID)
}

func TestA660_A661(t *testing.T) {
	ValidateRuneRange(t, 0xA660, 0xA661, pUNASSIGNED)
}

func TestA662(t *testing.T) {
	ValidateRune(t, 0xA662, pDISALLOWED)
}

func TestA663(t *testing.T) {
	ValidateRune(t, 0xA663, pPVALID)
}

func TestA664(t *testing.T) {
	ValidateRune(t, 0xA664, pDISALLOWED)
}

func TestA665(t *testing.T) {
	ValidateRune(t, 0xA665, pPVALID)
}

func TestA666(t *testing.T) {
	ValidateRune(t, 0xA666, pDISALLOWED)
}

func TestA667(t *testing.T) {
	ValidateRune(t, 0xA667, pPVALID)
}

func TestA668(t *testing.T) {
	ValidateRune(t, 0xA668, pDISALLOWED)
}

func TestA669(t *testing.T) {
	ValidateRune(t, 0xA669, pPVALID)
}

func TestA66A(t *testing.T) {
	ValidateRune(t, 0xA66A, pDISALLOWED)
}

func TestA66B(t *testing.T) {
	ValidateRune(t, 0xA66B, pPVALID)
}

func TestA66C(t *testing.T) {
	ValidateRune(t, 0xA66C, pDISALLOWED)
}

func TestA66D_A66F(t *testing.T) {
	ValidateRuneRange(t, 0xA66D, 0xA66F, pPVALID)
}

func TestA670_A673(t *testing.T) {
	ValidateRuneRange(t, 0xA670, 0xA673, pDISALLOWED)
}

func TestA674_A67B(t *testing.T) {
	ValidateRuneRange(t, 0xA674, 0xA67B, pUNASSIGNED)
}

func TestA67C_A67D(t *testing.T) {
	ValidateRuneRange(t, 0xA67C, 0xA67D, pPVALID)
}

func TestA67E(t *testing.T) {
	ValidateRune(t, 0xA67E, pDISALLOWED)
}

func TestA67F(t *testing.T) {
	ValidateRune(t, 0xA67F, pPVALID)
}

func TestA680(t *testing.T) {
	ValidateRune(t, 0xA680, pDISALLOWED)
}

func TestA681(t *testing.T) {
	ValidateRune(t, 0xA681, pPVALID)
}

func TestA682(t *testing.T) {
	ValidateRune(t, 0xA682, pDISALLOWED)
}

func TestA683(t *testing.T) {
	ValidateRune(t, 0xA683, pPVALID)
}

func TestA684(t *testing.T) {
	ValidateRune(t, 0xA684, pDISALLOWED)
}

func TestA685(t *testing.T) {
	ValidateRune(t, 0xA685, pPVALID)
}

func TestA686(t *testing.T) {
	ValidateRune(t, 0xA686, pDISALLOWED)
}

func TestA687(t *testing.T) {
	ValidateRune(t, 0xA687, pPVALID)
}

func TestA688(t *testing.T) {
	ValidateRune(t, 0xA688, pDISALLOWED)
}

func TestA689(t *testing.T) {
	ValidateRune(t, 0xA689, pPVALID)
}

func TestA68A(t *testing.T) {
	ValidateRune(t, 0xA68A, pDISALLOWED)
}

func TestA68B(t *testing.T) {
	ValidateRune(t, 0xA68B, pPVALID)
}

func TestA68C(t *testing.T) {
	ValidateRune(t, 0xA68C, pDISALLOWED)
}

func TestA68D(t *testing.T) {
	ValidateRune(t, 0xA68D, pPVALID)
}

func TestA68E(t *testing.T) {
	ValidateRune(t, 0xA68E, pDISALLOWED)
}

func TestA68F(t *testing.T) {
	ValidateRune(t, 0xA68F, pPVALID)
}

func TestA690(t *testing.T) {
	ValidateRune(t, 0xA690, pDISALLOWED)
}

func TestA691(t *testing.T) {
	ValidateRune(t, 0xA691, pPVALID)
}

func TestA692(t *testing.T) {
	ValidateRune(t, 0xA692, pDISALLOWED)
}

func TestA693(t *testing.T) {
	ValidateRune(t, 0xA693, pPVALID)
}

func TestA694(t *testing.T) {
	ValidateRune(t, 0xA694, pDISALLOWED)
}

func TestA695(t *testing.T) {
	ValidateRune(t, 0xA695, pPVALID)
}

func TestA696(t *testing.T) {
	ValidateRune(t, 0xA696, pDISALLOWED)
}

func TestA697(t *testing.T) {
	ValidateRune(t, 0xA697, pPVALID)
}

func TestA698_A69F(t *testing.T) {
	ValidateRuneRange(t, 0xA698, 0xA69F, pUNASSIGNED)
}

func TestA6A0_A6E5(t *testing.T) {
	ValidateRuneRange(t, 0xA6A0, 0xA6E5, pPVALID)
}

func TestA6E6_A6EF(t *testing.T) {
	ValidateRuneRange(t, 0xA6E6, 0xA6EF, pDISALLOWED)
}

func TestA6F0_A6F1(t *testing.T) {
	ValidateRuneRange(t, 0xA6F0, 0xA6F1, pPVALID)
}

func TestA6F2_A6F7(t *testing.T) {
	ValidateRuneRange(t, 0xA6F2, 0xA6F7, pDISALLOWED)
}

func TestA6F8_A6FF(t *testing.T) {
	ValidateRuneRange(t, 0xA6F8, 0xA6FF, pUNASSIGNED)
}

func TestA700_A716(t *testing.T) {
	ValidateRuneRange(t, 0xA700, 0xA716, pDISALLOWED)
}

func TestA717_A71F(t *testing.T) {
	ValidateRuneRange(t, 0xA717, 0xA71F, pPVALID)
}

func TestA720_A722(t *testing.T) {
	ValidateRuneRange(t, 0xA720, 0xA722, pDISALLOWED)
}

func TestA723(t *testing.T) {
	ValidateRune(t, 0xA723, pPVALID)
}

func TestA724(t *testing.T) {
	ValidateRune(t, 0xA724, pDISALLOWED)
}

func TestA725(t *testing.T) {
	ValidateRune(t, 0xA725, pPVALID)
}

func TestA726(t *testing.T) {
	ValidateRune(t, 0xA726, pDISALLOWED)
}

func TestA727(t *testing.T) {
	ValidateRune(t, 0xA727, pPVALID)
}

func TestA728(t *testing.T) {
	ValidateRune(t, 0xA728, pDISALLOWED)
}

func TestA729(t *testing.T) {
	ValidateRune(t, 0xA729, pPVALID)
}

func TestA72A(t *testing.T) {
	ValidateRune(t, 0xA72A, pDISALLOWED)
}

func TestA72B(t *testing.T) {
	ValidateRune(t, 0xA72B, pPVALID)
}

func TestA72C(t *testing.T) {
	ValidateRune(t, 0xA72C, pDISALLOWED)
}

func TestA72D(t *testing.T) {
	ValidateRune(t, 0xA72D, pPVALID)
}

func TestA72E(t *testing.T) {
	ValidateRune(t, 0xA72E, pDISALLOWED)
}

func TestA72F_A731(t *testing.T) {
	ValidateRuneRange(t, 0xA72F, 0xA731, pPVALID)
}

func TestA732(t *testing.T) {
	ValidateRune(t, 0xA732, pDISALLOWED)
}

func TestA733(t *testing.T) {
	ValidateRune(t, 0xA733, pPVALID)
}

func TestA734(t *testing.T) {
	ValidateRune(t, 0xA734, pDISALLOWED)
}

func TestA735(t *testing.T) {
	ValidateRune(t, 0xA735, pPVALID)
}

func TestA736(t *testing.T) {
	ValidateRune(t, 0xA736, pDISALLOWED)
}

func TestA737(t *testing.T) {
	ValidateRune(t, 0xA737, pPVALID)
}

func TestA738(t *testing.T) {
	ValidateRune(t, 0xA738, pDISALLOWED)
}

func TestA739(t *testing.T) {
	ValidateRune(t, 0xA739, pPVALID)
}

func TestA73A(t *testing.T) {
	ValidateRune(t, 0xA73A, pDISALLOWED)
}

func TestA73B(t *testing.T) {
	ValidateRune(t, 0xA73B, pPVALID)
}

func TestA73C(t *testing.T) {
	ValidateRune(t, 0xA73C, pDISALLOWED)
}

func TestA73D(t *testing.T) {
	ValidateRune(t, 0xA73D, pPVALID)
}

func TestA73E(t *testing.T) {
	ValidateRune(t, 0xA73E, pDISALLOWED)
}

func TestA73F(t *testing.T) {
	ValidateRune(t, 0xA73F, pPVALID)
}

func TestA740(t *testing.T) {
	ValidateRune(t, 0xA740, pDISALLOWED)
}

func TestA741(t *testing.T) {
	ValidateRune(t, 0xA741, pPVALID)
}

func TestA742(t *testing.T) {
	ValidateRune(t, 0xA742, pDISALLOWED)
}

func TestA743(t *testing.T) {
	ValidateRune(t, 0xA743, pPVALID)
}

func TestA744(t *testing.T) {
	ValidateRune(t, 0xA744, pDISALLOWED)
}

func TestA745(t *testing.T) {
	ValidateRune(t, 0xA745, pPVALID)
}

func TestA746(t *testing.T) {
	ValidateRune(t, 0xA746, pDISALLOWED)
}

func TestA747(t *testing.T) {
	ValidateRune(t, 0xA747, pPVALID)
}

func TestA748(t *testing.T) {
	ValidateRune(t, 0xA748, pDISALLOWED)
}

func TestA749(t *testing.T) {
	ValidateRune(t, 0xA749, pPVALID)
}

func TestA74A(t *testing.T) {
	ValidateRune(t, 0xA74A, pDISALLOWED)
}

func TestA74B(t *testing.T) {
	ValidateRune(t, 0xA74B, pPVALID)
}

func TestA74C(t *testing.T) {
	ValidateRune(t, 0xA74C, pDISALLOWED)
}

func TestA74D(t *testing.T) {
	ValidateRune(t, 0xA74D, pPVALID)
}

func TestA74E(t *testing.T) {
	ValidateRune(t, 0xA74E, pDISALLOWED)
}

func TestA74F(t *testing.T) {
	ValidateRune(t, 0xA74F, pPVALID)
}

func TestA750(t *testing.T) {
	ValidateRune(t, 0xA750, pDISALLOWED)
}

func TestA751(t *testing.T) {
	ValidateRune(t, 0xA751, pPVALID)
}

func TestA752(t *testing.T) {
	ValidateRune(t, 0xA752, pDISALLOWED)
}

func TestA753(t *testing.T) {
	ValidateRune(t, 0xA753, pPVALID)
}

func TestA754(t *testing.T) {
	ValidateRune(t, 0xA754, pDISALLOWED)
}

func TestA755(t *testing.T) {
	ValidateRune(t, 0xA755, pPVALID)
}

func TestA756(t *testing.T) {
	ValidateRune(t, 0xA756, pDISALLOWED)
}

func TestA757(t *testing.T) {
	ValidateRune(t, 0xA757, pPVALID)
}

func TestA758(t *testing.T) {
	ValidateRune(t, 0xA758, pDISALLOWED)
}

func TestA759(t *testing.T) {
	ValidateRune(t, 0xA759, pPVALID)
}

func TestA75A(t *testing.T) {
	ValidateRune(t, 0xA75A, pDISALLOWED)
}

func TestA75B(t *testing.T) {
	ValidateRune(t, 0xA75B, pPVALID)
}

func TestA75C(t *testing.T) {
	ValidateRune(t, 0xA75C, pDISALLOWED)
}

func TestA75D(t *testing.T) {
	ValidateRune(t, 0xA75D, pPVALID)
}

func TestA75E(t *testing.T) {
	ValidateRune(t, 0xA75E, pDISALLOWED)
}

func TestA75F(t *testing.T) {
	ValidateRune(t, 0xA75F, pPVALID)
}

func TestA760(t *testing.T) {
	ValidateRune(t, 0xA760, pDISALLOWED)
}

func TestA761(t *testing.T) {
	ValidateRune(t, 0xA761, pPVALID)
}

func TestA762(t *testing.T) {
	ValidateRune(t, 0xA762, pDISALLOWED)
}

func TestA763(t *testing.T) {
	ValidateRune(t, 0xA763, pPVALID)
}

func TestA764(t *testing.T) {
	ValidateRune(t, 0xA764, pDISALLOWED)
}

func TestA765(t *testing.T) {
	ValidateRune(t, 0xA765, pPVALID)
}

func TestA766(t *testing.T) {
	ValidateRune(t, 0xA766, pDISALLOWED)
}

func TestA767(t *testing.T) {
	ValidateRune(t, 0xA767, pPVALID)
}

func TestA768(t *testing.T) {
	ValidateRune(t, 0xA768, pDISALLOWED)
}

func TestA769(t *testing.T) {
	ValidateRune(t, 0xA769, pPVALID)
}

func TestA76A(t *testing.T) {
	ValidateRune(t, 0xA76A, pDISALLOWED)
}

func TestA76B(t *testing.T) {
	ValidateRune(t, 0xA76B, pPVALID)
}

func TestA76C(t *testing.T) {
	ValidateRune(t, 0xA76C, pDISALLOWED)
}

func TestA76D(t *testing.T) {
	ValidateRune(t, 0xA76D, pPVALID)
}

func TestA76E(t *testing.T) {
	ValidateRune(t, 0xA76E, pDISALLOWED)
}

func TestA76F(t *testing.T) {
	ValidateRune(t, 0xA76F, pPVALID)
}

func TestA770(t *testing.T) {
	ValidateRune(t, 0xA770, pDISALLOWED)
}

func TestA771_A778(t *testing.T) {
	ValidateRuneRange(t, 0xA771, 0xA778, pPVALID)
}

func TestA779(t *testing.T) {
	ValidateRune(t, 0xA779, pDISALLOWED)
}

func TestA77A(t *testing.T) {
	ValidateRune(t, 0xA77A, pPVALID)
}

func TestA77B(t *testing.T) {
	ValidateRune(t, 0xA77B, pDISALLOWED)
}

func TestA77C(t *testing.T) {
	ValidateRune(t, 0xA77C, pPVALID)
}

func TestA77D_A77E(t *testing.T) {
	ValidateRuneRange(t, 0xA77D, 0xA77E, pDISALLOWED)
}

func TestA77F(t *testing.T) {
	ValidateRune(t, 0xA77F, pPVALID)
}

func TestA780(t *testing.T) {
	ValidateRune(t, 0xA780, pDISALLOWED)
}

func TestA781(t *testing.T) {
	ValidateRune(t, 0xA781, pPVALID)
}

func TestA782(t *testing.T) {
	ValidateRune(t, 0xA782, pDISALLOWED)
}

func TestA783(t *testing.T) {
	ValidateRune(t, 0xA783, pPVALID)
}

func TestA784(t *testing.T) {
	ValidateRune(t, 0xA784, pDISALLOWED)
}

func TestA785(t *testing.T) {
	ValidateRune(t, 0xA785, pPVALID)
}

func TestA786(t *testing.T) {
	ValidateRune(t, 0xA786, pDISALLOWED)
}

func TestA787_A788(t *testing.T) {
	ValidateRuneRange(t, 0xA787, 0xA788, pPVALID)
}

func TestA789_A78B(t *testing.T) {
	ValidateRuneRange(t, 0xA789, 0xA78B, pDISALLOWED)
}

func TestA78C(t *testing.T) {
	ValidateRune(t, 0xA78C, pPVALID)
}

func TestA78D_A7FA(t *testing.T) {
	ValidateRuneRange(t, 0xA78D, 0xA7FA, pUNASSIGNED)
}

func TestA7FB_A827(t *testing.T) {
	ValidateRuneRange(t, 0xA7FB, 0xA827, pPVALID)
}

func TestA828_A82B(t *testing.T) {
	ValidateRuneRange(t, 0xA828, 0xA82B, pDISALLOWED)
}

func TestA82C_A82F(t *testing.T) {
	ValidateRuneRange(t, 0xA82C, 0xA82F, pUNASSIGNED)
}

func TestA830_A839(t *testing.T) {
	ValidateRuneRange(t, 0xA830, 0xA839, pDISALLOWED)
}

func TestA83A_A83F(t *testing.T) {
	ValidateRuneRange(t, 0xA83A, 0xA83F, pUNASSIGNED)
}

func TestA840_A873(t *testing.T) {
	ValidateRuneRange(t, 0xA840, 0xA873, pPVALID)
}

func TestA874_A877(t *testing.T) {
	ValidateRuneRange(t, 0xA874, 0xA877, pDISALLOWED)
}

func TestA878_A87F(t *testing.T) {
	ValidateRuneRange(t, 0xA878, 0xA87F, pUNASSIGNED)
}

func TestA880_A8C4(t *testing.T) {
	ValidateRuneRange(t, 0xA880, 0xA8C4, pPVALID)
}

func TestA8C5_A8CD(t *testing.T) {
	ValidateRuneRange(t, 0xA8C5, 0xA8CD, pUNASSIGNED)
}

func TestA8CE_A8CF(t *testing.T) {
	ValidateRuneRange(t, 0xA8CE, 0xA8CF, pDISALLOWED)
}

func TestA8D0_A8D9(t *testing.T) {
	ValidateRuneRange(t, 0xA8D0, 0xA8D9, pPVALID)
}

func TestA8DA_A8DF(t *testing.T) {
	ValidateRuneRange(t, 0xA8DA, 0xA8DF, pUNASSIGNED)
}

func TestA8E0_A8F7(t *testing.T) {
	ValidateRuneRange(t, 0xA8E0, 0xA8F7, pPVALID)
}

func TestA8F8_A8FA(t *testing.T) {
	ValidateRuneRange(t, 0xA8F8, 0xA8FA, pDISALLOWED)
}

func TestA8FB(t *testing.T) {
	ValidateRune(t, 0xA8FB, pPVALID)
}

func TestA8FC_A8FF(t *testing.T) {
	ValidateRuneRange(t, 0xA8FC, 0xA8FF, pUNASSIGNED)
}

func TestA900_A92D(t *testing.T) {
	ValidateRuneRange(t, 0xA900, 0xA92D, pPVALID)
}

func TestA92E_A92F(t *testing.T) {
	ValidateRuneRange(t, 0xA92E, 0xA92F, pDISALLOWED)
}

func TestA930_A953(t *testing.T) {
	ValidateRuneRange(t, 0xA930, 0xA953, pPVALID)
}

func TestA954_A95E(t *testing.T) {
	ValidateRuneRange(t, 0xA954, 0xA95E, pUNASSIGNED)
}

func TestA95F_A97C(t *testing.T) {
	ValidateRuneRange(t, 0xA95F, 0xA97C, pDISALLOWED)
}

func TestA97D_A97F(t *testing.T) {
	ValidateRuneRange(t, 0xA97D, 0xA97F, pUNASSIGNED)
}

func TestA980_A9C0(t *testing.T) {
	ValidateRuneRange(t, 0xA980, 0xA9C0, pPVALID)
}

func TestA9C1_A9CD(t *testing.T) {
	ValidateRuneRange(t, 0xA9C1, 0xA9CD, pDISALLOWED)
}

func TestA9CE(t *testing.T) {
	ValidateRune(t, 0xA9CE, pUNASSIGNED)
}

func TestA9CF_A9D9(t *testing.T) {
	ValidateRuneRange(t, 0xA9CF, 0xA9D9, pPVALID)
}

func TestA9DA_A9DD(t *testing.T) {
	ValidateRuneRange(t, 0xA9DA, 0xA9DD, pUNASSIGNED)
}

func TestA9DE_A9DF(t *testing.T) {
	ValidateRuneRange(t, 0xA9DE, 0xA9DF, pDISALLOWED)
}

func TestA9E0_A9FF(t *testing.T) {
	ValidateRuneRange(t, 0xA9E0, 0xA9FF, pUNASSIGNED)
}

func TestAA00_AA36(t *testing.T) {
	ValidateRuneRange(t, 0xAA00, 0xAA36, pPVALID)
}

func TestAA37_AA3F(t *testing.T) {
	ValidateRuneRange(t, 0xAA37, 0xAA3F, pUNASSIGNED)
}

func TestAA40_AA4D(t *testing.T) {
	ValidateRuneRange(t, 0xAA40, 0xAA4D, pPVALID)
}

func TestAA4E_AA4F(t *testing.T) {
	ValidateRuneRange(t, 0xAA4E, 0xAA4F, pUNASSIGNED)
}

func TestAA50_AA59(t *testing.T) {
	ValidateRuneRange(t, 0xAA50, 0xAA59, pPVALID)
}

func TestAA5A_AA5B(t *testing.T) {
	ValidateRuneRange(t, 0xAA5A, 0xAA5B, pUNASSIGNED)
}

func TestAA5C_AA5F(t *testing.T) {
	ValidateRuneRange(t, 0xAA5C, 0xAA5F, pDISALLOWED)
}

func TestAA60_AA76(t *testing.T) {
	ValidateRuneRange(t, 0xAA60, 0xAA76, pPVALID)
}

func TestAA77_AA79(t *testing.T) {
	ValidateRuneRange(t, 0xAA77, 0xAA79, pDISALLOWED)
}

func TestAA7A_AA7B(t *testing.T) {
	ValidateRuneRange(t, 0xAA7A, 0xAA7B, pPVALID)
}

func TestAA7C_AA7F(t *testing.T) {
	ValidateRuneRange(t, 0xAA7C, 0xAA7F, pUNASSIGNED)
}

func TestAA80_AAC2(t *testing.T) {
	ValidateRuneRange(t, 0xAA80, 0xAAC2, pPVALID)
}

func TestAAC3_AADA(t *testing.T) {
	ValidateRuneRange(t, 0xAAC3, 0xAADA, pUNASSIGNED)
}

func TestAADB_AADD(t *testing.T) {
	ValidateRuneRange(t, 0xAADB, 0xAADD, pPVALID)
}

func TestAADE_AADF(t *testing.T) {
	ValidateRuneRange(t, 0xAADE, 0xAADF, pDISALLOWED)
}

func TestAAE0_ABBF(t *testing.T) {
	ValidateRuneRange(t, 0xAAE0, 0xABBF, pUNASSIGNED)
}

func TestABC0_ABEA(t *testing.T) {
	ValidateRuneRange(t, 0xABC0, 0xABEA, pPVALID)
}

func TestABEB(t *testing.T) {
	ValidateRune(t, 0xABEB, pDISALLOWED)
}

func TestABEC_ABED(t *testing.T) {
	ValidateRuneRange(t, 0xABEC, 0xABED, pPVALID)
}

func TestABEE_ABEF(t *testing.T) {
	ValidateRuneRange(t, 0xABEE, 0xABEF, pUNASSIGNED)
}

func TestABF0_ABF9(t *testing.T) {
	ValidateRuneRange(t, 0xABF0, 0xABF9, pPVALID)
}

func TestABFA_ABFF(t *testing.T) {
	ValidateRuneRange(t, 0xABFA, 0xABFF, pUNASSIGNED)
}

func TestAC00_D7A3(t *testing.T) {
	ValidateRuneRange(t, 0xAC00, 0xD7A3, pPVALID)
}

func TestD7A4_D7AF(t *testing.T) {
	ValidateRuneRange(t, 0xD7A4, 0xD7AF, pUNASSIGNED)
}

func TestD7B0_D7C6(t *testing.T) {
	ValidateRuneRange(t, 0xD7B0, 0xD7C6, pDISALLOWED)
}

func TestD7C7_D7CA(t *testing.T) {
	ValidateRuneRange(t, 0xD7C7, 0xD7CA, pUNASSIGNED)
}

func TestD7CB_D7FB(t *testing.T) {
	ValidateRuneRange(t, 0xD7CB, 0xD7FB, pDISALLOWED)
}

func TestD7FC_D7FF(t *testing.T) {
	ValidateRuneRange(t, 0xD7FC, 0xD7FF, pUNASSIGNED)
}

func TestD800_FA0D(t *testing.T) {
	ValidateRuneRange(t, 0xD800, 0xFA0D, pDISALLOWED)
}

func TestFA0E_FA0F(t *testing.T) {
	ValidateRuneRange(t, 0xFA0E, 0xFA0F, pPVALID)
}

func TestFA10(t *testing.T) {
	ValidateRune(t, 0xFA10, pDISALLOWED)
}

func TestFA11(t *testing.T) {
	ValidateRune(t, 0xFA11, pPVALID)
}

func TestFA12(t *testing.T) {
	ValidateRune(t, 0xFA12, pDISALLOWED)
}

func TestFA13_FA14(t *testing.T) {
	ValidateRuneRange(t, 0xFA13, 0xFA14, pPVALID)
}

func TestFA15_FA1E(t *testing.T) {
	ValidateRuneRange(t, 0xFA15, 0xFA1E, pDISALLOWED)
}

func TestFA1F(t *testing.T) {
	ValidateRune(t, 0xFA1F, pPVALID)
}

func TestFA20(t *testing.T) {
	ValidateRune(t, 0xFA20, pDISALLOWED)
}

func TestFA21(t *testing.T) {
	ValidateRune(t, 0xFA21, pPVALID)
}

func TestFA22(t *testing.T) {
	ValidateRune(t, 0xFA22, pDISALLOWED)
}

func TestFA23_FA24(t *testing.T) {
	ValidateRuneRange(t, 0xFA23, 0xFA24, pPVALID)
}

func TestFA25_FA26(t *testing.T) {
	ValidateRuneRange(t, 0xFA25, 0xFA26, pDISALLOWED)
}

func TestFA27_FA29(t *testing.T) {
	ValidateRuneRange(t, 0xFA27, 0xFA29, pPVALID)
}

func TestFA2A_FA2D(t *testing.T) {
	ValidateRuneRange(t, 0xFA2A, 0xFA2D, pDISALLOWED)
}

func TestFA2E_FA2F(t *testing.T) {
	ValidateRuneRange(t, 0xFA2E, 0xFA2F, pUNASSIGNED)
}

func TestFA30_FA6D(t *testing.T) {
	ValidateRuneRange(t, 0xFA30, 0xFA6D, pDISALLOWED)
}

func TestFA6E_FA6F(t *testing.T) {
	ValidateRuneRange(t, 0xFA6E, 0xFA6F, pUNASSIGNED)
}

func TestFA70_FAD9(t *testing.T) {
	ValidateRuneRange(t, 0xFA70, 0xFAD9, pDISALLOWED)
}

func TestFADA_FAFF(t *testing.T) {
	ValidateRuneRange(t, 0xFADA, 0xFAFF, pUNASSIGNED)
}

func TestFB00_FB06(t *testing.T) {
	ValidateRuneRange(t, 0xFB00, 0xFB06, pDISALLOWED)
}

func TestFB07_FB12(t *testing.T) {
	ValidateRuneRange(t, 0xFB07, 0xFB12, pUNASSIGNED)
}

func TestFB13_FB17(t *testing.T) {
	ValidateRuneRange(t, 0xFB13, 0xFB17, pDISALLOWED)
}

func TestFB18_FB1C(t *testing.T) {
	ValidateRuneRange(t, 0xFB18, 0xFB1C, pUNASSIGNED)
}

func TestFB1D(t *testing.T) {
	ValidateRune(t, 0xFB1D, pDISALLOWED)
}

func TestFB1E(t *testing.T) {
	ValidateRune(t, 0xFB1E, pPVALID)
}

func TestFB1F_FB36(t *testing.T) {
	ValidateRuneRange(t, 0xFB1F, 0xFB36, pDISALLOWED)
}

func TestFB37(t *testing.T) {
	ValidateRune(t, 0xFB37, pUNASSIGNED)
}

func TestFB38_FB3C(t *testing.T) {
	ValidateRuneRange(t, 0xFB38, 0xFB3C, pDISALLOWED)
}

func TestFB3D(t *testing.T) {
	ValidateRune(t, 0xFB3D, pUNASSIGNED)
}

func TestFB3E(t *testing.T) {
	ValidateRune(t, 0xFB3E, pDISALLOWED)
}

func TestFB3F(t *testing.T) {
	ValidateRune(t, 0xFB3F, pUNASSIGNED)
}

func TestFB40_FB41(t *testing.T) {
	ValidateRuneRange(t, 0xFB40, 0xFB41, pDISALLOWED)
}

func TestFB42(t *testing.T) {
	ValidateRune(t, 0xFB42, pUNASSIGNED)
}

func TestFB43_FB44(t *testing.T) {
	ValidateRuneRange(t, 0xFB43, 0xFB44, pDISALLOWED)
}

func TestFB45(t *testing.T) {
	ValidateRune(t, 0xFB45, pUNASSIGNED)
}

func TestFB46_FBB1(t *testing.T) {
	ValidateRuneRange(t, 0xFB46, 0xFBB1, pDISALLOWED)
}

func TestFBB2_FBD2(t *testing.T) {
	ValidateRuneRange(t, 0xFBB2, 0xFBD2, pUNASSIGNED)
}

func TestFBD3_FD3F(t *testing.T) {
	ValidateRuneRange(t, 0xFBD3, 0xFD3F, pDISALLOWED)
}

func TestFD40_FD4F(t *testing.T) {
	ValidateRuneRange(t, 0xFD40, 0xFD4F, pUNASSIGNED)
}

func TestFD50_FD8F(t *testing.T) {
	ValidateRuneRange(t, 0xFD50, 0xFD8F, pDISALLOWED)
}

func TestFD90_FD91(t *testing.T) {
	ValidateRuneRange(t, 0xFD90, 0xFD91, pUNASSIGNED)
}

func TestFD92_FDC7(t *testing.T) {
	ValidateRuneRange(t, 0xFD92, 0xFDC7, pDISALLOWED)
}

func TestFDC8_FDCF(t *testing.T) {
	ValidateRuneRange(t, 0xFDC8, 0xFDCF, pUNASSIGNED)
}

func TestFDD0_FDFD(t *testing.T) {
	ValidateRuneRange(t, 0xFDD0, 0xFDFD, pDISALLOWED)
}

func TestFDFE_FDFF(t *testing.T) {
	ValidateRuneRange(t, 0xFDFE, 0xFDFF, pUNASSIGNED)
}

func TestFE00_FE19(t *testing.T) {
	ValidateRuneRange(t, 0xFE00, 0xFE19, pDISALLOWED)
}

func TestFE1A_FE1F(t *testing.T) {
	ValidateRuneRange(t, 0xFE1A, 0xFE1F, pUNASSIGNED)
}

func TestFE20_FE26(t *testing.T) {
	ValidateRuneRange(t, 0xFE20, 0xFE26, pPVALID)
}

func TestFE27_FE2F(t *testing.T) {
	ValidateRuneRange(t, 0xFE27, 0xFE2F, pUNASSIGNED)
}

func TestFE30_FE52(t *testing.T) {
	ValidateRuneRange(t, 0xFE30, 0xFE52, pDISALLOWED)
}

func TestFE53(t *testing.T) {
	ValidateRune(t, 0xFE53, pUNASSIGNED)
}

func TestFE54_FE66(t *testing.T) {
	ValidateRuneRange(t, 0xFE54, 0xFE66, pDISALLOWED)
}

func TestFE67(t *testing.T) {
	ValidateRune(t, 0xFE67, pUNASSIGNED)
}

func TestFE68_FE6B(t *testing.T) {
	ValidateRuneRange(t, 0xFE68, 0xFE6B, pDISALLOWED)
}

func TestFE6C_FE6F(t *testing.T) {
	ValidateRuneRange(t, 0xFE6C, 0xFE6F, pUNASSIGNED)
}

func TestFE70_FE72(t *testing.T) {
	ValidateRuneRange(t, 0xFE70, 0xFE72, pDISALLOWED)
}

func TestFE73(t *testing.T) {
	ValidateRune(t, 0xFE73, pPVALID)
}

func TestFE74(t *testing.T) {
	ValidateRune(t, 0xFE74, pDISALLOWED)
}

func TestFE75(t *testing.T) {
	ValidateRune(t, 0xFE75, pUNASSIGNED)
}

func TestFE76_FEFC(t *testing.T) {
	ValidateRuneRange(t, 0xFE76, 0xFEFC, pDISALLOWED)
}

func TestFEFD_FEFE(t *testing.T) {
	ValidateRuneRange(t, 0xFEFD, 0xFEFE, pUNASSIGNED)
}

func TestFEFF(t *testing.T) {
	ValidateRune(t, 0xFEFF, pDISALLOWED)
}

func TestFF00(t *testing.T) {
	ValidateRune(t, 0xFF00, pUNASSIGNED)
}

func TestFF01_FFBE(t *testing.T) {
	ValidateRuneRange(t, 0xFF01, 0xFFBE, pDISALLOWED)
}

func TestFFBF_FFC1(t *testing.T) {
	ValidateRuneRange(t, 0xFFBF, 0xFFC1, pUNASSIGNED)
}

func TestFFC2_FFC7(t *testing.T) {
	ValidateRuneRange(t, 0xFFC2, 0xFFC7, pDISALLOWED)
}

func TestFFC8_FFC9(t *testing.T) {
	ValidateRuneRange(t, 0xFFC8, 0xFFC9, pUNASSIGNED)
}

func TestFFCA_FFCF(t *testing.T) {
	ValidateRuneRange(t, 0xFFCA, 0xFFCF, pDISALLOWED)
}

func TestFFD0_FFD1(t *testing.T) {
	ValidateRuneRange(t, 0xFFD0, 0xFFD1, pUNASSIGNED)
}

func TestFFD2_FFD7(t *testing.T) {
	ValidateRuneRange(t, 0xFFD2, 0xFFD7, pDISALLOWED)
}

func TestFFD8_FFD9(t *testing.T) {
	ValidateRuneRange(t, 0xFFD8, 0xFFD9, pUNASSIGNED)
}

func TestFFDA_FFDC(t *testing.T) {
	ValidateRuneRange(t, 0xFFDA, 0xFFDC, pDISALLOWED)
}

func TestFFDD_FFDF(t *testing.T) {
	ValidateRuneRange(t, 0xFFDD, 0xFFDF, pUNASSIGNED)
}

func TestFFE0_FFE6(t *testing.T) {
	ValidateRuneRange(t, 0xFFE0, 0xFFE6, pDISALLOWED)
}

func TestFFE7(t *testing.T) {
	ValidateRune(t, 0xFFE7, pUNASSIGNED)
}

func TestFFE8_FFEE(t *testing.T) {
	ValidateRuneRange(t, 0xFFE8, 0xFFEE, pDISALLOWED)
}

func TestFFEF_FFF8(t *testing.T) {
	ValidateRuneRange(t, 0xFFEF, 0xFFF8, pUNASSIGNED)
}

func TestFFF9_FFFF(t *testing.T) {
	ValidateRuneRange(t, 0xFFF9, 0xFFFF, pDISALLOWED)
}

func Test10000_1000B(t *testing.T) {
	ValidateRuneRange(t, 0x10000, 0x1000B, pPVALID)
}

func Test1000C(t *testing.T) {
	ValidateRune(t, 0x1000C, pUNASSIGNED)
}

func Test1000D_10026(t *testing.T) {
	ValidateRuneRange(t, 0x1000D, 0x10026, pPVALID)
}

func Test10028_1003A(t *testing.T) {
	ValidateRuneRange(t, 0x10028, 0x1003A, pPVALID)
}

func Test1003B(t *testing.T) {
	ValidateRune(t, 0x1003B, pUNASSIGNED)
}

func Test1003C_1003D(t *testing.T) {
	ValidateRuneRange(t, 0x1003C, 0x1003D, pPVALID)
}

func Test1003E(t *testing.T) {
	ValidateRune(t, 0x1003E, pUNASSIGNED)
}

func Test1003F_1004D(t *testing.T) {
	ValidateRuneRange(t, 0x1003F, 0x1004D, pPVALID)
}

func Test1004E_1004F(t *testing.T) {
	ValidateRuneRange(t, 0x1004E, 0x1004F, pUNASSIGNED)
}

func Test10050_1005D(t *testing.T) {
	ValidateRuneRange(t, 0x10050, 0x1005D, pPVALID)
}

func Test1005E_1007F(t *testing.T) {
	ValidateRuneRange(t, 0x1005E, 0x1007F, pUNASSIGNED)
}

func Test10080_100FA(t *testing.T) {
	ValidateRuneRange(t, 0x10080, 0x100FA, pPVALID)
}

func Test100FB_100FF(t *testing.T) {
	ValidateRuneRange(t, 0x100FB, 0x100FF, pUNASSIGNED)
}

func Test10100_10102(t *testing.T) {
	ValidateRuneRange(t, 0x10100, 0x10102, pDISALLOWED)
}

func Test10103_10106(t *testing.T) {
	ValidateRuneRange(t, 0x10103, 0x10106, pUNASSIGNED)
}

func Test10107_10133(t *testing.T) {
	ValidateRuneRange(t, 0x10107, 0x10133, pDISALLOWED)
}

func Test10134_10136(t *testing.T) {
	ValidateRuneRange(t, 0x10134, 0x10136, pUNASSIGNED)
}

func Test10137_1018A(t *testing.T) {
	ValidateRuneRange(t, 0x10137, 0x1018A, pDISALLOWED)
}

func Test1018B_1018F(t *testing.T) {
	ValidateRuneRange(t, 0x1018B, 0x1018F, pUNASSIGNED)
}

func Test10190_1019B(t *testing.T) {
	ValidateRuneRange(t, 0x10190, 0x1019B, pDISALLOWED)
}

func Test1019C_101CF(t *testing.T) {
	ValidateRuneRange(t, 0x1019C, 0x101CF, pUNASSIGNED)
}

func Test101D0_101FC(t *testing.T) {
	ValidateRuneRange(t, 0x101D0, 0x101FC, pDISALLOWED)
}

func Test101FD(t *testing.T) {
	ValidateRune(t, 0x101FD, pPVALID)
}

func Test101FE_1027F(t *testing.T) {
	ValidateRuneRange(t, 0x101FE, 0x1027F, pUNASSIGNED)
}

func Test10280_1029C(t *testing.T) {
	ValidateRuneRange(t, 0x10280, 0x1029C, pPVALID)
}

func Test1029D_1029F(t *testing.T) {
	ValidateRuneRange(t, 0x1029D, 0x1029F, pUNASSIGNED)
}

func Test102A0_102D0(t *testing.T) {
	ValidateRuneRange(t, 0x102A0, 0x102D0, pPVALID)
}

func Test102D1_102FF(t *testing.T) {
	ValidateRuneRange(t, 0x102D1, 0x102FF, pUNASSIGNED)
}

func Test10300_1031E(t *testing.T) {
	ValidateRuneRange(t, 0x10300, 0x1031E, pPVALID)
}

func Test1031F(t *testing.T) {
	ValidateRune(t, 0x1031F, pUNASSIGNED)
}

func Test10320_10323(t *testing.T) {
	ValidateRuneRange(t, 0x10320, 0x10323, pDISALLOWED)
}

func Test10324_1032F(t *testing.T) {
	ValidateRuneRange(t, 0x10324, 0x1032F, pUNASSIGNED)
}

func Test10330_10340(t *testing.T) {
	ValidateRuneRange(t, 0x10330, 0x10340, pPVALID)
}

func Test10342_10349(t *testing.T) {
	ValidateRuneRange(t, 0x10342, 0x10349, pPVALID)
}

func Test1034A(t *testing.T) {
	ValidateRune(t, 0x1034A, pDISALLOWED)
}

func Test1034B_1037F(t *testing.T) {
	ValidateRuneRange(t, 0x1034B, 0x1037F, pUNASSIGNED)
}

func Test10380_1039D(t *testing.T) {
	ValidateRuneRange(t, 0x10380, 0x1039D, pPVALID)
}

func Test1039E(t *testing.T) {
	ValidateRune(t, 0x1039E, pUNASSIGNED)
}

func Test1039F(t *testing.T) {
	ValidateRune(t, 0x1039F, pDISALLOWED)
}

func Test103A0_103C3(t *testing.T) {
	ValidateRuneRange(t, 0x103A0, 0x103C3, pPVALID)
}

func Test103C4_103C7(t *testing.T) {
	ValidateRuneRange(t, 0x103C4, 0x103C7, pUNASSIGNED)
}

func Test103C8_103CF(t *testing.T) {
	ValidateRuneRange(t, 0x103C8, 0x103CF, pPVALID)
}

func Test103D0_103D5(t *testing.T) {
	ValidateRuneRange(t, 0x103D0, 0x103D5, pDISALLOWED)
}

func Test103D6_103FF(t *testing.T) {
	ValidateRuneRange(t, 0x103D6, 0x103FF, pUNASSIGNED)
}

func Test10400_10427(t *testing.T) {
	ValidateRuneRange(t, 0x10400, 0x10427, pDISALLOWED)
}

func Test10428_1049D(t *testing.T) {
	ValidateRuneRange(t, 0x10428, 0x1049D, pPVALID)
}

func Test1049E_1049F(t *testing.T) {
	ValidateRuneRange(t, 0x1049E, 0x1049F, pUNASSIGNED)
}

func Test104A0_104A9(t *testing.T) {
	ValidateRuneRange(t, 0x104A0, 0x104A9, pPVALID)
}

func Test104AA_107FF(t *testing.T) {
	ValidateRuneRange(t, 0x104AA, 0x107FF, pUNASSIGNED)
}

func Test10800_10805(t *testing.T) {
	ValidateRuneRange(t, 0x10800, 0x10805, pPVALID)
}

func Test10806_10807(t *testing.T) {
	ValidateRuneRange(t, 0x10806, 0x10807, pUNASSIGNED)
}

func Test1080A_10835(t *testing.T) {
	ValidateRuneRange(t, 0x1080A, 0x10835, pPVALID)
}

func Test10837_10838(t *testing.T) {
	ValidateRuneRange(t, 0x10837, 0x10838, pPVALID)
}

func Test10839_1083B(t *testing.T) {
	ValidateRuneRange(t, 0x10839, 0x1083B, pUNASSIGNED)
}

func Test1083C(t *testing.T) {
	ValidateRune(t, 0x1083C, pPVALID)
}

func Test1083D_1083E(t *testing.T) {
	ValidateRuneRange(t, 0x1083D, 0x1083E, pUNASSIGNED)
}

func Test1083F_10855(t *testing.T) {
	ValidateRuneRange(t, 0x1083F, 0x10855, pPVALID)
}

func Test10857_1085F(t *testing.T) {
	ValidateRuneRange(t, 0x10857, 0x1085F, pDISALLOWED)
}

func Test10860_108FF(t *testing.T) {
	ValidateRuneRange(t, 0x10860, 0x108FF, pUNASSIGNED)
}

func Test10900_10915(t *testing.T) {
	ValidateRuneRange(t, 0x10900, 0x10915, pPVALID)
}

func Test10916_1091B(t *testing.T) {
	ValidateRuneRange(t, 0x10916, 0x1091B, pDISALLOWED)
}

func Test1091C_1091E(t *testing.T) {
	ValidateRuneRange(t, 0x1091C, 0x1091E, pUNASSIGNED)
}

func Test1091F(t *testing.T) {
	ValidateRune(t, 0x1091F, pDISALLOWED)
}

func Test10920_10939(t *testing.T) {
	ValidateRuneRange(t, 0x10920, 0x10939, pPVALID)
}

func Test1093A_1093E(t *testing.T) {
	ValidateRuneRange(t, 0x1093A, 0x1093E, pUNASSIGNED)
}

func Test1093F(t *testing.T) {
	ValidateRune(t, 0x1093F, pDISALLOWED)
}

func Test10940_109FF(t *testing.T) {
	ValidateRuneRange(t, 0x10940, 0x109FF, pUNASSIGNED)
}

func Test10A00_10A03(t *testing.T) {
	ValidateRuneRange(t, 0x10A00, 0x10A03, pPVALID)
}

func Test10A04(t *testing.T) {
	ValidateRune(t, 0x10A04, pUNASSIGNED)
}

func Test10A05_10A06(t *testing.T) {
	ValidateRuneRange(t, 0x10A05, 0x10A06, pPVALID)
}

func Test10A07_10A0B(t *testing.T) {
	ValidateRuneRange(t, 0x10A07, 0x10A0B, pUNASSIGNED)
}

func Test10A0C_10A13(t *testing.T) {
	ValidateRuneRange(t, 0x10A0C, 0x10A13, pPVALID)
}

func Test10A14(t *testing.T) {
	ValidateRune(t, 0x10A14, pUNASSIGNED)
}

func Test10A15_10A17(t *testing.T) {
	ValidateRuneRange(t, 0x10A15, 0x10A17, pPVALID)
}

func Test10A18(t *testing.T) {
	ValidateRune(t, 0x10A18, pUNASSIGNED)
}

func Test10A19_10A33(t *testing.T) {
	ValidateRuneRange(t, 0x10A19, 0x10A33, pPVALID)
}

func Test10A34_10A37(t *testing.T) {
	ValidateRuneRange(t, 0x10A34, 0x10A37, pUNASSIGNED)
}

func Test10A38_10A3A(t *testing.T) {
	ValidateRuneRange(t, 0x10A38, 0x10A3A, pPVALID)
}

func Test10A3B_10A3E(t *testing.T) {
	ValidateRuneRange(t, 0x10A3B, 0x10A3E, pUNASSIGNED)
}

func Test10A3F(t *testing.T) {
	ValidateRune(t, 0x10A3F, pPVALID)
}

func Test10A40_10A47(t *testing.T) {
	ValidateRuneRange(t, 0x10A40, 0x10A47, pDISALLOWED)
}

func Test10A48_10A4F(t *testing.T) {
	ValidateRuneRange(t, 0x10A48, 0x10A4F, pUNASSIGNED)
}

func Test10A50_10A58(t *testing.T) {
	ValidateRuneRange(t, 0x10A50, 0x10A58, pDISALLOWED)
}

func Test10A59_10A5F(t *testing.T) {
	ValidateRuneRange(t, 0x10A59, 0x10A5F, pUNASSIGNED)
}

func Test10A60_10A7C(t *testing.T) {
	ValidateRuneRange(t, 0x10A60, 0x10A7C, pPVALID)
}

func Test10A7D_10A7F(t *testing.T) {
	ValidateRuneRange(t, 0x10A7D, 0x10A7F, pDISALLOWED)
}

func Test10A80_10AFF(t *testing.T) {
	ValidateRuneRange(t, 0x10A80, 0x10AFF, pUNASSIGNED)
}

func Test10B00_10B35(t *testing.T) {
	ValidateRuneRange(t, 0x10B00, 0x10B35, pPVALID)
}

func Test10B36_10B38(t *testing.T) {
	ValidateRuneRange(t, 0x10B36, 0x10B38, pUNASSIGNED)
}

func Test10B39_10B3F(t *testing.T) {
	ValidateRuneRange(t, 0x10B39, 0x10B3F, pDISALLOWED)
}

func Test10B40_10B55(t *testing.T) {
	ValidateRuneRange(t, 0x10B40, 0x10B55, pPVALID)
}

func Test10B56_10B57(t *testing.T) {
	ValidateRuneRange(t, 0x10B56, 0x10B57, pUNASSIGNED)
}

func Test10B58_10B5F(t *testing.T) {
	ValidateRuneRange(t, 0x10B58, 0x10B5F, pDISALLOWED)
}

func Test10B60_10B72(t *testing.T) {
	ValidateRuneRange(t, 0x10B60, 0x10B72, pPVALID)
}

func Test10B73_10B77(t *testing.T) {
	ValidateRuneRange(t, 0x10B73, 0x10B77, pUNASSIGNED)
}

func Test10B78_10B7F(t *testing.T) {
	ValidateRuneRange(t, 0x10B78, 0x10B7F, pDISALLOWED)
}

func Test10B80_10BFF(t *testing.T) {
	ValidateRuneRange(t, 0x10B80, 0x10BFF, pUNASSIGNED)
}

func Test10C00_10C48(t *testing.T) {
	ValidateRuneRange(t, 0x10C00, 0x10C48, pPVALID)
}

func Test10C49_10E5F(t *testing.T) {
	ValidateRuneRange(t, 0x10C49, 0x10E5F, pUNASSIGNED)
}

func Test10E60_10E7E(t *testing.T) {
	ValidateRuneRange(t, 0x10E60, 0x10E7E, pDISALLOWED)
}

func Test10E7F_1107F(t *testing.T) {
	ValidateRuneRange(t, 0x10E7F, 0x1107F, pUNASSIGNED)
}

func Test11080_110BA(t *testing.T) {
	ValidateRuneRange(t, 0x11080, 0x110BA, pPVALID)
}

func Test110BB_110C1(t *testing.T) {
	ValidateRuneRange(t, 0x110BB, 0x110C1, pDISALLOWED)
}

func Test110C2_11FFF(t *testing.T) {
	ValidateRuneRange(t, 0x110C2, 0x11FFF, pUNASSIGNED)
}

func Test12000_1236E(t *testing.T) {
	ValidateRuneRange(t, 0x12000, 0x1236E, pPVALID)
}

func Test1236F_123FF(t *testing.T) {
	ValidateRuneRange(t, 0x1236F, 0x123FF, pUNASSIGNED)
}

func Test12400_12462(t *testing.T) {
	ValidateRuneRange(t, 0x12400, 0x12462, pDISALLOWED)
}

func Test12463_1246F(t *testing.T) {
	ValidateRuneRange(t, 0x12463, 0x1246F, pUNASSIGNED)
}

func Test12470_12473(t *testing.T) {
	ValidateRuneRange(t, 0x12470, 0x12473, pDISALLOWED)
}

func Test12474_12FFF(t *testing.T) {
	ValidateRuneRange(t, 0x12474, 0x12FFF, pUNASSIGNED)
}

func Test13000_1342E(t *testing.T) {
	ValidateRuneRange(t, 0x13000, 0x1342E, pPVALID)
}

func Test1342F_1CFFF(t *testing.T) {
	ValidateRuneRange(t, 0x1342F, 0x1CFFF, pUNASSIGNED)
}

func Test1D000_1D0F5(t *testing.T) {
	ValidateRuneRange(t, 0x1D000, 0x1D0F5, pDISALLOWED)
}

func Test1D0F6_1D0FF(t *testing.T) {
	ValidateRuneRange(t, 0x1D0F6, 0x1D0FF, pUNASSIGNED)
}

func Test1D100_1D126(t *testing.T) {
	ValidateRuneRange(t, 0x1D100, 0x1D126, pDISALLOWED)
}

func Test1D127_1D128(t *testing.T) {
	ValidateRuneRange(t, 0x1D127, 0x1D128, pUNASSIGNED)
}

func Test1D129_1D1DD(t *testing.T) {
	ValidateRuneRange(t, 0x1D129, 0x1D1DD, pDISALLOWED)
}

func Test1D1DE_1D1FF(t *testing.T) {
	ValidateRuneRange(t, 0x1D1DE, 0x1D1FF, pUNASSIGNED)
}

func Test1D200_1D245(t *testing.T) {
	ValidateRuneRange(t, 0x1D200, 0x1D245, pDISALLOWED)
}

func Test1D246_1D2FF(t *testing.T) {
	ValidateRuneRange(t, 0x1D246, 0x1D2FF, pUNASSIGNED)
}

func Test1D300_1D356(t *testing.T) {
	ValidateRuneRange(t, 0x1D300, 0x1D356, pDISALLOWED)
}

func Test1D357_1D35F(t *testing.T) {
	ValidateRuneRange(t, 0x1D357, 0x1D35F, pUNASSIGNED)
}

func Test1D360_1D371(t *testing.T) {
	ValidateRuneRange(t, 0x1D360, 0x1D371, pDISALLOWED)
}

func Test1D372_1D3FF(t *testing.T) {
	ValidateRuneRange(t, 0x1D372, 0x1D3FF, pUNASSIGNED)
}

func Test1D400_1D454(t *testing.T) {
	ValidateRuneRange(t, 0x1D400, 0x1D454, pDISALLOWED)
}

func Test1D455(t *testing.T) {
	ValidateRune(t, 0x1D455, pUNASSIGNED)
}

func Test1D456_1D49C(t *testing.T) {
	ValidateRuneRange(t, 0x1D456, 0x1D49C, pDISALLOWED)
}

func Test1D49D(t *testing.T) {
	ValidateRune(t, 0x1D49D, pUNASSIGNED)
}

func Test1D49E_1D49F(t *testing.T) {
	ValidateRuneRange(t, 0x1D49E, 0x1D49F, pDISALLOWED)
}

func Test1D4A0_1D4A1(t *testing.T) {
	ValidateRuneRange(t, 0x1D4A0, 0x1D4A1, pUNASSIGNED)
}

func Test1D4A2(t *testing.T) {
	ValidateRune(t, 0x1D4A2, pDISALLOWED)
}

func Test1D4A3_1D4A4(t *testing.T) {
	ValidateRuneRange(t, 0x1D4A3, 0x1D4A4, pUNASSIGNED)
}

func Test1D4A5_1D4A6(t *testing.T) {
	ValidateRuneRange(t, 0x1D4A5, 0x1D4A6, pDISALLOWED)
}

func Test1D4A7_1D4A8(t *testing.T) {
	ValidateRuneRange(t, 0x1D4A7, 0x1D4A8, pUNASSIGNED)
}

func Test1D4A9_1D4AC(t *testing.T) {
	ValidateRuneRange(t, 0x1D4A9, 0x1D4AC, pDISALLOWED)
}

func Test1D4AD(t *testing.T) {
	ValidateRune(t, 0x1D4AD, pUNASSIGNED)
}

func Test1D4AE_1D4B9(t *testing.T) {
	ValidateRuneRange(t, 0x1D4AE, 0x1D4B9, pDISALLOWED)
}

func Test1D4BA(t *testing.T) {
	ValidateRune(t, 0x1D4BA, pUNASSIGNED)
}

func Test1D4BB(t *testing.T) {
	ValidateRune(t, 0x1D4BB, pDISALLOWED)
}

func Test1D4BC(t *testing.T) {
	ValidateRune(t, 0x1D4BC, pUNASSIGNED)
}

func Test1D4BD_1D4C3(t *testing.T) {
	ValidateRuneRange(t, 0x1D4BD, 0x1D4C3, pDISALLOWED)
}

func Test1D4C4(t *testing.T) {
	ValidateRune(t, 0x1D4C4, pUNASSIGNED)
}

func Test1D4C5_1D505(t *testing.T) {
	ValidateRuneRange(t, 0x1D4C5, 0x1D505, pDISALLOWED)
}

func Test1D506(t *testing.T) {
	ValidateRune(t, 0x1D506, pUNASSIGNED)
}

func Test1D507_1D50A(t *testing.T) {
	ValidateRuneRange(t, 0x1D507, 0x1D50A, pDISALLOWED)
}

func Test1D50B_1D50C(t *testing.T) {
	ValidateRuneRange(t, 0x1D50B, 0x1D50C, pUNASSIGNED)
}

func Test1D50D_1D514(t *testing.T) {
	ValidateRuneRange(t, 0x1D50D, 0x1D514, pDISALLOWED)
}

func Test1D515(t *testing.T) {
	ValidateRune(t, 0x1D515, pUNASSIGNED)
}

func Test1D516_1D51C(t *testing.T) {
	ValidateRuneRange(t, 0x1D516, 0x1D51C, pDISALLOWED)
}

func Test1D51D(t *testing.T) {
	ValidateRune(t, 0x1D51D, pUNASSIGNED)
}

func Test1D51E_1D539(t *testing.T) {
	ValidateRuneRange(t, 0x1D51E, 0x1D539, pDISALLOWED)
}

func Test1D53A(t *testing.T) {
	ValidateRune(t, 0x1D53A, pUNASSIGNED)
}

func Test1D53B_1D53E(t *testing.T) {
	ValidateRuneRange(t, 0x1D53B, 0x1D53E, pDISALLOWED)
}

func Test1D53F(t *testing.T) {
	ValidateRune(t, 0x1D53F, pUNASSIGNED)
}

func Test1D540_1D544(t *testing.T) {
	ValidateRuneRange(t, 0x1D540, 0x1D544, pDISALLOWED)
}

func Test1D545(t *testing.T) {
	ValidateRune(t, 0x1D545, pUNASSIGNED)
}

func Test1D546(t *testing.T) {
	ValidateRune(t, 0x1D546, pDISALLOWED)
}

func Test1D547_1D549(t *testing.T) {
	ValidateRuneRange(t, 0x1D547, 0x1D549, pUNASSIGNED)
}

func Test1D54A_1D550(t *testing.T) {
	ValidateRuneRange(t, 0x1D54A, 0x1D550, pDISALLOWED)
}

func Test1D551(t *testing.T) {
	ValidateRune(t, 0x1D551, pUNASSIGNED)
}

func Test1D552_1D6A5(t *testing.T) {
	ValidateRuneRange(t, 0x1D552, 0x1D6A5, pDISALLOWED)
}

func Test1D6A6_1D6A7(t *testing.T) {
	ValidateRuneRange(t, 0x1D6A6, 0x1D6A7, pUNASSIGNED)
}

func Test1D6A8_1D7CB(t *testing.T) {
	ValidateRuneRange(t, 0x1D6A8, 0x1D7CB, pDISALLOWED)
}

func Test1D7CC_1D7CD(t *testing.T) {
	ValidateRuneRange(t, 0x1D7CC, 0x1D7CD, pUNASSIGNED)
}

func Test1D7CE_1D7FF(t *testing.T) {
	ValidateRuneRange(t, 0x1D7CE, 0x1D7FF, pDISALLOWED)
}

func Test1D800_1EFFF(t *testing.T) {
	ValidateRuneRange(t, 0x1D800, 0x1EFFF, pUNASSIGNED)
}

func Test1F000_1F02B(t *testing.T) {
	ValidateRuneRange(t, 0x1F000, 0x1F02B, pDISALLOWED)
}

func Test1F02C_1F02F(t *testing.T) {
	ValidateRuneRange(t, 0x1F02C, 0x1F02F, pUNASSIGNED)
}

func Test1F030_1F093(t *testing.T) {
	ValidateRuneRange(t, 0x1F030, 0x1F093, pDISALLOWED)
}

func Test1F094_1F0FF(t *testing.T) {
	ValidateRuneRange(t, 0x1F094, 0x1F0FF, pUNASSIGNED)
}

func Test1F100_1F10A(t *testing.T) {
	ValidateRuneRange(t, 0x1F100, 0x1F10A, pDISALLOWED)
}

func Test1F10B_1F10F(t *testing.T) {
	ValidateRuneRange(t, 0x1F10B, 0x1F10F, pUNASSIGNED)
}

func Test1F110_1F12E(t *testing.T) {
	ValidateRuneRange(t, 0x1F110, 0x1F12E, pDISALLOWED)
}

func Test1F12F_1F130(t *testing.T) {
	ValidateRuneRange(t, 0x1F12F, 0x1F130, pUNASSIGNED)
}

func Test1F131(t *testing.T) {
	ValidateRune(t, 0x1F131, pDISALLOWED)
}

func Test1F132_1F13C(t *testing.T) {
	ValidateRuneRange(t, 0x1F132, 0x1F13C, pUNASSIGNED)
}

func Test1F13D(t *testing.T) {
	ValidateRune(t, 0x1F13D, pDISALLOWED)
}

func Test1F13E(t *testing.T) {
	ValidateRune(t, 0x1F13E, pUNASSIGNED)
}

func Test1F13F(t *testing.T) {
	ValidateRune(t, 0x1F13F, pDISALLOWED)
}

func Test1F140_1F141(t *testing.T) {
	ValidateRuneRange(t, 0x1F140, 0x1F141, pUNASSIGNED)
}

func Test1F142(t *testing.T) {
	ValidateRune(t, 0x1F142, pDISALLOWED)
}

func Test1F143_1F145(t *testing.T) {
	ValidateRuneRange(t, 0x1F143, 0x1F145, pUNASSIGNED)
}

func Test1F146(t *testing.T) {
	ValidateRune(t, 0x1F146, pDISALLOWED)
}

func Test1F147_1F149(t *testing.T) {
	ValidateRuneRange(t, 0x1F147, 0x1F149, pUNASSIGNED)
}

func Test1F14A_1F14E(t *testing.T) {
	ValidateRuneRange(t, 0x1F14A, 0x1F14E, pDISALLOWED)
}

func Test1F14F_1F156(t *testing.T) {
	ValidateRuneRange(t, 0x1F14F, 0x1F156, pUNASSIGNED)
}

func Test1F157(t *testing.T) {
	ValidateRune(t, 0x1F157, pDISALLOWED)
}

func Test1F158_1F15E(t *testing.T) {
	ValidateRuneRange(t, 0x1F158, 0x1F15E, pUNASSIGNED)
}

func Test1F15F(t *testing.T) {
	ValidateRune(t, 0x1F15F, pDISALLOWED)
}

func Test1F160_1F178(t *testing.T) {
	ValidateRuneRange(t, 0x1F160, 0x1F178, pUNASSIGNED)
}

func Test1F179(t *testing.T) {
	ValidateRune(t, 0x1F179, pDISALLOWED)
}

func Test1F17A(t *testing.T) {
	ValidateRune(t, 0x1F17A, pUNASSIGNED)
}

func Test1F17B_1F17C(t *testing.T) {
	ValidateRuneRange(t, 0x1F17B, 0x1F17C, pDISALLOWED)
}

func Test1F17D_1F17E(t *testing.T) {
	ValidateRuneRange(t, 0x1F17D, 0x1F17E, pUNASSIGNED)
}

func Test1F17F(t *testing.T) {
	ValidateRune(t, 0x1F17F, pDISALLOWED)
}

func Test1F180_1F189(t *testing.T) {
	ValidateRuneRange(t, 0x1F180, 0x1F189, pUNASSIGNED)
}

func Test1F18A_1F18D(t *testing.T) {
	ValidateRuneRange(t, 0x1F18A, 0x1F18D, pDISALLOWED)
}

func Test1F18E_1F18F(t *testing.T) {
	ValidateRuneRange(t, 0x1F18E, 0x1F18F, pUNASSIGNED)
}

func Test1F190(t *testing.T) {
	ValidateRune(t, 0x1F190, pDISALLOWED)
}

func Test1F191_1F1FF(t *testing.T) {
	ValidateRuneRange(t, 0x1F191, 0x1F1FF, pUNASSIGNED)
}

func Test1F200(t *testing.T) {
	ValidateRune(t, 0x1F200, pDISALLOWED)
}

func Test1F201_1F20F(t *testing.T) {
	ValidateRuneRange(t, 0x1F201, 0x1F20F, pUNASSIGNED)
}

func Test1F210_1F231(t *testing.T) {
	ValidateRuneRange(t, 0x1F210, 0x1F231, pDISALLOWED)
}

func Test1F232_1F23F(t *testing.T) {
	ValidateRuneRange(t, 0x1F232, 0x1F23F, pUNASSIGNED)
}

func Test1F240_1F248(t *testing.T) {
	ValidateRuneRange(t, 0x1F240, 0x1F248, pDISALLOWED)
}

func Test1F249_1FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x1F249, 0x1FFFD, pUNASSIGNED)
}

func Test1FFFE_1FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x1FFFE, 0x1FFFF, pDISALLOWED)
}

func Test20000_2A6D6(t *testing.T) {
	ValidateRuneRange(t, 0x20000, 0x2A6D6, pPVALID)
}

func Test2A6D7_2A6FF(t *testing.T) {
	ValidateRuneRange(t, 0x2A6D7, 0x2A6FF, pUNASSIGNED)
}

func Test2A700_2B734(t *testing.T) {
	ValidateRuneRange(t, 0x2A700, 0x2B734, pPVALID)
}

func Test2B735_2F7FF(t *testing.T) {
	ValidateRuneRange(t, 0x2B735, 0x2F7FF, pUNASSIGNED)
}

func Test2F800_2FA1D(t *testing.T) {
	ValidateRuneRange(t, 0x2F800, 0x2FA1D, pDISALLOWED)
}

func Test2FA1E_2FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x2FA1E, 0x2FFFD, pUNASSIGNED)
}

func Test2FFFE_2FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x2FFFE, 0x2FFFF, pDISALLOWED)
}

func Test30000_3FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x30000, 0x3FFFD, pUNASSIGNED)
}

func Test3FFFE_3FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x3FFFE, 0x3FFFF, pDISALLOWED)
}

func Test40000_4FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x40000, 0x4FFFD, pUNASSIGNED)
}

func Test4FFFE_4FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x4FFFE, 0x4FFFF, pDISALLOWED)
}

func Test50000_5FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x50000, 0x5FFFD, pUNASSIGNED)
}

func Test5FFFE_5FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x5FFFE, 0x5FFFF, pDISALLOWED)
}

func Test60000_6FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x60000, 0x6FFFD, pUNASSIGNED)
}

func Test6FFFE_6FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x6FFFE, 0x6FFFF, pDISALLOWED)
}

func Test70000_7FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x70000, 0x7FFFD, pUNASSIGNED)
}

func Test7FFFE_7FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x7FFFE, 0x7FFFF, pDISALLOWED)
}

func Test80000_8FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x80000, 0x8FFFD, pUNASSIGNED)
}

func Test8FFFE_8FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x8FFFE, 0x8FFFF, pDISALLOWED)
}

func Test90000_9FFFD(t *testing.T) {
	ValidateRuneRange(t, 0x90000, 0x9FFFD, pUNASSIGNED)
}

func Test9FFFE_9FFFF(t *testing.T) {
	ValidateRuneRange(t, 0x9FFFE, 0x9FFFF, pDISALLOWED)
}

func TestA0000_AFFFD(t *testing.T) {
	ValidateRuneRange(t, 0xA0000, 0xAFFFD, pUNASSIGNED)
}

func TestAFFFE_AFFFF(t *testing.T) {
	ValidateRuneRange(t, 0xAFFFE, 0xAFFFF, pDISALLOWED)
}

func TestB0000_BFFFD(t *testing.T) {
	ValidateRuneRange(t, 0xB0000, 0xBFFFD, pUNASSIGNED)
}

func TestBFFFE_BFFFF(t *testing.T) {
	ValidateRuneRange(t, 0xBFFFE, 0xBFFFF, pDISALLOWED)
}

func TestC0000_CFFFD(t *testing.T) {
	ValidateRuneRange(t, 0xC0000, 0xCFFFD, pUNASSIGNED)
}

func TestCFFFE_CFFFF(t *testing.T) {
	ValidateRuneRange(t, 0xCFFFE, 0xCFFFF, pDISALLOWED)
}

func TestD0000_DFFFD(t *testing.T) {
	ValidateRuneRange(t, 0xD0000, 0xDFFFD, pUNASSIGNED)
}

func TestDFFFE_DFFFF(t *testing.T) {
	ValidateRuneRange(t, 0xDFFFE, 0xDFFFF, pDISALLOWED)
}

func TestE0000(t *testing.T) {
	ValidateRune(t, 0xE0000, pUNASSIGNED)
}

func TestE0001(t *testing.T) {
	ValidateRune(t, 0xE0001, pDISALLOWED)
}

func TestE0002_E001F(t *testing.T) {
	ValidateRuneRange(t, 0xE0002, 0xE001F, pUNASSIGNED)
}

func TestE0020_E007F(t *testing.T) {
	ValidateRuneRange(t, 0xE0020, 0xE007F, pDISALLOWED)
}

func TestE0080_E00FF(t *testing.T) {
	ValidateRuneRange(t, 0xE0080, 0xE00FF, pUNASSIGNED)
}

func TestE0100_E01EF(t *testing.T) {
	ValidateRuneRange(t, 0xE0100, 0xE01EF, pDISALLOWED)
}

func TestE01F0_EFFFD(t *testing.T) {
	ValidateRuneRange(t, 0xE01F0, 0xEFFFD, pUNASSIGNED)
}

func TestEFFFE_10FFFF(t *testing.T) {
	ValidateRuneRange(t, 0xEFFFE, 0x10FFFF, pDISALLOWED)
}
