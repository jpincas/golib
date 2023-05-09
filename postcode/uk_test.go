package postcode

import "testing"

func TestParseUKPostcode(t *testing.T) {
	tests := []struct {
		name                            string
		in                              string
		expectedOutward, expectedInward string
		expectError                     bool
	}{
		// Easily identifiable incorrect Formatting
		{"blank", "", "", "", true},
		{"too short", "NW", "", "", true},

		// According to Wikipedia, these are the only valid foramts
		{"wiki1", "AA9A 9AA", "AA9A", "9AA", false},
		{"wiki2", "A9 9AA", "A9", "9AA", false},
		{"wiki3", "A99 9AA", "A99", "9AA", false},
		{"wiki4", "AA9 9AA", "AA9", "9AA", false},
		{"wiki5", "AA99 9AA", "AA99", "9AA", false},

		// Subtle variations on the above that shouldn't parse
		{"inward too short", "AA9A 9A", "", "", true},
		{"inward starts with digit", "A9 BAA", "", "", true},
		{"outward too short", "A 9AA", "", "", true},
		{"outward starts with digit", "9AA 9AA", "", "", true},

		// Formatting variations
		{"no space", "NW98QU", "NW9", "8QU", false},
		{"surrounding spaces", " NW98QU ", "NW9", "8QU", false},
		{"surrounding spaces and internal space", " NW9 8QU ", "NW9", "8QU", false},
		{"two digits in outward", "GX11 1AA", "GX11", "1AA", false},
		{"one char in outward", "N19GU", "N1", "9GU", false},
	}

	for _, test := range tests {
		res, err := ParseUKPostcode(test.in)
		if !test.expectError && err != nil {
			t.Fatalf("Testing %s. Did not expect error but error parsing postcode: %s", test.name, err)
		}

		if test.expectError && err == nil {
			t.Fatalf("Testing %s. Expected error but did not get one", test.name)
		}

		if res.OutwardCode != test.expectedOutward {
			t.Errorf("Testing %s. Outward. Expected %s but got %s", test.name, test.expectedOutward, res.OutwardCode)
		}

		if res.InwardCode != test.expectedInward {
			t.Errorf("Testing %s. Inward. Expected %s but got %s", test.name, test.expectedInward, res.InwardCode)
		}
	}

}
