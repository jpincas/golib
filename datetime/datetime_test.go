package datetime

import (
	"testing"
	"time"
)

func TestTimeIsWithin(t *testing.T) {
	tests := []struct {
		name                 string
		t                    time.Time
		rangeStart, rangeEnd time.Time
		expect               bool
	}{
		{
			name:       "all blank",
			t:          time.Time{},
			rangeStart: time.Time{},
			rangeEnd:   time.Time{},
			expect:     true,
		},
		{
			name:       "blank range, target",
			t:          time.Date(2020, time.January, 2, 13, 0, 0, 0, time.UTC),
			rangeStart: time.Time{},
			rangeEnd:   time.Time{},
			expect:     false,
		},
		{
			name:       "range, target is perfectly inside",
			t:          time.Date(2020, time.January, 3, 13, 0, 0, 0, time.UTC),
			rangeStart: time.Date(2020, time.January, 2, 13, 0, 0, 0, time.UTC),
			rangeEnd:   time.Date(2020, time.January, 5, 13, 0, 0, 0, time.UTC),
			expect:     true,
		},
		{
			name:       "range, target is on start",
			t:          time.Date(2020, time.January, 2, 13, 0, 0, 0, time.UTC),
			rangeStart: time.Date(2020, time.January, 2, 13, 0, 0, 0, time.UTC),
			rangeEnd:   time.Date(2020, time.January, 5, 13, 0, 0, 0, time.UTC),
			expect:     true,
		},

		{
			name:       "range, target is on end",
			t:          time.Date(2020, time.January, 5, 13, 0, 0, 0, time.UTC),
			rangeStart: time.Date(2020, time.January, 2, 13, 0, 0, 0, time.UTC),
			rangeEnd:   time.Date(2020, time.January, 5, 13, 0, 0, 0, time.UTC),
			expect:     true,
		},
		{
			name:       "blank range, target is outside",
			t:          time.Date(2020, time.January, 6, 13, 0, 0, 0, time.UTC),
			rangeStart: time.Date(2020, time.January, 2, 13, 0, 0, 0, time.UTC),
			rangeEnd:   time.Date(2020, time.January, 5, 13, 0, 0, 0, time.UTC),
			expect:     false,
		},
	}

	for _, test := range tests {
		res := TimeIsWithin(test.t, test.rangeStart, test.rangeEnd)
		if res != test.expect {
			t.Errorf("Testing %s.  Result was not as expected", test.name)
		}
	}

}
