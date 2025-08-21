package time

import (
	"reflect"
	"testing"
	"time"
)

// time.Now mock
var now = time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)

func TestStartOfDay(t *testing.T) {
	t.Parallel()

	type args struct {
		input time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success - Start of the day",
			args: args{time.Date(2024, 12, 25, 15, 30, 45, 123456789, time.UTC)},
			want: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "success - Start of the day",
			args: args{time.Date(2023, 1, 1, 23, 59, 59, 999999999, time.UTC)},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "success - Start of the day",
			args: args{time.Date(2022, 6, 15, 12, 0, 0, 0, time.FixedZone("TestZone", 3600))},
			want: time.Date(2022, 6, 15, 0, 0, 0, 0, time.FixedZone("TestZone", 3600)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := StartOfDay(tc.args.input)
			if !result.Equal(tc.want) {
				t.Errorf("StartOfDay(%v) = %v; want %v", tc.args.input, result, tc.want)
			}
		})
	}
}

func TestEndOfDay(t *testing.T) {
	t.Parallel()

	type args struct {
		input time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success - End of the day",
			args: args{time.Date(2024, 12, 25, 15, 30, 45, 123456789, time.UTC)},
			want: time.Date(2024, 12, 25, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name: "success - End of the day",
			args: args{time.Date(2023, 1, 1, 23, 59, 59, 999999999, time.UTC)},
			want: time.Date(2023, 1, 1, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name: "success - End of the day",
			args: args{time.Date(2022, 6, 15, 12, 0, 0, 0, time.FixedZone("TestZone", 3600))},
			want: time.Date(2022, 6, 15, 23, 59, 59, 999999999, time.FixedZone("TestZone", 3600)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := EndOfDay(tc.args.input)
			if !result.Equal(tc.want) {
				t.Errorf("EndOfDay(%v) = %v; want %v", tc.args.input, result, tc.want)
			}
		})
	}
}

func TestAddBusinessDays(t *testing.T) {
	t.Parallel()

	type args struct {
		t    time.Time
		days int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success - add 1 business day",
			args: args{
				t:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				days: 1,
			},
			want: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "success - add 2 business days",
			args: args{
				t:    time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC),
				days: 2,
			},
			want: time.Date(2023, 1, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "success - add 0 business days",
			args: args{
				t:    time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC),
				days: 0,
			},
			want: time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "success - add 20 business days",
			args: args{
				t:    time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
				days: 20,
			},
			want: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := AddBusinessDays(tt.args.t, tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBusinessDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - is weekend (Sunday)",
			args: args{t: time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "success - is not weekend (Monday)",
			args: args{t: time.Date(2024, 12, 2, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "success - is weekend (Saturday)",
			args: args{t: time.Date(2024, 12, 7, 0, 0, 0, 0, time.FixedZone("TestZone", 3600))},
			want: true,
		},
		{
			name: "success - is not weekend (Empty time)",
			args: args{t: time.Time{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsWeekend(tt.args.t); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeDifferenceHumanReadable(t *testing.T) {
	t.Parallel()

	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - exact same time",
			args: args{
				from: time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
			},
			want: "in 0 hour(s)",
		},
		{
			name: "success - future within 24 hours",
			args: args{
				from: time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 25, 15, 0, 0, 0, time.UTC),
			},
			want: "in 3 hour(s)",
		},
		{
			name: "success - future beyond 24 hours",
			args: args{
				from: time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 27, 12, 0, 0, 0, time.UTC),
			},
			want: "in 2 day(s)",
		},
		{
			name: "success - past within 24 hours",
			args: args{
				from: time.Date(2024, time.December, 25, 15, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
			},
			want: "in 3 hour(s)",
		},
		{
			name: "success - past beyond 24 hours",
			args: args{
				from: time.Date(2024, time.December, 27, 12, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
			},
			want: "2 day(s) ago",
		},
		{
			name: "success - negative time difference within 24 hours",
			args: args{
				from: time.Date(2024, time.December, 25, 15, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC),
			},
			want: "in 3 hour(s)",
		},
		{
			name: "success - negative time difference beyond 24 hours",
			args: args{
				from: time.Date(2024, time.December, 27, 15, 0, 0, 0, time.UTC),
				to:   time.Date(2024, time.December, 25, 15, 0, 0, 0, time.UTC),
			},
			want: "2 day(s) ago",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := TimeDifferenceHumanReadable(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("TimeDifferenceHumanReadable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDurationUntilNext(t *testing.T) {
	t.Parallel()

	type args struct {
		day time.Weekday
		t   time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "Next Monday from Friday",
			args: args{
				day: time.Monday,
				t:   time.Date(2023, 12, 22, 12, 0, 0, 0, time.UTC),
			},
			want: 3 * 24 * time.Hour,
		},
		{
			name: "Next Sunday from Saturday",
			args: args{
				day: time.Sunday,
				t:   time.Date(2023, 12, 23, 12, 0, 0, 0, time.UTC),
			},
			want: 24 * time.Hour,
		},
		{
			name: "Next Wednesday from Wednesday",
			args: args{
				day: time.Wednesday,
				t:   time.Date(2023, 12, 20, 12, 0, 0, 0, time.UTC),
			},
			want: 7 * 24 * time.Hour,
		},
		{
			name: "Next Friday from Thursday",
			args: args{
				day: time.Friday,
				t:   time.Date(2023, 12, 21, 12, 0, 0, 0, time.UTC),
			},
			want: 24 * time.Hour,
		},
		{
			name: "Next Monday from Monday",
			args: args{
				day: time.Monday,
				t:   time.Date(2023, 12, 18, 12, 0, 0, 0, time.UTC),
			},
			want: 7 * 24 * time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := DurationUntilNext(tt.args.day, tt.args.t); got != tt.want {
				t.Errorf("DurationUntilNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToTimeZone(t *testing.T) {
	t.Parallel()

	type args struct {
		t        time.Time
		location string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "Convert UTC to PST",
			args: args{
				t:        time.Date(2023, 12, 25, 15, 0, 0, 0, time.UTC), // 3 PM UTC
				location: "America/Los_Angeles",                          // PST
			},
			want:    time.Date(2023, 12, 25, 7, 0, 0, 0, time.FixedZone("PST", -8*60*60)),
			wantErr: false,
		},
		{
			name: "Convert UTC to GMT+5",
			args: args{
				t:        time.Date(2023, 12, 25, 15, 0, 0, 0, time.UTC), // 3 PM UTC
				location: "Asia/Karachi",                                 // GMT+5
			},
			want:    time.Date(2023, 12, 25, 20, 0, 0, 0, time.FixedZone("PKT", 5*60*60)),
			wantErr: false,
		},
		{
			name: "Convert UTC to invalid timezone",
			args: args{
				t:        time.Date(2023, 12, 25, 15, 0, 0, 0, time.UTC),
				location: "Invalid/Timezone",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "Convert UTC to local timezone",
			args: args{
				t:        time.Date(2023, 12, 25, 15, 0, 0, 0, time.UTC),
				location: "Local",
			},
			wantErr: false,
			want:    time.Date(2023, 12, 25, 15, 0, 0, 0, time.UTC).In(now.Location()),
		},
		{
			name: "Convert PST to IST",
			args: args{
				t:        time.Date(2023, 12, 25, 7, 0, 0, 0, time.FixedZone("PST", -8*60*60)),
				location: "Asia/Kolkata", // IST
			},
			want:    time.Date(2023, 12, 25, 20, 30, 0, 0, time.FixedZone("IST", 5*60*60+30*60)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ConvertToTimeZone(tt.args.t, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToTimeZone() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ConvertToTimeZone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanReadableDuration(t *testing.T) {
	t.Parallel()

	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Zero duration",
			args: args{
				d: 0,
			},
			want: "0h 0m 0s",
		},
		{
			name: "Only seconds",
			args: args{
				d: 45 * time.Second,
			},
			want: "0h 0m 45s",
		},
		{
			name: "Minutes and seconds",
			args: args{
				d: 90 * time.Second,
			},
			want: "0h 1m 30s",
		},
		{
			name: "Hours, minutes, and seconds",
			args: args{
				d: 2*time.Hour + 15*time.Minute + 42*time.Second,
			},
			want: "2h 15m 42s",
		},
		{
			name: "More than a day",
			args: args{
				d: 26*time.Hour + 5*time.Minute + 10*time.Second,
			},
			want: "26h 5m 10s",
		},
		{
			name: "Negative duration",
			args: args{
				d: -2*time.Hour - 10*time.Minute - 30*time.Second,
			},
			want: "-2h -10m -30s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := HumanReadableDuration(tt.args.d); got != tt.want {
				t.Errorf("HumanReadableDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateAge(t *testing.T) {
	t.Parallel()

	testTime := time.Now() // Fixed reference time
	type args struct {
		birthDate time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Birthday today, age remains the same",
			args: args{
				birthDate: time.Date(testTime.Year()-30, testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC),
			},
			want: 30,
		},
		{
			name: "Birthday not yet reached this year",
			args: args{
				birthDate: time.Date(testTime.Year()-25, testTime.Month()%12+1, testTime.Day(), 0, 0, 0, 0, time.UTC),
			},
			want: 24,
		},
		{
			name: "Birthday already passed this year",
			args: args{
				birthDate: time.Date(testTime.Year()-40, (testTime.Month()+11)%12+1, testTime.Day(), 0, 0, 0, 0, time.UTC),
			},
			want: 40,
		},
		{
			name: "Birthdate is exactly one year ago",
			args: args{
				birthDate: time.Date(testTime.Year()-1, testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC),
			},
			want: func() int {
				age := 1
				// If the birthday hasn't passed yet this year, decrement the age
				if testTime.Before(time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC)) {
					age--
				}
				return age
			}(),
		},
		{
			name: "Leap year adjustment, after birthday",
			args: args{
				birthDate: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			},
			want: func() int {
				leapYearAge := testTime.Year() - 2000

				// Check if the current year is not a leap year AND today is before March 1st
				if (testTime.Year()%4 != 0 || (testTime.Year()%100 == 0 && testTime.Year()%400 != 0)) &&
					testTime.Before(time.Date(testTime.Year(), 3, 1, 0, 0, 0, 0, time.UTC)) {
					leapYearAge--
				}

				return leapYearAge
			}(),
		},
		{
			name: "Future date, invalid age",
			args: args{
				birthDate: time.Date(testTime.Year()+5, testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC),
			},
			want: -5,
		},
		{
			name: "Birthdate today, zero age",
			args: args{
				birthDate: time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := CalculateAge(tt.args.birthDate); got != tt.want {
				t.Errorf("CalculateAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	t.Parallel()

	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Leap year divisible by 4 but not 100",
			args: args{year: 2024},
			want: true,
		},
		{
			name: "Leap year divisible by 400",
			args: args{year: 2000},
			want: true,
		},
		{
			name: "Not a leap year divisible by 100 but not 400",
			args: args{year: 1900},
			want: false,
		},
		{
			name: "Not a leap year divisible by 4 but not 100",
			args: args{year: 2023},
			want: false,
		},
		{
			name: "Year 0 (edge case)",
			args: args{year: 0},
			want: true,
		},
		{
			name: "Very far future year",
			args: args{year: 10000},
			want: true,
		},
		{
			name: "Very far future year divisible by 100 but not 400",
			args: args{year: 2100},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsLeapYear(tt.args.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextOccurrence(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("UTC")
	currentTime := time.Date(2024, time.December, 29, 10, 30, 0, 0, loc)

	type args struct {
		hour   int
		minute int
		second int
		t      time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Next occurrence after current time",
			args: args{
				hour:   11,
				minute: 30,
				second: 0,
				t:      currentTime,
			},
			want: time.Date(2024, time.December, 29, 11, 30, 0, 0, loc),
		},
		{
			name: "Next occurrence next day (time is in the past)",
			args: args{
				hour:   9,
				minute: 30,
				second: 0,
				t:      currentTime,
			},
			want: time.Date(2024, time.December, 30, 9, 30, 0, 0, loc),
		},
		{
			name: "Exact next second after current time",
			args: args{
				hour:   10,
				minute: 30,
				second: 1,
				t:      currentTime,
			},
			want: time.Date(2024, time.December, 29, 10, 30, 1, 0, loc),
		},
		{
			name: "Midnight next occurrence",
			args: args{
				hour:   0,
				minute: 0,
				second: 0,
				t:      currentTime,
			},
			want: time.Date(2024, time.December, 30, 0, 0, 0, 0, loc),
		},
		{
			name: "Same day, time is already passed",
			args: args{
				hour:   9,
				minute: 0,
				second: 0,
				t:      currentTime,
			},
			want: time.Date(2024, time.December, 30, 9, 0, 0, 0, loc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NextOccurrence(tt.args.hour, tt.args.minute, tt.args.second, tt.args.t)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextOccurrence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name     string
		args     args
		wantYear int
		wantWeek int
	}{
		{
			name: "Week 1 of the year 2024",
			args: args{
				t: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantYear: 2024,
			wantWeek: 1,
		},
		{
			name: "Last week of the year 2024",
			args: args{
				t: time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
			},
			wantYear: 2024,
			wantWeek: 52,
		},
		{
			name: "Middle of the year 2024",
			args: args{
				t: time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC),
			},
			wantYear: 2024,
			wantWeek: 24,
		},
		{
			name: "New Year's Eve 2025",
			args: args{
				t: time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			wantYear: 2025,
			wantWeek: 1,
		},
		{
			name: "Start of the week May 2024",
			args: args{
				t: time.Date(2024, 5, 5, 0, 0, 0, 0, time.UTC),
			},
			wantYear: 2024,
			wantWeek: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			year, week := WeekNumber(tt.args.t)
			if year != tt.wantYear {
				t.Errorf("WeekNumber() gotYear = %v, wantYear %v", year, tt.wantYear)
			}
			if week != tt.wantWeek {
				t.Errorf("WeekNumber() gotWeek = %v, wantWeek %v", week, tt.wantWeek)
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	t.Parallel()

	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Same day",
			args: args{
				start: time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 23, 59, 59, 0, time.UTC),
			},
			want: 0, // Same day, 0 days between
		},
		{
			name: "One day difference",
			args: args{
				start: time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC),
			},
			want: 1, // One day difference
		},
		{
			name: "Multiple days",
			args: args{
				start: time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			},
			want: 2, // Two days difference
		},
		{
			name: "Leap year day",
			args: args{
				start: time.Date(2020, 2, 28, 0, 0, 0, 0, time.UTC), // Feb 28, 2020 (Leap Year)
				end:   time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),  // Mar 1, 2020
			},
			want: 2, // 2 days difference (Leap year: Feb 29th included)
		},
		{
			name: "Negative days (start after end)",
			args: args{
				start: time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
			},
			want: -1, // Negative 1 day difference
		},
		{
			name: "Crossing months",
			args: args{
				start: time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 2, 2, 0, 0, 0, 0, time.UTC),
			},
			want: 2, // 2 days difference between Jan 31 and Feb 2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := DaysBetween(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("DaysBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTimeBetween(t *testing.T) {
	t.Parallel()

	type args struct {
		t     time.Time
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Time between start and end",
			args: args{
				t:     time.Date(2024, 12, 29, 10, 30, 0, 0, time.UTC),
				start: time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
			},
			want: true, // t is between start and end
		},
		{
			name: "Time exactly equal to start",
			args: args{
				t:     time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
				start: time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
			},
			want: true, // t is equal to start, should be considered inside the range
		},
		{
			name: "Time exactly equal to end",
			args: args{
				t:     time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
				start: time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
			},
			want: false, // t is equal to end, should not be considered inside the range
		},
		{
			name: "Time before start",
			args: args{
				t:     time.Date(2024, 12, 29, 9, 0, 0, 0, time.UTC),
				start: time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
			},
			want: false, // t is before start
		},
		{
			name: "Time after end",
			args: args{
				t:     time.Date(2024, 12, 29, 12, 0, 0, 0, time.UTC),
				start: time.Date(2024, 12, 29, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 12, 29, 11, 0, 0, 0, time.UTC),
			},
			want: false, // t is after end
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsTimeBetween(tt.args.t, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("IsTimeBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnixMilliToTime(t *testing.T) {
	t.Parallel()

	type args struct {
		ms int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Epoch time (0 milliseconds)",
			args: args{
				ms: 0,
			},
			want: time.Unix(0, 0).UTC(),
		},
		{
			name: "Specific time in the future",
			args: args{
				ms: time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC).UnixMilli(),
			},
			want: time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC),
		},
		{
			name: "Specific time in the past",
			args: args{
				ms: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(),
			},
			want: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Time with a fraction of second",
			args: args{
				ms: time.Date(2021, 1, 1, 12, 0, 1, 0, time.UTC).UnixMilli(),
			},
			want: time.Date(2021, 1, 1, 12, 0, 1, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := UnixMilliToTime(tt.args.ms).UTC(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnixMilliToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitDuration(t *testing.T) {
	t.Parallel()

	type args struct {
		d time.Duration
	}
	tests := []struct {
		name        string
		args        args
		wantDays    int
		wantHours   int
		wantMinutes int
		wantSeconds int
	}{
		{
			name: "Zero duration",
			args: args{
				d: 0,
			},
			wantDays:    0,
			wantHours:   0,
			wantMinutes: 0,
			wantSeconds: 0,
		},
		{
			name: "1 day 2 hours 30 minutes 45 seconds",
			args: args{
				d: 1*24*time.Hour + 2*time.Hour + 30*time.Minute + 45*time.Second,
			},
			wantDays:    1,
			wantHours:   2,
			wantMinutes: 30,
			wantSeconds: 45,
		},
		{
			name: "3 days 5 hours",
			args: args{
				d: 3*24*time.Hour + 5*time.Hour,
			},
			wantDays:    3,
			wantHours:   5,
			wantMinutes: 0,
			wantSeconds: 0,
		},
		{
			name: "No days, 1 hour 30 minutes",
			args: args{
				d: 1*time.Hour + 30*time.Minute,
			},
			wantDays:    0,
			wantHours:   1,
			wantMinutes: 30,
			wantSeconds: 0,
		},
		{
			name: "Negative duration",
			args: args{
				d: -26 * time.Hour,
			},
			wantDays:    -1,
			wantHours:   -2,
			wantMinutes: 0,
			wantSeconds: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotDays, gotHours, gotMinutes, gotSeconds := SplitDuration(tt.args.d)
			if gotDays != tt.wantDays {
				t.Errorf("SplitDuration() gotDays = %v, want %v", gotDays, tt.wantDays)
			}
			if gotHours != tt.wantHours {
				t.Errorf("SplitDuration() gotHours = %v, want %v", gotHours, tt.wantHours)
			}
			if gotMinutes != tt.wantMinutes {
				t.Errorf("SplitDuration() gotMinutes = %v, want %v", gotMinutes, tt.wantMinutes)
			}
			if gotSeconds != tt.wantSeconds {
				t.Errorf("SplitDuration() gotSeconds = %v, want %v", gotSeconds, tt.wantSeconds)
			}
		})
	}
}

func TestGetMonthName(t *testing.T) {
	t.Parallel()

	type args struct {
		monthNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid month number 1",
			args: args{
				monthNumber: 1,
			},
			want:    "January",
			wantErr: false,
		},
		{
			name: "Valid month number 12",
			args: args{
				monthNumber: 12,
			},
			want:    "December",
			wantErr: false,
		},
		{
			name: "Invalid month number 0",
			args: args{
				monthNumber: 0,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid month number 13",
			args: args{
				monthNumber: 13,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid month number 6",
			args: args{
				monthNumber: 6,
			},
			want:    "June",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetMonthName(tt.args.monthNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMonthName() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("GetMonthName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDayName(t *testing.T) {
	t.Parallel()

	type args struct {
		dayNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid day number 0",
			args: args{
				dayNumber: 0,
			},
			want:    "Sunday",
			wantErr: false,
		},
		{
			name: "Valid day number 3",
			args: args{
				dayNumber: 3,
			},
			want:    "Wednesday",
			wantErr: false,
		},
		{
			name: "Valid day number 6",
			args: args{
				dayNumber: 6,
			},
			want:    "Saturday",
			wantErr: false,
		},
		{
			name: "Invalid day number -1",
			args: args{
				dayNumber: -1,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid day number 7",
			args: args{
				dayNumber: 7,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetDayName(tt.args.dayNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDayName() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("GetDayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatForDisplay(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid date",
			args: args{
				t: time.Date(2024, time.December, 30, 10, 30, 0, 0, time.UTC),
			},
			want: "Monday, 30 Dec 2024",
		},
		{
			name: "Edge case: January 1st, 2023",
			args: args{
				t: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			want: "Sunday, 1 Jan 2023",
		},
		{
			name: "Edge case: December 31st, 2024",
			args: args{
				t: time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			want: "Tuesday, 31 Dec 2024",
		},
		{
			name: "Edge case: Empty time",
			args: args{
				t: time.Time{},
			},
			want: "Monday, 1 Jan 0001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := FormatForDisplay(tt.args.t); got != tt.want {
				t.Errorf("FormatForDisplay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsToday(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is today",
			args: args{
				t: time.Now(),
			},
			want: true,
		},
		{
			name: "Yesterday",
			args: args{
				t: time.Now().AddDate(0, 0, -1),
			},
			want: false,
		},
		{
			name: "Tomorrow",
			args: args{
				t: time.Now().AddDate(0, 0, 1),
			},
			want: false,
		},
		{
			name: "Same date, different time",
			args: args{
				t: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.UTC),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsToday(tt.args.t); got != tt.want {
				t.Errorf("IsToday() = %v, want %v", got, tt.want)
			}
		})
	}
}
