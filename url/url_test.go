package url

import "testing"

func TestStripPort(t *testing.T) {
	tests := []struct {
		src      string
		expected string
	}{
		{"", ""},
		{"123.456.789.10", "123.456.789.10"},
		{"123.456.789.10:1234", "123.456.789.10"},
		{"mydomain.com", "mydomain.com"},
		{"mydomain.com:1234", "mydomain.com"},
		{"domainwithnumber99.com", "domainwithnumber99.com"},
		{"domainwithnumber99.com:1234", "domainwithnumber99.com"},
	}

	for _, test := range tests {
		if res := StripPort(test.src); res != test.expected {
			t.Errorf("Result of strip port on '%s', not as expected. Expected '%s'; got '%s'",
				test.src,
				test.expected,
				res,
			)
		}
	}
}
