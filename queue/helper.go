package queue

// grow doubles capacity. Caller MUST hold q.mu.
func (q *Queue[T]) grow() {
	newCapacity := len(q.data) * 2
	if newCapacity == 0 {
		newCapacity = MinCapacity
	}
	newData := make([]T, newCapacity)

	// Copy elements in FIFO order from circular buffer
	// Start from head and wrap around using modulo
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.head+i)%len(q.data)]
	}

	q.data = newData
	q.head = 0
	q.tail = q.size
}

func (q *Queue[T]) shrink() {
	newCapacity := len(q.data) / 2

	if newCapacity < MinCapacity {
		newCapacity = MinCapacity
	}

	newData := make([]T, newCapacity)

	// Copy elements in FIFO order from circular buffer
	// Start from head and wrap around using modulo
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.head+i)%len(q.data)]
	}

	q.data = newData
	q.head = 0
	q.tail = q.size
}
