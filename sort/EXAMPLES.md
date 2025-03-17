## Sort Function Examples

### Bubble Sort

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/sort"
)

func main() {
    numbers := []int{5, 3, 1, 2, 4}
    sortedNumbers := sort.BubbleSort(numbers)
    fmt.Println("Sorted Numbers:", sortedNumbers)
}
```

#### Output

```
Sorted Numbers: [1 2 3 4 5]
```

### Selection Sort

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/sort"
)

func main() {
    numbers := []int{5, 3, 1, 2, 4}
    sortedNumbers := sort.SelectionSort(numbers)
    fmt.Println("Sorted Numbers:", sortedNumbers)
}
```

#### Output

```
Sorted Numbers: [1 2 3 4 5]
```

### Insertion Sort

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/sort"
)

func main() {
    numbers := []int{5, 3, 1, 2, 4}
    sortedNumbers := sort.InsertionSort(numbers)
    fmt.Println("Sorted Numbers:", sortedNumbers)
}
```

#### Output

```
Sorted Numbers: [1 2 3 4 5]
```

### Merge Sort

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/sort"
)

func main() {
    numbers := []int{5, 3, 1, 2, 4}
    sortedNumbers := sort.MergeSort(numbers)
    fmt.Println("Sorted Numbers:", sortedNumbers)
}
```

#### Output

```
Sorted Numbers: [1 2 3 4 5]
```

### Quick Sort

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/sort"
)

func main() {
    numbers := []int{5, 3, 1, 2, 4}
    sortedNumbers := sort.QuickSort(numbers)
    fmt.Println("Sorted Numbers:", sortedNumbers)
}
```

#### Output

```
Sorted Numbers: [1 2 3 4 5]
```

### Heap Sort

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/sort"
)

func main() {
    numbers := []int{5, 3, 1, 2, 4}
    sortedNumbers := sort.HeapSort(numbers)
    fmt.Println("Sorted Numbers:", sortedNumbers)
}
```

#### Output

```
Sorted Numbers: [1 2 3 4 5]
```
