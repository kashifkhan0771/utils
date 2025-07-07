# Stack Test Examples

This file summarizes and shows the test code for the generic stack implementation.

---

## TestStackBasic

```go
func TestStackBasic(t *testing.T) {
    stack := New[int]()

    if !stack.IsEmpty() {
        t.Errorf("Expected stack to be empty")
    }

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    if stack.Size() != 3 {
        t.Errorf("Expected stack size 3, got %d", stack.Size())
    }

    val, ok := stack.Peek()
    if !ok || val != 30 {
        t.Errorf("Peek() = %v, %v; want 30, true", val, ok)
    }

    val, ok = stack.Pop()
    if !ok || val != 30 {
        t.Errorf("Pop() = %v, %v; want 30, true", val, ok)
    }

    val, ok = stack.Pop()
    if !ok || val != 20 {
        t.Errorf("Pop() = %v, %v; want 20, true", val, ok)
    }

    val, ok = stack.Pop()
    if !ok || val != 10 {
        t.Errorf("Pop() = %v, %v; want 10, true", val, ok)
    }

    if !stack.IsEmpty() {
        t.Errorf("Expected stack to be empty after popping all elements")
    }

    _, ok = stack.Pop()
    if ok {
        t.Errorf("Pop() on empty stack should return ok=false")
    }
}