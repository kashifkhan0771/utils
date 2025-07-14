package stack

import (
	"testing"
)

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

	val, ok = stack.PeekNthElement(1)
	if !ok || val != 20 {
		t.Errorf("Peek() = %v, %v; want 20, true", val, ok)
	}

	val, ok = stack.PeekNthElement(2)
	if !ok || val != 10 {
		t.Errorf("Peek() = %v, %v; want 10, true", val, ok)
	}

	val, ok = stack.Peek()
	if !ok || val != 30 {
		t.Errorf("Peek() = %v, %v; want 30, true", val, ok)
	}

	val, ok = stack.Peek()
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

func TestStackWithStrings(t *testing.T) {
	stack := New[string]()
	stack.Push("a")
	stack.Push("b")

	val, ok := stack.Pop()
	if !ok || val != "b" {
		t.Errorf("Pop() = %v, %v; want 'b', true", val, ok)
	}

	val, ok = stack.Peek()
	if !ok || val != "a" {
		t.Errorf("Peek(0) = %v, %v; want 'a', true", val, ok)
	}
}

func TestStackSize(t *testing.T) {
	stack := New[int]()
	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}
	stack.Push(1)
	stack.Push(2)
	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
	stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Expected size 1, got %d", stack.Size())
	}
}

func TestStackWithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	stack := New[Person]()
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}

	stack.Push(p1)
	stack.Push(p2)

	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}

	val, ok := stack.Pop()
	if !ok || val != p2 {
		t.Errorf("Pop() = %v, %v; want %v, true", val, ok, p2)
	}

	val, ok = stack.Peek()
	if !ok || val != p1 {
		t.Errorf("Peek(0) = %v, %v; want %v, true", val, ok, p1)
	}
}
