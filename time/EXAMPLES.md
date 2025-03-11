## Time Function Examples

## 1. `StartOfDay`

### Get the start of the day for the given time

```go
package main

import (
	"fmt"
	"time"

	utils "github.com/kashifkhan0771/utils/time"
)

func main() {
	t := time.Now()
	fmt.Println(utils.StartOfDay(t))
}
```

#### Output:

```
2024-12-29 00:00:00 +0500 PKT
```

---

## 2. `EndOfDay`

### Get the end of the day for the given time

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    fmt.Println(utils.EndOfDay(t))
}
```

#### Output:

```
2024-12-29 23:59:59.999999999 +0500 PKT
```

---

## 3. `AddBusinessDays`

### Add business days to a date (skipping weekends)

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Date(2024, 12, 27, 0, 0, 0, 0, time.Local) // Friday
    // Add 3 business days
    result := utils.AddBusinessDays(t, 3)
    fmt.Println(result)
}
```

#### Output:

```
2025-01-01 00:00:00 +0500 PKT
```

---

## 4. `IsWeekend`

### Check if a given date is a weekend

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    saturday := time.Date(2024, 12, 28, 0, 0, 0, 0, time.Local)
    monday := time.Date(2024, 12, 30, 0, 0, 0, 0, time.Local)

    fmt.Printf("Is Saturday a weekend? %v\n", utils.IsWeekend(saturday))
    fmt.Printf("Is Monday a weekend? %v\n", utils.IsWeekend(monday))
}
```

#### Output:

```
Is Saturday a weekend? true
Is Monday a weekend? false
```

---

## 5. `TimeDifferenceHumanReadable`

### Get human-readable time difference

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    future := now.Add(72 * time.Hour)
    past := now.Add(-48 * time.Hour)

    fmt.Println(utils.TimeDifferenceHumanReadable(now, future))
    fmt.Println(utils.TimeDifferenceHumanReadable(now, past))
}
```

#### Output:

```
in 3 day(s)
2 day(s) ago
```

---

## 6. `DurationUntilNext`

### Calculate duration until next specified weekday

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    nextMonday := utils.DurationUntilNext(time.Monday, now)
    fmt.Printf("Duration until next Monday: %v\n", nextMonday)
}
```

#### Output:

```
Duration until next Monday: 24h0m0s
```

---

## 7. `ConvertToTimeZone`

### Convert time to different timezone

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    nyTime, err := utils.ConvertToTimeZone(t, "America/New_York")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(nyTime)
}
```

#### Output:

```
2024-12-29 14:00:00 -0500 EST
```

---

## 8. `HumanReadableDuration`

### Format duration in human-readable format

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    d := 3*time.Hour + 25*time.Minute + 45*time.Second
    fmt.Println(utils.HumanReadableDuration(d))
}
```

#### Output:

```
3h 25m 45s
```

---

## 9. `CalculateAge`

### Calculate age from birthdate

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.Local)
    age := utils.CalculateAge(birthDate)
    fmt.Printf("Age: %d years\n", age)
}
```

#### Output:

```
Age: 34 years
```

---

## 10. `IsLeapYear`

### Check if a year is a leap year

```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    fmt.Printf("Is 2024 a leap year? %v\n", utils.IsLeapYear(2024))
    fmt.Printf("Is 2023 a leap year? %v\n", utils.IsLeapYear(2023))
}
```

#### Output:

```
Is 2024 a leap year? true
Is 2023 a leap year? false
```

---

## 11. `NextOccurrence`

### Find next occurrence of a specific time

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    nextNoon := utils.NextOccurrence(12, 0, 0, now)
    fmt.Println("Next noon:", nextNoon)
}
```

#### Output:

```
Next noon: 2024-12-30 12:00:00 +0500 PKT
```

---

## 12. `WeekNumber`

### Get ISO year and week number

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    year, week := utils.WeekNumber(t)
    fmt.Printf("Year: %d, Week: %d\n", year, week)
}
```

#### Output:

```
Year: 2024, Week: 52
```

---

## 13. `DaysBetween`

### Calculate days between two dates

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
    end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local)
    days := utils.DaysBetween(start, end)
    fmt.Printf("Days between: %d\n", days)
}
```

#### Output:

```
Days between: 365
```

---

## 14. `IsTimeBetween`

### Check if time is between two other times

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    start := now.Add(-1 * time.Hour)
    end := now.Add(1 * time.Hour)
    fmt.Printf("Is current time between? %v\n", utils.IsTimeBetween(now, start, end))
}
```

#### Output:

```
Is current time between? true
```

---

## 15. `UnixMilliToTime`

### Convert Unix milliseconds to time

```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    ms := int64(1703836800000) // 2024-12-29 00:00:00
    t := utils.UnixMilliToTime(ms)
    fmt.Println(t)
}
```

#### Output:

```
2024-12-29 00:00:00 +0000 UTC
```

---

## 16. `SplitDuration`

### Split duration into components

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    d := 50*time.Hour + 30*time.Minute + 15*time.Second
    days, hours, minutes, seconds := utils.SplitDuration(d)
    fmt.Printf("Days: %d, Hours: %d, Minutes: %d, Seconds: %d\n",
        days, hours, minutes, seconds)
}
```

#### Output:

```
Days: 2, Hours: 2, Minutes: 30, Seconds: 15
```

---

## 17. `GetMonthName`

### Get month name from number

```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    monthName, err := utils.GetMonthName(12)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Month 12 is: %s\n", monthName)
}
```

#### Output:

```
Month 12 is: December
```

---

## 18. `GetDayName`

### Get day name from number

```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    dayName, err := utils.GetDayName(1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Day 1 is: %s\n", dayName)
}
```

#### Output:

```
Day 1 is: Monday
```

---

## 19. `FormatForDisplay`

### Format time for display

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    formatted := utils.FormatForDisplay(t)
    fmt.Println(formatted)
}
```

#### Output:

```
Sunday, 29 Dec 2024
```

---

## 20. `IsToday`

### Check if date is today

```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    tomorrow := now.AddDate(0, 0, 1)

    fmt.Printf("Is now today? %v\n", utils.IsToday(now))
    fmt.Printf("Is tomorrow today? %v\n", utils.IsToday(tomorrow))
}
```

#### Output:

```
Is now today? true
Is tomorrow today? false
```

---
