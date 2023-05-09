package diacritic

import "testing"

func TestPrepareForRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"single letter", "o", "[oòóôõö]"},
		{"two letters", "oo", "[oòóôõö][oòóôõö]"},
		{"two letters, interspersed", "oko", "[oòóôõö]k[oòóôõö]"},
		{"single diacritic letter", "õ", "[oòóôõö]"},
		{"two diacritic letters", "òó", "[oòóôõö][oòóôõö]"},
		{"full word", "Ñòt", "[NÑ][oòóôõö]t"},
	}

	for _, test := range tests {
		res := PrepareForRegex(test.input)
		if res != test.expected {
			t.Errorf("Testing %s, got unexpected result. Expected %s; got %s",
				test.name,
				test.expected,
				res,
			)
		}
	}
}
