package validators

/*
bidi_classes is generated with :

curl -o - http://www.unicode.org/Public/8.0.0/ucd/extracted/DerivedBidiClass.txt 2>/dev/null | grep -E '^[0-9A-F]{4,8}' | sed -e 's/;/ ;/'| sort -s -k3,3 | (echo -e "package validators\nimport \"unicode\"\n\nvar (\n"; gawk '{if($3 != last) { if(last) { print "\t\t},\n\t}\n"}; print "\tbidiClass"$3" = unicode.RangeTable{\n\t\tR16: []unicode.Range16{"} last=$3} /^[0-9A-F]{4}[[:space:]\.]/ {sixteen=1; } /^[0-9A-F]{5,8}[[:space:]\.]/ { if(sixteen) { sixteen=0; print "\t\t},\n\t\tR32: []unicode.Range32{"; } } {if(index($1,"..")) { n1=substr($1,0,index($1,"..")-1); n2=substr($1,index($1,"..")+2); } else { n1 = n2 = $1 } print "\t\t\t{\n\t\t\t\tLo: 0x"n1",\n\t\t\t\tHi: 0x"n2",\n\t\t\t\tStride: 1,\n\t\t\t}," }'; echo -e "\t\t},\n\t}\n)")  | gofmt >> THIS_FILE
*/

// --------------------- cut here --------------------------

import "unicode"

var (
	bidiClassAL = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0608,
				Hi:     0x0608,
				Stride: 1,
			},
			{
				Lo:     0x060B,
				Hi:     0x060B,
				Stride: 1,
			},
			{
				Lo:     0x060D,
				Hi:     0x060D,
				Stride: 1,
			},
			{
				Lo:     0x061B,
				Hi:     0x061B,
				Stride: 1,
			},
			{
				Lo:     0x061C,
				Hi:     0x061C,
				Stride: 1,
			},
			{
				Lo:     0x061D,
				Hi:     0x061D,
				Stride: 1,
			},
			{
				Lo:     0x061E,
				Hi:     0x061F,
				Stride: 1,
			},
			{
				Lo:     0x0620,
				Hi:     0x063F,
				Stride: 1,
			},
			{
				Lo:     0x0640,
				Hi:     0x0640,
				Stride: 1,
			},
			{
				Lo:     0x0641,
				Hi:     0x064A,
				Stride: 1,
			},
			{
				Lo:     0x066D,
				Hi:     0x066D,
				Stride: 1,
			},
			{
				Lo:     0x066E,
				Hi:     0x066F,
				Stride: 1,
			},
			{
				Lo:     0x0671,
				Hi:     0x06D3,
				Stride: 1,
			},
			{
				Lo:     0x06D4,
				Hi:     0x06D4,
				Stride: 1,
			},
			{
				Lo:     0x06D5,
				Hi:     0x06D5,
				Stride: 1,
			},
			{
				Lo:     0x06E5,
				Hi:     0x06E6,
				Stride: 1,
			},
			{
				Lo:     0x06EE,
				Hi:     0x06EF,
				Stride: 1,
			},
			{
				Lo:     0x06FA,
				Hi:     0x06FC,
				Stride: 1,
			},
			{
				Lo:     0x06FD,
				Hi:     0x06FE,
				Stride: 1,
			},
			{
				Lo:     0x06FF,
				Hi:     0x06FF,
				Stride: 1,
			},
			{
				Lo:     0x0700,
				Hi:     0x070D,
				Stride: 1,
			},
			{
				Lo:     0x070E,
				Hi:     0x070E,
				Stride: 1,
			},
			{
				Lo:     0x070F,
				Hi:     0x070F,
				Stride: 1,
			},
			{
				Lo:     0x0710,
				Hi:     0x0710,
				Stride: 1,
			},
			{
				Lo:     0x0712,
				Hi:     0x072F,
				Stride: 1,
			},
			{
				Lo:     0x074B,
				Hi:     0x074C,
				Stride: 1,
			},
			{
				Lo:     0x074D,
				Hi:     0x07A5,
				Stride: 1,
			},
			{
				Lo:     0x07B1,
				Hi:     0x07B1,
				Stride: 1,
			},
			{
				Lo:     0x07B2,
				Hi:     0x07BF,
				Stride: 1,
			},
			{
				Lo:     0x08A0,
				Hi:     0x08B4,
				Stride: 1,
			},
			{
				Lo:     0x08B5,
				Hi:     0x08E2,
				Stride: 1,
			},
			{
				Lo:     0xFB50,
				Hi:     0xFBB1,
				Stride: 1,
			},
			{
				Lo:     0xFBB2,
				Hi:     0xFBC1,
				Stride: 1,
			},
			{
				Lo:     0xFBC2,
				Hi:     0xFBD2,
				Stride: 1,
			},
			{
				Lo:     0xFBD3,
				Hi:     0xFD3D,
				Stride: 1,
			},
			{
				Lo:     0xFD40,
				Hi:     0xFD4F,
				Stride: 1,
			},
			{
				Lo:     0xFD50,
				Hi:     0xFD8F,
				Stride: 1,
			},
			{
				Lo:     0xFD90,
				Hi:     0xFD91,
				Stride: 1,
			},
			{
				Lo:     0xFD92,
				Hi:     0xFDC7,
				Stride: 1,
			},
			{
				Lo:     0xFDC8,
				Hi:     0xFDCF,
				Stride: 1,
			},
			{
				Lo:     0xFDF0,
				Hi:     0xFDFB,
				Stride: 1,
			},
			{
				Lo:     0xFDFC,
				Hi:     0xFDFC,
				Stride: 1,
			},
			{
				Lo:     0xFDFE,
				Hi:     0xFDFF,
				Stride: 1,
			},
			{
				Lo:     0xFE70,
				Hi:     0xFE74,
				Stride: 1,
			},
			{
				Lo:     0xFE75,
				Hi:     0xFE75,
				Stride: 1,
			},
			{
				Lo:     0xFE76,
				Hi:     0xFEFC,
				Stride: 1,
			},
			{
				Lo:     0xFEFD,
				Hi:     0xFEFE,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x1EE00,
				Hi:     0x1EE03,
				Stride: 1,
			},
			{
				Lo:     0x1EE04,
				Hi:     0x1EE04,
				Stride: 1,
			},
			{
				Lo:     0x1EE05,
				Hi:     0x1EE1F,
				Stride: 1,
			},
			{
				Lo:     0x1EE20,
				Hi:     0x1EE20,
				Stride: 1,
			},
			{
				Lo:     0x1EE21,
				Hi:     0x1EE22,
				Stride: 1,
			},
			{
				Lo:     0x1EE23,
				Hi:     0x1EE23,
				Stride: 1,
			},
			{
				Lo:     0x1EE24,
				Hi:     0x1EE24,
				Stride: 1,
			},
			{
				Lo:     0x1EE25,
				Hi:     0x1EE26,
				Stride: 1,
			},
			{
				Lo:     0x1EE27,
				Hi:     0x1EE27,
				Stride: 1,
			},
			{
				Lo:     0x1EE28,
				Hi:     0x1EE28,
				Stride: 1,
			},
			{
				Lo:     0x1EE29,
				Hi:     0x1EE32,
				Stride: 1,
			},
			{
				Lo:     0x1EE33,
				Hi:     0x1EE33,
				Stride: 1,
			},
			{
				Lo:     0x1EE34,
				Hi:     0x1EE37,
				Stride: 1,
			},
			{
				Lo:     0x1EE38,
				Hi:     0x1EE38,
				Stride: 1,
			},
			{
				Lo:     0x1EE39,
				Hi:     0x1EE39,
				Stride: 1,
			},
			{
				Lo:     0x1EE3A,
				Hi:     0x1EE3A,
				Stride: 1,
			},
			{
				Lo:     0x1EE3B,
				Hi:     0x1EE3B,
				Stride: 1,
			},
			{
				Lo:     0x1EE3C,
				Hi:     0x1EE41,
				Stride: 1,
			},
			{
				Lo:     0x1EE42,
				Hi:     0x1EE42,
				Stride: 1,
			},
			{
				Lo:     0x1EE43,
				Hi:     0x1EE46,
				Stride: 1,
			},
			{
				Lo:     0x1EE47,
				Hi:     0x1EE47,
				Stride: 1,
			},
			{
				Lo:     0x1EE48,
				Hi:     0x1EE48,
				Stride: 1,
			},
			{
				Lo:     0x1EE49,
				Hi:     0x1EE49,
				Stride: 1,
			},
			{
				Lo:     0x1EE4A,
				Hi:     0x1EE4A,
				Stride: 1,
			},
			{
				Lo:     0x1EE4B,
				Hi:     0x1EE4B,
				Stride: 1,
			},
			{
				Lo:     0x1EE4C,
				Hi:     0x1EE4C,
				Stride: 1,
			},
			{
				Lo:     0x1EE4D,
				Hi:     0x1EE4F,
				Stride: 1,
			},
			{
				Lo:     0x1EE50,
				Hi:     0x1EE50,
				Stride: 1,
			},
			{
				Lo:     0x1EE51,
				Hi:     0x1EE52,
				Stride: 1,
			},
			{
				Lo:     0x1EE53,
				Hi:     0x1EE53,
				Stride: 1,
			},
			{
				Lo:     0x1EE54,
				Hi:     0x1EE54,
				Stride: 1,
			},
			{
				Lo:     0x1EE55,
				Hi:     0x1EE56,
				Stride: 1,
			},
			{
				Lo:     0x1EE57,
				Hi:     0x1EE57,
				Stride: 1,
			},
			{
				Lo:     0x1EE58,
				Hi:     0x1EE58,
				Stride: 1,
			},
			{
				Lo:     0x1EE59,
				Hi:     0x1EE59,
				Stride: 1,
			},
			{
				Lo:     0x1EE5A,
				Hi:     0x1EE5A,
				Stride: 1,
			},
			{
				Lo:     0x1EE5B,
				Hi:     0x1EE5B,
				Stride: 1,
			},
			{
				Lo:     0x1EE5C,
				Hi:     0x1EE5C,
				Stride: 1,
			},
			{
				Lo:     0x1EE5D,
				Hi:     0x1EE5D,
				Stride: 1,
			},
			{
				Lo:     0x1EE5E,
				Hi:     0x1EE5E,
				Stride: 1,
			},
			{
				Lo:     0x1EE5F,
				Hi:     0x1EE5F,
				Stride: 1,
			},
			{
				Lo:     0x1EE60,
				Hi:     0x1EE60,
				Stride: 1,
			},
			{
				Lo:     0x1EE61,
				Hi:     0x1EE62,
				Stride: 1,
			},
			{
				Lo:     0x1EE63,
				Hi:     0x1EE63,
				Stride: 1,
			},
			{
				Lo:     0x1EE64,
				Hi:     0x1EE64,
				Stride: 1,
			},
			{
				Lo:     0x1EE65,
				Hi:     0x1EE66,
				Stride: 1,
			},
			{
				Lo:     0x1EE67,
				Hi:     0x1EE6A,
				Stride: 1,
			},
			{
				Lo:     0x1EE6B,
				Hi:     0x1EE6B,
				Stride: 1,
			},
			{
				Lo:     0x1EE6C,
				Hi:     0x1EE72,
				Stride: 1,
			},
			{
				Lo:     0x1EE73,
				Hi:     0x1EE73,
				Stride: 1,
			},
			{
				Lo:     0x1EE74,
				Hi:     0x1EE77,
				Stride: 1,
			},
			{
				Lo:     0x1EE78,
				Hi:     0x1EE78,
				Stride: 1,
			},
			{
				Lo:     0x1EE79,
				Hi:     0x1EE7C,
				Stride: 1,
			},
			{
				Lo:     0x1EE7D,
				Hi:     0x1EE7D,
				Stride: 1,
			},
			{
				Lo:     0x1EE7E,
				Hi:     0x1EE7E,
				Stride: 1,
			},
			{
				Lo:     0x1EE7F,
				Hi:     0x1EE7F,
				Stride: 1,
			},
			{
				Lo:     0x1EE80,
				Hi:     0x1EE89,
				Stride: 1,
			},
			{
				Lo:     0x1EE8A,
				Hi:     0x1EE8A,
				Stride: 1,
			},
			{
				Lo:     0x1EE8B,
				Hi:     0x1EE9B,
				Stride: 1,
			},
			{
				Lo:     0x1EE9C,
				Hi:     0x1EEA0,
				Stride: 1,
			},
			{
				Lo:     0x1EEA1,
				Hi:     0x1EEA3,
				Stride: 1,
			},
			{
				Lo:     0x1EEA4,
				Hi:     0x1EEA4,
				Stride: 1,
			},
			{
				Lo:     0x1EEA5,
				Hi:     0x1EEA9,
				Stride: 1,
			},
			{
				Lo:     0x1EEAA,
				Hi:     0x1EEAA,
				Stride: 1,
			},
			{
				Lo:     0x1EEAB,
				Hi:     0x1EEBB,
				Stride: 1,
			},
			{
				Lo:     0x1EEBC,
				Hi:     0x1EEEF,
				Stride: 1,
			},
			{
				Lo:     0x1EEF2,
				Hi:     0x1EEFF,
				Stride: 1,
			},
		},
	}

	bidiClassAN = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0600,
				Hi:     0x0605,
				Stride: 1,
			},
			{
				Lo:     0x0660,
				Hi:     0x0669,
				Stride: 1,
			},
			{
				Lo:     0x066B,
				Hi:     0x066C,
				Stride: 1,
			},
			{
				Lo:     0x06DD,
				Hi:     0x06DD,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x10E60,
				Hi:     0x10E7E,
				Stride: 1,
			},
		},
	}

	bidiClassB = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x000A,
				Hi:     0x000A,
				Stride: 1,
			},
			{
				Lo:     0x000D,
				Hi:     0x000D,
				Stride: 1,
			},
			{
				Lo:     0x001C,
				Hi:     0x001E,
				Stride: 1,
			},
			{
				Lo:     0x0085,
				Hi:     0x0085,
				Stride: 1,
			},
			{
				Lo:     0x2029,
				Hi:     0x2029,
				Stride: 1,
			},
		},
	}

	bidiClassBN = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0000,
				Hi:     0x0008,
				Stride: 1,
			},
			{
				Lo:     0x000E,
				Hi:     0x001B,
				Stride: 1,
			},
			{
				Lo:     0x007F,
				Hi:     0x0084,
				Stride: 1,
			},
			{
				Lo:     0x0086,
				Hi:     0x009F,
				Stride: 1,
			},
			{
				Lo:     0x00AD,
				Hi:     0x00AD,
				Stride: 1,
			},
			{
				Lo:     0x180E,
				Hi:     0x180E,
				Stride: 1,
			},
			{
				Lo:     0x200B,
				Hi:     0x200D,
				Stride: 1,
			},
			{
				Lo:     0x2060,
				Hi:     0x2064,
				Stride: 1,
			},
			{
				Lo:     0x2065,
				Hi:     0x2065,
				Stride: 1,
			},
			{
				Lo:     0x206A,
				Hi:     0x206F,
				Stride: 1,
			},
			{
				Lo:     0xFDD0,
				Hi:     0xFDEF,
				Stride: 1,
			},
			{
				Lo:     0xFEFF,
				Hi:     0xFEFF,
				Stride: 1,
			},
			{
				Lo:     0xFFF0,
				Hi:     0xFFF8,
				Stride: 1,
			},
			{
				Lo:     0xFFFE,
				Hi:     0xFFFF,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x1BCA0,
				Hi:     0x1BCA3,
				Stride: 1,
			},
			{
				Lo:     0x1D173,
				Hi:     0x1D17A,
				Stride: 1,
			},
			{
				Lo:     0x1FFFE,
				Hi:     0x1FFFF,
				Stride: 1,
			},
			{
				Lo:     0x2FFFE,
				Hi:     0x2FFFF,
				Stride: 1,
			},
			{
				Lo:     0x3FFFE,
				Hi:     0x3FFFF,
				Stride: 1,
			},
			{
				Lo:     0x4FFFE,
				Hi:     0x4FFFF,
				Stride: 1,
			},
			{
				Lo:     0x5FFFE,
				Hi:     0x5FFFF,
				Stride: 1,
			},
			{
				Lo:     0x6FFFE,
				Hi:     0x6FFFF,
				Stride: 1,
			},
			{
				Lo:     0x7FFFE,
				Hi:     0x7FFFF,
				Stride: 1,
			},
			{
				Lo:     0x8FFFE,
				Hi:     0x8FFFF,
				Stride: 1,
			},
			{
				Lo:     0x9FFFE,
				Hi:     0x9FFFF,
				Stride: 1,
			},
			{
				Lo:     0xAFFFE,
				Hi:     0xAFFFF,
				Stride: 1,
			},
			{
				Lo:     0xBFFFE,
				Hi:     0xBFFFF,
				Stride: 1,
			},
			{
				Lo:     0xCFFFE,
				Hi:     0xCFFFF,
				Stride: 1,
			},
			{
				Lo:     0xDFFFE,
				Hi:     0xE0000,
				Stride: 1,
			},
			{
				Lo:     0xE0001,
				Hi:     0xE0001,
				Stride: 1,
			},
			{
				Lo:     0xE0002,
				Hi:     0xE001F,
				Stride: 1,
			},
			{
				Lo:     0xE0020,
				Hi:     0xE007F,
				Stride: 1,
			},
			{
				Lo:     0xE0080,
				Hi:     0xE00FF,
				Stride: 1,
			},
			{
				Lo:     0xE01F0,
				Hi:     0xE0FFF,
				Stride: 1,
			},
			{
				Lo:     0xEFFFE,
				Hi:     0xEFFFF,
				Stride: 1,
			},
			{
				Lo:     0xFFFFE,
				Hi:     0xFFFFF,
				Stride: 1,
			},
			{
				Lo:     0x10FFFE,
				Hi:     0x10FFFF,
				Stride: 1,
			},
		},
	}

	bidiClassCS = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x002C,
				Hi:     0x002C,
				Stride: 1,
			},
			{
				Lo:     0x002E,
				Hi:     0x002F,
				Stride: 1,
			},
			{
				Lo:     0x003A,
				Hi:     0x003A,
				Stride: 1,
			},
			{
				Lo:     0x00A0,
				Hi:     0x00A0,
				Stride: 1,
			},
			{
				Lo:     0x060C,
				Hi:     0x060C,
				Stride: 1,
			},
			{
				Lo:     0x202F,
				Hi:     0x202F,
				Stride: 1,
			},
			{
				Lo:     0x2044,
				Hi:     0x2044,
				Stride: 1,
			},
			{
				Lo:     0xFE50,
				Hi:     0xFE50,
				Stride: 1,
			},
			{
				Lo:     0xFE52,
				Hi:     0xFE52,
				Stride: 1,
			},
			{
				Lo:     0xFE55,
				Hi:     0xFE55,
				Stride: 1,
			},
			{
				Lo:     0xFF0C,
				Hi:     0xFF0C,
				Stride: 1,
			},
			{
				Lo:     0xFF0E,
				Hi:     0xFF0F,
				Stride: 1,
			},
			{
				Lo:     0xFF1A,
				Hi:     0xFF1A,
				Stride: 1,
			},
		},
	}

	bidiClassEN = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0030,
				Hi:     0x0039,
				Stride: 1,
			},
			{
				Lo:     0x00B2,
				Hi:     0x00B3,
				Stride: 1,
			},
			{
				Lo:     0x00B9,
				Hi:     0x00B9,
				Stride: 1,
			},
			{
				Lo:     0x06F0,
				Hi:     0x06F9,
				Stride: 1,
			},
			{
				Lo:     0x2070,
				Hi:     0x2070,
				Stride: 1,
			},
			{
				Lo:     0x2074,
				Hi:     0x2079,
				Stride: 1,
			},
			{
				Lo:     0x2080,
				Hi:     0x2089,
				Stride: 1,
			},
			{
				Lo:     0x2488,
				Hi:     0x249B,
				Stride: 1,
			},
			{
				Lo:     0xFF10,
				Hi:     0xFF19,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x102E1,
				Hi:     0x102FB,
				Stride: 1,
			},
			{
				Lo:     0x1D7CE,
				Hi:     0x1D7FF,
				Stride: 1,
			},
			{
				Lo:     0x1F100,
				Hi:     0x1F10A,
				Stride: 1,
			},
		},
	}

	bidiClassES = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x002B,
				Hi:     0x002B,
				Stride: 1,
			},
			{
				Lo:     0x002D,
				Hi:     0x002D,
				Stride: 1,
			},
			{
				Lo:     0x207A,
				Hi:     0x207B,
				Stride: 1,
			},
			{
				Lo:     0x208A,
				Hi:     0x208B,
				Stride: 1,
			},
			{
				Lo:     0x2212,
				Hi:     0x2212,
				Stride: 1,
			},
			{
				Lo:     0xFB29,
				Hi:     0xFB29,
				Stride: 1,
			},
			{
				Lo:     0xFE62,
				Hi:     0xFE62,
				Stride: 1,
			},
			{
				Lo:     0xFE63,
				Hi:     0xFE63,
				Stride: 1,
			},
			{
				Lo:     0xFF0B,
				Hi:     0xFF0B,
				Stride: 1,
			},
			{
				Lo:     0xFF0D,
				Hi:     0xFF0D,
				Stride: 1,
			},
		},
	}

	bidiClassET = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0023,
				Hi:     0x0023,
				Stride: 1,
			},
			{
				Lo:     0x0024,
				Hi:     0x0024,
				Stride: 1,
			},
			{
				Lo:     0x0025,
				Hi:     0x0025,
				Stride: 1,
			},
			{
				Lo:     0x00A2,
				Hi:     0x00A5,
				Stride: 1,
			},
			{
				Lo:     0x00B0,
				Hi:     0x00B0,
				Stride: 1,
			},
			{
				Lo:     0x00B1,
				Hi:     0x00B1,
				Stride: 1,
			},
			{
				Lo:     0x058F,
				Hi:     0x058F,
				Stride: 1,
			},
			{
				Lo:     0x0609,
				Hi:     0x060A,
				Stride: 1,
			},
			{
				Lo:     0x066A,
				Hi:     0x066A,
				Stride: 1,
			},
			{
				Lo:     0x09F2,
				Hi:     0x09F3,
				Stride: 1,
			},
			{
				Lo:     0x09FB,
				Hi:     0x09FB,
				Stride: 1,
			},
			{
				Lo:     0x0AF1,
				Hi:     0x0AF1,
				Stride: 1,
			},
			{
				Lo:     0x0BF9,
				Hi:     0x0BF9,
				Stride: 1,
			},
			{
				Lo:     0x0E3F,
				Hi:     0x0E3F,
				Stride: 1,
			},
			{
				Lo:     0x17DB,
				Hi:     0x17DB,
				Stride: 1,
			},
			{
				Lo:     0x2030,
				Hi:     0x2034,
				Stride: 1,
			},
			{
				Lo:     0x20A0,
				Hi:     0x20BE,
				Stride: 1,
			},
			{
				Lo:     0x20BF,
				Hi:     0x20CF,
				Stride: 1,
			},
			{
				Lo:     0x212E,
				Hi:     0x212E,
				Stride: 1,
			},
			{
				Lo:     0x2213,
				Hi:     0x2213,
				Stride: 1,
			},
			{
				Lo:     0xA838,
				Hi:     0xA838,
				Stride: 1,
			},
			{
				Lo:     0xA839,
				Hi:     0xA839,
				Stride: 1,
			},
			{
				Lo:     0xFE5F,
				Hi:     0xFE5F,
				Stride: 1,
			},
			{
				Lo:     0xFE69,
				Hi:     0xFE69,
				Stride: 1,
			},
			{
				Lo:     0xFE6A,
				Hi:     0xFE6A,
				Stride: 1,
			},
			{
				Lo:     0xFF03,
				Hi:     0xFF03,
				Stride: 1,
			},
			{
				Lo:     0xFF04,
				Hi:     0xFF04,
				Stride: 1,
			},
			{
				Lo:     0xFF05,
				Hi:     0xFF05,
				Stride: 1,
			},
			{
				Lo:     0xFFE0,
				Hi:     0xFFE1,
				Stride: 1,
			},
			{
				Lo:     0xFFE5,
				Hi:     0xFFE6,
				Stride: 1,
			},
		},
	}

	bidiClassFSI = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x2068,
				Hi:     0x2068,
				Stride: 1,
			},
		},
	}

	bidiClassL = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0041,
				Hi:     0x005A,
				Stride: 1,
			},
			{
				Lo:     0x0061,
				Hi:     0x007A,
				Stride: 1,
			},
			{
				Lo:     0x00AA,
				Hi:     0x00AA,
				Stride: 1,
			},
			{
				Lo:     0x00B5,
				Hi:     0x00B5,
				Stride: 1,
			},
			{
				Lo:     0x00BA,
				Hi:     0x00BA,
				Stride: 1,
			},
			{
				Lo:     0x00C0,
				Hi:     0x00D6,
				Stride: 1,
			},
			{
				Lo:     0x00D8,
				Hi:     0x00F6,
				Stride: 1,
			},
			{
				Lo:     0x00F8,
				Hi:     0x01BA,
				Stride: 1,
			},
			{
				Lo:     0x01BB,
				Hi:     0x01BB,
				Stride: 1,
			},
			{
				Lo:     0x01BC,
				Hi:     0x01BF,
				Stride: 1,
			},
			{
				Lo:     0x01C0,
				Hi:     0x01C3,
				Stride: 1,
			},
			{
				Lo:     0x01C4,
				Hi:     0x0293,
				Stride: 1,
			},
			{
				Lo:     0x0294,
				Hi:     0x0294,
				Stride: 1,
			},
			{
				Lo:     0x0295,
				Hi:     0x02AF,
				Stride: 1,
			},
			{
				Lo:     0x02B0,
				Hi:     0x02B8,
				Stride: 1,
			},
			{
				Lo:     0x02BB,
				Hi:     0x02C1,
				Stride: 1,
			},
			{
				Lo:     0x02D0,
				Hi:     0x02D1,
				Stride: 1,
			},
			{
				Lo:     0x02E0,
				Hi:     0x02E4,
				Stride: 1,
			},
			{
				Lo:     0x02EE,
				Hi:     0x02EE,
				Stride: 1,
			},
			{
				Lo:     0x0370,
				Hi:     0x0373,
				Stride: 1,
			},
			{
				Lo:     0x0376,
				Hi:     0x0377,
				Stride: 1,
			},
			{
				Lo:     0x037A,
				Hi:     0x037A,
				Stride: 1,
			},
			{
				Lo:     0x037B,
				Hi:     0x037D,
				Stride: 1,
			},
			{
				Lo:     0x037F,
				Hi:     0x037F,
				Stride: 1,
			},
			{
				Lo:     0x0386,
				Hi:     0x0386,
				Stride: 1,
			},
			{
				Lo:     0x0388,
				Hi:     0x038A,
				Stride: 1,
			},
			{
				Lo:     0x038C,
				Hi:     0x038C,
				Stride: 1,
			},
			{
				Lo:     0x038E,
				Hi:     0x03A1,
				Stride: 1,
			},
			{
				Lo:     0x03A3,
				Hi:     0x03F5,
				Stride: 1,
			},
			{
				Lo:     0x03F7,
				Hi:     0x0481,
				Stride: 1,
			},
			{
				Lo:     0x0482,
				Hi:     0x0482,
				Stride: 1,
			},
			{
				Lo:     0x048A,
				Hi:     0x052F,
				Stride: 1,
			},
			{
				Lo:     0x0531,
				Hi:     0x0556,
				Stride: 1,
			},
			{
				Lo:     0x0559,
				Hi:     0x0559,
				Stride: 1,
			},
			{
				Lo:     0x055A,
				Hi:     0x055F,
				Stride: 1,
			},
			{
				Lo:     0x0561,
				Hi:     0x0587,
				Stride: 1,
			},
			{
				Lo:     0x0589,
				Hi:     0x0589,
				Stride: 1,
			},
			{
				Lo:     0x0903,
				Hi:     0x0903,
				Stride: 1,
			},
			{
				Lo:     0x0904,
				Hi:     0x0939,
				Stride: 1,
			},
			{
				Lo:     0x093B,
				Hi:     0x093B,
				Stride: 1,
			},
			{
				Lo:     0x093D,
				Hi:     0x093D,
				Stride: 1,
			},
			{
				Lo:     0x093E,
				Hi:     0x0940,
				Stride: 1,
			},
			{
				Lo:     0x0949,
				Hi:     0x094C,
				Stride: 1,
			},
			{
				Lo:     0x094E,
				Hi:     0x094F,
				Stride: 1,
			},
			{
				Lo:     0x0950,
				Hi:     0x0950,
				Stride: 1,
			},
			{
				Lo:     0x0958,
				Hi:     0x0961,
				Stride: 1,
			},
			{
				Lo:     0x0964,
				Hi:     0x0965,
				Stride: 1,
			},
			{
				Lo:     0x0966,
				Hi:     0x096F,
				Stride: 1,
			},
			{
				Lo:     0x0970,
				Hi:     0x0970,
				Stride: 1,
			},
			{
				Lo:     0x0971,
				Hi:     0x0971,
				Stride: 1,
			},
			{
				Lo:     0x0972,
				Hi:     0x0980,
				Stride: 1,
			},
			{
				Lo:     0x0982,
				Hi:     0x0983,
				Stride: 1,
			},
			{
				Lo:     0x0985,
				Hi:     0x098C,
				Stride: 1,
			},
			{
				Lo:     0x098F,
				Hi:     0x0990,
				Stride: 1,
			},
			{
				Lo:     0x0993,
				Hi:     0x09A8,
				Stride: 1,
			},
			{
				Lo:     0x09AA,
				Hi:     0x09B0,
				Stride: 1,
			},
			{
				Lo:     0x09B2,
				Hi:     0x09B2,
				Stride: 1,
			},
			{
				Lo:     0x09B6,
				Hi:     0x09B9,
				Stride: 1,
			},
			{
				Lo:     0x09BD,
				Hi:     0x09BD,
				Stride: 1,
			},
			{
				Lo:     0x09BE,
				Hi:     0x09C0,
				Stride: 1,
			},
			{
				Lo:     0x09C7,
				Hi:     0x09C8,
				Stride: 1,
			},
			{
				Lo:     0x09CB,
				Hi:     0x09CC,
				Stride: 1,
			},
			{
				Lo:     0x09CE,
				Hi:     0x09CE,
				Stride: 1,
			},
			{
				Lo:     0x09D7,
				Hi:     0x09D7,
				Stride: 1,
			},
			{
				Lo:     0x09DC,
				Hi:     0x09DD,
				Stride: 1,
			},
			{
				Lo:     0x09DF,
				Hi:     0x09E1,
				Stride: 1,
			},
			{
				Lo:     0x09E6,
				Hi:     0x09EF,
				Stride: 1,
			},
			{
				Lo:     0x09F0,
				Hi:     0x09F1,
				Stride: 1,
			},
			{
				Lo:     0x09F4,
				Hi:     0x09F9,
				Stride: 1,
			},
			{
				Lo:     0x09FA,
				Hi:     0x09FA,
				Stride: 1,
			},
			{
				Lo:     0x0A03,
				Hi:     0x0A03,
				Stride: 1,
			},
			{
				Lo:     0x0A05,
				Hi:     0x0A0A,
				Stride: 1,
			},
			{
				Lo:     0x0A0F,
				Hi:     0x0A10,
				Stride: 1,
			},
			{
				Lo:     0x0A13,
				Hi:     0x0A28,
				Stride: 1,
			},
			{
				Lo:     0x0A2A,
				Hi:     0x0A30,
				Stride: 1,
			},
			{
				Lo:     0x0A32,
				Hi:     0x0A33,
				Stride: 1,
			},
			{
				Lo:     0x0A35,
				Hi:     0x0A36,
				Stride: 1,
			},
			{
				Lo:     0x0A38,
				Hi:     0x0A39,
				Stride: 1,
			},
			{
				Lo:     0x0A3E,
				Hi:     0x0A40,
				Stride: 1,
			},
			{
				Lo:     0x0A59,
				Hi:     0x0A5C,
				Stride: 1,
			},
			{
				Lo:     0x0A5E,
				Hi:     0x0A5E,
				Stride: 1,
			},
			{
				Lo:     0x0A66,
				Hi:     0x0A6F,
				Stride: 1,
			},
			{
				Lo:     0x0A72,
				Hi:     0x0A74,
				Stride: 1,
			},
			{
				Lo:     0x0A83,
				Hi:     0x0A83,
				Stride: 1,
			},
			{
				Lo:     0x0A85,
				Hi:     0x0A8D,
				Stride: 1,
			},
			{
				Lo:     0x0A8F,
				Hi:     0x0A91,
				Stride: 1,
			},
			{
				Lo:     0x0A93,
				Hi:     0x0AA8,
				Stride: 1,
			},
			{
				Lo:     0x0AAA,
				Hi:     0x0AB0,
				Stride: 1,
			},
			{
				Lo:     0x0AB2,
				Hi:     0x0AB3,
				Stride: 1,
			},
			{
				Lo:     0x0AB5,
				Hi:     0x0AB9,
				Stride: 1,
			},
			{
				Lo:     0x0ABD,
				Hi:     0x0ABD,
				Stride: 1,
			},
			{
				Lo:     0x0ABE,
				Hi:     0x0AC0,
				Stride: 1,
			},
			{
				Lo:     0x0AC9,
				Hi:     0x0AC9,
				Stride: 1,
			},
			{
				Lo:     0x0ACB,
				Hi:     0x0ACC,
				Stride: 1,
			},
			{
				Lo:     0x0AD0,
				Hi:     0x0AD0,
				Stride: 1,
			},
			{
				Lo:     0x0AE0,
				Hi:     0x0AE1,
				Stride: 1,
			},
			{
				Lo:     0x0AE6,
				Hi:     0x0AEF,
				Stride: 1,
			},
			{
				Lo:     0x0AF0,
				Hi:     0x0AF0,
				Stride: 1,
			},
			{
				Lo:     0x0AF9,
				Hi:     0x0AF9,
				Stride: 1,
			},
			{
				Lo:     0x0B02,
				Hi:     0x0B03,
				Stride: 1,
			},
			{
				Lo:     0x0B05,
				Hi:     0x0B0C,
				Stride: 1,
			},
			{
				Lo:     0x0B0F,
				Hi:     0x0B10,
				Stride: 1,
			},
			{
				Lo:     0x0B13,
				Hi:     0x0B28,
				Stride: 1,
			},
			{
				Lo:     0x0B2A,
				Hi:     0x0B30,
				Stride: 1,
			},
			{
				Lo:     0x0B32,
				Hi:     0x0B33,
				Stride: 1,
			},
			{
				Lo:     0x0B35,
				Hi:     0x0B39,
				Stride: 1,
			},
			{
				Lo:     0x0B3D,
				Hi:     0x0B3D,
				Stride: 1,
			},
			{
				Lo:     0x0B3E,
				Hi:     0x0B3E,
				Stride: 1,
			},
			{
				Lo:     0x0B40,
				Hi:     0x0B40,
				Stride: 1,
			},
			{
				Lo:     0x0B47,
				Hi:     0x0B48,
				Stride: 1,
			},
			{
				Lo:     0x0B4B,
				Hi:     0x0B4C,
				Stride: 1,
			},
			{
				Lo:     0x0B57,
				Hi:     0x0B57,
				Stride: 1,
			},
			{
				Lo:     0x0B5C,
				Hi:     0x0B5D,
				Stride: 1,
			},
			{
				Lo:     0x0B5F,
				Hi:     0x0B61,
				Stride: 1,
			},
			{
				Lo:     0x0B66,
				Hi:     0x0B6F,
				Stride: 1,
			},
			{
				Lo:     0x0B70,
				Hi:     0x0B70,
				Stride: 1,
			},
			{
				Lo:     0x0B71,
				Hi:     0x0B71,
				Stride: 1,
			},
			{
				Lo:     0x0B72,
				Hi:     0x0B77,
				Stride: 1,
			},
			{
				Lo:     0x0B83,
				Hi:     0x0B83,
				Stride: 1,
			},
			{
				Lo:     0x0B85,
				Hi:     0x0B8A,
				Stride: 1,
			},
			{
				Lo:     0x0B8E,
				Hi:     0x0B90,
				Stride: 1,
			},
			{
				Lo:     0x0B92,
				Hi:     0x0B95,
				Stride: 1,
			},
			{
				Lo:     0x0B99,
				Hi:     0x0B9A,
				Stride: 1,
			},
			{
				Lo:     0x0B9C,
				Hi:     0x0B9C,
				Stride: 1,
			},
			{
				Lo:     0x0B9E,
				Hi:     0x0B9F,
				Stride: 1,
			},
			{
				Lo:     0x0BA3,
				Hi:     0x0BA4,
				Stride: 1,
			},
			{
				Lo:     0x0BA8,
				Hi:     0x0BAA,
				Stride: 1,
			},
			{
				Lo:     0x0BAE,
				Hi:     0x0BB9,
				Stride: 1,
			},
			{
				Lo:     0x0BBE,
				Hi:     0x0BBF,
				Stride: 1,
			},
			{
				Lo:     0x0BC1,
				Hi:     0x0BC2,
				Stride: 1,
			},
			{
				Lo:     0x0BC6,
				Hi:     0x0BC8,
				Stride: 1,
			},
			{
				Lo:     0x0BCA,
				Hi:     0x0BCC,
				Stride: 1,
			},
			{
				Lo:     0x0BD0,
				Hi:     0x0BD0,
				Stride: 1,
			},
			{
				Lo:     0x0BD7,
				Hi:     0x0BD7,
				Stride: 1,
			},
			{
				Lo:     0x0BE6,
				Hi:     0x0BEF,
				Stride: 1,
			},
			{
				Lo:     0x0BF0,
				Hi:     0x0BF2,
				Stride: 1,
			},
			{
				Lo:     0x0C01,
				Hi:     0x0C03,
				Stride: 1,
			},
			{
				Lo:     0x0C05,
				Hi:     0x0C0C,
				Stride: 1,
			},
			{
				Lo:     0x0C0E,
				Hi:     0x0C10,
				Stride: 1,
			},
			{
				Lo:     0x0C12,
				Hi:     0x0C28,
				Stride: 1,
			},
			{
				Lo:     0x0C2A,
				Hi:     0x0C39,
				Stride: 1,
			},
			{
				Lo:     0x0C3D,
				Hi:     0x0C3D,
				Stride: 1,
			},
			{
				Lo:     0x0C41,
				Hi:     0x0C44,
				Stride: 1,
			},
			{
				Lo:     0x0C58,
				Hi:     0x0C5A,
				Stride: 1,
			},
			{
				Lo:     0x0C60,
				Hi:     0x0C61,
				Stride: 1,
			},
			{
				Lo:     0x0C66,
				Hi:     0x0C6F,
				Stride: 1,
			},
			{
				Lo:     0x0C7F,
				Hi:     0x0C7F,
				Stride: 1,
			},
			{
				Lo:     0x0C82,
				Hi:     0x0C83,
				Stride: 1,
			},
			{
				Lo:     0x0C85,
				Hi:     0x0C8C,
				Stride: 1,
			},
			{
				Lo:     0x0C8E,
				Hi:     0x0C90,
				Stride: 1,
			},
			{
				Lo:     0x0C92,
				Hi:     0x0CA8,
				Stride: 1,
			},
			{
				Lo:     0x0CAA,
				Hi:     0x0CB3,
				Stride: 1,
			},
			{
				Lo:     0x0CB5,
				Hi:     0x0CB9,
				Stride: 1,
			},
			{
				Lo:     0x0CBD,
				Hi:     0x0CBD,
				Stride: 1,
			},
			{
				Lo:     0x0CBE,
				Hi:     0x0CBE,
				Stride: 1,
			},
			{
				Lo:     0x0CBF,
				Hi:     0x0CBF,
				Stride: 1,
			},
			{
				Lo:     0x0CC0,
				Hi:     0x0CC4,
				Stride: 1,
			},
			{
				Lo:     0x0CC6,
				Hi:     0x0CC6,
				Stride: 1,
			},
			{
				Lo:     0x0CC7,
				Hi:     0x0CC8,
				Stride: 1,
			},
			{
				Lo:     0x0CCA,
				Hi:     0x0CCB,
				Stride: 1,
			},
			{
				Lo:     0x0CD5,
				Hi:     0x0CD6,
				Stride: 1,
			},
			{
				Lo:     0x0CDE,
				Hi:     0x0CDE,
				Stride: 1,
			},
			{
				Lo:     0x0CE0,
				Hi:     0x0CE1,
				Stride: 1,
			},
			{
				Lo:     0x0CE6,
				Hi:     0x0CEF,
				Stride: 1,
			},
			{
				Lo:     0x0CF1,
				Hi:     0x0CF2,
				Stride: 1,
			},
			{
				Lo:     0x0D02,
				Hi:     0x0D03,
				Stride: 1,
			},
			{
				Lo:     0x0D05,
				Hi:     0x0D0C,
				Stride: 1,
			},
			{
				Lo:     0x0D0E,
				Hi:     0x0D10,
				Stride: 1,
			},
			{
				Lo:     0x0D12,
				Hi:     0x0D3A,
				Stride: 1,
			},
			{
				Lo:     0x0D3D,
				Hi:     0x0D3D,
				Stride: 1,
			},
			{
				Lo:     0x0D3E,
				Hi:     0x0D40,
				Stride: 1,
			},
			{
				Lo:     0x0D46,
				Hi:     0x0D48,
				Stride: 1,
			},
			{
				Lo:     0x0D4A,
				Hi:     0x0D4C,
				Stride: 1,
			},
			{
				Lo:     0x0D4E,
				Hi:     0x0D4E,
				Stride: 1,
			},
			{
				Lo:     0x0D57,
				Hi:     0x0D57,
				Stride: 1,
			},
			{
				Lo:     0x0D5F,
				Hi:     0x0D61,
				Stride: 1,
			},
			{
				Lo:     0x0D66,
				Hi:     0x0D6F,
				Stride: 1,
			},
			{
				Lo:     0x0D70,
				Hi:     0x0D75,
				Stride: 1,
			},
			{
				Lo:     0x0D79,
				Hi:     0x0D79,
				Stride: 1,
			},
			{
				Lo:     0x0D7A,
				Hi:     0x0D7F,
				Stride: 1,
			},
			{
				Lo:     0x0D82,
				Hi:     0x0D83,
				Stride: 1,
			},
			{
				Lo:     0x0D85,
				Hi:     0x0D96,
				Stride: 1,
			},
			{
				Lo:     0x0D9A,
				Hi:     0x0DB1,
				Stride: 1,
			},
			{
				Lo:     0x0DB3,
				Hi:     0x0DBB,
				Stride: 1,
			},
			{
				Lo:     0x0DBD,
				Hi:     0x0DBD,
				Stride: 1,
			},
			{
				Lo:     0x0DC0,
				Hi:     0x0DC6,
				Stride: 1,
			},
			{
				Lo:     0x0DCF,
				Hi:     0x0DD1,
				Stride: 1,
			},
			{
				Lo:     0x0DD8,
				Hi:     0x0DDF,
				Stride: 1,
			},
			{
				Lo:     0x0DE6,
				Hi:     0x0DEF,
				Stride: 1,
			},
			{
				Lo:     0x0DF2,
				Hi:     0x0DF3,
				Stride: 1,
			},
			{
				Lo:     0x0DF4,
				Hi:     0x0DF4,
				Stride: 1,
			},
			{
				Lo:     0x0E01,
				Hi:     0x0E30,
				Stride: 1,
			},
			{
				Lo:     0x0E32,
				Hi:     0x0E33,
				Stride: 1,
			},
			{
				Lo:     0x0E40,
				Hi:     0x0E45,
				Stride: 1,
			},
			{
				Lo:     0x0E46,
				Hi:     0x0E46,
				Stride: 1,
			},
			{
				Lo:     0x0E4F,
				Hi:     0x0E4F,
				Stride: 1,
			},
			{
				Lo:     0x0E50,
				Hi:     0x0E59,
				Stride: 1,
			},
			{
				Lo:     0x0E5A,
				Hi:     0x0E5B,
				Stride: 1,
			},
			{
				Lo:     0x0E81,
				Hi:     0x0E82,
				Stride: 1,
			},
			{
				Lo:     0x0E84,
				Hi:     0x0E84,
				Stride: 1,
			},
			{
				Lo:     0x0E87,
				Hi:     0x0E88,
				Stride: 1,
			},
			{
				Lo:     0x0E8A,
				Hi:     0x0E8A,
				Stride: 1,
			},
			{
				Lo:     0x0E8D,
				Hi:     0x0E8D,
				Stride: 1,
			},
			{
				Lo:     0x0E94,
				Hi:     0x0E97,
				Stride: 1,
			},
			{
				Lo:     0x0E99,
				Hi:     0x0E9F,
				Stride: 1,
			},
			{
				Lo:     0x0EA1,
				Hi:     0x0EA3,
				Stride: 1,
			},
			{
				Lo:     0x0EA5,
				Hi:     0x0EA5,
				Stride: 1,
			},
			{
				Lo:     0x0EA7,
				Hi:     0x0EA7,
				Stride: 1,
			},
			{
				Lo:     0x0EAA,
				Hi:     0x0EAB,
				Stride: 1,
			},
			{
				Lo:     0x0EAD,
				Hi:     0x0EB0,
				Stride: 1,
			},
			{
				Lo:     0x0EB2,
				Hi:     0x0EB3,
				Stride: 1,
			},
			{
				Lo:     0x0EBD,
				Hi:     0x0EBD,
				Stride: 1,
			},
			{
				Lo:     0x0EC0,
				Hi:     0x0EC4,
				Stride: 1,
			},
			{
				Lo:     0x0EC6,
				Hi:     0x0EC6,
				Stride: 1,
			},
			{
				Lo:     0x0ED0,
				Hi:     0x0ED9,
				Stride: 1,
			},
			{
				Lo:     0x0EDC,
				Hi:     0x0EDF,
				Stride: 1,
			},
			{
				Lo:     0x0F00,
				Hi:     0x0F00,
				Stride: 1,
			},
			{
				Lo:     0x0F01,
				Hi:     0x0F03,
				Stride: 1,
			},
			{
				Lo:     0x0F04,
				Hi:     0x0F12,
				Stride: 1,
			},
			{
				Lo:     0x0F13,
				Hi:     0x0F13,
				Stride: 1,
			},
			{
				Lo:     0x0F14,
				Hi:     0x0F14,
				Stride: 1,
			},
			{
				Lo:     0x0F15,
				Hi:     0x0F17,
				Stride: 1,
			},
			{
				Lo:     0x0F1A,
				Hi:     0x0F1F,
				Stride: 1,
			},
			{
				Lo:     0x0F20,
				Hi:     0x0F29,
				Stride: 1,
			},
			{
				Lo:     0x0F2A,
				Hi:     0x0F33,
				Stride: 1,
			},
			{
				Lo:     0x0F34,
				Hi:     0x0F34,
				Stride: 1,
			},
			{
				Lo:     0x0F36,
				Hi:     0x0F36,
				Stride: 1,
			},
			{
				Lo:     0x0F38,
				Hi:     0x0F38,
				Stride: 1,
			},
			{
				Lo:     0x0F3E,
				Hi:     0x0F3F,
				Stride: 1,
			},
			{
				Lo:     0x0F40,
				Hi:     0x0F47,
				Stride: 1,
			},
			{
				Lo:     0x0F49,
				Hi:     0x0F6C,
				Stride: 1,
			},
			{
				Lo:     0x0F7F,
				Hi:     0x0F7F,
				Stride: 1,
			},
			{
				Lo:     0x0F85,
				Hi:     0x0F85,
				Stride: 1,
			},
			{
				Lo:     0x0F88,
				Hi:     0x0F8C,
				Stride: 1,
			},
			{
				Lo:     0x0FBE,
				Hi:     0x0FC5,
				Stride: 1,
			},
			{
				Lo:     0x0FC7,
				Hi:     0x0FCC,
				Stride: 1,
			},
			{
				Lo:     0x0FCE,
				Hi:     0x0FCF,
				Stride: 1,
			},
			{
				Lo:     0x0FD0,
				Hi:     0x0FD4,
				Stride: 1,
			},
			{
				Lo:     0x0FD5,
				Hi:     0x0FD8,
				Stride: 1,
			},
			{
				Lo:     0x0FD9,
				Hi:     0x0FDA,
				Stride: 1,
			},
			{
				Lo:     0x1000,
				Hi:     0x102A,
				Stride: 1,
			},
			{
				Lo:     0x102B,
				Hi:     0x102C,
				Stride: 1,
			},
			{
				Lo:     0x1031,
				Hi:     0x1031,
				Stride: 1,
			},
			{
				Lo:     0x1038,
				Hi:     0x1038,
				Stride: 1,
			},
			{
				Lo:     0x103B,
				Hi:     0x103C,
				Stride: 1,
			},
			{
				Lo:     0x103F,
				Hi:     0x103F,
				Stride: 1,
			},
			{
				Lo:     0x1040,
				Hi:     0x1049,
				Stride: 1,
			},
			{
				Lo:     0x104A,
				Hi:     0x104F,
				Stride: 1,
			},
			{
				Lo:     0x1050,
				Hi:     0x1055,
				Stride: 1,
			},
			{
				Lo:     0x1056,
				Hi:     0x1057,
				Stride: 1,
			},
			{
				Lo:     0x105A,
				Hi:     0x105D,
				Stride: 1,
			},
			{
				Lo:     0x1061,
				Hi:     0x1061,
				Stride: 1,
			},
			{
				Lo:     0x1062,
				Hi:     0x1064,
				Stride: 1,
			},
			{
				Lo:     0x1065,
				Hi:     0x1066,
				Stride: 1,
			},
			{
				Lo:     0x1067,
				Hi:     0x106D,
				Stride: 1,
			},
			{
				Lo:     0x106E,
				Hi:     0x1070,
				Stride: 1,
			},
			{
				Lo:     0x1075,
				Hi:     0x1081,
				Stride: 1,
			},
			{
				Lo:     0x1083,
				Hi:     0x1084,
				Stride: 1,
			},
			{
				Lo:     0x1087,
				Hi:     0x108C,
				Stride: 1,
			},
			{
				Lo:     0x108E,
				Hi:     0x108E,
				Stride: 1,
			},
			{
				Lo:     0x108F,
				Hi:     0x108F,
				Stride: 1,
			},
			{
				Lo:     0x1090,
				Hi:     0x1099,
				Stride: 1,
			},
			{
				Lo:     0x109A,
				Hi:     0x109C,
				Stride: 1,
			},
			{
				Lo:     0x109E,
				Hi:     0x109F,
				Stride: 1,
			},
			{
				Lo:     0x10A0,
				Hi:     0x10C5,
				Stride: 1,
			},
			{
				Lo:     0x10C7,
				Hi:     0x10C7,
				Stride: 1,
			},
			{
				Lo:     0x10CD,
				Hi:     0x10CD,
				Stride: 1,
			},
			{
				Lo:     0x10D0,
				Hi:     0x10FA,
				Stride: 1,
			},
			{
				Lo:     0x10FB,
				Hi:     0x10FB,
				Stride: 1,
			},
			{
				Lo:     0x10FC,
				Hi:     0x10FC,
				Stride: 1,
			},
			{
				Lo:     0x10FD,
				Hi:     0x1248,
				Stride: 1,
			},
			{
				Lo:     0x124A,
				Hi:     0x124D,
				Stride: 1,
			},
			{
				Lo:     0x1250,
				Hi:     0x1256,
				Stride: 1,
			},
			{
				Lo:     0x1258,
				Hi:     0x1258,
				Stride: 1,
			},
			{
				Lo:     0x125A,
				Hi:     0x125D,
				Stride: 1,
			},
			{
				Lo:     0x1260,
				Hi:     0x1288,
				Stride: 1,
			},
			{
				Lo:     0x128A,
				Hi:     0x128D,
				Stride: 1,
			},
			{
				Lo:     0x1290,
				Hi:     0x12B0,
				Stride: 1,
			},
			{
				Lo:     0x12B2,
				Hi:     0x12B5,
				Stride: 1,
			},
			{
				Lo:     0x12B8,
				Hi:     0x12BE,
				Stride: 1,
			},
			{
				Lo:     0x12C0,
				Hi:     0x12C0,
				Stride: 1,
			},
			{
				Lo:     0x12C2,
				Hi:     0x12C5,
				Stride: 1,
			},
			{
				Lo:     0x12C8,
				Hi:     0x12D6,
				Stride: 1,
			},
			{
				Lo:     0x12D8,
				Hi:     0x1310,
				Stride: 1,
			},
			{
				Lo:     0x1312,
				Hi:     0x1315,
				Stride: 1,
			},
			{
				Lo:     0x1318,
				Hi:     0x135A,
				Stride: 1,
			},
			{
				Lo:     0x1360,
				Hi:     0x1368,
				Stride: 1,
			},
			{
				Lo:     0x1369,
				Hi:     0x137C,
				Stride: 1,
			},
			{
				Lo:     0x1380,
				Hi:     0x138F,
				Stride: 1,
			},
			{
				Lo:     0x13A0,
				Hi:     0x13F5,
				Stride: 1,
			},
			{
				Lo:     0x13F8,
				Hi:     0x13FD,
				Stride: 1,
			},
			{
				Lo:     0x1401,
				Hi:     0x166C,
				Stride: 1,
			},
			{
				Lo:     0x166D,
				Hi:     0x166E,
				Stride: 1,
			},
			{
				Lo:     0x166F,
				Hi:     0x167F,
				Stride: 1,
			},
			{
				Lo:     0x1681,
				Hi:     0x169A,
				Stride: 1,
			},
			{
				Lo:     0x16A0,
				Hi:     0x16EA,
				Stride: 1,
			},
			{
				Lo:     0x16EB,
				Hi:     0x16ED,
				Stride: 1,
			},
			{
				Lo:     0x16EE,
				Hi:     0x16F0,
				Stride: 1,
			},
			{
				Lo:     0x16F1,
				Hi:     0x16F8,
				Stride: 1,
			},
			{
				Lo:     0x1700,
				Hi:     0x170C,
				Stride: 1,
			},
			{
				Lo:     0x170E,
				Hi:     0x1711,
				Stride: 1,
			},
			{
				Lo:     0x1720,
				Hi:     0x1731,
				Stride: 1,
			},
			{
				Lo:     0x1735,
				Hi:     0x1736,
				Stride: 1,
			},
			{
				Lo:     0x1740,
				Hi:     0x1751,
				Stride: 1,
			},
			{
				Lo:     0x1760,
				Hi:     0x176C,
				Stride: 1,
			},
			{
				Lo:     0x176E,
				Hi:     0x1770,
				Stride: 1,
			},
			{
				Lo:     0x1780,
				Hi:     0x17B3,
				Stride: 1,
			},
			{
				Lo:     0x17B6,
				Hi:     0x17B6,
				Stride: 1,
			},
			{
				Lo:     0x17BE,
				Hi:     0x17C5,
				Stride: 1,
			},
			{
				Lo:     0x17C7,
				Hi:     0x17C8,
				Stride: 1,
			},
			{
				Lo:     0x17D4,
				Hi:     0x17D6,
				Stride: 1,
			},
			{
				Lo:     0x17D7,
				Hi:     0x17D7,
				Stride: 1,
			},
			{
				Lo:     0x17D8,
				Hi:     0x17DA,
				Stride: 1,
			},
			{
				Lo:     0x17DC,
				Hi:     0x17DC,
				Stride: 1,
			},
			{
				Lo:     0x17E0,
				Hi:     0x17E9,
				Stride: 1,
			},
			{
				Lo:     0x1810,
				Hi:     0x1819,
				Stride: 1,
			},
			{
				Lo:     0x1820,
				Hi:     0x1842,
				Stride: 1,
			},
			{
				Lo:     0x1843,
				Hi:     0x1843,
				Stride: 1,
			},
			{
				Lo:     0x1844,
				Hi:     0x1877,
				Stride: 1,
			},
			{
				Lo:     0x1880,
				Hi:     0x18A8,
				Stride: 1,
			},
			{
				Lo:     0x18AA,
				Hi:     0x18AA,
				Stride: 1,
			},
			{
				Lo:     0x18B0,
				Hi:     0x18F5,
				Stride: 1,
			},
			{
				Lo:     0x1900,
				Hi:     0x191E,
				Stride: 1,
			},
			{
				Lo:     0x1923,
				Hi:     0x1926,
				Stride: 1,
			},
			{
				Lo:     0x1929,
				Hi:     0x192B,
				Stride: 1,
			},
			{
				Lo:     0x1930,
				Hi:     0x1931,
				Stride: 1,
			},
			{
				Lo:     0x1933,
				Hi:     0x1938,
				Stride: 1,
			},
			{
				Lo:     0x1946,
				Hi:     0x194F,
				Stride: 1,
			},
			{
				Lo:     0x1950,
				Hi:     0x196D,
				Stride: 1,
			},
			{
				Lo:     0x1970,
				Hi:     0x1974,
				Stride: 1,
			},
			{
				Lo:     0x1980,
				Hi:     0x19AB,
				Stride: 1,
			},
			{
				Lo:     0x19B0,
				Hi:     0x19C9,
				Stride: 1,
			},
			{
				Lo:     0x19D0,
				Hi:     0x19D9,
				Stride: 1,
			},
			{
				Lo:     0x19DA,
				Hi:     0x19DA,
				Stride: 1,
			},
			{
				Lo:     0x1A00,
				Hi:     0x1A16,
				Stride: 1,
			},
			{
				Lo:     0x1A19,
				Hi:     0x1A1A,
				Stride: 1,
			},
			{
				Lo:     0x1A1E,
				Hi:     0x1A1F,
				Stride: 1,
			},
			{
				Lo:     0x1A20,
				Hi:     0x1A54,
				Stride: 1,
			},
			{
				Lo:     0x1A55,
				Hi:     0x1A55,
				Stride: 1,
			},
			{
				Lo:     0x1A57,
				Hi:     0x1A57,
				Stride: 1,
			},
			{
				Lo:     0x1A61,
				Hi:     0x1A61,
				Stride: 1,
			},
			{
				Lo:     0x1A63,
				Hi:     0x1A64,
				Stride: 1,
			},
			{
				Lo:     0x1A6D,
				Hi:     0x1A72,
				Stride: 1,
			},
			{
				Lo:     0x1A80,
				Hi:     0x1A89,
				Stride: 1,
			},
			{
				Lo:     0x1A90,
				Hi:     0x1A99,
				Stride: 1,
			},
			{
				Lo:     0x1AA0,
				Hi:     0x1AA6,
				Stride: 1,
			},
			{
				Lo:     0x1AA7,
				Hi:     0x1AA7,
				Stride: 1,
			},
			{
				Lo:     0x1AA8,
				Hi:     0x1AAD,
				Stride: 1,
			},
			{
				Lo:     0x1B04,
				Hi:     0x1B04,
				Stride: 1,
			},
			{
				Lo:     0x1B05,
				Hi:     0x1B33,
				Stride: 1,
			},
			{
				Lo:     0x1B35,
				Hi:     0x1B35,
				Stride: 1,
			},
			{
				Lo:     0x1B3B,
				Hi:     0x1B3B,
				Stride: 1,
			},
			{
				Lo:     0x1B3D,
				Hi:     0x1B41,
				Stride: 1,
			},
			{
				Lo:     0x1B43,
				Hi:     0x1B44,
				Stride: 1,
			},
			{
				Lo:     0x1B45,
				Hi:     0x1B4B,
				Stride: 1,
			},
			{
				Lo:     0x1B50,
				Hi:     0x1B59,
				Stride: 1,
			},
			{
				Lo:     0x1B5A,
				Hi:     0x1B60,
				Stride: 1,
			},
			{
				Lo:     0x1B61,
				Hi:     0x1B6A,
				Stride: 1,
			},
			{
				Lo:     0x1B74,
				Hi:     0x1B7C,
				Stride: 1,
			},
			{
				Lo:     0x1B82,
				Hi:     0x1B82,
				Stride: 1,
			},
			{
				Lo:     0x1B83,
				Hi:     0x1BA0,
				Stride: 1,
			},
			{
				Lo:     0x1BA1,
				Hi:     0x1BA1,
				Stride: 1,
			},
			{
				Lo:     0x1BA6,
				Hi:     0x1BA7,
				Stride: 1,
			},
			{
				Lo:     0x1BAA,
				Hi:     0x1BAA,
				Stride: 1,
			},
			{
				Lo:     0x1BAE,
				Hi:     0x1BAF,
				Stride: 1,
			},
			{
				Lo:     0x1BB0,
				Hi:     0x1BB9,
				Stride: 1,
			},
			{
				Lo:     0x1BBA,
				Hi:     0x1BE5,
				Stride: 1,
			},
			{
				Lo:     0x1BE7,
				Hi:     0x1BE7,
				Stride: 1,
			},
			{
				Lo:     0x1BEA,
				Hi:     0x1BEC,
				Stride: 1,
			},
			{
				Lo:     0x1BEE,
				Hi:     0x1BEE,
				Stride: 1,
			},
			{
				Lo:     0x1BF2,
				Hi:     0x1BF3,
				Stride: 1,
			},
			{
				Lo:     0x1BFC,
				Hi:     0x1BFF,
				Stride: 1,
			},
			{
				Lo:     0x1C00,
				Hi:     0x1C23,
				Stride: 1,
			},
			{
				Lo:     0x1C24,
				Hi:     0x1C2B,
				Stride: 1,
			},
			{
				Lo:     0x1C34,
				Hi:     0x1C35,
				Stride: 1,
			},
			{
				Lo:     0x1C3B,
				Hi:     0x1C3F,
				Stride: 1,
			},
			{
				Lo:     0x1C40,
				Hi:     0x1C49,
				Stride: 1,
			},
			{
				Lo:     0x1C4D,
				Hi:     0x1C4F,
				Stride: 1,
			},
			{
				Lo:     0x1C50,
				Hi:     0x1C59,
				Stride: 1,
			},
			{
				Lo:     0x1C5A,
				Hi:     0x1C77,
				Stride: 1,
			},
			{
				Lo:     0x1C78,
				Hi:     0x1C7D,
				Stride: 1,
			},
			{
				Lo:     0x1C7E,
				Hi:     0x1C7F,
				Stride: 1,
			},
			{
				Lo:     0x1CC0,
				Hi:     0x1CC7,
				Stride: 1,
			},
			{
				Lo:     0x1CD3,
				Hi:     0x1CD3,
				Stride: 1,
			},
			{
				Lo:     0x1CE1,
				Hi:     0x1CE1,
				Stride: 1,
			},
			{
				Lo:     0x1CE9,
				Hi:     0x1CEC,
				Stride: 1,
			},
			{
				Lo:     0x1CEE,
				Hi:     0x1CF1,
				Stride: 1,
			},
			{
				Lo:     0x1CF2,
				Hi:     0x1CF3,
				Stride: 1,
			},
			{
				Lo:     0x1CF5,
				Hi:     0x1CF6,
				Stride: 1,
			},
			{
				Lo:     0x1D00,
				Hi:     0x1D2B,
				Stride: 1,
			},
			{
				Lo:     0x1D2C,
				Hi:     0x1D6A,
				Stride: 1,
			},
			{
				Lo:     0x1D6B,
				Hi:     0x1D77,
				Stride: 1,
			},
			{
				Lo:     0x1D78,
				Hi:     0x1D78,
				Stride: 1,
			},
			{
				Lo:     0x1D79,
				Hi:     0x1D9A,
				Stride: 1,
			},
			{
				Lo:     0x1D9B,
				Hi:     0x1DBF,
				Stride: 1,
			},
			{
				Lo:     0x1E00,
				Hi:     0x1F15,
				Stride: 1,
			},
			{
				Lo:     0x1F18,
				Hi:     0x1F1D,
				Stride: 1,
			},
			{
				Lo:     0x1F20,
				Hi:     0x1F45,
				Stride: 1,
			},
			{
				Lo:     0x1F48,
				Hi:     0x1F4D,
				Stride: 1,
			},
			{
				Lo:     0x1F50,
				Hi:     0x1F57,
				Stride: 1,
			},
			{
				Lo:     0x1F59,
				Hi:     0x1F59,
				Stride: 1,
			},
			{
				Lo:     0x1F5B,
				Hi:     0x1F5B,
				Stride: 1,
			},
			{
				Lo:     0x1F5D,
				Hi:     0x1F5D,
				Stride: 1,
			},
			{
				Lo:     0x1F5F,
				Hi:     0x1F7D,
				Stride: 1,
			},
			{
				Lo:     0x1F80,
				Hi:     0x1FB4,
				Stride: 1,
			},
			{
				Lo:     0x1FB6,
				Hi:     0x1FBC,
				Stride: 1,
			},
			{
				Lo:     0x1FBE,
				Hi:     0x1FBE,
				Stride: 1,
			},
			{
				Lo:     0x1FC2,
				Hi:     0x1FC4,
				Stride: 1,
			},
			{
				Lo:     0x1FC6,
				Hi:     0x1FCC,
				Stride: 1,
			},
			{
				Lo:     0x1FD0,
				Hi:     0x1FD3,
				Stride: 1,
			},
			{
				Lo:     0x1FD6,
				Hi:     0x1FDB,
				Stride: 1,
			},
			{
				Lo:     0x1FE0,
				Hi:     0x1FEC,
				Stride: 1,
			},
			{
				Lo:     0x1FF2,
				Hi:     0x1FF4,
				Stride: 1,
			},
			{
				Lo:     0x1FF6,
				Hi:     0x1FFC,
				Stride: 1,
			},
			{
				Lo:     0x200E,
				Hi:     0x200E,
				Stride: 1,
			},
			{
				Lo:     0x2071,
				Hi:     0x2071,
				Stride: 1,
			},
			{
				Lo:     0x207F,
				Hi:     0x207F,
				Stride: 1,
			},
			{
				Lo:     0x2090,
				Hi:     0x209C,
				Stride: 1,
			},
			{
				Lo:     0x2102,
				Hi:     0x2102,
				Stride: 1,
			},
			{
				Lo:     0x2107,
				Hi:     0x2107,
				Stride: 1,
			},
			{
				Lo:     0x210A,
				Hi:     0x2113,
				Stride: 1,
			},
			{
				Lo:     0x2115,
				Hi:     0x2115,
				Stride: 1,
			},
			{
				Lo:     0x2119,
				Hi:     0x211D,
				Stride: 1,
			},
			{
				Lo:     0x2124,
				Hi:     0x2124,
				Stride: 1,
			},
			{
				Lo:     0x2126,
				Hi:     0x2126,
				Stride: 1,
			},
			{
				Lo:     0x2128,
				Hi:     0x2128,
				Stride: 1,
			},
			{
				Lo:     0x212A,
				Hi:     0x212D,
				Stride: 1,
			},
			{
				Lo:     0x212F,
				Hi:     0x2134,
				Stride: 1,
			},
			{
				Lo:     0x2135,
				Hi:     0x2138,
				Stride: 1,
			},
			{
				Lo:     0x2139,
				Hi:     0x2139,
				Stride: 1,
			},
			{
				Lo:     0x213C,
				Hi:     0x213F,
				Stride: 1,
			},
			{
				Lo:     0x2145,
				Hi:     0x2149,
				Stride: 1,
			},
			{
				Lo:     0x214E,
				Hi:     0x214E,
				Stride: 1,
			},
			{
				Lo:     0x214F,
				Hi:     0x214F,
				Stride: 1,
			},
			{
				Lo:     0x2160,
				Hi:     0x2182,
				Stride: 1,
			},
			{
				Lo:     0x2183,
				Hi:     0x2184,
				Stride: 1,
			},
			{
				Lo:     0x2185,
				Hi:     0x2188,
				Stride: 1,
			},
			{
				Lo:     0x2336,
				Hi:     0x237A,
				Stride: 1,
			},
			{
				Lo:     0x2395,
				Hi:     0x2395,
				Stride: 1,
			},
			{
				Lo:     0x249C,
				Hi:     0x24E9,
				Stride: 1,
			},
			{
				Lo:     0x26AC,
				Hi:     0x26AC,
				Stride: 1,
			},
			{
				Lo:     0x2800,
				Hi:     0x28FF,
				Stride: 1,
			},
			{
				Lo:     0x2C00,
				Hi:     0x2C2E,
				Stride: 1,
			},
			{
				Lo:     0x2C30,
				Hi:     0x2C5E,
				Stride: 1,
			},
			{
				Lo:     0x2C60,
				Hi:     0x2C7B,
				Stride: 1,
			},
			{
				Lo:     0x2C7C,
				Hi:     0x2C7D,
				Stride: 1,
			},
			{
				Lo:     0x2C7E,
				Hi:     0x2CE4,
				Stride: 1,
			},
			{
				Lo:     0x2CEB,
				Hi:     0x2CEE,
				Stride: 1,
			},
			{
				Lo:     0x2CF2,
				Hi:     0x2CF3,
				Stride: 1,
			},
			{
				Lo:     0x2D00,
				Hi:     0x2D25,
				Stride: 1,
			},
			{
				Lo:     0x2D27,
				Hi:     0x2D27,
				Stride: 1,
			},
			{
				Lo:     0x2D2D,
				Hi:     0x2D2D,
				Stride: 1,
			},
			{
				Lo:     0x2D30,
				Hi:     0x2D67,
				Stride: 1,
			},
			{
				Lo:     0x2D6F,
				Hi:     0x2D6F,
				Stride: 1,
			},
			{
				Lo:     0x2D70,
				Hi:     0x2D70,
				Stride: 1,
			},
			{
				Lo:     0x2D80,
				Hi:     0x2D96,
				Stride: 1,
			},
			{
				Lo:     0x2DA0,
				Hi:     0x2DA6,
				Stride: 1,
			},
			{
				Lo:     0x2DA8,
				Hi:     0x2DAE,
				Stride: 1,
			},
			{
				Lo:     0x2DB0,
				Hi:     0x2DB6,
				Stride: 1,
			},
			{
				Lo:     0x2DB8,
				Hi:     0x2DBE,
				Stride: 1,
			},
			{
				Lo:     0x2DC0,
				Hi:     0x2DC6,
				Stride: 1,
			},
			{
				Lo:     0x2DC8,
				Hi:     0x2DCE,
				Stride: 1,
			},
			{
				Lo:     0x2DD0,
				Hi:     0x2DD6,
				Stride: 1,
			},
			{
				Lo:     0x2DD8,
				Hi:     0x2DDE,
				Stride: 1,
			},
			{
				Lo:     0x3005,
				Hi:     0x3005,
				Stride: 1,
			},
			{
				Lo:     0x3006,
				Hi:     0x3006,
				Stride: 1,
			},
			{
				Lo:     0x3007,
				Hi:     0x3007,
				Stride: 1,
			},
			{
				Lo:     0x3021,
				Hi:     0x3029,
				Stride: 1,
			},
			{
				Lo:     0x302E,
				Hi:     0x302F,
				Stride: 1,
			},
			{
				Lo:     0x3031,
				Hi:     0x3035,
				Stride: 1,
			},
			{
				Lo:     0x3038,
				Hi:     0x303A,
				Stride: 1,
			},
			{
				Lo:     0x303B,
				Hi:     0x303B,
				Stride: 1,
			},
			{
				Lo:     0x303C,
				Hi:     0x303C,
				Stride: 1,
			},
			{
				Lo:     0x3041,
				Hi:     0x3096,
				Stride: 1,
			},
			{
				Lo:     0x309D,
				Hi:     0x309E,
				Stride: 1,
			},
			{
				Lo:     0x309F,
				Hi:     0x309F,
				Stride: 1,
			},
			{
				Lo:     0x30A1,
				Hi:     0x30FA,
				Stride: 1,
			},
			{
				Lo:     0x30FC,
				Hi:     0x30FE,
				Stride: 1,
			},
			{
				Lo:     0x30FF,
				Hi:     0x30FF,
				Stride: 1,
			},
			{
				Lo:     0x3105,
				Hi:     0x312D,
				Stride: 1,
			},
			{
				Lo:     0x3131,
				Hi:     0x318E,
				Stride: 1,
			},
			{
				Lo:     0x3190,
				Hi:     0x3191,
				Stride: 1,
			},
			{
				Lo:     0x3192,
				Hi:     0x3195,
				Stride: 1,
			},
			{
				Lo:     0x3196,
				Hi:     0x319F,
				Stride: 1,
			},
			{
				Lo:     0x31A0,
				Hi:     0x31BA,
				Stride: 1,
			},
			{
				Lo:     0x31F0,
				Hi:     0x31FF,
				Stride: 1,
			},
			{
				Lo:     0x3200,
				Hi:     0x321C,
				Stride: 1,
			},
			{
				Lo:     0x3220,
				Hi:     0x3229,
				Stride: 1,
			},
			{
				Lo:     0x322A,
				Hi:     0x3247,
				Stride: 1,
			},
			{
				Lo:     0x3248,
				Hi:     0x324F,
				Stride: 1,
			},
			{
				Lo:     0x3260,
				Hi:     0x327B,
				Stride: 1,
			},
			{
				Lo:     0x327F,
				Hi:     0x327F,
				Stride: 1,
			},
			{
				Lo:     0x3280,
				Hi:     0x3289,
				Stride: 1,
			},
			{
				Lo:     0x328A,
				Hi:     0x32B0,
				Stride: 1,
			},
			{
				Lo:     0x32C0,
				Hi:     0x32CB,
				Stride: 1,
			},
			{
				Lo:     0x32D0,
				Hi:     0x32FE,
				Stride: 1,
			},
			{
				Lo:     0x3300,
				Hi:     0x3376,
				Stride: 1,
			},
			{
				Lo:     0x337B,
				Hi:     0x33DD,
				Stride: 1,
			},
			{
				Lo:     0x33E0,
				Hi:     0x33FE,
				Stride: 1,
			},
			{
				Lo:     0x3400,
				Hi:     0x4DB5,
				Stride: 1,
			},
			{
				Lo:     0x4E00,
				Hi:     0x9FD5,
				Stride: 1,
			},
			{
				Lo:     0xA000,
				Hi:     0xA014,
				Stride: 1,
			},
			{
				Lo:     0xA015,
				Hi:     0xA015,
				Stride: 1,
			},
			{
				Lo:     0xA016,
				Hi:     0xA48C,
				Stride: 1,
			},
			{
				Lo:     0xA4D0,
				Hi:     0xA4F7,
				Stride: 1,
			},
			{
				Lo:     0xA4F8,
				Hi:     0xA4FD,
				Stride: 1,
			},
			{
				Lo:     0xA4FE,
				Hi:     0xA4FF,
				Stride: 1,
			},
			{
				Lo:     0xA500,
				Hi:     0xA60B,
				Stride: 1,
			},
			{
				Lo:     0xA60C,
				Hi:     0xA60C,
				Stride: 1,
			},
			{
				Lo:     0xA610,
				Hi:     0xA61F,
				Stride: 1,
			},
			{
				Lo:     0xA620,
				Hi:     0xA629,
				Stride: 1,
			},
			{
				Lo:     0xA62A,
				Hi:     0xA62B,
				Stride: 1,
			},
			{
				Lo:     0xA640,
				Hi:     0xA66D,
				Stride: 1,
			},
			{
				Lo:     0xA66E,
				Hi:     0xA66E,
				Stride: 1,
			},
			{
				Lo:     0xA680,
				Hi:     0xA69B,
				Stride: 1,
			},
			{
				Lo:     0xA69C,
				Hi:     0xA69D,
				Stride: 1,
			},
			{
				Lo:     0xA6A0,
				Hi:     0xA6E5,
				Stride: 1,
			},
			{
				Lo:     0xA6E6,
				Hi:     0xA6EF,
				Stride: 1,
			},
			{
				Lo:     0xA6F2,
				Hi:     0xA6F7,
				Stride: 1,
			},
			{
				Lo:     0xA722,
				Hi:     0xA76F,
				Stride: 1,
			},
			{
				Lo:     0xA770,
				Hi:     0xA770,
				Stride: 1,
			},
			{
				Lo:     0xA771,
				Hi:     0xA787,
				Stride: 1,
			},
			{
				Lo:     0xA789,
				Hi:     0xA78A,
				Stride: 1,
			},
			{
				Lo:     0xA78B,
				Hi:     0xA78E,
				Stride: 1,
			},
			{
				Lo:     0xA78F,
				Hi:     0xA78F,
				Stride: 1,
			},
			{
				Lo:     0xA790,
				Hi:     0xA7AD,
				Stride: 1,
			},
			{
				Lo:     0xA7B0,
				Hi:     0xA7B7,
				Stride: 1,
			},
			{
				Lo:     0xA7F7,
				Hi:     0xA7F7,
				Stride: 1,
			},
			{
				Lo:     0xA7F8,
				Hi:     0xA7F9,
				Stride: 1,
			},
			{
				Lo:     0xA7FA,
				Hi:     0xA7FA,
				Stride: 1,
			},
			{
				Lo:     0xA7FB,
				Hi:     0xA801,
				Stride: 1,
			},
			{
				Lo:     0xA803,
				Hi:     0xA805,
				Stride: 1,
			},
			{
				Lo:     0xA807,
				Hi:     0xA80A,
				Stride: 1,
			},
			{
				Lo:     0xA80C,
				Hi:     0xA822,
				Stride: 1,
			},
			{
				Lo:     0xA823,
				Hi:     0xA824,
				Stride: 1,
			},
			{
				Lo:     0xA827,
				Hi:     0xA827,
				Stride: 1,
			},
			{
				Lo:     0xA830,
				Hi:     0xA835,
				Stride: 1,
			},
			{
				Lo:     0xA836,
				Hi:     0xA837,
				Stride: 1,
			},
			{
				Lo:     0xA840,
				Hi:     0xA873,
				Stride: 1,
			},
			{
				Lo:     0xA880,
				Hi:     0xA881,
				Stride: 1,
			},
			{
				Lo:     0xA882,
				Hi:     0xA8B3,
				Stride: 1,
			},
			{
				Lo:     0xA8B4,
				Hi:     0xA8C3,
				Stride: 1,
			},
			{
				Lo:     0xA8CE,
				Hi:     0xA8CF,
				Stride: 1,
			},
			{
				Lo:     0xA8D0,
				Hi:     0xA8D9,
				Stride: 1,
			},
			{
				Lo:     0xA8F2,
				Hi:     0xA8F7,
				Stride: 1,
			},
			{
				Lo:     0xA8F8,
				Hi:     0xA8FA,
				Stride: 1,
			},
			{
				Lo:     0xA8FB,
				Hi:     0xA8FB,
				Stride: 1,
			},
			{
				Lo:     0xA8FC,
				Hi:     0xA8FC,
				Stride: 1,
			},
			{
				Lo:     0xA8FD,
				Hi:     0xA8FD,
				Stride: 1,
			},
			{
				Lo:     0xA900,
				Hi:     0xA909,
				Stride: 1,
			},
			{
				Lo:     0xA90A,
				Hi:     0xA925,
				Stride: 1,
			},
			{
				Lo:     0xA92E,
				Hi:     0xA92F,
				Stride: 1,
			},
			{
				Lo:     0xA930,
				Hi:     0xA946,
				Stride: 1,
			},
			{
				Lo:     0xA952,
				Hi:     0xA953,
				Stride: 1,
			},
			{
				Lo:     0xA95F,
				Hi:     0xA95F,
				Stride: 1,
			},
			{
				Lo:     0xA960,
				Hi:     0xA97C,
				Stride: 1,
			},
			{
				Lo:     0xA983,
				Hi:     0xA983,
				Stride: 1,
			},
			{
				Lo:     0xA984,
				Hi:     0xA9B2,
				Stride: 1,
			},
			{
				Lo:     0xA9B4,
				Hi:     0xA9B5,
				Stride: 1,
			},
			{
				Lo:     0xA9BA,
				Hi:     0xA9BB,
				Stride: 1,
			},
			{
				Lo:     0xA9BD,
				Hi:     0xA9C0,
				Stride: 1,
			},
			{
				Lo:     0xA9C1,
				Hi:     0xA9CD,
				Stride: 1,
			},
			{
				Lo:     0xA9CF,
				Hi:     0xA9CF,
				Stride: 1,
			},
			{
				Lo:     0xA9D0,
				Hi:     0xA9D9,
				Stride: 1,
			},
			{
				Lo:     0xA9DE,
				Hi:     0xA9DF,
				Stride: 1,
			},
			{
				Lo:     0xA9E0,
				Hi:     0xA9E4,
				Stride: 1,
			},
			{
				Lo:     0xA9E6,
				Hi:     0xA9E6,
				Stride: 1,
			},
			{
				Lo:     0xA9E7,
				Hi:     0xA9EF,
				Stride: 1,
			},
			{
				Lo:     0xA9F0,
				Hi:     0xA9F9,
				Stride: 1,
			},
			{
				Lo:     0xA9FA,
				Hi:     0xA9FE,
				Stride: 1,
			},
			{
				Lo:     0xAA00,
				Hi:     0xAA28,
				Stride: 1,
			},
			{
				Lo:     0xAA2F,
				Hi:     0xAA30,
				Stride: 1,
			},
			{
				Lo:     0xAA33,
				Hi:     0xAA34,
				Stride: 1,
			},
			{
				Lo:     0xAA40,
				Hi:     0xAA42,
				Stride: 1,
			},
			{
				Lo:     0xAA44,
				Hi:     0xAA4B,
				Stride: 1,
			},
			{
				Lo:     0xAA4D,
				Hi:     0xAA4D,
				Stride: 1,
			},
			{
				Lo:     0xAA50,
				Hi:     0xAA59,
				Stride: 1,
			},
			{
				Lo:     0xAA5C,
				Hi:     0xAA5F,
				Stride: 1,
			},
			{
				Lo:     0xAA60,
				Hi:     0xAA6F,
				Stride: 1,
			},
			{
				Lo:     0xAA70,
				Hi:     0xAA70,
				Stride: 1,
			},
			{
				Lo:     0xAA71,
				Hi:     0xAA76,
				Stride: 1,
			},
			{
				Lo:     0xAA77,
				Hi:     0xAA79,
				Stride: 1,
			},
			{
				Lo:     0xAA7A,
				Hi:     0xAA7A,
				Stride: 1,
			},
			{
				Lo:     0xAA7B,
				Hi:     0xAA7B,
				Stride: 1,
			},
			{
				Lo:     0xAA7D,
				Hi:     0xAA7D,
				Stride: 1,
			},
			{
				Lo:     0xAA7E,
				Hi:     0xAAAF,
				Stride: 1,
			},
			{
				Lo:     0xAAB1,
				Hi:     0xAAB1,
				Stride: 1,
			},
			{
				Lo:     0xAAB5,
				Hi:     0xAAB6,
				Stride: 1,
			},
			{
				Lo:     0xAAB9,
				Hi:     0xAABD,
				Stride: 1,
			},
			{
				Lo:     0xAAC0,
				Hi:     0xAAC0,
				Stride: 1,
			},
			{
				Lo:     0xAAC2,
				Hi:     0xAAC2,
				Stride: 1,
			},
			{
				Lo:     0xAADB,
				Hi:     0xAADC,
				Stride: 1,
			},
			{
				Lo:     0xAADD,
				Hi:     0xAADD,
				Stride: 1,
			},
			{
				Lo:     0xAADE,
				Hi:     0xAADF,
				Stride: 1,
			},
			{
				Lo:     0xAAE0,
				Hi:     0xAAEA,
				Stride: 1,
			},
			{
				Lo:     0xAAEB,
				Hi:     0xAAEB,
				Stride: 1,
			},
			{
				Lo:     0xAAEE,
				Hi:     0xAAEF,
				Stride: 1,
			},
			{
				Lo:     0xAAF0,
				Hi:     0xAAF1,
				Stride: 1,
			},
			{
				Lo:     0xAAF2,
				Hi:     0xAAF2,
				Stride: 1,
			},
			{
				Lo:     0xAAF3,
				Hi:     0xAAF4,
				Stride: 1,
			},
			{
				Lo:     0xAAF5,
				Hi:     0xAAF5,
				Stride: 1,
			},
			{
				Lo:     0xAB01,
				Hi:     0xAB06,
				Stride: 1,
			},
			{
				Lo:     0xAB09,
				Hi:     0xAB0E,
				Stride: 1,
			},
			{
				Lo:     0xAB11,
				Hi:     0xAB16,
				Stride: 1,
			},
			{
				Lo:     0xAB20,
				Hi:     0xAB26,
				Stride: 1,
			},
			{
				Lo:     0xAB28,
				Hi:     0xAB2E,
				Stride: 1,
			},
			{
				Lo:     0xAB30,
				Hi:     0xAB5A,
				Stride: 1,
			},
			{
				Lo:     0xAB5B,
				Hi:     0xAB5B,
				Stride: 1,
			},
			{
				Lo:     0xAB5C,
				Hi:     0xAB5F,
				Stride: 1,
			},
			{
				Lo:     0xAB60,
				Hi:     0xAB65,
				Stride: 1,
			},
			{
				Lo:     0xAB70,
				Hi:     0xABBF,
				Stride: 1,
			},
			{
				Lo:     0xABC0,
				Hi:     0xABE2,
				Stride: 1,
			},
			{
				Lo:     0xABE3,
				Hi:     0xABE4,
				Stride: 1,
			},
			{
				Lo:     0xABE6,
				Hi:     0xABE7,
				Stride: 1,
			},
			{
				Lo:     0xABE9,
				Hi:     0xABEA,
				Stride: 1,
			},
			{
				Lo:     0xABEB,
				Hi:     0xABEB,
				Stride: 1,
			},
			{
				Lo:     0xABEC,
				Hi:     0xABEC,
				Stride: 1,
			},
			{
				Lo:     0xABF0,
				Hi:     0xABF9,
				Stride: 1,
			},
			{
				Lo:     0xAC00,
				Hi:     0xD7A3,
				Stride: 1,
			},
			{
				Lo:     0xD7B0,
				Hi:     0xD7C6,
				Stride: 1,
			},
			{
				Lo:     0xD7CB,
				Hi:     0xD7FB,
				Stride: 1,
			},
			{
				Lo:     0xE000,
				Hi:     0xF8FF,
				Stride: 1,
			},
			{
				Lo:     0xF900,
				Hi:     0xFA6D,
				Stride: 1,
			},
			{
				Lo:     0xFA70,
				Hi:     0xFAD9,
				Stride: 1,
			},
			{
				Lo:     0xFB00,
				Hi:     0xFB06,
				Stride: 1,
			},
			{
				Lo:     0xFB13,
				Hi:     0xFB17,
				Stride: 1,
			},
			{
				Lo:     0xFF21,
				Hi:     0xFF3A,
				Stride: 1,
			},
			{
				Lo:     0xFF41,
				Hi:     0xFF5A,
				Stride: 1,
			},
			{
				Lo:     0xFF66,
				Hi:     0xFF6F,
				Stride: 1,
			},
			{
				Lo:     0xFF70,
				Hi:     0xFF70,
				Stride: 1,
			},
			{
				Lo:     0xFF71,
				Hi:     0xFF9D,
				Stride: 1,
			},
			{
				Lo:     0xFF9E,
				Hi:     0xFF9F,
				Stride: 1,
			},
			{
				Lo:     0xFFA0,
				Hi:     0xFFBE,
				Stride: 1,
			},
			{
				Lo:     0xFFC2,
				Hi:     0xFFC7,
				Stride: 1,
			},
			{
				Lo:     0xFFCA,
				Hi:     0xFFCF,
				Stride: 1,
			},
			{
				Lo:     0xFFD2,
				Hi:     0xFFD7,
				Stride: 1,
			},
			{
				Lo:     0xFFDA,
				Hi:     0xFFDC,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x10000,
				Hi:     0x1000B,
				Stride: 1,
			},
			{
				Lo:     0x1000D,
				Hi:     0x10026,
				Stride: 1,
			},
			{
				Lo:     0x10028,
				Hi:     0x1003A,
				Stride: 1,
			},
			{
				Lo:     0x1003C,
				Hi:     0x1003D,
				Stride: 1,
			},
			{
				Lo:     0x1003F,
				Hi:     0x1004D,
				Stride: 1,
			},
			{
				Lo:     0x10050,
				Hi:     0x1005D,
				Stride: 1,
			},
			{
				Lo:     0x10080,
				Hi:     0x100FA,
				Stride: 1,
			},
			{
				Lo:     0x10100,
				Hi:     0x10100,
				Stride: 1,
			},
			{
				Lo:     0x10102,
				Hi:     0x10102,
				Stride: 1,
			},
			{
				Lo:     0x10107,
				Hi:     0x10133,
				Stride: 1,
			},
			{
				Lo:     0x10137,
				Hi:     0x1013F,
				Stride: 1,
			},
			{
				Lo:     0x101D0,
				Hi:     0x101FC,
				Stride: 1,
			},
			{
				Lo:     0x10280,
				Hi:     0x1029C,
				Stride: 1,
			},
			{
				Lo:     0x102A0,
				Hi:     0x102D0,
				Stride: 1,
			},
			{
				Lo:     0x10300,
				Hi:     0x1031F,
				Stride: 1,
			},
			{
				Lo:     0x10320,
				Hi:     0x10323,
				Stride: 1,
			},
			{
				Lo:     0x10330,
				Hi:     0x10340,
				Stride: 1,
			},
			{
				Lo:     0x10341,
				Hi:     0x10341,
				Stride: 1,
			},
			{
				Lo:     0x10342,
				Hi:     0x10349,
				Stride: 1,
			},
			{
				Lo:     0x1034A,
				Hi:     0x1034A,
				Stride: 1,
			},
			{
				Lo:     0x10350,
				Hi:     0x10375,
				Stride: 1,
			},
			{
				Lo:     0x10380,
				Hi:     0x1039D,
				Stride: 1,
			},
			{
				Lo:     0x1039F,
				Hi:     0x1039F,
				Stride: 1,
			},
			{
				Lo:     0x103A0,
				Hi:     0x103C3,
				Stride: 1,
			},
			{
				Lo:     0x103C8,
				Hi:     0x103CF,
				Stride: 1,
			},
			{
				Lo:     0x103D0,
				Hi:     0x103D0,
				Stride: 1,
			},
			{
				Lo:     0x103D1,
				Hi:     0x103D5,
				Stride: 1,
			},
			{
				Lo:     0x10400,
				Hi:     0x1044F,
				Stride: 1,
			},
			{
				Lo:     0x10450,
				Hi:     0x1049D,
				Stride: 1,
			},
			{
				Lo:     0x104A0,
				Hi:     0x104A9,
				Stride: 1,
			},
			{
				Lo:     0x10500,
				Hi:     0x10527,
				Stride: 1,
			},
			{
				Lo:     0x10530,
				Hi:     0x10563,
				Stride: 1,
			},
			{
				Lo:     0x1056F,
				Hi:     0x1056F,
				Stride: 1,
			},
			{
				Lo:     0x10600,
				Hi:     0x10736,
				Stride: 1,
			},
			{
				Lo:     0x10740,
				Hi:     0x10755,
				Stride: 1,
			},
			{
				Lo:     0x10760,
				Hi:     0x10767,
				Stride: 1,
			},
			{
				Lo:     0x11000,
				Hi:     0x11000,
				Stride: 1,
			},
			{
				Lo:     0x11002,
				Hi:     0x11002,
				Stride: 1,
			},
			{
				Lo:     0x11003,
				Hi:     0x11037,
				Stride: 1,
			},
			{
				Lo:     0x11047,
				Hi:     0x1104D,
				Stride: 1,
			},
			{
				Lo:     0x11066,
				Hi:     0x1106F,
				Stride: 1,
			},
			{
				Lo:     0x11082,
				Hi:     0x11082,
				Stride: 1,
			},
			{
				Lo:     0x11083,
				Hi:     0x110AF,
				Stride: 1,
			},
			{
				Lo:     0x110B0,
				Hi:     0x110B2,
				Stride: 1,
			},
			{
				Lo:     0x110B7,
				Hi:     0x110B8,
				Stride: 1,
			},
			{
				Lo:     0x110BB,
				Hi:     0x110BC,
				Stride: 1,
			},
			{
				Lo:     0x110BD,
				Hi:     0x110BD,
				Stride: 1,
			},
			{
				Lo:     0x110BE,
				Hi:     0x110C1,
				Stride: 1,
			},
			{
				Lo:     0x110D0,
				Hi:     0x110E8,
				Stride: 1,
			},
			{
				Lo:     0x110F0,
				Hi:     0x110F9,
				Stride: 1,
			},
			{
				Lo:     0x11103,
				Hi:     0x11126,
				Stride: 1,
			},
			{
				Lo:     0x1112C,
				Hi:     0x1112C,
				Stride: 1,
			},
			{
				Lo:     0x11136,
				Hi:     0x1113F,
				Stride: 1,
			},
			{
				Lo:     0x11140,
				Hi:     0x11143,
				Stride: 1,
			},
			{
				Lo:     0x11150,
				Hi:     0x11172,
				Stride: 1,
			},
			{
				Lo:     0x11174,
				Hi:     0x11175,
				Stride: 1,
			},
			{
				Lo:     0x11176,
				Hi:     0x11176,
				Stride: 1,
			},
			{
				Lo:     0x11182,
				Hi:     0x11182,
				Stride: 1,
			},
			{
				Lo:     0x11183,
				Hi:     0x111B2,
				Stride: 1,
			},
			{
				Lo:     0x111B3,
				Hi:     0x111B5,
				Stride: 1,
			},
			{
				Lo:     0x111BF,
				Hi:     0x111C0,
				Stride: 1,
			},
			{
				Lo:     0x111C1,
				Hi:     0x111C4,
				Stride: 1,
			},
			{
				Lo:     0x111C5,
				Hi:     0x111C9,
				Stride: 1,
			},
			{
				Lo:     0x111CD,
				Hi:     0x111CD,
				Stride: 1,
			},
			{
				Lo:     0x111D0,
				Hi:     0x111D9,
				Stride: 1,
			},
			{
				Lo:     0x111DA,
				Hi:     0x111DA,
				Stride: 1,
			},
			{
				Lo:     0x111DB,
				Hi:     0x111DB,
				Stride: 1,
			},
			{
				Lo:     0x111DC,
				Hi:     0x111DC,
				Stride: 1,
			},
			{
				Lo:     0x111DD,
				Hi:     0x111DF,
				Stride: 1,
			},
			{
				Lo:     0x111E1,
				Hi:     0x111F4,
				Stride: 1,
			},
			{
				Lo:     0x11200,
				Hi:     0x11211,
				Stride: 1,
			},
			{
				Lo:     0x11213,
				Hi:     0x1122B,
				Stride: 1,
			},
			{
				Lo:     0x1122C,
				Hi:     0x1122E,
				Stride: 1,
			},
			{
				Lo:     0x11232,
				Hi:     0x11233,
				Stride: 1,
			},
			{
				Lo:     0x11235,
				Hi:     0x11235,
				Stride: 1,
			},
			{
				Lo:     0x11238,
				Hi:     0x1123D,
				Stride: 1,
			},
			{
				Lo:     0x11280,
				Hi:     0x11286,
				Stride: 1,
			},
			{
				Lo:     0x11288,
				Hi:     0x11288,
				Stride: 1,
			},
			{
				Lo:     0x1128A,
				Hi:     0x1128D,
				Stride: 1,
			},
			{
				Lo:     0x1128F,
				Hi:     0x1129D,
				Stride: 1,
			},
			{
				Lo:     0x1129F,
				Hi:     0x112A8,
				Stride: 1,
			},
			{
				Lo:     0x112A9,
				Hi:     0x112A9,
				Stride: 1,
			},
			{
				Lo:     0x112B0,
				Hi:     0x112DE,
				Stride: 1,
			},
			{
				Lo:     0x112E0,
				Hi:     0x112E2,
				Stride: 1,
			},
			{
				Lo:     0x112F0,
				Hi:     0x112F9,
				Stride: 1,
			},
			{
				Lo:     0x11302,
				Hi:     0x11303,
				Stride: 1,
			},
			{
				Lo:     0x11305,
				Hi:     0x1130C,
				Stride: 1,
			},
			{
				Lo:     0x1130F,
				Hi:     0x11310,
				Stride: 1,
			},
			{
				Lo:     0x11313,
				Hi:     0x11328,
				Stride: 1,
			},
			{
				Lo:     0x1132A,
				Hi:     0x11330,
				Stride: 1,
			},
			{
				Lo:     0x11332,
				Hi:     0x11333,
				Stride: 1,
			},
			{
				Lo:     0x11335,
				Hi:     0x11339,
				Stride: 1,
			},
			{
				Lo:     0x1133D,
				Hi:     0x1133D,
				Stride: 1,
			},
			{
				Lo:     0x1133E,
				Hi:     0x1133F,
				Stride: 1,
			},
			{
				Lo:     0x11341,
				Hi:     0x11344,
				Stride: 1,
			},
			{
				Lo:     0x11347,
				Hi:     0x11348,
				Stride: 1,
			},
			{
				Lo:     0x1134B,
				Hi:     0x1134D,
				Stride: 1,
			},
			{
				Lo:     0x11350,
				Hi:     0x11350,
				Stride: 1,
			},
			{
				Lo:     0x11357,
				Hi:     0x11357,
				Stride: 1,
			},
			{
				Lo:     0x1135D,
				Hi:     0x11361,
				Stride: 1,
			},
			{
				Lo:     0x11362,
				Hi:     0x11363,
				Stride: 1,
			},
			{
				Lo:     0x11480,
				Hi:     0x114AF,
				Stride: 1,
			},
			{
				Lo:     0x114B0,
				Hi:     0x114B2,
				Stride: 1,
			},
			{
				Lo:     0x114B9,
				Hi:     0x114B9,
				Stride: 1,
			},
			{
				Lo:     0x114BB,
				Hi:     0x114BE,
				Stride: 1,
			},
			{
				Lo:     0x114C1,
				Hi:     0x114C1,
				Stride: 1,
			},
			{
				Lo:     0x114C4,
				Hi:     0x114C5,
				Stride: 1,
			},
			{
				Lo:     0x114C6,
				Hi:     0x114C6,
				Stride: 1,
			},
			{
				Lo:     0x114C7,
				Hi:     0x114C7,
				Stride: 1,
			},
			{
				Lo:     0x114D0,
				Hi:     0x114D9,
				Stride: 1,
			},
			{
				Lo:     0x11580,
				Hi:     0x115AE,
				Stride: 1,
			},
			{
				Lo:     0x115AF,
				Hi:     0x115B1,
				Stride: 1,
			},
			{
				Lo:     0x115B8,
				Hi:     0x115BB,
				Stride: 1,
			},
			{
				Lo:     0x115BE,
				Hi:     0x115BE,
				Stride: 1,
			},
			{
				Lo:     0x115C1,
				Hi:     0x115D7,
				Stride: 1,
			},
			{
				Lo:     0x115D8,
				Hi:     0x115DB,
				Stride: 1,
			},
			{
				Lo:     0x11600,
				Hi:     0x1162F,
				Stride: 1,
			},
			{
				Lo:     0x11630,
				Hi:     0x11632,
				Stride: 1,
			},
			{
				Lo:     0x1163B,
				Hi:     0x1163C,
				Stride: 1,
			},
			{
				Lo:     0x1163E,
				Hi:     0x1163E,
				Stride: 1,
			},
			{
				Lo:     0x11641,
				Hi:     0x11643,
				Stride: 1,
			},
			{
				Lo:     0x11644,
				Hi:     0x11644,
				Stride: 1,
			},
			{
				Lo:     0x11650,
				Hi:     0x11659,
				Stride: 1,
			},
			{
				Lo:     0x11680,
				Hi:     0x116AA,
				Stride: 1,
			},
			{
				Lo:     0x116AC,
				Hi:     0x116AC,
				Stride: 1,
			},
			{
				Lo:     0x116AE,
				Hi:     0x116AF,
				Stride: 1,
			},
			{
				Lo:     0x116B6,
				Hi:     0x116B6,
				Stride: 1,
			},
			{
				Lo:     0x116C0,
				Hi:     0x116C9,
				Stride: 1,
			},
			{
				Lo:     0x11700,
				Hi:     0x11719,
				Stride: 1,
			},
			{
				Lo:     0x11720,
				Hi:     0x11721,
				Stride: 1,
			},
			{
				Lo:     0x11726,
				Hi:     0x11726,
				Stride: 1,
			},
			{
				Lo:     0x11730,
				Hi:     0x11739,
				Stride: 1,
			},
			{
				Lo:     0x1173A,
				Hi:     0x1173B,
				Stride: 1,
			},
			{
				Lo:     0x1173C,
				Hi:     0x1173E,
				Stride: 1,
			},
			{
				Lo:     0x1173F,
				Hi:     0x1173F,
				Stride: 1,
			},
			{
				Lo:     0x118A0,
				Hi:     0x118DF,
				Stride: 1,
			},
			{
				Lo:     0x118E0,
				Hi:     0x118E9,
				Stride: 1,
			},
			{
				Lo:     0x118EA,
				Hi:     0x118F2,
				Stride: 1,
			},
			{
				Lo:     0x118FF,
				Hi:     0x118FF,
				Stride: 1,
			},
			{
				Lo:     0x11AC0,
				Hi:     0x11AF8,
				Stride: 1,
			},
			{
				Lo:     0x12000,
				Hi:     0x12399,
				Stride: 1,
			},
			{
				Lo:     0x12400,
				Hi:     0x1246E,
				Stride: 1,
			},
			{
				Lo:     0x12470,
				Hi:     0x12474,
				Stride: 1,
			},
			{
				Lo:     0x12480,
				Hi:     0x12543,
				Stride: 1,
			},
			{
				Lo:     0x13000,
				Hi:     0x1342E,
				Stride: 1,
			},
			{
				Lo:     0x14400,
				Hi:     0x14646,
				Stride: 1,
			},
			{
				Lo:     0x16800,
				Hi:     0x16A38,
				Stride: 1,
			},
			{
				Lo:     0x16A40,
				Hi:     0x16A5E,
				Stride: 1,
			},
			{
				Lo:     0x16A60,
				Hi:     0x16A69,
				Stride: 1,
			},
			{
				Lo:     0x16A6E,
				Hi:     0x16A6F,
				Stride: 1,
			},
			{
				Lo:     0x16AD0,
				Hi:     0x16AED,
				Stride: 1,
			},
			{
				Lo:     0x16AF5,
				Hi:     0x16AF5,
				Stride: 1,
			},
			{
				Lo:     0x16B00,
				Hi:     0x16B2F,
				Stride: 1,
			},
			{
				Lo:     0x16B37,
				Hi:     0x16B3B,
				Stride: 1,
			},
			{
				Lo:     0x16B3C,
				Hi:     0x16B3F,
				Stride: 1,
			},
			{
				Lo:     0x16B40,
				Hi:     0x16B43,
				Stride: 1,
			},
			{
				Lo:     0x16B44,
				Hi:     0x16B44,
				Stride: 1,
			},
			{
				Lo:     0x16B45,
				Hi:     0x16B45,
				Stride: 1,
			},
			{
				Lo:     0x16B50,
				Hi:     0x16B59,
				Stride: 1,
			},
			{
				Lo:     0x16B5B,
				Hi:     0x16B61,
				Stride: 1,
			},
			{
				Lo:     0x16B63,
				Hi:     0x16B77,
				Stride: 1,
			},
			{
				Lo:     0x16B7D,
				Hi:     0x16B8F,
				Stride: 1,
			},
			{
				Lo:     0x16F00,
				Hi:     0x16F44,
				Stride: 1,
			},
			{
				Lo:     0x16F50,
				Hi:     0x16F50,
				Stride: 1,
			},
			{
				Lo:     0x16F51,
				Hi:     0x16F7E,
				Stride: 1,
			},
			{
				Lo:     0x16F93,
				Hi:     0x16F9F,
				Stride: 1,
			},
			{
				Lo:     0x1B000,
				Hi:     0x1B001,
				Stride: 1,
			},
			{
				Lo:     0x1BC00,
				Hi:     0x1BC6A,
				Stride: 1,
			},
			{
				Lo:     0x1BC70,
				Hi:     0x1BC7C,
				Stride: 1,
			},
			{
				Lo:     0x1BC80,
				Hi:     0x1BC88,
				Stride: 1,
			},
			{
				Lo:     0x1BC90,
				Hi:     0x1BC99,
				Stride: 1,
			},
			{
				Lo:     0x1BC9C,
				Hi:     0x1BC9C,
				Stride: 1,
			},
			{
				Lo:     0x1BC9F,
				Hi:     0x1BC9F,
				Stride: 1,
			},
			{
				Lo:     0x1D000,
				Hi:     0x1D0F5,
				Stride: 1,
			},
			{
				Lo:     0x1D100,
				Hi:     0x1D126,
				Stride: 1,
			},
			{
				Lo:     0x1D129,
				Hi:     0x1D164,
				Stride: 1,
			},
			{
				Lo:     0x1D165,
				Hi:     0x1D166,
				Stride: 1,
			},
			{
				Lo:     0x1D16A,
				Hi:     0x1D16C,
				Stride: 1,
			},
			{
				Lo:     0x1D16D,
				Hi:     0x1D172,
				Stride: 1,
			},
			{
				Lo:     0x1D183,
				Hi:     0x1D184,
				Stride: 1,
			},
			{
				Lo:     0x1D18C,
				Hi:     0x1D1A9,
				Stride: 1,
			},
			{
				Lo:     0x1D1AE,
				Hi:     0x1D1E8,
				Stride: 1,
			},
			{
				Lo:     0x1D360,
				Hi:     0x1D371,
				Stride: 1,
			},
			{
				Lo:     0x1D400,
				Hi:     0x1D454,
				Stride: 1,
			},
			{
				Lo:     0x1D456,
				Hi:     0x1D49C,
				Stride: 1,
			},
			{
				Lo:     0x1D49E,
				Hi:     0x1D49F,
				Stride: 1,
			},
			{
				Lo:     0x1D4A2,
				Hi:     0x1D4A2,
				Stride: 1,
			},
			{
				Lo:     0x1D4A5,
				Hi:     0x1D4A6,
				Stride: 1,
			},
			{
				Lo:     0x1D4A9,
				Hi:     0x1D4AC,
				Stride: 1,
			},
			{
				Lo:     0x1D4AE,
				Hi:     0x1D4B9,
				Stride: 1,
			},
			{
				Lo:     0x1D4BB,
				Hi:     0x1D4BB,
				Stride: 1,
			},
			{
				Lo:     0x1D4BD,
				Hi:     0x1D4C3,
				Stride: 1,
			},
			{
				Lo:     0x1D4C5,
				Hi:     0x1D505,
				Stride: 1,
			},
			{
				Lo:     0x1D507,
				Hi:     0x1D50A,
				Stride: 1,
			},
			{
				Lo:     0x1D50D,
				Hi:     0x1D514,
				Stride: 1,
			},
			{
				Lo:     0x1D516,
				Hi:     0x1D51C,
				Stride: 1,
			},
			{
				Lo:     0x1D51E,
				Hi:     0x1D539,
				Stride: 1,
			},
			{
				Lo:     0x1D53B,
				Hi:     0x1D53E,
				Stride: 1,
			},
			{
				Lo:     0x1D540,
				Hi:     0x1D544,
				Stride: 1,
			},
			{
				Lo:     0x1D546,
				Hi:     0x1D546,
				Stride: 1,
			},
			{
				Lo:     0x1D54A,
				Hi:     0x1D550,
				Stride: 1,
			},
			{
				Lo:     0x1D552,
				Hi:     0x1D6A5,
				Stride: 1,
			},
			{
				Lo:     0x1D6A8,
				Hi:     0x1D6C0,
				Stride: 1,
			},
			{
				Lo:     0x1D6C1,
				Hi:     0x1D6C1,
				Stride: 1,
			},
			{
				Lo:     0x1D6C2,
				Hi:     0x1D6DA,
				Stride: 1,
			},
			{
				Lo:     0x1D6DC,
				Hi:     0x1D6FA,
				Stride: 1,
			},
			{
				Lo:     0x1D6FB,
				Hi:     0x1D6FB,
				Stride: 1,
			},
			{
				Lo:     0x1D6FC,
				Hi:     0x1D714,
				Stride: 1,
			},
			{
				Lo:     0x1D716,
				Hi:     0x1D734,
				Stride: 1,
			},
			{
				Lo:     0x1D735,
				Hi:     0x1D735,
				Stride: 1,
			},
			{
				Lo:     0x1D736,
				Hi:     0x1D74E,
				Stride: 1,
			},
			{
				Lo:     0x1D750,
				Hi:     0x1D76E,
				Stride: 1,
			},
			{
				Lo:     0x1D76F,
				Hi:     0x1D76F,
				Stride: 1,
			},
			{
				Lo:     0x1D770,
				Hi:     0x1D788,
				Stride: 1,
			},
			{
				Lo:     0x1D78A,
				Hi:     0x1D7A8,
				Stride: 1,
			},
			{
				Lo:     0x1D7A9,
				Hi:     0x1D7A9,
				Stride: 1,
			},
			{
				Lo:     0x1D7AA,
				Hi:     0x1D7C2,
				Stride: 1,
			},
			{
				Lo:     0x1D7C4,
				Hi:     0x1D7CB,
				Stride: 1,
			},
			{
				Lo:     0x1D800,
				Hi:     0x1D9FF,
				Stride: 1,
			},
			{
				Lo:     0x1DA37,
				Hi:     0x1DA3A,
				Stride: 1,
			},
			{
				Lo:     0x1DA6D,
				Hi:     0x1DA74,
				Stride: 1,
			},
			{
				Lo:     0x1DA76,
				Hi:     0x1DA83,
				Stride: 1,
			},
			{
				Lo:     0x1DA85,
				Hi:     0x1DA86,
				Stride: 1,
			},
			{
				Lo:     0x1DA87,
				Hi:     0x1DA8B,
				Stride: 1,
			},
			{
				Lo:     0x1F110,
				Hi:     0x1F12E,
				Stride: 1,
			},
			{
				Lo:     0x1F130,
				Hi:     0x1F169,
				Stride: 1,
			},
			{
				Lo:     0x1F170,
				Hi:     0x1F19A,
				Stride: 1,
			},
			{
				Lo:     0x1F1E6,
				Hi:     0x1F202,
				Stride: 1,
			},
			{
				Lo:     0x1F210,
				Hi:     0x1F23A,
				Stride: 1,
			},
			{
				Lo:     0x1F240,
				Hi:     0x1F248,
				Stride: 1,
			},
			{
				Lo:     0x1F250,
				Hi:     0x1F251,
				Stride: 1,
			},
			{
				Lo:     0x20000,
				Hi:     0x2A6D6,
				Stride: 1,
			},
			{
				Lo:     0x2A700,
				Hi:     0x2B734,
				Stride: 1,
			},
			{
				Lo:     0x2B740,
				Hi:     0x2B81D,
				Stride: 1,
			},
			{
				Lo:     0x2B820,
				Hi:     0x2CEA1,
				Stride: 1,
			},
			{
				Lo:     0x2F800,
				Hi:     0x2FA1D,
				Stride: 1,
			},
			{
				Lo:     0xF0000,
				Hi:     0xFFFFD,
				Stride: 1,
			},
			{
				Lo:     0x100000,
				Hi:     0x10FFFD,
				Stride: 1,
			},
		},
	}

	bidiClassLRE = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x202A,
				Hi:     0x202A,
				Stride: 1,
			},
		},
	}

	bidiClassLRI = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x2066,
				Hi:     0x2066,
				Stride: 1,
			},
		},
	}

	bidiClassLRO = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x202D,
				Hi:     0x202D,
				Stride: 1,
			},
		},
	}

	bidiClassNSM = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0300,
				Hi:     0x036F,
				Stride: 1,
			},
			{
				Lo:     0x0483,
				Hi:     0x0487,
				Stride: 1,
			},
			{
				Lo:     0x0488,
				Hi:     0x0489,
				Stride: 1,
			},
			{
				Lo:     0x0591,
				Hi:     0x05BD,
				Stride: 1,
			},
			{
				Lo:     0x05BF,
				Hi:     0x05BF,
				Stride: 1,
			},
			{
				Lo:     0x05C1,
				Hi:     0x05C2,
				Stride: 1,
			},
			{
				Lo:     0x05C4,
				Hi:     0x05C5,
				Stride: 1,
			},
			{
				Lo:     0x05C7,
				Hi:     0x05C7,
				Stride: 1,
			},
			{
				Lo:     0x0610,
				Hi:     0x061A,
				Stride: 1,
			},
			{
				Lo:     0x064B,
				Hi:     0x065F,
				Stride: 1,
			},
			{
				Lo:     0x0670,
				Hi:     0x0670,
				Stride: 1,
			},
			{
				Lo:     0x06D6,
				Hi:     0x06DC,
				Stride: 1,
			},
			{
				Lo:     0x06DF,
				Hi:     0x06E4,
				Stride: 1,
			},
			{
				Lo:     0x06E7,
				Hi:     0x06E8,
				Stride: 1,
			},
			{
				Lo:     0x06EA,
				Hi:     0x06ED,
				Stride: 1,
			},
			{
				Lo:     0x0711,
				Hi:     0x0711,
				Stride: 1,
			},
			{
				Lo:     0x0730,
				Hi:     0x074A,
				Stride: 1,
			},
			{
				Lo:     0x07A6,
				Hi:     0x07B0,
				Stride: 1,
			},
			{
				Lo:     0x07EB,
				Hi:     0x07F3,
				Stride: 1,
			},
			{
				Lo:     0x0816,
				Hi:     0x0819,
				Stride: 1,
			},
			{
				Lo:     0x081B,
				Hi:     0x0823,
				Stride: 1,
			},
			{
				Lo:     0x0825,
				Hi:     0x0827,
				Stride: 1,
			},
			{
				Lo:     0x0829,
				Hi:     0x082D,
				Stride: 1,
			},
			{
				Lo:     0x0859,
				Hi:     0x085B,
				Stride: 1,
			},
			{
				Lo:     0x08E3,
				Hi:     0x0902,
				Stride: 1,
			},
			{
				Lo:     0x093A,
				Hi:     0x093A,
				Stride: 1,
			},
			{
				Lo:     0x093C,
				Hi:     0x093C,
				Stride: 1,
			},
			{
				Lo:     0x0941,
				Hi:     0x0948,
				Stride: 1,
			},
			{
				Lo:     0x094D,
				Hi:     0x094D,
				Stride: 1,
			},
			{
				Lo:     0x0951,
				Hi:     0x0957,
				Stride: 1,
			},
			{
				Lo:     0x0962,
				Hi:     0x0963,
				Stride: 1,
			},
			{
				Lo:     0x0981,
				Hi:     0x0981,
				Stride: 1,
			},
			{
				Lo:     0x09BC,
				Hi:     0x09BC,
				Stride: 1,
			},
			{
				Lo:     0x09C1,
				Hi:     0x09C4,
				Stride: 1,
			},
			{
				Lo:     0x09CD,
				Hi:     0x09CD,
				Stride: 1,
			},
			{
				Lo:     0x09E2,
				Hi:     0x09E3,
				Stride: 1,
			},
			{
				Lo:     0x0A01,
				Hi:     0x0A02,
				Stride: 1,
			},
			{
				Lo:     0x0A3C,
				Hi:     0x0A3C,
				Stride: 1,
			},
			{
				Lo:     0x0A41,
				Hi:     0x0A42,
				Stride: 1,
			},
			{
				Lo:     0x0A47,
				Hi:     0x0A48,
				Stride: 1,
			},
			{
				Lo:     0x0A4B,
				Hi:     0x0A4D,
				Stride: 1,
			},
			{
				Lo:     0x0A51,
				Hi:     0x0A51,
				Stride: 1,
			},
			{
				Lo:     0x0A70,
				Hi:     0x0A71,
				Stride: 1,
			},
			{
				Lo:     0x0A75,
				Hi:     0x0A75,
				Stride: 1,
			},
			{
				Lo:     0x0A81,
				Hi:     0x0A82,
				Stride: 1,
			},
			{
				Lo:     0x0ABC,
				Hi:     0x0ABC,
				Stride: 1,
			},
			{
				Lo:     0x0AC1,
				Hi:     0x0AC5,
				Stride: 1,
			},
			{
				Lo:     0x0AC7,
				Hi:     0x0AC8,
				Stride: 1,
			},
			{
				Lo:     0x0ACD,
				Hi:     0x0ACD,
				Stride: 1,
			},
			{
				Lo:     0x0AE2,
				Hi:     0x0AE3,
				Stride: 1,
			},
			{
				Lo:     0x0B01,
				Hi:     0x0B01,
				Stride: 1,
			},
			{
				Lo:     0x0B3C,
				Hi:     0x0B3C,
				Stride: 1,
			},
			{
				Lo:     0x0B3F,
				Hi:     0x0B3F,
				Stride: 1,
			},
			{
				Lo:     0x0B41,
				Hi:     0x0B44,
				Stride: 1,
			},
			{
				Lo:     0x0B4D,
				Hi:     0x0B4D,
				Stride: 1,
			},
			{
				Lo:     0x0B56,
				Hi:     0x0B56,
				Stride: 1,
			},
			{
				Lo:     0x0B62,
				Hi:     0x0B63,
				Stride: 1,
			},
			{
				Lo:     0x0B82,
				Hi:     0x0B82,
				Stride: 1,
			},
			{
				Lo:     0x0BC0,
				Hi:     0x0BC0,
				Stride: 1,
			},
			{
				Lo:     0x0BCD,
				Hi:     0x0BCD,
				Stride: 1,
			},
			{
				Lo:     0x0C00,
				Hi:     0x0C00,
				Stride: 1,
			},
			{
				Lo:     0x0C3E,
				Hi:     0x0C40,
				Stride: 1,
			},
			{
				Lo:     0x0C46,
				Hi:     0x0C48,
				Stride: 1,
			},
			{
				Lo:     0x0C4A,
				Hi:     0x0C4D,
				Stride: 1,
			},
			{
				Lo:     0x0C55,
				Hi:     0x0C56,
				Stride: 1,
			},
			{
				Lo:     0x0C62,
				Hi:     0x0C63,
				Stride: 1,
			},
			{
				Lo:     0x0C81,
				Hi:     0x0C81,
				Stride: 1,
			},
			{
				Lo:     0x0CBC,
				Hi:     0x0CBC,
				Stride: 1,
			},
			{
				Lo:     0x0CCC,
				Hi:     0x0CCD,
				Stride: 1,
			},
			{
				Lo:     0x0CE2,
				Hi:     0x0CE3,
				Stride: 1,
			},
			{
				Lo:     0x0D01,
				Hi:     0x0D01,
				Stride: 1,
			},
			{
				Lo:     0x0D41,
				Hi:     0x0D44,
				Stride: 1,
			},
			{
				Lo:     0x0D4D,
				Hi:     0x0D4D,
				Stride: 1,
			},
			{
				Lo:     0x0D62,
				Hi:     0x0D63,
				Stride: 1,
			},
			{
				Lo:     0x0DCA,
				Hi:     0x0DCA,
				Stride: 1,
			},
			{
				Lo:     0x0DD2,
				Hi:     0x0DD4,
				Stride: 1,
			},
			{
				Lo:     0x0DD6,
				Hi:     0x0DD6,
				Stride: 1,
			},
			{
				Lo:     0x0E31,
				Hi:     0x0E31,
				Stride: 1,
			},
			{
				Lo:     0x0E34,
				Hi:     0x0E3A,
				Stride: 1,
			},
			{
				Lo:     0x0E47,
				Hi:     0x0E4E,
				Stride: 1,
			},
			{
				Lo:     0x0EB1,
				Hi:     0x0EB1,
				Stride: 1,
			},
			{
				Lo:     0x0EB4,
				Hi:     0x0EB9,
				Stride: 1,
			},
			{
				Lo:     0x0EBB,
				Hi:     0x0EBC,
				Stride: 1,
			},
			{
				Lo:     0x0EC8,
				Hi:     0x0ECD,
				Stride: 1,
			},
			{
				Lo:     0x0F18,
				Hi:     0x0F19,
				Stride: 1,
			},
			{
				Lo:     0x0F35,
				Hi:     0x0F35,
				Stride: 1,
			},
			{
				Lo:     0x0F37,
				Hi:     0x0F37,
				Stride: 1,
			},
			{
				Lo:     0x0F39,
				Hi:     0x0F39,
				Stride: 1,
			},
			{
				Lo:     0x0F71,
				Hi:     0x0F7E,
				Stride: 1,
			},
			{
				Lo:     0x0F80,
				Hi:     0x0F84,
				Stride: 1,
			},
			{
				Lo:     0x0F86,
				Hi:     0x0F87,
				Stride: 1,
			},
			{
				Lo:     0x0F8D,
				Hi:     0x0F97,
				Stride: 1,
			},
			{
				Lo:     0x0F99,
				Hi:     0x0FBC,
				Stride: 1,
			},
			{
				Lo:     0x0FC6,
				Hi:     0x0FC6,
				Stride: 1,
			},
			{
				Lo:     0x102D,
				Hi:     0x1030,
				Stride: 1,
			},
			{
				Lo:     0x1032,
				Hi:     0x1037,
				Stride: 1,
			},
			{
				Lo:     0x1039,
				Hi:     0x103A,
				Stride: 1,
			},
			{
				Lo:     0x103D,
				Hi:     0x103E,
				Stride: 1,
			},
			{
				Lo:     0x1058,
				Hi:     0x1059,
				Stride: 1,
			},
			{
				Lo:     0x105E,
				Hi:     0x1060,
				Stride: 1,
			},
			{
				Lo:     0x1071,
				Hi:     0x1074,
				Stride: 1,
			},
			{
				Lo:     0x1082,
				Hi:     0x1082,
				Stride: 1,
			},
			{
				Lo:     0x1085,
				Hi:     0x1086,
				Stride: 1,
			},
			{
				Lo:     0x108D,
				Hi:     0x108D,
				Stride: 1,
			},
			{
				Lo:     0x109D,
				Hi:     0x109D,
				Stride: 1,
			},
			{
				Lo:     0x135D,
				Hi:     0x135F,
				Stride: 1,
			},
			{
				Lo:     0x1712,
				Hi:     0x1714,
				Stride: 1,
			},
			{
				Lo:     0x1732,
				Hi:     0x1734,
				Stride: 1,
			},
			{
				Lo:     0x1752,
				Hi:     0x1753,
				Stride: 1,
			},
			{
				Lo:     0x1772,
				Hi:     0x1773,
				Stride: 1,
			},
			{
				Lo:     0x17B4,
				Hi:     0x17B5,
				Stride: 1,
			},
			{
				Lo:     0x17B7,
				Hi:     0x17BD,
				Stride: 1,
			},
			{
				Lo:     0x17C6,
				Hi:     0x17C6,
				Stride: 1,
			},
			{
				Lo:     0x17C9,
				Hi:     0x17D3,
				Stride: 1,
			},
			{
				Lo:     0x17DD,
				Hi:     0x17DD,
				Stride: 1,
			},
			{
				Lo:     0x180B,
				Hi:     0x180D,
				Stride: 1,
			},
			{
				Lo:     0x18A9,
				Hi:     0x18A9,
				Stride: 1,
			},
			{
				Lo:     0x1920,
				Hi:     0x1922,
				Stride: 1,
			},
			{
				Lo:     0x1927,
				Hi:     0x1928,
				Stride: 1,
			},
			{
				Lo:     0x1932,
				Hi:     0x1932,
				Stride: 1,
			},
			{
				Lo:     0x1939,
				Hi:     0x193B,
				Stride: 1,
			},
			{
				Lo:     0x1A17,
				Hi:     0x1A18,
				Stride: 1,
			},
			{
				Lo:     0x1A1B,
				Hi:     0x1A1B,
				Stride: 1,
			},
			{
				Lo:     0x1A56,
				Hi:     0x1A56,
				Stride: 1,
			},
			{
				Lo:     0x1A58,
				Hi:     0x1A5E,
				Stride: 1,
			},
			{
				Lo:     0x1A60,
				Hi:     0x1A60,
				Stride: 1,
			},
			{
				Lo:     0x1A62,
				Hi:     0x1A62,
				Stride: 1,
			},
			{
				Lo:     0x1A65,
				Hi:     0x1A6C,
				Stride: 1,
			},
			{
				Lo:     0x1A73,
				Hi:     0x1A7C,
				Stride: 1,
			},
			{
				Lo:     0x1A7F,
				Hi:     0x1A7F,
				Stride: 1,
			},
			{
				Lo:     0x1AB0,
				Hi:     0x1ABD,
				Stride: 1,
			},
			{
				Lo:     0x1ABE,
				Hi:     0x1ABE,
				Stride: 1,
			},
			{
				Lo:     0x1B00,
				Hi:     0x1B03,
				Stride: 1,
			},
			{
				Lo:     0x1B34,
				Hi:     0x1B34,
				Stride: 1,
			},
			{
				Lo:     0x1B36,
				Hi:     0x1B3A,
				Stride: 1,
			},
			{
				Lo:     0x1B3C,
				Hi:     0x1B3C,
				Stride: 1,
			},
			{
				Lo:     0x1B42,
				Hi:     0x1B42,
				Stride: 1,
			},
			{
				Lo:     0x1B6B,
				Hi:     0x1B73,
				Stride: 1,
			},
			{
				Lo:     0x1B80,
				Hi:     0x1B81,
				Stride: 1,
			},
			{
				Lo:     0x1BA2,
				Hi:     0x1BA5,
				Stride: 1,
			},
			{
				Lo:     0x1BA8,
				Hi:     0x1BA9,
				Stride: 1,
			},
			{
				Lo:     0x1BAB,
				Hi:     0x1BAD,
				Stride: 1,
			},
			{
				Lo:     0x1BE6,
				Hi:     0x1BE6,
				Stride: 1,
			},
			{
				Lo:     0x1BE8,
				Hi:     0x1BE9,
				Stride: 1,
			},
			{
				Lo:     0x1BED,
				Hi:     0x1BED,
				Stride: 1,
			},
			{
				Lo:     0x1BEF,
				Hi:     0x1BF1,
				Stride: 1,
			},
			{
				Lo:     0x1C2C,
				Hi:     0x1C33,
				Stride: 1,
			},
			{
				Lo:     0x1C36,
				Hi:     0x1C37,
				Stride: 1,
			},
			{
				Lo:     0x1CD0,
				Hi:     0x1CD2,
				Stride: 1,
			},
			{
				Lo:     0x1CD4,
				Hi:     0x1CE0,
				Stride: 1,
			},
			{
				Lo:     0x1CE2,
				Hi:     0x1CE8,
				Stride: 1,
			},
			{
				Lo:     0x1CED,
				Hi:     0x1CED,
				Stride: 1,
			},
			{
				Lo:     0x1CF4,
				Hi:     0x1CF4,
				Stride: 1,
			},
			{
				Lo:     0x1CF8,
				Hi:     0x1CF9,
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
				Hi:     0x20DC,
				Stride: 1,
			},
			{
				Lo:     0x20DD,
				Hi:     0x20E0,
				Stride: 1,
			},
			{
				Lo:     0x20E1,
				Hi:     0x20E1,
				Stride: 1,
			},
			{
				Lo:     0x20E2,
				Hi:     0x20E4,
				Stride: 1,
			},
			{
				Lo:     0x20E5,
				Hi:     0x20F0,
				Stride: 1,
			},
			{
				Lo:     0x2CEF,
				Hi:     0x2CF1,
				Stride: 1,
			},
			{
				Lo:     0x2D7F,
				Hi:     0x2D7F,
				Stride: 1,
			},
			{
				Lo:     0x2DE0,
				Hi:     0x2DFF,
				Stride: 1,
			},
			{
				Lo:     0x302A,
				Hi:     0x302D,
				Stride: 1,
			},
			{
				Lo:     0x3099,
				Hi:     0x309A,
				Stride: 1,
			},
			{
				Lo:     0xA66F,
				Hi:     0xA66F,
				Stride: 1,
			},
			{
				Lo:     0xA670,
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
				Lo:     0xA802,
				Hi:     0xA802,
				Stride: 1,
			},
			{
				Lo:     0xA806,
				Hi:     0xA806,
				Stride: 1,
			},
			{
				Lo:     0xA80B,
				Hi:     0xA80B,
				Stride: 1,
			},
			{
				Lo:     0xA825,
				Hi:     0xA826,
				Stride: 1,
			},
			{
				Lo:     0xA8C4,
				Hi:     0xA8C4,
				Stride: 1,
			},
			{
				Lo:     0xA8E0,
				Hi:     0xA8F1,
				Stride: 1,
			},
			{
				Lo:     0xA926,
				Hi:     0xA92D,
				Stride: 1,
			},
			{
				Lo:     0xA947,
				Hi:     0xA951,
				Stride: 1,
			},
			{
				Lo:     0xA980,
				Hi:     0xA982,
				Stride: 1,
			},
			{
				Lo:     0xA9B3,
				Hi:     0xA9B3,
				Stride: 1,
			},
			{
				Lo:     0xA9B6,
				Hi:     0xA9B9,
				Stride: 1,
			},
			{
				Lo:     0xA9BC,
				Hi:     0xA9BC,
				Stride: 1,
			},
			{
				Lo:     0xA9E5,
				Hi:     0xA9E5,
				Stride: 1,
			},
			{
				Lo:     0xAA29,
				Hi:     0xAA2E,
				Stride: 1,
			},
			{
				Lo:     0xAA31,
				Hi:     0xAA32,
				Stride: 1,
			},
			{
				Lo:     0xAA35,
				Hi:     0xAA36,
				Stride: 1,
			},
			{
				Lo:     0xAA43,
				Hi:     0xAA43,
				Stride: 1,
			},
			{
				Lo:     0xAA4C,
				Hi:     0xAA4C,
				Stride: 1,
			},
			{
				Lo:     0xAA7C,
				Hi:     0xAA7C,
				Stride: 1,
			},
			{
				Lo:     0xAAB0,
				Hi:     0xAAB0,
				Stride: 1,
			},
			{
				Lo:     0xAAB2,
				Hi:     0xAAB4,
				Stride: 1,
			},
			{
				Lo:     0xAAB7,
				Hi:     0xAAB8,
				Stride: 1,
			},
			{
				Lo:     0xAABE,
				Hi:     0xAABF,
				Stride: 1,
			},
			{
				Lo:     0xAAC1,
				Hi:     0xAAC1,
				Stride: 1,
			},
			{
				Lo:     0xAAEC,
				Hi:     0xAAED,
				Stride: 1,
			},
			{
				Lo:     0xAAF6,
				Hi:     0xAAF6,
				Stride: 1,
			},
			{
				Lo:     0xABE5,
				Hi:     0xABE5,
				Stride: 1,
			},
			{
				Lo:     0xABE8,
				Hi:     0xABE8,
				Stride: 1,
			},
			{
				Lo:     0xABED,
				Hi:     0xABED,
				Stride: 1,
			},
			{
				Lo:     0xFB1E,
				Hi:     0xFB1E,
				Stride: 1,
			},
			{
				Lo:     0xFE00,
				Hi:     0xFE0F,
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
				Lo:     0x101FD,
				Hi:     0x101FD,
				Stride: 1,
			},
			{
				Lo:     0x102E0,
				Hi:     0x102E0,
				Stride: 1,
			},
			{
				Lo:     0x10376,
				Hi:     0x1037A,
				Stride: 1,
			},
			{
				Lo:     0x10A01,
				Hi:     0x10A03,
				Stride: 1,
			},
			{
				Lo:     0x10A05,
				Hi:     0x10A06,
				Stride: 1,
			},
			{
				Lo:     0x10A0C,
				Hi:     0x10A0F,
				Stride: 1,
			},
			{
				Lo:     0x10A38,
				Hi:     0x10A3A,
				Stride: 1,
			},
			{
				Lo:     0x10A3F,
				Hi:     0x10A3F,
				Stride: 1,
			},
			{
				Lo:     0x10AE5,
				Hi:     0x10AE6,
				Stride: 1,
			},
			{
				Lo:     0x11001,
				Hi:     0x11001,
				Stride: 1,
			},
			{
				Lo:     0x11038,
				Hi:     0x11046,
				Stride: 1,
			},
			{
				Lo:     0x1107F,
				Hi:     0x11081,
				Stride: 1,
			},
			{
				Lo:     0x110B3,
				Hi:     0x110B6,
				Stride: 1,
			},
			{
				Lo:     0x110B9,
				Hi:     0x110BA,
				Stride: 1,
			},
			{
				Lo:     0x11100,
				Hi:     0x11102,
				Stride: 1,
			},
			{
				Lo:     0x11127,
				Hi:     0x1112B,
				Stride: 1,
			},
			{
				Lo:     0x1112D,
				Hi:     0x11134,
				Stride: 1,
			},
			{
				Lo:     0x11173,
				Hi:     0x11173,
				Stride: 1,
			},
			{
				Lo:     0x11180,
				Hi:     0x11181,
				Stride: 1,
			},
			{
				Lo:     0x111B6,
				Hi:     0x111BE,
				Stride: 1,
			},
			{
				Lo:     0x111CA,
				Hi:     0x111CC,
				Stride: 1,
			},
			{
				Lo:     0x1122F,
				Hi:     0x11231,
				Stride: 1,
			},
			{
				Lo:     0x11234,
				Hi:     0x11234,
				Stride: 1,
			},
			{
				Lo:     0x11236,
				Hi:     0x11237,
				Stride: 1,
			},
			{
				Lo:     0x112DF,
				Hi:     0x112DF,
				Stride: 1,
			},
			{
				Lo:     0x112E3,
				Hi:     0x112EA,
				Stride: 1,
			},
			{
				Lo:     0x11300,
				Hi:     0x11301,
				Stride: 1,
			},
			{
				Lo:     0x1133C,
				Hi:     0x1133C,
				Stride: 1,
			},
			{
				Lo:     0x11340,
				Hi:     0x11340,
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
				Lo:     0x114B3,
				Hi:     0x114B8,
				Stride: 1,
			},
			{
				Lo:     0x114BA,
				Hi:     0x114BA,
				Stride: 1,
			},
			{
				Lo:     0x114BF,
				Hi:     0x114C0,
				Stride: 1,
			},
			{
				Lo:     0x114C2,
				Hi:     0x114C3,
				Stride: 1,
			},
			{
				Lo:     0x115B2,
				Hi:     0x115B5,
				Stride: 1,
			},
			{
				Lo:     0x115BC,
				Hi:     0x115BD,
				Stride: 1,
			},
			{
				Lo:     0x115BF,
				Hi:     0x115C0,
				Stride: 1,
			},
			{
				Lo:     0x115DC,
				Hi:     0x115DD,
				Stride: 1,
			},
			{
				Lo:     0x11633,
				Hi:     0x1163A,
				Stride: 1,
			},
			{
				Lo:     0x1163D,
				Hi:     0x1163D,
				Stride: 1,
			},
			{
				Lo:     0x1163F,
				Hi:     0x11640,
				Stride: 1,
			},
			{
				Lo:     0x116AB,
				Hi:     0x116AB,
				Stride: 1,
			},
			{
				Lo:     0x116AD,
				Hi:     0x116AD,
				Stride: 1,
			},
			{
				Lo:     0x116B0,
				Hi:     0x116B5,
				Stride: 1,
			},
			{
				Lo:     0x116B7,
				Hi:     0x116B7,
				Stride: 1,
			},
			{
				Lo:     0x1171D,
				Hi:     0x1171F,
				Stride: 1,
			},
			{
				Lo:     0x11722,
				Hi:     0x11725,
				Stride: 1,
			},
			{
				Lo:     0x11727,
				Hi:     0x1172B,
				Stride: 1,
			},
			{
				Lo:     0x16AF0,
				Hi:     0x16AF4,
				Stride: 1,
			},
			{
				Lo:     0x16B30,
				Hi:     0x16B36,
				Stride: 1,
			},
			{
				Lo:     0x16F8F,
				Hi:     0x16F92,
				Stride: 1,
			},
			{
				Lo:     0x1BC9D,
				Hi:     0x1BC9E,
				Stride: 1,
			},
			{
				Lo:     0x1D167,
				Hi:     0x1D169,
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
			{
				Lo:     0x1DA00,
				Hi:     0x1DA36,
				Stride: 1,
			},
			{
				Lo:     0x1DA3B,
				Hi:     0x1DA6C,
				Stride: 1,
			},
			{
				Lo:     0x1DA75,
				Hi:     0x1DA75,
				Stride: 1,
			},
			{
				Lo:     0x1DA84,
				Hi:     0x1DA84,
				Stride: 1,
			},
			{
				Lo:     0x1DA9B,
				Hi:     0x1DA9F,
				Stride: 1,
			},
			{
				Lo:     0x1DAA1,
				Hi:     0x1DAAF,
				Stride: 1,
			},
			{
				Lo:     0x1E8D0,
				Hi:     0x1E8D6,
				Stride: 1,
			},
			{
				Lo:     0xE0100,
				Hi:     0xE01EF,
				Stride: 1,
			},
		},
	}

	bidiClassON = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0021,
				Hi:     0x0022,
				Stride: 1,
			},
			{
				Lo:     0x0026,
				Hi:     0x0027,
				Stride: 1,
			},
			{
				Lo:     0x0028,
				Hi:     0x0028,
				Stride: 1,
			},
			{
				Lo:     0x0029,
				Hi:     0x0029,
				Stride: 1,
			},
			{
				Lo:     0x002A,
				Hi:     0x002A,
				Stride: 1,
			},
			{
				Lo:     0x003B,
				Hi:     0x003B,
				Stride: 1,
			},
			{
				Lo:     0x003C,
				Hi:     0x003E,
				Stride: 1,
			},
			{
				Lo:     0x003F,
				Hi:     0x0040,
				Stride: 1,
			},
			{
				Lo:     0x005B,
				Hi:     0x005B,
				Stride: 1,
			},
			{
				Lo:     0x005C,
				Hi:     0x005C,
				Stride: 1,
			},
			{
				Lo:     0x005D,
				Hi:     0x005D,
				Stride: 1,
			},
			{
				Lo:     0x005E,
				Hi:     0x005E,
				Stride: 1,
			},
			{
				Lo:     0x005F,
				Hi:     0x005F,
				Stride: 1,
			},
			{
				Lo:     0x0060,
				Hi:     0x0060,
				Stride: 1,
			},
			{
				Lo:     0x007B,
				Hi:     0x007B,
				Stride: 1,
			},
			{
				Lo:     0x007C,
				Hi:     0x007C,
				Stride: 1,
			},
			{
				Lo:     0x007D,
				Hi:     0x007D,
				Stride: 1,
			},
			{
				Lo:     0x007E,
				Hi:     0x007E,
				Stride: 1,
			},
			{
				Lo:     0x00A1,
				Hi:     0x00A1,
				Stride: 1,
			},
			{
				Lo:     0x00A6,
				Hi:     0x00A6,
				Stride: 1,
			},
			{
				Lo:     0x00A7,
				Hi:     0x00A7,
				Stride: 1,
			},
			{
				Lo:     0x00A8,
				Hi:     0x00A8,
				Stride: 1,
			},
			{
				Lo:     0x00A9,
				Hi:     0x00A9,
				Stride: 1,
			},
			{
				Lo:     0x00AB,
				Hi:     0x00AB,
				Stride: 1,
			},
			{
				Lo:     0x00AC,
				Hi:     0x00AC,
				Stride: 1,
			},
			{
				Lo:     0x00AE,
				Hi:     0x00AE,
				Stride: 1,
			},
			{
				Lo:     0x00AF,
				Hi:     0x00AF,
				Stride: 1,
			},
			{
				Lo:     0x00B4,
				Hi:     0x00B4,
				Stride: 1,
			},
			{
				Lo:     0x00B6,
				Hi:     0x00B7,
				Stride: 1,
			},
			{
				Lo:     0x00B8,
				Hi:     0x00B8,
				Stride: 1,
			},
			{
				Lo:     0x00BB,
				Hi:     0x00BB,
				Stride: 1,
			},
			{
				Lo:     0x00BC,
				Hi:     0x00BE,
				Stride: 1,
			},
			{
				Lo:     0x00BF,
				Hi:     0x00BF,
				Stride: 1,
			},
			{
				Lo:     0x00D7,
				Hi:     0x00D7,
				Stride: 1,
			},
			{
				Lo:     0x00F7,
				Hi:     0x00F7,
				Stride: 1,
			},
			{
				Lo:     0x02B9,
				Hi:     0x02BA,
				Stride: 1,
			},
			{
				Lo:     0x02C2,
				Hi:     0x02C5,
				Stride: 1,
			},
			{
				Lo:     0x02C6,
				Hi:     0x02CF,
				Stride: 1,
			},
			{
				Lo:     0x02D2,
				Hi:     0x02DF,
				Stride: 1,
			},
			{
				Lo:     0x02E5,
				Hi:     0x02EB,
				Stride: 1,
			},
			{
				Lo:     0x02EC,
				Hi:     0x02EC,
				Stride: 1,
			},
			{
				Lo:     0x02ED,
				Hi:     0x02ED,
				Stride: 1,
			},
			{
				Lo:     0x02EF,
				Hi:     0x02FF,
				Stride: 1,
			},
			{
				Lo:     0x0374,
				Hi:     0x0374,
				Stride: 1,
			},
			{
				Lo:     0x0375,
				Hi:     0x0375,
				Stride: 1,
			},
			{
				Lo:     0x037E,
				Hi:     0x037E,
				Stride: 1,
			},
			{
				Lo:     0x0384,
				Hi:     0x0385,
				Stride: 1,
			},
			{
				Lo:     0x0387,
				Hi:     0x0387,
				Stride: 1,
			},
			{
				Lo:     0x03F6,
				Hi:     0x03F6,
				Stride: 1,
			},
			{
				Lo:     0x058A,
				Hi:     0x058A,
				Stride: 1,
			},
			{
				Lo:     0x058D,
				Hi:     0x058E,
				Stride: 1,
			},
			{
				Lo:     0x0606,
				Hi:     0x0607,
				Stride: 1,
			},
			{
				Lo:     0x060E,
				Hi:     0x060F,
				Stride: 1,
			},
			{
				Lo:     0x06DE,
				Hi:     0x06DE,
				Stride: 1,
			},
			{
				Lo:     0x06E9,
				Hi:     0x06E9,
				Stride: 1,
			},
			{
				Lo:     0x07F6,
				Hi:     0x07F6,
				Stride: 1,
			},
			{
				Lo:     0x07F7,
				Hi:     0x07F9,
				Stride: 1,
			},
			{
				Lo:     0x0BF3,
				Hi:     0x0BF8,
				Stride: 1,
			},
			{
				Lo:     0x0BFA,
				Hi:     0x0BFA,
				Stride: 1,
			},
			{
				Lo:     0x0C78,
				Hi:     0x0C7E,
				Stride: 1,
			},
			{
				Lo:     0x0F3A,
				Hi:     0x0F3A,
				Stride: 1,
			},
			{
				Lo:     0x0F3B,
				Hi:     0x0F3B,
				Stride: 1,
			},
			{
				Lo:     0x0F3C,
				Hi:     0x0F3C,
				Stride: 1,
			},
			{
				Lo:     0x0F3D,
				Hi:     0x0F3D,
				Stride: 1,
			},
			{
				Lo:     0x1390,
				Hi:     0x1399,
				Stride: 1,
			},
			{
				Lo:     0x1400,
				Hi:     0x1400,
				Stride: 1,
			},
			{
				Lo:     0x169B,
				Hi:     0x169B,
				Stride: 1,
			},
			{
				Lo:     0x169C,
				Hi:     0x169C,
				Stride: 1,
			},
			{
				Lo:     0x17F0,
				Hi:     0x17F9,
				Stride: 1,
			},
			{
				Lo:     0x1800,
				Hi:     0x1805,
				Stride: 1,
			},
			{
				Lo:     0x1806,
				Hi:     0x1806,
				Stride: 1,
			},
			{
				Lo:     0x1807,
				Hi:     0x180A,
				Stride: 1,
			},
			{
				Lo:     0x1940,
				Hi:     0x1940,
				Stride: 1,
			},
			{
				Lo:     0x1944,
				Hi:     0x1945,
				Stride: 1,
			},
			{
				Lo:     0x19DE,
				Hi:     0x19FF,
				Stride: 1,
			},
			{
				Lo:     0x1FBD,
				Hi:     0x1FBD,
				Stride: 1,
			},
			{
				Lo:     0x1FBF,
				Hi:     0x1FC1,
				Stride: 1,
			},
			{
				Lo:     0x1FCD,
				Hi:     0x1FCF,
				Stride: 1,
			},
			{
				Lo:     0x1FDD,
				Hi:     0x1FDF,
				Stride: 1,
			},
			{
				Lo:     0x1FED,
				Hi:     0x1FEF,
				Stride: 1,
			},
			{
				Lo:     0x1FFD,
				Hi:     0x1FFE,
				Stride: 1,
			},
			{
				Lo:     0x2010,
				Hi:     0x2015,
				Stride: 1,
			},
			{
				Lo:     0x2016,
				Hi:     0x2017,
				Stride: 1,
			},
			{
				Lo:     0x2018,
				Hi:     0x2018,
				Stride: 1,
			},
			{
				Lo:     0x2019,
				Hi:     0x2019,
				Stride: 1,
			},
			{
				Lo:     0x201A,
				Hi:     0x201A,
				Stride: 1,
			},
			{
				Lo:     0x201B,
				Hi:     0x201C,
				Stride: 1,
			},
			{
				Lo:     0x201D,
				Hi:     0x201D,
				Stride: 1,
			},
			{
				Lo:     0x201E,
				Hi:     0x201E,
				Stride: 1,
			},
			{
				Lo:     0x201F,
				Hi:     0x201F,
				Stride: 1,
			},
			{
				Lo:     0x2020,
				Hi:     0x2027,
				Stride: 1,
			},
			{
				Lo:     0x2035,
				Hi:     0x2038,
				Stride: 1,
			},
			{
				Lo:     0x2039,
				Hi:     0x2039,
				Stride: 1,
			},
			{
				Lo:     0x203A,
				Hi:     0x203A,
				Stride: 1,
			},
			{
				Lo:     0x203B,
				Hi:     0x203E,
				Stride: 1,
			},
			{
				Lo:     0x203F,
				Hi:     0x2040,
				Stride: 1,
			},
			{
				Lo:     0x2041,
				Hi:     0x2043,
				Stride: 1,
			},
			{
				Lo:     0x2045,
				Hi:     0x2045,
				Stride: 1,
			},
			{
				Lo:     0x2046,
				Hi:     0x2046,
				Stride: 1,
			},
			{
				Lo:     0x2047,
				Hi:     0x2051,
				Stride: 1,
			},
			{
				Lo:     0x2052,
				Hi:     0x2052,
				Stride: 1,
			},
			{
				Lo:     0x2053,
				Hi:     0x2053,
				Stride: 1,
			},
			{
				Lo:     0x2054,
				Hi:     0x2054,
				Stride: 1,
			},
			{
				Lo:     0x2055,
				Hi:     0x205E,
				Stride: 1,
			},
			{
				Lo:     0x207C,
				Hi:     0x207C,
				Stride: 1,
			},
			{
				Lo:     0x207D,
				Hi:     0x207D,
				Stride: 1,
			},
			{
				Lo:     0x207E,
				Hi:     0x207E,
				Stride: 1,
			},
			{
				Lo:     0x208C,
				Hi:     0x208C,
				Stride: 1,
			},
			{
				Lo:     0x208D,
				Hi:     0x208D,
				Stride: 1,
			},
			{
				Lo:     0x208E,
				Hi:     0x208E,
				Stride: 1,
			},
			{
				Lo:     0x2100,
				Hi:     0x2101,
				Stride: 1,
			},
			{
				Lo:     0x2103,
				Hi:     0x2106,
				Stride: 1,
			},
			{
				Lo:     0x2108,
				Hi:     0x2109,
				Stride: 1,
			},
			{
				Lo:     0x2114,
				Hi:     0x2114,
				Stride: 1,
			},
			{
				Lo:     0x2116,
				Hi:     0x2117,
				Stride: 1,
			},
			{
				Lo:     0x2118,
				Hi:     0x2118,
				Stride: 1,
			},
			{
				Lo:     0x211E,
				Hi:     0x2123,
				Stride: 1,
			},
			{
				Lo:     0x2125,
				Hi:     0x2125,
				Stride: 1,
			},
			{
				Lo:     0x2127,
				Hi:     0x2127,
				Stride: 1,
			},
			{
				Lo:     0x2129,
				Hi:     0x2129,
				Stride: 1,
			},
			{
				Lo:     0x213A,
				Hi:     0x213B,
				Stride: 1,
			},
			{
				Lo:     0x2140,
				Hi:     0x2144,
				Stride: 1,
			},
			{
				Lo:     0x214A,
				Hi:     0x214A,
				Stride: 1,
			},
			{
				Lo:     0x214B,
				Hi:     0x214B,
				Stride: 1,
			},
			{
				Lo:     0x214C,
				Hi:     0x214D,
				Stride: 1,
			},
			{
				Lo:     0x2150,
				Hi:     0x215F,
				Stride: 1,
			},
			{
				Lo:     0x2189,
				Hi:     0x2189,
				Stride: 1,
			},
			{
				Lo:     0x218A,
				Hi:     0x218B,
				Stride: 1,
			},
			{
				Lo:     0x2190,
				Hi:     0x2194,
				Stride: 1,
			},
			{
				Lo:     0x2195,
				Hi:     0x2199,
				Stride: 1,
			},
			{
				Lo:     0x219A,
				Hi:     0x219B,
				Stride: 1,
			},
			{
				Lo:     0x219C,
				Hi:     0x219F,
				Stride: 1,
			},
			{
				Lo:     0x21A0,
				Hi:     0x21A0,
				Stride: 1,
			},
			{
				Lo:     0x21A1,
				Hi:     0x21A2,
				Stride: 1,
			},
			{
				Lo:     0x21A3,
				Hi:     0x21A3,
				Stride: 1,
			},
			{
				Lo:     0x21A4,
				Hi:     0x21A5,
				Stride: 1,
			},
			{
				Lo:     0x21A6,
				Hi:     0x21A6,
				Stride: 1,
			},
			{
				Lo:     0x21A7,
				Hi:     0x21AD,
				Stride: 1,
			},
			{
				Lo:     0x21AE,
				Hi:     0x21AE,
				Stride: 1,
			},
			{
				Lo:     0x21AF,
				Hi:     0x21CD,
				Stride: 1,
			},
			{
				Lo:     0x21CE,
				Hi:     0x21CF,
				Stride: 1,
			},
			{
				Lo:     0x21D0,
				Hi:     0x21D1,
				Stride: 1,
			},
			{
				Lo:     0x21D2,
				Hi:     0x21D2,
				Stride: 1,
			},
			{
				Lo:     0x21D3,
				Hi:     0x21D3,
				Stride: 1,
			},
			{
				Lo:     0x21D4,
				Hi:     0x21D4,
				Stride: 1,
			},
			{
				Lo:     0x21D5,
				Hi:     0x21F3,
				Stride: 1,
			},
			{
				Lo:     0x21F4,
				Hi:     0x2211,
				Stride: 1,
			},
			{
				Lo:     0x2214,
				Hi:     0x22FF,
				Stride: 1,
			},
			{
				Lo:     0x2300,
				Hi:     0x2307,
				Stride: 1,
			},
			{
				Lo:     0x2308,
				Hi:     0x2308,
				Stride: 1,
			},
			{
				Lo:     0x2309,
				Hi:     0x2309,
				Stride: 1,
			},
			{
				Lo:     0x230A,
				Hi:     0x230A,
				Stride: 1,
			},
			{
				Lo:     0x230B,
				Hi:     0x230B,
				Stride: 1,
			},
			{
				Lo:     0x230C,
				Hi:     0x231F,
				Stride: 1,
			},
			{
				Lo:     0x2320,
				Hi:     0x2321,
				Stride: 1,
			},
			{
				Lo:     0x2322,
				Hi:     0x2328,
				Stride: 1,
			},
			{
				Lo:     0x2329,
				Hi:     0x2329,
				Stride: 1,
			},
			{
				Lo:     0x232A,
				Hi:     0x232A,
				Stride: 1,
			},
			{
				Lo:     0x232B,
				Hi:     0x2335,
				Stride: 1,
			},
			{
				Lo:     0x237B,
				Hi:     0x237B,
				Stride: 1,
			},
			{
				Lo:     0x237C,
				Hi:     0x237C,
				Stride: 1,
			},
			{
				Lo:     0x237D,
				Hi:     0x2394,
				Stride: 1,
			},
			{
				Lo:     0x2396,
				Hi:     0x239A,
				Stride: 1,
			},
			{
				Lo:     0x239B,
				Hi:     0x23B3,
				Stride: 1,
			},
			{
				Lo:     0x23B4,
				Hi:     0x23DB,
				Stride: 1,
			},
			{
				Lo:     0x23DC,
				Hi:     0x23E1,
				Stride: 1,
			},
			{
				Lo:     0x23E2,
				Hi:     0x23FA,
				Stride: 1,
			},
			{
				Lo:     0x2400,
				Hi:     0x2426,
				Stride: 1,
			},
			{
				Lo:     0x2440,
				Hi:     0x244A,
				Stride: 1,
			},
			{
				Lo:     0x2460,
				Hi:     0x2487,
				Stride: 1,
			},
			{
				Lo:     0x24EA,
				Hi:     0x24FF,
				Stride: 1,
			},
			{
				Lo:     0x2500,
				Hi:     0x25B6,
				Stride: 1,
			},
			{
				Lo:     0x25B7,
				Hi:     0x25B7,
				Stride: 1,
			},
			{
				Lo:     0x25B8,
				Hi:     0x25C0,
				Stride: 1,
			},
			{
				Lo:     0x25C1,
				Hi:     0x25C1,
				Stride: 1,
			},
			{
				Lo:     0x25C2,
				Hi:     0x25F7,
				Stride: 1,
			},
			{
				Lo:     0x25F8,
				Hi:     0x25FF,
				Stride: 1,
			},
			{
				Lo:     0x2600,
				Hi:     0x266E,
				Stride: 1,
			},
			{
				Lo:     0x266F,
				Hi:     0x266F,
				Stride: 1,
			},
			{
				Lo:     0x2670,
				Hi:     0x26AB,
				Stride: 1,
			},
			{
				Lo:     0x26AD,
				Hi:     0x2767,
				Stride: 1,
			},
			{
				Lo:     0x2768,
				Hi:     0x2768,
				Stride: 1,
			},
			{
				Lo:     0x2769,
				Hi:     0x2769,
				Stride: 1,
			},
			{
				Lo:     0x276A,
				Hi:     0x276A,
				Stride: 1,
			},
			{
				Lo:     0x276B,
				Hi:     0x276B,
				Stride: 1,
			},
			{
				Lo:     0x276C,
				Hi:     0x276C,
				Stride: 1,
			},
			{
				Lo:     0x276D,
				Hi:     0x276D,
				Stride: 1,
			},
			{
				Lo:     0x276E,
				Hi:     0x276E,
				Stride: 1,
			},
			{
				Lo:     0x276F,
				Hi:     0x276F,
				Stride: 1,
			},
			{
				Lo:     0x2770,
				Hi:     0x2770,
				Stride: 1,
			},
			{
				Lo:     0x2771,
				Hi:     0x2771,
				Stride: 1,
			},
			{
				Lo:     0x2772,
				Hi:     0x2772,
				Stride: 1,
			},
			{
				Lo:     0x2773,
				Hi:     0x2773,
				Stride: 1,
			},
			{
				Lo:     0x2774,
				Hi:     0x2774,
				Stride: 1,
			},
			{
				Lo:     0x2775,
				Hi:     0x2775,
				Stride: 1,
			},
			{
				Lo:     0x2776,
				Hi:     0x2793,
				Stride: 1,
			},
			{
				Lo:     0x2794,
				Hi:     0x27BF,
				Stride: 1,
			},
			{
				Lo:     0x27C0,
				Hi:     0x27C4,
				Stride: 1,
			},
			{
				Lo:     0x27C5,
				Hi:     0x27C5,
				Stride: 1,
			},
			{
				Lo:     0x27C6,
				Hi:     0x27C6,
				Stride: 1,
			},
			{
				Lo:     0x27C7,
				Hi:     0x27E5,
				Stride: 1,
			},
			{
				Lo:     0x27E6,
				Hi:     0x27E6,
				Stride: 1,
			},
			{
				Lo:     0x27E7,
				Hi:     0x27E7,
				Stride: 1,
			},
			{
				Lo:     0x27E8,
				Hi:     0x27E8,
				Stride: 1,
			},
			{
				Lo:     0x27E9,
				Hi:     0x27E9,
				Stride: 1,
			},
			{
				Lo:     0x27EA,
				Hi:     0x27EA,
				Stride: 1,
			},
			{
				Lo:     0x27EB,
				Hi:     0x27EB,
				Stride: 1,
			},
			{
				Lo:     0x27EC,
				Hi:     0x27EC,
				Stride: 1,
			},
			{
				Lo:     0x27ED,
				Hi:     0x27ED,
				Stride: 1,
			},
			{
				Lo:     0x27EE,
				Hi:     0x27EE,
				Stride: 1,
			},
			{
				Lo:     0x27EF,
				Hi:     0x27EF,
				Stride: 1,
			},
			{
				Lo:     0x27F0,
				Hi:     0x27FF,
				Stride: 1,
			},
			{
				Lo:     0x2900,
				Hi:     0x2982,
				Stride: 1,
			},
			{
				Lo:     0x2983,
				Hi:     0x2983,
				Stride: 1,
			},
			{
				Lo:     0x2984,
				Hi:     0x2984,
				Stride: 1,
			},
			{
				Lo:     0x2985,
				Hi:     0x2985,
				Stride: 1,
			},
			{
				Lo:     0x2986,
				Hi:     0x2986,
				Stride: 1,
			},
			{
				Lo:     0x2987,
				Hi:     0x2987,
				Stride: 1,
			},
			{
				Lo:     0x2988,
				Hi:     0x2988,
				Stride: 1,
			},
			{
				Lo:     0x2989,
				Hi:     0x2989,
				Stride: 1,
			},
			{
				Lo:     0x298A,
				Hi:     0x298A,
				Stride: 1,
			},
			{
				Lo:     0x298B,
				Hi:     0x298B,
				Stride: 1,
			},
			{
				Lo:     0x298C,
				Hi:     0x298C,
				Stride: 1,
			},
			{
				Lo:     0x298D,
				Hi:     0x298D,
				Stride: 1,
			},
			{
				Lo:     0x298E,
				Hi:     0x298E,
				Stride: 1,
			},
			{
				Lo:     0x298F,
				Hi:     0x298F,
				Stride: 1,
			},
			{
				Lo:     0x2990,
				Hi:     0x2990,
				Stride: 1,
			},
			{
				Lo:     0x2991,
				Hi:     0x2991,
				Stride: 1,
			},
			{
				Lo:     0x2992,
				Hi:     0x2992,
				Stride: 1,
			},
			{
				Lo:     0x2993,
				Hi:     0x2993,
				Stride: 1,
			},
			{
				Lo:     0x2994,
				Hi:     0x2994,
				Stride: 1,
			},
			{
				Lo:     0x2995,
				Hi:     0x2995,
				Stride: 1,
			},
			{
				Lo:     0x2996,
				Hi:     0x2996,
				Stride: 1,
			},
			{
				Lo:     0x2997,
				Hi:     0x2997,
				Stride: 1,
			},
			{
				Lo:     0x2998,
				Hi:     0x2998,
				Stride: 1,
			},
			{
				Lo:     0x2999,
				Hi:     0x29D7,
				Stride: 1,
			},
			{
				Lo:     0x29D8,
				Hi:     0x29D8,
				Stride: 1,
			},
			{
				Lo:     0x29D9,
				Hi:     0x29D9,
				Stride: 1,
			},
			{
				Lo:     0x29DA,
				Hi:     0x29DA,
				Stride: 1,
			},
			{
				Lo:     0x29DB,
				Hi:     0x29DB,
				Stride: 1,
			},
			{
				Lo:     0x29DC,
				Hi:     0x29FB,
				Stride: 1,
			},
			{
				Lo:     0x29FC,
				Hi:     0x29FC,
				Stride: 1,
			},
			{
				Lo:     0x29FD,
				Hi:     0x29FD,
				Stride: 1,
			},
			{
				Lo:     0x29FE,
				Hi:     0x2AFF,
				Stride: 1,
			},
			{
				Lo:     0x2B00,
				Hi:     0x2B2F,
				Stride: 1,
			},
			{
				Lo:     0x2B30,
				Hi:     0x2B44,
				Stride: 1,
			},
			{
				Lo:     0x2B45,
				Hi:     0x2B46,
				Stride: 1,
			},
			{
				Lo:     0x2B47,
				Hi:     0x2B4C,
				Stride: 1,
			},
			{
				Lo:     0x2B4D,
				Hi:     0x2B73,
				Stride: 1,
			},
			{
				Lo:     0x2B76,
				Hi:     0x2B95,
				Stride: 1,
			},
			{
				Lo:     0x2B98,
				Hi:     0x2BB9,
				Stride: 1,
			},
			{
				Lo:     0x2BBD,
				Hi:     0x2BC8,
				Stride: 1,
			},
			{
				Lo:     0x2BCA,
				Hi:     0x2BD1,
				Stride: 1,
			},
			{
				Lo:     0x2BEC,
				Hi:     0x2BEF,
				Stride: 1,
			},
			{
				Lo:     0x2CE5,
				Hi:     0x2CEA,
				Stride: 1,
			},
			{
				Lo:     0x2CF9,
				Hi:     0x2CFC,
				Stride: 1,
			},
			{
				Lo:     0x2CFD,
				Hi:     0x2CFD,
				Stride: 1,
			},
			{
				Lo:     0x2CFE,
				Hi:     0x2CFF,
				Stride: 1,
			},
			{
				Lo:     0x2E00,
				Hi:     0x2E01,
				Stride: 1,
			},
			{
				Lo:     0x2E02,
				Hi:     0x2E02,
				Stride: 1,
			},
			{
				Lo:     0x2E03,
				Hi:     0x2E03,
				Stride: 1,
			},
			{
				Lo:     0x2E04,
				Hi:     0x2E04,
				Stride: 1,
			},
			{
				Lo:     0x2E05,
				Hi:     0x2E05,
				Stride: 1,
			},
			{
				Lo:     0x2E06,
				Hi:     0x2E08,
				Stride: 1,
			},
			{
				Lo:     0x2E09,
				Hi:     0x2E09,
				Stride: 1,
			},
			{
				Lo:     0x2E0A,
				Hi:     0x2E0A,
				Stride: 1,
			},
			{
				Lo:     0x2E0B,
				Hi:     0x2E0B,
				Stride: 1,
			},
			{
				Lo:     0x2E0C,
				Hi:     0x2E0C,
				Stride: 1,
			},
			{
				Lo:     0x2E0D,
				Hi:     0x2E0D,
				Stride: 1,
			},
			{
				Lo:     0x2E0E,
				Hi:     0x2E16,
				Stride: 1,
			},
			{
				Lo:     0x2E17,
				Hi:     0x2E17,
				Stride: 1,
			},
			{
				Lo:     0x2E18,
				Hi:     0x2E19,
				Stride: 1,
			},
			{
				Lo:     0x2E1A,
				Hi:     0x2E1A,
				Stride: 1,
			},
			{
				Lo:     0x2E1B,
				Hi:     0x2E1B,
				Stride: 1,
			},
			{
				Lo:     0x2E1C,
				Hi:     0x2E1C,
				Stride: 1,
			},
			{
				Lo:     0x2E1D,
				Hi:     0x2E1D,
				Stride: 1,
			},
			{
				Lo:     0x2E1E,
				Hi:     0x2E1F,
				Stride: 1,
			},
			{
				Lo:     0x2E20,
				Hi:     0x2E20,
				Stride: 1,
			},
			{
				Lo:     0x2E21,
				Hi:     0x2E21,
				Stride: 1,
			},
			{
				Lo:     0x2E22,
				Hi:     0x2E22,
				Stride: 1,
			},
			{
				Lo:     0x2E23,
				Hi:     0x2E23,
				Stride: 1,
			},
			{
				Lo:     0x2E24,
				Hi:     0x2E24,
				Stride: 1,
			},
			{
				Lo:     0x2E25,
				Hi:     0x2E25,
				Stride: 1,
			},
			{
				Lo:     0x2E26,
				Hi:     0x2E26,
				Stride: 1,
			},
			{
				Lo:     0x2E27,
				Hi:     0x2E27,
				Stride: 1,
			},
			{
				Lo:     0x2E28,
				Hi:     0x2E28,
				Stride: 1,
			},
			{
				Lo:     0x2E29,
				Hi:     0x2E29,
				Stride: 1,
			},
			{
				Lo:     0x2E2A,
				Hi:     0x2E2E,
				Stride: 1,
			},
			{
				Lo:     0x2E2F,
				Hi:     0x2E2F,
				Stride: 1,
			},
			{
				Lo:     0x2E30,
				Hi:     0x2E39,
				Stride: 1,
			},
			{
				Lo:     0x2E3A,
				Hi:     0x2E3B,
				Stride: 1,
			},
			{
				Lo:     0x2E3C,
				Hi:     0x2E3F,
				Stride: 1,
			},
			{
				Lo:     0x2E40,
				Hi:     0x2E40,
				Stride: 1,
			},
			{
				Lo:     0x2E41,
				Hi:     0x2E41,
				Stride: 1,
			},
			{
				Lo:     0x2E42,
				Hi:     0x2E42,
				Stride: 1,
			},
			{
				Lo:     0x2E80,
				Hi:     0x2E99,
				Stride: 1,
			},
			{
				Lo:     0x2E9B,
				Hi:     0x2EF3,
				Stride: 1,
			},
			{
				Lo:     0x2F00,
				Hi:     0x2FD5,
				Stride: 1,
			},
			{
				Lo:     0x2FF0,
				Hi:     0x2FFB,
				Stride: 1,
			},
			{
				Lo:     0x3001,
				Hi:     0x3003,
				Stride: 1,
			},
			{
				Lo:     0x3004,
				Hi:     0x3004,
				Stride: 1,
			},
			{
				Lo:     0x3008,
				Hi:     0x3008,
				Stride: 1,
			},
			{
				Lo:     0x3009,
				Hi:     0x3009,
				Stride: 1,
			},
			{
				Lo:     0x300A,
				Hi:     0x300A,
				Stride: 1,
			},
			{
				Lo:     0x300B,
				Hi:     0x300B,
				Stride: 1,
			},
			{
				Lo:     0x300C,
				Hi:     0x300C,
				Stride: 1,
			},
			{
				Lo:     0x300D,
				Hi:     0x300D,
				Stride: 1,
			},
			{
				Lo:     0x300E,
				Hi:     0x300E,
				Stride: 1,
			},
			{
				Lo:     0x300F,
				Hi:     0x300F,
				Stride: 1,
			},
			{
				Lo:     0x3010,
				Hi:     0x3010,
				Stride: 1,
			},
			{
				Lo:     0x3011,
				Hi:     0x3011,
				Stride: 1,
			},
			{
				Lo:     0x3012,
				Hi:     0x3013,
				Stride: 1,
			},
			{
				Lo:     0x3014,
				Hi:     0x3014,
				Stride: 1,
			},
			{
				Lo:     0x3015,
				Hi:     0x3015,
				Stride: 1,
			},
			{
				Lo:     0x3016,
				Hi:     0x3016,
				Stride: 1,
			},
			{
				Lo:     0x3017,
				Hi:     0x3017,
				Stride: 1,
			},
			{
				Lo:     0x3018,
				Hi:     0x3018,
				Stride: 1,
			},
			{
				Lo:     0x3019,
				Hi:     0x3019,
				Stride: 1,
			},
			{
				Lo:     0x301A,
				Hi:     0x301A,
				Stride: 1,
			},
			{
				Lo:     0x301B,
				Hi:     0x301B,
				Stride: 1,
			},
			{
				Lo:     0x301C,
				Hi:     0x301C,
				Stride: 1,
			},
			{
				Lo:     0x301D,
				Hi:     0x301D,
				Stride: 1,
			},
			{
				Lo:     0x301E,
				Hi:     0x301F,
				Stride: 1,
			},
			{
				Lo:     0x3020,
				Hi:     0x3020,
				Stride: 1,
			},
			{
				Lo:     0x3030,
				Hi:     0x3030,
				Stride: 1,
			},
			{
				Lo:     0x3036,
				Hi:     0x3037,
				Stride: 1,
			},
			{
				Lo:     0x303D,
				Hi:     0x303D,
				Stride: 1,
			},
			{
				Lo:     0x303E,
				Hi:     0x303F,
				Stride: 1,
			},
			{
				Lo:     0x309B,
				Hi:     0x309C,
				Stride: 1,
			},
			{
				Lo:     0x30A0,
				Hi:     0x30A0,
				Stride: 1,
			},
			{
				Lo:     0x30FB,
				Hi:     0x30FB,
				Stride: 1,
			},
			{
				Lo:     0x31C0,
				Hi:     0x31E3,
				Stride: 1,
			},
			{
				Lo:     0x321D,
				Hi:     0x321E,
				Stride: 1,
			},
			{
				Lo:     0x3250,
				Hi:     0x3250,
				Stride: 1,
			},
			{
				Lo:     0x3251,
				Hi:     0x325F,
				Stride: 1,
			},
			{
				Lo:     0x327C,
				Hi:     0x327E,
				Stride: 1,
			},
			{
				Lo:     0x32B1,
				Hi:     0x32BF,
				Stride: 1,
			},
			{
				Lo:     0x32CC,
				Hi:     0x32CF,
				Stride: 1,
			},
			{
				Lo:     0x3377,
				Hi:     0x337A,
				Stride: 1,
			},
			{
				Lo:     0x33DE,
				Hi:     0x33DF,
				Stride: 1,
			},
			{
				Lo:     0x33FF,
				Hi:     0x33FF,
				Stride: 1,
			},
			{
				Lo:     0x4DC0,
				Hi:     0x4DFF,
				Stride: 1,
			},
			{
				Lo:     0xA490,
				Hi:     0xA4C6,
				Stride: 1,
			},
			{
				Lo:     0xA60D,
				Hi:     0xA60F,
				Stride: 1,
			},
			{
				Lo:     0xA673,
				Hi:     0xA673,
				Stride: 1,
			},
			{
				Lo:     0xA67E,
				Hi:     0xA67E,
				Stride: 1,
			},
			{
				Lo:     0xA67F,
				Hi:     0xA67F,
				Stride: 1,
			},
			{
				Lo:     0xA700,
				Hi:     0xA716,
				Stride: 1,
			},
			{
				Lo:     0xA717,
				Hi:     0xA71F,
				Stride: 1,
			},
			{
				Lo:     0xA720,
				Hi:     0xA721,
				Stride: 1,
			},
			{
				Lo:     0xA788,
				Hi:     0xA788,
				Stride: 1,
			},
			{
				Lo:     0xA828,
				Hi:     0xA82B,
				Stride: 1,
			},
			{
				Lo:     0xA874,
				Hi:     0xA877,
				Stride: 1,
			},
			{
				Lo:     0xFD3E,
				Hi:     0xFD3E,
				Stride: 1,
			},
			{
				Lo:     0xFD3F,
				Hi:     0xFD3F,
				Stride: 1,
			},
			{
				Lo:     0xFDFD,
				Hi:     0xFDFD,
				Stride: 1,
			},
			{
				Lo:     0xFE10,
				Hi:     0xFE16,
				Stride: 1,
			},
			{
				Lo:     0xFE17,
				Hi:     0xFE17,
				Stride: 1,
			},
			{
				Lo:     0xFE18,
				Hi:     0xFE18,
				Stride: 1,
			},
			{
				Lo:     0xFE19,
				Hi:     0xFE19,
				Stride: 1,
			},
			{
				Lo:     0xFE30,
				Hi:     0xFE30,
				Stride: 1,
			},
			{
				Lo:     0xFE31,
				Hi:     0xFE32,
				Stride: 1,
			},
			{
				Lo:     0xFE33,
				Hi:     0xFE34,
				Stride: 1,
			},
			{
				Lo:     0xFE35,
				Hi:     0xFE35,
				Stride: 1,
			},
			{
				Lo:     0xFE36,
				Hi:     0xFE36,
				Stride: 1,
			},
			{
				Lo:     0xFE37,
				Hi:     0xFE37,
				Stride: 1,
			},
			{
				Lo:     0xFE38,
				Hi:     0xFE38,
				Stride: 1,
			},
			{
				Lo:     0xFE39,
				Hi:     0xFE39,
				Stride: 1,
			},
			{
				Lo:     0xFE3A,
				Hi:     0xFE3A,
				Stride: 1,
			},
			{
				Lo:     0xFE3B,
				Hi:     0xFE3B,
				Stride: 1,
			},
			{
				Lo:     0xFE3C,
				Hi:     0xFE3C,
				Stride: 1,
			},
			{
				Lo:     0xFE3D,
				Hi:     0xFE3D,
				Stride: 1,
			},
			{
				Lo:     0xFE3E,
				Hi:     0xFE3E,
				Stride: 1,
			},
			{
				Lo:     0xFE3F,
				Hi:     0xFE3F,
				Stride: 1,
			},
			{
				Lo:     0xFE40,
				Hi:     0xFE40,
				Stride: 1,
			},
			{
				Lo:     0xFE41,
				Hi:     0xFE41,
				Stride: 1,
			},
			{
				Lo:     0xFE42,
				Hi:     0xFE42,
				Stride: 1,
			},
			{
				Lo:     0xFE43,
				Hi:     0xFE43,
				Stride: 1,
			},
			{
				Lo:     0xFE44,
				Hi:     0xFE44,
				Stride: 1,
			},
			{
				Lo:     0xFE45,
				Hi:     0xFE46,
				Stride: 1,
			},
			{
				Lo:     0xFE47,
				Hi:     0xFE47,
				Stride: 1,
			},
			{
				Lo:     0xFE48,
				Hi:     0xFE48,
				Stride: 1,
			},
			{
				Lo:     0xFE49,
				Hi:     0xFE4C,
				Stride: 1,
			},
			{
				Lo:     0xFE4D,
				Hi:     0xFE4F,
				Stride: 1,
			},
			{
				Lo:     0xFE51,
				Hi:     0xFE51,
				Stride: 1,
			},
			{
				Lo:     0xFE54,
				Hi:     0xFE54,
				Stride: 1,
			},
			{
				Lo:     0xFE56,
				Hi:     0xFE57,
				Stride: 1,
			},
			{
				Lo:     0xFE58,
				Hi:     0xFE58,
				Stride: 1,
			},
			{
				Lo:     0xFE59,
				Hi:     0xFE59,
				Stride: 1,
			},
			{
				Lo:     0xFE5A,
				Hi:     0xFE5A,
				Stride: 1,
			},
			{
				Lo:     0xFE5B,
				Hi:     0xFE5B,
				Stride: 1,
			},
			{
				Lo:     0xFE5C,
				Hi:     0xFE5C,
				Stride: 1,
			},
			{
				Lo:     0xFE5D,
				Hi:     0xFE5D,
				Stride: 1,
			},
			{
				Lo:     0xFE5E,
				Hi:     0xFE5E,
				Stride: 1,
			},
			{
				Lo:     0xFE60,
				Hi:     0xFE61,
				Stride: 1,
			},
			{
				Lo:     0xFE64,
				Hi:     0xFE66,
				Stride: 1,
			},
			{
				Lo:     0xFE68,
				Hi:     0xFE68,
				Stride: 1,
			},
			{
				Lo:     0xFE6B,
				Hi:     0xFE6B,
				Stride: 1,
			},
			{
				Lo:     0xFF01,
				Hi:     0xFF02,
				Stride: 1,
			},
			{
				Lo:     0xFF06,
				Hi:     0xFF07,
				Stride: 1,
			},
			{
				Lo:     0xFF08,
				Hi:     0xFF08,
				Stride: 1,
			},
			{
				Lo:     0xFF09,
				Hi:     0xFF09,
				Stride: 1,
			},
			{
				Lo:     0xFF0A,
				Hi:     0xFF0A,
				Stride: 1,
			},
			{
				Lo:     0xFF1B,
				Hi:     0xFF1B,
				Stride: 1,
			},
			{
				Lo:     0xFF1C,
				Hi:     0xFF1E,
				Stride: 1,
			},
			{
				Lo:     0xFF1F,
				Hi:     0xFF20,
				Stride: 1,
			},
			{
				Lo:     0xFF3B,
				Hi:     0xFF3B,
				Stride: 1,
			},
			{
				Lo:     0xFF3C,
				Hi:     0xFF3C,
				Stride: 1,
			},
			{
				Lo:     0xFF3D,
				Hi:     0xFF3D,
				Stride: 1,
			},
			{
				Lo:     0xFF3E,
				Hi:     0xFF3E,
				Stride: 1,
			},
			{
				Lo:     0xFF3F,
				Hi:     0xFF3F,
				Stride: 1,
			},
			{
				Lo:     0xFF40,
				Hi:     0xFF40,
				Stride: 1,
			},
			{
				Lo:     0xFF5B,
				Hi:     0xFF5B,
				Stride: 1,
			},
			{
				Lo:     0xFF5C,
				Hi:     0xFF5C,
				Stride: 1,
			},
			{
				Lo:     0xFF5D,
				Hi:     0xFF5D,
				Stride: 1,
			},
			{
				Lo:     0xFF5E,
				Hi:     0xFF5E,
				Stride: 1,
			},
			{
				Lo:     0xFF5F,
				Hi:     0xFF5F,
				Stride: 1,
			},
			{
				Lo:     0xFF60,
				Hi:     0xFF60,
				Stride: 1,
			},
			{
				Lo:     0xFF61,
				Hi:     0xFF61,
				Stride: 1,
			},
			{
				Lo:     0xFF62,
				Hi:     0xFF62,
				Stride: 1,
			},
			{
				Lo:     0xFF63,
				Hi:     0xFF63,
				Stride: 1,
			},
			{
				Lo:     0xFF64,
				Hi:     0xFF65,
				Stride: 1,
			},
			{
				Lo:     0xFFE2,
				Hi:     0xFFE2,
				Stride: 1,
			},
			{
				Lo:     0xFFE3,
				Hi:     0xFFE3,
				Stride: 1,
			},
			{
				Lo:     0xFFE4,
				Hi:     0xFFE4,
				Stride: 1,
			},
			{
				Lo:     0xFFE8,
				Hi:     0xFFE8,
				Stride: 1,
			},
			{
				Lo:     0xFFE9,
				Hi:     0xFFEC,
				Stride: 1,
			},
			{
				Lo:     0xFFED,
				Hi:     0xFFEE,
				Stride: 1,
			},
			{
				Lo:     0xFFF9,
				Hi:     0xFFFB,
				Stride: 1,
			},
			{
				Lo:     0xFFFC,
				Hi:     0xFFFD,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x10101,
				Hi:     0x10101,
				Stride: 1,
			},
			{
				Lo:     0x10140,
				Hi:     0x10174,
				Stride: 1,
			},
			{
				Lo:     0x10175,
				Hi:     0x10178,
				Stride: 1,
			},
			{
				Lo:     0x10179,
				Hi:     0x10189,
				Stride: 1,
			},
			{
				Lo:     0x1018A,
				Hi:     0x1018B,
				Stride: 1,
			},
			{
				Lo:     0x1018C,
				Hi:     0x1018C,
				Stride: 1,
			},
			{
				Lo:     0x10190,
				Hi:     0x1019B,
				Stride: 1,
			},
			{
				Lo:     0x101A0,
				Hi:     0x101A0,
				Stride: 1,
			},
			{
				Lo:     0x1091F,
				Hi:     0x1091F,
				Stride: 1,
			},
			{
				Lo:     0x10B39,
				Hi:     0x10B3F,
				Stride: 1,
			},
			{
				Lo:     0x11052,
				Hi:     0x11065,
				Stride: 1,
			},
			{
				Lo:     0x1D200,
				Hi:     0x1D241,
				Stride: 1,
			},
			{
				Lo:     0x1D245,
				Hi:     0x1D245,
				Stride: 1,
			},
			{
				Lo:     0x1D300,
				Hi:     0x1D356,
				Stride: 1,
			},
			{
				Lo:     0x1D6DB,
				Hi:     0x1D6DB,
				Stride: 1,
			},
			{
				Lo:     0x1D715,
				Hi:     0x1D715,
				Stride: 1,
			},
			{
				Lo:     0x1D74F,
				Hi:     0x1D74F,
				Stride: 1,
			},
			{
				Lo:     0x1D789,
				Hi:     0x1D789,
				Stride: 1,
			},
			{
				Lo:     0x1D7C3,
				Hi:     0x1D7C3,
				Stride: 1,
			},
			{
				Lo:     0x1EEF0,
				Hi:     0x1EEF1,
				Stride: 1,
			},
			{
				Lo:     0x1F000,
				Hi:     0x1F02B,
				Stride: 1,
			},
			{
				Lo:     0x1F030,
				Hi:     0x1F093,
				Stride: 1,
			},
			{
				Lo:     0x1F0A0,
				Hi:     0x1F0AE,
				Stride: 1,
			},
			{
				Lo:     0x1F0B1,
				Hi:     0x1F0BF,
				Stride: 1,
			},
			{
				Lo:     0x1F0C1,
				Hi:     0x1F0CF,
				Stride: 1,
			},
			{
				Lo:     0x1F0D1,
				Hi:     0x1F0F5,
				Stride: 1,
			},
			{
				Lo:     0x1F10B,
				Hi:     0x1F10C,
				Stride: 1,
			},
			{
				Lo:     0x1F16A,
				Hi:     0x1F16B,
				Stride: 1,
			},
			{
				Lo:     0x1F300,
				Hi:     0x1F3FA,
				Stride: 1,
			},
			{
				Lo:     0x1F3FB,
				Hi:     0x1F3FF,
				Stride: 1,
			},
			{
				Lo:     0x1F400,
				Hi:     0x1F579,
				Stride: 1,
			},
			{
				Lo:     0x1F57B,
				Hi:     0x1F5A3,
				Stride: 1,
			},
			{
				Lo:     0x1F5A5,
				Hi:     0x1F6D0,
				Stride: 1,
			},
			{
				Lo:     0x1F6E0,
				Hi:     0x1F6EC,
				Stride: 1,
			},
			{
				Lo:     0x1F6F0,
				Hi:     0x1F6F3,
				Stride: 1,
			},
			{
				Lo:     0x1F700,
				Hi:     0x1F773,
				Stride: 1,
			},
			{
				Lo:     0x1F780,
				Hi:     0x1F7D4,
				Stride: 1,
			},
			{
				Lo:     0x1F800,
				Hi:     0x1F80B,
				Stride: 1,
			},
			{
				Lo:     0x1F810,
				Hi:     0x1F847,
				Stride: 1,
			},
			{
				Lo:     0x1F850,
				Hi:     0x1F859,
				Stride: 1,
			},
			{
				Lo:     0x1F860,
				Hi:     0x1F887,
				Stride: 1,
			},
			{
				Lo:     0x1F890,
				Hi:     0x1F8AD,
				Stride: 1,
			},
			{
				Lo:     0x1F910,
				Hi:     0x1F918,
				Stride: 1,
			},
			{
				Lo:     0x1F980,
				Hi:     0x1F984,
				Stride: 1,
			},
			{
				Lo:     0x1F9C0,
				Hi:     0x1F9C0,
				Stride: 1,
			},
		},
	}

	bidiClassPDF = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x202C,
				Hi:     0x202C,
				Stride: 1,
			},
		},
	}

	bidiClassPDI = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x2069,
				Hi:     0x2069,
				Stride: 1,
			},
		},
	}

	bidiClassR = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0590,
				Hi:     0x0590,
				Stride: 1,
			},
			{
				Lo:     0x05BE,
				Hi:     0x05BE,
				Stride: 1,
			},
			{
				Lo:     0x05C0,
				Hi:     0x05C0,
				Stride: 1,
			},
			{
				Lo:     0x05C3,
				Hi:     0x05C3,
				Stride: 1,
			},
			{
				Lo:     0x05C6,
				Hi:     0x05C6,
				Stride: 1,
			},
			{
				Lo:     0x05C8,
				Hi:     0x05CF,
				Stride: 1,
			},
			{
				Lo:     0x05D0,
				Hi:     0x05EA,
				Stride: 1,
			},
			{
				Lo:     0x05EB,
				Hi:     0x05EF,
				Stride: 1,
			},
			{
				Lo:     0x05F0,
				Hi:     0x05F2,
				Stride: 1,
			},
			{
				Lo:     0x05F3,
				Hi:     0x05F4,
				Stride: 1,
			},
			{
				Lo:     0x05F5,
				Hi:     0x05FF,
				Stride: 1,
			},
			{
				Lo:     0x07C0,
				Hi:     0x07C9,
				Stride: 1,
			},
			{
				Lo:     0x07CA,
				Hi:     0x07EA,
				Stride: 1,
			},
			{
				Lo:     0x07F4,
				Hi:     0x07F5,
				Stride: 1,
			},
			{
				Lo:     0x07FA,
				Hi:     0x07FA,
				Stride: 1,
			},
			{
				Lo:     0x07FB,
				Hi:     0x07FF,
				Stride: 1,
			},
			{
				Lo:     0x0800,
				Hi:     0x0815,
				Stride: 1,
			},
			{
				Lo:     0x081A,
				Hi:     0x081A,
				Stride: 1,
			},
			{
				Lo:     0x0824,
				Hi:     0x0824,
				Stride: 1,
			},
			{
				Lo:     0x0828,
				Hi:     0x0828,
				Stride: 1,
			},
			{
				Lo:     0x082E,
				Hi:     0x082F,
				Stride: 1,
			},
			{
				Lo:     0x0830,
				Hi:     0x083E,
				Stride: 1,
			},
			{
				Lo:     0x083F,
				Hi:     0x083F,
				Stride: 1,
			},
			{
				Lo:     0x0840,
				Hi:     0x0858,
				Stride: 1,
			},
			{
				Lo:     0x085C,
				Hi:     0x085D,
				Stride: 1,
			},
			{
				Lo:     0x085E,
				Hi:     0x085E,
				Stride: 1,
			},
			{
				Lo:     0x085F,
				Hi:     0x089F,
				Stride: 1,
			},
			{
				Lo:     0x200F,
				Hi:     0x200F,
				Stride: 1,
			},
			{
				Lo:     0xFB1D,
				Hi:     0xFB1D,
				Stride: 1,
			},
			{
				Lo:     0xFB1F,
				Hi:     0xFB28,
				Stride: 1,
			},
			{
				Lo:     0xFB2A,
				Hi:     0xFB36,
				Stride: 1,
			},
			{
				Lo:     0xFB37,
				Hi:     0xFB37,
				Stride: 1,
			},
			{
				Lo:     0xFB38,
				Hi:     0xFB3C,
				Stride: 1,
			},
			{
				Lo:     0xFB3D,
				Hi:     0xFB3D,
				Stride: 1,
			},
			{
				Lo:     0xFB3E,
				Hi:     0xFB3E,
				Stride: 1,
			},
			{
				Lo:     0xFB3F,
				Hi:     0xFB3F,
				Stride: 1,
			},
			{
				Lo:     0xFB40,
				Hi:     0xFB41,
				Stride: 1,
			},
			{
				Lo:     0xFB42,
				Hi:     0xFB42,
				Stride: 1,
			},
			{
				Lo:     0xFB43,
				Hi:     0xFB44,
				Stride: 1,
			},
			{
				Lo:     0xFB45,
				Hi:     0xFB45,
				Stride: 1,
			},
			{
				Lo:     0xFB46,
				Hi:     0xFB4F,
				Stride: 1,
			},
		},
		R32: []unicode.Range32{
			{
				Lo:     0x10800,
				Hi:     0x10805,
				Stride: 1,
			},
			{
				Lo:     0x10806,
				Hi:     0x10807,
				Stride: 1,
			},
			{
				Lo:     0x10808,
				Hi:     0x10808,
				Stride: 1,
			},
			{
				Lo:     0x10809,
				Hi:     0x10809,
				Stride: 1,
			},
			{
				Lo:     0x1080A,
				Hi:     0x10835,
				Stride: 1,
			},
			{
				Lo:     0x10836,
				Hi:     0x10836,
				Stride: 1,
			},
			{
				Lo:     0x10837,
				Hi:     0x10838,
				Stride: 1,
			},
			{
				Lo:     0x10839,
				Hi:     0x1083B,
				Stride: 1,
			},
			{
				Lo:     0x1083C,
				Hi:     0x1083C,
				Stride: 1,
			},
			{
				Lo:     0x1083D,
				Hi:     0x1083E,
				Stride: 1,
			},
			{
				Lo:     0x1083F,
				Hi:     0x10855,
				Stride: 1,
			},
			{
				Lo:     0x10856,
				Hi:     0x10856,
				Stride: 1,
			},
			{
				Lo:     0x10857,
				Hi:     0x10857,
				Stride: 1,
			},
			{
				Lo:     0x10858,
				Hi:     0x1085F,
				Stride: 1,
			},
			{
				Lo:     0x10860,
				Hi:     0x10876,
				Stride: 1,
			},
			{
				Lo:     0x10877,
				Hi:     0x10878,
				Stride: 1,
			},
			{
				Lo:     0x10879,
				Hi:     0x1087F,
				Stride: 1,
			},
			{
				Lo:     0x10880,
				Hi:     0x1089E,
				Stride: 1,
			},
			{
				Lo:     0x1089F,
				Hi:     0x108A6,
				Stride: 1,
			},
			{
				Lo:     0x108A7,
				Hi:     0x108AF,
				Stride: 1,
			},
			{
				Lo:     0x108B0,
				Hi:     0x108DF,
				Stride: 1,
			},
			{
				Lo:     0x108E0,
				Hi:     0x108F2,
				Stride: 1,
			},
			{
				Lo:     0x108F3,
				Hi:     0x108F3,
				Stride: 1,
			},
			{
				Lo:     0x108F4,
				Hi:     0x108F5,
				Stride: 1,
			},
			{
				Lo:     0x108F6,
				Hi:     0x108FA,
				Stride: 1,
			},
			{
				Lo:     0x108FB,
				Hi:     0x108FF,
				Stride: 1,
			},
			{
				Lo:     0x10900,
				Hi:     0x10915,
				Stride: 1,
			},
			{
				Lo:     0x10916,
				Hi:     0x1091B,
				Stride: 1,
			},
			{
				Lo:     0x1091C,
				Hi:     0x1091E,
				Stride: 1,
			},
			{
				Lo:     0x10920,
				Hi:     0x10939,
				Stride: 1,
			},
			{
				Lo:     0x1093A,
				Hi:     0x1093E,
				Stride: 1,
			},
			{
				Lo:     0x1093F,
				Hi:     0x1093F,
				Stride: 1,
			},
			{
				Lo:     0x10940,
				Hi:     0x1097F,
				Stride: 1,
			},
			{
				Lo:     0x10980,
				Hi:     0x109B7,
				Stride: 1,
			},
			{
				Lo:     0x109B8,
				Hi:     0x109BB,
				Stride: 1,
			},
			{
				Lo:     0x109BC,
				Hi:     0x109BD,
				Stride: 1,
			},
			{
				Lo:     0x109BE,
				Hi:     0x109BF,
				Stride: 1,
			},
			{
				Lo:     0x109C0,
				Hi:     0x109CF,
				Stride: 1,
			},
			{
				Lo:     0x109D0,
				Hi:     0x109D1,
				Stride: 1,
			},
			{
				Lo:     0x109D2,
				Hi:     0x109FF,
				Stride: 1,
			},
			{
				Lo:     0x10A00,
				Hi:     0x10A00,
				Stride: 1,
			},
			{
				Lo:     0x10A04,
				Hi:     0x10A04,
				Stride: 1,
			},
			{
				Lo:     0x10A07,
				Hi:     0x10A0B,
				Stride: 1,
			},
			{
				Lo:     0x10A10,
				Hi:     0x10A13,
				Stride: 1,
			},
			{
				Lo:     0x10A14,
				Hi:     0x10A14,
				Stride: 1,
			},
			{
				Lo:     0x10A15,
				Hi:     0x10A17,
				Stride: 1,
			},
			{
				Lo:     0x10A18,
				Hi:     0x10A18,
				Stride: 1,
			},
			{
				Lo:     0x10A19,
				Hi:     0x10A33,
				Stride: 1,
			},
			{
				Lo:     0x10A34,
				Hi:     0x10A37,
				Stride: 1,
			},
			{
				Lo:     0x10A3B,
				Hi:     0x10A3E,
				Stride: 1,
			},
			{
				Lo:     0x10A40,
				Hi:     0x10A47,
				Stride: 1,
			},
			{
				Lo:     0x10A48,
				Hi:     0x10A4F,
				Stride: 1,
			},
			{
				Lo:     0x10A50,
				Hi:     0x10A58,
				Stride: 1,
			},
			{
				Lo:     0x10A59,
				Hi:     0x10A5F,
				Stride: 1,
			},
			{
				Lo:     0x10A60,
				Hi:     0x10A7C,
				Stride: 1,
			},
			{
				Lo:     0x10A7D,
				Hi:     0x10A7E,
				Stride: 1,
			},
			{
				Lo:     0x10A7F,
				Hi:     0x10A7F,
				Stride: 1,
			},
			{
				Lo:     0x10A80,
				Hi:     0x10A9C,
				Stride: 1,
			},
			{
				Lo:     0x10A9D,
				Hi:     0x10A9F,
				Stride: 1,
			},
			{
				Lo:     0x10AA0,
				Hi:     0x10ABF,
				Stride: 1,
			},
			{
				Lo:     0x10AC0,
				Hi:     0x10AC7,
				Stride: 1,
			},
			{
				Lo:     0x10AC8,
				Hi:     0x10AC8,
				Stride: 1,
			},
			{
				Lo:     0x10AC9,
				Hi:     0x10AE4,
				Stride: 1,
			},
			{
				Lo:     0x10AE7,
				Hi:     0x10AEA,
				Stride: 1,
			},
			{
				Lo:     0x10AEB,
				Hi:     0x10AEF,
				Stride: 1,
			},
			{
				Lo:     0x10AF0,
				Hi:     0x10AF6,
				Stride: 1,
			},
			{
				Lo:     0x10AF7,
				Hi:     0x10AFF,
				Stride: 1,
			},
			{
				Lo:     0x10B00,
				Hi:     0x10B35,
				Stride: 1,
			},
			{
				Lo:     0x10B36,
				Hi:     0x10B38,
				Stride: 1,
			},
			{
				Lo:     0x10B40,
				Hi:     0x10B55,
				Stride: 1,
			},
			{
				Lo:     0x10B56,
				Hi:     0x10B57,
				Stride: 1,
			},
			{
				Lo:     0x10B58,
				Hi:     0x10B5F,
				Stride: 1,
			},
			{
				Lo:     0x10B60,
				Hi:     0x10B72,
				Stride: 1,
			},
			{
				Lo:     0x10B73,
				Hi:     0x10B77,
				Stride: 1,
			},
			{
				Lo:     0x10B78,
				Hi:     0x10B7F,
				Stride: 1,
			},
			{
				Lo:     0x10B80,
				Hi:     0x10B91,
				Stride: 1,
			},
			{
				Lo:     0x10B92,
				Hi:     0x10B98,
				Stride: 1,
			},
			{
				Lo:     0x10B99,
				Hi:     0x10B9C,
				Stride: 1,
			},
			{
				Lo:     0x10B9D,
				Hi:     0x10BA8,
				Stride: 1,
			},
			{
				Lo:     0x10BA9,
				Hi:     0x10BAF,
				Stride: 1,
			},
			{
				Lo:     0x10BB0,
				Hi:     0x10BFF,
				Stride: 1,
			},
			{
				Lo:     0x10C00,
				Hi:     0x10C48,
				Stride: 1,
			},
			{
				Lo:     0x10C49,
				Hi:     0x10C7F,
				Stride: 1,
			},
			{
				Lo:     0x10C80,
				Hi:     0x10CB2,
				Stride: 1,
			},
			{
				Lo:     0x10CB3,
				Hi:     0x10CBF,
				Stride: 1,
			},
			{
				Lo:     0x10CC0,
				Hi:     0x10CF2,
				Stride: 1,
			},
			{
				Lo:     0x10CF3,
				Hi:     0x10CF9,
				Stride: 1,
			},
			{
				Lo:     0x10CFA,
				Hi:     0x10CFF,
				Stride: 1,
			},
			{
				Lo:     0x10D00,
				Hi:     0x10E5F,
				Stride: 1,
			},
			{
				Lo:     0x10E7F,
				Hi:     0x10FFF,
				Stride: 1,
			},
			{
				Lo:     0x1E800,
				Hi:     0x1E8C4,
				Stride: 1,
			},
			{
				Lo:     0x1E8C5,
				Hi:     0x1E8C6,
				Stride: 1,
			},
			{
				Lo:     0x1E8C7,
				Hi:     0x1E8CF,
				Stride: 1,
			},
			{
				Lo:     0x1E8D7,
				Hi:     0x1EDFF,
				Stride: 1,
			},
			{
				Lo:     0x1EF00,
				Hi:     0x1EFFF,
				Stride: 1,
			},
		},
	}

	bidiClassRLE = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x202B,
				Hi:     0x202B,
				Stride: 1,
			},
		},
	}

	bidiClassRLI = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x2067,
				Hi:     0x2067,
				Stride: 1,
			},
		},
	}

	bidiClassRLO = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x202E,
				Hi:     0x202E,
				Stride: 1,
			},
		},
	}

	bidiClassS = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x0009,
				Hi:     0x0009,
				Stride: 1,
			},
			{
				Lo:     0x000B,
				Hi:     0x000B,
				Stride: 1,
			},
			{
				Lo:     0x001F,
				Hi:     0x001F,
				Stride: 1,
			},
		},
	}

	bidiClassWS = unicode.RangeTable{
		R16: []unicode.Range16{
			{
				Lo:     0x000C,
				Hi:     0x000C,
				Stride: 1,
			},
			{
				Lo:     0x0020,
				Hi:     0x0020,
				Stride: 1,
			},
			{
				Lo:     0x1680,
				Hi:     0x1680,
				Stride: 1,
			},
			{
				Lo:     0x2000,
				Hi:     0x200A,
				Stride: 1,
			},
			{
				Lo:     0x2028,
				Hi:     0x2028,
				Stride: 1,
			},
			{
				Lo:     0x205F,
				Hi:     0x205F,
				Stride: 1,
			},
			{
				Lo:     0x3000,
				Hi:     0x3000,
				Stride: 1,
			},
		},
	}
)
