package slice

import "testing"

func TestCompareStringSlicesOrderIrrelevant(t *testing.T) {
	testCases := []struct {
		name           string
		ss1            []string
		ss2            []string
		expectedResult bool
	}{
		{
			"two blank slices",
			[]string{},
			[]string{},
			true,
		},
		{
			"one blank, one single member",
			[]string{},
			[]string{"t1"},
			false,
		},
		{
			"two single members, identical",
			[]string{"t1"},
			[]string{"t1"},
			true,
		},
		{
			"two single members, different",
			[]string{"t1"},
			[]string{"t2"},
			false,
		},
		{
			"one single member, one double member, no overlap",
			[]string{"t1"},
			[]string{"t2", "t3"},
			false,
		},
		{
			"one single member, one double member, overlap",
			[]string{"t1"},
			[]string{"t2", "t1"},
			false,
		},
		{
			"one single member, one double member, changed order, overlap",
			[]string{"t1"},
			[]string{"t1", "t2"},
			false,
		},

		{
			"two double members, no overlap",
			[]string{"t1", "t2"},
			[]string{"t3", "t4"},
			false,
		},
		{
			"two double members, overlap",
			[]string{"t1", "t2"},
			[]string{"t3", "t2"},
			false,
		},
		{
			"two double members, same, different order",
			[]string{"t1", "t2"},
			[]string{"t2", "t1"},
			true,
		},
		{
			"two double members, same, same order",
			[]string{"t1", "t2"},
			[]string{"t1", "t2"},
			true,
		},
	}

	for _, testCase := range testCases {
		result := CompareStringSlicesOrderIrrelevant(testCase.ss1, testCase.ss2)
		if result != testCase.expectedResult {
			t.Errorf("Testing %s. Expected %v, got %v",
				testCase.name,
				testCase.expectedResult,
				result,
			)
		}
	}
}

func TestStringSliceReturnLastNonBlank(t *testing.T) {
	tests := []struct {
		src      []string
		expected string
	}{
		{
			src:      []string{},
			expected: "",
		},
		{
			src:      []string{""},
			expected: "",
		},
		{
			src:      []string{"", "jon"},
			expected: "jon",
		},
		{
			src:      []string{"jon", ""},
			expected: "jon",
		},
		{
			src:      []string{"jon1", "", "jon2"},
			expected: "jon2",
		},
	}

	for _, test := range tests {
		if res := LastNonEmptyString(test.src...); res != test.expected {
			t.Errorf("Testing %v, expected %s, got %s", test.src, test.expected, res)
		}
	}
}
