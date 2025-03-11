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
13. [Fake](/fake/EXAMPLES.md)
14. [Time](/time/EXAMPLES.md)
15. [Logging](/logging/EXAMPLES.md)
16. [File System Utilities](#16-fsutils)
17. [Caching](#15-caching)

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
