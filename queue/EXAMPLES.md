# Queue Examples

This document provides comprehensive examples for all functions in the queue package.

## Table of Contents

- [NewQueue](#newqueue)
- [Enqueue](#enqueue)
- [Dequeue](#dequeue)
- [Peek](#peek)
- [Size, Capacity, IsEmpty](#size-capacity-isempty)
- [Dynamic Resizing](#dynamic-resizing)
- [Different Types](#different-types)
- [Concurrent Usage](#concurrent-usage)
- [Error Handling](#error-handling)
- [Complete Usage Example](#complete-usage-example)

## NewQueue

Creates a new queue with the specified initial capacity. If capacity is zero or negative, defaults to 16.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    // Create a queue with initial capacity of 10
    q1 := queue.NewQueue[int](10)
    fmt.Printf("Queue capacity: %d, size: %d, empty: %v\n", q1.Capacity(), q1.Size(), q1.IsEmpty())
    
    // Create a queue with default capacity (zero becomes 16)
    q2 := queue.NewQueue[string](0)
    fmt.Printf("Default queue capacity: %d\n", q2.Capacity())
    
    // Negative capacity also becomes default
    q3 := queue.NewQueue[float64](-5)
    fmt.Printf("Negative capacity queue: %d\n", q3.Capacity())
}
```

**Output:**
```
Queue capacity: 10, size: 0, empty: true
Default queue capacity: 16
Negative capacity queue: 16
```

## Enqueue

Adds elements to the end of the queue. The queue automatically grows when needed.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[int](4) // Small capacity to demonstrate growth
    
    fmt.Printf("Initial: capacity=%d, size=%d\n", q.Capacity(), q.Size())
    
    // Add elements one by one
    for i := 1; i <= 8; i++ {
        q.Enqueue(i)
        fmt.Printf("After enqueue %d: capacity=%d, size=%d\n", i, q.Capacity(), q.Size())
    }
}
```

**Output:**
```
Initial: capacity=4, size=0
After enqueue 1: capacity=4, size=1
After enqueue 2: capacity=4, size=2
After enqueue 3: capacity=4, size=3
After enqueue 4: capacity=4, size=4
After enqueue 5: capacity=8, size=5
After enqueue 6: capacity=8, size=6
After enqueue 7: capacity=8, size=7
After enqueue 8: capacity=8, size=8
```

## Dequeue

Removes and returns the front element from the queue. Returns an error if the queue is empty.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[string](4)
    
    // Add some elements
    items := []string{"first", "second", "third", "fourth"}
    for _, item := range items {
        q.Enqueue(item)
    }
    
    fmt.Printf("Queue size: %d\n", q.Size())
    
    // Dequeue all elements
    for !q.IsEmpty() {
        item, err := q.Dequeue()
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            break
        }
        fmt.Printf("Dequeued: %s, remaining size: %d\n", item, q.Size())
    }
    
    // Try to dequeue from empty queue
    _, err := q.Dequeue()
    if err != nil {
        fmt.Printf("Empty queue error: %v\n", err)
    }
}
```

**Output:**
```
Queue size: 4
Dequeued: first, remaining size: 3
Dequeued: second, remaining size: 2
Dequeued: third, remaining size: 1
Dequeued: fourth, remaining size: 0
Empty queue error: queue is empty
```

## Peek

Returns the front element without removing it from the queue. Returns an error if the queue is empty.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[int](4)
    
    // Try to peek empty queue
    val, err := q.Peek()
    if err != nil {
        fmt.Printf("Empty queue peek error: %v\n", err)
    }
    
    // Add elements and peek
    q.Enqueue(100)
    q.Enqueue(200)
    q.Enqueue(300)
    
    // Multiple peeks should return the same value
    for i := 1; i <= 3; i++ {
        val, err := q.Peek()
        if err != nil {
            fmt.Printf("Peek error: %v\n", err)
        } else {
            fmt.Printf("Peek %d: %d (size: %d)\n", i, val, q.Size())
        }
    }
    
    // Dequeue one element and peek again
    q.Dequeue()
    val, err = q.Peek()
    if err != nil {
        fmt.Printf("Peek error: %v\n", err)
    } else {
        fmt.Printf("After dequeue, peek: %d\n", val)
    }
}
```

**Output:**
```
Empty queue peek error: queue is empty
Peek 1: 100 (size: 3)
Peek 2: 100 (size: 3)
Peek 3: 100 (size: 3)
After dequeue, peek: 200
```

## Size, Capacity, IsEmpty

Query methods to get information about the queue's current state.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[int](8)
    
    fmt.Printf("New queue - Size: %d, Capacity: %d, Empty: %v\n", 
        q.Size(), q.Capacity(), q.IsEmpty())
    
    // Add some elements
    for i := 1; i <= 5; i++ {
        q.Enqueue(i * 10)
        fmt.Printf("After adding %d - Size: %d, Capacity: %d, Empty: %v\n", 
            i*10, q.Size(), q.Capacity(), q.IsEmpty())
    }
    
    // Remove some elements
    for i := 1; i <= 3; i++ {
        q.Dequeue()
        fmt.Printf("After dequeue %d - Size: %d, Capacity: %d, Empty: %v\n", 
            i, q.Size(), q.Capacity(), q.IsEmpty())
    }
    
    // Remove remaining elements
    for !q.IsEmpty() {
        q.Dequeue()
    }
    fmt.Printf("After emptying - Size: %d, Capacity: %d, Empty: %v\n", 
        q.Size(), q.Capacity(), q.IsEmpty())
}
```

**Output:**
```
New queue - Size: 0, Capacity: 8, Empty: true
After adding 10 - Size: 1, Capacity: 8, Empty: false
After adding 20 - Size: 2, Capacity: 8, Empty: false
After adding 30 - Size: 3, Capacity: 8, Empty: false
After adding 40 - Size: 4, Capacity: 8, Empty: false
After adding 50 - Size: 5, Capacity: 8, Empty: false
After dequeue 1 - Size: 4, Capacity: 8, Empty: false
After dequeue 2 - Size: 3, Capacity: 8, Empty: false
After dequeue 3 - Size: 2, Capacity: 8, Empty: false
After emptying - Size: 0, Capacity: 8, Empty: true
```

## Dynamic Resizing

The queue automatically grows when full and shrinks when utilization is low.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[int](4)
    
    fmt.Printf("Initial capacity: %d\n", q.Capacity())
    
    // Fill queue to trigger growth
    fmt.Println("\nFilling queue to trigger growth:")
    for i := 1; i <= 10; i++ {
        q.Enqueue(i)
        fmt.Printf("Added %d: size=%d, capacity=%d\n", i, q.Size(), q.Capacity())
    }
    
    // Add many more to see multiple growth cycles
    fmt.Println("\nAdding more elements:")
    for i := 11; i <= 20; i++ {
        q.Enqueue(i)
    }
    fmt.Printf("After adding 20 elements: size=%d, capacity=%d\n", q.Size(), q.Capacity())
    
    // Remove most elements to trigger shrinking
    fmt.Println("\nRemoving elements to trigger shrinking:")
    for i := 1; i <= 17; i++ {
        q.Dequeue()
        if i%5 == 0 || i == 17 { // Show every 5th removal and the last one
            fmt.Printf("After removing %d elements: size=%d, capacity=%d\n", 
                i, q.Size(), q.Capacity())
        }
    }
}
```

**Output:**
```
Initial capacity: 4

Filling queue to trigger growth:
Added 1: size=1, capacity=4
Added 2: size=2, capacity=4
Added 3: size=3, capacity=4
Added 4: size=4, capacity=4
Added 5: size=5, capacity=8
Added 6: size=6, capacity=8
Added 7: size=7, capacity=8
Added 8: size=8, capacity=8
Added 9: size=9, capacity=16
Added 10: size=10, capacity=16

Adding more elements:
After adding 20 elements: size=20, capacity=32

Removing elements to trigger shrinking:
After removing 5 elements: size=15, capacity=32
After removing 10 elements: size=10, capacity=32
After removing 15 elements: size=5, capacity=16
After removing 17 elements: size=3, capacity=16
```

## Different Types

The queue works with any type thanks to Go generics.

```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    // String queue
    fmt.Println("String queue:")
    strQueue := queue.NewQueue[string](4)
    words := []string{"hello", "world", "golang", "queue"}
    
    for _, word := range words {
        strQueue.Enqueue(word)
    }
    
    for !strQueue.IsEmpty() {
        word, _ := strQueue.Dequeue()
        fmt.Printf("String: %s\n", word)
    }
    
    // Struct queue
    fmt.Println("\nStruct queue:")
    type Person struct {
        Name string
        Age  int
    }
    
    personQueue := queue.NewQueue[Person](2)
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    
    for _, person := range people {
        personQueue.Enqueue(person)
    }
    
    for !personQueue.IsEmpty() {
        person, _ := personQueue.Dequeue()
        fmt.Printf("Person: %s (age %d)\n", person.Name, person.Age)
    }
    
    // Pointer queue
    fmt.Println("\nPointer queue:")
    ptrQueue := queue.NewQueue[*int](2)
    
    values := []int{10, 20, 30}
    for _, v := range values {
        val := v // Important: capture value in loop
        ptrQueue.Enqueue(&val)
    }
    
    for !ptrQueue.IsEmpty() {
        ptr, _ := ptrQueue.Dequeue()
        if ptr != nil {
            fmt.Printf("Pointer value: %d\n", *ptr)
        }
    }
}
```

**Output:**
```
String queue:
String: hello
String: world
String: golang
String: queue

Struct queue:
Person: Alice (age 30)
Person: Bob (age 25)
Person: Charlie (age 35)

Pointer queue:
Pointer value: 10
Pointer value: 20
Pointer value: 30
```

## Concurrent Usage

The queue is thread-safe and can be used safely from multiple goroutines.

```go
package main

import (
    "fmt"
    "errors"
    "sync"
    "time"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[int](4)
    var wg sync.WaitGroup
    
    // Producer goroutines
    fmt.Println("Starting concurrent producers and consumers...")
    
    // Start 3 producers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(producerID int) {
            defer wg.Done()
            for j := 1; j <= 5; j++ {
                value := producerID*100 + j
                q.Enqueue(value)
                fmt.Printf("Producer %d: enqueued %d\n", producerID, value)
                time.Sleep(50 * time.Millisecond)
            }
        }(i)
    }
    
    // Start 2 consumers
    for i := 1; i <= 2; i++ {
        wg.Add(1)
        go func(consumerID int) {
            defer wg.Done()
            consumed := 0
            for consumed < 7 { // Each consumer will try to consume 7 items
              value, err := q.Dequeue()
                if errors.Is(err, queue.ErrEmptyQueue) {
                    time.Sleep(10 * time.Millisecond) // Wait for items
                    continue
                } else if err != nil {
                    fmt.Printf("Unexpected error: %v\n", err)
                    continue
                }
                fmt.Printf("Consumer %d: dequeued %d\n", consumerID,value)
                consumed++
                time.Sleep(75 * time.Millisecond)
            }
        }(i)
    }
    
    // Monitor queue size
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 20; i++ {
            fmt.Printf("Queue size: %d\n", q.Size())
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    wg.Wait()
    fmt.Printf("Final queue size: %d\n", q.Size())
}
```

**Example output (abridged; order may vary):*
```
Starting concurrent producers and consumers...
Producer 1: enqueued 101
Queue size: 1
Producer 2: enqueued 201
Producer 3: enqueued 301
Consumer 1: dequeued 101
Queue size: 2
Producer 1: enqueued 102
Consumer 2: dequeued 201
Producer 2: enqueued 202
...
Final queue size: 1
```

## Error Handling

Proper error handling for empty queue operations.

```go
package main

import (
    "fmt"
    "errors"
    "github.com/kashifkhan0771/utils/queue"
)

func main() {
    q := queue.NewQueue[int](4)
    
    // Function to safely dequeue with error handling
    safeDequeue := func(q *queue.Queue[int]) {
        value, err := q.Dequeue()
        if err != nil {
            if errors.Is(err, queue.ErrEmptyQueue) {
                fmt.Println("Cannot dequeue: queue is empty")
            } else {
                fmt.Printf("Unexpected error: %v\n", err)
            }
            return
        }
        fmt.Printf("Successfully dequeued: %d\n", value)
    }
    
    // Function to safely peek with error handling
    safePeek := func(q *queue.Queue[int]) {
        value, err := q.Peek()
        if err != nil {
            if errors.Is(err, queue.ErrEmptyQueue) {
                fmt.Println("Cannot peek: queue is empty")
            } else {
                fmt.Printf("Unexpected error: %v\n", err)
            }
            return
        }
        fmt.Printf("Peeked value: %d\n", value)
    }
    
    // Test error cases
    fmt.Println("Testing empty queue operations:")
    safeDequeue(q)
    safePeek(q)
    
    // Add some items and test success cases
    fmt.Println("\nAdding items and testing success cases:")
    q.Enqueue(42)
    q.Enqueue(84)
    
    safePeek(q)
    safeDequeue(q)
    safePeek(q)
    safeDequeue(q)
    
    // Test empty again
    fmt.Println("\nTesting empty queue again:")
    safeDequeue(q)
    safePeek(q)
}
```

**Output:**
```
Testing empty queue operations:
Cannot dequeue: queue is empty
Cannot peek: queue is empty

Adding items and testing success cases:
Peeked value: 42
Successfully dequeued: 42
Peeked value: 84
Successfully dequeued: 84

Testing empty queue again:
Cannot dequeue: queue is empty
Cannot peek: queue is empty
```

## Complete Usage Example

Here's a comprehensive example showing a typical use case - implementing a work queue for task processing:

```go
package main

import (
    "fmt"
    "errors"
    "sync"
    "time"
    "github.com/kashifkhan0771/utils/queue"
)

type Task struct {
    ID       int
    Name     string
    Priority int
    Data     string
}

func (t Task) String() string {
    return fmt.Sprintf("Task{ID: %d, Name: %s, Priority: %d}", t.ID, t.Name, t.Priority)
}

func main() {
    // Create a work queue
    workQueue := queue.NewQueue[Task](10)
    var wg sync.WaitGroup
    
    fmt.Println("Starting task processing system...")
    
    // Task generator
    wg.Add(1)
    go func() {
        defer wg.Done()
        tasks := []Task{
            {1, "Email Processing", 1, "process emails"},
            {2, "Database Backup", 3, "backup database"},
            {3, "Report Generation", 2, "generate reports"},
            {4, "Cache Cleanup", 1, "clean cache"},
            {5, "Log Rotation", 2, "rotate logs"},
            {6, "Security Scan", 3, "run security scan"},
            {7, "Index Rebuild", 3, "rebuild search index"},
            {8, "Health Check", 1, "system health check"},
        }
        
        for _, task := range tasks {
            workQueue.Enqueue(task)
            fmt.Printf("📝 Queued: %s\n", task)
            time.Sleep(200 * time.Millisecond)
        }
        fmt.Println("✅ All tasks queued")
        close(done)
    }()
    // Worker processes
    numWorkers := 3
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            processed := 0
            workerLoop:
            for {
                task, err := workQueue.Dequeue()
                if err != nil {
                    if errors.Is(err, queue.ErrEmptyQueue) {
                        // If producer is done and queue drained, exit
                        select {
                        case <-done:
                            if workQueue.IsEmpty() {
                                break workerLoop
                            }
                        default:
                        }
                        time.Sleep(100 * time.Millisecond)
                        continue
                    }
                    // Unexpected error path
                    fmt.Printf("dequeue error: %v\n", err)
                    time.Sleep(50 * time.Millisecond)
                    continue
                }

                // Simulate task processing
                fmt.Printf("🔄 Worker %d processing: %s\n", workerID, task)
                processingTime := time.Duration(task.Priority*200) * time.Millisecond
                time.Sleep(processingTime)

                fmt.Printf(
                    "✅ Worker %d completed: %s (took %v)\n",
                    workerID, task, processingTime,
                )
                processed++
            }
            fmt.Printf("🛑 Worker %d finished (processed %d tasks)\n", workerID, processed)
        }(i)
    }
    
    // Queue monitor
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 15; i++ {
            size := workQueue.Size()
            capacity := workQueue.Capacity()
            fmt.Printf("📊 Queue status: %d/%d tasks pending\n", size, capacity)
            time.Sleep(500 * time.Millisecond)
        }
    }()
    
    wg.Wait()
    
    // Final status
    remaining := workQueue.Size()
    fmt.Printf("\n🏁 Processing complete. Remaining tasks: %d\n", remaining)
    
    if remaining > 0 {
        fmt.Println("Remaining tasks:")
        for !workQueue.IsEmpty() {
            task, _ := workQueue.Dequeue()
            fmt.Printf("  - %s\n", task)
        }
    }
}
```

**Output:**
```
Starting task processing system...
📝 Queued: Task{ID: 1, Name: Email Processing, Priority: 1}
📊 Queue status: 1/10 tasks pending
📝 Queued: Task{ID: 2, Name: Database Backup, Priority: 3}
🔄 Worker 1 processing: Task{ID: 1, Name: Email Processing, Priority: 1}
🔄 Worker 2 processing: Task{ID: 2, Name: Database Backup, Priority: 3}
📝 Queued: Task{ID: 3, Name: Report Generation, Priority: 2}
✅ Worker 1 completed: Task{ID: 1, Name: Email Processing, Priority: 1} (took 200ms)
🔄 Worker 1 processing: Task{ID: 3, Name: Report Generation, Priority: 2}
📊 Queue status: 2/10 tasks pending
...
🏁 Processing complete. Remaining tasks: 0
```

---
