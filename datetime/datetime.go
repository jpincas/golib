package datetime

import "time"

const (
	BritishDate = "02/01/2006"
)

func FormatBritishDate(dt time.Time) string {
	return dt.Format(BritishDate)
}

func AddDays(t time.Time, days int) time.Time {
	return t.Add(time.Hour * 24 * time.Duration(days))
}

// DaysFromNowIgnoringTime works out the day differential from now but based on actual days, rather than units of 24hrs
func DaysOverdueIgnoringTime(due time.Time) int {
	return daysDiffIgnoringTime(due, false)
}

func DueInDaysIgnoringTime(due time.Time) int {
	return daysDiffIgnoringTime(due, true)
}

func daysDiffIgnoringTime(due time.Time, dueInDaysStyle bool) int {
	// By default, do the calculation 'days overdue' style
	daysDiff := time.Now().YearDay() - due.YearDay()
	yearsDiff := time.Now().Year() - due.Year()

	// But can also do it 'due in days style' (just subtract the other way round)
	if dueInDaysStyle {
		daysDiff = due.YearDay() - time.Now().YearDay()
		yearsDiff = due.Year() - time.Now().Year()
	}

	totalDaysDiff := (yearsDiff * 365) + daysDiff

	if totalDaysDiff <= 0 {
		totalDaysDiff = 0
	}

	return totalDaysDiff
}

func TimeIsWithin(t, rangeStart, rangeEnd time.Time) bool {
	return (t == rangeStart || t.After(rangeStart)) && (t == rangeEnd || t.Before(rangeEnd))
}

func TimesAreEqualIgnoreMilliseconds(t1, t2 time.Time) bool {
	// Just pick the right output format and compare strings
	f := time.RFC850
	return t1.Format(f) == t2.Format(f)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func InNDaysFrom(now time.Time, n int) time.Time {
	return now.AddDate(0, 0, n)
}
