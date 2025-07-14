# Stack

This package provides a generic stack data structure for Go, supporting any type.

## Features

- **Generic**: Works with any type (`Stack[T any]`).
- **Push**: Add an element to the top of the stack.
- **Pop**: Remove and return the top element (returns zero value and `false` if empty).
- **Peek**: View the top element without removing it.
- **PeekNthElement**: View the n-th element from the top without removing it.
- **IsEmpty**: Check if the stack is empty.
- **Size**: Get the number of elements in the stack.
