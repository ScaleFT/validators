package validators

/*

An implementation of unicode labal validation as described in RFC5892 and RFC6452
*/

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

const (
	// RFC5890_4_2LabelMaxLength defines the maximum length of a single IDNA label
	RFC5890_4_2LabelMaxLength = 252
)

// singleCP16 abbreviates the process of defining a single codepoint range
// c: the codepoint to make into a range
// returns: the constructed Range16
func singleCP16(c uint16) unicode.Range16 {
	return makeRange16(c, c)
}

// singleCP32  abbreviates the process of defining a single codepoint range
// c: the codepoint to make into a range
// returns: the constructed Range32
func singleCP32(c uint32) unicode.Range32 {
	return makeRange32(c, c)
}

func makeRange16(start uint16, end uint16) unicode.Range16 {
	return unicode.Range16{
		Lo:     start,
		Hi:     end,
		Stride: 1,
	}
}

func makeRange32(start uint32, end uint32) unicode.Range32 {
	return unicode.Range32{
		Lo:     start,
		Hi:     end,
		Stride: 1,
	}
}

type pVALUE int

const (
	pUNKNOWN    pVALUE = iota // unknown
	pPVALID                   // a valid codepoint
	pCONTEXTJ                 // requires join-context validation
	pCONTEXTO                 // requires other-context validation
	pDISALLOWED               // not allowd
	pUNASSIGNED               // not a valid character currently
)

var (
	exceptionpPVALID = unicode.RangeTable{
		R16: []unicode.Range16{
			singleCP16(0x00DF), // pPVALID     # LATIN SMALL LETTER SHARP S
			singleCP16(0x03C2), // pPVALID     # GREEK SMALL LETTER FINAL SIGMA
			singleCP16(0x06FD), // pPVALID     # ARABIC SIGN SINDHI AMPERSAND
			singleCP16(0x06FE), // pPVALID     # ARABIC SIGN SINDHI POSTPOSITION MEN
			singleCP16(0x0F0B), // pPVALID     # TIBETAN MARK INTERSYLLABIC TSHEG
			singleCP16(0x3007), // pPVALID     # IDEOGRAPHIC NUMBER ZERO
		},
	}

	exceptionpDISALLOWED = unicode.RangeTable{
		R16: []unicode.Range16{
			singleCP16(0x0640), // pDISALLOWED # ARABIC TATWEEL
			singleCP16(0x07FA), // pDISALLOWED # NKO LAJANYALAN
			singleCP16(0x302E), // pDISALLOWED # HANGUL SINGLE DOT TONE MARK
			singleCP16(0x302F), // pDISALLOWED # HANGUL DOUBLE DOT TONE MARK
			singleCP16(0x3031), // pDISALLOWED # VERTICAL KANA REPEAT MARK
			singleCP16(0x3032), // pDISALLOWED # VERTICAL KANA REPEAT WITH VOICED SOUND MARK
			singleCP16(0x3033), // pDISALLOWED # VERTICAL KANA REPEAT MARK UPPER HALF
			singleCP16(0x3034), // pDISALLOWED # VERTICAL KANA REPEAT WITH VOICED SOUND MARK UPPER HA
			singleCP16(0x3035), // pDISALLOWED # VERTICAL KANA REPEAT MARK LOWER HALF
			singleCP16(0x303B), // pDISALLOWED # VERTICAL IDEOGRAPHIC ITERATION MARKk
		},
	}

	exceptionpCONTEXTOPPVALID = unicode.RangeTable{
		R16: []unicode.Range16{
			singleCP16(0x00B7), // pCONTEXTO   # MIDDLE DOT
			singleCP16(0x0375), // pCONTEXTO   # GREEK LOWER NUMERAL SIGN (KERAIA)
			singleCP16(0x05F3), // pCONTEXTO   # HEBREW PUNCTUATION GERESH
			singleCP16(0x05F4), // pCONTEXTO   # HEBREW PUNCTUATION GERSHAYIM
			singleCP16(0x30FB), // pCONTEXTO   # KATAKANA MIDDLE DOT
		},
	}

	arabicIndicDigits = unicode.Range16{
		Lo:     0x0660,
		Hi:     0x0669,
		Stride: 1,
	}

	extendedArabicIndicDigits = unicode.Range16{
		Lo:     0x06F0,
		Hi:     0x06F9,
		Stride: 1,
	}

	exceptionpCONTEXTOPDISALLOWED = unicode.RangeTable{
		R16: []unicode.Range16{
			arabicIndicDigits,
			extendedArabicIndicDigits,
		},
	}

	// used to check for hangul/Jamo
	hangulSyllableType = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x1100, // L
				Hi:     0x11FF,
				Stride: 1,
			},
			{
				Lo:     0xA960, // L
				Hi:     0xA97C,
				Stride: 1,
			},
			{
				Lo:     0x1160, // V
				Hi:     0x11A7,
				Stride: 1,
			},
			{
				Lo:     0xD7B0, // V
				Hi:     0xD7C6,
				Stride: 1,
			},
			{
				Lo:     0x11A8, // T
				Hi:     0x11FF,
				Stride: 1,
			},
			{
				Lo:     0xD7CB,
				Hi:     0xD7FB,
				Stride: 1,
			},
		},
	}

	// test for isLDH (letter, digit, hyphen)
	ldhpoints = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x002D,
				Hi:     0x002D,
				Stride: 1,
			},
			{
				Lo:     0x0030,
				Hi:     0x0039,
				Stride: 1,
			},
			{
				Lo:     0x0061,
				Hi:     0x007A,
				Stride: 1,
			},
		},
	}

	assignedCodepoints []*unicode.RangeTable

	ignorableBlocksTable = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x20D0, // Combining Diacritical Marks for Symbols
				Hi:     0x20FF,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x1D100, // Musical Symbols
				Hi:     0x1D1FF,
				Stride: 1,
			},
			{
				Lo:     0x1D200, // Ancient Greek Musical Notation
				Hi:     0x1D24F,
				Stride: 1,
			},
		},
	}

	exceptionsForDefaultIgnorable = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0600,
				Hi:     0x0605,
				Stride: 1,
			},
			singleCP16(0x06DD),
			singleCP16(0x070F),
		},
		R32: []unicode.Range32{
			singleCP32(0x110BD),
		},
	}

	combiningCharactersAndMarks = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0300,
				Hi:     0x036F,
				Stride: 1,
			},
			{
				Lo:     0x0483,
				Hi:     0x0489,
				Stride: 1,
			},
			{
				Lo:     0x07EB,
				Hi:     0x07F3,
				Stride: 1,
			},
			{
				Lo:     0x135D,
				Hi:     0x135F,
				Stride: 1,
			},
			{
				Lo:     0x1AB0,
				Hi:     0x1ABE,
				Stride: 1,
			},
			{
				Lo:     0x1B6B,
				Hi:     0x1B73,
				Stride: 1,
			},
			{
				Lo:     0x1DC0,
				Hi:     0x1DF5,
				Stride: 1,
			},
			{
				Lo:     0x1DFC,
				Hi:     0x1DFF,
				Stride: 1,
			},
			{
				Lo:     0x20D0,
				Hi:     0x20F0,
				Stride: 1,
			},
			{
				Lo:     0x2CEF,
				Hi:     0x2CF1,
				Stride: 1,
			},
			{
				Lo:     0x2DE0,
				Hi:     0x2DFF,
				Stride: 1,
			},
			{
				Lo:     0x3099,
				Hi:     0x309A,
				Stride: 1,
			},
			{
				Lo:     0xA66F,
				Hi:     0xA672,
				Stride: 1,
			},
			{
				Lo:     0xA674,
				Hi:     0xA67D,
				Stride: 1,
			},
			{
				Lo:     0xA69E,
				Hi:     0xA69F,
				Stride: 1,
			},
			{
				Lo:     0xA6F0,
				Hi:     0xA6F1,
				Stride: 1,
			},
			{
				Lo:     0xA8E0,
				Hi:     0xA8F1,
				Stride: 1,
			},
			{
				Lo:     0xFE20,
				Hi:     0xFE2F,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{

			{
				Lo:     0x10376,
				Hi:     0x1037A,
				Stride: 1,
			},
			{
				Lo:     0x11366,
				Hi:     0x1136C,
				Stride: 1,
			},
			{
				Lo:     0x11370,
				Hi:     0x11374,
				Stride: 1,
			},
			{
				Lo:     0x16AF0,
				Hi:     0x16AF4,
				Stride: 1,
			},
			{
				Lo:     0x1D165,
				Hi:     0x1D169,
				Stride: 1,
			},
			{
				Lo:     0x1D16D,
				Hi:     0x1D172,
				Stride: 1,
			},
			{
				Lo:     0x1D17B,
				Hi:     0x1D182,
				Stride: 1,
			},
			{
				Lo:     0x1D185,
				Hi:     0x1D18B,
				Stride: 1,
			},
			{
				Lo:     0x1D1AA,
				Hi:     0x1D1AD,
				Stride: 1,
			},
			{
				Lo:     0x1D242,
				Hi:     0x1D244,
				Stride: 1,
			},
		},
	}
)

// Golang was not kind enough to include a Cn property, so I have to make
// something similar myself
func init() {
	assignedCodepoints = make([]*unicode.RangeTable, 0,
		len(unicode.Properties)+len(unicode.Categories))

	// FIXME: codepoints are duplicated, and because the implementation is O(N)
	// iteration, this is bad. An optimization would be to perform set-union
	// and compose a new range from that union (and pay the price only once)

	for _, v := range unicode.Properties {
		assignedCodepoints = append(assignedCodepoints, v)
	}
	for _, v := range unicode.Categories {
		assignedCodepoints = append(assignedCodepoints, v)
	}
}

// letterDigits (A) is from RFC5892
func letterDigits(c rune) bool {
	return unicode.In(c,
		unicode.Categories["Ll"],
		unicode.Categories["Lu"],
		unicode.Categories["Lo"],
		unicode.Categories["Nd"],
		unicode.Categories["Lm"],
		unicode.Categories["Mn"],
		unicode.Categories["Mc"])
}

// ToNFKC uses the NFKC normalization for individual runes
func ToNFKC(c rune) rune {
	s := []byte(string([]rune{c}))
	d := norm.NFKC.Bytes(s)
	runes := []rune(string(d))
	return runes[0]
}

// unstable (B) is from RFC5892
func unstable(cp rune) bool {
	if folded, ok := SimpleCaseFolding[ToNFKC(cp)]; ok {
		return ToNFKC(folded) != cp
	}
	return ToNFKC(cp) != cp
}

// ignorableProperties (C) is from RFC5892
func ignorableProperties(c rune) bool {
	return unicode.In(c,
		unicode.Properties["Other_Default_Ignorable_Code_Point"],
		unicode.Properties["White_Space"],
		unicode.Properties["Noncharacter_Code_Point"],
		unicode.Categories["Cf"],
		unicode.Properties["Variation_Selector"]) &&
		!unicode.In(c, &exceptionsForDefaultIgnorable)
}

// ignorableBlocks (D) is from RFC5892
func ignorableBlocks(c rune) bool {
	return unicode.In(c, &ignorableBlocksTable)
}

// isLDH (E) is from RFC5892
func isLDH(c rune) bool {
	return unicode.In(c, &ldhpoints)
}

// exceptions (F) is from RFC5892
func exceptions(c rune) (pVALUE, bool) {
	switch {
	case unicode.In(c, &exceptionpPVALID):
		return pPVALID, true
	case unicode.In(c, &exceptionpDISALLOWED):
		return pDISALLOWED, true
	case unicode.In(c, &exceptionpCONTEXTOPPVALID):
		return pCONTEXTO, true
	case unicode.In(c, &exceptionpCONTEXTOPDISALLOWED):
		return pCONTEXTO, true
	}
	return pUNKNOWN, false
}

// gackwardCompatible (G) is from RFC5892
func gackwardCompatible(c rune) (pVALUE, bool) {
	return pUNKNOWN, false
}

// joinControl (H) is from RFC5892
func joinControl(c rune) bool {
	return unicode.In(c, unicode.Properties["Join_Control"])
}

// oldHangulJamo (I) is from RFC5892
func oldHangulJamo(c rune) bool {
	if unicode.In(c, &hangulSyllableType) &&
		unicode.In(c, unicode.Categories["L"]) {
		return true
	}
	return false
}

// unassigned (J) is from RFC5892
func unassigned(c rune) bool {
	return !unicode.IsOneOf(assignedCodepoints, c)
}

// classifyRune is an implementation of the calculation of properties described if RFC5892 section 3
func classifyRune(cp rune) pVALUE {
	if v, ok := exceptions(cp); ok {
		return v
	}
	if v, ok := gackwardCompatible(cp); ok {
		return v
	}

	switch {
	case unassigned(cp):
		return pUNASSIGNED
	case isLDH(cp):
		return pPVALID
	case joinControl(cp):
		return pCONTEXTJ
	case unstable(cp):
		return pDISALLOWED
	case ignorableProperties(cp):
		return pDISALLOWED
	case ignorableBlocks(cp):
		return pDISALLOWED
	case oldHangulJamo(cp):
		return pDISALLOWED
	case letterDigits(cp):
		return pPVALID
	default:
		return pDISALLOWED
	}
}

// ContextualRules is an implementation of the contextual rules registery found in RFC5892 Appendix A.1-A.9
// val: a rune-slice of the label being validated
// idx: the index of the currently considered rune
// returns: true if pPVALID, false otherwise
func ContextualRules(val []rune, idx int) bool {
	var before, cp, after rune

	cp = val[idx]
	if idx > 0 {
		before = val[idx-1]
	}

	if idx < len(val) {
		after = val[idx+1]
	}

	switch cp {
	case 0x200C: // zero-width non-joiner
		return unicode.In(before, &viramaCombiningClass) ||
			ZeroWidthNonJoiner(val, idx)
	case 0x200D: // zero-width joiner
		return unicode.In(before, &viramaCombiningClass)
	case 0x00B7: // middle dot
		return before == 0x006C && after == 0x006C
	case 0x0375: // Greek lower numeral sign (keraia)
		return unicode.In(after, unicode.Scripts["Greek"])
	case 0x05F3, 0x05F4: // Hebrew puncuation geresh and Gershayim
		return unicode.In(before, unicode.Scripts["Hebrew"])
	case 0x30FB: // Katakana middle dot
		return unicode.In(cp, unicode.Scripts["Hiragana"],
			unicode.Scripts["Katakana"], unicode.Scripts["Han"])
	}

	switch {
	case unicode.In(cp, &unicode.RangeTable{R16: []unicode.Range16{arabicIndicDigits}}):
		// FIXME: This can be optimized ... a lot...
		for _, c := range val {
			if unicode.In(c, &unicode.RangeTable{R16: []unicode.Range16{extendedArabicIndicDigits}}) {
				return false
			}
		}
	case unicode.In(cp, &unicode.RangeTable{R16: []unicode.Range16{extendedArabicIndicDigits}}):
		for _, c := range val {
			if unicode.In(c, &unicode.RangeTable{R16: []unicode.Range16{arabicIndicDigits}}) {
				return false
			}
		}
	}

	return false
}

// ZeroWidthNonJoiner Helper for the schematic parsing for zero width non-joiners
func ZeroWidthNonJoiner(val []rune, idx int) (ok bool) {
	for i := idx; i > 0; i-- {
		jtype := GetJoinType(val[i])
		switch jtype {
		case JoinT:
			continue
		case JoinL, JoinD:
			ok = true
			break
		default:
			ok = false
			break
		}
	}

	if !ok {
		return ok
	}

	for i := idx; i < len(val); i++ {
		jtype := GetJoinType(val[i])
		switch jtype {
		case JoinT:
			continue
		case JoinR, JoinD:
			ok = true
			break
		default:
			ok = false
			break
		}
	}
	return ok
}

// verifyRFC5891_4_2_3_1 ensures there are no leading, terminating, or double-hypens in the label
func verifyRFC5891_4_2_3_1(label string) error {
	var p, q rune
	var i int
	runeslice := []rune(label)
	llen := len(runeslice) - 1

	if unicode.Is(unicode.Properties["Hyphen"], runeslice[0]) ||
		unicode.Is(unicode.Properties["Hyphen"], runeslice[llen]) {
		return errors.New("leading or trailing hyphen")
	}

	for i, p = range runeslice {
		if unicode.Is(unicode.Properties["Hyphen"], p) &&
			unicode.Is(unicode.Properties["Hyphen"], q) && i==3 {
			return errors.New("Consecutive hyphens")
		}
		q = p
	}
	return nil
}

// verifyRFC5891_4_2_3_2 ensures there aren't empty labels nor combining marks in a label
func verifyRFC5891_4_2_3_2(label string) error {
	runeslice := []rune(label)

	if len(runeslice) == 0 {
		return errors.New("label is empty")
	}

	if unicode.In(runeslice[0], &combiningCharactersAndMarks) {
		return errors.New("Cannot begin with combining mark")
	}

	return nil

}

type bidiType int

const (
	bidiLTR bidiType = iota // Left to right
	bidiRTL                 // Right to left
)

// verifyRFC5891_4_2_3_3 is actually mostly defined in rfc5893 #2 And it ensures that
// that a label begins with an LTR or RTL codepoint, and that that all
// characters within the label are compatible with it (so if LTR, everything is
// LTR, ad if RTL, everything is RTL). Finally it ensures that the label is
// terminated appropriately for the direction. (zero or more NSM cp's, preceded
// by one of the valid terminators for the direction)
func verifyRFC5891_4_2_3_3(label string) error {
	runeslice := []rune(label)
	var t bidiType
	var hasEN, hasAN bool

	if unicode.In(runeslice[0], &bidiClassL) {
		t = bidiLTR
	} else if unicode.In(runeslice[0], &bidiClassR, &bidiClassAL) {
		t = bidiRTL
	} else {
		return errors.New("Invalid Bidi class for first character of label")
	}

	for _, r := range runeslice {
		if t == bidiRTL {
			if !unicode.In(r, &bidiClassR, &bidiClassAL, &bidiClassAN,
				&bidiClassEN, &bidiClassES, &bidiClassCS, &bidiClassET, &bidiClassON,
				&bidiClassBN, &bidiClassNSM) {
				return errors.New("Invalid bidi for RTL label")
			}
			if unicode.In(r, &bidiClassEN) {
				hasEN = true
				if hasAN {
					return errors.New("label cannot have both AN and EN Bidi")
				}
			}
			if unicode.In(r, &bidiClassAN) {
				hasAN = true
				if hasEN {
					return errors.New("label cannot have both AN and EN Bidi")
				}
			}
		} else {
			if !unicode.In(r, &bidiClassL, &bidiClassEN, &bidiClassES,
				&bidiClassES, &bidiClassCS, &bidiClassET, &bidiClassON, &bidiClassBN,
				&bidiClassNSM) {
				return errors.New("Invalid bidi for LTR label")
			}
		}
	}

	for i := len(runeslice) - 1; i >= 0; i-- {
		if unicode.In(runeslice[i], &bidiClassNSM) {
			continue
		}
		if t == bidiRTL {
			if !unicode.In(runeslice[i], &bidiClassR, &bidiClassAL, &bidiClassEN,
				&bidiClassAN) {
				return errors.New("Invalid termination for RTL label")
			}
		} else {
			if !unicode.In(runeslice[i], &bidiClassL, &bidiClassEN) {
				return fmt.Errorf("Invalid termination for LTR label (%.4X)", runeslice[i])
			}
		}
	}
	return nil
}

// RFC5891DNSIgnoreCase ensures that a label is RFC5891 compliant, while
// ignoring "case", which is probably useful for validation, since most clients
// will kindly lowercase things for their users. The strict RFC compliant is
// used after lowercasing. Both are exposed in case using this turns out to
// have been a bad idea
func RFC5891DNSIgnoreCase(record string) error {
	normal := string(norm.NFC.Bytes([]byte(record)))
	normal = strings.ToLower(normal)
	return RFC5891DNS(normal)
}

// RFC5891DNS is the RFC5891 complaint validator for hostnames DNS records, and ensures
// compliance with IDNA
func RFC5891DNS(record string) error {
	normal := string(norm.NFC.Bytes([]byte(record)))
	var foundNull bool
	ldhOnly := true

	labels := strings.Split(normal, ".")

	for _, r := range normal {
		if !isLDH(r) && r != '.' {
			ldhOnly = false
			break
		}
	}

	for _, label := range labels {
		if foundNull {
			return fmt.Errorf("%s: Cannot contain non-terminating null label", record)
		}
		if len(label) == 0 {
			foundNull = true
			continue
		}
		if err := RFC5892(label); err != nil {
			return fmt.Errorf("%s: %s", record, err.Error())
		}
		if err := verifyRFC5891_4_2_3_1(label); err != nil {
			return fmt.Errorf("%s: %s", record, err.Error())
		}
		if !ldhOnly {
			if err := verifyRFC5891_4_2_3_3(label); err != nil {
				return fmt.Errorf("%s: %s", record, err.Error())
			}
		}
	}
	return nil
}

// RFC5892 ensures a label complies with IDNP for labels (less strict than the
// rules for IDNA DNS records, but used by it)
func RFC5892(label string) error {
	runeslice := []rune(label)
	if len(runeslice) > RFC5890_4_2LabelMaxLength {
		return errors.New("label too long")
	}

	for idx, cp := range label {
		switch classifyRune(cp) {
		case pPVALID:
			continue
		case pDISALLOWED:
			return fmt.Errorf("disallowed value '%c'", cp)
		case pCONTEXTO, pCONTEXTJ:
			if !ContextualRules(runeslice, idx) {
				return errors.New("disallowed based on context")
			}
		case pUNASSIGNED:
			return errors.New("unassigned")
		}
	}
	return nil
}
