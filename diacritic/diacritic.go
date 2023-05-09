package diacritic

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var Equivalences = map[string]string{
	// Lower
	"o": "oòóôõö",
	"a": "aàáâãäåã",
	"e": "eèéêë",
	"i": "iìíîï",
	"u": "uùúûü",
	"n": "nñ",
	"c": "cç",
	"s": "sš",
	"z": "zž",
	"y": "yÿ",
	// Upper
	"O": "OÒÓÔÕÖ",
	"A": "AÀÁÂÃÄÅÃ",
	"E": "EÈÉÊË",
	"I": "IÌÍÎÏ",
	"U": "UÙÚÛÜ",
	"N": "NÑ",
	"C": "CÇ",
	"S": "SŠ",
	"Z": "ZŽ",
	"Y": "YŸ",
}

func RemoveDiacritics(input string) string {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	result, _, _ := transform.String(t, input)
	return result
}

func PrepareForRegex(input string) (res string) {
	// Removing the diacritics completely
	// That way we only have to specify one-way equivalences
	input = RemoveDiacritics(input)

	for _, c := range strings.Split(input, "") {
		trans := c
		eq, ok := Equivalences[c]
		if ok {
			trans = fmt.Sprintf("[%s]", eq)
		}

		res = fmt.Sprintf("%s%s", res, trans)
	}

	return
}
