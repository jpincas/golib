package email

import "testing"

func TestValidEmail(t *testing.T) {
	tests := []struct {
		email       string
		expectValid bool
	}{
		{"", false},
		{"jon@pincas.co.uk", true},
		{"jon@pincas.com", true},
		{"jon.pincas@pinacs.co.uk", true},
		{"jon,pincas@pincas.co.uk", false},
		{"jon@pincas", false},
		{"jon@pincas.co,uk", false},
		{"<jon@pincas.co.uk", false},
	}

	for _, test := range tests {
		res := IsValidEmail(test.email)

		if test.expectValid && res == false {
			t.Errorf("Testing email '%s'. Expected valid, but got invalid", test.email)
		}

		if !test.expectValid && res == true {
			t.Errorf("Testing email '%s'. Expected invalid, but got valid", test.email)
		}
	}
}
