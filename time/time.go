package time

import (
	"fmt"
	"time"
)

// StartOfDay returns the start of the day for the given time
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day for the given time
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// AddBusinessDays adds the given number of business days to the given time
func AddBusinessDays(t time.Time, days int) time.Time {
	for days > 0 {
		t = t.AddDate(0, 0, 1)
		if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
			days--
		}
	}

	return t
}

// IsWeekend checks if the given time is a weekend
func IsWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

// TimeDifferenceHumanReadable returns a human-readable string representing the time difference between two times
func TimeDifferenceHumanReadable(from, to time.Time) string {
	diff := to.Sub(from)
	if diff < 0 {
		diff = -diff
		if diff.Hours() > 24 {
			return fmt.Sprintf("%d day(s) ago", int(diff.Hours()/24))
		}
		return fmt.Sprintf("in %d hour(s)", int(diff.Hours()))
	}
	if diff.Hours() > 24 {
		return fmt.Sprintf("in %d day(s)", int(diff.Hours()/24))
	}
	return fmt.Sprintf("in %d hour(s)", int(diff.Hours()))
}

// DurationUntilNext calculates the duration from the given time `t` to the next occurrence
// of the specified weekday `day`. If the target day is the same as the current day,
// it assumes the next occurrence of that day is in 7 days.
func DurationUntilNext(day time.Weekday, t time.Time) time.Duration {
	daysAhead := (int(day) - int(t.Weekday()) + 7) % 7
	if daysAhead == 0 {
		daysAhead = 7
	}

	next := t.AddDate(0, 0, daysAhead)
	return next.Sub(t)
}

// ConvertToTimeZone converts the given time to the specified time zone
func ConvertToTimeZone(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}

// HumanReadableDuration returns a human-readable string representing the given duration
func HumanReadableDuration(d time.Duration) string {
	hours := d / time.Hour
	minutes := (d % time.Hour) / time.Minute
	seconds := (d % time.Minute) / time.Second

	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

// CalculateAge calculates the age of a person given their birth date
func CalculateAge(birthDate time.Time) int {
	today := time.Now()
	age := today.Year() - birthDate.Year()
	if today.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}

// IsLeapYear checks if the given year is a leap year
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// NextOccurrence returns the next occurrence of the specified time on the same day as the given time
func NextOccurrence(hour, minute, second int, t time.Time) time.Time {
	next := time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, 0, t.Location())
	if !next.After(t) {
		next = next.Add(24 * time.Hour)
	}
	return next
}

// WeekNumber returns the year and week number for the given time
func WeekNumber(t time.Time) (int, int) {
	year, week := t.ISOWeek()
	return year, week
}

// DaysBetween returns the number of days between two times
func DaysBetween(start, end time.Time) int {
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
	return int(end.Sub(start).Hours() / 24)
}

// IsTimeBetween checks if the given time is between the start and end times
func IsTimeBetween(t, start, end time.Time) bool {
	return (t.After(start) || t.Equal(start)) && t.Before(end)
}

// UnixMilliToTime converts milliseconds since epoch to time.Time
func UnixMilliToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}

// SplitDuration splits the given duration into days, hours, minutes, and seconds
func SplitDuration(d time.Duration) (days, hours, minutes, seconds int) {
	days = int(d.Hours()) / 24
	hours = int(d.Hours()) % 24
	minutes = int(d.Minutes()) % 60
	seconds = int(d.Seconds()) % 60
	return
}

// GetMonthName returns the name of the month for the given month number
func GetMonthName(monthNumber int) (string, error) {
	months := []string{
		"",          // Index 0 (placeholder, as months start from 1)
		"January",   // 1
		"February",  // 2
		"March",     // 3
		"April",     // 4
		"May",       // 5
		"June",      // 6
		"July",      // 7
		"August",    // 8
		"September", // 9
		"October",   // 10
		"November",  // 11
		"December",  // 12
	}

	if monthNumber < 1 || monthNumber > 12 {
		return "", fmt.Errorf("invalid month number: %d", monthNumber)
	}

	return months[monthNumber], nil
}

// GetDayName returns the name of the day for the given day number
func GetDayName(dayNumber int) (string, error) {
	days := []string{
		"Sunday",    // 0
		"Monday",    // 1
		"Tuesday",   // 2
		"Wednesday", // 3
		"Thursday",  // 4
		"Friday",    // 5
		"Saturday",  // 6
	}

	if dayNumber < 0 || dayNumber > 6 {
		return "", fmt.Errorf("invalid day number: %d", dayNumber)
	}

	return days[dayNumber], nil
}
