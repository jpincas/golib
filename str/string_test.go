package str

import (
	"regexp"
	"testing"
)

func TestZeroPad(t *testing.T) {
	expected := "00000999"

	if res := ZeroPad("999", 8); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestZeroPad0(t *testing.T) {
	expected := "999"

	if res := ZeroPad("999", 0); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_ExcludePrefixFromCount(t *testing.T) {
	expected := "PRE00000999"

	if res := PrefixAndZeroPad("999", "PRE", 8, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_IncludePrefixInCount(t *testing.T) {
	expected := "PRE00999"

	if res := PrefixAndZeroPad("999", "PRE", 8, true); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad0(t *testing.T) {
	expected := "PRE999"

	if res := PrefixAndZeroPad("999", "PRE", 0, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_TooShort(t *testing.T) {
	// With the specified count set to 5, and 'include' set to true,
	// there are not enough characters to fully express the prefix + ref combo
	// so default to the original
	expected := "PRE999"

	if res := PrefixAndZeroPad("999", "PRE", 5, true); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_NoPrefix(t *testing.T) {
	expected := "00000999"

	if res := PrefixAndZeroPad("999", "", 8, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestPrefixAndZeroPad_0NoPrefix(t *testing.T) {
	expected := "999"

	// shouldn't make any difference what we pass for 'include'
	if res := PrefixAndZeroPad("999", "", 0, true); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}

	if res := PrefixAndZeroPad("999", "", 0, false); res != expected {
		t.Errorf("Expected %s; got %s", expected, res)
	}
}

func TestReplaceAllWithHyphen(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		toReplace []string
		expected  string
	}{
		{"blank", "", []string{}, ""},
		{"blank, try replace space", "", []string{" "}, ""},
		{"space, try replace space", " ", []string{" "}, "-"},
		{"simple, try replace space", "word1 word2", []string{" "}, "word1-word2"},
		{"multi replace", "word1 word2(word3)", []string{" ", "(", ")"}, "word1-word2-word3-"},
	}

	for _, testCase := range testCases {
		res := ReplaceAllWithHyphen(testCase.input, testCase.toReplace...)
		if res != testCase.expected {
			t.Errorf("Testing %s. Expected %s, got %s", testCase.name, testCase.expected, res)
		}
	}

}

func TestExtractInitials(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name     string
		expected string
	}{
		{"John Smith", "JS"},
		{"John", "JO"},
		{"John Michael Smith", "JM"},
		{"John M", "JM"},
		{"J Michael", "JM"},
		{"J ", "JX"}, // has a space,
		{"", "XX"},
		{"a", "AX"},
		{"ab", "AB"},
	}

	// Run the tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initials := MustExtractTwoInitialsFromName(tc.name)
			if initials != tc.expected {
				t.Errorf("For name %q, expected initials %q but got %q", tc.name, tc.expected, initials)
			}
		})
	}
}

func TestCutAtRegex(t *testing.T) {
	// Taking advantage to test my email reply stripper regex
	re_en := regexp.MustCompile("(?s)On.*wrote:")
	re_es := regexp.MustCompile("(?s)El.*escribió:")

	testCases := []struct {
		name                          string
		regex                         *regexp.Regexp
		input                         string
		expectedBefore, expectedAfter string
		expectedOK                    bool
	}{
		{
			name:           "blank",
			regex:          re_en,
			input:          "",
			expectedBefore: "",
			expectedAfter:  "",
			expectedOK:     false,
		},
		{
			name:           "no match",
			regex:          re_en,
			input:          "no match",
			expectedBefore: "no match",
			expectedAfter:  "",
			expectedOK:     false,
		},
		{
			name:           "similar to regex (no colon) but no match",
			regex:          re_en,
			input:          "On 27th Jan I wrote",
			expectedBefore: "On 27th Jan I wrote",
			expectedAfter:  "",
			expectedOK:     false,
		},
		{
			name:           "only the regex",
			regex:          re_en,
			input:          "On 27 Jan 2023, at 10:14, Jon (Pakk) <jon@pakk.io> wrote:",
			expectedBefore: "",
			expectedAfter:  "",
			expectedOK:     true,
		},
		{
			name:           "text before and after the regex",
			regex:          re_en,
			input:          "BEFOREOn 27 Jan 2023, at 10:14, Jon (Pakk) <jon@pakk.io> wrote:AFTER",
			expectedBefore: "BEFORE",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
		{
			name:           "text before the regex",
			regex:          re_en,
			input:          "BEFOREOn 27 Jan 2023, at 10:14, Jon (Pakk) <jon@pakk.io> wrote:",
			expectedBefore: "BEFORE",
			expectedAfter:  "",
			expectedOK:     true,
		},
		{
			name:           "text after the regex",
			regex:          re_en,
			input:          "On 27 Jan 2023, at 10:14, Jon (Pakk) <jon@pakk.io> wrote:AFTER",
			expectedBefore: "",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
		{
			name:  "text before and after the regex, with line break",
			regex: re_en,
			input: `BEFOREOn 27 Jan 2023, at 10:14, 
			Jon (Pakk) <jon@pakk.io> wrote:AFTER`,
			expectedBefore: "BEFORE",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
		{
			name:           "spanish regex",
			regex:          re_es,
			input:          `BEFOREEl jue., 26 ene. 2023 19:42, Dev Ticket System mail@pakk.store escribió:AFTER`,
			expectedBefore: "BEFORE",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
		{
			name:  "spanish regex, with line break",
			regex: re_es,
			input: `BEFOREEl jue., 26 ene. 2023 19:42,
			Dev Ticket System mail@pakk.store escribió:AFTER`,
			expectedBefore: "BEFORE",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
	}

	for _, tc := range testCases {
		before, after, ok := CutAtRegex(tc.input, tc.regex)
		if tc.expectedBefore != before {
			t.Errorf("Testing %s. Expected before: %s; got %s", tc.name, tc.expectedBefore, before)
		}
		if tc.expectedAfter != after {
			t.Errorf("Testing %s. Expected before: %s; got %s", tc.name, tc.expectedAfter, after)
		}
		if tc.expectedOK != ok {
			t.Errorf("Testing %s. Expected ok: %v; got %v", tc.name, tc.expectedOK, ok)
		}
	}

}

func TestCutAtAnyRegex(t *testing.T) {
	// Taking advantage to test my email reply stripper regex
	re_en := regexp.MustCompile("(?s)On.*wrote:")
	re_es := regexp.MustCompile("(?s)El.*escribió:")

	testCases := []struct {
		name                          string
		input                         string
		expectedBefore, expectedAfter string
		expectedOK                    bool
	}{
		{
			name:           "english regex",
			input:          "BEFOREOn 27 Jan 2023, at 10:14, Jon (Pakk) <jon@pakk.io> wrote:AFTER",
			expectedBefore: "BEFORE",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
		{
			name:           "spanish regex",
			input:          `BEFOREEl jue., 26 ene. 2023 19:42, Dev Ticket System mail@pakk.store escribió:AFTER`,
			expectedBefore: "BEFORE",
			expectedAfter:  "AFTER",
			expectedOK:     true,
		},
	}

	for _, tc := range testCases {
		before, after, ok := CutAtAnyRegex(tc.input, re_en, re_es)
		if tc.expectedBefore != before {
			t.Errorf("Testing %s. Expected before: %s; got %s", tc.name, tc.expectedBefore, before)
		}
		if tc.expectedAfter != after {
			t.Errorf("Testing %s. Expected before: %s; got %s", tc.name, tc.expectedAfter, after)
		}
		if tc.expectedOK != ok {
			t.Errorf("Testing %s. Expected ok: %v; got %v", tc.name, tc.expectedOK, ok)
		}
	}
}
