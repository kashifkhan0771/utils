# Examples

This document provides practical examples of how to use the library's features. Each section demonstrates a specific use case with clear, concise code snippets.

## Table of Contents

1. [Boolean](/boolean/EXAMPLES.md)
2. [Context (ctxutils)](/ctxutils/EXAMPLES.md)
3. [Error (errutils)](/errutils/EXAMPLES.md)
4. [Maps](/maps/EXAMPLES.md)
5. [Pointers](/pointers/EXAMPLES.md)
6. [Random (rand)](/rand/EXAMPLES.md)
7. [Slice](/slice/EXAMPLES.md)
8. [Strings](/strings/EXAMPLES.md)
9. [Structs](/structs/EXAMPLES.md)
10. [Templates](/templates/EXAMPLES.md)
11. [URLs](/url/EXAMPLES.md)
12. [Math](/math/EXAMPLES.md)
13. [Fake](#13-fake)
14. [Time](#14-time)
15. [Loggin](#15-logging)
16. [File System Utilities](#16-fsutils)
17. [Loggin](#15-logging)
18. [File System Utilities](#16-fsutils)
19. [Caching](#15-caching)

## 13. Fake

### Generate a random UUID of version 4 and variant 2

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	uuid, err := fake.RandomUUID()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(uuid)
}

```

#### Output:

```
93a540eb-46e4-4e52-b0d5-cb63a7c361f9
```

### Generate a random date between 1st January 1970 and the current date

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	date, err := fake.RandomDate()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date)
}

```

#### Output:

```
2006-06-13 21:31:17.312528419 +0200 CEST
```

### Generate a random US phone number

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	num, err := fake.RandomPhoneNumber()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num)
}

```

#### Output:

```
+1 (965) 419-5534
```

### Generates a random US address

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	address, err := fake.RandomAddress()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(address)
}

```

#### Output:

```
81 Broadway, Rivertown, CT 12345, USA
```

---

## 14. Time

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

## 15. Logging

### Create and use a logger

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a new logger with prefix "MyApp", minimum level INFO, and output to stdout
	logger := logging.NewLogger("MyApp", logging.INFO, os.Stdout)

	// Log messages of different levels
	logger.Debug("This is a debug message.") // Ignored because minLevel is INFO
	logger.Info("Application started.")      // Printed with blue color
	logger.Warn("Low disk space.")           // Printed with yellow color
	logger.Error("Failed to connect to DB.") // Printed with red color
}
```

#### Output:

```
[2025-01-09 12:34:56] [INFO] MyApp: Application started.
[2025-01-09 12:34:56] [WARN] MyApp: Low disk space.
[2025-01-09 12:34:56] [ERROR] MyApp: Failed to connect to DB.
```

### Log without colors (useful for plain text logs)

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger and disable colors
	logger := logging.NewLogger("MyApp", logging.DEBUG, os.Stdout)
	logger.disableColors = true

	// Log messages of different levels
	logger.Debug("Debugging without colors.")
	logger.Info("Information without colors.")
	logger.Warn("Warning without colors.")
	logger.Error("Error without colors.")
}
```

#### Output:

```
[2025-01-09 12:34:56] [DEBUG] MyApp: Debugging without colors.
[2025-01-09 12:34:56] [INFO] MyApp: Information without colors.
[2025-01-09 12:34:56] [WARN] MyApp: Warning without colors.
[2025-01-09 12:34:56] [ERROR] MyApp: Error without colors.
```

### Log messages to a file

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Open a log file for writing
	file, err := os.Create("app.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a logger with file output
	logger := logging.NewLogger("MyApp", logging.DEBUG, file)

	// Log messages
	logger.Debug("Writing debug logs to file.")
	logger.Info("Application log stored in file.")
	logger.Warn("This is a warning.")
	logger.Error("This is an error.")
}
```

#### Output (in `app.log` file):

```
[2025-01-09 12:34:56] [DEBUG] MyApp: Writing debug logs to file.
[2025-01-09 12:34:56] [INFO] MyApp: Application log stored in file.
[2025-01-09 12:34:56] [WARN] MyApp: This is a warning.
[2025-01-09 12:34:56] [ERROR] MyApp: This is an error.
```

### Filter logs by minimum log level

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger with minimum level WARN
	logger := logging.NewLogger("MyApp", logging.WARN, os.Stdout)

	// Log messages
	logger.Debug("This is a debug message.") // Ignored
	logger.Info("This is an info message.")  // Ignored
	logger.Warn("This is a warning.")        // Printed
	logger.Error("This is an error.")        // Printed
}
```

#### Output:

```
[2025-01-09 12:34:56] [WARN] MyApp: This is a warning.
[2025-01-09 12:34:56] [ERROR] MyApp: This is an error.
```

### Customize log prefixes

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger with a custom prefix
	logger := logging.NewLogger("CustomPrefix", logging.INFO, os.Stdout)

	// Log messages
	logger.Info("This message has a custom prefix.")
}
```

#### Output:

```
[2025-01-09 12:34:56] [INFO] CustomPrefix: This message has a custom prefix.
```

## 16. Fsutils

### Format a file size given in bytes into a human-readable format

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	sizes := []int64{0, 512, 1024, 1048576, 1073741824, 1099511627776}

	for _, size := range sizes {
		fmt.Println(fsutils.FormatFileSize(size))
	}
}
```

#### Output:

```
0 B
512 B
1.00 KB
1.00 MB
1.00 GB
1.00 TB
```

### Search for files with the specified extension

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir := "/path/to/your/dir"

	txtFiles, err := fsutils.FindFiles(dir, ".txt")
	if err != nil {
		log.Fatalf("Error finding .txt files: %v", err)
	}

	fmt.Println("TXT Files:", txtFiles)

	logFiles, err := fsutils.FindFiles(dir, ".log")
	if err != nil {
		log.Fatalf("Error finding .log files: %v", err)
	}

	fmt.Println("LOG Files:", logFiles)

	allFiles, err := fsutils.FindFiles(dir, "")
	if err != nil {
		log.Fatalf("Error finding all files: %v", err)
	}

	fmt.Println("All Files:", allFiles)
}

```

#### Output:

```
TXT Files: [/path/to/your/dir/file1.txt /path/to/your/dir/file2.txt /path/to/your/dir/file4.txt]
LOG Files: [/path/to/your/dir/file3.log]
All Files: [/path/to/your/dir/file1.txt /path/to/your/dir/file2.txt /path/to/your/dir/file3.log /path/to/your/dir/file4.txt]
```

### Calculate the total size (in bytes) of all files in a directory

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir := "/path/to/your/dir"

	size, err := fsutils.GetDirectorySize(dir)
	if err != nil {
		log.Fatalf("Error calculating directory size: %v", err)
	}

	fmt.Printf("The total size of directory \"%s\" is %dB\n", dir, size)
}

```

#### Output:

```
The total size of directory "/path/to/your/dir" is 6406B
```

### Compare two files

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	file1 := "/path/to/your/file1.txt"
	file2 := "/path/to/your/file2.txt"

	identical, err := fsutils.FilesIdentical(file1, file2)
	if err != nil {
		log.Fatalf("Error comparing files: %v", err)
	}

	if identical {
		fmt.Printf("The files %s and %s are identical\n", file1, file2)
	} else {
		fmt.Printf("The files %s and %s are not identical\n", file1, file2)
	}
}

```

#### Output:

```
The files /path/to/your/file1.txt and /path/to/your/file2.txt are identical
```

### Compare two directories

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir1 := "/path/to/your/dir1"
	dir2 := "/path/to/your/dir2"

	identical, err := fsutils.DirsIdentical(dir1, dir2)
	if err != nil {
		log.Fatalf("Error comparing directories: %v", err)
	}

	if identical {
		fmt.Printf("The directories %s and %s are identical.\n", dir1, dir2)
	} else {
		fmt.Printf("The directories %s and %s are not identical.\n", dir1, dir2)
	}
}

```

#### Output:

```
The directories /path/to/your/dir1 and /path/to/your/dir2 are identical.
```

### Get File Metadata

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	file := "example.txt"
	metadata, err := fsutils.GetFileMetadata(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(
		"Name: %s, Size: %d, IsDir: %t, ModTime: %s, Mode: %v, Path: %s, Ext: %s, Owner: %s\n",
		metadata.Name, metadata.Size,
		metadata.IsDir, metadata.ModTime.String(),
		metadata.Mode, metadata.Path,
		metadata.Ext, metadata.Owner,
	)
}

```

#### Output:

```
Name: example.txt, Size: 172, IsDir: false, ModTime: 2025-01-20 15:03:00.189199994 +0100 CET, Mode: -rw-rw-r--, Path: /path/to/your/dir/example.txt, Ext: .txt, Owner: owner
```

### Get Directory Metadata

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir := "example/"
	metadata, err := fsutils.GetFileMetadata(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(
		"Name: %s, Size: %d, IsDir: %t, ModTime: %s, Mode: %v, Path: %s, Ext: %s, Owner: %s\n",
		metadata.Name, metadata.Size,
		metadata.IsDir, metadata.ModTime.String(),
		metadata.Mode, metadata.Path,
		metadata.Ext, metadata.Owner,
	)
}

```

#### Output:

```
Name: example, Size: 4096, IsDir: true, ModTime: 2025-01-20 15:06:23.057206656 +0100 CET, Mode: drwxrwxr-x, Path: /path/to/your/dir/example, Ext: , Owner: owner
```

### Marshal File's Metadata to JSON

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	file := "example.txt"
	metadata, err := fsutils.GetFileMetadata(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	json, err := json.Marshal(&metadata)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(json))
}

```

#### Output:

```json
{
  "name": "example.txt",
  "size": 172,
  "is_dir": false,
  "mod_time": "2025-01-20T15:06:34.812677487+01:00",
  "mode": 436,
  "path": "/path/to/your/dir/example.txt",
  "ext": ".txt"
}
```

# 17. Caching

## `CacheWrapper`

### A non-thread-safe caching decorator

```go
package main

import (
	"fmt"
	"math/big"
	"github.com/kashifkhan0771/utils/math"
)

// Example function: Compute factorial
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	cachedFactorial := utils.CacheWrapper(factorial)
	fmt.Println(cachedFactorial(10))
}
```

#### Output:

```
3628800
```

---

## SafeCacheWrapper

### A thread-safe caching decorator

```go
package main

import (
	"fmt"
	"math/big"
	"github.com/kashifkhan0771/utils/math"
)

// Example function: Compute factorial
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	cachedFactorial := utils.SafeCacheWrapper(factorial)
	fmt.Println(cachedFactorial(10))
}
```

#### Output:

```
3628800
```
