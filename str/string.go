package str

import (
	"fmt"
	"regexp"
	"strings"
)

// Compare

func CompareCaseInsensitive(input, compareTo string) (bool, string) {
	if strings.ToLower(input) == strings.ToLower(compareTo) {
		return true, compareTo
	}

	return false, input
}

// Padding

func PrefixAndZeroPad(original interface{}, prefix string, count int, includePrefixInPadLength bool) string {
	padLength := count
	if includePrefixInPadLength {
		padLength = count - len(prefix)
	}

	return fmt.Sprintf("%s%s", prefix, LeftPad(original, "0", padLength))
}

func ZeroPad(original interface{}, count int) string {
	return LeftPad(original, "0", count)
}

func LeftPad(original interface{}, padder string, count int) string {
	padFunc := func(padding, original string) (string, string) { return padding, original }
	return pad(original, padder, padFunc, count)
}

func RightPad(original interface{}, padder string, count int) string {
	padFunc := func(padding, original string) (string, string) { return original, padding }
	return pad(original, padder, padFunc, count)
}

func pad(original interface{}, padder string, orderFunc func(string, string) (string, string), count int) string {
	originalStr := fmt.Sprintf("%v", original)
	if count <= 0 {
		return originalStr
	}

	var needed int
	needed = count - len(originalStr)
	if needed <= 0 {
		return originalStr
	}

	padding := strings.Repeat(padder, needed)
	left, right := orderFunc(padding, originalStr)
	return fmt.Sprintf("%s%s", left, right)
}

// Names

func FirstNameWithDefault(s, def string) string {
	if s == "" {
		return def
	}

	names := strings.Split(strings.TrimSpace(s), " ")
	if len(names) == 0 {
		return def
	}

	return names[0]
}

func MustExtractTwoInitialsFromName(name string) string {
	// Clean up whitespace before splitting
	name = strings.TrimSpace(name)

	initials := ""

	if len(name) == 0 {
		initials = "XX"
	} else if len(name) == 1 {
		initials = fmt.Sprintf("%sX", name)
	} else {
		initials = name[:2]
	}

	// Split the name into words
	words := strings.Fields(name)

	// Extract the first letter of each word
	if len(words) >= 2 {
		initials = string(rune(words[0][0])) + string(rune(words[1][0]))
	}

	return strings.ToUpper(initials)
}

// Replace

func ReplaceAllWithHyphen(s string, toReplace ...string) string {
	return ReplaceAllWith(s, "-", toReplace...)
}

func RemoveAll(s string, toReplace ...string) string {
	return ReplaceAllWith(s, "", toReplace...)
}

func RemoveAllRegexSymbols(s string) string {
	return RemoveAll(s, "^", ".", "?", "{", "}", "(", ")", "[", "]", "-", "$", "|", "+", "*", `"`, "/", `\`)
}

func ReplaceAllWith(s string, replacement string, toReplace ...string) string {
	if len(toReplace) == 0 {
		return s
	}

	for _, r := range toReplace {
		s = strings.ReplaceAll(s, r, replacement)
	}

	return s
}

func ReplaceSpacesWithCommas(s string) string {
	return strings.ReplaceAll(s, " ", ",")
}

// IterativeCutAbove separates a string at a point specified by a separator,
// returns only the portion above, then feeds that in again
func IterativeCutAbove(s string, seps ...string) string {
	res := s

	for _, sep := range seps {
		above, _, _ := strings.Cut(res, sep)
		res = above
	}

	return res
}

func CutAtRegex(s string, re *regexp.Regexp) (string, string, bool) {
	res := re.FindStringIndex(s)
	if res == nil {
		return s, "", false
	}

	before := s[:res[0]]
	after := s[res[1]:]

	return before, after, true
}

func CutAtAnyRegex(s string, res ...*regexp.Regexp) (string, string, bool) {
	for _, re := range res {
		if before, after, ok := CutAtRegex(s, re); ok {
			return before, after, true
		}
	}

	return s, "", false
}
